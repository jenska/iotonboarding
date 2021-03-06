// Code generated by go-swagger; DO NOT EDIT.

package users

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

// GetUsersUsingGETReader is a Reader for the GetUsersUsingGET structure.
type GetUsersUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsersUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsersUsingGETOK creates a GetUsersUsingGETOK with default headers values
func NewGetUsersUsingGETOK() *GetUsersUsingGETOK {
	return &GetUsersUsingGETOK{}
}

/*GetUsersUsingGETOK handles this case with default header values.

Successfully returned users.
*/
type GetUsersUsingGETOK struct {
	Payload []*GetUsersUsingGETOKBodyItems0
}

func (o *GetUsersUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] getUsersUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetUsersUsingGETOK) GetPayload() []*GetUsersUsingGETOKBodyItems0 {
	return o.Payload
}

func (o *GetUsersUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUsingGETBadRequest creates a GetUsersUsingGETBadRequest with default headers values
func NewGetUsersUsingGETBadRequest() *GetUsersUsingGETBadRequest {
	return &GetUsersUsingGETBadRequest{}
}

/*GetUsersUsingGETBadRequest handles this case with default header values.

HTTP request is malformed.
*/
type GetUsersUsingGETBadRequest struct {
	Payload *GetUsersUsingGETBadRequestBody
}

func (o *GetUsersUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] getUsersUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetUsersUsingGETBadRequest) GetPayload() *GetUsersUsingGETBadRequestBody {
	return o.Payload
}

func (o *GetUsersUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetUsersUsingGETBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetUsersUsingGETBadRequestBody Response
swagger:model GetUsersUsingGETBadRequestBody
*/
type GetUsersUsingGETBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this get users using g e t bad request body
func (o *GetUsersUsingGETBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUsersUsingGETBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getUsersUsingGETBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUsersUsingGETBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUsersUsingGETBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetUsersUsingGETBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUsersUsingGETOKBodyItems0 UserBean
swagger:model GetUsersUsingGETOKBodyItems0
*/
type GetUsersUsingGETOKBodyItems0 struct {

	// This value identifies if a user has been locked or not.
	// Read Only: true
	AccountNonLocked *bool `json:"accountNonLocked,omitempty"`

	// Authentications of the user.
	// Required: true
	Authentications []*GetUsersUsingGETOKBodyItems0AuthenticationsItems0 `json:"authentications"`

	// Unix time in milliseconds. The timestamp indicates when the user was added to the the platform. This field is an immutable.
	// Read Only: true
	// Format: date-time
	//	CreationTimestamp strfmt.DateTime `json:"creationTimestamp,omitempty"`

	// A set of user defined properties represented as key value pair.
	CustomProperties []*GetUsersUsingGETOKBodyItems0CustomPropertiesItems0 `json:"customProperties"`

	// A unique identifier of a user. Is generated by the system and unique across platform tenant. This field is an immutable.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// This value identifies when the user was last updated.
	// Read Only: true
	// Format: date-time
	//	LastUpdatedTimestamp strfmt.DateTime `json:"lastUpdatedTimestamp,omitempty"`

	// This value identifies the locked timestamp if the account is locked.
	// Read Only: true
	// Format: date-time
	//	LockedTimestamp strfmt.DateTime `json:"lockedTimestamp,omitempty"`

	// This value identifies the number of login attempts before being locked.
	// Read Only: true
	LoginAttempts int32 `json:"loginAttempts,omitempty"`

	// The name of the user. This field is not case sensitive and is unique across platform tenant.
	// Required: true
	Name *string `json:"name"`

	// A set of global user roles
	Roles []*GetUsersUsingGETOKBodyItems0RolesItems0 `json:"roles"`
}

// Validate validates this get users using g e t o k body items0
func (o *GetUsersUsingGETOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAuthentications(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCustomProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLastUpdatedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLockedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUsersUsingGETOKBodyItems0) validateAuthentications(formats strfmt.Registry) error {

	if err := validate.Required("authentications", "body", o.Authentications); err != nil {
		return err
	}

	for i := 0; i < len(o.Authentications); i++ {
		if swag.IsZero(o.Authentications[i]) { // not required
			continue
		}

		if o.Authentications[i] != nil {
			if err := o.Authentications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("authentications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetUsersUsingGETOKBodyItems0) validateCreationTimestamp(formats strfmt.Registry) error {
	/*
		if swag.IsZero(o.CreationTimestamp) { // not required
			return nil
		}

		if err := validate.FormatOf("creationTimestamp", "body", "date-time", o.CreationTimestamp.String(), formats); err != nil {
			return err
		}
	*/
	return nil
}

func (o *GetUsersUsingGETOKBodyItems0) validateCustomProperties(formats strfmt.Registry) error {

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

func (o *GetUsersUsingGETOKBodyItems0) validateLastUpdatedTimestamp(formats strfmt.Registry) error {
	/*
		if swag.IsZero(o.LastUpdatedTimestamp) { // not required
			return nil
		}

		if err := validate.FormatOf("lastUpdatedTimestamp", "body", "date-time", o.LastUpdatedTimestamp.String(), formats); err != nil {
			return err
		}
	*/
	return nil
}

func (o *GetUsersUsingGETOKBodyItems0) validateLockedTimestamp(formats strfmt.Registry) error {
	/*
		if swag.IsZero(o.LockedTimestamp) { // not required
			return nil
		}

		if err := validate.FormatOf("lockedTimestamp", "body", "date-time", o.LockedTimestamp.String(), formats); err != nil {
			return err
		}
	*/
	return nil
}

func (o *GetUsersUsingGETOKBodyItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetUsersUsingGETOKBodyItems0) validateRoles(formats strfmt.Registry) error {

	if swag.IsZero(o.Roles) { // not required
		return nil
	}

	for i := 0; i < len(o.Roles); i++ {
		if swag.IsZero(o.Roles[i]) { // not required
			continue
		}

		if o.Roles[i] != nil {
			if err := o.Roles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUsersUsingGETOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUsersUsingGETOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetUsersUsingGETOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUsersUsingGETOKBodyItems0AuthenticationsItems0 AuthenticationBean
swagger:model GetUsersUsingGETOKBodyItems0AuthenticationsItems0
*/
type GetUsersUsingGETOKBodyItems0AuthenticationsItems0 struct {

	// The password for the user.
	Password string `json:"password,omitempty"`

	// The authentication type. Not required when updating user password.
	// Enum: [basic clientCertificate jwt]
	Type string `json:"type,omitempty"`
}

// Validate validates this get users using g e t o k body items0 authentications items0
func (o *GetUsersUsingGETOKBodyItems0AuthenticationsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getUsersUsingGETOKBodyItems0AuthenticationsItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["basic","clientCertificate","jwt"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getUsersUsingGETOKBodyItems0AuthenticationsItems0TypeTypePropEnum = append(getUsersUsingGETOKBodyItems0AuthenticationsItems0TypeTypePropEnum, v)
	}
}

const (

	// GetUsersUsingGETOKBodyItems0AuthenticationsItems0TypeBasic captures enum value "basic"
	GetUsersUsingGETOKBodyItems0AuthenticationsItems0TypeBasic string = "basic"

	// GetUsersUsingGETOKBodyItems0AuthenticationsItems0TypeClientCertificate captures enum value "clientCertificate"
	GetUsersUsingGETOKBodyItems0AuthenticationsItems0TypeClientCertificate string = "clientCertificate"

	// GetUsersUsingGETOKBodyItems0AuthenticationsItems0TypeJwt captures enum value "jwt"
	GetUsersUsingGETOKBodyItems0AuthenticationsItems0TypeJwt string = "jwt"
)

// prop value enum
func (o *GetUsersUsingGETOKBodyItems0AuthenticationsItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getUsersUsingGETOKBodyItems0AuthenticationsItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetUsersUsingGETOKBodyItems0AuthenticationsItems0) validateType(formats strfmt.Registry) error {

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
func (o *GetUsersUsingGETOKBodyItems0AuthenticationsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUsersUsingGETOKBodyItems0AuthenticationsItems0) UnmarshalBinary(b []byte) error {
	var res GetUsersUsingGETOKBodyItems0AuthenticationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUsersUsingGETOKBodyItems0CustomPropertiesItems0 CustomPropertyBean
swagger:model GetUsersUsingGETOKBodyItems0CustomPropertiesItems0
*/
type GetUsersUsingGETOKBodyItems0CustomPropertiesItems0 struct {

	// A unique identifier of a custom property. Is defined by the user. This field is an immutable.
	// Required: true
	Key *string `json:"key"`

	// Value of the custom property.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this get users using g e t o k body items0 custom properties items0
func (o *GetUsersUsingGETOKBodyItems0CustomPropertiesItems0) Validate(formats strfmt.Registry) error {
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

func (o *GetUsersUsingGETOKBodyItems0CustomPropertiesItems0) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", o.Key); err != nil {
		return err
	}

	return nil
}

func (o *GetUsersUsingGETOKBodyItems0CustomPropertiesItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUsersUsingGETOKBodyItems0CustomPropertiesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUsersUsingGETOKBodyItems0CustomPropertiesItems0) UnmarshalBinary(b []byte) error {
	var res GetUsersUsingGETOKBodyItems0CustomPropertiesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUsersUsingGETOKBodyItems0RolesItems0 UserRoleBean
swagger:model GetUsersUsingGETOKBodyItems0RolesItems0
*/
type GetUsersUsingGETOKBodyItems0RolesItems0 struct {

	// The role of the user.
	// Required: true
	// Enum: [instanceOwner]
	Role *string `json:"role"`
}

// Validate validates this get users using g e t o k body items0 roles items0
func (o *GetUsersUsingGETOKBodyItems0RolesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getUsersUsingGETOKBodyItems0RolesItems0TypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["instanceOwner"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getUsersUsingGETOKBodyItems0RolesItems0TypeRolePropEnum = append(getUsersUsingGETOKBodyItems0RolesItems0TypeRolePropEnum, v)
	}
}

const (

	// GetUsersUsingGETOKBodyItems0RolesItems0RoleInstanceOwner captures enum value "instanceOwner"
	GetUsersUsingGETOKBodyItems0RolesItems0RoleInstanceOwner string = "instanceOwner"
)

// prop value enum
func (o *GetUsersUsingGETOKBodyItems0RolesItems0) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getUsersUsingGETOKBodyItems0RolesItems0TypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetUsersUsingGETOKBodyItems0RolesItems0) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", o.Role); err != nil {
		return err
	}

	// value enum
	if err := o.validateRoleEnum("role", "body", *o.Role); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUsersUsingGETOKBodyItems0RolesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUsersUsingGETOKBodyItems0RolesItems0) UnmarshalBinary(b []byte) error {
	var res GetUsersUsingGETOKBodyItems0RolesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
