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

// DeleteCustomPropertyForTenantReader is a Reader for the DeleteCustomPropertyForTenant structure.
type DeleteCustomPropertyForTenantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomPropertyForTenantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCustomPropertyForTenantOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCustomPropertyForTenantBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCustomPropertyForTenantNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCustomPropertyForTenantOK creates a DeleteCustomPropertyForTenantOK with default headers values
func NewDeleteCustomPropertyForTenantOK() *DeleteCustomPropertyForTenantOK {
	return &DeleteCustomPropertyForTenantOK{}
}

/*DeleteCustomPropertyForTenantOK handles this case with default header values.

Successfully deleted custom property.
*/
type DeleteCustomPropertyForTenantOK struct {
	Payload *DeleteCustomPropertyForTenantOKBody
}

func (o *DeleteCustomPropertyForTenantOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/tenants/{tenantId}/customProperties/{key}][%d] deleteCustomPropertyForTenantOK  %+v", 200, o.Payload)
}

func (o *DeleteCustomPropertyForTenantOK) GetPayload() *DeleteCustomPropertyForTenantOKBody {
	return o.Payload
}

func (o *DeleteCustomPropertyForTenantOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteCustomPropertyForTenantOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomPropertyForTenantBadRequest creates a DeleteCustomPropertyForTenantBadRequest with default headers values
func NewDeleteCustomPropertyForTenantBadRequest() *DeleteCustomPropertyForTenantBadRequest {
	return &DeleteCustomPropertyForTenantBadRequest{}
}

/*DeleteCustomPropertyForTenantBadRequest handles this case with default header values.

HTTP request is malformed.
*/
type DeleteCustomPropertyForTenantBadRequest struct {
	Payload *DeleteCustomPropertyForTenantBadRequestBody
}

func (o *DeleteCustomPropertyForTenantBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/tenants/{tenantId}/customProperties/{key}][%d] deleteCustomPropertyForTenantBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCustomPropertyForTenantBadRequest) GetPayload() *DeleteCustomPropertyForTenantBadRequestBody {
	return o.Payload
}

func (o *DeleteCustomPropertyForTenantBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteCustomPropertyForTenantBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomPropertyForTenantNotFound creates a DeleteCustomPropertyForTenantNotFound with default headers values
func NewDeleteCustomPropertyForTenantNotFound() *DeleteCustomPropertyForTenantNotFound {
	return &DeleteCustomPropertyForTenantNotFound{}
}

/*DeleteCustomPropertyForTenantNotFound handles this case with default header values.

Tenant or custom property with specified id or key does not exist.
*/
type DeleteCustomPropertyForTenantNotFound struct {
	Payload *DeleteCustomPropertyForTenantNotFoundBody
}

func (o *DeleteCustomPropertyForTenantNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/tenants/{tenantId}/customProperties/{key}][%d] deleteCustomPropertyForTenantNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCustomPropertyForTenantNotFound) GetPayload() *DeleteCustomPropertyForTenantNotFoundBody {
	return o.Payload
}

func (o *DeleteCustomPropertyForTenantNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteCustomPropertyForTenantNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteCustomPropertyForTenantBadRequestBody Response
swagger:model DeleteCustomPropertyForTenantBadRequestBody
*/
type DeleteCustomPropertyForTenantBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this delete custom property for tenant bad request body
func (o *DeleteCustomPropertyForTenantBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteCustomPropertyForTenantBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("deleteCustomPropertyForTenantBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteCustomPropertyForTenantBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteCustomPropertyForTenantBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DeleteCustomPropertyForTenantBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteCustomPropertyForTenantNotFoundBody Response
swagger:model DeleteCustomPropertyForTenantNotFoundBody
*/
type DeleteCustomPropertyForTenantNotFoundBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this delete custom property for tenant not found body
func (o *DeleteCustomPropertyForTenantNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteCustomPropertyForTenantNotFoundBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("deleteCustomPropertyForTenantNotFound"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteCustomPropertyForTenantNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteCustomPropertyForTenantNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteCustomPropertyForTenantNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteCustomPropertyForTenantOKBody Response
swagger:model DeleteCustomPropertyForTenantOKBody
*/
type DeleteCustomPropertyForTenantOKBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this delete custom property for tenant o k body
func (o *DeleteCustomPropertyForTenantOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteCustomPropertyForTenantOKBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("deleteCustomPropertyForTenantOK"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteCustomPropertyForTenantOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteCustomPropertyForTenantOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteCustomPropertyForTenantOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}