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

// CloudInstanceInstanceDetail InstanceDetail
// swagger:model cloud.Instance.InstanceDetail
type CloudInstanceInstanceDetail struct {

	// Instance creation date
	// Required: true
	// Read Only: true
	Created strfmt.DateTime `json:"created"`

	// flavor
	// Required: true
	Flavor *CloudFlavorFlavor `json:"flavor"`

	// Instance id
	// Required: true
	// Read Only: true
	ID string `json:"id"`

	// image
	// Required: true
	Image *CloudImageImage `json:"image"`

	// ip addresses
	// Required: true
	IPAddresses CloudInstanceInstanceDetailIPAddresses `json:"ipAddresses"`

	// monthly billing
	MonthlyBilling *CloudInstanceMonthlyBilling `json:"monthlyBilling,omitempty"`

	// Instance name
	// Required: true
	// Read Only: true
	Name string `json:"name"`

	// Instance id
	// Required: true
	// Read Only: true
	Region string `json:"region"`

	// ssh key
	// Required: true
	SSHKey *CloudSshkeySSHKeyDetail `json:"sshKey"`

	// Instance status
	// Required: true
	// Read Only: true
	Status string `json:"status"`
}

// Validate validates this cloud instance instance detail
func (m *CloudInstanceInstanceDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFlavor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPAddresses(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMonthlyBilling(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSSHKey(formats); err != nil {
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

func (m *CloudInstanceInstanceDetail) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CloudInstanceInstanceDetail) validateFlavor(formats strfmt.Registry) error {

	if err := validate.Required("flavor", "body", m.Flavor); err != nil {
		return err
	}

	if m.Flavor != nil {

		if err := m.Flavor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flavor")
			}
			return err
		}
	}

	return nil
}

func (m *CloudInstanceInstanceDetail) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *CloudInstanceInstanceDetail) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	if m.Image != nil {

		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *CloudInstanceInstanceDetail) validateIPAddresses(formats strfmt.Registry) error {

	if err := validate.Required("ipAddresses", "body", m.IPAddresses); err != nil {
		return err
	}

	if err := m.IPAddresses.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ipAddresses")
		}
		return err
	}

	return nil
}

func (m *CloudInstanceInstanceDetail) validateMonthlyBilling(formats strfmt.Registry) error {

	if swag.IsZero(m.MonthlyBilling) { // not required
		return nil
	}

	if m.MonthlyBilling != nil {

		if err := m.MonthlyBilling.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monthlyBilling")
			}
			return err
		}
	}

	return nil
}

func (m *CloudInstanceInstanceDetail) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *CloudInstanceInstanceDetail) validateRegion(formats strfmt.Registry) error {

	if err := validate.RequiredString("region", "body", string(m.Region)); err != nil {
		return err
	}

	return nil
}

func (m *CloudInstanceInstanceDetail) validateSSHKey(formats strfmt.Registry) error {

	if err := validate.Required("sshKey", "body", m.SSHKey); err != nil {
		return err
	}

	if m.SSHKey != nil {

		if err := m.SSHKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sshKey")
			}
			return err
		}
	}

	return nil
}

var cloudInstanceInstanceDetailTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","BUILDING","DELETED","ERROR","HARD_REBOOT","PASSWORD","PAUSED","REBOOT","REBUILD","RESCUED","RESIZED","REVERT_RESIZE","SOFT_DELETED","STOPPED","SUSPENDED","UNKNOWN","VERIFY_RESIZE","MIGRATING","RESIZE","BUILD","SHUTOFF","RESCUE","SHELVED","SHELVED_OFFLOADED","RESCUING","UNRESCUING","SNAPSHOTTING","RESUMING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudInstanceInstanceDetailTypeStatusPropEnum = append(cloudInstanceInstanceDetailTypeStatusPropEnum, v)
	}
}

const (
	// CloudInstanceInstanceDetailStatusACTIVE captures enum value "ACTIVE"
	CloudInstanceInstanceDetailStatusACTIVE string = "ACTIVE"
	// CloudInstanceInstanceDetailStatusBUILDING captures enum value "BUILDING"
	CloudInstanceInstanceDetailStatusBUILDING string = "BUILDING"
	// CloudInstanceInstanceDetailStatusDELETED captures enum value "DELETED"
	CloudInstanceInstanceDetailStatusDELETED string = "DELETED"
	// CloudInstanceInstanceDetailStatusERROR captures enum value "ERROR"
	CloudInstanceInstanceDetailStatusERROR string = "ERROR"
	// CloudInstanceInstanceDetailStatusHARDREBOOT captures enum value "HARD_REBOOT"
	CloudInstanceInstanceDetailStatusHARDREBOOT string = "HARD_REBOOT"
	// CloudInstanceInstanceDetailStatusPASSWORD captures enum value "PASSWORD"
	CloudInstanceInstanceDetailStatusPASSWORD string = "PASSWORD"
	// CloudInstanceInstanceDetailStatusPAUSED captures enum value "PAUSED"
	CloudInstanceInstanceDetailStatusPAUSED string = "PAUSED"
	// CloudInstanceInstanceDetailStatusREBOOT captures enum value "REBOOT"
	CloudInstanceInstanceDetailStatusREBOOT string = "REBOOT"
	// CloudInstanceInstanceDetailStatusREBUILD captures enum value "REBUILD"
	CloudInstanceInstanceDetailStatusREBUILD string = "REBUILD"
	// CloudInstanceInstanceDetailStatusRESCUED captures enum value "RESCUED"
	CloudInstanceInstanceDetailStatusRESCUED string = "RESCUED"
	// CloudInstanceInstanceDetailStatusRESIZED captures enum value "RESIZED"
	CloudInstanceInstanceDetailStatusRESIZED string = "RESIZED"
	// CloudInstanceInstanceDetailStatusREVERTRESIZE captures enum value "REVERT_RESIZE"
	CloudInstanceInstanceDetailStatusREVERTRESIZE string = "REVERT_RESIZE"
	// CloudInstanceInstanceDetailStatusSOFTDELETED captures enum value "SOFT_DELETED"
	CloudInstanceInstanceDetailStatusSOFTDELETED string = "SOFT_DELETED"
	// CloudInstanceInstanceDetailStatusSTOPPED captures enum value "STOPPED"
	CloudInstanceInstanceDetailStatusSTOPPED string = "STOPPED"
	// CloudInstanceInstanceDetailStatusSUSPENDED captures enum value "SUSPENDED"
	CloudInstanceInstanceDetailStatusSUSPENDED string = "SUSPENDED"
	// CloudInstanceInstanceDetailStatusUNKNOWN captures enum value "UNKNOWN"
	CloudInstanceInstanceDetailStatusUNKNOWN string = "UNKNOWN"
	// CloudInstanceInstanceDetailStatusVERIFYRESIZE captures enum value "VERIFY_RESIZE"
	CloudInstanceInstanceDetailStatusVERIFYRESIZE string = "VERIFY_RESIZE"
	// CloudInstanceInstanceDetailStatusMIGRATING captures enum value "MIGRATING"
	CloudInstanceInstanceDetailStatusMIGRATING string = "MIGRATING"
	// CloudInstanceInstanceDetailStatusRESIZE captures enum value "RESIZE"
	CloudInstanceInstanceDetailStatusRESIZE string = "RESIZE"
	// CloudInstanceInstanceDetailStatusBUILD captures enum value "BUILD"
	CloudInstanceInstanceDetailStatusBUILD string = "BUILD"
	// CloudInstanceInstanceDetailStatusSHUTOFF captures enum value "SHUTOFF"
	CloudInstanceInstanceDetailStatusSHUTOFF string = "SHUTOFF"
	// CloudInstanceInstanceDetailStatusRESCUE captures enum value "RESCUE"
	CloudInstanceInstanceDetailStatusRESCUE string = "RESCUE"
	// CloudInstanceInstanceDetailStatusSHELVED captures enum value "SHELVED"
	CloudInstanceInstanceDetailStatusSHELVED string = "SHELVED"
	// CloudInstanceInstanceDetailStatusSHELVEDOFFLOADED captures enum value "SHELVED_OFFLOADED"
	CloudInstanceInstanceDetailStatusSHELVEDOFFLOADED string = "SHELVED_OFFLOADED"
	// CloudInstanceInstanceDetailStatusRESCUING captures enum value "RESCUING"
	CloudInstanceInstanceDetailStatusRESCUING string = "RESCUING"
	// CloudInstanceInstanceDetailStatusUNRESCUING captures enum value "UNRESCUING"
	CloudInstanceInstanceDetailStatusUNRESCUING string = "UNRESCUING"
	// CloudInstanceInstanceDetailStatusSNAPSHOTTING captures enum value "SNAPSHOTTING"
	CloudInstanceInstanceDetailStatusSNAPSHOTTING string = "SNAPSHOTTING"
	// CloudInstanceInstanceDetailStatusRESUMING captures enum value "RESUMING"
	CloudInstanceInstanceDetailStatusRESUMING string = "RESUMING"
)

// prop value enum
func (m *CloudInstanceInstanceDetail) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cloudInstanceInstanceDetailTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CloudInstanceInstanceDetail) validateStatus(formats strfmt.Registry) error {

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
func (m *CloudInstanceInstanceDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudInstanceInstanceDetail) UnmarshalBinary(b []byte) error {
	var res CloudInstanceInstanceDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
