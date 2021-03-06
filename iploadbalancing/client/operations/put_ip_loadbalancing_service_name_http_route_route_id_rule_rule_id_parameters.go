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

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/appscode/go-ovh/iploadbalancing/models"
)

// NewPutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams creates a new PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams object
// with the default values initialized.
func NewPutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams() *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams {
	var ()
	return &PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParamsWithTimeout creates a new PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParamsWithTimeout(timeout time.Duration) *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams {
	var ()
	return &PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams{

		timeout: timeout,
	}
}

// NewPutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParamsWithContext creates a new PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParamsWithContext(ctx context.Context) *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams {
	var ()
	return &PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams{

		Context: ctx,
	}
}

// NewPutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParamsWithHTTPClient creates a new PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParamsWithHTTPClient(client *http.Client) *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams {
	var ()
	return &PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams{
		HTTPClient: client,
	}
}

/*PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams contains all the parameters to send to the API endpoint
for the put IP loadbalancing service name HTTP route route ID rule rule ID operation typically these are written to a http.Request
*/
type PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams struct {

	/*IPLBHTTPRouteRulePut*/
	IPLBHTTPRouteRulePut *models.IPLBRouteRuleRouteRule
	/*RouteID*/
	RouteID int64
	/*RuleID*/
	RuleID int64
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put IP loadbalancing service name HTTP route route ID rule rule ID params
func (o *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams) WithTimeout(timeout time.Duration) *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put IP loadbalancing service name HTTP route route ID rule rule ID params
func (o *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put IP loadbalancing service name HTTP route route ID rule rule ID params
func (o *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams) WithContext(ctx context.Context) *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put IP loadbalancing service name HTTP route route ID rule rule ID params
func (o *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put IP loadbalancing service name HTTP route route ID rule rule ID params
func (o *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams) WithHTTPClient(client *http.Client) *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put IP loadbalancing service name HTTP route route ID rule rule ID params
func (o *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIPLBHTTPRouteRulePut adds the iPLBHTTPRouteRulePut to the put IP loadbalancing service name HTTP route route ID rule rule ID params
func (o *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams) WithIPLBHTTPRouteRulePut(iPLBHTTPRouteRulePut *models.IPLBRouteRuleRouteRule) *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams {
	o.SetIPLBHTTPRouteRulePut(iPLBHTTPRouteRulePut)
	return o
}

// SetIPLBHTTPRouteRulePut adds the iplbHttpRouteRulePut to the put IP loadbalancing service name HTTP route route ID rule rule ID params
func (o *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams) SetIPLBHTTPRouteRulePut(iPLBHTTPRouteRulePut *models.IPLBRouteRuleRouteRule) {
	o.IPLBHTTPRouteRulePut = iPLBHTTPRouteRulePut
}

// WithRouteID adds the routeID to the put IP loadbalancing service name HTTP route route ID rule rule ID params
func (o *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams) WithRouteID(routeID int64) *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams {
	o.SetRouteID(routeID)
	return o
}

// SetRouteID adds the routeId to the put IP loadbalancing service name HTTP route route ID rule rule ID params
func (o *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams) SetRouteID(routeID int64) {
	o.RouteID = routeID
}

// WithRuleID adds the ruleID to the put IP loadbalancing service name HTTP route route ID rule rule ID params
func (o *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams) WithRuleID(ruleID int64) *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the put IP loadbalancing service name HTTP route route ID rule rule ID params
func (o *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams) SetRuleID(ruleID int64) {
	o.RuleID = ruleID
}

// WithServiceName adds the serviceName to the put IP loadbalancing service name HTTP route route ID rule rule ID params
func (o *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams) WithServiceName(serviceName string) *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the put IP loadbalancing service name HTTP route route ID rule rule ID params
func (o *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PutIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IPLBHTTPRouteRulePut != nil {
		if err := r.SetBodyParam(o.IPLBHTTPRouteRulePut); err != nil {
			return err
		}
	}

	// path param routeId
	if err := r.SetPathParam("routeId", swag.FormatInt64(o.RouteID)); err != nil {
		return err
	}

	// path param ruleId
	if err := r.SetPathParam("ruleId", swag.FormatInt64(o.RuleID)); err != nil {
		return err
	}

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
