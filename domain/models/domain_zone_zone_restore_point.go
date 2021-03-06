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

// DomainZoneZoneRestorePoint Zone restore point
// swagger:model domain.Zone.ZoneRestorePoint
type DomainZoneZoneRestorePoint struct {

	// Date of backup creation
	// Required: true
	// Read Only: true
	CreationDate strfmt.DateTime `json:"creationDate"`

	// URL to get backup content
	// Required: true
	// Read Only: true
	ZoneFileURL string `json:"zoneFileUrl"`
}

// Validate validates this domain zone zone restore point
func (m *DomainZoneZoneRestorePoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateZoneFileURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainZoneZoneRestorePoint) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", strfmt.DateTime(m.CreationDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainZoneZoneRestorePoint) validateZoneFileURL(formats strfmt.Registry) error {

	if err := validate.RequiredString("zoneFileUrl", "body", string(m.ZoneFileURL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainZoneZoneRestorePoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainZoneZoneRestorePoint) UnmarshalBinary(b []byte) error {
	var res DomainZoneZoneRestorePoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
