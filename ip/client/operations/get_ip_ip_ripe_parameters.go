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

// NewGetIPIPRipeParams creates a new GetIPIPRipeParams object
// with the default values initialized.
func NewGetIPIPRipeParams() *GetIPIPRipeParams {
	var ()
	return &GetIPIPRipeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIPIPRipeParamsWithTimeout creates a new GetIPIPRipeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIPIPRipeParamsWithTimeout(timeout time.Duration) *GetIPIPRipeParams {
	var ()
	return &GetIPIPRipeParams{

		timeout: timeout,
	}
}

// NewGetIPIPRipeParamsWithContext creates a new GetIPIPRipeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIPIPRipeParamsWithContext(ctx context.Context) *GetIPIPRipeParams {
	var ()
	return &GetIPIPRipeParams{

		Context: ctx,
	}
}

// NewGetIPIPRipeParamsWithHTTPClient creates a new GetIPIPRipeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIPIPRipeParamsWithHTTPClient(client *http.Client) *GetIPIPRipeParams {
	var ()
	return &GetIPIPRipeParams{
		HTTPClient: client,
	}
}

/*GetIPIPRipeParams contains all the parameters to send to the API endpoint
for the get IP IP ripe operation typically these are written to a http.Request
*/
type GetIPIPRipeParams struct {

	/*IP*/
	IP string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get IP IP ripe params
func (o *GetIPIPRipeParams) WithTimeout(timeout time.Duration) *GetIPIPRipeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get IP IP ripe params
func (o *GetIPIPRipeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get IP IP ripe params
func (o *GetIPIPRipeParams) WithContext(ctx context.Context) *GetIPIPRipeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get IP IP ripe params
func (o *GetIPIPRipeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get IP IP ripe params
func (o *GetIPIPRipeParams) WithHTTPClient(client *http.Client) *GetIPIPRipeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get IP IP ripe params
func (o *GetIPIPRipeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIP adds the ip to the get IP IP ripe params
func (o *GetIPIPRipeParams) WithIP(ip string) *GetIPIPRipeParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the get IP IP ripe params
func (o *GetIPIPRipeParams) SetIP(ip string) {
	o.IP = ip
}

// WriteToRequest writes these params to a swagger request
func (o *GetIPIPRipeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ip
	if err := r.SetPathParam("ip", o.IP); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
