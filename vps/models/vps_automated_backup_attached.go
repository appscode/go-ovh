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
)

// VpsAutomatedBackupAttached A backup attached to your VPS
// swagger:model vps.AutomatedBackup.Attached
type VpsAutomatedBackupAttached struct {

	// access
	Access *VpsAutomatedBackupAttachedInfos `json:"access,omitempty"`

	// restore point
	RestorePoint strfmt.DateTime `json:"restorePoint,omitempty"`
}

// Validate validates this vps automated backup attached
func (m *VpsAutomatedBackupAttached) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VpsAutomatedBackupAttached) validateAccess(formats strfmt.Registry) error {

	if swag.IsZero(m.Access) { // not required
		return nil
	}

	if m.Access != nil {

		if err := m.Access.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("access")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VpsAutomatedBackupAttached) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VpsAutomatedBackupAttached) UnmarshalBinary(b []byte) error {
	var res VpsAutomatedBackupAttached
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
