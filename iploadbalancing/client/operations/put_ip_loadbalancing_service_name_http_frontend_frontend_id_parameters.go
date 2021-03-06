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

// NewPutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams creates a new PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams object
// with the default values initialized.
func NewPutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams() *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams {
	var ()
	return &PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParamsWithTimeout creates a new PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParamsWithTimeout(timeout time.Duration) *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams {
	var ()
	return &PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams{

		timeout: timeout,
	}
}

// NewPutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParamsWithContext creates a new PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParamsWithContext(ctx context.Context) *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams {
	var ()
	return &PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams{

		Context: ctx,
	}
}

// NewPutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParamsWithHTTPClient creates a new PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParamsWithHTTPClient(client *http.Client) *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams {
	var ()
	return &PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams{
		HTTPClient: client,
	}
}

/*PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams contains all the parameters to send to the API endpoint
for the put IP loadbalancing service name HTTP frontend frontend ID operation typically these are written to a http.Request
*/
type PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams struct {

	/*FrontendID*/
	FrontendID int64
	/*IPLBHTTPFrontendPut*/
	IPLBHTTPFrontendPut *models.IPLBFrontendHTTPFrontendHTTP
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put IP loadbalancing service name HTTP frontend frontend ID params
func (o *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) WithTimeout(timeout time.Duration) *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put IP loadbalancing service name HTTP frontend frontend ID params
func (o *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put IP loadbalancing service name HTTP frontend frontend ID params
func (o *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) WithContext(ctx context.Context) *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put IP loadbalancing service name HTTP frontend frontend ID params
func (o *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put IP loadbalancing service name HTTP frontend frontend ID params
func (o *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) WithHTTPClient(client *http.Client) *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put IP loadbalancing service name HTTP frontend frontend ID params
func (o *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrontendID adds the frontendID to the put IP loadbalancing service name HTTP frontend frontend ID params
func (o *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) WithFrontendID(frontendID int64) *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams {
	o.SetFrontendID(frontendID)
	return o
}

// SetFrontendID adds the frontendId to the put IP loadbalancing service name HTTP frontend frontend ID params
func (o *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) SetFrontendID(frontendID int64) {
	o.FrontendID = frontendID
}

// WithIPLBHTTPFrontendPut adds the iPLBHTTPFrontendPut to the put IP loadbalancing service name HTTP frontend frontend ID params
func (o *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) WithIPLBHTTPFrontendPut(iPLBHTTPFrontendPut *models.IPLBFrontendHTTPFrontendHTTP) *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams {
	o.SetIPLBHTTPFrontendPut(iPLBHTTPFrontendPut)
	return o
}

// SetIPLBHTTPFrontendPut adds the iplbHttpFrontendPut to the put IP loadbalancing service name HTTP frontend frontend ID params
func (o *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) SetIPLBHTTPFrontendPut(iPLBHTTPFrontendPut *models.IPLBFrontendHTTPFrontendHTTP) {
	o.IPLBHTTPFrontendPut = iPLBHTTPFrontendPut
}

// WithServiceName adds the serviceName to the put IP loadbalancing service name HTTP frontend frontend ID params
func (o *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) WithServiceName(serviceName string) *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the put IP loadbalancing service name HTTP frontend frontend ID params
func (o *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PutIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param frontendId
	if err := r.SetPathParam("frontendId", swag.FormatInt64(o.FrontendID)); err != nil {
		return err
	}

	if o.IPLBHTTPFrontendPut != nil {
		if err := r.SetBodyParam(o.IPLBHTTPFrontendPut); err != nil {
			return err
		}
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
