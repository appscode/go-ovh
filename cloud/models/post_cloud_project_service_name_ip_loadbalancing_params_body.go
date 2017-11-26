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

// PostCloudProjectServiceNameIPLoadbalancingParamsBody post cloud project service name Ip loadbalancing params body
// swagger:model postCloudProjectServiceNameIpLoadbalancingParamsBody
type PostCloudProjectServiceNameIPLoadbalancingParamsBody struct {

	// ip loadbalancing service name
	// Required: true
	IPLoadbalancingServiceName *string `json:"ipLoadbalancingServiceName"`

	// redirection
	// Required: true
	Redirection *string `json:"redirection"`
}

// Validate validates this post cloud project service name Ip loadbalancing params body
func (m *PostCloudProjectServiceNameIPLoadbalancingParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPLoadbalancingServiceName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRedirection(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostCloudProjectServiceNameIPLoadbalancingParamsBody) validateIPLoadbalancingServiceName(formats strfmt.Registry) error {

	if err := validate.Required("ipLoadbalancingServiceName", "body", m.IPLoadbalancingServiceName); err != nil {
		return err
	}

	return nil
}

func (m *PostCloudProjectServiceNameIPLoadbalancingParamsBody) validateRedirection(formats strfmt.Registry) error {

	if err := validate.Required("redirection", "body", m.Redirection); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostCloudProjectServiceNameIPLoadbalancingParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostCloudProjectServiceNameIPLoadbalancingParamsBody) UnmarshalBinary(b []byte) error {
	var res PostCloudProjectServiceNameIPLoadbalancingParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}