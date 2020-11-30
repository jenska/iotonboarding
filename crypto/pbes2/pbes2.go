package pbes2

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/asn1"
	"errors"
	"hash"
	"strconv"

	"golang.org/x/crypto/pbkdf2"
)

const maxIter = 100000

var (
	oidRSADSI              = asn1.ObjectIdentifier{1, 2, 840, 113549}
	oidPKCS5               = appendOID(oidRSADSI, 1, 5)
	oidPBKDF2              = appendOID(oidPKCS5, 12)
	oidPBES2               = appendOID(oidPKCS5, 13)
	oidDigestAlgorithm     = appendOID(oidRSADSI, 2)
	oidHMACWithSHA1        = appendOID(oidDigestAlgorithm, 7)
	oidHMACWithSHA224      = appendOID(oidDigestAlgorithm, 8)
	oidHMACWithSHA256      = appendOID(oidDigestAlgorithm, 9)
	oidHMACWithSHA384      = appendOID(oidDigestAlgorithm, 10)
	oidHMACWithSHA512      = appendOID(oidDigestAlgorithm, 11)
	oidHMACWithSHA512_224  = appendOID(oidDigestAlgorithm, 12)
	oidHMACWithSHA512_256  = appendOID(oidDigestAlgorithm, 13)
	oidEncryptionAlgorithm = appendOID(oidRSADSI, 3)
	oidDESCBC              = asn1.ObjectIdentifier{1, 3, 14, 3, 2, 7}
	oidDESEDE3CBC          = appendOID(oidEncryptionAlgorithm, 7)
	oidAES                 = asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 1}
	oidAES128CBCPAD        = appendOID(oidAES, 2)
	oidAES192CBCPAD        = appendOID(oidAES, 22)
	oidAES256CBCPAD        = appendOID(oidAES, 42)
)

var (
	errIncorrectPassword = errors.New("possilbly incorrect encryption password")
	errInvalidDataLen    = errors.New("data size is not a multiple of cipher block size")
	errInvalidIVLen      = errors.New("invalid IV size")
	errInvalidIter       = errors.New("invalid iteration count")
	errNoData            = errors.New("no data to decrypt")
	errNotPBES2          = errors.New("not PBES2 data")
	errUnsupportedKDF    = errors.New("unsupported KDF")
	errUnsupportedPRF    = errors.New("unsupported PRF")
	errUnsupportedEncS   = errors.New("unsupported encryption scheme")
)

type errTooManyIterations int

func (err errTooManyIterations) Error() string {
	return "too many PBKDF2 iterations: " + strconv.Itoa(int(err))
}

func appendOID(b asn1.ObjectIdentifier, v ...int) asn1.ObjectIdentifier {
	n := make(asn1.ObjectIdentifier, len(b), len(b)+len(v))
	copy(n, b)
	return append(n, v...)
}

func prfByOID(oid asn1.ObjectIdentifier) func() hash.Hash {
	if len(oid) == 0 {
		return sha1.New
	}
	if oid.Equal(oidHMACWithSHA1) {
		return sha1.New
	}
	if oid.Equal(oidHMACWithSHA224) {
		return sha256.New224
	}
	if oid.Equal(oidHMACWithSHA256) {
		return sha256.New
	}
	if oid.Equal(oidHMACWithSHA384) {
		return sha512.New384
	}
	if oid.Equal(oidHMACWithSHA512) {
		return sha512.New
	}
	if oid.Equal(oidHMACWithSHA512_224) {
		return sha512.New512_224
	}
	if oid.Equal(oidHMACWithSHA512_256) {
		return sha512.New512_256
	}
	return nil
}

func encsByOID(oid asn1.ObjectIdentifier) (func([]byte) (cipher.Block, error), func(cipher.Block, []byte) cipher.BlockMode, int) {
	if oid.Equal(oidDESCBC) {
		return des.NewCipher, cipher.NewCBCDecrypter, 8
	}
	if oid.Equal(oidDESEDE3CBC) {
		return des.NewTripleDESCipher, cipher.NewCBCDecrypter, 24
	}
	if oid.Equal(oidAES128CBCPAD) {
		return aes.NewCipher, cipher.NewCBCDecrypter, 16
	}
	if oid.Equal(oidAES192CBCPAD) {
		return aes.NewCipher, cipher.NewCBCDecrypter, 24
	}
	if oid.Equal(oidAES256CBCPAD) {
		return aes.NewCipher, cipher.NewCBCDecrypter, 32
	}
	return nil, nil, 0
}

func DecryptPBES2(b, password []byte) (data, rest []byte, err error) {
	var p struct {
		ES struct {
			ID     asn1.ObjectIdentifier
			Params struct {
				KDF struct {
					ID     asn1.ObjectIdentifier
					Params struct {
						Salt      []byte
						Iter      int
						KeyLength int `asn1:"optional"`
						PRF       struct {
							ID     asn1.ObjectIdentifier
							Params asn1.RawValue
						} `asn1:"optional"`
					}
				}
				EncS struct {
					ID     asn1.ObjectIdentifier
					Params []byte
				}
			}
		}
		Data []byte
	}
	rest, err = asn1.Unmarshal(b, &p)
	if err != nil {
		return
	}
	if !p.ES.ID.Equal(oidPBES2) {
		err = errNotPBES2
		return
	}
	if !p.ES.Params.KDF.ID.Equal(oidPBKDF2) {
		err = errUnsupportedKDF
		return
	}
	if p.ES.Params.KDF.Params.Iter < 1 {
		err = errInvalidIter
		return
	}
	prf := prfByOID(p.ES.Params.KDF.Params.PRF.ID)
	if prf == nil {
		err = errUnsupportedPRF
		return
	}
	bcf, bmf, kl := encsByOID(p.ES.Params.EncS.ID)
	if bcf == nil || bmf == nil {
		err = errUnsupportedEncS
		return
	}
	if len(p.Data) == 0 {
		err = errNoData
		return
	}
	if maxIter > 0 && p.ES.Params.KDF.Params.Iter > maxIter {
		err = errTooManyIterations(p.ES.Params.KDF.Params.Iter)
		return
	}
	key := pbkdf2.Key(password, p.ES.Params.KDF.Params.Salt, p.ES.Params.KDF.Params.Iter, kl, prf)
	var bc cipher.Block
	bc, err = bcf(key)
	if err != nil {
		return
	}
	if len(p.ES.Params.EncS.Params) != bc.BlockSize() {
		err = errInvalidIVLen
		return
	}
	bm := bmf(bc, p.ES.Params.EncS.Params)
	if len(p.Data)%bm.BlockSize() != 0 {
		err = errInvalidDataLen
		return
	}
	data = make([]byte, len(p.Data))
	bm.CryptBlocks(data, p.Data)
	pl := data[len(data)-1]
	if pl == 0 || int(pl) > bm.BlockSize() {
		err = errIncorrectPassword
		return
	}
	dl := len(data) - int(pl)
	for _, b := range data[dl:] {
		if b != pl {
			err = errIncorrectPassword
			return
		}
	}
	data = data[:dl]
	return
}
