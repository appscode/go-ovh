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

// NewGetDomainParams creates a new GetDomainParams object
// with the default values initialized.
func NewGetDomainParams() *GetDomainParams {
	var ()
	return &GetDomainParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomainParamsWithTimeout creates a new GetDomainParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDomainParamsWithTimeout(timeout time.Duration) *GetDomainParams {
	var ()
	return &GetDomainParams{

		timeout: timeout,
	}
}

// NewGetDomainParamsWithContext creates a new GetDomainParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDomainParamsWithContext(ctx context.Context) *GetDomainParams {
	var ()
	return &GetDomainParams{

		Context: ctx,
	}
}

// NewGetDomainParamsWithHTTPClient creates a new GetDomainParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDomainParamsWithHTTPClient(client *http.Client) *GetDomainParams {
	var ()
	return &GetDomainParams{
		HTTPClient: client,
	}
}

/*GetDomainParams contains all the parameters to send to the API endpoint
for the get domain operation typically these are written to a http.Request
*/
type GetDomainParams struct {

	/*WhoisOwner*/
	WhoisOwner *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get domain params
func (o *GetDomainParams) WithTimeout(timeout time.Duration) *GetDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domain params
func (o *GetDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domain params
func (o *GetDomainParams) WithContext(ctx context.Context) *GetDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domain params
func (o *GetDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domain params
func (o *GetDomainParams) WithHTTPClient(client *http.Client) *GetDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domain params
func (o *GetDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWhoisOwner adds the whoisOwner to the get domain params
func (o *GetDomainParams) WithWhoisOwner(whoisOwner *string) *GetDomainParams {
	o.SetWhoisOwner(whoisOwner)
	return o
}

// SetWhoisOwner adds the whoisOwner to the get domain params
func (o *GetDomainParams) SetWhoisOwner(whoisOwner *string) {
	o.WhoisOwner = whoisOwner
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.WhoisOwner != nil {

		// query param whoisOwner
		var qrWhoisOwner string
		if o.WhoisOwner != nil {
			qrWhoisOwner = *o.WhoisOwner
		}
		qWhoisOwner := qrWhoisOwner
		if qWhoisOwner != "" {
			if err := r.SetQueryParam("whoisOwner", qWhoisOwner); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
