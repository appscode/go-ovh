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

// GetMeAccessRestrictionSmsIDDefaultBody get me access restriction sms Id default body
// swagger:model getMeAccessRestrictionSmsIdDefaultBody
type GetMeAccessRestrictionSmsIDDefaultBody struct {

	// error code
	ErrorCode int32 `json:"errorCode,omitempty"`

	// http code
	HTTPCode int32 `json:"httpCode,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get me access restriction sms Id default body
func (m *GetMeAccessRestrictionSmsIDDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetMeAccessRestrictionSmsIDDefaultBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetMeAccessRestrictionSmsIDDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetMeAccessRestrictionSmsIDDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
