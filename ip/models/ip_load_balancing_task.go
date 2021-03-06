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

// IPLoadBalancingTask List of tasks associated with your IP load balancing
// swagger:model ip.LoadBalancingTask
type IPLoadBalancingTask struct {

	// The action made
	// Required: true
	// Read Only: true
	Action string `json:"action"`

	// Creation date of your task
	// Required: true
	// Read Only: true
	CreationDate strfmt.DateTime `json:"creationDate"`

	// Identifier of your task
	// Required: true
	// Read Only: true
	ID int64 `json:"id"`

	// Current status of your task
	// Required: true
	// Read Only: true
	Status string `json:"status"`
}

// Validate validates this ip load balancing task
func (m *IPLoadBalancingTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		// prop
		res = append(res, err)
	}

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

var ipLoadBalancingTaskTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["activateSsl","addBackend","addIpToBackend","announceIpLoadBalancing","backupStateSet","backupStateUnset","changeProbe","delBackend","desactivateSsl","removeIpFromBackend","setPortRedirection","setStickiness","setWeight","unannounceIpLoadBalancing","unsetPortRedirection"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipLoadBalancingTaskTypeActionPropEnum = append(ipLoadBalancingTaskTypeActionPropEnum, v)
	}
}

const (
	// IPLoadBalancingTaskActionActivateSsl captures enum value "activateSsl"
	IPLoadBalancingTaskActionActivateSsl string = "activateSsl"
	// IPLoadBalancingTaskActionAddBackend captures enum value "addBackend"
	IPLoadBalancingTaskActionAddBackend string = "addBackend"
	// IPLoadBalancingTaskActionAddIPToBackend captures enum value "addIpToBackend"
	IPLoadBalancingTaskActionAddIPToBackend string = "addIpToBackend"
	// IPLoadBalancingTaskActionAnnounceIPLoadBalancing captures enum value "announceIpLoadBalancing"
	IPLoadBalancingTaskActionAnnounceIPLoadBalancing string = "announceIpLoadBalancing"
	// IPLoadBalancingTaskActionBackupStateSet captures enum value "backupStateSet"
	IPLoadBalancingTaskActionBackupStateSet string = "backupStateSet"
	// IPLoadBalancingTaskActionBackupStateUnset captures enum value "backupStateUnset"
	IPLoadBalancingTaskActionBackupStateUnset string = "backupStateUnset"
	// IPLoadBalancingTaskActionChangeProbe captures enum value "changeProbe"
	IPLoadBalancingTaskActionChangeProbe string = "changeProbe"
	// IPLoadBalancingTaskActionDelBackend captures enum value "delBackend"
	IPLoadBalancingTaskActionDelBackend string = "delBackend"
	// IPLoadBalancingTaskActionDesactivateSsl captures enum value "desactivateSsl"
	IPLoadBalancingTaskActionDesactivateSsl string = "desactivateSsl"
	// IPLoadBalancingTaskActionRemoveIPFromBackend captures enum value "removeIpFromBackend"
	IPLoadBalancingTaskActionRemoveIPFromBackend string = "removeIpFromBackend"
	// IPLoadBalancingTaskActionSetPortRedirection captures enum value "setPortRedirection"
	IPLoadBalancingTaskActionSetPortRedirection string = "setPortRedirection"
	// IPLoadBalancingTaskActionSetStickiness captures enum value "setStickiness"
	IPLoadBalancingTaskActionSetStickiness string = "setStickiness"
	// IPLoadBalancingTaskActionSetWeight captures enum value "setWeight"
	IPLoadBalancingTaskActionSetWeight string = "setWeight"
	// IPLoadBalancingTaskActionUnannounceIPLoadBalancing captures enum value "unannounceIpLoadBalancing"
	IPLoadBalancingTaskActionUnannounceIPLoadBalancing string = "unannounceIpLoadBalancing"
	// IPLoadBalancingTaskActionUnsetPortRedirection captures enum value "unsetPortRedirection"
	IPLoadBalancingTaskActionUnsetPortRedirection string = "unsetPortRedirection"
)

// prop value enum
func (m *IPLoadBalancingTask) validateActionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ipLoadBalancingTaskTypeActionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IPLoadBalancingTask) validateAction(formats strfmt.Registry) error {

	if err := validate.RequiredString("action", "body", string(m.Action)); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *IPLoadBalancingTask) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", strfmt.DateTime(m.CreationDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IPLoadBalancingTask) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *IPLoadBalancingTask) validateStatus(formats strfmt.Registry) error {

	if err := validate.RequiredString("status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPLoadBalancingTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPLoadBalancingTask) UnmarshalBinary(b []byte) error {
	var res IPLoadBalancingTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
