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
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PcaFile cloud archives files
// swagger:model pca.File
type PcaFile struct {

	// File MD5 hash
	// Required: true
	// Read Only: true
	MD5 string `json:"MD5"`

	// File SHA1 hash
	// Required: true
	// Read Only: true
	SHA1 string `json:"SHA1"`

	// File SHA256 hash
	// Required: true
	// Read Only: true
	SHA256 string `json:"SHA256"`

	// File id
	// Required: true
	// Read Only: true
	ID string `json:"id"`

	// File name
	// Required: true
	// Read Only: true
	Name string `json:"name"`

	// File size, in bytes
	// Required: true
	// Read Only: true
	Size int64 `json:"size"`

	// File state
	// Required: true
	// Read Only: true
	State string `json:"state"`

	// File type
	// Required: true
	// Read Only: true
	Type string `json:"type"`
}

// Validate validates this pca file
func (m *PcaFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMD5(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSHA1(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSHA256(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *PcaFile) validateMD5(formats strfmt.Registry) error {

	if err := validate.RequiredString("MD5", "body", string(m.MD5)); err != nil {
		return err
	}

	return nil
}

func (m *PcaFile) validateSHA1(formats strfmt.Registry) error {

	if err := validate.RequiredString("SHA1", "body", string(m.SHA1)); err != nil {
		return err
	}

	return nil
}

func (m *PcaFile) validateSHA256(formats strfmt.Registry) error {

	if err := validate.RequiredString("SHA256", "body", string(m.SHA256)); err != nil {
		return err
	}

	return nil
}

func (m *PcaFile) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *PcaFile) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *PcaFile) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", int64(m.Size)); err != nil {
		return err
	}

	return nil
}

func (m *PcaFile) validateState(formats strfmt.Registry) error {

	if err := validate.RequiredString("state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

func (m *PcaFile) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PcaFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PcaFile) UnmarshalBinary(b []byte) error {
	var res PcaFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
