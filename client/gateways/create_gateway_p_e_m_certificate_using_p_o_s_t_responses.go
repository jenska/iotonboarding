// Code generated by go-swagger; DO NOT EDIT.

package gateways

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

// CreateGatewayPEMCertificateUsingPOSTReader is a Reader for the CreateGatewayPEMCertificateUsingPOST structure.
type CreateGatewayPEMCertificateUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGatewayPEMCertificateUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateGatewayPEMCertificateUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateGatewayPEMCertificateUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateGatewayPEMCertificateUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateGatewayPEMCertificateUsingPOSTOK creates a CreateGatewayPEMCertificateUsingPOSTOK with default headers values
func NewCreateGatewayPEMCertificateUsingPOSTOK() *CreateGatewayPEMCertificateUsingPOSTOK {
	return &CreateGatewayPEMCertificateUsingPOSTOK{}
}

/*CreateGatewayPEMCertificateUsingPOSTOK handles this case with default header values.

Successfully created gateway certificate.
*/
type CreateGatewayPEMCertificateUsingPOSTOK struct {
	Payload *CreateGatewayPEMCertificateUsingPOSTOKBody
}

func (o *CreateGatewayPEMCertificateUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /v1/tenant/{tenantId}/gateways/{gatewayId}/authentications/clientCertificate/pem][%d] createGatewayPEMCertificateUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateGatewayPEMCertificateUsingPOSTOK) GetPayload() *CreateGatewayPEMCertificateUsingPOSTOKBody {
	return o.Payload
}

func (o *CreateGatewayPEMCertificateUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateGatewayPEMCertificateUsingPOSTOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGatewayPEMCertificateUsingPOSTBadRequest creates a CreateGatewayPEMCertificateUsingPOSTBadRequest with default headers values
func NewCreateGatewayPEMCertificateUsingPOSTBadRequest() *CreateGatewayPEMCertificateUsingPOSTBadRequest {
	return &CreateGatewayPEMCertificateUsingPOSTBadRequest{}
}

/*CreateGatewayPEMCertificateUsingPOSTBadRequest handles this case with default header values.

HTTP request is malformed.
*/
type CreateGatewayPEMCertificateUsingPOSTBadRequest struct {
	Payload *CreateGatewayPEMCertificateUsingPOSTBadRequestBody
}

func (o *CreateGatewayPEMCertificateUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/tenant/{tenantId}/gateways/{gatewayId}/authentications/clientCertificate/pem][%d] createGatewayPEMCertificateUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *CreateGatewayPEMCertificateUsingPOSTBadRequest) GetPayload() *CreateGatewayPEMCertificateUsingPOSTBadRequestBody {
	return o.Payload
}

func (o *CreateGatewayPEMCertificateUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateGatewayPEMCertificateUsingPOSTBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGatewayPEMCertificateUsingPOSTNotFound creates a CreateGatewayPEMCertificateUsingPOSTNotFound with default headers values
func NewCreateGatewayPEMCertificateUsingPOSTNotFound() *CreateGatewayPEMCertificateUsingPOSTNotFound {
	return &CreateGatewayPEMCertificateUsingPOSTNotFound{}
}

/*CreateGatewayPEMCertificateUsingPOSTNotFound handles this case with default header values.

Gateway with specified id does not exist.
*/
type CreateGatewayPEMCertificateUsingPOSTNotFound struct {
	Payload *CreateGatewayPEMCertificateUsingPOSTNotFoundBody
}

func (o *CreateGatewayPEMCertificateUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/tenant/{tenantId}/gateways/{gatewayId}/authentications/clientCertificate/pem][%d] createGatewayPEMCertificateUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *CreateGatewayPEMCertificateUsingPOSTNotFound) GetPayload() *CreateGatewayPEMCertificateUsingPOSTNotFoundBody {
	return o.Payload
}

func (o *CreateGatewayPEMCertificateUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateGatewayPEMCertificateUsingPOSTNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateGatewayPEMCertificateUsingPOSTBadRequestBody Response
swagger:model CreateGatewayPEMCertificateUsingPOSTBadRequestBody
*/
type CreateGatewayPEMCertificateUsingPOSTBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this create gateway p e m certificate using p o s t bad request body
func (o *CreateGatewayPEMCertificateUsingPOSTBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateGatewayPEMCertificateUsingPOSTBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("createGatewayPEMCertificateUsingPOSTBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateGatewayPEMCertificateUsingPOSTBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateGatewayPEMCertificateUsingPOSTBadRequestBody) UnmarshalBinary(b []byte) error {
	var res CreateGatewayPEMCertificateUsingPOSTBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateGatewayPEMCertificateUsingPOSTBody ClientCertificateAuthenticationBean
swagger:model CreateGatewayPEMCertificateUsingPOSTBody
*/
type CreateGatewayPEMCertificateUsingPOSTBody struct {

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

// Validate validates this create gateway p e m certificate using p o s t body
func (o *CreateGatewayPEMCertificateUsingPOSTBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createGatewayPEMCertificateUsingPOSTBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["clientCertificate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createGatewayPEMCertificateUsingPOSTBodyTypeTypePropEnum = append(createGatewayPEMCertificateUsingPOSTBodyTypeTypePropEnum, v)
	}
}

const (

	// CreateGatewayPEMCertificateUsingPOSTBodyTypeClientCertificate captures enum value "clientCertificate"
	CreateGatewayPEMCertificateUsingPOSTBodyTypeClientCertificate string = "clientCertificate"
)

// prop value enum
func (o *CreateGatewayPEMCertificateUsingPOSTBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createGatewayPEMCertificateUsingPOSTBodyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateGatewayPEMCertificateUsingPOSTBody) validateType(formats strfmt.Registry) error {

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
func (o *CreateGatewayPEMCertificateUsingPOSTBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateGatewayPEMCertificateUsingPOSTBody) UnmarshalBinary(b []byte) error {
	var res CreateGatewayPEMCertificateUsingPOSTBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateGatewayPEMCertificateUsingPOSTNotFoundBody Response
swagger:model CreateGatewayPEMCertificateUsingPOSTNotFoundBody
*/
type CreateGatewayPEMCertificateUsingPOSTNotFoundBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this create gateway p e m certificate using p o s t not found body
func (o *CreateGatewayPEMCertificateUsingPOSTNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateGatewayPEMCertificateUsingPOSTNotFoundBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("createGatewayPEMCertificateUsingPOSTNotFound"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateGatewayPEMCertificateUsingPOSTNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateGatewayPEMCertificateUsingPOSTNotFoundBody) UnmarshalBinary(b []byte) error {
	var res CreateGatewayPEMCertificateUsingPOSTNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateGatewayPEMCertificateUsingPOSTOKBody ClientCertificateAuthenticationBean
swagger:model CreateGatewayPEMCertificateUsingPOSTOKBody
*/
type CreateGatewayPEMCertificateUsingPOSTOKBody struct {

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

// Validate validates this create gateway p e m certificate using p o s t o k body
func (o *CreateGatewayPEMCertificateUsingPOSTOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createGatewayPEMCertificateUsingPOSTOKBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["clientCertificate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createGatewayPEMCertificateUsingPOSTOKBodyTypeTypePropEnum = append(createGatewayPEMCertificateUsingPOSTOKBodyTypeTypePropEnum, v)
	}
}

const (

	// CreateGatewayPEMCertificateUsingPOSTOKBodyTypeClientCertificate captures enum value "clientCertificate"
	CreateGatewayPEMCertificateUsingPOSTOKBodyTypeClientCertificate string = "clientCertificate"
)

// prop value enum
func (o *CreateGatewayPEMCertificateUsingPOSTOKBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createGatewayPEMCertificateUsingPOSTOKBodyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateGatewayPEMCertificateUsingPOSTOKBody) validateType(formats strfmt.Registry) error {

	if err := validate.Required("createGatewayPEMCertificateUsingPOSTOK"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("createGatewayPEMCertificateUsingPOSTOK"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateGatewayPEMCertificateUsingPOSTOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateGatewayPEMCertificateUsingPOSTOKBody) UnmarshalBinary(b []byte) error {
	var res CreateGatewayPEMCertificateUsingPOSTOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
