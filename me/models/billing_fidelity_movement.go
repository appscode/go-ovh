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

// BillingFidelityMovement Details about a fidelity account
// swagger:model billing.FidelityMovement
type BillingFidelityMovement struct {

	// amount
	// Required: true
	// Read Only: true
	Amount int64 `json:"amount"`

	// balance
	// Required: true
	// Read Only: true
	Balance int64 `json:"balance"`

	// date
	// Required: true
	// Read Only: true
	Date strfmt.DateTime `json:"date"`

	// description
	// Required: true
	// Read Only: true
	Description string `json:"description"`

	// movement Id
	// Required: true
	// Read Only: true
	MovementID int64 `json:"movementId"`

	// operation
	// Required: true
	// Read Only: true
	Operation string `json:"operation"`

	// order
	// Required: true
	// Read Only: true
	Order int64 `json:"order"`

	// previous balance
	// Required: true
	// Read Only: true
	PreviousBalance int64 `json:"previousBalance"`
}

// Validate validates this billing fidelity movement
func (m *BillingFidelityMovement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBalance(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMovementID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOperation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrder(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePreviousBalance(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingFidelityMovement) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", int64(m.Amount)); err != nil {
		return err
	}

	return nil
}

func (m *BillingFidelityMovement) validateBalance(formats strfmt.Registry) error {

	if err := validate.Required("balance", "body", int64(m.Balance)); err != nil {
		return err
	}

	return nil
}

func (m *BillingFidelityMovement) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", strfmt.DateTime(m.Date)); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BillingFidelityMovement) validateDescription(formats strfmt.Registry) error {

	if err := validate.RequiredString("description", "body", string(m.Description)); err != nil {
		return err
	}

	return nil
}

func (m *BillingFidelityMovement) validateMovementID(formats strfmt.Registry) error {

	if err := validate.Required("movementId", "body", int64(m.MovementID)); err != nil {
		return err
	}

	return nil
}

var billingFidelityMovementTypeOperationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bonus","cancel-bonus","cancel-credit","cancel-debit","cancel-pre-debit","credit","debit","pre-credit","pre-debit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billingFidelityMovementTypeOperationPropEnum = append(billingFidelityMovementTypeOperationPropEnum, v)
	}
}

const (
	// BillingFidelityMovementOperationBonus captures enum value "bonus"
	BillingFidelityMovementOperationBonus string = "bonus"
	// BillingFidelityMovementOperationCancelBonus captures enum value "cancel-bonus"
	BillingFidelityMovementOperationCancelBonus string = "cancel-bonus"
	// BillingFidelityMovementOperationCancelCredit captures enum value "cancel-credit"
	BillingFidelityMovementOperationCancelCredit string = "cancel-credit"
	// BillingFidelityMovementOperationCancelDebit captures enum value "cancel-debit"
	BillingFidelityMovementOperationCancelDebit string = "cancel-debit"
	// BillingFidelityMovementOperationCancelPreDebit captures enum value "cancel-pre-debit"
	BillingFidelityMovementOperationCancelPreDebit string = "cancel-pre-debit"
	// BillingFidelityMovementOperationCredit captures enum value "credit"
	BillingFidelityMovementOperationCredit string = "credit"
	// BillingFidelityMovementOperationDebit captures enum value "debit"
	BillingFidelityMovementOperationDebit string = "debit"
	// BillingFidelityMovementOperationPreCredit captures enum value "pre-credit"
	BillingFidelityMovementOperationPreCredit string = "pre-credit"
	// BillingFidelityMovementOperationPreDebit captures enum value "pre-debit"
	BillingFidelityMovementOperationPreDebit string = "pre-debit"
)

// prop value enum
func (m *BillingFidelityMovement) validateOperationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, billingFidelityMovementTypeOperationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BillingFidelityMovement) validateOperation(formats strfmt.Registry) error {

	if err := validate.RequiredString("operation", "body", string(m.Operation)); err != nil {
		return err
	}

	// value enum
	if err := m.validateOperationEnum("operation", "body", m.Operation); err != nil {
		return err
	}

	return nil
}

func (m *BillingFidelityMovement) validateOrder(formats strfmt.Registry) error {

	if err := validate.Required("order", "body", int64(m.Order)); err != nil {
		return err
	}

	return nil
}

func (m *BillingFidelityMovement) validatePreviousBalance(formats strfmt.Registry) error {

	if err := validate.Required("previousBalance", "body", int64(m.PreviousBalance)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillingFidelityMovement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingFidelityMovement) UnmarshalBinary(b []byte) error {
	var res BillingFidelityMovement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
