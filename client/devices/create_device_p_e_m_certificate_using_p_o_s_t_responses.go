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

// CreateDevicePEMCertificateUsingPOSTReader is a Reader for the CreateDevicePEMCertificateUsingPOST structure.
type CreateDevicePEMCertificateUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDevicePEMCertificateUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDevicePEMCertificateUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDevicePEMCertificateUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateDevicePEMCertificateUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDevicePEMCertificateUsingPOSTOK creates a CreateDevicePEMCertificateUsingPOSTOK with default headers values
func NewCreateDevicePEMCertificateUsingPOSTOK() *CreateDevicePEMCertificateUsingPOSTOK {
	return &CreateDevicePEMCertificateUsingPOSTOK{}
}

/*CreateDevicePEMCertificateUsingPOSTOK handles this case with default header values.

Successfully created device certificate.
*/
type CreateDevicePEMCertificateUsingPOSTOK struct {
	Payload *CreateDevicePEMCertificateUsingPOSTOKBody
}

func (o *CreateDevicePEMCertificateUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /v1/tenant/{tenantId}/devices/{deviceId}/authentications/clientCertificate/pem][%d] createDevicePEMCertificateUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateDevicePEMCertificateUsingPOSTOK) GetPayload() *CreateDevicePEMCertificateUsingPOSTOKBody {
	return o.Payload
}

func (o *CreateDevicePEMCertificateUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateDevicePEMCertificateUsingPOSTOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDevicePEMCertificateUsingPOSTBadRequest creates a CreateDevicePEMCertificateUsingPOSTBadRequest with default headers values
func NewCreateDevicePEMCertificateUsingPOSTBadRequest() *CreateDevicePEMCertificateUsingPOSTBadRequest {
	return &CreateDevicePEMCertificateUsingPOSTBadRequest{}
}

/*CreateDevicePEMCertificateUsingPOSTBadRequest handles this case with default header values.

HTTP request is malformed.
*/
type CreateDevicePEMCertificateUsingPOSTBadRequest struct {
	Payload *CreateDevicePEMCertificateUsingPOSTBadRequestBody
}

func (o *CreateDevicePEMCertificateUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/tenant/{tenantId}/devices/{deviceId}/authentications/clientCertificate/pem][%d] createDevicePEMCertificateUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDevicePEMCertificateUsingPOSTBadRequest) GetPayload() *CreateDevicePEMCertificateUsingPOSTBadRequestBody {
	return o.Payload
}

func (o *CreateDevicePEMCertificateUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateDevicePEMCertificateUsingPOSTBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDevicePEMCertificateUsingPOSTNotFound creates a CreateDevicePEMCertificateUsingPOSTNotFound with default headers values
func NewCreateDevicePEMCertificateUsingPOSTNotFound() *CreateDevicePEMCertificateUsingPOSTNotFound {
	return &CreateDevicePEMCertificateUsingPOSTNotFound{}
}

/*CreateDevicePEMCertificateUsingPOSTNotFound handles this case with default header values.

Device with specified id does not exist.
*/
type CreateDevicePEMCertificateUsingPOSTNotFound struct {
	Payload *CreateDevicePEMCertificateUsingPOSTNotFoundBody
}

func (o *CreateDevicePEMCertificateUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/tenant/{tenantId}/devices/{deviceId}/authentications/clientCertificate/pem][%d] createDevicePEMCertificateUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *CreateDevicePEMCertificateUsingPOSTNotFound) GetPayload() *CreateDevicePEMCertificateUsingPOSTNotFoundBody {
	return o.Payload
}

func (o *CreateDevicePEMCertificateUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateDevicePEMCertificateUsingPOSTNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateDevicePEMCertificateUsingPOSTBadRequestBody Response
swagger:model CreateDevicePEMCertificateUsingPOSTBadRequestBody
*/
type CreateDevicePEMCertificateUsingPOSTBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this create device p e m certificate using p o s t bad request body
func (o *CreateDevicePEMCertificateUsingPOSTBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateDevicePEMCertificateUsingPOSTBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("createDevicePEMCertificateUsingPOSTBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateDevicePEMCertificateUsingPOSTBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateDevicePEMCertificateUsingPOSTBadRequestBody) UnmarshalBinary(b []byte) error {
	var res CreateDevicePEMCertificateUsingPOSTBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateDevicePEMCertificateUsingPOSTBody ClientCertificateAuthenticationBean
swagger:model CreateDevicePEMCertificateUsingPOSTBody
*/
type CreateDevicePEMCertificateUsingPOSTBody struct {

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

// Validate validates this create device p e m certificate using p o s t body
func (o *CreateDevicePEMCertificateUsingPOSTBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createDevicePEMCertificateUsingPOSTBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["clientCertificate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createDevicePEMCertificateUsingPOSTBodyTypeTypePropEnum = append(createDevicePEMCertificateUsingPOSTBodyTypeTypePropEnum, v)
	}
}

const (

	// CreateDevicePEMCertificateUsingPOSTBodyTypeClientCertificate captures enum value "clientCertificate"
	CreateDevicePEMCertificateUsingPOSTBodyTypeClientCertificate string = "clientCertificate"
)

// prop value enum
func (o *CreateDevicePEMCertificateUsingPOSTBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createDevicePEMCertificateUsingPOSTBodyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateDevicePEMCertificateUsingPOSTBody) validateType(formats strfmt.Registry) error {

	if err := validate.Required("request"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("request"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateDevicePEMCertificateUsingPOSTBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateDevicePEMCertificateUsingPOSTBody) UnmarshalBinary(b []byte) error {
	var res CreateDevicePEMCertificateUsingPOSTBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateDevicePEMCertificateUsingPOSTNotFoundBody Response
swagger:model CreateDevicePEMCertificateUsingPOSTNotFoundBody
*/
type CreateDevicePEMCertificateUsingPOSTNotFoundBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this create device p e m certificate using p o s t not found body
func (o *CreateDevicePEMCertificateUsingPOSTNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateDevicePEMCertificateUsingPOSTNotFoundBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("createDevicePEMCertificateUsingPOSTNotFound"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateDevicePEMCertificateUsingPOSTNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateDevicePEMCertificateUsingPOSTNotFoundBody) UnmarshalBinary(b []byte) error {
	var res CreateDevicePEMCertificateUsingPOSTNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateDevicePEMCertificateUsingPOSTOKBody ClientCertificateAuthenticationBean
swagger:model CreateDevicePEMCertificateUsingPOSTOKBody
*/
type CreateDevicePEMCertificateUsingPOSTOKBody struct {

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

// Validate validates this create device p e m certificate using p o s t o k body
func (o *CreateDevicePEMCertificateUsingPOSTOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createDevicePEMCertificateUsingPOSTOKBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["clientCertificate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createDevicePEMCertificateUsingPOSTOKBodyTypeTypePropEnum = append(createDevicePEMCertificateUsingPOSTOKBodyTypeTypePropEnum, v)
	}
}

const (

	// CreateDevicePEMCertificateUsingPOSTOKBodyTypeClientCertificate captures enum value "clientCertificate"
	CreateDevicePEMCertificateUsingPOSTOKBodyTypeClientCertificate string = "clientCertificate"
)

// prop value enum
func (o *CreateDevicePEMCertificateUsingPOSTOKBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createDevicePEMCertificateUsingPOSTOKBodyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateDevicePEMCertificateUsingPOSTOKBody) validateType(formats strfmt.Registry) error {

	if err := validate.Required("createDevicePEMCertificateUsingPOSTOK"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("createDevicePEMCertificateUsingPOSTOK"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateDevicePEMCertificateUsingPOSTOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateDevicePEMCertificateUsingPOSTOKBody) UnmarshalBinary(b []byte) error {
	var res CreateDevicePEMCertificateUsingPOSTOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}