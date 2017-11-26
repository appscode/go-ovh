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

// IPFirewallIP Your IP on firewall
// swagger:model ip.FirewallIp
type IPFirewallIP struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// ip on firewall
	// Required: true
	// Read Only: true
	IPOnFirewall string `json:"ipOnFirewall"`

	// Current state of your ip on firewall
	// Required: true
	// Read Only: true
	State string `json:"state"`
}

// Validate validates this ip firewall Ip
func (m *IPFirewallIP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPOnFirewall(formats); err != nil {
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

func (m *IPFirewallIP) validateIPOnFirewall(formats strfmt.Registry) error {

	if err := validate.RequiredString("ipOnFirewall", "body", string(m.IPOnFirewall)); err != nil {
		return err
	}

	return nil
}

var ipFirewallIpTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["disableFirewallPending","enableFirewallPending","ok"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipFirewallIpTypeStatePropEnum = append(ipFirewallIpTypeStatePropEnum, v)
	}
}

const (
	// IPFirewallIPStateDisableFirewallPending captures enum value "disableFirewallPending"
	IPFirewallIPStateDisableFirewallPending string = "disableFirewallPending"
	// IPFirewallIPStateEnableFirewallPending captures enum value "enableFirewallPending"
	IPFirewallIPStateEnableFirewallPending string = "enableFirewallPending"
	// IPFirewallIPStateOk captures enum value "ok"
	IPFirewallIPStateOk string = "ok"
)

// prop value enum
func (m *IPFirewallIP) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ipFirewallIpTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IPFirewallIP) validateState(formats strfmt.Registry) error {

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
func (m *IPFirewallIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPFirewallIP) UnmarshalBinary(b []byte) error {
	var res IPFirewallIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}