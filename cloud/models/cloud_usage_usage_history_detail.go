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

// CloudUsageUsageHistoryDetail UsageHistoryDetail
// swagger:model cloud.Usage.UsageHistoryDetail
type CloudUsageUsageHistoryDetail struct {

	// hourly usage
	HourlyUsage *CloudBillingViewHourlyResources `json:"hourlyUsage,omitempty"`

	// Usage id
	// Required: true
	// Read Only: true
	ID string `json:"id"`

	// Entry last update
	// Required: true
	// Read Only: true
	LastUpdate strfmt.DateTime `json:"lastUpdate"`

	// monthly usage
	MonthlyUsage *CloudBillingViewMonthlyResources `json:"monthlyUsage,omitempty"`

	// period
	// Required: true
	Period *CloudUsagePeriod `json:"period"`
}

// Validate validates this cloud usage usage history detail
func (m *CloudUsageUsageHistoryDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHourlyUsage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastUpdate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMonthlyUsage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePeriod(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudUsageUsageHistoryDetail) validateHourlyUsage(formats strfmt.Registry) error {

	if swag.IsZero(m.HourlyUsage) { // not required
		return nil
	}

	if m.HourlyUsage != nil {

		if err := m.HourlyUsage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hourlyUsage")
			}
			return err
		}
	}

	return nil
}

func (m *CloudUsageUsageHistoryDetail) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *CloudUsageUsageHistoryDetail) validateLastUpdate(formats strfmt.Registry) error {

	if err := validate.Required("lastUpdate", "body", strfmt.DateTime(m.LastUpdate)); err != nil {
		return err
	}

	if err := validate.FormatOf("lastUpdate", "body", "date-time", m.LastUpdate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CloudUsageUsageHistoryDetail) validateMonthlyUsage(formats strfmt.Registry) error {

	if swag.IsZero(m.MonthlyUsage) { // not required
		return nil
	}

	if m.MonthlyUsage != nil {

		if err := m.MonthlyUsage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monthlyUsage")
			}
			return err
		}
	}

	return nil
}

func (m *CloudUsageUsageHistoryDetail) validatePeriod(formats strfmt.Registry) error {

	if err := validate.Required("period", "body", m.Period); err != nil {
		return err
	}

	if m.Period != nil {

		if err := m.Period.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("period")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudUsageUsageHistoryDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudUsageUsageHistoryDetail) UnmarshalBinary(b []byte) error {
	var res CloudUsageUsageHistoryDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
