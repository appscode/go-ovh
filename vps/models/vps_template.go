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

// VpsTemplate Installation template for a VPS Virtual Machine
// swagger:model vps.Template
type VpsTemplate struct {

	// available language
	// Required: true
	// Read Only: true
	AvailableLanguage []string `json:"availableLanguage"`

	// bit format
	// Required: true
	// Read Only: true
	BitFormat int64 `json:"bitFormat"`

	// distribution
	// Required: true
	// Read Only: true
	Distribution string `json:"distribution"`

	// id
	// Required: true
	// Read Only: true
	ID int64 `json:"id"`

	// locale
	// Required: true
	// Read Only: true
	Locale string `json:"locale"`

	// name
	// Required: true
	// Read Only: true
	Name string `json:"name"`
}

// Validate validates this vps template
func (m *VpsTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableLanguage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBitFormat(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDistribution(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLocale(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VpsTemplate) validateAvailableLanguage(formats strfmt.Registry) error {

	if err := validate.Required("availableLanguage", "body", m.AvailableLanguage); err != nil {
		return err
	}

	return nil
}

var vpsTemplateTypeBitFormatPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[32,64]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vpsTemplateTypeBitFormatPropEnum = append(vpsTemplateTypeBitFormatPropEnum, v)
	}
}

// prop value enum
func (m *VpsTemplate) validateBitFormatEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, vpsTemplateTypeBitFormatPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VpsTemplate) validateBitFormat(formats strfmt.Registry) error {

	if err := validate.Required("bitFormat", "body", int64(m.BitFormat)); err != nil {
		return err
	}

	// value enum
	if err := m.validateBitFormatEnum("bitFormat", "body", m.BitFormat); err != nil {
		return err
	}

	return nil
}

func (m *VpsTemplate) validateDistribution(formats strfmt.Registry) error {

	if err := validate.RequiredString("distribution", "body", string(m.Distribution)); err != nil {
		return err
	}

	return nil
}

func (m *VpsTemplate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *VpsTemplate) validateLocale(formats strfmt.Registry) error {

	if err := validate.RequiredString("locale", "body", string(m.Locale)); err != nil {
		return err
	}

	return nil
}

func (m *VpsTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VpsTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VpsTemplate) UnmarshalBinary(b []byte) error {
	var res VpsTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}