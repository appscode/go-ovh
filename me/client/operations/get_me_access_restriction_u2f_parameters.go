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

// NewGetMeAccessRestrictionU2fParams creates a new GetMeAccessRestrictionU2fParams object
// with the default values initialized.
func NewGetMeAccessRestrictionU2fParams() *GetMeAccessRestrictionU2fParams {

	return &GetMeAccessRestrictionU2fParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeAccessRestrictionU2fParamsWithTimeout creates a new GetMeAccessRestrictionU2fParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeAccessRestrictionU2fParamsWithTimeout(timeout time.Duration) *GetMeAccessRestrictionU2fParams {

	return &GetMeAccessRestrictionU2fParams{

		timeout: timeout,
	}
}

// NewGetMeAccessRestrictionU2fParamsWithContext creates a new GetMeAccessRestrictionU2fParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeAccessRestrictionU2fParamsWithContext(ctx context.Context) *GetMeAccessRestrictionU2fParams {

	return &GetMeAccessRestrictionU2fParams{

		Context: ctx,
	}
}

// NewGetMeAccessRestrictionU2fParamsWithHTTPClient creates a new GetMeAccessRestrictionU2fParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeAccessRestrictionU2fParamsWithHTTPClient(client *http.Client) *GetMeAccessRestrictionU2fParams {

	return &GetMeAccessRestrictionU2fParams{
		HTTPClient: client,
	}
}

/*GetMeAccessRestrictionU2fParams contains all the parameters to send to the API endpoint
for the get me access restriction u2f operation typically these are written to a http.Request
*/
type GetMeAccessRestrictionU2fParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me access restriction u2f params
func (o *GetMeAccessRestrictionU2fParams) WithTimeout(timeout time.Duration) *GetMeAccessRestrictionU2fParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me access restriction u2f params
func (o *GetMeAccessRestrictionU2fParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me access restriction u2f params
func (o *GetMeAccessRestrictionU2fParams) WithContext(ctx context.Context) *GetMeAccessRestrictionU2fParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me access restriction u2f params
func (o *GetMeAccessRestrictionU2fParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me access restriction u2f params
func (o *GetMeAccessRestrictionU2fParams) WithHTTPClient(client *http.Client) *GetMeAccessRestrictionU2fParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me access restriction u2f params
func (o *GetMeAccessRestrictionU2fParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeAccessRestrictionU2fParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
