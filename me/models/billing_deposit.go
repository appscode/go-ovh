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

// BillingDeposit Details about a deposit
// swagger:model billing.Deposit
type BillingDeposit struct {

	// amount
	// Required: true
	Amount *OrderPrice `json:"amount"`

	// date
	// Required: true
	// Read Only: true
	Date strfmt.DateTime `json:"date"`

	// deposit Id
	// Required: true
	// Read Only: true
	DepositID string `json:"depositId"`

	// order Id
	// Required: true
	// Read Only: true
	OrderID int64 `json:"orderId"`

	// password
	// Required: true
	// Read Only: true
	Password string `json:"password"`

	// payment info
	PaymentInfo *DebtAssociatedObjectPaymentInfo `json:"paymentInfo,omitempty"`

	// pdf Url
	// Required: true
	// Read Only: true
	PdfURL string `json:"pdfUrl"`

	// url
	// Required: true
	// Read Only: true
	URL string `json:"url"`
}

// Validate validates this billing deposit
func (m *BillingDeposit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDepositID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrderID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePdfURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingDeposit) validateAmount(formats strfmt.Registry) error {

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

func (m *BillingDeposit) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", strfmt.DateTime(m.Date)); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BillingDeposit) validateDepositID(formats strfmt.Registry) error {

	if err := validate.RequiredString("depositId", "body", string(m.DepositID)); err != nil {
		return err
	}

	return nil
}

func (m *BillingDeposit) validateOrderID(formats strfmt.Registry) error {

	if err := validate.Required("orderId", "body", int64(m.OrderID)); err != nil {
		return err
	}

	return nil
}

func (m *BillingDeposit) validatePassword(formats strfmt.Registry) error {

	if err := validate.RequiredString("password", "body", string(m.Password)); err != nil {
		return err
	}

	return nil
}

func (m *BillingDeposit) validatePaymentInfo(formats strfmt.Registry) error {

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

func (m *BillingDeposit) validatePdfURL(formats strfmt.Registry) error {

	if err := validate.RequiredString("pdfUrl", "body", string(m.PdfURL)); err != nil {
		return err
	}

	return nil
}

func (m *BillingDeposit) validateURL(formats strfmt.Registry) error {

	if err := validate.RequiredString("url", "body", string(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillingDeposit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingDeposit) UnmarshalBinary(b []byte) error {
	var res BillingDeposit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
