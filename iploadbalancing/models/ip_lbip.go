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

// IPLBIP Your IP load balancing
// swagger:model iplb.Ip
type IPLBIP struct {

	// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
	DisplayName string `json:"displayName,omitempty"`

	// Your IP load balancing
	// Required: true
	// Read Only: true
	IPLB string `json:"iplb"`

	// The IPV4 associated to your IP load balancing
	// Read Only: true
	IPV4 string `json:"ipv4,omitempty"`

	// The IPV6 associated to your IP load balancing
	// Read Only: true
	IPV6 string `json:"ipv6,omitempty"`

	// The metrics token associated with your IP load balancing
	// Read Only: true
	MetricsToken string `json:"metricsToken,omitempty"`

	// The offer of your IP load balancing
	// Required: true
	// Read Only: true
	Offer string `json:"offer"`

	// The internal name of your IP load balancing
	// Required: true
	// Read Only: true
	ServiceName string `json:"serviceName"`

	// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null.
	SslConfiguration string `json:"sslConfiguration,omitempty"`

	// Current state of your IP
	// Required: true
	// Read Only: true
	State string `json:"state"`

	// Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product
	// Read Only: true
	VrackName string `json:"vrackName,omitempty"`

	// Location where your service is
	// Required: true
	// Read Only: true
	Zone []string `json:"zone"`
}

// Validate validates this iplb Ip
func (m *IPLBIP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPLB(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOffer(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServiceName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSslConfiguration(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateZone(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPLBIP) validateIPLB(formats strfmt.Registry) error {

	if err := validate.RequiredString("iplb", "body", string(m.IPLB)); err != nil {
		return err
	}

	return nil
}

func (m *IPLBIP) validateOffer(formats strfmt.Registry) error {

	if err := validate.RequiredString("offer", "body", string(m.Offer)); err != nil {
		return err
	}

	return nil
}

func (m *IPLBIP) validateServiceName(formats strfmt.Registry) error {

	if err := validate.RequiredString("serviceName", "body", string(m.ServiceName)); err != nil {
		return err
	}

	return nil
}

var iplbIpTypeSslConfigurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["intermediate","modern"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iplbIpTypeSslConfigurationPropEnum = append(iplbIpTypeSslConfigurationPropEnum, v)
	}
}

const (
	// IPLBIPSslConfigurationIntermediate captures enum value "intermediate"
	IPLBIPSslConfigurationIntermediate string = "intermediate"
	// IPLBIPSslConfigurationModern captures enum value "modern"
	IPLBIPSslConfigurationModern string = "modern"
)

// prop value enum
func (m *IPLBIP) validateSslConfigurationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, iplbIpTypeSslConfigurationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IPLBIP) validateSslConfiguration(formats strfmt.Registry) error {

	if swag.IsZero(m.SslConfiguration) { // not required
		return nil
	}

	// value enum
	if err := m.validateSslConfigurationEnum("sslConfiguration", "body", m.SslConfiguration); err != nil {
		return err
	}

	return nil
}

var iplbIpTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["blacklisted","deleted","free","ok","quarantined","suspended"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iplbIpTypeStatePropEnum = append(iplbIpTypeStatePropEnum, v)
	}
}

const (
	// IPLBIPStateBlacklisted captures enum value "blacklisted"
	IPLBIPStateBlacklisted string = "blacklisted"
	// IPLBIPStateDeleted captures enum value "deleted"
	IPLBIPStateDeleted string = "deleted"
	// IPLBIPStateFree captures enum value "free"
	IPLBIPStateFree string = "free"
	// IPLBIPStateOk captures enum value "ok"
	IPLBIPStateOk string = "ok"
	// IPLBIPStateQuarantined captures enum value "quarantined"
	IPLBIPStateQuarantined string = "quarantined"
	// IPLBIPStateSuspended captures enum value "suspended"
	IPLBIPStateSuspended string = "suspended"
)

// prop value enum
func (m *IPLBIP) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, iplbIpTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IPLBIP) validateState(formats strfmt.Registry) error {

	if err := validate.RequiredString("state", "body", string(m.State)); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *IPLBIP) validateZone(formats strfmt.Registry) error {

	if err := validate.Required("zone", "body", m.Zone); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPLBIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPLBIP) UnmarshalBinary(b []byte) error {
	var res IPLBIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
