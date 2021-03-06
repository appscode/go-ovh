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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetIPLoadbalancingServiceNameUDPFarmParams creates a new GetIPLoadbalancingServiceNameUDPFarmParams object
// with the default values initialized.
func NewGetIPLoadbalancingServiceNameUDPFarmParams() *GetIPLoadbalancingServiceNameUDPFarmParams {
	var ()
	return &GetIPLoadbalancingServiceNameUDPFarmParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIPLoadbalancingServiceNameUDPFarmParamsWithTimeout creates a new GetIPLoadbalancingServiceNameUDPFarmParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIPLoadbalancingServiceNameUDPFarmParamsWithTimeout(timeout time.Duration) *GetIPLoadbalancingServiceNameUDPFarmParams {
	var ()
	return &GetIPLoadbalancingServiceNameUDPFarmParams{

		timeout: timeout,
	}
}

// NewGetIPLoadbalancingServiceNameUDPFarmParamsWithContext creates a new GetIPLoadbalancingServiceNameUDPFarmParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIPLoadbalancingServiceNameUDPFarmParamsWithContext(ctx context.Context) *GetIPLoadbalancingServiceNameUDPFarmParams {
	var ()
	return &GetIPLoadbalancingServiceNameUDPFarmParams{

		Context: ctx,
	}
}

// NewGetIPLoadbalancingServiceNameUDPFarmParamsWithHTTPClient creates a new GetIPLoadbalancingServiceNameUDPFarmParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIPLoadbalancingServiceNameUDPFarmParamsWithHTTPClient(client *http.Client) *GetIPLoadbalancingServiceNameUDPFarmParams {
	var ()
	return &GetIPLoadbalancingServiceNameUDPFarmParams{
		HTTPClient: client,
	}
}

/*GetIPLoadbalancingServiceNameUDPFarmParams contains all the parameters to send to the API endpoint
for the get IP loadbalancing service name UDP farm operation typically these are written to a http.Request
*/
type GetIPLoadbalancingServiceNameUDPFarmParams struct {

	/*ServiceName*/
	ServiceName string
	/*Zone*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get IP loadbalancing service name UDP farm params
func (o *GetIPLoadbalancingServiceNameUDPFarmParams) WithTimeout(timeout time.Duration) *GetIPLoadbalancingServiceNameUDPFarmParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get IP loadbalancing service name UDP farm params
func (o *GetIPLoadbalancingServiceNameUDPFarmParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get IP loadbalancing service name UDP farm params
func (o *GetIPLoadbalancingServiceNameUDPFarmParams) WithContext(ctx context.Context) *GetIPLoadbalancingServiceNameUDPFarmParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get IP loadbalancing service name UDP farm params
func (o *GetIPLoadbalancingServiceNameUDPFarmParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get IP loadbalancing service name UDP farm params
func (o *GetIPLoadbalancingServiceNameUDPFarmParams) WithHTTPClient(client *http.Client) *GetIPLoadbalancingServiceNameUDPFarmParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get IP loadbalancing service name UDP farm params
func (o *GetIPLoadbalancingServiceNameUDPFarmParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceName adds the serviceName to the get IP loadbalancing service name UDP farm params
func (o *GetIPLoadbalancingServiceNameUDPFarmParams) WithServiceName(serviceName string) *GetIPLoadbalancingServiceNameUDPFarmParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get IP loadbalancing service name UDP farm params
func (o *GetIPLoadbalancingServiceNameUDPFarmParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WithZone adds the zone to the get IP loadbalancing service name UDP farm params
func (o *GetIPLoadbalancingServiceNameUDPFarmParams) WithZone(zone *string) *GetIPLoadbalancingServiceNameUDPFarmParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the get IP loadbalancing service name UDP farm params
func (o *GetIPLoadbalancingServiceNameUDPFarmParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *GetIPLoadbalancingServiceNameUDPFarmParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	if o.Zone != nil {

		// query param zone
		var qrZone string
		if o.Zone != nil {
			qrZone = *o.Zone
		}
		qZone := qrZone
		if qZone != "" {
			if err := r.SetQueryParam("zone", qZone); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
