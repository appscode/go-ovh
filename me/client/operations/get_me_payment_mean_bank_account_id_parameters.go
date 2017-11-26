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

// NewGetMePaymentMeanBankAccountIDParams creates a new GetMePaymentMeanBankAccountIDParams object
// with the default values initialized.
func NewGetMePaymentMeanBankAccountIDParams() *GetMePaymentMeanBankAccountIDParams {
	var ()
	return &GetMePaymentMeanBankAccountIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMePaymentMeanBankAccountIDParamsWithTimeout creates a new GetMePaymentMeanBankAccountIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMePaymentMeanBankAccountIDParamsWithTimeout(timeout time.Duration) *GetMePaymentMeanBankAccountIDParams {
	var ()
	return &GetMePaymentMeanBankAccountIDParams{

		timeout: timeout,
	}
}

// NewGetMePaymentMeanBankAccountIDParamsWithContext creates a new GetMePaymentMeanBankAccountIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMePaymentMeanBankAccountIDParamsWithContext(ctx context.Context) *GetMePaymentMeanBankAccountIDParams {
	var ()
	return &GetMePaymentMeanBankAccountIDParams{

		Context: ctx,
	}
}

// NewGetMePaymentMeanBankAccountIDParamsWithHTTPClient creates a new GetMePaymentMeanBankAccountIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMePaymentMeanBankAccountIDParamsWithHTTPClient(client *http.Client) *GetMePaymentMeanBankAccountIDParams {
	var ()
	return &GetMePaymentMeanBankAccountIDParams{
		HTTPClient: client,
	}
}

/*GetMePaymentMeanBankAccountIDParams contains all the parameters to send to the API endpoint
for the get me payment mean bank account ID operation typically these are written to a http.Request
*/
type GetMePaymentMeanBankAccountIDParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me payment mean bank account ID params
func (o *GetMePaymentMeanBankAccountIDParams) WithTimeout(timeout time.Duration) *GetMePaymentMeanBankAccountIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me payment mean bank account ID params
func (o *GetMePaymentMeanBankAccountIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me payment mean bank account ID params
func (o *GetMePaymentMeanBankAccountIDParams) WithContext(ctx context.Context) *GetMePaymentMeanBankAccountIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me payment mean bank account ID params
func (o *GetMePaymentMeanBankAccountIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me payment mean bank account ID params
func (o *GetMePaymentMeanBankAccountIDParams) WithHTTPClient(client *http.Client) *GetMePaymentMeanBankAccountIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me payment mean bank account ID params
func (o *GetMePaymentMeanBankAccountIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get me payment mean bank account ID params
func (o *GetMePaymentMeanBankAccountIDParams) WithID(id int64) *GetMePaymentMeanBankAccountIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get me payment mean bank account ID params
func (o *GetMePaymentMeanBankAccountIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetMePaymentMeanBankAccountIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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