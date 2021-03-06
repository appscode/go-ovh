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

// CloudProjectInstanceUsageDetail Instance usage
// swagger:model cloud.Project.InstanceUsageDetail
type CloudProjectInstanceUsageDetail struct {

	// hourly
	Hourly *OrderPrice `json:"hourly,omitempty"`

	// Instance id
	InstanceID string `json:"instanceId,omitempty"`

	// monthly
	Monthly *CloudProjectInstanceMonthlyBilling `json:"monthly,omitempty"`

	// Is monthly billing enabled
	MonthlyBilling bool `json:"monthlyBilling,omitempty"`

	// Reference
	Reference string `json:"reference,omitempty"`
}

// Validate validates this cloud project instance usage detail
func (m *CloudProjectInstanceUsageDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHourly(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMonthly(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudProjectInstanceUsageDetail) validateHourly(formats strfmt.Registry) error {

	if swag.IsZero(m.Hourly) { // not required
		return nil
	}

	if m.Hourly != nil {

		if err := m.Hourly.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hourly")
			}
			return err
		}
	}

	return nil
}

func (m *CloudProjectInstanceUsageDetail) validateMonthly(formats strfmt.Registry) error {

	if swag.IsZero(m.Monthly) { // not required
		return nil
	}

	if m.Monthly != nil {

		if err := m.Monthly.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monthly")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudProjectInstanceUsageDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudProjectInstanceUsageDetail) UnmarshalBinary(b []byte) error {
	var res CloudProjectInstanceUsageDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
