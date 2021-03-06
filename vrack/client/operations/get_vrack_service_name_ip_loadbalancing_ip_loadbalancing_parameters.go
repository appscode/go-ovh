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

// NewGetVrackServiceNameIPLoadbalancingIPLoadbalancingParams creates a new GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams object
// with the default values initialized.
func NewGetVrackServiceNameIPLoadbalancingIPLoadbalancingParams() *GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams {
	var ()
	return &GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVrackServiceNameIPLoadbalancingIPLoadbalancingParamsWithTimeout creates a new GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVrackServiceNameIPLoadbalancingIPLoadbalancingParamsWithTimeout(timeout time.Duration) *GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams {
	var ()
	return &GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams{

		timeout: timeout,
	}
}

// NewGetVrackServiceNameIPLoadbalancingIPLoadbalancingParamsWithContext creates a new GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVrackServiceNameIPLoadbalancingIPLoadbalancingParamsWithContext(ctx context.Context) *GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams {
	var ()
	return &GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams{

		Context: ctx,
	}
}

// NewGetVrackServiceNameIPLoadbalancingIPLoadbalancingParamsWithHTTPClient creates a new GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVrackServiceNameIPLoadbalancingIPLoadbalancingParamsWithHTTPClient(client *http.Client) *GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams {
	var ()
	return &GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams{
		HTTPClient: client,
	}
}

/*GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams contains all the parameters to send to the API endpoint
for the get vrack service name IP loadbalancing IP loadbalancing operation typically these are written to a http.Request
*/
type GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams struct {

	/*IPLoadbalancing*/
	IPLoadbalancing string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vrack service name IP loadbalancing IP loadbalancing params
func (o *GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams) WithTimeout(timeout time.Duration) *GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vrack service name IP loadbalancing IP loadbalancing params
func (o *GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vrack service name IP loadbalancing IP loadbalancing params
func (o *GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams) WithContext(ctx context.Context) *GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vrack service name IP loadbalancing IP loadbalancing params
func (o *GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vrack service name IP loadbalancing IP loadbalancing params
func (o *GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams) WithHTTPClient(client *http.Client) *GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vrack service name IP loadbalancing IP loadbalancing params
func (o *GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIPLoadbalancing adds the iPLoadbalancing to the get vrack service name IP loadbalancing IP loadbalancing params
func (o *GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams) WithIPLoadbalancing(iPLoadbalancing string) *GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams {
	o.SetIPLoadbalancing(iPLoadbalancing)
	return o
}

// SetIPLoadbalancing adds the ipLoadbalancing to the get vrack service name IP loadbalancing IP loadbalancing params
func (o *GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams) SetIPLoadbalancing(iPLoadbalancing string) {
	o.IPLoadbalancing = iPLoadbalancing
}

// WithServiceName adds the serviceName to the get vrack service name IP loadbalancing IP loadbalancing params
func (o *GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams) WithServiceName(serviceName string) *GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get vrack service name IP loadbalancing IP loadbalancing params
func (o *GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetVrackServiceNameIPLoadbalancingIPLoadbalancingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ipLoadbalancing
	if err := r.SetPathParam("ipLoadbalancing", o.IPLoadbalancing); err != nil {
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
