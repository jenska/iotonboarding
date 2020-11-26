// Code generated by go-swagger; DO NOT EDIT.

package gateways

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

// UpdateGatewayConfigurationUsingPUTReader is a Reader for the UpdateGatewayConfigurationUsingPUT structure.
type UpdateGatewayConfigurationUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGatewayConfigurationUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateGatewayConfigurationUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateGatewayConfigurationUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateGatewayConfigurationUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateGatewayConfigurationUsingPUTOK creates a UpdateGatewayConfigurationUsingPUTOK with default headers values
func NewUpdateGatewayConfigurationUsingPUTOK() *UpdateGatewayConfigurationUsingPUTOK {
	return &UpdateGatewayConfigurationUsingPUTOK{}
}

/*UpdateGatewayConfigurationUsingPUTOK handles this case with default header values.

Successfully updated gateway configuration.
*/
type UpdateGatewayConfigurationUsingPUTOK struct {
	Payload string
}

func (o *UpdateGatewayConfigurationUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /v1/tenant/{tenantId}/gateways/{gatewayId}/configuration][%d] updateGatewayConfigurationUsingPUTOK  %+v", 200, o.Payload)
}

func (o *UpdateGatewayConfigurationUsingPUTOK) GetPayload() string {
	return o.Payload
}

func (o *UpdateGatewayConfigurationUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGatewayConfigurationUsingPUTBadRequest creates a UpdateGatewayConfigurationUsingPUTBadRequest with default headers values
func NewUpdateGatewayConfigurationUsingPUTBadRequest() *UpdateGatewayConfigurationUsingPUTBadRequest {
	return &UpdateGatewayConfigurationUsingPUTBadRequest{}
}

/*UpdateGatewayConfigurationUsingPUTBadRequest handles this case with default header values.

HTTP request is malformed.
*/
type UpdateGatewayConfigurationUsingPUTBadRequest struct {
	Payload *UpdateGatewayConfigurationUsingPUTBadRequestBody
}

func (o *UpdateGatewayConfigurationUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/tenant/{tenantId}/gateways/{gatewayId}/configuration][%d] updateGatewayConfigurationUsingPUTBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateGatewayConfigurationUsingPUTBadRequest) GetPayload() *UpdateGatewayConfigurationUsingPUTBadRequestBody {
	return o.Payload
}

func (o *UpdateGatewayConfigurationUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateGatewayConfigurationUsingPUTBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGatewayConfigurationUsingPUTNotFound creates a UpdateGatewayConfigurationUsingPUTNotFound with default headers values
func NewUpdateGatewayConfigurationUsingPUTNotFound() *UpdateGatewayConfigurationUsingPUTNotFound {
	return &UpdateGatewayConfigurationUsingPUTNotFound{}
}

/*UpdateGatewayConfigurationUsingPUTNotFound handles this case with default header values.

Gateway with specified id does not exist.
*/
type UpdateGatewayConfigurationUsingPUTNotFound struct {
	Payload *UpdateGatewayConfigurationUsingPUTNotFoundBody
}

func (o *UpdateGatewayConfigurationUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/tenant/{tenantId}/gateways/{gatewayId}/configuration][%d] updateGatewayConfigurationUsingPUTNotFound  %+v", 404, o.Payload)
}

func (o *UpdateGatewayConfigurationUsingPUTNotFound) GetPayload() *UpdateGatewayConfigurationUsingPUTNotFoundBody {
	return o.Payload
}

func (o *UpdateGatewayConfigurationUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateGatewayConfigurationUsingPUTNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateGatewayConfigurationUsingPUTBadRequestBody Response
swagger:model UpdateGatewayConfigurationUsingPUTBadRequestBody
*/
type UpdateGatewayConfigurationUsingPUTBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this update gateway configuration using p u t bad request body
func (o *UpdateGatewayConfigurationUsingPUTBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateGatewayConfigurationUsingPUTBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("updateGatewayConfigurationUsingPUTBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateGatewayConfigurationUsingPUTBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateGatewayConfigurationUsingPUTBadRequestBody) UnmarshalBinary(b []byte) error {
	var res UpdateGatewayConfigurationUsingPUTBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateGatewayConfigurationUsingPUTNotFoundBody Response
swagger:model UpdateGatewayConfigurationUsingPUTNotFoundBody
*/
type UpdateGatewayConfigurationUsingPUTNotFoundBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this update gateway configuration using p u t not found body
func (o *UpdateGatewayConfigurationUsingPUTNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateGatewayConfigurationUsingPUTNotFoundBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("updateGatewayConfigurationUsingPUTNotFound"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateGatewayConfigurationUsingPUTNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateGatewayConfigurationUsingPUTNotFoundBody) UnmarshalBinary(b []byte) error {
	var res UpdateGatewayConfigurationUsingPUTNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
