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

// NewGetMePaymentMeanCreditCardParams creates a new GetMePaymentMeanCreditCardParams object
// with the default values initialized.
func NewGetMePaymentMeanCreditCardParams() *GetMePaymentMeanCreditCardParams {

	return &GetMePaymentMeanCreditCardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMePaymentMeanCreditCardParamsWithTimeout creates a new GetMePaymentMeanCreditCardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMePaymentMeanCreditCardParamsWithTimeout(timeout time.Duration) *GetMePaymentMeanCreditCardParams {

	return &GetMePaymentMeanCreditCardParams{

		timeout: timeout,
	}
}

// NewGetMePaymentMeanCreditCardParamsWithContext creates a new GetMePaymentMeanCreditCardParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMePaymentMeanCreditCardParamsWithContext(ctx context.Context) *GetMePaymentMeanCreditCardParams {

	return &GetMePaymentMeanCreditCardParams{

		Context: ctx,
	}
}

// NewGetMePaymentMeanCreditCardParamsWithHTTPClient creates a new GetMePaymentMeanCreditCardParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMePaymentMeanCreditCardParamsWithHTTPClient(client *http.Client) *GetMePaymentMeanCreditCardParams {

	return &GetMePaymentMeanCreditCardParams{
		HTTPClient: client,
	}
}

/*GetMePaymentMeanCreditCardParams contains all the parameters to send to the API endpoint
for the get me payment mean credit card operation typically these are written to a http.Request
*/
type GetMePaymentMeanCreditCardParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me payment mean credit card params
func (o *GetMePaymentMeanCreditCardParams) WithTimeout(timeout time.Duration) *GetMePaymentMeanCreditCardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me payment mean credit card params
func (o *GetMePaymentMeanCreditCardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me payment mean credit card params
func (o *GetMePaymentMeanCreditCardParams) WithContext(ctx context.Context) *GetMePaymentMeanCreditCardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me payment mean credit card params
func (o *GetMePaymentMeanCreditCardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me payment mean credit card params
func (o *GetMePaymentMeanCreditCardParams) WithHTTPClient(client *http.Client) *GetMePaymentMeanCreditCardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me payment mean credit card params
func (o *GetMePaymentMeanCreditCardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetMePaymentMeanCreditCardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
