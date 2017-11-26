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

// DebtEntryAssociatedObject The object linked to this debt entry
// swagger:model debt.Entry.AssociatedObject
type DebtEntryAssociatedObject struct {

	// id
	ID string `json:"id,omitempty"`

	// payment info
	PaymentInfo *DebtAssociatedObjectPaymentInfo `json:"paymentInfo,omitempty"`

	// sub Id
	SubID string `json:"subId,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this debt entry associated object
func (m *DebtEntryAssociatedObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentInfo(formats); err != nil {
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

func (m *DebtEntryAssociatedObject) validatePaymentInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentInfo) { // not required
		return nil
	}

	if m.PaymentInfo != nil {

		if err := m.PaymentInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("paymentInfo")
			}
			return err
		}
	}

	return nil
}

var debtEntryAssociatedObjectTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Bill","DebtOperation","Deposit","Order","OvhAccountMovement","Refund","Withdrawal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		debtEntryAssociatedObjectTypeTypePropEnum = append(debtEntryAssociatedObjectTypeTypePropEnum, v)
	}
}

const (
	// DebtEntryAssociatedObjectTypeBill captures enum value "Bill"
	DebtEntryAssociatedObjectTypeBill string = "Bill"
	// DebtEntryAssociatedObjectTypeDebtOperation captures enum value "DebtOperation"
	DebtEntryAssociatedObjectTypeDebtOperation string = "DebtOperation"
	// DebtEntryAssociatedObjectTypeDeposit captures enum value "Deposit"
	DebtEntryAssociatedObjectTypeDeposit string = "Deposit"
	// DebtEntryAssociatedObjectTypeOrder captures enum value "Order"
	DebtEntryAssociatedObjectTypeOrder string = "Order"
	// DebtEntryAssociatedObjectTypeOvhAccountMovement captures enum value "OvhAccountMovement"
	DebtEntryAssociatedObjectTypeOvhAccountMovement string = "OvhAccountMovement"
	// DebtEntryAssociatedObjectTypeRefund captures enum value "Refund"
	DebtEntryAssociatedObjectTypeRefund string = "Refund"
	// DebtEntryAssociatedObjectTypeWithdrawal captures enum value "Withdrawal"
	DebtEntryAssociatedObjectTypeWithdrawal string = "Withdrawal"
)

// prop value enum
func (m *DebtEntryAssociatedObject) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, debtEntryAssociatedObjectTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DebtEntryAssociatedObject) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DebtEntryAssociatedObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DebtEntryAssociatedObject) UnmarshalBinary(b []byte) error {
	var res DebtEntryAssociatedObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
