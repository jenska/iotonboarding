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

// BasicAuthenticationBean BasicAuthenticationBean
//
// swagger:model BasicAuthenticationBean
type BasicAuthenticationBean struct {

	// password
	// Required: true
	Password *string `json:"password"`

	// The type of the authentication mechanism
	// Required: true
	// Read Only: true
	// Enum: [basic]
	Type string `json:"type"`
}

// Validate validates this basic authentication bean
func (m *BasicAuthenticationBean) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BasicAuthenticationBean) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

var basicAuthenticationBeanTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["basic"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		basicAuthenticationBeanTypeTypePropEnum = append(basicAuthenticationBeanTypeTypePropEnum, v)
	}
}

const (

	// BasicAuthenticationBeanTypeBasic captures enum value "basic"
	BasicAuthenticationBeanTypeBasic string = "basic"
)

// prop value enum
func (m *BasicAuthenticationBean) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, basicAuthenticationBeanTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BasicAuthenticationBean) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BasicAuthenticationBean) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BasicAuthenticationBean) UnmarshalBinary(b []byte) error {
	var res BasicAuthenticationBean
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}