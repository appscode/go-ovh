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

// PostMeAccessRestrictionSmsParamsBody post me access restriction sms params body
// swagger:model postMeAccessRestrictionSmsParamsBody
type PostMeAccessRestrictionSmsParamsBody struct {

	// phone
	// Required: true
	Phone *string `json:"phone"`
}

// Validate validates this post me access restriction sms params body
func (m *PostMeAccessRestrictionSmsParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhone(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostMeAccessRestrictionSmsParamsBody) validatePhone(formats strfmt.Registry) error {

	if err := validate.Required("phone", "body", m.Phone); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostMeAccessRestrictionSmsParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostMeAccessRestrictionSmsParamsBody) UnmarshalBinary(b []byte) error {
	var res PostMeAccessRestrictionSmsParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
