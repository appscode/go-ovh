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

// NichandleAccessRestrictionU2FAccount U2F Two-Factor Authentication
// swagger:model nichandle.AccessRestriction.U2FAccount
type NichandleAccessRestrictionU2FAccount struct {

	// Creation date
	// Required: true
	// Read Only: true
	CreationDate strfmt.DateTime `json:"creationDate"`

	// Description of this U2F key
	Description string `json:"description,omitempty"`

	// The Id of the restriction
	// Required: true
	// Read Only: true
	ID int64 `json:"id"`

	// Last used date
	// Read Only: true
	LastUsedDate strfmt.DateTime `json:"lastUsedDate,omitempty"`

	// Status of this account
	// Required: true
	// Read Only: true
	Status string `json:"status"`
}

// Validate validates this nichandle access restriction u2 f account
func (m *NichandleAccessRestrictionU2FAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
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

func (m *NichandleAccessRestrictionU2FAccount) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", strfmt.DateTime(m.CreationDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NichandleAccessRestrictionU2FAccount) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

var nichandleAccessRestrictionU2FAccountTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["disabled","enabled","needCodeValidation"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nichandleAccessRestrictionU2FAccountTypeStatusPropEnum = append(nichandleAccessRestrictionU2FAccountTypeStatusPropEnum, v)
	}
}

const (
	// NichandleAccessRestrictionU2FAccountStatusDisabled captures enum value "disabled"
	NichandleAccessRestrictionU2FAccountStatusDisabled string = "disabled"
	// NichandleAccessRestrictionU2FAccountStatusEnabled captures enum value "enabled"
	NichandleAccessRestrictionU2FAccountStatusEnabled string = "enabled"
	// NichandleAccessRestrictionU2FAccountStatusNeedCodeValidation captures enum value "needCodeValidation"
	NichandleAccessRestrictionU2FAccountStatusNeedCodeValidation string = "needCodeValidation"
)

// prop value enum
func (m *NichandleAccessRestrictionU2FAccount) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, nichandleAccessRestrictionU2FAccountTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NichandleAccessRestrictionU2FAccount) validateStatus(formats strfmt.Registry) error {

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
func (m *NichandleAccessRestrictionU2FAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NichandleAccessRestrictionU2FAccount) UnmarshalBinary(b []byte) error {
	var res NichandleAccessRestrictionU2FAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
