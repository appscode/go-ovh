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

// BillingPayment Details about a payment
// swagger:model billing.Payment
type BillingPayment struct {

	// payment date
	// Required: true
	// Read Only: true
	PaymentDate strfmt.DateTime `json:"paymentDate"`

	// payment identifier
	// Read Only: true
	PaymentIdentifier string `json:"paymentIdentifier,omitempty"`

	// payment type
	// Required: true
	// Read Only: true
	PaymentType string `json:"paymentType"`
}

// Validate validates this billing payment
func (m *BillingPayment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingPayment) validatePaymentDate(formats strfmt.Registry) error {

	if err := validate.Required("paymentDate", "body", strfmt.DateTime(m.PaymentDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("paymentDate", "body", "date-time", m.PaymentDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var billingPaymentTypePaymentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cash","chargeback","cheque","creditCard","debtAccount","deposit","digitalLaunchPad","edinar","fidelityPoints","free","ideal","incubatorAccount","mandat","multibanco","none","ovhAccount","paymentMandate","paypal","payu","platnosci","refund","transfer","withdrawal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billingPaymentTypePaymentTypePropEnum = append(billingPaymentTypePaymentTypePropEnum, v)
	}
}

const (
	// BillingPaymentPaymentTypeCash captures enum value "cash"
	BillingPaymentPaymentTypeCash string = "cash"
	// BillingPaymentPaymentTypeChargeback captures enum value "chargeback"
	BillingPaymentPaymentTypeChargeback string = "chargeback"
	// BillingPaymentPaymentTypeCheque captures enum value "cheque"
	BillingPaymentPaymentTypeCheque string = "cheque"
	// BillingPaymentPaymentTypeCreditCard captures enum value "creditCard"
	BillingPaymentPaymentTypeCreditCard string = "creditCard"
	// BillingPaymentPaymentTypeDebtAccount captures enum value "debtAccount"
	BillingPaymentPaymentTypeDebtAccount string = "debtAccount"
	// BillingPaymentPaymentTypeDeposit captures enum value "deposit"
	BillingPaymentPaymentTypeDeposit string = "deposit"
	// BillingPaymentPaymentTypeDigitalLaunchPad captures enum value "digitalLaunchPad"
	BillingPaymentPaymentTypeDigitalLaunchPad string = "digitalLaunchPad"
	// BillingPaymentPaymentTypeEdinar captures enum value "edinar"
	BillingPaymentPaymentTypeEdinar string = "edinar"
	// BillingPaymentPaymentTypeFidelityPoints captures enum value "fidelityPoints"
	BillingPaymentPaymentTypeFidelityPoints string = "fidelityPoints"
	// BillingPaymentPaymentTypeFree captures enum value "free"
	BillingPaymentPaymentTypeFree string = "free"
	// BillingPaymentPaymentTypeIdeal captures enum value "ideal"
	BillingPaymentPaymentTypeIdeal string = "ideal"
	// BillingPaymentPaymentTypeIncubatorAccount captures enum value "incubatorAccount"
	BillingPaymentPaymentTypeIncubatorAccount string = "incubatorAccount"
	// BillingPaymentPaymentTypeMandat captures enum value "mandat"
	BillingPaymentPaymentTypeMandat string = "mandat"
	// BillingPaymentPaymentTypeMultibanco captures enum value "multibanco"
	BillingPaymentPaymentTypeMultibanco string = "multibanco"
	// BillingPaymentPaymentTypeNone captures enum value "none"
	BillingPaymentPaymentTypeNone string = "none"
	// BillingPaymentPaymentTypeOvhAccount captures enum value "ovhAccount"
	BillingPaymentPaymentTypeOvhAccount string = "ovhAccount"
	// BillingPaymentPaymentTypePaymentMandate captures enum value "paymentMandate"
	BillingPaymentPaymentTypePaymentMandate string = "paymentMandate"
	// BillingPaymentPaymentTypePaypal captures enum value "paypal"
	BillingPaymentPaymentTypePaypal string = "paypal"
	// BillingPaymentPaymentTypePayu captures enum value "payu"
	BillingPaymentPaymentTypePayu string = "payu"
	// BillingPaymentPaymentTypePlatnosci captures enum value "platnosci"
	BillingPaymentPaymentTypePlatnosci string = "platnosci"
	// BillingPaymentPaymentTypeRefund captures enum value "refund"
	BillingPaymentPaymentTypeRefund string = "refund"
	// BillingPaymentPaymentTypeTransfer captures enum value "transfer"
	BillingPaymentPaymentTypeTransfer string = "transfer"
	// BillingPaymentPaymentTypeWithdrawal captures enum value "withdrawal"
	BillingPaymentPaymentTypeWithdrawal string = "withdrawal"
)

// prop value enum
func (m *BillingPayment) validatePaymentTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, billingPaymentTypePaymentTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BillingPayment) validatePaymentType(formats strfmt.Registry) error {

	if err := validate.RequiredString("paymentType", "body", string(m.PaymentType)); err != nil {
		return err
	}

	// value enum
	if err := m.validatePaymentTypeEnum("paymentType", "body", m.PaymentType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillingPayment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingPayment) UnmarshalBinary(b []byte) error {
	var res BillingPayment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
