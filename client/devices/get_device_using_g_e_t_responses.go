// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetDeviceUsingGETReader is a Reader for the GetDeviceUsingGET structure.
type GetDeviceUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDeviceUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeviceUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeviceUsingGETOK creates a GetDeviceUsingGETOK with default headers values
func NewGetDeviceUsingGETOK() *GetDeviceUsingGETOK {
	return &GetDeviceUsingGETOK{}
}

/*GetDeviceUsingGETOK handles this case with default header values.

Successfully returned device.
*/
type GetDeviceUsingGETOK struct {
	Payload *GetDeviceUsingGETOKBody
}

func (o *GetDeviceUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /v1/tenant/{tenantId}/devices/{deviceId}][%d] getDeviceUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetDeviceUsingGETOK) GetPayload() *GetDeviceUsingGETOKBody {
	return o.Payload
}

func (o *GetDeviceUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetDeviceUsingGETOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceUsingGETBadRequest creates a GetDeviceUsingGETBadRequest with default headers values
func NewGetDeviceUsingGETBadRequest() *GetDeviceUsingGETBadRequest {
	return &GetDeviceUsingGETBadRequest{}
}

/*GetDeviceUsingGETBadRequest handles this case with default header values.

HTTP request is malformed.
*/
type GetDeviceUsingGETBadRequest struct {
	Payload *GetDeviceUsingGETBadRequestBody
}

func (o *GetDeviceUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/tenant/{tenantId}/devices/{deviceId}][%d] getDeviceUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetDeviceUsingGETBadRequest) GetPayload() *GetDeviceUsingGETBadRequestBody {
	return o.Payload
}

func (o *GetDeviceUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetDeviceUsingGETBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceUsingGETNotFound creates a GetDeviceUsingGETNotFound with default headers values
func NewGetDeviceUsingGETNotFound() *GetDeviceUsingGETNotFound {
	return &GetDeviceUsingGETNotFound{}
}

/*GetDeviceUsingGETNotFound handles this case with default header values.

Device with specified id does not exist.
*/
type GetDeviceUsingGETNotFound struct {
	Payload *GetDeviceUsingGETNotFoundBody
}

func (o *GetDeviceUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/tenant/{tenantId}/devices/{deviceId}][%d] getDeviceUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetDeviceUsingGETNotFound) GetPayload() *GetDeviceUsingGETNotFoundBody {
	return o.Payload
}

func (o *GetDeviceUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetDeviceUsingGETNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetDeviceUsingGETBadRequestBody Response
swagger:model GetDeviceUsingGETBadRequestBody
*/
type GetDeviceUsingGETBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this get device using g e t bad request body
func (o *GetDeviceUsingGETBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDeviceUsingGETBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getDeviceUsingGETBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDeviceUsingGETBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeviceUsingGETBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetDeviceUsingGETBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDeviceUsingGETNotFoundBody Response
swagger:model GetDeviceUsingGETNotFoundBody
*/
type GetDeviceUsingGETNotFoundBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this get device using g e t not found body
func (o *GetDeviceUsingGETNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDeviceUsingGETNotFoundBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getDeviceUsingGETNotFound"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDeviceUsingGETNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeviceUsingGETNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetDeviceUsingGETNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDeviceUsingGETOKBody DeviceBean
swagger:model GetDeviceUsingGETOKBody
*/
type GetDeviceUsingGETOKBody struct {

	// The alternate ID of the device (e.g., serial number crafted into the device). This field is unique across gateway.
	AlternateID string `json:"alternateId,omitempty"`

	// Authentications of the device.
	// Read Only: true
	Authentications []*GetDeviceUsingGETOKBodyAuthenticationsItems0 `json:"authentications"`

	// Authorizations of the device.
	// Read Only: true
	Authorizations []*GetDeviceUsingGETOKBodyAuthorizationsItems0 `json:"authorizations"`

	// Unix time in milliseconds. The timestamp indicates when the device was added to the the platform. This field is an immutable.
	// Read Only: true
	// Format: date-time
	CreationTimestamp strfmt.DateTime `json:"creationTimestamp,omitempty"`

	// A set of user defined properties represented as key value pair.
	CustomProperties []*GetDeviceUsingGETOKBodyCustomPropertiesItems0 `json:"customProperties"`

	// The ID of the associated gateway. This field is an immutable.
	// Required: true
	GatewayID *string `json:"gatewayId"`

	// A unique identifier of a device. Is generated by the system and unique across platform tenant. This field is an immutable.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The name of the device.
	// Required: true
	Name *string `json:"name"`

	// The online status
	// Read Only: true
	Online *bool `json:"online,omitempty"`

	// The set of sensors that compose the device. DeviceId is not required to be given to sensor while creating device and sensor together
	Sensors []*GetDeviceUsingGETOKBodySensorsItems0 `json:"sensors"`
}

// Validate validates this get device using g e t o k body
func (o *GetDeviceUsingGETOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAuthentications(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAuthorizations(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCustomProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGatewayID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSensors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDeviceUsingGETOKBody) validateAuthentications(formats strfmt.Registry) error {

	if swag.IsZero(o.Authentications) { // not required
		return nil
	}

	for i := 0; i < len(o.Authentications); i++ {
		if swag.IsZero(o.Authentications[i]) { // not required
			continue
		}

		if o.Authentications[i] != nil {
			if err := o.Authentications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDeviceUsingGETOK" + "." + "authentications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetDeviceUsingGETOKBody) validateAuthorizations(formats strfmt.Registry) error {

	if swag.IsZero(o.Authorizations) { // not required
		return nil
	}

	for i := 0; i < len(o.Authorizations); i++ {
		if swag.IsZero(o.Authorizations[i]) { // not required
			continue
		}

		if o.Authorizations[i] != nil {
			if err := o.Authorizations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDeviceUsingGETOK" + "." + "authorizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetDeviceUsingGETOKBody) validateCreationTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(o.CreationTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("getDeviceUsingGETOK"+"."+"creationTimestamp", "body", "date-time", o.CreationTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetDeviceUsingGETOKBody) validateCustomProperties(formats strfmt.Registry) error {

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
					return ve.ValidateName("getDeviceUsingGETOK" + "." + "customProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetDeviceUsingGETOKBody) validateGatewayID(formats strfmt.Registry) error {

	if err := validate.Required("getDeviceUsingGETOK"+"."+"gatewayId", "body", o.GatewayID); err != nil {
		return err
	}

	return nil
}

func (o *GetDeviceUsingGETOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getDeviceUsingGETOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetDeviceUsingGETOKBody) validateSensors(formats strfmt.Registry) error {

	if swag.IsZero(o.Sensors) { // not required
		return nil
	}

	for i := 0; i < len(o.Sensors); i++ {
		if swag.IsZero(o.Sensors[i]) { // not required
			continue
		}

		if o.Sensors[i] != nil {
			if err := o.Sensors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDeviceUsingGETOK" + "." + "sensors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDeviceUsingGETOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeviceUsingGETOKBody) UnmarshalBinary(b []byte) error {
	var res GetDeviceUsingGETOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDeviceUsingGETOKBodyAuthenticationsItems0 AuthenticationBean
swagger:model GetDeviceUsingGETOKBodyAuthenticationsItems0
*/
type GetDeviceUsingGETOKBodyAuthenticationsItems0 struct {

	// The password for the user.
	Password string `json:"password,omitempty"`

	// The authentication type. Not required when updating user password.
	// Enum: [basic clientCertificate jwt]
	Type string `json:"type,omitempty"`
}

// Validate validates this get device using g e t o k body authentications items0
func (o *GetDeviceUsingGETOKBodyAuthenticationsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getDeviceUsingGETOKBodyAuthenticationsItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["basic","clientCertificate","jwt"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getDeviceUsingGETOKBodyAuthenticationsItems0TypeTypePropEnum = append(getDeviceUsingGETOKBodyAuthenticationsItems0TypeTypePropEnum, v)
	}
}

const (

	// GetDeviceUsingGETOKBodyAuthenticationsItems0TypeBasic captures enum value "basic"
	GetDeviceUsingGETOKBodyAuthenticationsItems0TypeBasic string = "basic"

	// GetDeviceUsingGETOKBodyAuthenticationsItems0TypeClientCertificate captures enum value "clientCertificate"
	GetDeviceUsingGETOKBodyAuthenticationsItems0TypeClientCertificate string = "clientCertificate"

	// GetDeviceUsingGETOKBodyAuthenticationsItems0TypeJwt captures enum value "jwt"
	GetDeviceUsingGETOKBodyAuthenticationsItems0TypeJwt string = "jwt"
)

// prop value enum
func (o *GetDeviceUsingGETOKBodyAuthenticationsItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getDeviceUsingGETOKBodyAuthenticationsItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetDeviceUsingGETOKBodyAuthenticationsItems0) validateType(formats strfmt.Registry) error {

	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDeviceUsingGETOKBodyAuthenticationsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeviceUsingGETOKBodyAuthenticationsItems0) UnmarshalBinary(b []byte) error {
	var res GetDeviceUsingGETOKBodyAuthenticationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDeviceUsingGETOKBodyAuthorizationsItems0 AuthorizationBean
swagger:model GetDeviceUsingGETOKBodyAuthorizationsItems0
*/
type GetDeviceUsingGETOKBodyAuthorizationsItems0 struct {

	// The type of device authorization policies
	// Enum: [router]
	Type string `json:"type,omitempty"`
}

// Validate validates this get device using g e t o k body authorizations items0
func (o *GetDeviceUsingGETOKBodyAuthorizationsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getDeviceUsingGETOKBodyAuthorizationsItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["router"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getDeviceUsingGETOKBodyAuthorizationsItems0TypeTypePropEnum = append(getDeviceUsingGETOKBodyAuthorizationsItems0TypeTypePropEnum, v)
	}
}

const (

	// GetDeviceUsingGETOKBodyAuthorizationsItems0TypeRouter captures enum value "router"
	GetDeviceUsingGETOKBodyAuthorizationsItems0TypeRouter string = "router"
)

// prop value enum
func (o *GetDeviceUsingGETOKBodyAuthorizationsItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getDeviceUsingGETOKBodyAuthorizationsItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetDeviceUsingGETOKBodyAuthorizationsItems0) validateType(formats strfmt.Registry) error {

	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDeviceUsingGETOKBodyAuthorizationsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeviceUsingGETOKBodyAuthorizationsItems0) UnmarshalBinary(b []byte) error {
	var res GetDeviceUsingGETOKBodyAuthorizationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDeviceUsingGETOKBodyCustomPropertiesItems0 CustomPropertyBean
swagger:model GetDeviceUsingGETOKBodyCustomPropertiesItems0
*/
type GetDeviceUsingGETOKBodyCustomPropertiesItems0 struct {

	// A unique identifier of a custom property. Is defined by the user. This field is an immutable.
	// Required: true
	Key *string `json:"key"`

	// Value of the custom property.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this get device using g e t o k body custom properties items0
func (o *GetDeviceUsingGETOKBodyCustomPropertiesItems0) Validate(formats strfmt.Registry) error {
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

func (o *GetDeviceUsingGETOKBodyCustomPropertiesItems0) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", o.Key); err != nil {
		return err
	}

	return nil
}

func (o *GetDeviceUsingGETOKBodyCustomPropertiesItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDeviceUsingGETOKBodyCustomPropertiesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeviceUsingGETOKBodyCustomPropertiesItems0) UnmarshalBinary(b []byte) error {
	var res GetDeviceUsingGETOKBodyCustomPropertiesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDeviceUsingGETOKBodySensorsItems0 SensorBean
swagger:model GetDeviceUsingGETOKBodySensorsItems0
*/
type GetDeviceUsingGETOKBodySensorsItems0 struct {

	// The alternate ID of the sensor. This field is unique across device.
	AlternateID string `json:"alternateId,omitempty"`

	// A set of user defined properties represented as key value pair.
	CustomProperties []*GetDeviceUsingGETOKBodySensorsItems0CustomPropertiesItems0 `json:"customProperties"`

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

// Validate validates this get device using g e t o k body sensors items0
func (o *GetDeviceUsingGETOKBodySensorsItems0) Validate(formats strfmt.Registry) error {
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

func (o *GetDeviceUsingGETOKBodySensorsItems0) validateCustomProperties(formats strfmt.Registry) error {

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
					return ve.ValidateName("customProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetDeviceUsingGETOKBodySensorsItems0) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("deviceId", "body", o.DeviceID); err != nil {
		return err
	}

	return nil
}

func (o *GetDeviceUsingGETOKBodySensorsItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetDeviceUsingGETOKBodySensorsItems0) validateSensorTypeID(formats strfmt.Registry) error {

	if err := validate.Required("sensorTypeId", "body", o.SensorTypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDeviceUsingGETOKBodySensorsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeviceUsingGETOKBodySensorsItems0) UnmarshalBinary(b []byte) error {
	var res GetDeviceUsingGETOKBodySensorsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDeviceUsingGETOKBodySensorsItems0CustomPropertiesItems0 CustomPropertyBean
swagger:model GetDeviceUsingGETOKBodySensorsItems0CustomPropertiesItems0
*/
type GetDeviceUsingGETOKBodySensorsItems0CustomPropertiesItems0 struct {

	// A unique identifier of a custom property. Is defined by the user. This field is an immutable.
	// Required: true
	Key *string `json:"key"`

	// Value of the custom property.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this get device using g e t o k body sensors items0 custom properties items0
func (o *GetDeviceUsingGETOKBodySensorsItems0CustomPropertiesItems0) Validate(formats strfmt.Registry) error {
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

func (o *GetDeviceUsingGETOKBodySensorsItems0CustomPropertiesItems0) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", o.Key); err != nil {
		return err
	}

	return nil
}

func (o *GetDeviceUsingGETOKBodySensorsItems0CustomPropertiesItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDeviceUsingGETOKBodySensorsItems0CustomPropertiesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeviceUsingGETOKBodySensorsItems0CustomPropertiesItems0) UnmarshalBinary(b []byte) error {
	var res GetDeviceUsingGETOKBodySensorsItems0CustomPropertiesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
