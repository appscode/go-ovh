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

// ServicesNonExpiringService Details about a non-expiring Service
// swagger:model services.NonExpiringService
type ServicesNonExpiringService struct {

	// contact admin
	// Required: true
	// Read Only: true
	ContactAdmin string `json:"contactAdmin"`

	// contact billing
	// Required: true
	// Read Only: true
	ContactBilling string `json:"contactBilling"`

	// contact tech
	// Required: true
	// Read Only: true
	ContactTech string `json:"contactTech"`

	// creation
	// Required: true
	// Read Only: true
	Creation strfmt.Date `json:"creation"`

	// domain
	// Required: true
	// Read Only: true
	Domain string `json:"domain"`

	// service Id
	// Required: true
	// Read Only: true
	ServiceID int64 `json:"serviceId"`

	// status
	// Required: true
	// Read Only: true
	Status string `json:"status"`
}

// Validate validates this services non expiring service
func (m *ServicesNonExpiringService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContactAdmin(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateContactBilling(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateContactTech(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServiceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServicesNonExpiringService) validateContactAdmin(formats strfmt.Registry) error {

	if err := validate.RequiredString("contactAdmin", "body", string(m.ContactAdmin)); err != nil {
		return err
	}

	return nil
}

func (m *ServicesNonExpiringService) validateContactBilling(formats strfmt.Registry) error {

	if err := validate.RequiredString("contactBilling", "body", string(m.ContactBilling)); err != nil {
		return err
	}

	return nil
}

func (m *ServicesNonExpiringService) validateContactTech(formats strfmt.Registry) error {

	if err := validate.RequiredString("contactTech", "body", string(m.ContactTech)); err != nil {
		return err
	}

	return nil
}

func (m *ServicesNonExpiringService) validateCreation(formats strfmt.Registry) error {

	if err := validate.Required("creation", "body", strfmt.Date(m.Creation)); err != nil {
		return err
	}

	if err := validate.FormatOf("creation", "body", "date", m.Creation.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServicesNonExpiringService) validateDomain(formats strfmt.Registry) error {

	if err := validate.RequiredString("domain", "body", string(m.Domain)); err != nil {
		return err
	}

	return nil
}

func (m *ServicesNonExpiringService) validateServiceID(formats strfmt.Registry) error {

	if err := validate.Required("serviceId", "body", int64(m.ServiceID)); err != nil {
		return err
	}

	return nil
}

var servicesNonExpiringServiceTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["expired","inCreation","ok","unPaid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		servicesNonExpiringServiceTypeStatusPropEnum = append(servicesNonExpiringServiceTypeStatusPropEnum, v)
	}
}

const (
	// ServicesNonExpiringServiceStatusExpired captures enum value "expired"
	ServicesNonExpiringServiceStatusExpired string = "expired"
	// ServicesNonExpiringServiceStatusInCreation captures enum value "inCreation"
	ServicesNonExpiringServiceStatusInCreation string = "inCreation"
	// ServicesNonExpiringServiceStatusOk captures enum value "ok"
	ServicesNonExpiringServiceStatusOk string = "ok"
	// ServicesNonExpiringServiceStatusUnPaid captures enum value "unPaid"
	ServicesNonExpiringServiceStatusUnPaid string = "unPaid"
)

// prop value enum
func (m *ServicesNonExpiringService) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, servicesNonExpiringServiceTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServicesNonExpiringService) validateStatus(formats strfmt.Registry) error {

	if err := validate.RequiredString("status", "body", string(m.Status)); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServicesNonExpiringService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServicesNonExpiringService) UnmarshalBinary(b []byte) error {
	var res ServicesNonExpiringService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
