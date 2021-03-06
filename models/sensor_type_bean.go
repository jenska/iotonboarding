// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SensorTypeBean SensorTypeBean
//
// swagger:model SensorTypeBean
type SensorTypeBean struct {

	// The alternate ID of the sensor type. This field is unique across application tenant.
	AlternateID string `json:"alternateId,omitempty"`

	// The array of capability references. A capability is either of type measure or command.
	// Required: true
	Capabilities []*SensorTypeBeanCapabilitiesItems0 `json:"capabilities"`

	// A unique identifier of a sensor type. Is generated by the system and unique across platform tenant. This field is an immutable.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The name of the sensor type
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this sensor type bean
func (m *SensorTypeBean) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SensorTypeBean) validateCapabilities(formats strfmt.Registry) error {

	if err := validate.Required("capabilities", "body", m.Capabilities); err != nil {
		return err
	}

	for i := 0; i < len(m.Capabilities); i++ {
		if swag.IsZero(m.Capabilities[i]) { // not required
			continue
		}

		if m.Capabilities[i] != nil {
			if err := m.Capabilities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("capabilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SensorTypeBean) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SensorTypeBean) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SensorTypeBean) UnmarshalBinary(b []byte) error {
	var res SensorTypeBean
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SensorTypeBeanCapabilitiesItems0 CapabilityAssignmentBean
//
// swagger:model SensorTypeBeanCapabilitiesItems0
type SensorTypeBeanCapabilitiesItems0 struct {

	// References the unique identifier of a capability
	// Required: true
	ID *string `json:"id"`

	// The type of the referenced capability
	// Required: true
	// Enum: [measure command]
	Type *string `json:"type"`
}

// Validate validates this sensor type bean capabilities items0
func (m *SensorTypeBeanCapabilitiesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
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

func (m *SensorTypeBeanCapabilitiesItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var sensorTypeBeanCapabilitiesItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["measure","command"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sensorTypeBeanCapabilitiesItems0TypeTypePropEnum = append(sensorTypeBeanCapabilitiesItems0TypeTypePropEnum, v)
	}
}

const (

	// SensorTypeBeanCapabilitiesItems0TypeMeasure captures enum value "measure"
	SensorTypeBeanCapabilitiesItems0TypeMeasure string = "measure"

	// SensorTypeBeanCapabilitiesItems0TypeCommand captures enum value "command"
	SensorTypeBeanCapabilitiesItems0TypeCommand string = "command"
)

// prop value enum
func (m *SensorTypeBeanCapabilitiesItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sensorTypeBeanCapabilitiesItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SensorTypeBeanCapabilitiesItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SensorTypeBeanCapabilitiesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SensorTypeBeanCapabilitiesItems0) UnmarshalBinary(b []byte) error {
	var res SensorTypeBeanCapabilitiesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
