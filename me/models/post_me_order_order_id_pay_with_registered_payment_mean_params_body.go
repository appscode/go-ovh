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

// PostMeOrderOrderIDPayWithRegisteredPaymentMeanParamsBody post me order order Id pay with registered payment mean params body
// swagger:model postMeOrderOrderIdPayWithRegisteredPaymentMeanParamsBody
type PostMeOrderOrderIDPayWithRegisteredPaymentMeanParamsBody struct {

	// payment mean
	// Required: true
	PaymentMean *string `json:"paymentMean"`

	// payment mean Id
	PaymentMeanID int64 `json:"paymentMeanId,omitempty"`
}

// Validate validates this post me order order Id pay with registered payment mean params body
func (m *PostMeOrderOrderIDPayWithRegisteredPaymentMeanParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentMean(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postMeOrderOrderIdPayWithRegisteredPaymentMeanParamsBodyTypePaymentMeanPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bankAccount","creditCard","fidelityAccount","ovhAccount","paypal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postMeOrderOrderIdPayWithRegisteredPaymentMeanParamsBodyTypePaymentMeanPropEnum = append(postMeOrderOrderIdPayWithRegisteredPaymentMeanParamsBodyTypePaymentMeanPropEnum, v)
	}
}

const (
	// PostMeOrderOrderIDPayWithRegisteredPaymentMeanParamsBodyPaymentMeanBankAccount captures enum value "bankAccount"
	PostMeOrderOrderIDPayWithRegisteredPaymentMeanParamsBodyPaymentMeanBankAccount string = "bankAccount"
	// PostMeOrderOrderIDPayWithRegisteredPaymentMeanParamsBodyPaymentMeanCreditCard captures enum value "creditCard"
	PostMeOrderOrderIDPayWithRegisteredPaymentMeanParamsBodyPaymentMeanCreditCard string = "creditCard"
	// PostMeOrderOrderIDPayWithRegisteredPaymentMeanParamsBodyPaymentMeanFidelityAccount captures enum value "fidelityAccount"
	PostMeOrderOrderIDPayWithRegisteredPaymentMeanParamsBodyPaymentMeanFidelityAccount string = "fidelityAccount"
	// PostMeOrderOrderIDPayWithRegisteredPaymentMeanParamsBodyPaymentMeanOvhAccount captures enum value "ovhAccount"
	PostMeOrderOrderIDPayWithRegisteredPaymentMeanParamsBodyPaymentMeanOvhAccount string = "ovhAccount"
	// PostMeOrderOrderIDPayWithRegisteredPaymentMeanParamsBodyPaymentMeanPaypal captures enum value "paypal"
	PostMeOrderOrderIDPayWithRegisteredPaymentMeanParamsBodyPaymentMeanPaypal string = "paypal"
)

// prop value enum
func (m *PostMeOrderOrderIDPayWithRegisteredPaymentMeanParamsBody) validatePaymentMeanEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postMeOrderOrderIdPayWithRegisteredPaymentMeanParamsBodyTypePaymentMeanPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostMeOrderOrderIDPayWithRegisteredPaymentMeanParamsBody) validatePaymentMean(formats strfmt.Registry) error {

	if err := validate.Required("paymentMean", "body", m.PaymentMean); err != nil {
		return err
	}

	// value enum
	if err := m.validatePaymentMeanEnum("paymentMean", "body", *m.PaymentMean); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostMeOrderOrderIDPayWithRegisteredPaymentMeanParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostMeOrderOrderIDPayWithRegisteredPaymentMeanParamsBody) UnmarshalBinary(b []byte) error {
	var res PostMeOrderOrderIDPayWithRegisteredPaymentMeanParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}