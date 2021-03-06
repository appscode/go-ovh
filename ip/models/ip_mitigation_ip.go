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

// IPMitigationIP Your IP on mitigation
// swagger:model ip.MitigationIp
type IPMitigationIP struct {

	// Set on true if your ip is on auto-mitigation
	// Required: true
	// Read Only: true
	Auto bool `json:"auto"`

	// ip on mitigation
	// Required: true
	// Read Only: true
	IPOnMitigation string `json:"ipOnMitigation"`

	// Set on true if your ip is on permanent mitigation
	Permanent bool `json:"permanent,omitempty"`

	// Current state of your ip on mitigation
	// Required: true
	// Read Only: true
	State string `json:"state"`
}

// Validate validates this ip mitigation Ip
func (m *IPMitigationIP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuto(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPOnMitigation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPMitigationIP) validateAuto(formats strfmt.Registry) error {

	if err := validate.Required("auto", "body", bool(m.Auto)); err != nil {
		return err
	}

	return nil
}

func (m *IPMitigationIP) validateIPOnMitigation(formats strfmt.Registry) error {

	if err := validate.RequiredString("ipOnMitigation", "body", string(m.IPOnMitigation)); err != nil {
		return err
	}

	return nil
}

var ipMitigationIpTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["creationPending","ok","removalPending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipMitigationIpTypeStatePropEnum = append(ipMitigationIpTypeStatePropEnum, v)
	}
}

const (
	// IPMitigationIPStateCreationPending captures enum value "creationPending"
	IPMitigationIPStateCreationPending string = "creationPending"
	// IPMitigationIPStateOk captures enum value "ok"
	IPMitigationIPStateOk string = "ok"
	// IPMitigationIPStateRemovalPending captures enum value "removalPending"
	IPMitigationIPStateRemovalPending string = "removalPending"
)

// prop value enum
func (m *IPMitigationIP) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ipMitigationIpTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IPMitigationIP) validateState(formats strfmt.Registry) error {

	if err := validate.RequiredString("state", "body", string(m.State)); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPMitigationIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPMitigationIP) UnmarshalBinary(b []byte) error {
	var res IPMitigationIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
