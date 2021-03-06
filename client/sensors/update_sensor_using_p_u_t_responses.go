// Code generated by go-swagger; DO NOT EDIT.

package sensors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateSensorUsingPUTReader is a Reader for the UpdateSensorUsingPUT structure.
type UpdateSensorUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSensorUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSensorUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSensorUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateSensorUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateSensorUsingPUTConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateSensorUsingPUTOK creates a UpdateSensorUsingPUTOK with default headers values
func NewUpdateSensorUsingPUTOK() *UpdateSensorUsingPUTOK {
	return &UpdateSensorUsingPUTOK{}
}

/*UpdateSensorUsingPUTOK handles this case with default header values.

Successfully updated sensor.
*/
type UpdateSensorUsingPUTOK struct {
	Payload *UpdateSensorUsingPUTOKBody
}

func (o *UpdateSensorUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /v1/tenant/{tenantId}/sensors/{sensorId}][%d] updateSensorUsingPUTOK  %+v", 200, o.Payload)
}

func (o *UpdateSensorUsingPUTOK) GetPayload() *UpdateSensorUsingPUTOKBody {
	return o.Payload
}

func (o *UpdateSensorUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateSensorUsingPUTOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSensorUsingPUTBadRequest creates a UpdateSensorUsingPUTBadRequest with default headers values
func NewUpdateSensorUsingPUTBadRequest() *UpdateSensorUsingPUTBadRequest {
	return &UpdateSensorUsingPUTBadRequest{}
}

/*UpdateSensorUsingPUTBadRequest handles this case with default header values.

HTTP request is malformed.
*/
type UpdateSensorUsingPUTBadRequest struct {
	Payload *UpdateSensorUsingPUTBadRequestBody
}

func (o *UpdateSensorUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/tenant/{tenantId}/sensors/{sensorId}][%d] updateSensorUsingPUTBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSensorUsingPUTBadRequest) GetPayload() *UpdateSensorUsingPUTBadRequestBody {
	return o.Payload
}

func (o *UpdateSensorUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateSensorUsingPUTBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSensorUsingPUTNotFound creates a UpdateSensorUsingPUTNotFound with default headers values
func NewUpdateSensorUsingPUTNotFound() *UpdateSensorUsingPUTNotFound {
	return &UpdateSensorUsingPUTNotFound{}
}

/*UpdateSensorUsingPUTNotFound handles this case with default header values.

Sensor with specified id does not exist.
*/
type UpdateSensorUsingPUTNotFound struct {
	Payload *UpdateSensorUsingPUTNotFoundBody
}

func (o *UpdateSensorUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/tenant/{tenantId}/sensors/{sensorId}][%d] updateSensorUsingPUTNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSensorUsingPUTNotFound) GetPayload() *UpdateSensorUsingPUTNotFoundBody {
	return o.Payload
}

func (o *UpdateSensorUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateSensorUsingPUTNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSensorUsingPUTConflict creates a UpdateSensorUsingPUTConflict with default headers values
func NewUpdateSensorUsingPUTConflict() *UpdateSensorUsingPUTConflict {
	return &UpdateSensorUsingPUTConflict{}
}

/*UpdateSensorUsingPUTConflict handles this case with default header values.

Sensor cannot be updated with the specified sensor type.
*/
type UpdateSensorUsingPUTConflict struct {
	Payload *UpdateSensorUsingPUTConflictBody
}

func (o *UpdateSensorUsingPUTConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/tenant/{tenantId}/sensors/{sensorId}][%d] updateSensorUsingPUTConflict  %+v", 409, o.Payload)
}

func (o *UpdateSensorUsingPUTConflict) GetPayload() *UpdateSensorUsingPUTConflictBody {
	return o.Payload
}

func (o *UpdateSensorUsingPUTConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateSensorUsingPUTConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateSensorUsingPUTBadRequestBody Response
swagger:model UpdateSensorUsingPUTBadRequestBody
*/
type UpdateSensorUsingPUTBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this update sensor using p u t bad request body
func (o *UpdateSensorUsingPUTBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateSensorUsingPUTBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("updateSensorUsingPUTBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateSensorUsingPUTBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateSensorUsingPUTBadRequestBody) UnmarshalBinary(b []byte) error {
	var res UpdateSensorUsingPUTBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateSensorUsingPUTBody SensorBean
swagger:model UpdateSensorUsingPUTBody
*/
type UpdateSensorUsingPUTBody struct {

	// The alternate ID of the sensor. This field is unique across device.
	AlternateID string `json:"alternateId,omitempty"`

	// A set of user defined properties represented as key value pair.
	CustomProperties []*UpdateSensorUsingPUTParamsBodyCustomPropertiesItems0 `json:"customProperties"`

	// The ID of the related device. DeviceId is not required to be given to sensor while creating device and sensor together
	// Required: true
	DeviceID *string `json:"deviceId"`

	// Identifier of a sensor. Is generated by the system and unique across platform tenant. This field is an immutable.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The name of the sensor
	// Required: true
	Name *string `json:"name"`

	// The ID of the SensorType
	// Required: true
	SensorTypeID *string `json:"sensorTypeId"`
}

// Validate validates this update sensor using p u t body
func (o *UpdateSensorUsingPUTBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCustomProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSensorTypeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateSensorUsingPUTBody) validateCustomProperties(formats strfmt.Registry) error {

	if swag.IsZero(o.CustomProperties) { // not required
		return nil
	}

	for i := 0; i < len(o.CustomProperties); i++ {
		if swag.IsZero(o.CustomProperties[i]) { // not required
			continue
		}

		if o.CustomProperties[i] != nil {
			if err := o.CustomProperties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("request" + "." + "customProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UpdateSensorUsingPUTBody) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("request"+"."+"deviceId", "body", o.DeviceID); err != nil {
		return err
	}

	return nil
}

func (o *UpdateSensorUsingPUTBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("request"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *UpdateSensorUsingPUTBody) validateSensorTypeID(formats strfmt.Registry) error {

	if err := validate.Required("request"+"."+"sensorTypeId", "body", o.SensorTypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateSensorUsingPUTBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateSensorUsingPUTBody) UnmarshalBinary(b []byte) error {
	var res UpdateSensorUsingPUTBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateSensorUsingPUTConflictBody Response
swagger:model UpdateSensorUsingPUTConflictBody
*/
type UpdateSensorUsingPUTConflictBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this update sensor using p u t conflict body
func (o *UpdateSensorUsingPUTConflictBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateSensorUsingPUTConflictBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("updateSensorUsingPUTConflict"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateSensorUsingPUTConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateSensorUsingPUTConflictBody) UnmarshalBinary(b []byte) error {
	var res UpdateSensorUsingPUTConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateSensorUsingPUTNotFoundBody Response
swagger:model UpdateSensorUsingPUTNotFoundBody
*/
type UpdateSensorUsingPUTNotFoundBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this update sensor using p u t not found body
func (o *UpdateSensorUsingPUTNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateSensorUsingPUTNotFoundBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("updateSensorUsingPUTNotFound"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateSensorUsingPUTNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateSensorUsingPUTNotFoundBody) UnmarshalBinary(b []byte) error {
	var res UpdateSensorUsingPUTNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateSensorUsingPUTOKBody SensorBean
swagger:model UpdateSensorUsingPUTOKBody
*/
type UpdateSensorUsingPUTOKBody struct {

	// The alternate ID of the sensor. This field is unique across device.
	AlternateID string `json:"alternateId,omitempty"`

	// A set of user defined properties represented as key value pair.
	CustomProperties []*UpdateSensorUsingPUTOKBodyCustomPropertiesItems0 `json:"customProperties"`

	// The ID of the related device. DeviceId is not required to be given to sensor while creating device and sensor together
	// Required: true
	DeviceID *string `json:"deviceId"`

	// Identifier of a sensor. Is generated by the system and unique across platform tenant. This field is an immutable.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The name of the sensor
	// Required: true
	Name *string `json:"name"`

	// The ID of the SensorType
	// Required: true
	SensorTypeID *string `json:"sensorTypeId"`
}

// Validate validates this update sensor using p u t o k body
func (o *UpdateSensorUsingPUTOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCustomProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSensorTypeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateSensorUsingPUTOKBody) validateCustomProperties(formats strfmt.Registry) error {

	if swag.IsZero(o.CustomProperties) { // not required
		return nil
	}

	for i := 0; i < len(o.CustomProperties); i++ {
		if swag.IsZero(o.CustomProperties[i]) { // not required
			continue
		}

		if o.CustomProperties[i] != nil {
			if err := o.CustomProperties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateSensorUsingPUTOK" + "." + "customProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UpdateSensorUsingPUTOKBody) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("updateSensorUsingPUTOK"+"."+"deviceId", "body", o.DeviceID); err != nil {
		return err
	}

	return nil
}

func (o *UpdateSensorUsingPUTOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("updateSensorUsingPUTOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *UpdateSensorUsingPUTOKBody) validateSensorTypeID(formats strfmt.Registry) error {

	if err := validate.Required("updateSensorUsingPUTOK"+"."+"sensorTypeId", "body", o.SensorTypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateSensorUsingPUTOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateSensorUsingPUTOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateSensorUsingPUTOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateSensorUsingPUTOKBodyCustomPropertiesItems0 CustomPropertyBean
swagger:model UpdateSensorUsingPUTOKBodyCustomPropertiesItems0
*/
type UpdateSensorUsingPUTOKBodyCustomPropertiesItems0 struct {

	// A unique identifier of a custom property. Is defined by the user. This field is an immutable.
	// Required: true
	Key *string `json:"key"`

	// Value of the custom property.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this update sensor using p u t o k body custom properties items0
func (o *UpdateSensorUsingPUTOKBodyCustomPropertiesItems0) Validate(formats strfmt.Registry) error {
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

func (o *UpdateSensorUsingPUTOKBodyCustomPropertiesItems0) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", o.Key); err != nil {
		return err
	}

	return nil
}

func (o *UpdateSensorUsingPUTOKBodyCustomPropertiesItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateSensorUsingPUTOKBodyCustomPropertiesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateSensorUsingPUTOKBodyCustomPropertiesItems0) UnmarshalBinary(b []byte) error {
	var res UpdateSensorUsingPUTOKBodyCustomPropertiesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateSensorUsingPUTParamsBodyCustomPropertiesItems0 CustomPropertyBean
swagger:model UpdateSensorUsingPUTParamsBodyCustomPropertiesItems0
*/
type UpdateSensorUsingPUTParamsBodyCustomPropertiesItems0 struct {

	// A unique identifier of a custom property. Is defined by the user. This field is an immutable.
	// Required: true
	Key *string `json:"key"`

	// Value of the custom property.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this update sensor using p u t params body custom properties items0
func (o *UpdateSensorUsingPUTParamsBodyCustomPropertiesItems0) Validate(formats strfmt.Registry) error {
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

func (o *UpdateSensorUsingPUTParamsBodyCustomPropertiesItems0) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", o.Key); err != nil {
		return err
	}

	return nil
}

func (o *UpdateSensorUsingPUTParamsBodyCustomPropertiesItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateSensorUsingPUTParamsBodyCustomPropertiesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateSensorUsingPUTParamsBodyCustomPropertiesItems0) UnmarshalBinary(b []byte) error {
	var res UpdateSensorUsingPUTParamsBodyCustomPropertiesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
