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

// DebtDebt State of a debt
// swagger:model debt.Debt
type DebtDebt struct {

	// amount
	// Required: true
	Amount *OrderPrice `json:"amount"`

	// Date the debt was created on
	// Required: true
	// Read Only: true
	Date strfmt.DateTime `json:"date"`

	// debt Id
	// Required: true
	// Read Only: true
	DebtID int64 `json:"debtId"`

	// due amount
	// Required: true
	DueAmount *OrderPrice `json:"dueAmount"`

	// If specified, the debt will not be recovered before that date
	// Read Only: true
	DueDate strfmt.DateTime `json:"dueDate,omitempty"`

	// The order the debt relates to
	// Required: true
	// Read Only: true
	OrderID int64 `json:"orderId"`

	// pending amount
	// Required: true
	PendingAmount *OrderPrice `json:"pendingAmount"`

	// todo amount
	// Required: true
	TodoAmount *OrderPrice `json:"todoAmount"`

	// unmatured amount
	// Required: true
	UnmaturedAmount *OrderPrice `json:"unmaturedAmount"`
}

// Validate validates this debt debt
func (m *DebtDebt) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDebtID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDueAmount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrderID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePendingAmount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTodoAmount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUnmaturedAmount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DebtDebt) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	if m.Amount != nil {

		if err := m.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amount")
			}
			return err
		}
	}

	return nil
}

func (m *DebtDebt) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", strfmt.DateTime(m.Date)); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DebtDebt) validateDebtID(formats strfmt.Registry) error {

	if err := validate.Required("debtId", "body", int64(m.DebtID)); err != nil {
		return err
	}

	return nil
}

func (m *DebtDebt) validateDueAmount(formats strfmt.Registry) error {

	if err := validate.Required("dueAmount", "body", m.DueAmount); err != nil {
		return err
	}

	if m.DueAmount != nil {

		if err := m.DueAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dueAmount")
			}
			return err
		}
	}

	return nil
}

func (m *DebtDebt) validateOrderID(formats strfmt.Registry) error {

	if err := validate.Required("orderId", "body", int64(m.OrderID)); err != nil {
		return err
	}

	return nil
}

func (m *DebtDebt) validatePendingAmount(formats strfmt.Registry) error {

	if err := validate.Required("pendingAmount", "body", m.PendingAmount); err != nil {
		return err
	}

	if m.PendingAmount != nil {

		if err := m.PendingAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pendingAmount")
			}
			return err
		}
	}

	return nil
}

func (m *DebtDebt) validateTodoAmount(formats strfmt.Registry) error {

	if err := validate.Required("todoAmount", "body", m.TodoAmount); err != nil {
		return err
	}

	if m.TodoAmount != nil {

		if err := m.TodoAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("todoAmount")
			}
			return err
		}
	}

	return nil
}

func (m *DebtDebt) validateUnmaturedAmount(formats strfmt.Registry) error {

	if err := validate.Required("unmaturedAmount", "body", m.UnmaturedAmount); err != nil {
		return err
	}

	if m.UnmaturedAmount != nil {

		if err := m.UnmaturedAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unmaturedAmount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DebtDebt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DebtDebt) UnmarshalBinary(b []byte) error {
	var res DebtDebt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}