// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017 The go-ovh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ZoneResetRecord Resource record
// swagger:model zone.ResetRecord
type ZoneResetRecord struct {

	// field type
	FieldType string `json:"fieldType,omitempty"`

	// Resource record target
	Target string `json:"target,omitempty"`
}

// Validate validates this zone reset record
func (m *ZoneResetRecord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFieldType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var zoneResetRecordTypeFieldTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["A","MX"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		zoneResetRecordTypeFieldTypePropEnum = append(zoneResetRecordTypeFieldTypePropEnum, v)
	}
}

const (
	// ZoneResetRecordFieldTypeA captures enum value "A"
	ZoneResetRecordFieldTypeA string = "A"
	// ZoneResetRecordFieldTypeMX captures enum value "MX"
	ZoneResetRecordFieldTypeMX string = "MX"
)

// prop value enum
func (m *ZoneResetRecord) validateFieldTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, zoneResetRecordTypeFieldTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ZoneResetRecord) validateFieldType(formats strfmt.Registry) error {

	if swag.IsZero(m.FieldType) { // not required
		return nil
	}

	// value enum
	if err := m.validateFieldTypeEnum("fieldType", "body", m.FieldType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ZoneResetRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZoneResetRecord) UnmarshalBinary(b []byte) error {
	var res ZoneResetRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
