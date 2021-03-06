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

// CloudQuotaQuotas Quotas
// swagger:model cloud.Quota.Quotas
type CloudQuotaQuotas struct {

	// instance
	// Required: true
	Instance *CloudQuotaInstanceQuotas `json:"instance"`

	// keypair
	// Required: true
	Keypair *CloudQuotaKeypairQuotas `json:"keypair"`

	// Region
	// Required: true
	// Read Only: true
	Region string `json:"region"`

	// volume
	// Required: true
	Volume *CloudQuotaVolumeQuotas `json:"volume"`
}

// Validate validates this cloud quota quotas
func (m *CloudQuotaQuotas) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstance(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKeypair(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVolume(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudQuotaQuotas) validateInstance(formats strfmt.Registry) error {

	if err := validate.Required("instance", "body", m.Instance); err != nil {
		return err
	}

	if m.Instance != nil {

		if err := m.Instance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instance")
			}
			return err
		}
	}

	return nil
}

func (m *CloudQuotaQuotas) validateKeypair(formats strfmt.Registry) error {

	if err := validate.Required("keypair", "body", m.Keypair); err != nil {
		return err
	}

	if m.Keypair != nil {

		if err := m.Keypair.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keypair")
			}
			return err
		}
	}

	return nil
}

func (m *CloudQuotaQuotas) validateRegion(formats strfmt.Registry) error {

	if err := validate.RequiredString("region", "body", string(m.Region)); err != nil {
		return err
	}

	return nil
}

func (m *CloudQuotaQuotas) validateVolume(formats strfmt.Registry) error {

	if err := validate.Required("volume", "body", m.Volume); err != nil {
		return err
	}

	if m.Volume != nil {

		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudQuotaQuotas) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudQuotaQuotas) UnmarshalBinary(b []byte) error {
	var res CloudQuotaQuotas
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
