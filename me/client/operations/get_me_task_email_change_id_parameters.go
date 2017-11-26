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

// NewGetMeTaskEmailChangeIDParams creates a new GetMeTaskEmailChangeIDParams object
// with the default values initialized.
func NewGetMeTaskEmailChangeIDParams() *GetMeTaskEmailChangeIDParams {
	var ()
	return &GetMeTaskEmailChangeIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeTaskEmailChangeIDParamsWithTimeout creates a new GetMeTaskEmailChangeIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeTaskEmailChangeIDParamsWithTimeout(timeout time.Duration) *GetMeTaskEmailChangeIDParams {
	var ()
	return &GetMeTaskEmailChangeIDParams{

		timeout: timeout,
	}
}

// NewGetMeTaskEmailChangeIDParamsWithContext creates a new GetMeTaskEmailChangeIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeTaskEmailChangeIDParamsWithContext(ctx context.Context) *GetMeTaskEmailChangeIDParams {
	var ()
	return &GetMeTaskEmailChangeIDParams{

		Context: ctx,
	}
}

// NewGetMeTaskEmailChangeIDParamsWithHTTPClient creates a new GetMeTaskEmailChangeIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeTaskEmailChangeIDParamsWithHTTPClient(client *http.Client) *GetMeTaskEmailChangeIDParams {
	var ()
	return &GetMeTaskEmailChangeIDParams{
		HTTPClient: client,
	}
}

/*GetMeTaskEmailChangeIDParams contains all the parameters to send to the API endpoint
for the get me task email change ID operation typically these are written to a http.Request
*/
type GetMeTaskEmailChangeIDParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me task email change ID params
func (o *GetMeTaskEmailChangeIDParams) WithTimeout(timeout time.Duration) *GetMeTaskEmailChangeIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me task email change ID params
func (o *GetMeTaskEmailChangeIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me task email change ID params
func (o *GetMeTaskEmailChangeIDParams) WithContext(ctx context.Context) *GetMeTaskEmailChangeIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me task email change ID params
func (o *GetMeTaskEmailChangeIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me task email change ID params
func (o *GetMeTaskEmailChangeIDParams) WithHTTPClient(client *http.Client) *GetMeTaskEmailChangeIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me task email change ID params
func (o *GetMeTaskEmailChangeIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get me task email change ID params
func (o *GetMeTaskEmailChangeIDParams) WithID(id int64) *GetMeTaskEmailChangeIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get me task email change ID params
func (o *GetMeTaskEmailChangeIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeTaskEmailChangeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}