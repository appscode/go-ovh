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

// NewGetIPLoadbalancingServiceNameAvailableRouteActionsParams creates a new GetIPLoadbalancingServiceNameAvailableRouteActionsParams object
// with the default values initialized.
func NewGetIPLoadbalancingServiceNameAvailableRouteActionsParams() *GetIPLoadbalancingServiceNameAvailableRouteActionsParams {
	var ()
	return &GetIPLoadbalancingServiceNameAvailableRouteActionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIPLoadbalancingServiceNameAvailableRouteActionsParamsWithTimeout creates a new GetIPLoadbalancingServiceNameAvailableRouteActionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIPLoadbalancingServiceNameAvailableRouteActionsParamsWithTimeout(timeout time.Duration) *GetIPLoadbalancingServiceNameAvailableRouteActionsParams {
	var ()
	return &GetIPLoadbalancingServiceNameAvailableRouteActionsParams{

		timeout: timeout,
	}
}

// NewGetIPLoadbalancingServiceNameAvailableRouteActionsParamsWithContext creates a new GetIPLoadbalancingServiceNameAvailableRouteActionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIPLoadbalancingServiceNameAvailableRouteActionsParamsWithContext(ctx context.Context) *GetIPLoadbalancingServiceNameAvailableRouteActionsParams {
	var ()
	return &GetIPLoadbalancingServiceNameAvailableRouteActionsParams{

		Context: ctx,
	}
}

// NewGetIPLoadbalancingServiceNameAvailableRouteActionsParamsWithHTTPClient creates a new GetIPLoadbalancingServiceNameAvailableRouteActionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIPLoadbalancingServiceNameAvailableRouteActionsParamsWithHTTPClient(client *http.Client) *GetIPLoadbalancingServiceNameAvailableRouteActionsParams {
	var ()
	return &GetIPLoadbalancingServiceNameAvailableRouteActionsParams{
		HTTPClient: client,
	}
}

/*GetIPLoadbalancingServiceNameAvailableRouteActionsParams contains all the parameters to send to the API endpoint
for the get IP loadbalancing service name available route actions operation typically these are written to a http.Request
*/
type GetIPLoadbalancingServiceNameAvailableRouteActionsParams struct {

	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get IP loadbalancing service name available route actions params
func (o *GetIPLoadbalancingServiceNameAvailableRouteActionsParams) WithTimeout(timeout time.Duration) *GetIPLoadbalancingServiceNameAvailableRouteActionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get IP loadbalancing service name available route actions params
func (o *GetIPLoadbalancingServiceNameAvailableRouteActionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get IP loadbalancing service name available route actions params
func (o *GetIPLoadbalancingServiceNameAvailableRouteActionsParams) WithContext(ctx context.Context) *GetIPLoadbalancingServiceNameAvailableRouteActionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get IP loadbalancing service name available route actions params
func (o *GetIPLoadbalancingServiceNameAvailableRouteActionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get IP loadbalancing service name available route actions params
func (o *GetIPLoadbalancingServiceNameAvailableRouteActionsParams) WithHTTPClient(client *http.Client) *GetIPLoadbalancingServiceNameAvailableRouteActionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get IP loadbalancing service name available route actions params
func (o *GetIPLoadbalancingServiceNameAvailableRouteActionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceName adds the serviceName to the get IP loadbalancing service name available route actions params
func (o *GetIPLoadbalancingServiceNameAvailableRouteActionsParams) WithServiceName(serviceName string) *GetIPLoadbalancingServiceNameAvailableRouteActionsParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get IP loadbalancing service name available route actions params
func (o *GetIPLoadbalancingServiceNameAvailableRouteActionsParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetIPLoadbalancingServiceNameAvailableRouteActionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
