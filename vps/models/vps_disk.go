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

// VpsDisk Information about a disk of a VPS Virtual Machine
// swagger:model vps.Disk
type VpsDisk struct {

	// bandwidth limit
	// Required: true
	// Read Only: true
	BandwidthLimit int64 `json:"bandwidthLimit"`

	// id
	// Required: true
	// Read Only: true
	ID int64 `json:"id"`

	// The low disk free space threshold in MiB
	LowFreeSpaceThreshold int64 `json:"lowFreeSpaceThreshold,omitempty"`

	// The monitoring state of this disk
	Monitoring bool `json:"monitoring,omitempty"`

	// size
	// Required: true
	// Read Only: true
	Size int64 `json:"size"`

	// state
	// Required: true
	// Read Only: true
	State string `json:"state"`

	// type
	// Required: true
	// Read Only: true
	Type string `json:"type"`
}

// Validate validates this vps disk
func (m *VpsDisk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBandwidthLimit(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
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

func (m *VpsDisk) validateBandwidthLimit(formats strfmt.Registry) error {

	if err := validate.Required("bandwidthLimit", "body", int64(m.BandwidthLimit)); err != nil {
		return err
	}

	return nil
}

func (m *VpsDisk) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *VpsDisk) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", int64(m.Size)); err != nil {
		return err
	}

	return nil
}

var vpsDiskTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["connected","disconnected","pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vpsDiskTypeStatePropEnum = append(vpsDiskTypeStatePropEnum, v)
	}
}

const (
	// VpsDiskStateConnected captures enum value "connected"
	VpsDiskStateConnected string = "connected"
	// VpsDiskStateDisconnected captures enum value "disconnected"
	VpsDiskStateDisconnected string = "disconnected"
	// VpsDiskStatePending captures enum value "pending"
	VpsDiskStatePending string = "pending"
)

// prop value enum
func (m *VpsDisk) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, vpsDiskTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VpsDisk) validateState(formats strfmt.Registry) error {

	if err := validate.RequiredString("state", "body", string(m.State)); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

var vpsDiskTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["additional","primary"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vpsDiskTypeTypePropEnum = append(vpsDiskTypeTypePropEnum, v)
	}
}

const (
	// VpsDiskTypeAdditional captures enum value "additional"
	VpsDiskTypeAdditional string = "additional"
	// VpsDiskTypePrimary captures enum value "primary"
	VpsDiskTypePrimary string = "primary"
)

// prop value enum
func (m *VpsDisk) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, vpsDiskTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VpsDisk) validateType(formats strfmt.Registry) error {

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
func (m *VpsDisk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VpsDisk) UnmarshalBinary(b []byte) error {
	var res VpsDisk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
