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

// TenantUserAssignmentBean TenantUserAssignmentBean
//
// swagger:model TenantUserAssignmentBean
type TenantUserAssignmentBean struct {

	// role
	// Enum: [administrator user]
	Role string `json:"role,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`
}

// Validate validates this tenant user assignment bean
func (m *TenantUserAssignmentBean) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var tenantUserAssignmentBeanTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["administrator","user"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tenantUserAssignmentBeanTypeRolePropEnum = append(tenantUserAssignmentBeanTypeRolePropEnum, v)
	}
}

const (

	// TenantUserAssignmentBeanRoleAdministrator captures enum value "administrator"
	TenantUserAssignmentBeanRoleAdministrator string = "administrator"

	// TenantUserAssignmentBeanRoleUser captures enum value "user"
	TenantUserAssignmentBeanRoleUser string = "user"
)

// prop value enum
func (m *TenantUserAssignmentBean) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tenantUserAssignmentBeanTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TenantUserAssignmentBean) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TenantUserAssignmentBean) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TenantUserAssignmentBean) UnmarshalBinary(b []byte) error {
	var res TenantUserAssignmentBean
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
