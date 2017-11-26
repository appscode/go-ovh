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

// BillingDepositDetail Information about a Deposit entry
// swagger:model billing.DepositDetail
type BillingDepositDetail struct {

	// deposit detail Id
	// Required: true
	// Read Only: true
	DepositDetailID string `json:"depositDetailId"`

	// description
	// Required: true
	// Read Only: true
	Description string `json:"description"`

	// domain
	// Required: true
	// Read Only: true
	Domain string `json:"domain"`

	// quantity
	// Required: true
	// Read Only: true
	Quantity string `json:"quantity"`

	// total price
	// Required: true
	TotalPrice *OrderPrice `json:"totalPrice"`

	// unit price
	// Required: true
	UnitPrice *OrderPrice `json:"unitPrice"`
}

// Validate validates this billing deposit detail
func (m *BillingDepositDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDepositDetailID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTotalPrice(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUnitPrice(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingDepositDetail) validateDepositDetailID(formats strfmt.Registry) error {

	if err := validate.RequiredString("depositDetailId", "body", string(m.DepositDetailID)); err != nil {
		return err
	}

	return nil
}

func (m *BillingDepositDetail) validateDescription(formats strfmt.Registry) error {

	if err := validate.RequiredString("description", "body", string(m.Description)); err != nil {
		return err
	}

	return nil
}

func (m *BillingDepositDetail) validateDomain(formats strfmt.Registry) error {

	if err := validate.RequiredString("domain", "body", string(m.Domain)); err != nil {
		return err
	}

	return nil
}

func (m *BillingDepositDetail) validateQuantity(formats strfmt.Registry) error {

	if err := validate.RequiredString("quantity", "body", string(m.Quantity)); err != nil {
		return err
	}

	return nil
}

func (m *BillingDepositDetail) validateTotalPrice(formats strfmt.Registry) error {

	if err := validate.Required("totalPrice", "body", m.TotalPrice); err != nil {
		return err
	}

	if m.TotalPrice != nil {

		if err := m.TotalPrice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("totalPrice")
			}
			return err
		}
	}

	return nil
}

func (m *BillingDepositDetail) validateUnitPrice(formats strfmt.Registry) error {

	if err := validate.Required("unitPrice", "body", m.UnitPrice); err != nil {
		return err
	}

	if m.UnitPrice != nil {

		if err := m.UnitPrice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unitPrice")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillingDepositDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingDepositDetail) UnmarshalBinary(b []byte) error {
	var res BillingDepositDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
