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

// PostCloudProjectServiceNameInstanceBulkParamsBody post cloud project service name instance bulk params body
// swagger:model postCloudProjectServiceNameInstanceBulkParamsBody
type PostCloudProjectServiceNameInstanceBulkParamsBody struct {

	// flavor Id
	// Required: true
	FlavorID *string `json:"flavorId"`

	// group Id
	GroupID string `json:"groupId,omitempty"`

	// image Id
	ImageID string `json:"imageId,omitempty"`

	// monthly billing
	MonthlyBilling *bool `json:"monthlyBilling,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// networks
	Networks PostCloudProjectServiceNameInstanceBulkParamsBodyNetworks `json:"networks"`

	// number
	// Required: true
	Number *int64 `json:"number"`

	// region
	// Required: true
	Region *string `json:"region"`

	// ssh key Id
	SSHKeyID string `json:"sshKeyId,omitempty"`

	// user data
	UserData string `json:"userData,omitempty"`

	// volume Id
	VolumeID string `json:"volumeId,omitempty"`
}

// Validate validates this post cloud project service name instance bulk params body
func (m *PostCloudProjectServiceNameInstanceBulkParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlavorID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostCloudProjectServiceNameInstanceBulkParamsBody) validateFlavorID(formats strfmt.Registry) error {

	if err := validate.Required("flavorId", "body", m.FlavorID); err != nil {
		return err
	}

	return nil
}

func (m *PostCloudProjectServiceNameInstanceBulkParamsBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PostCloudProjectServiceNameInstanceBulkParamsBody) validateNumber(formats strfmt.Registry) error {

	if err := validate.Required("number", "body", m.Number); err != nil {
		return err
	}

	return nil
}

func (m *PostCloudProjectServiceNameInstanceBulkParamsBody) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostCloudProjectServiceNameInstanceBulkParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostCloudProjectServiceNameInstanceBulkParamsBody) UnmarshalBinary(b []byte) error {
	var res PostCloudProjectServiceNameInstanceBulkParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
