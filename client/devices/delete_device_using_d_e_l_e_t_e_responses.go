// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// DeleteDeviceUsingDELETEReader is a Reader for the DeleteDeviceUsingDELETE structure.
type DeleteDeviceUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDeviceUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDeviceUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteDeviceUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDeviceUsingDELETENotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDeviceUsingDELETEOK creates a DeleteDeviceUsingDELETEOK with default headers values
func NewDeleteDeviceUsingDELETEOK() *DeleteDeviceUsingDELETEOK {
	return &DeleteDeviceUsingDELETEOK{}
}

/*DeleteDeviceUsingDELETEOK handles this case with default header values.

Successfully deleted device.
*/
type DeleteDeviceUsingDELETEOK struct {
	Payload *DeleteDeviceUsingDELETEOKBody
}

func (o *DeleteDeviceUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/tenant/{tenantId}/devices/{deviceId}][%d] deleteDeviceUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *DeleteDeviceUsingDELETEOK) GetPayload() *DeleteDeviceUsingDELETEOKBody {
	return o.Payload
}

func (o *DeleteDeviceUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteDeviceUsingDELETEOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceUsingDELETEBadRequest creates a DeleteDeviceUsingDELETEBadRequest with default headers values
func NewDeleteDeviceUsingDELETEBadRequest() *DeleteDeviceUsingDELETEBadRequest {
	return &DeleteDeviceUsingDELETEBadRequest{}
}

/*DeleteDeviceUsingDELETEBadRequest handles this case with default header values.

HTTP request is malformed.
*/
type DeleteDeviceUsingDELETEBadRequest struct {
	Payload *DeleteDeviceUsingDELETEBadRequestBody
}

func (o *DeleteDeviceUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/tenant/{tenantId}/devices/{deviceId}][%d] deleteDeviceUsingDELETEBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteDeviceUsingDELETEBadRequest) GetPayload() *DeleteDeviceUsingDELETEBadRequestBody {
	return o.Payload
}

func (o *DeleteDeviceUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteDeviceUsingDELETEBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceUsingDELETENotFound creates a DeleteDeviceUsingDELETENotFound with default headers values
func NewDeleteDeviceUsingDELETENotFound() *DeleteDeviceUsingDELETENotFound {
	return &DeleteDeviceUsingDELETENotFound{}
}

/*DeleteDeviceUsingDELETENotFound handles this case with default header values.

Device with specified id does not exist.
*/
type DeleteDeviceUsingDELETENotFound struct {
	Payload *DeleteDeviceUsingDELETENotFoundBody
}

func (o *DeleteDeviceUsingDELETENotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/tenant/{tenantId}/devices/{deviceId}][%d] deleteDeviceUsingDELETENotFound  %+v", 404, o.Payload)
}

func (o *DeleteDeviceUsingDELETENotFound) GetPayload() *DeleteDeviceUsingDELETENotFoundBody {
	return o.Payload
}

func (o *DeleteDeviceUsingDELETENotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteDeviceUsingDELETENotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteDeviceUsingDELETEBadRequestBody Response
swagger:model DeleteDeviceUsingDELETEBadRequestBody
*/
type DeleteDeviceUsingDELETEBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this delete device using d e l e t e bad request body
func (o *DeleteDeviceUsingDELETEBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteDeviceUsingDELETEBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("deleteDeviceUsingDELETEBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteDeviceUsingDELETEBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteDeviceUsingDELETEBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DeleteDeviceUsingDELETEBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteDeviceUsingDELETENotFoundBody Response
swagger:model DeleteDeviceUsingDELETENotFoundBody
*/
type DeleteDeviceUsingDELETENotFoundBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this delete device using d e l e t e not found body
func (o *DeleteDeviceUsingDELETENotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteDeviceUsingDELETENotFoundBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("deleteDeviceUsingDELETENotFound"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteDeviceUsingDELETENotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteDeviceUsingDELETENotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteDeviceUsingDELETENotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteDeviceUsingDELETEOKBody Response
swagger:model DeleteDeviceUsingDELETEOKBody
*/
type DeleteDeviceUsingDELETEOKBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this delete device using d e l e t e o k body
func (o *DeleteDeviceUsingDELETEOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteDeviceUsingDELETEOKBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("deleteDeviceUsingDELETEOK"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteDeviceUsingDELETEOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteDeviceUsingDELETEOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteDeviceUsingDELETEOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
