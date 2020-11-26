// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TrustedCACertificateBean TrustedCACertificateBean
//
// swagger:model TrustedCACertificateBean
type TrustedCACertificateBean struct {

	// pem
	// Required: true
	Pem *string `json:"pem"`
}

// Validate validates this trusted c a certificate bean
func (m *TrustedCACertificateBean) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrustedCACertificateBean) validatePem(formats strfmt.Registry) error {

	if err := validate.Required("pem", "body", m.Pem); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrustedCACertificateBean) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrustedCACertificateBean) UnmarshalBinary(b []byte) error {
	var res TrustedCACertificateBean
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
