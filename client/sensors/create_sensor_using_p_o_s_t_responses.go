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

// CreateSensorUsingPOSTReader is a Reader for the CreateSensorUsingPOST structure.
type CreateSensorUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSensorUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSensorUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSensorUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateSensorUsingPOSTConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSensorUsingPOSTOK creates a CreateSensorUsingPOSTOK with default headers values
func NewCreateSensorUsingPOSTOK() *CreateSensorUsingPOSTOK {
	return &CreateSensorUsingPOSTOK{}
}

/*CreateSensorUsingPOSTOK handles this case with default header values.

Successfully created sensor.
*/
type CreateSensorUsingPOSTOK struct {
	Payload *CreateSensorUsingPOSTOKBody
}

func (o *CreateSensorUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /v1/tenant/{tenantId}/sensors][%d] createSensorUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateSensorUsingPOSTOK) GetPayload() *CreateSensorUsingPOSTOKBody {
	return o.Payload
}

func (o *CreateSensorUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateSensorUsingPOSTOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSensorUsingPOSTBadRequest creates a CreateSensorUsingPOSTBadRequest with default headers values
func NewCreateSensorUsingPOSTBadRequest() *CreateSensorUsingPOSTBadRequest {
	return &CreateSensorUsingPOSTBadRequest{}
}

/*CreateSensorUsingPOSTBadRequest handles this case with default header values.

HTTP request is malformed.
*/
type CreateSensorUsingPOSTBadRequest struct {
	Payload *CreateSensorUsingPOSTBadRequestBody
}

func (o *CreateSensorUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/tenant/{tenantId}/sensors][%d] createSensorUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSensorUsingPOSTBadRequest) GetPayload() *CreateSensorUsingPOSTBadRequestBody {
	return o.Payload
}

func (o *CreateSensorUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateSensorUsingPOSTBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSensorUsingPOSTConflict creates a CreateSensorUsingPOSTConflict with default headers values
func NewCreateSensorUsingPOSTConflict() *CreateSensorUsingPOSTConflict {
	return &CreateSensorUsingPOSTConflict{}
}

/*CreateSensorUsingPOSTConflict handles this case with default header values.

Sensor cannot be created with the specified sensor type or device.
*/
type CreateSensorUsingPOSTConflict struct {
	Payload *CreateSensorUsingPOSTConflictBody
}

func (o *CreateSensorUsingPOSTConflict) Error() string {
	return fmt.Sprintf("[POST /v1/tenant/{tenantId}/sensors][%d] createSensorUsingPOSTConflict  %+v", 409, o.Payload)
}

func (o *CreateSensorUsingPOSTConflict) GetPayload() *CreateSensorUsingPOSTConflictBody {
	return o.Payload
}

func (o *CreateSensorUsingPOSTConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateSensorUsingPOSTConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateSensorUsingPOSTBadRequestBody Response
swagger:model CreateSensorUsingPOSTBadRequestBody
*/
type CreateSensorUsingPOSTBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this create sensor using p o s t bad request body
func (o *CreateSensorUsingPOSTBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateSensorUsingPOSTBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("createSensorUsingPOSTBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateSensorUsingPOSTBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateSensorUsingPOSTBadRequestBody) UnmarshalBinary(b []byte) error {
	var res CreateSensorUsingPOSTBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateSensorUsingPOSTBody SensorBean
swagger:model CreateSensorUsingPOSTBody
*/
type CreateSensorUsingPOSTBody struct {

	// The alternate ID of the sensor. This field is unique across device.
	AlternateID string `json:"alternateId,omitempty"`

	// A set of user defined properties represented as key value pair.
	CustomProperties []*CreateSensorUsingPOSTParamsBodyCustomPropertiesItems0 `json:"customProperties"`

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

// Validate validates this create sensor using p o s t body
func (o *CreateSensorUsingPOSTBody) Validate(formats strfmt.Registry) error {
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

func (o *CreateSensorUsingPOSTBody) validateCustomProperties(formats strfmt.Registry) error {

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

func (o *CreateSensorUsingPOSTBody) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("request"+"."+"deviceId", "body", o.DeviceID); err != nil {
		return err
	}

	return nil
}

func (o *CreateSensorUsingPOSTBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("request"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *CreateSensorUsingPOSTBody) validateSensorTypeID(formats strfmt.Registry) error {

	if err := validate.Required("request"+"."+"sensorTypeId", "body", o.SensorTypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateSensorUsingPOSTBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateSensorUsingPOSTBody) UnmarshalBinary(b []byte) error {
	var res CreateSensorUsingPOSTBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateSensorUsingPOSTConflictBody Response
swagger:model CreateSensorUsingPOSTConflictBody
*/
type CreateSensorUsingPOSTConflictBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this create sensor using p o s t conflict body
func (o *CreateSensorUsingPOSTConflictBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateSensorUsingPOSTConflictBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("createSensorUsingPOSTConflict"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateSensorUsingPOSTConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateSensorUsingPOSTConflictBody) UnmarshalBinary(b []byte) error {
	var res CreateSensorUsingPOSTConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateSensorUsingPOSTOKBody SensorBean
swagger:model CreateSensorUsingPOSTOKBody
*/
type CreateSensorUsingPOSTOKBody struct {

	// The alternate ID of the sensor. This field is unique across device.
	AlternateID string `json:"alternateId,omitempty"`

	// A set of user defined properties represented as key value pair.
	CustomProperties []*CreateSensorUsingPOSTOKBodyCustomPropertiesItems0 `json:"customProperties"`

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

// Validate validates this create sensor using p o s t o k body
func (o *CreateSensorUsingPOSTOKBody) Validate(formats strfmt.Registry) error {
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

func (o *CreateSensorUsingPOSTOKBody) validateCustomProperties(formats strfmt.Registry) error {

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
					return ve.ValidateName("createSensorUsingPOSTOK" + "." + "customProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *CreateSensorUsingPOSTOKBody) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("createSensorUsingPOSTOK"+"."+"deviceId", "body", o.DeviceID); err != nil {
		return err
	}

	return nil
}

func (o *CreateSensorUsingPOSTOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("createSensorUsingPOSTOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *CreateSensorUsingPOSTOKBody) validateSensorTypeID(formats strfmt.Registry) error {

	if err := validate.Required("createSensorUsingPOSTOK"+"."+"sensorTypeId", "body", o.SensorTypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateSensorUsingPOSTOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateSensorUsingPOSTOKBody) UnmarshalBinary(b []byte) error {
	var res CreateSensorUsingPOSTOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateSensorUsingPOSTOKBodyCustomPropertiesItems0 CustomPropertyBean
swagger:model CreateSensorUsingPOSTOKBodyCustomPropertiesItems0
*/
type CreateSensorUsingPOSTOKBodyCustomPropertiesItems0 struct {

	// A unique identifier of a custom property. Is defined by the user. This field is an immutable.
	// Required: true
	Key *string `json:"key"`

	// Value of the custom property.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this create sensor using p o s t o k body custom properties items0
func (o *CreateSensorUsingPOSTOKBodyCustomPropertiesItems0) Validate(formats strfmt.Registry) error {
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

func (o *CreateSensorUsingPOSTOKBodyCustomPropertiesItems0) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", o.Key); err != nil {
		return err
	}

	return nil
}

func (o *CreateSensorUsingPOSTOKBodyCustomPropertiesItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateSensorUsingPOSTOKBodyCustomPropertiesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateSensorUsingPOSTOKBodyCustomPropertiesItems0) UnmarshalBinary(b []byte) error {
	var res CreateSensorUsingPOSTOKBodyCustomPropertiesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateSensorUsingPOSTParamsBodyCustomPropertiesItems0 CustomPropertyBean
swagger:model CreateSensorUsingPOSTParamsBodyCustomPropertiesItems0
*/
type CreateSensorUsingPOSTParamsBodyCustomPropertiesItems0 struct {

	// A unique identifier of a custom property. Is defined by the user. This field is an immutable.
	// Required: true
	Key *string `json:"key"`

	// Value of the custom property.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this create sensor using p o s t params body custom properties items0
func (o *CreateSensorUsingPOSTParamsBodyCustomPropertiesItems0) Validate(formats strfmt.Registry) error {
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

func (o *CreateSensorUsingPOSTParamsBodyCustomPropertiesItems0) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", o.Key); err != nil {
		return err
	}

	return nil
}

func (o *CreateSensorUsingPOSTParamsBodyCustomPropertiesItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateSensorUsingPOSTParamsBodyCustomPropertiesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateSensorUsingPOSTParamsBodyCustomPropertiesItems0) UnmarshalBinary(b []byte) error {
	var res CreateSensorUsingPOSTParamsBodyCustomPropertiesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
