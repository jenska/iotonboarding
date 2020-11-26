// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetPEMCertificateUsingGETReader is a Reader for the GetPEMCertificateUsingGET structure.
type GetPEMCertificateUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPEMCertificateUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPEMCertificateUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPEMCertificateUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPEMCertificateUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPEMCertificateUsingGETOK creates a GetPEMCertificateUsingGETOK with default headers values
func NewGetPEMCertificateUsingGETOK() *GetPEMCertificateUsingGETOK {
	return &GetPEMCertificateUsingGETOK{}
}

/*GetPEMCertificateUsingGETOK handles this case with default header values.

Successfully downloaded device private key and certificate.
*/
type GetPEMCertificateUsingGETOK struct {
	Payload *GetPEMCertificateUsingGETOKBody
}

func (o *GetPEMCertificateUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /v1/tenant/{tenantId}/devices/{deviceId}/authentications/clientCertificate/pem][%d] getPEMCertificateUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetPEMCertificateUsingGETOK) GetPayload() *GetPEMCertificateUsingGETOKBody {
	return o.Payload
}

func (o *GetPEMCertificateUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPEMCertificateUsingGETOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPEMCertificateUsingGETBadRequest creates a GetPEMCertificateUsingGETBadRequest with default headers values
func NewGetPEMCertificateUsingGETBadRequest() *GetPEMCertificateUsingGETBadRequest {
	return &GetPEMCertificateUsingGETBadRequest{}
}

/*GetPEMCertificateUsingGETBadRequest handles this case with default header values.

HTTP request is malformed.
*/
type GetPEMCertificateUsingGETBadRequest struct {
	Payload *GetPEMCertificateUsingGETBadRequestBody
}

func (o *GetPEMCertificateUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/tenant/{tenantId}/devices/{deviceId}/authentications/clientCertificate/pem][%d] getPEMCertificateUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetPEMCertificateUsingGETBadRequest) GetPayload() *GetPEMCertificateUsingGETBadRequestBody {
	return o.Payload
}

func (o *GetPEMCertificateUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPEMCertificateUsingGETBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPEMCertificateUsingGETNotFound creates a GetPEMCertificateUsingGETNotFound with default headers values
func NewGetPEMCertificateUsingGETNotFound() *GetPEMCertificateUsingGETNotFound {
	return &GetPEMCertificateUsingGETNotFound{}
}

/*GetPEMCertificateUsingGETNotFound handles this case with default header values.

Device with specified id does not exist.
*/
type GetPEMCertificateUsingGETNotFound struct {
	Payload *GetPEMCertificateUsingGETNotFoundBody
}

func (o *GetPEMCertificateUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/tenant/{tenantId}/devices/{deviceId}/authentications/clientCertificate/pem][%d] getPEMCertificateUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetPEMCertificateUsingGETNotFound) GetPayload() *GetPEMCertificateUsingGETNotFoundBody {
	return o.Payload
}

func (o *GetPEMCertificateUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPEMCertificateUsingGETNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetPEMCertificateUsingGETBadRequestBody Response
swagger:model GetPEMCertificateUsingGETBadRequestBody
*/
type GetPEMCertificateUsingGETBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this get p e m certificate using g e t bad request body
func (o *GetPEMCertificateUsingGETBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPEMCertificateUsingGETBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getPEMCertificateUsingGETBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPEMCertificateUsingGETBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPEMCertificateUsingGETBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetPEMCertificateUsingGETBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPEMCertificateUsingGETNotFoundBody Response
swagger:model GetPEMCertificateUsingGETNotFoundBody
*/
type GetPEMCertificateUsingGETNotFoundBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this get p e m certificate using g e t not found body
func (o *GetPEMCertificateUsingGETNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPEMCertificateUsingGETNotFoundBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getPEMCertificateUsingGETNotFound"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPEMCertificateUsingGETNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPEMCertificateUsingGETNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetPEMCertificateUsingGETNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPEMCertificateUsingGETOKBody ClientCertificateAuthenticationBean
swagger:model GetPEMCertificateUsingGETOKBody
*/
type GetPEMCertificateUsingGETOKBody struct {

	// csr
	Csr string `json:"csr,omitempty"`

	// p12
	// Read Only: true
	P12 string `json:"p12,omitempty"`

	// pem
	// Read Only: true
	Pem string `json:"pem,omitempty"`

	// secret
	// Read Only: true
	Secret string `json:"secret,omitempty"`

	// The type of the authentication mechanism
	// Required: true
	// Enum: [clientCertificate]
	Type *string `json:"type"`
}

// Validate validates this get p e m certificate using g e t o k body
func (o *GetPEMCertificateUsingGETOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getPEMCertificateUsingGETOKBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["clientCertificate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getPEMCertificateUsingGETOKBodyTypeTypePropEnum = append(getPEMCertificateUsingGETOKBodyTypeTypePropEnum, v)
	}
}

const (

	// GetPEMCertificateUsingGETOKBodyTypeClientCertificate captures enum value "clientCertificate"
	GetPEMCertificateUsingGETOKBodyTypeClientCertificate string = "clientCertificate"
)

// prop value enum
func (o *GetPEMCertificateUsingGETOKBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getPEMCertificateUsingGETOKBodyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetPEMCertificateUsingGETOKBody) validateType(formats strfmt.Registry) error {

	if err := validate.Required("getPEMCertificateUsingGETOK"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("getPEMCertificateUsingGETOK"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPEMCertificateUsingGETOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPEMCertificateUsingGETOKBody) UnmarshalBinary(b []byte) error {
	var res GetPEMCertificateUsingGETOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}