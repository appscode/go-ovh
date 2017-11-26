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

// NewPutMePaymentMeanDeferredPaymentAccountIDParams creates a new PutMePaymentMeanDeferredPaymentAccountIDParams object
// with the default values initialized.
func NewPutMePaymentMeanDeferredPaymentAccountIDParams() *PutMePaymentMeanDeferredPaymentAccountIDParams {
	var ()
	return &PutMePaymentMeanDeferredPaymentAccountIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutMePaymentMeanDeferredPaymentAccountIDParamsWithTimeout creates a new PutMePaymentMeanDeferredPaymentAccountIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutMePaymentMeanDeferredPaymentAccountIDParamsWithTimeout(timeout time.Duration) *PutMePaymentMeanDeferredPaymentAccountIDParams {
	var ()
	return &PutMePaymentMeanDeferredPaymentAccountIDParams{

		timeout: timeout,
	}
}

// NewPutMePaymentMeanDeferredPaymentAccountIDParamsWithContext creates a new PutMePaymentMeanDeferredPaymentAccountIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutMePaymentMeanDeferredPaymentAccountIDParamsWithContext(ctx context.Context) *PutMePaymentMeanDeferredPaymentAccountIDParams {
	var ()
	return &PutMePaymentMeanDeferredPaymentAccountIDParams{

		Context: ctx,
	}
}

// NewPutMePaymentMeanDeferredPaymentAccountIDParamsWithHTTPClient creates a new PutMePaymentMeanDeferredPaymentAccountIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutMePaymentMeanDeferredPaymentAccountIDParamsWithHTTPClient(client *http.Client) *PutMePaymentMeanDeferredPaymentAccountIDParams {
	var ()
	return &PutMePaymentMeanDeferredPaymentAccountIDParams{
		HTTPClient: client,
	}
}

/*PutMePaymentMeanDeferredPaymentAccountIDParams contains all the parameters to send to the API endpoint
for the put me payment mean deferred payment account ID operation typically these are written to a http.Request
*/
type PutMePaymentMeanDeferredPaymentAccountIDParams struct {

	/*ID*/
	ID int64
	/*MePaymentMeanDeferredPaymentAccountPut*/
	MePaymentMeanDeferredPaymentAccountPut *models.BillingDeferredPaymentAccount

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put me payment mean deferred payment account ID params
func (o *PutMePaymentMeanDeferredPaymentAccountIDParams) WithTimeout(timeout time.Duration) *PutMePaymentMeanDeferredPaymentAccountIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put me payment mean deferred payment account ID params
func (o *PutMePaymentMeanDeferredPaymentAccountIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put me payment mean deferred payment account ID params
func (o *PutMePaymentMeanDeferredPaymentAccountIDParams) WithContext(ctx context.Context) *PutMePaymentMeanDeferredPaymentAccountIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put me payment mean deferred payment account ID params
func (o *PutMePaymentMeanDeferredPaymentAccountIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put me payment mean deferred payment account ID params
func (o *PutMePaymentMeanDeferredPaymentAccountIDParams) WithHTTPClient(client *http.Client) *PutMePaymentMeanDeferredPaymentAccountIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put me payment mean deferred payment account ID params
func (o *PutMePaymentMeanDeferredPaymentAccountIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put me payment mean deferred payment account ID params
func (o *PutMePaymentMeanDeferredPaymentAccountIDParams) WithID(id int64) *PutMePaymentMeanDeferredPaymentAccountIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put me payment mean deferred payment account ID params
func (o *PutMePaymentMeanDeferredPaymentAccountIDParams) SetID(id int64) {
	o.ID = id
}

// WithMePaymentMeanDeferredPaymentAccountPut adds the mePaymentMeanDeferredPaymentAccountPut to the put me payment mean deferred payment account ID params
func (o *PutMePaymentMeanDeferredPaymentAccountIDParams) WithMePaymentMeanDeferredPaymentAccountPut(mePaymentMeanDeferredPaymentAccountPut *models.BillingDeferredPaymentAccount) *PutMePaymentMeanDeferredPaymentAccountIDParams {
	o.SetMePaymentMeanDeferredPaymentAccountPut(mePaymentMeanDeferredPaymentAccountPut)
	return o
}

// SetMePaymentMeanDeferredPaymentAccountPut adds the mePaymentMeanDeferredPaymentAccountPut to the put me payment mean deferred payment account ID params
func (o *PutMePaymentMeanDeferredPaymentAccountIDParams) SetMePaymentMeanDeferredPaymentAccountPut(mePaymentMeanDeferredPaymentAccountPut *models.BillingDeferredPaymentAccount) {
	o.MePaymentMeanDeferredPaymentAccountPut = mePaymentMeanDeferredPaymentAccountPut
}

// WriteToRequest writes these params to a swagger request
func (o *PutMePaymentMeanDeferredPaymentAccountIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.MePaymentMeanDeferredPaymentAccountPut != nil {
		if err := r.SetBodyParam(o.MePaymentMeanDeferredPaymentAccountPut); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
