// Code generated by go-swagger; DO NOT EDIT.

package users

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

// UpdateCustomPropertyForUserReader is a Reader for the UpdateCustomPropertyForUser structure.
type UpdateCustomPropertyForUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCustomPropertyForUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCustomPropertyForUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCustomPropertyForUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateCustomPropertyForUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateCustomPropertyForUserOK creates a UpdateCustomPropertyForUserOK with default headers values
func NewUpdateCustomPropertyForUserOK() *UpdateCustomPropertyForUserOK {
	return &UpdateCustomPropertyForUserOK{}
}

/*UpdateCustomPropertyForUserOK handles this case with default header values.

Successfully updated custom property.
*/
type UpdateCustomPropertyForUserOK struct {
	Payload *UpdateCustomPropertyForUserOKBody
}

func (o *UpdateCustomPropertyForUserOK) Error() string {
	return fmt.Sprintf("[PUT /v1/users/{userId}/customProperties/{key}][%d] updateCustomPropertyForUserOK  %+v", 200, o.Payload)
}

func (o *UpdateCustomPropertyForUserOK) GetPayload() *UpdateCustomPropertyForUserOKBody {
	return o.Payload
}

func (o *UpdateCustomPropertyForUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateCustomPropertyForUserOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomPropertyForUserBadRequest creates a UpdateCustomPropertyForUserBadRequest with default headers values
func NewUpdateCustomPropertyForUserBadRequest() *UpdateCustomPropertyForUserBadRequest {
	return &UpdateCustomPropertyForUserBadRequest{}
}

/*UpdateCustomPropertyForUserBadRequest handles this case with default header values.

HTTP request is malformed.
*/
type UpdateCustomPropertyForUserBadRequest struct {
	Payload *UpdateCustomPropertyForUserBadRequestBody
}

func (o *UpdateCustomPropertyForUserBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/users/{userId}/customProperties/{key}][%d] updateCustomPropertyForUserBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateCustomPropertyForUserBadRequest) GetPayload() *UpdateCustomPropertyForUserBadRequestBody {
	return o.Payload
}

func (o *UpdateCustomPropertyForUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateCustomPropertyForUserBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomPropertyForUserNotFound creates a UpdateCustomPropertyForUserNotFound with default headers values
func NewUpdateCustomPropertyForUserNotFound() *UpdateCustomPropertyForUserNotFound {
	return &UpdateCustomPropertyForUserNotFound{}
}

/*UpdateCustomPropertyForUserNotFound handles this case with default header values.

User or custom property with specified id or key does not exist.
*/
type UpdateCustomPropertyForUserNotFound struct {
	Payload *UpdateCustomPropertyForUserNotFoundBody
}

func (o *UpdateCustomPropertyForUserNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/users/{userId}/customProperties/{key}][%d] updateCustomPropertyForUserNotFound  %+v", 404, o.Payload)
}

func (o *UpdateCustomPropertyForUserNotFound) GetPayload() *UpdateCustomPropertyForUserNotFoundBody {
	return o.Payload
}

func (o *UpdateCustomPropertyForUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateCustomPropertyForUserNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateCustomPropertyForUserBadRequestBody Response
swagger:model UpdateCustomPropertyForUserBadRequestBody
*/
type UpdateCustomPropertyForUserBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this update custom property for user bad request body
func (o *UpdateCustomPropertyForUserBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateCustomPropertyForUserBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("updateCustomPropertyForUserBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateCustomPropertyForUserBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateCustomPropertyForUserBadRequestBody) UnmarshalBinary(b []byte) error {
	var res UpdateCustomPropertyForUserBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateCustomPropertyForUserBody CustomPropertyBean
swagger:model UpdateCustomPropertyForUserBody
*/
type UpdateCustomPropertyForUserBody struct {

	// A unique identifier of a custom property. Is defined by the user. This field is an immutable.
	// Required: true
	Key *string `json:"key"`

	// Value of the custom property.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this update custom property for user body
func (o *UpdateCustomPropertyForUserBody) Validate(formats strfmt.Registry) error {
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

func (o *UpdateCustomPropertyForUserBody) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("request"+"."+"key", "body", o.Key); err != nil {
		return err
	}

	return nil
}

func (o *UpdateCustomPropertyForUserBody) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("request"+"."+"value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateCustomPropertyForUserBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateCustomPropertyForUserBody) UnmarshalBinary(b []byte) error {
	var res UpdateCustomPropertyForUserBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateCustomPropertyForUserNotFoundBody Response
swagger:model UpdateCustomPropertyForUserNotFoundBody
*/
type UpdateCustomPropertyForUserNotFoundBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this update custom property for user not found body
func (o *UpdateCustomPropertyForUserNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateCustomPropertyForUserNotFoundBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("updateCustomPropertyForUserNotFound"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateCustomPropertyForUserNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateCustomPropertyForUserNotFoundBody) UnmarshalBinary(b []byte) error {
	var res UpdateCustomPropertyForUserNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateCustomPropertyForUserOKBody CustomPropertyBean
swagger:model UpdateCustomPropertyForUserOKBody
*/
type UpdateCustomPropertyForUserOKBody struct {

	// A unique identifier of a custom property. Is defined by the user. This field is an immutable.
	// Required: true
	Key *string `json:"key"`

	// Value of the custom property.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this update custom property for user o k body
func (o *UpdateCustomPropertyForUserOKBody) Validate(formats strfmt.Registry) error {
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

func (o *UpdateCustomPropertyForUserOKBody) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("updateCustomPropertyForUserOK"+"."+"key", "body", o.Key); err != nil {
		return err
	}

	return nil
}

func (o *UpdateCustomPropertyForUserOKBody) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("updateCustomPropertyForUserOK"+"."+"value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateCustomPropertyForUserOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateCustomPropertyForUserOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateCustomPropertyForUserOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
