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

// NewGetVpsServiceNameOptionOptionParams creates a new GetVpsServiceNameOptionOptionParams object
// with the default values initialized.
func NewGetVpsServiceNameOptionOptionParams() *GetVpsServiceNameOptionOptionParams {
	var ()
	return &GetVpsServiceNameOptionOptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVpsServiceNameOptionOptionParamsWithTimeout creates a new GetVpsServiceNameOptionOptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVpsServiceNameOptionOptionParamsWithTimeout(timeout time.Duration) *GetVpsServiceNameOptionOptionParams {
	var ()
	return &GetVpsServiceNameOptionOptionParams{

		timeout: timeout,
	}
}

// NewGetVpsServiceNameOptionOptionParamsWithContext creates a new GetVpsServiceNameOptionOptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVpsServiceNameOptionOptionParamsWithContext(ctx context.Context) *GetVpsServiceNameOptionOptionParams {
	var ()
	return &GetVpsServiceNameOptionOptionParams{

		Context: ctx,
	}
}

// NewGetVpsServiceNameOptionOptionParamsWithHTTPClient creates a new GetVpsServiceNameOptionOptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVpsServiceNameOptionOptionParamsWithHTTPClient(client *http.Client) *GetVpsServiceNameOptionOptionParams {
	var ()
	return &GetVpsServiceNameOptionOptionParams{
		HTTPClient: client,
	}
}

/*GetVpsServiceNameOptionOptionParams contains all the parameters to send to the API endpoint
for the get vps service name option option operation typically these are written to a http.Request
*/
type GetVpsServiceNameOptionOptionParams struct {

	/*Option*/
	Option string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vps service name option option params
func (o *GetVpsServiceNameOptionOptionParams) WithTimeout(timeout time.Duration) *GetVpsServiceNameOptionOptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vps service name option option params
func (o *GetVpsServiceNameOptionOptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vps service name option option params
func (o *GetVpsServiceNameOptionOptionParams) WithContext(ctx context.Context) *GetVpsServiceNameOptionOptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vps service name option option params
func (o *GetVpsServiceNameOptionOptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vps service name option option params
func (o *GetVpsServiceNameOptionOptionParams) WithHTTPClient(client *http.Client) *GetVpsServiceNameOptionOptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vps service name option option params
func (o *GetVpsServiceNameOptionOptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOption adds the option to the get vps service name option option params
func (o *GetVpsServiceNameOptionOptionParams) WithOption(option string) *GetVpsServiceNameOptionOptionParams {
	o.SetOption(option)
	return o
}

// SetOption adds the option to the get vps service name option option params
func (o *GetVpsServiceNameOptionOptionParams) SetOption(option string) {
	o.Option = option
}

// WithServiceName adds the serviceName to the get vps service name option option params
func (o *GetVpsServiceNameOptionOptionParams) WithServiceName(serviceName string) *GetVpsServiceNameOptionOptionParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get vps service name option option params
func (o *GetVpsServiceNameOptionOptionParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetVpsServiceNameOptionOptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param option
	if err := r.SetPathParam("option", o.Option); err != nil {
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