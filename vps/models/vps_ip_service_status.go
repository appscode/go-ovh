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

// VpsIPServiceStatus Service states for an Ip
// swagger:model vps.Ip.ServiceStatus
type VpsIPServiceStatus struct {

	// dns
	DNS *VpsIPServiceStatusService `json:"dns,omitempty"`

	// http
	HTTP *VpsIPServiceStatusService `json:"http,omitempty"`

	// https
	HTTPS *VpsIPServiceStatusService `json:"https,omitempty"`

	// ping
	Ping string `json:"ping,omitempty"`

	// smtp
	SMTP *VpsIPServiceStatusService `json:"smtp,omitempty"`

	// ssh
	SSH *VpsIPServiceStatusService `json:"ssh,omitempty"`

	// tools
	Tools string `json:"tools,omitempty"`
}

// Validate validates this vps Ip service status
func (m *VpsIPServiceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDNS(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHTTP(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHTTPS(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePing(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSMTP(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSSH(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTools(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VpsIPServiceStatus) validateDNS(formats strfmt.Registry) error {

	if swag.IsZero(m.DNS) { // not required
		return nil
	}

	if m.DNS != nil {

		if err := m.DNS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns")
			}
			return err
		}
	}

	return nil
}

func (m *VpsIPServiceStatus) validateHTTP(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTP) { // not required
		return nil
	}

	if m.HTTP != nil {

		if err := m.HTTP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http")
			}
			return err
		}
	}

	return nil
}

func (m *VpsIPServiceStatus) validateHTTPS(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPS) { // not required
		return nil
	}

	if m.HTTPS != nil {

		if err := m.HTTPS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("https")
			}
			return err
		}
	}

	return nil
}

var vpsIpServiceStatusTypePingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["down","up"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vpsIpServiceStatusTypePingPropEnum = append(vpsIpServiceStatusTypePingPropEnum, v)
	}
}

const (
	// VpsIPServiceStatusPingDown captures enum value "down"
	VpsIPServiceStatusPingDown string = "down"
	// VpsIPServiceStatusPingUp captures enum value "up"
	VpsIPServiceStatusPingUp string = "up"
)

// prop value enum
func (m *VpsIPServiceStatus) validatePingEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, vpsIpServiceStatusTypePingPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VpsIPServiceStatus) validatePing(formats strfmt.Registry) error {

	if swag.IsZero(m.Ping) { // not required
		return nil
	}

	// value enum
	if err := m.validatePingEnum("ping", "body", m.Ping); err != nil {
		return err
	}

	return nil
}

func (m *VpsIPServiceStatus) validateSMTP(formats strfmt.Registry) error {

	if swag.IsZero(m.SMTP) { // not required
		return nil
	}

	if m.SMTP != nil {

		if err := m.SMTP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smtp")
			}
			return err
		}
	}

	return nil
}

func (m *VpsIPServiceStatus) validateSSH(formats strfmt.Registry) error {

	if swag.IsZero(m.SSH) { // not required
		return nil
	}

	if m.SSH != nil {

		if err := m.SSH.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ssh")
			}
			return err
		}
	}

	return nil
}

var vpsIpServiceStatusTypeToolsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["down","up"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vpsIpServiceStatusTypeToolsPropEnum = append(vpsIpServiceStatusTypeToolsPropEnum, v)
	}
}

const (
	// VpsIPServiceStatusToolsDown captures enum value "down"
	VpsIPServiceStatusToolsDown string = "down"
	// VpsIPServiceStatusToolsUp captures enum value "up"
	VpsIPServiceStatusToolsUp string = "up"
)

// prop value enum
func (m *VpsIPServiceStatus) validateToolsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, vpsIpServiceStatusTypeToolsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VpsIPServiceStatus) validateTools(formats strfmt.Registry) error {

	if swag.IsZero(m.Tools) { // not required
		return nil
	}

	// value enum
	if err := m.validateToolsEnum("tools", "body", m.Tools); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VpsIPServiceStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VpsIPServiceStatus) UnmarshalBinary(b []byte) error {
	var res VpsIPServiceStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
