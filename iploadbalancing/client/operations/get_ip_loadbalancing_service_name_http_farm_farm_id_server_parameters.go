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
)

// NewGetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams creates a new GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams object
// with the default values initialized.
func NewGetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams() *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	var ()
	return &GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParamsWithTimeout creates a new GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParamsWithTimeout(timeout time.Duration) *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	var ()
	return &GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams{

		timeout: timeout,
	}
}

// NewGetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParamsWithContext creates a new GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParamsWithContext(ctx context.Context) *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	var ()
	return &GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams{

		Context: ctx,
	}
}

// NewGetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParamsWithHTTPClient creates a new GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParamsWithHTTPClient(client *http.Client) *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	var ()
	return &GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams{
		HTTPClient: client,
	}
}

/*GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams contains all the parameters to send to the API endpoint
for the get IP loadbalancing service name HTTP farm farm ID server operation typically these are written to a http.Request
*/
type GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams struct {

	/*Address*/
	Address *string
	/*Cookie*/
	Cookie *string
	/*FarmID*/
	FarmID int64
	/*ServiceName*/
	ServiceName string
	/*Status*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get IP loadbalancing service name HTTP farm farm ID server params
func (o *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) WithTimeout(timeout time.Duration) *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get IP loadbalancing service name HTTP farm farm ID server params
func (o *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get IP loadbalancing service name HTTP farm farm ID server params
func (o *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) WithContext(ctx context.Context) *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get IP loadbalancing service name HTTP farm farm ID server params
func (o *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get IP loadbalancing service name HTTP farm farm ID server params
func (o *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) WithHTTPClient(client *http.Client) *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get IP loadbalancing service name HTTP farm farm ID server params
func (o *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the get IP loadbalancing service name HTTP farm farm ID server params
func (o *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) WithAddress(address *string) *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the get IP loadbalancing service name HTTP farm farm ID server params
func (o *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) SetAddress(address *string) {
	o.Address = address
}

// WithCookie adds the cookie to the get IP loadbalancing service name HTTP farm farm ID server params
func (o *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) WithCookie(cookie *string) *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	o.SetCookie(cookie)
	return o
}

// SetCookie adds the cookie to the get IP loadbalancing service name HTTP farm farm ID server params
func (o *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) SetCookie(cookie *string) {
	o.Cookie = cookie
}

// WithFarmID adds the farmID to the get IP loadbalancing service name HTTP farm farm ID server params
func (o *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) WithFarmID(farmID int64) *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	o.SetFarmID(farmID)
	return o
}

// SetFarmID adds the farmId to the get IP loadbalancing service name HTTP farm farm ID server params
func (o *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) SetFarmID(farmID int64) {
	o.FarmID = farmID
}

// WithServiceName adds the serviceName to the get IP loadbalancing service name HTTP farm farm ID server params
func (o *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) WithServiceName(serviceName string) *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get IP loadbalancing service name HTTP farm farm ID server params
func (o *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WithStatus adds the status to the get IP loadbalancing service name HTTP farm farm ID server params
func (o *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) WithStatus(status *string) *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get IP loadbalancing service name HTTP farm farm ID server params
func (o *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Address != nil {

		// query param address
		var qrAddress string
		if o.Address != nil {
			qrAddress = *o.Address
		}
		qAddress := qrAddress
		if qAddress != "" {
			if err := r.SetQueryParam("address", qAddress); err != nil {
				return err
			}
		}

	}

	if o.Cookie != nil {

		// query param cookie
		var qrCookie string
		if o.Cookie != nil {
			qrCookie = *o.Cookie
		}
		qCookie := qrCookie
		if qCookie != "" {
			if err := r.SetQueryParam("cookie", qCookie); err != nil {
				return err
			}
		}

	}

	// path param farmId
	if err := r.SetPathParam("farmId", swag.FormatInt64(o.FarmID)); err != nil {
		return err
	}

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}