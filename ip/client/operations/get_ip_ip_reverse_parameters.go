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

// NewGetIPIPReverseParams creates a new GetIPIPReverseParams object
// with the default values initialized.
func NewGetIPIPReverseParams() *GetIPIPReverseParams {
	var ()
	return &GetIPIPReverseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIPIPReverseParamsWithTimeout creates a new GetIPIPReverseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIPIPReverseParamsWithTimeout(timeout time.Duration) *GetIPIPReverseParams {
	var ()
	return &GetIPIPReverseParams{

		timeout: timeout,
	}
}

// NewGetIPIPReverseParamsWithContext creates a new GetIPIPReverseParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIPIPReverseParamsWithContext(ctx context.Context) *GetIPIPReverseParams {
	var ()
	return &GetIPIPReverseParams{

		Context: ctx,
	}
}

// NewGetIPIPReverseParamsWithHTTPClient creates a new GetIPIPReverseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIPIPReverseParamsWithHTTPClient(client *http.Client) *GetIPIPReverseParams {
	var ()
	return &GetIPIPReverseParams{
		HTTPClient: client,
	}
}

/*GetIPIPReverseParams contains all the parameters to send to the API endpoint
for the get IP IP reverse operation typically these are written to a http.Request
*/
type GetIPIPReverseParams struct {

	/*IP*/
	IP string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get IP IP reverse params
func (o *GetIPIPReverseParams) WithTimeout(timeout time.Duration) *GetIPIPReverseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get IP IP reverse params
func (o *GetIPIPReverseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get IP IP reverse params
func (o *GetIPIPReverseParams) WithContext(ctx context.Context) *GetIPIPReverseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get IP IP reverse params
func (o *GetIPIPReverseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get IP IP reverse params
func (o *GetIPIPReverseParams) WithHTTPClient(client *http.Client) *GetIPIPReverseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get IP IP reverse params
func (o *GetIPIPReverseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIP adds the ip to the get IP IP reverse params
func (o *GetIPIPReverseParams) WithIP(ip string) *GetIPIPReverseParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the get IP IP reverse params
func (o *GetIPIPReverseParams) SetIP(ip string) {
	o.IP = ip
}

// WriteToRequest writes these params to a swagger request
func (o *GetIPIPReverseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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