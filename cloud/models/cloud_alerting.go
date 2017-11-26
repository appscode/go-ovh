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

// CloudAlerting Cloud alerting consumption
// swagger:model cloud.Alerting
type CloudAlerting struct {

	// Alerting creation date
	// Required: true
	// Read Only: true
	CreationDate strfmt.DateTime `json:"creationDate"`

	// Delay between alerts in seconds
	Delay int64 `json:"delay,omitempty"`

	// Email to contact
	Email string `json:"email,omitempty"`

	// formatted monthly threshold
	// Required: true
	FormattedMonthlyThreshold *OrderPrice `json:"formattedMonthlyThreshold"`

	// Alerting unique UUID
	// Required: true
	// Read Only: true
	ID string `json:"id"`

	// Monthly threshold for this alerting
	MonthlyThreshold int64 `json:"monthlyThreshold,omitempty"`
}

// Validate validates this cloud alerting
func (m *CloudAlerting) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDelay(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFormattedMonthlyThreshold(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAlerting) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", strfmt.DateTime(m.CreationDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var cloudAlertingTypeDelayPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[10800,172800,21600,259200,3600,43200,604800,86400]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudAlertingTypeDelayPropEnum = append(cloudAlertingTypeDelayPropEnum, v)
	}
}

// prop value enum
func (m *CloudAlerting) validateDelayEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, cloudAlertingTypeDelayPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CloudAlerting) validateDelay(formats strfmt.Registry) error {

	if swag.IsZero(m.Delay) { // not required
		return nil
	}

	// value enum
	if err := m.validateDelayEnum("delay", "body", m.Delay); err != nil {
		return err
	}

	return nil
}

func (m *CloudAlerting) validateFormattedMonthlyThreshold(formats strfmt.Registry) error {

	if err := validate.Required("formattedMonthlyThreshold", "body", m.FormattedMonthlyThreshold); err != nil {
		return err
	}

	if m.FormattedMonthlyThreshold != nil {

		if err := m.FormattedMonthlyThreshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("formattedMonthlyThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *CloudAlerting) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudAlerting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudAlerting) UnmarshalBinary(b []byte) error {
	var res CloudAlerting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}