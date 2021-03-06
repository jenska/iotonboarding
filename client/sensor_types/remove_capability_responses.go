// Code generated by go-swagger; DO NOT EDIT.

package sensor_types

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

// RemoveCapabilityReader is a Reader for the RemoveCapability structure.
type RemoveCapabilityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveCapabilityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveCapabilityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRemoveCapabilityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRemoveCapabilityConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveCapabilityOK creates a RemoveCapabilityOK with default headers values
func NewRemoveCapabilityOK() *RemoveCapabilityOK {
	return &RemoveCapabilityOK{}
}

/*RemoveCapabilityOK handles this case with default header values.

Successfully removed capability.
*/
type RemoveCapabilityOK struct {
	Payload *RemoveCapabilityOKBody
}

func (o *RemoveCapabilityOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/tenant/{tenantId}/sensorTypes/{sensorTypeId}/capabilities/{capabilityId}][%d] removeCapabilityOK  %+v", 200, o.Payload)
}

func (o *RemoveCapabilityOK) GetPayload() *RemoveCapabilityOKBody {
	return o.Payload
}

func (o *RemoveCapabilityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RemoveCapabilityOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveCapabilityNotFound creates a RemoveCapabilityNotFound with default headers values
func NewRemoveCapabilityNotFound() *RemoveCapabilityNotFound {
	return &RemoveCapabilityNotFound{}
}

/*RemoveCapabilityNotFound handles this case with default header values.

Capability with specified id does not exist.
*/
type RemoveCapabilityNotFound struct {
	Payload *RemoveCapabilityNotFoundBody
}

func (o *RemoveCapabilityNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/tenant/{tenantId}/sensorTypes/{sensorTypeId}/capabilities/{capabilityId}][%d] removeCapabilityNotFound  %+v", 404, o.Payload)
}

func (o *RemoveCapabilityNotFound) GetPayload() *RemoveCapabilityNotFoundBody {
	return o.Payload
}

func (o *RemoveCapabilityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RemoveCapabilityNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveCapabilityConflict creates a RemoveCapabilityConflict with default headers values
func NewRemoveCapabilityConflict() *RemoveCapabilityConflict {
	return &RemoveCapabilityConflict{}
}

/*RemoveCapabilityConflict handles this case with default header values.

Capability is still used by sensor types.
*/
type RemoveCapabilityConflict struct {
	Payload *RemoveCapabilityConflictBody
}

func (o *RemoveCapabilityConflict) Error() string {
	return fmt.Sprintf("[DELETE /v1/tenant/{tenantId}/sensorTypes/{sensorTypeId}/capabilities/{capabilityId}][%d] removeCapabilityConflict  %+v", 409, o.Payload)
}

func (o *RemoveCapabilityConflict) GetPayload() *RemoveCapabilityConflictBody {
	return o.Payload
}

func (o *RemoveCapabilityConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RemoveCapabilityConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RemoveCapabilityConflictBody Response
swagger:model RemoveCapabilityConflictBody
*/
type RemoveCapabilityConflictBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this remove capability conflict body
func (o *RemoveCapabilityConflictBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveCapabilityConflictBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("removeCapabilityConflict"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RemoveCapabilityConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveCapabilityConflictBody) UnmarshalBinary(b []byte) error {
	var res RemoveCapabilityConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoveCapabilityNotFoundBody Response
swagger:model RemoveCapabilityNotFoundBody
*/
type RemoveCapabilityNotFoundBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this remove capability not found body
func (o *RemoveCapabilityNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveCapabilityNotFoundBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("removeCapabilityNotFound"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RemoveCapabilityNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveCapabilityNotFoundBody) UnmarshalBinary(b []byte) error {
	var res RemoveCapabilityNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoveCapabilityOKBody Response
swagger:model RemoveCapabilityOKBody
*/
type RemoveCapabilityOKBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this remove capability o k body
func (o *RemoveCapabilityOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveCapabilityOKBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("removeCapabilityOK"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RemoveCapabilityOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveCapabilityOKBody) UnmarshalBinary(b []byte) error {
	var res RemoveCapabilityOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
