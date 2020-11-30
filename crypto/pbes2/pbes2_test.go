package pbes2

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"testing"

	"golang.org/x/crypto/ssh"
)

const testPem = "./../../testdata/simple_vacuum_pump_iot_device_simulator-device_certificate.pem"
const secrect = "2rYQqXsHZDwEQnf5rlcp7TckoZsbBQKvRPD3"

func TestDecryptPBES2(t *testing.T) {
	binKey, err := ioutil.ReadFile(testPem)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ssh.ParsePrivateKey(binKey)
	if err == nil {
		t.Fatal("Encrypted private key is assumed to be unsupported by golang")
	}

	block, _ := pem.Decode(binKey)
	if block == nil {
		t.Fatal("failed to decode PEM block")
	}

	var derKey []byte
	derKey, _, err = DecryptPBES2(block.Bytes, []byte(secrect))
	if err != nil {
		t.Fatal(err)
	}

	key, err := x509.ParsePKCS8PrivateKey(derKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("key type: %T\n", key)
}
