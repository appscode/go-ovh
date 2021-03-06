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

// NichandleDomainTask Domain tasks
// swagger:model nichandle.DomainTask
type NichandleDomainTask struct {

	// Can accelerate the task
	// Required: true
	// Read Only: true
	CanAccelerate bool `json:"canAccelerate"`

	// Can cancel the task
	// Required: true
	// Read Only: true
	CanCancel bool `json:"canCancel"`

	// Can relaunch the task
	// Required: true
	// Read Only: true
	CanRelaunch bool `json:"canRelaunch"`

	// Comment about the task
	// Read Only: true
	Comment string `json:"comment,omitempty"`

	// Creation date of the task
	// Required: true
	// Read Only: true
	CreationDate strfmt.DateTime `json:"creationDate"`

	// Domain of the task
	// Required: true
	// Read Only: true
	Domain string `json:"domain"`

	// Done date of the task
	// Read Only: true
	DoneDate strfmt.DateTime `json:"doneDate,omitempty"`

	// Function of the task
	// Required: true
	// Read Only: true
	Function string `json:"function"`

	// Id of the task
	// Required: true
	// Read Only: true
	ID int64 `json:"id"`

	// Last update date of the task
	// Required: true
	// Read Only: true
	LastUpdate strfmt.DateTime `json:"lastUpdate"`

	// Status of the task
	// Required: true
	// Read Only: true
	Status string `json:"status"`

	// Todo date of the task
	// Required: true
	// Read Only: true
	TodoDate strfmt.DateTime `json:"todoDate"`
}

// Validate validates this nichandle domain task
func (m *NichandleDomainTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanAccelerate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCanCancel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCanRelaunch(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFunction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastUpdate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTodoDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NichandleDomainTask) validateCanAccelerate(formats strfmt.Registry) error {

	if err := validate.Required("canAccelerate", "body", bool(m.CanAccelerate)); err != nil {
		return err
	}

	return nil
}

func (m *NichandleDomainTask) validateCanCancel(formats strfmt.Registry) error {

	if err := validate.Required("canCancel", "body", bool(m.CanCancel)); err != nil {
		return err
	}

	return nil
}

func (m *NichandleDomainTask) validateCanRelaunch(formats strfmt.Registry) error {

	if err := validate.Required("canRelaunch", "body", bool(m.CanRelaunch)); err != nil {
		return err
	}

	return nil
}

func (m *NichandleDomainTask) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", strfmt.DateTime(m.CreationDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NichandleDomainTask) validateDomain(formats strfmt.Registry) error {

	if err := validate.RequiredString("domain", "body", string(m.Domain)); err != nil {
		return err
	}

	return nil
}

var nichandleDomainTaskTypeFunctionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ContactControl","DnsAnycastActivate","DnsAnycastDeactivate","DnssecDisable","DnssecEnable","DnssecResigning","DnssecRollKsk","DnssecRollZsk","DomainContactControl","DomainContactUpdate","DomainControl","DomainCreate","DomainDelete","DomainDnsUpdate","DomainDsUpdate","DomainHold","DomainHostCreate","DomainHostDelete","DomainHostUpdate","DomainIncomingTransfer","DomainLock","DomainOutgoingTransfer","DomainRenew","DomainRestore","DomainTrade","ZoneImport"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nichandleDomainTaskTypeFunctionPropEnum = append(nichandleDomainTaskTypeFunctionPropEnum, v)
	}
}

const (
	// NichandleDomainTaskFunctionContactControl captures enum value "ContactControl"
	NichandleDomainTaskFunctionContactControl string = "ContactControl"
	// NichandleDomainTaskFunctionDNSAnycastActivate captures enum value "DnsAnycastActivate"
	NichandleDomainTaskFunctionDNSAnycastActivate string = "DnsAnycastActivate"
	// NichandleDomainTaskFunctionDNSAnycastDeactivate captures enum value "DnsAnycastDeactivate"
	NichandleDomainTaskFunctionDNSAnycastDeactivate string = "DnsAnycastDeactivate"
	// NichandleDomainTaskFunctionDnssecDisable captures enum value "DnssecDisable"
	NichandleDomainTaskFunctionDnssecDisable string = "DnssecDisable"
	// NichandleDomainTaskFunctionDnssecEnable captures enum value "DnssecEnable"
	NichandleDomainTaskFunctionDnssecEnable string = "DnssecEnable"
	// NichandleDomainTaskFunctionDnssecResigning captures enum value "DnssecResigning"
	NichandleDomainTaskFunctionDnssecResigning string = "DnssecResigning"
	// NichandleDomainTaskFunctionDnssecRollKsk captures enum value "DnssecRollKsk"
	NichandleDomainTaskFunctionDnssecRollKsk string = "DnssecRollKsk"
	// NichandleDomainTaskFunctionDnssecRollZsk captures enum value "DnssecRollZsk"
	NichandleDomainTaskFunctionDnssecRollZsk string = "DnssecRollZsk"
	// NichandleDomainTaskFunctionDomainContactControl captures enum value "DomainContactControl"
	NichandleDomainTaskFunctionDomainContactControl string = "DomainContactControl"
	// NichandleDomainTaskFunctionDomainContactUpdate captures enum value "DomainContactUpdate"
	NichandleDomainTaskFunctionDomainContactUpdate string = "DomainContactUpdate"
	// NichandleDomainTaskFunctionDomainControl captures enum value "DomainControl"
	NichandleDomainTaskFunctionDomainControl string = "DomainControl"
	// NichandleDomainTaskFunctionDomainCreate captures enum value "DomainCreate"
	NichandleDomainTaskFunctionDomainCreate string = "DomainCreate"
	// NichandleDomainTaskFunctionDomainDelete captures enum value "DomainDelete"
	NichandleDomainTaskFunctionDomainDelete string = "DomainDelete"
	// NichandleDomainTaskFunctionDomainDNSUpdate captures enum value "DomainDnsUpdate"
	NichandleDomainTaskFunctionDomainDNSUpdate string = "DomainDnsUpdate"
	// NichandleDomainTaskFunctionDomainDsUpdate captures enum value "DomainDsUpdate"
	NichandleDomainTaskFunctionDomainDsUpdate string = "DomainDsUpdate"
	// NichandleDomainTaskFunctionDomainHold captures enum value "DomainHold"
	NichandleDomainTaskFunctionDomainHold string = "DomainHold"
	// NichandleDomainTaskFunctionDomainHostCreate captures enum value "DomainHostCreate"
	NichandleDomainTaskFunctionDomainHostCreate string = "DomainHostCreate"
	// NichandleDomainTaskFunctionDomainHostDelete captures enum value "DomainHostDelete"
	NichandleDomainTaskFunctionDomainHostDelete string = "DomainHostDelete"
	// NichandleDomainTaskFunctionDomainHostUpdate captures enum value "DomainHostUpdate"
	NichandleDomainTaskFunctionDomainHostUpdate string = "DomainHostUpdate"
	// NichandleDomainTaskFunctionDomainIncomingTransfer captures enum value "DomainIncomingTransfer"
	NichandleDomainTaskFunctionDomainIncomingTransfer string = "DomainIncomingTransfer"
	// NichandleDomainTaskFunctionDomainLock captures enum value "DomainLock"
	NichandleDomainTaskFunctionDomainLock string = "DomainLock"
	// NichandleDomainTaskFunctionDomainOutgoingTransfer captures enum value "DomainOutgoingTransfer"
	NichandleDomainTaskFunctionDomainOutgoingTransfer string = "DomainOutgoingTransfer"
	// NichandleDomainTaskFunctionDomainRenew captures enum value "DomainRenew"
	NichandleDomainTaskFunctionDomainRenew string = "DomainRenew"
	// NichandleDomainTaskFunctionDomainRestore captures enum value "DomainRestore"
	NichandleDomainTaskFunctionDomainRestore string = "DomainRestore"
	// NichandleDomainTaskFunctionDomainTrade captures enum value "DomainTrade"
	NichandleDomainTaskFunctionDomainTrade string = "DomainTrade"
	// NichandleDomainTaskFunctionZoneImport captures enum value "ZoneImport"
	NichandleDomainTaskFunctionZoneImport string = "ZoneImport"
)

// prop value enum
func (m *NichandleDomainTask) validateFunctionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, nichandleDomainTaskTypeFunctionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NichandleDomainTask) validateFunction(formats strfmt.Registry) error {

	if err := validate.RequiredString("function", "body", string(m.Function)); err != nil {
		return err
	}

	// value enum
	if err := m.validateFunctionEnum("function", "body", m.Function); err != nil {
		return err
	}

	return nil
}

func (m *NichandleDomainTask) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *NichandleDomainTask) validateLastUpdate(formats strfmt.Registry) error {

	if err := validate.Required("lastUpdate", "body", strfmt.DateTime(m.LastUpdate)); err != nil {
		return err
	}

	if err := validate.FormatOf("lastUpdate", "body", "date-time", m.LastUpdate.String(), formats); err != nil {
		return err
	}

	return nil
}

var nichandleDomainTaskTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cancelled","doing","done","error","todo"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nichandleDomainTaskTypeStatusPropEnum = append(nichandleDomainTaskTypeStatusPropEnum, v)
	}
}

const (
	// NichandleDomainTaskStatusCancelled captures enum value "cancelled"
	NichandleDomainTaskStatusCancelled string = "cancelled"
	// NichandleDomainTaskStatusDoing captures enum value "doing"
	NichandleDomainTaskStatusDoing string = "doing"
	// NichandleDomainTaskStatusDone captures enum value "done"
	NichandleDomainTaskStatusDone string = "done"
	// NichandleDomainTaskStatusError captures enum value "error"
	NichandleDomainTaskStatusError string = "error"
	// NichandleDomainTaskStatusTodo captures enum value "todo"
	NichandleDomainTaskStatusTodo string = "todo"
)

// prop value enum
func (m *NichandleDomainTask) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, nichandleDomainTaskTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NichandleDomainTask) validateStatus(formats strfmt.Registry) error {

	if err := validate.RequiredString("status", "body", string(m.Status)); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *NichandleDomainTask) validateTodoDate(formats strfmt.Registry) error {

	if err := validate.Required("todoDate", "body", strfmt.DateTime(m.TodoDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("todoDate", "body", "date-time", m.TodoDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NichandleDomainTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NichandleDomainTask) UnmarshalBinary(b []byte) error {
	var res NichandleDomainTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
