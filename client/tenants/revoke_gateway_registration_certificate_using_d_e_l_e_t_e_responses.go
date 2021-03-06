// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RevokeGatewayRegistrationCertificateUsingDELETEReader is a Reader for the RevokeGatewayRegistrationCertificateUsingDELETE structure.
type RevokeGatewayRegistrationCertificateUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeGatewayRegistrationCertificateUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRevokeGatewayRegistrationCertificateUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRevokeGatewayRegistrationCertificateUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRevokeGatewayRegistrationCertificateUsingDELETENotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRevokeGatewayRegistrationCertificateUsingDELETEOK creates a RevokeGatewayRegistrationCertificateUsingDELETEOK with default headers values
func NewRevokeGatewayRegistrationCertificateUsingDELETEOK() *RevokeGatewayRegistrationCertificateUsingDELETEOK {
	return &RevokeGatewayRegistrationCertificateUsingDELETEOK{}
}

/*RevokeGatewayRegistrationCertificateUsingDELETEOK handles this case with default header values.

Successfully revoked gateway registration certificate.
*/
type RevokeGatewayRegistrationCertificateUsingDELETEOK struct {
	Payload *RevokeGatewayRegistrationCertificateUsingDELETEOKBody
}

func (o *RevokeGatewayRegistrationCertificateUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/tenants/{tenantId}/gatewayRegistrations/clientCertificate/{fingerprint}][%d] revokeGatewayRegistrationCertificateUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *RevokeGatewayRegistrationCertificateUsingDELETEOK) GetPayload() *RevokeGatewayRegistrationCertificateUsingDELETEOKBody {
	return o.Payload
}

func (o *RevokeGatewayRegistrationCertificateUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RevokeGatewayRegistrationCertificateUsingDELETEOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeGatewayRegistrationCertificateUsingDELETEBadRequest creates a RevokeGatewayRegistrationCertificateUsingDELETEBadRequest with default headers values
func NewRevokeGatewayRegistrationCertificateUsingDELETEBadRequest() *RevokeGatewayRegistrationCertificateUsingDELETEBadRequest {
	return &RevokeGatewayRegistrationCertificateUsingDELETEBadRequest{}
}

/*RevokeGatewayRegistrationCertificateUsingDELETEBadRequest handles this case with default header values.

HTTP request is malformed.
*/
type RevokeGatewayRegistrationCertificateUsingDELETEBadRequest struct {
	Payload *RevokeGatewayRegistrationCertificateUsingDELETEBadRequestBody
}

func (o *RevokeGatewayRegistrationCertificateUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/tenants/{tenantId}/gatewayRegistrations/clientCertificate/{fingerprint}][%d] revokeGatewayRegistrationCertificateUsingDELETEBadRequest  %+v", 400, o.Payload)
}

func (o *RevokeGatewayRegistrationCertificateUsingDELETEBadRequest) GetPayload() *RevokeGatewayRegistrationCertificateUsingDELETEBadRequestBody {
	return o.Payload
}

func (o *RevokeGatewayRegistrationCertificateUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RevokeGatewayRegistrationCertificateUsingDELETEBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeGatewayRegistrationCertificateUsingDELETENotFound creates a RevokeGatewayRegistrationCertificateUsingDELETENotFound with default headers values
func NewRevokeGatewayRegistrationCertificateUsingDELETENotFound() *RevokeGatewayRegistrationCertificateUsingDELETENotFound {
	return &RevokeGatewayRegistrationCertificateUsingDELETENotFound{}
}

/*RevokeGatewayRegistrationCertificateUsingDELETENotFound handles this case with default header values.

Certificate does not exist.
*/
type RevokeGatewayRegistrationCertificateUsingDELETENotFound struct {
	Payload *RevokeGatewayRegistrationCertificateUsingDELETENotFoundBody
}

func (o *RevokeGatewayRegistrationCertificateUsingDELETENotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/tenants/{tenantId}/gatewayRegistrations/clientCertificate/{fingerprint}][%d] revokeGatewayRegistrationCertificateUsingDELETENotFound  %+v", 404, o.Payload)
}

func (o *RevokeGatewayRegistrationCertificateUsingDELETENotFound) GetPayload() *RevokeGatewayRegistrationCertificateUsingDELETENotFoundBody {
	return o.Payload
}

func (o *RevokeGatewayRegistrationCertificateUsingDELETENotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RevokeGatewayRegistrationCertificateUsingDELETENotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RevokeGatewayRegistrationCertificateUsingDELETEBadRequestBody Response
swagger:model RevokeGatewayRegistrationCertificateUsingDELETEBadRequestBody
*/
type RevokeGatewayRegistrationCertificateUsingDELETEBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this revoke gateway registration certificate using d e l e t e bad request body
func (o *RevokeGatewayRegistrationCertificateUsingDELETEBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RevokeGatewayRegistrationCertificateUsingDELETEBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("revokeGatewayRegistrationCertificateUsingDELETEBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RevokeGatewayRegistrationCertificateUsingDELETEBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RevokeGatewayRegistrationCertificateUsingDELETEBadRequestBody) UnmarshalBinary(b []byte) error {
	var res RevokeGatewayRegistrationCertificateUsingDELETEBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RevokeGatewayRegistrationCertificateUsingDELETENotFoundBody Response
swagger:model RevokeGatewayRegistrationCertificateUsingDELETENotFoundBody
*/
type RevokeGatewayRegistrationCertificateUsingDELETENotFoundBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this revoke gateway registration certificate using d e l e t e not found body
func (o *RevokeGatewayRegistrationCertificateUsingDELETENotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RevokeGatewayRegistrationCertificateUsingDELETENotFoundBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("revokeGatewayRegistrationCertificateUsingDELETENotFound"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RevokeGatewayRegistrationCertificateUsingDELETENotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RevokeGatewayRegistrationCertificateUsingDELETENotFoundBody) UnmarshalBinary(b []byte) error {
	var res RevokeGatewayRegistrationCertificateUsingDELETENotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RevokeGatewayRegistrationCertificateUsingDELETEOKBody Response
swagger:model RevokeGatewayRegistrationCertificateUsingDELETEOKBody
*/
type RevokeGatewayRegistrationCertificateUsingDELETEOKBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this revoke gateway registration certificate using d e l e t e o k body
func (o *RevokeGatewayRegistrationCertificateUsingDELETEOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RevokeGatewayRegistrationCertificateUsingDELETEOKBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("revokeGatewayRegistrationCertificateUsingDELETEOK"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RevokeGatewayRegistrationCertificateUsingDELETEOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RevokeGatewayRegistrationCertificateUsingDELETEOKBody) UnmarshalBinary(b []byte) error {
	var res RevokeGatewayRegistrationCertificateUsingDELETEOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
