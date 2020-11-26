// Code generated by go-swagger; DO NOT EDIT.

package sensors

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

// UpdateCustomPropertyForSensorReader is a Reader for the UpdateCustomPropertyForSensor structure.
type UpdateCustomPropertyForSensorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCustomPropertyForSensorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCustomPropertyForSensorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCustomPropertyForSensorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateCustomPropertyForSensorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateCustomPropertyForSensorOK creates a UpdateCustomPropertyForSensorOK with default headers values
func NewUpdateCustomPropertyForSensorOK() *UpdateCustomPropertyForSensorOK {
	return &UpdateCustomPropertyForSensorOK{}
}

/*UpdateCustomPropertyForSensorOK handles this case with default header values.

Successfully updated custom property.
*/
type UpdateCustomPropertyForSensorOK struct {
	Payload *UpdateCustomPropertyForSensorOKBody
}

func (o *UpdateCustomPropertyForSensorOK) Error() string {
	return fmt.Sprintf("[PUT /v1/tenant/{tenantId}/sensors/{sensorId}/customProperties/{key}][%d] updateCustomPropertyForSensorOK  %+v", 200, o.Payload)
}

func (o *UpdateCustomPropertyForSensorOK) GetPayload() *UpdateCustomPropertyForSensorOKBody {
	return o.Payload
}

func (o *UpdateCustomPropertyForSensorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateCustomPropertyForSensorOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomPropertyForSensorBadRequest creates a UpdateCustomPropertyForSensorBadRequest with default headers values
func NewUpdateCustomPropertyForSensorBadRequest() *UpdateCustomPropertyForSensorBadRequest {
	return &UpdateCustomPropertyForSensorBadRequest{}
}

/*UpdateCustomPropertyForSensorBadRequest handles this case with default header values.

HTTP request is malformed.
*/
type UpdateCustomPropertyForSensorBadRequest struct {
	Payload *UpdateCustomPropertyForSensorBadRequestBody
}

func (o *UpdateCustomPropertyForSensorBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/tenant/{tenantId}/sensors/{sensorId}/customProperties/{key}][%d] updateCustomPropertyForSensorBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateCustomPropertyForSensorBadRequest) GetPayload() *UpdateCustomPropertyForSensorBadRequestBody {
	return o.Payload
}

func (o *UpdateCustomPropertyForSensorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateCustomPropertyForSensorBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomPropertyForSensorNotFound creates a UpdateCustomPropertyForSensorNotFound with default headers values
func NewUpdateCustomPropertyForSensorNotFound() *UpdateCustomPropertyForSensorNotFound {
	return &UpdateCustomPropertyForSensorNotFound{}
}

/*UpdateCustomPropertyForSensorNotFound handles this case with default header values.

Sensor or custom property with specified id or key does not exist.
*/
type UpdateCustomPropertyForSensorNotFound struct {
	Payload *UpdateCustomPropertyForSensorNotFoundBody
}

func (o *UpdateCustomPropertyForSensorNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/tenant/{tenantId}/sensors/{sensorId}/customProperties/{key}][%d] updateCustomPropertyForSensorNotFound  %+v", 404, o.Payload)
}

func (o *UpdateCustomPropertyForSensorNotFound) GetPayload() *UpdateCustomPropertyForSensorNotFoundBody {
	return o.Payload
}

func (o *UpdateCustomPropertyForSensorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateCustomPropertyForSensorNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateCustomPropertyForSensorBadRequestBody Response
swagger:model UpdateCustomPropertyForSensorBadRequestBody
*/
type UpdateCustomPropertyForSensorBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this update custom property for sensor bad request body
func (o *UpdateCustomPropertyForSensorBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateCustomPropertyForSensorBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("updateCustomPropertyForSensorBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateCustomPropertyForSensorBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateCustomPropertyForSensorBadRequestBody) UnmarshalBinary(b []byte) error {
	var res UpdateCustomPropertyForSensorBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateCustomPropertyForSensorBody CustomPropertyBean
swagger:model UpdateCustomPropertyForSensorBody
*/
type UpdateCustomPropertyForSensorBody struct {

	// A unique identifier of a custom property. Is defined by the user. This field is an immutable.
	// Required: true
	Key *string `json:"key"`

	// Value of the custom property.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this update custom property for sensor body
func (o *UpdateCustomPropertyForSensorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateCustomPropertyForSensorBody) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("request"+"."+"key", "body", o.Key); err != nil {
		return err
	}

	return nil
}

func (o *UpdateCustomPropertyForSensorBody) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("request"+"."+"value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateCustomPropertyForSensorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateCustomPropertyForSensorBody) UnmarshalBinary(b []byte) error {
	var res UpdateCustomPropertyForSensorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateCustomPropertyForSensorNotFoundBody Response
swagger:model UpdateCustomPropertyForSensorNotFoundBody
*/
type UpdateCustomPropertyForSensorNotFoundBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this update custom property for sensor not found body
func (o *UpdateCustomPropertyForSensorNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateCustomPropertyForSensorNotFoundBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("updateCustomPropertyForSensorNotFound"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateCustomPropertyForSensorNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateCustomPropertyForSensorNotFoundBody) UnmarshalBinary(b []byte) error {
	var res UpdateCustomPropertyForSensorNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateCustomPropertyForSensorOKBody CustomPropertyBean
swagger:model UpdateCustomPropertyForSensorOKBody
*/
type UpdateCustomPropertyForSensorOKBody struct {

	// A unique identifier of a custom property. Is defined by the user. This field is an immutable.
	// Required: true
	Key *string `json:"key"`

	// Value of the custom property.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this update custom property for sensor o k body
func (o *UpdateCustomPropertyForSensorOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateCustomPropertyForSensorOKBody) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("updateCustomPropertyForSensorOK"+"."+"key", "body", o.Key); err != nil {
		return err
	}

	return nil
}

func (o *UpdateCustomPropertyForSensorOKBody) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("updateCustomPropertyForSensorOK"+"."+"value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateCustomPropertyForSensorOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateCustomPropertyForSensorOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateCustomPropertyForSensorOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}