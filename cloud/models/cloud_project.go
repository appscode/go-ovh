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

// CloudProject Project
// swagger:model cloud.Project
type CloudProject struct {

	// Project access
	// Required: true
	// Read Only: true
	Access string `json:"access"`

	// Project creation date
	// Required: true
	// Read Only: true
	CreationDate strfmt.DateTime `json:"creationDate"`

	// Description of your project
	Description string `json:"description,omitempty"`

	// Expiration date of your project. After this date, your project will be deleted
	// Read Only: true
	Expiration strfmt.DateTime `json:"expiration,omitempty"`

	// Project order id
	// Read Only: true
	OrderID int64 `json:"orderId,omitempty"`

	// Project id
	// Required: true
	// Read Only: true
	ProjectID string `json:"project_id"`

	// Current status
	// Required: true
	// Read Only: true
	Status string `json:"status"`

	// Project unleashed
	// Required: true
	// Read Only: true
	Unleash bool `json:"unleash"`
}

// Validate validates this cloud project
func (m *CloudProject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUnleash(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cloudProjectTypeAccessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["full","restricted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudProjectTypeAccessPropEnum = append(cloudProjectTypeAccessPropEnum, v)
	}
}

const (
	// CloudProjectAccessFull captures enum value "full"
	CloudProjectAccessFull string = "full"
	// CloudProjectAccessRestricted captures enum value "restricted"
	CloudProjectAccessRestricted string = "restricted"
)

// prop value enum
func (m *CloudProject) validateAccessEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cloudProjectTypeAccessPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CloudProject) validateAccess(formats strfmt.Registry) error {

	if err := validate.RequiredString("access", "body", string(m.Access)); err != nil {
		return err
	}

	// value enum
	if err := m.validateAccessEnum("access", "body", m.Access); err != nil {
		return err
	}

	return nil
}

func (m *CloudProject) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", strfmt.DateTime(m.CreationDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CloudProject) validateProjectID(formats strfmt.Registry) error {

	if err := validate.RequiredString("project_id", "body", string(m.ProjectID)); err != nil {
		return err
	}

	return nil
}

var cloudProjectTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["creating","deleted","deleting","ok","suspended"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudProjectTypeStatusPropEnum = append(cloudProjectTypeStatusPropEnum, v)
	}
}

const (
	// CloudProjectStatusCreating captures enum value "creating"
	CloudProjectStatusCreating string = "creating"
	// CloudProjectStatusDeleted captures enum value "deleted"
	CloudProjectStatusDeleted string = "deleted"
	// CloudProjectStatusDeleting captures enum value "deleting"
	CloudProjectStatusDeleting string = "deleting"
	// CloudProjectStatusOk captures enum value "ok"
	CloudProjectStatusOk string = "ok"
	// CloudProjectStatusSuspended captures enum value "suspended"
	CloudProjectStatusSuspended string = "suspended"
)

// prop value enum
func (m *CloudProject) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cloudProjectTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CloudProject) validateStatus(formats strfmt.Registry) error {

	if err := validate.RequiredString("status", "body", string(m.Status)); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *CloudProject) validateUnleash(formats strfmt.Registry) error {

	if err := validate.Required("unleash", "body", bool(m.Unleash)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudProject) UnmarshalBinary(b []byte) error {
	var res CloudProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
