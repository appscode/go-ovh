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

// PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBody post Ip loadbalancing service name Tcp farm farm Id server params body
// swagger:model postIpLoadbalancingServiceNameTcpFarmFarmIdServerParamsBody
type PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBody struct {

	// address
	// Required: true
	Address *string `json:"address"`

	// backup
	Backup *bool `json:"backup,omitempty"`

	// chain
	Chain string `json:"chain,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// port
	Port int64 `json:"port,omitempty"`

	// probe
	Probe *bool `json:"probe,omitempty"`

	// proxy protocol version
	ProxyProtocolVersion string `json:"proxyProtocolVersion,omitempty"`

	// ssl
	Ssl *bool `json:"ssl,omitempty"`

	// status
	// Required: true
	Status *string `json:"status"`

	// weight
	Weight int64 `json:"weight,omitempty"`
}

// Validate validates this post Ip loadbalancing service name Tcp farm farm Id server params body
func (m *PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProxyProtocolVersion(formats); err != nil {
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

func (m *PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBody) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

var postIpLoadbalancingServiceNameTcpFarmFarmIdServerParamsBodyTypeProxyProtocolVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["v1","v2","v2-ssl","v2-ssl-cn"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postIpLoadbalancingServiceNameTcpFarmFarmIdServerParamsBodyTypeProxyProtocolVersionPropEnum = append(postIpLoadbalancingServiceNameTcpFarmFarmIdServerParamsBodyTypeProxyProtocolVersionPropEnum, v)
	}
}

const (
	// PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBodyProxyProtocolVersionV1 captures enum value "v1"
	PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBodyProxyProtocolVersionV1 string = "v1"
	// PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBodyProxyProtocolVersionV2 captures enum value "v2"
	PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBodyProxyProtocolVersionV2 string = "v2"
	// PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBodyProxyProtocolVersionV2Ssl captures enum value "v2-ssl"
	PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBodyProxyProtocolVersionV2Ssl string = "v2-ssl"
	// PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBodyProxyProtocolVersionV2SslCn captures enum value "v2-ssl-cn"
	PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBodyProxyProtocolVersionV2SslCn string = "v2-ssl-cn"
)

// prop value enum
func (m *PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBody) validateProxyProtocolVersionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postIpLoadbalancingServiceNameTcpFarmFarmIdServerParamsBodyTypeProxyProtocolVersionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBody) validateProxyProtocolVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.ProxyProtocolVersion) { // not required
		return nil
	}

	// value enum
	if err := m.validateProxyProtocolVersionEnum("proxyProtocolVersion", "body", m.ProxyProtocolVersion); err != nil {
		return err
	}

	return nil
}

var postIpLoadbalancingServiceNameTcpFarmFarmIdServerParamsBodyTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postIpLoadbalancingServiceNameTcpFarmFarmIdServerParamsBodyTypeStatusPropEnum = append(postIpLoadbalancingServiceNameTcpFarmFarmIdServerParamsBodyTypeStatusPropEnum, v)
	}
}

const (
	// PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBodyStatusActive captures enum value "active"
	PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBodyStatusActive string = "active"
	// PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBodyStatusInactive captures enum value "inactive"
	PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBodyStatusInactive string = "inactive"
)

// prop value enum
func (m *PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBody) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postIpLoadbalancingServiceNameTcpFarmFarmIdServerParamsBodyTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBody) UnmarshalBinary(b []byte) error {
	var res PostIPLoadbalancingServiceNameTCPFarmFarmIDServerParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}