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

// PostMePasswordRecoverParamsBody post me password recover params body
// swagger:model postMePasswordRecoverParamsBody
type PostMePasswordRecoverParamsBody struct {

	// ovh company
	// Required: true
	OvhCompany *string `json:"ovhCompany"`

	// ovh Id
	// Required: true
	OvhID *string `json:"ovhId"`
}

// Validate validates this post me password recover params body
func (m *PostMePasswordRecoverParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOvhCompany(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOvhID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postMePasswordRecoverParamsBodyTypeOvhCompanyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kimsufi","ovh","soyoustart"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postMePasswordRecoverParamsBodyTypeOvhCompanyPropEnum = append(postMePasswordRecoverParamsBodyTypeOvhCompanyPropEnum, v)
	}
}

const (
	// PostMePasswordRecoverParamsBodyOvhCompanyKimsufi captures enum value "kimsufi"
	PostMePasswordRecoverParamsBodyOvhCompanyKimsufi string = "kimsufi"
	// PostMePasswordRecoverParamsBodyOvhCompanyOvh captures enum value "ovh"
	PostMePasswordRecoverParamsBodyOvhCompanyOvh string = "ovh"
	// PostMePasswordRecoverParamsBodyOvhCompanySoyoustart captures enum value "soyoustart"
	PostMePasswordRecoverParamsBodyOvhCompanySoyoustart string = "soyoustart"
)

// prop value enum
func (m *PostMePasswordRecoverParamsBody) validateOvhCompanyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postMePasswordRecoverParamsBodyTypeOvhCompanyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostMePasswordRecoverParamsBody) validateOvhCompany(formats strfmt.Registry) error {

	if err := validate.Required("ovhCompany", "body", m.OvhCompany); err != nil {
		return err
	}

	// value enum
	if err := m.validateOvhCompanyEnum("ovhCompany", "body", *m.OvhCompany); err != nil {
		return err
	}

	return nil
}

func (m *PostMePasswordRecoverParamsBody) validateOvhID(formats strfmt.Registry) error {

	if err := validate.Required("ovhId", "body", m.OvhID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostMePasswordRecoverParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostMePasswordRecoverParamsBody) UnmarshalBinary(b []byte) error {
	var res PostMePasswordRecoverParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
