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

	"github.com/appscode/go-ovh/me/models"
)

// NewPutMeSubAccountIDParams creates a new PutMeSubAccountIDParams object
// with the default values initialized.
func NewPutMeSubAccountIDParams() *PutMeSubAccountIDParams {
	var ()
	return &PutMeSubAccountIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutMeSubAccountIDParamsWithTimeout creates a new PutMeSubAccountIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutMeSubAccountIDParamsWithTimeout(timeout time.Duration) *PutMeSubAccountIDParams {
	var ()
	return &PutMeSubAccountIDParams{

		timeout: timeout,
	}
}

// NewPutMeSubAccountIDParamsWithContext creates a new PutMeSubAccountIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutMeSubAccountIDParamsWithContext(ctx context.Context) *PutMeSubAccountIDParams {
	var ()
	return &PutMeSubAccountIDParams{

		Context: ctx,
	}
}

// NewPutMeSubAccountIDParamsWithHTTPClient creates a new PutMeSubAccountIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutMeSubAccountIDParamsWithHTTPClient(client *http.Client) *PutMeSubAccountIDParams {
	var ()
	return &PutMeSubAccountIDParams{
		HTTPClient: client,
	}
}

/*PutMeSubAccountIDParams contains all the parameters to send to the API endpoint
for the put me sub account ID operation typically these are written to a http.Request
*/
type PutMeSubAccountIDParams struct {

	/*ID*/
	ID int64
	/*MeSubAccountPut*/
	MeSubAccountPut *models.NichandleSubAccount

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put me sub account ID params
func (o *PutMeSubAccountIDParams) WithTimeout(timeout time.Duration) *PutMeSubAccountIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put me sub account ID params
func (o *PutMeSubAccountIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put me sub account ID params
func (o *PutMeSubAccountIDParams) WithContext(ctx context.Context) *PutMeSubAccountIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put me sub account ID params
func (o *PutMeSubAccountIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put me sub account ID params
func (o *PutMeSubAccountIDParams) WithHTTPClient(client *http.Client) *PutMeSubAccountIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put me sub account ID params
func (o *PutMeSubAccountIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put me sub account ID params
func (o *PutMeSubAccountIDParams) WithID(id int64) *PutMeSubAccountIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put me sub account ID params
func (o *PutMeSubAccountIDParams) SetID(id int64) {
	o.ID = id
}

// WithMeSubAccountPut adds the meSubAccountPut to the put me sub account ID params
func (o *PutMeSubAccountIDParams) WithMeSubAccountPut(meSubAccountPut *models.NichandleSubAccount) *PutMeSubAccountIDParams {
	o.SetMeSubAccountPut(meSubAccountPut)
	return o
}

// SetMeSubAccountPut adds the meSubAccountPut to the put me sub account ID params
func (o *PutMeSubAccountIDParams) SetMeSubAccountPut(meSubAccountPut *models.NichandleSubAccount) {
	o.MeSubAccountPut = meSubAccountPut
}

// WriteToRequest writes these params to a swagger request
func (o *PutMeSubAccountIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.MeSubAccountPut != nil {
		if err := r.SetBodyParam(o.MeSubAccountPut); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
