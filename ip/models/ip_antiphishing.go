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

// IPAntiphishing Phishing URLs hosted on your IP
// swagger:model ip.Antiphishing
type IPAntiphishing struct {

	// Date of the event
	// Required: true
	// Read Only: true
	CreationDate strfmt.DateTime `json:"creationDate"`

	// Internal ID of the phishing entry
	// Required: true
	// Read Only: true
	ID int64 `json:"id"`

	// IP address hosting the phishing URL
	// Required: true
	// Read Only: true
	IPOnAntiphishing string `json:"ipOnAntiphishing"`

	// Current state of the phishing
	// Required: true
	// Read Only: true
	State string `json:"state"`

	// Phishing URL
	// Required: true
	// Read Only: true
	URLPhishing string `json:"urlPhishing"`
}

// Validate validates this ip antiphishing
func (m *IPAntiphishing) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPOnAntiphishing(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateURLPhishing(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPAntiphishing) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", strfmt.DateTime(m.CreationDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IPAntiphishing) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *IPAntiphishing) validateIPOnAntiphishing(formats strfmt.Registry) error {

	if err := validate.RequiredString("ipOnAntiphishing", "body", string(m.IPOnAntiphishing)); err != nil {
		return err
	}

	return nil
}

var ipAntiphishingTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["blocked","blocking","unblocked","unblocking"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipAntiphishingTypeStatePropEnum = append(ipAntiphishingTypeStatePropEnum, v)
	}
}

const (
	// IPAntiphishingStateBlocked captures enum value "blocked"
	IPAntiphishingStateBlocked string = "blocked"
	// IPAntiphishingStateBlocking captures enum value "blocking"
	IPAntiphishingStateBlocking string = "blocking"
	// IPAntiphishingStateUnblocked captures enum value "unblocked"
	IPAntiphishingStateUnblocked string = "unblocked"
	// IPAntiphishingStateUnblocking captures enum value "unblocking"
	IPAntiphishingStateUnblocking string = "unblocking"
)

// prop value enum
func (m *IPAntiphishing) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ipAntiphishingTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IPAntiphishing) validateState(formats strfmt.Registry) error {

	if err := validate.RequiredString("state", "body", string(m.State)); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *IPAntiphishing) validateURLPhishing(formats strfmt.Registry) error {

	if err := validate.RequiredString("urlPhishing", "body", string(m.URLPhishing)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPAntiphishing) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAntiphishing) UnmarshalBinary(b []byte) error {
	var res IPAntiphishing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
