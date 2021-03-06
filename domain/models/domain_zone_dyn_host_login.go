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

// DomainZoneDynHostLogin Manage DynHost login
// swagger:model domain.Zone.DynHostLogin
type DomainZoneDynHostLogin struct {

	// Login
	// Required: true
	// Read Only: true
	Login string `json:"login"`

	// Subdomain that the login will be allowed to update (* to allow all)
	SubDomain string `json:"subDomain,omitempty"`

	// Zone
	// Required: true
	// Read Only: true
	Zone string `json:"zone"`
}

// Validate validates this domain zone dyn host login
func (m *DomainZoneDynHostLogin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogin(formats); err != nil {
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

func (m *DomainZoneDynHostLogin) validateLogin(formats strfmt.Registry) error {

	if err := validate.RequiredString("login", "body", string(m.Login)); err != nil {
		return err
	}

	return nil
}

func (m *DomainZoneDynHostLogin) validateZone(formats strfmt.Registry) error {

	if err := validate.RequiredString("zone", "body", string(m.Zone)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainZoneDynHostLogin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainZoneDynHostLogin) UnmarshalBinary(b []byte) error {
	var res DomainZoneDynHostLogin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
