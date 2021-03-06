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

// TelephonyDefaultIPRestriction Default IP restriction of a VoIP line
// swagger:model telephony.DefaultIpRestriction
type TelephonyDefaultIPRestriction struct {

	// id
	// Required: true
	// Read Only: true
	ID int64 `json:"id"`

	// The IPv4 subnet you want to allow
	// Required: true
	// Read Only: true
	Subnet string `json:"subnet"`

	// The protocol you want to restrict (sip/mgcp)
	// Required: true
	// Read Only: true
	Type string `json:"type"`
}

// Validate validates this telephony default Ip restriction
func (m *TelephonyDefaultIPRestriction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TelephonyDefaultIPRestriction) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *TelephonyDefaultIPRestriction) validateSubnet(formats strfmt.Registry) error {

	if err := validate.RequiredString("subnet", "body", string(m.Subnet)); err != nil {
		return err
	}

	return nil
}

var telephonyDefaultIpRestrictionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["mgcp","sip"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		telephonyDefaultIpRestrictionTypeTypePropEnum = append(telephonyDefaultIpRestrictionTypeTypePropEnum, v)
	}
}

const (
	// TelephonyDefaultIPRestrictionTypeMgcp captures enum value "mgcp"
	TelephonyDefaultIPRestrictionTypeMgcp string = "mgcp"
	// TelephonyDefaultIPRestrictionTypeSip captures enum value "sip"
	TelephonyDefaultIPRestrictionTypeSip string = "sip"
)

// prop value enum
func (m *TelephonyDefaultIPRestriction) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, telephonyDefaultIpRestrictionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TelephonyDefaultIPRestriction) validateType(formats strfmt.Registry) error {

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
func (m *TelephonyDefaultIPRestriction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TelephonyDefaultIPRestriction) UnmarshalBinary(b []byte) error {
	var res TelephonyDefaultIPRestriction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
