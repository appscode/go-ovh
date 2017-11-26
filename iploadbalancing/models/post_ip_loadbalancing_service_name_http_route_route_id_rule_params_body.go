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

// PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBody post Ip loadbalancing service name Http route route Id rule params body
// swagger:model postIpLoadbalancingServiceNameHttpRouteRouteIdRuleParamsBody
type PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBody struct {

	// field
	// Required: true
	Field *string `json:"field"`

	// match
	// Required: true
	Match *string `json:"match"`

	// negate
	Negate *bool `json:"negate,omitempty"`

	// pattern
	Pattern string `json:"pattern,omitempty"`

	// sub field
	SubField string `json:"subField,omitempty"`
}

// Validate validates this post Ip loadbalancing service name Http route route Id rule params body
func (m *PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateField(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMatch(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBody) validateField(formats strfmt.Registry) error {

	if err := validate.Required("field", "body", m.Field); err != nil {
		return err
	}

	return nil
}

var postIpLoadbalancingServiceNameHttpRouteRouteIdRuleParamsBodyTypeMatchPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["contains","endswith","exists","in","internal","is","matches","startswith"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postIpLoadbalancingServiceNameHttpRouteRouteIdRuleParamsBodyTypeMatchPropEnum = append(postIpLoadbalancingServiceNameHttpRouteRouteIdRuleParamsBodyTypeMatchPropEnum, v)
	}
}

const (
	// PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBodyMatchContains captures enum value "contains"
	PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBodyMatchContains string = "contains"
	// PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBodyMatchEndswith captures enum value "endswith"
	PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBodyMatchEndswith string = "endswith"
	// PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBodyMatchExists captures enum value "exists"
	PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBodyMatchExists string = "exists"
	// PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBodyMatchIn captures enum value "in"
	PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBodyMatchIn string = "in"
	// PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBodyMatchInternal captures enum value "internal"
	PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBodyMatchInternal string = "internal"
	// PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBodyMatchIs captures enum value "is"
	PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBodyMatchIs string = "is"
	// PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBodyMatchMatches captures enum value "matches"
	PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBodyMatchMatches string = "matches"
	// PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBodyMatchStartswith captures enum value "startswith"
	PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBodyMatchStartswith string = "startswith"
)

// prop value enum
func (m *PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBody) validateMatchEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postIpLoadbalancingServiceNameHttpRouteRouteIdRuleParamsBodyTypeMatchPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBody) validateMatch(formats strfmt.Registry) error {

	if err := validate.Required("match", "body", m.Match); err != nil {
		return err
	}

	// value enum
	if err := m.validateMatchEnum("match", "body", *m.Match); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBody) UnmarshalBinary(b []byte) error {
	var res PostIPLoadbalancingServiceNameHTTPRouteRouteIDRuleParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}