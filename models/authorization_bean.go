// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AuthorizationBean AuthorizationBean
//
// swagger:model AuthorizationBean
type AuthorizationBean struct {

	// The type of device authorization policies
	// Enum: [router]
	Type string `json:"type,omitempty"`
}

// Validate validates this authorization bean
func (m *AuthorizationBean) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var authorizationBeanTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["router"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		authorizationBeanTypeTypePropEnum = append(authorizationBeanTypeTypePropEnum, v)
	}
}

const (

	// AuthorizationBeanTypeRouter captures enum value "router"
	AuthorizationBeanTypeRouter string = "router"
)

// prop value enum
func (m *AuthorizationBean) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, authorizationBeanTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AuthorizationBean) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthorizationBean) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthorizationBean) UnmarshalBinary(b []byte) error {
	var res AuthorizationBean
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}