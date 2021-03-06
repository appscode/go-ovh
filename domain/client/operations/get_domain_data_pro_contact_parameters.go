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

// NewGetDomainDataProContactParams creates a new GetDomainDataProContactParams object
// with the default values initialized.
func NewGetDomainDataProContactParams() *GetDomainDataProContactParams {

	return &GetDomainDataProContactParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomainDataProContactParamsWithTimeout creates a new GetDomainDataProContactParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDomainDataProContactParamsWithTimeout(timeout time.Duration) *GetDomainDataProContactParams {

	return &GetDomainDataProContactParams{

		timeout: timeout,
	}
}

// NewGetDomainDataProContactParamsWithContext creates a new GetDomainDataProContactParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDomainDataProContactParamsWithContext(ctx context.Context) *GetDomainDataProContactParams {

	return &GetDomainDataProContactParams{

		Context: ctx,
	}
}

// NewGetDomainDataProContactParamsWithHTTPClient creates a new GetDomainDataProContactParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDomainDataProContactParamsWithHTTPClient(client *http.Client) *GetDomainDataProContactParams {

	return &GetDomainDataProContactParams{
		HTTPClient: client,
	}
}

/*GetDomainDataProContactParams contains all the parameters to send to the API endpoint
for the get domain data pro contact operation typically these are written to a http.Request
*/
type GetDomainDataProContactParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get domain data pro contact params
func (o *GetDomainDataProContactParams) WithTimeout(timeout time.Duration) *GetDomainDataProContactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domain data pro contact params
func (o *GetDomainDataProContactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domain data pro contact params
func (o *GetDomainDataProContactParams) WithContext(ctx context.Context) *GetDomainDataProContactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domain data pro contact params
func (o *GetDomainDataProContactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domain data pro contact params
func (o *GetDomainDataProContactParams) WithHTTPClient(client *http.Client) *GetDomainDataProContactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domain data pro contact params
func (o *GetDomainDataProContactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomainDataProContactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
