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

// PostMeAccessRestrictionBackupCodeEnableParamsBody post me access restriction backup code enable params body
// swagger:model postMeAccessRestrictionBackupCodeEnableParamsBody
type PostMeAccessRestrictionBackupCodeEnableParamsBody struct {

	// code
	// Required: true
	Code *strfmt.Password `json:"code"`
}

// Validate validates this post me access restriction backup code enable params body
func (m *PostMeAccessRestrictionBackupCodeEnableParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostMeAccessRestrictionBackupCodeEnableParamsBody) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	if err := validate.FormatOf("code", "body", "password", m.Code.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostMeAccessRestrictionBackupCodeEnableParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostMeAccessRestrictionBackupCodeEnableParamsBody) UnmarshalBinary(b []byte) error {
	var res PostMeAccessRestrictionBackupCodeEnableParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
