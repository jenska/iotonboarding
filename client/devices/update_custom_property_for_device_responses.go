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

// UpdateCustomPropertyForDeviceReader is a Reader for the UpdateCustomPropertyForDevice structure.
type UpdateCustomPropertyForDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCustomPropertyForDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCustomPropertyForDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCustomPropertyForDeviceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateCustomPropertyForDeviceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateCustomPropertyForDeviceOK creates a UpdateCustomPropertyForDeviceOK with default headers values
func NewUpdateCustomPropertyForDeviceOK() *UpdateCustomPropertyForDeviceOK {
	return &UpdateCustomPropertyForDeviceOK{}
}

/*UpdateCustomPropertyForDeviceOK handles this case with default header values.

Successfully updated custom property.
*/
type UpdateCustomPropertyForDeviceOK struct {
	Payload *UpdateCustomPropertyForDeviceOKBody
}

func (o *UpdateCustomPropertyForDeviceOK) Error() string {
	return fmt.Sprintf("[PUT /v1/tenant/{tenantId}/devices/{deviceId}/customProperties/{key}][%d] updateCustomPropertyForDeviceOK  %+v", 200, o.Payload)
}

func (o *UpdateCustomPropertyForDeviceOK) GetPayload() *UpdateCustomPropertyForDeviceOKBody {
	return o.Payload
}

func (o *UpdateCustomPropertyForDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateCustomPropertyForDeviceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomPropertyForDeviceBadRequest creates a UpdateCustomPropertyForDeviceBadRequest with default headers values
func NewUpdateCustomPropertyForDeviceBadRequest() *UpdateCustomPropertyForDeviceBadRequest {
	return &UpdateCustomPropertyForDeviceBadRequest{}
}

/*UpdateCustomPropertyForDeviceBadRequest handles this case with default header values.

HTTP request is malformed.
*/
type UpdateCustomPropertyForDeviceBadRequest struct {
	Payload *UpdateCustomPropertyForDeviceBadRequestBody
}

func (o *UpdateCustomPropertyForDeviceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/tenant/{tenantId}/devices/{deviceId}/customProperties/{key}][%d] updateCustomPropertyForDeviceBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateCustomPropertyForDeviceBadRequest) GetPayload() *UpdateCustomPropertyForDeviceBadRequestBody {
	return o.Payload
}

func (o *UpdateCustomPropertyForDeviceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateCustomPropertyForDeviceBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomPropertyForDeviceNotFound creates a UpdateCustomPropertyForDeviceNotFound with default headers values
func NewUpdateCustomPropertyForDeviceNotFound() *UpdateCustomPropertyForDeviceNotFound {
	return &UpdateCustomPropertyForDeviceNotFound{}
}

/*UpdateCustomPropertyForDeviceNotFound handles this case with default header values.

Device or custom property with specified id or key does not exist.
*/
type UpdateCustomPropertyForDeviceNotFound struct {
	Payload *UpdateCustomPropertyForDeviceNotFoundBody
}

func (o *UpdateCustomPropertyForDeviceNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/tenant/{tenantId}/devices/{deviceId}/customProperties/{key}][%d] updateCustomPropertyForDeviceNotFound  %+v", 404, o.Payload)
}

func (o *UpdateCustomPropertyForDeviceNotFound) GetPayload() *UpdateCustomPropertyForDeviceNotFoundBody {
	return o.Payload
}

func (o *UpdateCustomPropertyForDeviceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateCustomPropertyForDeviceNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateCustomPropertyForDeviceBadRequestBody Response
swagger:model UpdateCustomPropertyForDeviceBadRequestBody
*/
type UpdateCustomPropertyForDeviceBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this update custom property for device bad request body
func (o *UpdateCustomPropertyForDeviceBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateCustomPropertyForDeviceBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("updateCustomPropertyForDeviceBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateCustomPropertyForDeviceBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateCustomPropertyForDeviceBadRequestBody) UnmarshalBinary(b []byte) error {
	var res UpdateCustomPropertyForDeviceBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateCustomPropertyForDeviceBody CustomPropertyBean
swagger:model UpdateCustomPropertyForDeviceBody
*/
type UpdateCustomPropertyForDeviceBody struct {

	// A unique identifier of a custom property. Is defined by the user. This field is an immutable.
	// Required: true
	Key *string `json:"key"`

	// Value of the custom property.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this update custom property for device body
func (o *UpdateCustomPropertyForDeviceBody) Validate(formats strfmt.Registry) error {
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

func (o *UpdateCustomPropertyForDeviceBody) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("request"+"."+"key", "body", o.Key); err != nil {
		return err
	}

	return nil
}

func (o *UpdateCustomPropertyForDeviceBody) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("request"+"."+"value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateCustomPropertyForDeviceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateCustomPropertyForDeviceBody) UnmarshalBinary(b []byte) error {
	var res UpdateCustomPropertyForDeviceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateCustomPropertyForDeviceNotFoundBody Response
swagger:model UpdateCustomPropertyForDeviceNotFoundBody
*/
type UpdateCustomPropertyForDeviceNotFoundBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this update custom property for device not found body
func (o *UpdateCustomPropertyForDeviceNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateCustomPropertyForDeviceNotFoundBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("updateCustomPropertyForDeviceNotFound"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateCustomPropertyForDeviceNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateCustomPropertyForDeviceNotFoundBody) UnmarshalBinary(b []byte) error {
	var res UpdateCustomPropertyForDeviceNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateCustomPropertyForDeviceOKBody CustomPropertyBean
swagger:model UpdateCustomPropertyForDeviceOKBody
*/
type UpdateCustomPropertyForDeviceOKBody struct {

	// A unique identifier of a custom property. Is defined by the user. This field is an immutable.
	// Required: true
	Key *string `json:"key"`

	// Value of the custom property.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this update custom property for device o k body
func (o *UpdateCustomPropertyForDeviceOKBody) Validate(formats strfmt.Registry) error {
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

func (o *UpdateCustomPropertyForDeviceOKBody) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("updateCustomPropertyForDeviceOK"+"."+"key", "body", o.Key); err != nil {
		return err
	}

	return nil
}

func (o *UpdateCustomPropertyForDeviceOKBody) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("updateCustomPropertyForDeviceOK"+"."+"value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateCustomPropertyForDeviceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateCustomPropertyForDeviceOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateCustomPropertyForDeviceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
