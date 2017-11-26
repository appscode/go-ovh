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

// NewGetMeAgreementsIDParams creates a new GetMeAgreementsIDParams object
// with the default values initialized.
func NewGetMeAgreementsIDParams() *GetMeAgreementsIDParams {
	var ()
	return &GetMeAgreementsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeAgreementsIDParamsWithTimeout creates a new GetMeAgreementsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeAgreementsIDParamsWithTimeout(timeout time.Duration) *GetMeAgreementsIDParams {
	var ()
	return &GetMeAgreementsIDParams{

		timeout: timeout,
	}
}

// NewGetMeAgreementsIDParamsWithContext creates a new GetMeAgreementsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeAgreementsIDParamsWithContext(ctx context.Context) *GetMeAgreementsIDParams {
	var ()
	return &GetMeAgreementsIDParams{

		Context: ctx,
	}
}

// NewGetMeAgreementsIDParamsWithHTTPClient creates a new GetMeAgreementsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeAgreementsIDParamsWithHTTPClient(client *http.Client) *GetMeAgreementsIDParams {
	var ()
	return &GetMeAgreementsIDParams{
		HTTPClient: client,
	}
}

/*GetMeAgreementsIDParams contains all the parameters to send to the API endpoint
for the get me agreements ID operation typically these are written to a http.Request
*/
type GetMeAgreementsIDParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me agreements ID params
func (o *GetMeAgreementsIDParams) WithTimeout(timeout time.Duration) *GetMeAgreementsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me agreements ID params
func (o *GetMeAgreementsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me agreements ID params
func (o *GetMeAgreementsIDParams) WithContext(ctx context.Context) *GetMeAgreementsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me agreements ID params
func (o *GetMeAgreementsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me agreements ID params
func (o *GetMeAgreementsIDParams) WithHTTPClient(client *http.Client) *GetMeAgreementsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me agreements ID params
func (o *GetMeAgreementsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get me agreements ID params
func (o *GetMeAgreementsIDParams) WithID(id int64) *GetMeAgreementsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get me agreements ID params
func (o *GetMeAgreementsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeAgreementsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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