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

// Property Property
//
// swagger:model Property
type Property struct {

	// capability Id
	CapabilityID string `json:"capabilityId,omitempty"`

	// data type
	// Enum: [integer long float double boolean string binary date]
	DataType string `json:"dataType,omitempty"`

	// formatter
	Formatter *PropertyFormatter `json:"formatter,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// uom
	Uom string `json:"uom,omitempty"`
}

// Validate validates this property
func (m *Property) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormatter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var propertyTypeDataTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["integer","long","float","double","boolean","string","binary","date"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		propertyTypeDataTypePropEnum = append(propertyTypeDataTypePropEnum, v)
	}
}

const (

	// PropertyDataTypeInteger captures enum value "integer"
	PropertyDataTypeInteger string = "integer"

	// PropertyDataTypeLong captures enum value "long"
	PropertyDataTypeLong string = "long"

	// PropertyDataTypeFloat captures enum value "float"
	PropertyDataTypeFloat string = "float"

	// PropertyDataTypeDouble captures enum value "double"
	PropertyDataTypeDouble string = "double"

	// PropertyDataTypeBoolean captures enum value "boolean"
	PropertyDataTypeBoolean string = "boolean"

	// PropertyDataTypeString captures enum value "string"
	PropertyDataTypeString string = "string"

	// PropertyDataTypeBinary captures enum value "binary"
	PropertyDataTypeBinary string = "binary"

	// PropertyDataTypeDate captures enum value "date"
	PropertyDataTypeDate string = "date"
)

// prop value enum
func (m *Property) validateDataTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, propertyTypeDataTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Property) validateDataType(formats strfmt.Registry) error {

	if swag.IsZero(m.DataType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDataTypeEnum("dataType", "body", m.DataType); err != nil {
		return err
	}

	return nil
}

func (m *Property) validateFormatter(formats strfmt.Registry) error {

	if swag.IsZero(m.Formatter) { // not required
		return nil
	}

	if m.Formatter != nil {
		if err := m.Formatter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("formatter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Property) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Property) UnmarshalBinary(b []byte) error {
	var res Property
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PropertyFormatter NumberFormatter
//
// swagger:model PropertyFormatter
type PropertyFormatter struct {

	// data type
	// Enum: [integer long float double boolean string binary date]
	DataType string `json:"dataType,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// scale
	Scale int32 `json:"scale,omitempty"`

	// shift
	Shift float32 `json:"shift,omitempty"`

	// swap
	Swap bool `json:"swap,omitempty"`
}

// Validate validates this property formatter
func (m *PropertyFormatter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var propertyFormatterTypeDataTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["integer","long","float","double","boolean","string","binary","date"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		propertyFormatterTypeDataTypePropEnum = append(propertyFormatterTypeDataTypePropEnum, v)
	}
}

const (

	// PropertyFormatterDataTypeInteger captures enum value "integer"
	PropertyFormatterDataTypeInteger string = "integer"

	// PropertyFormatterDataTypeLong captures enum value "long"
	PropertyFormatterDataTypeLong string = "long"

	// PropertyFormatterDataTypeFloat captures enum value "float"
	PropertyFormatterDataTypeFloat string = "float"

	// PropertyFormatterDataTypeDouble captures enum value "double"
	PropertyFormatterDataTypeDouble string = "double"

	// PropertyFormatterDataTypeBoolean captures enum value "boolean"
	PropertyFormatterDataTypeBoolean string = "boolean"

	// PropertyFormatterDataTypeString captures enum value "string"
	PropertyFormatterDataTypeString string = "string"

	// PropertyFormatterDataTypeBinary captures enum value "binary"
	PropertyFormatterDataTypeBinary string = "binary"

	// PropertyFormatterDataTypeDate captures enum value "date"
	PropertyFormatterDataTypeDate string = "date"
)

// prop value enum
func (m *PropertyFormatter) validateDataTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, propertyFormatterTypeDataTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PropertyFormatter) validateDataType(formats strfmt.Registry) error {

	if swag.IsZero(m.DataType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDataTypeEnum("formatter"+"."+"dataType", "body", m.DataType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PropertyFormatter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PropertyFormatter) UnmarshalBinary(b []byte) error {
	var res PropertyFormatter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
