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

// DomainDataClaimNoticeContact Contact definition of a claim notice holder
// swagger:model domain.Data.ClaimNotice.Contact
type DomainDataClaimNoticeContact struct {

	// address
	// Required: true
	Address *DomainDataClaimNoticeAddress `json:"address"`

	// Email address
	// Read Only: true
	Email string `json:"email,omitempty"`

	// Legitimacy of holder
	// Read Only: true
	Entitlement string `json:"entitlement,omitempty"`

	// Name of claim notice holder
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Organisation name
	// Read Only: true
	Organisation string `json:"organisation,omitempty"`

	// Type of contact
	// Read Only: true
	Type string `json:"type,omitempty"`
}

// Validate validates this domain data claim notice contact
func (m *DomainDataClaimNoticeContact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainDataClaimNoticeContact) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	if m.Address != nil {

		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainDataClaimNoticeContact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainDataClaimNoticeContact) UnmarshalBinary(b []byte) error {
	var res DomainDataClaimNoticeContact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}