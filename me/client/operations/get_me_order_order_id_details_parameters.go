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

// NewGetMeOrderOrderIDDetailsParams creates a new GetMeOrderOrderIDDetailsParams object
// with the default values initialized.
func NewGetMeOrderOrderIDDetailsParams() *GetMeOrderOrderIDDetailsParams {
	var ()
	return &GetMeOrderOrderIDDetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeOrderOrderIDDetailsParamsWithTimeout creates a new GetMeOrderOrderIDDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeOrderOrderIDDetailsParamsWithTimeout(timeout time.Duration) *GetMeOrderOrderIDDetailsParams {
	var ()
	return &GetMeOrderOrderIDDetailsParams{

		timeout: timeout,
	}
}

// NewGetMeOrderOrderIDDetailsParamsWithContext creates a new GetMeOrderOrderIDDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeOrderOrderIDDetailsParamsWithContext(ctx context.Context) *GetMeOrderOrderIDDetailsParams {
	var ()
	return &GetMeOrderOrderIDDetailsParams{

		Context: ctx,
	}
}

// NewGetMeOrderOrderIDDetailsParamsWithHTTPClient creates a new GetMeOrderOrderIDDetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeOrderOrderIDDetailsParamsWithHTTPClient(client *http.Client) *GetMeOrderOrderIDDetailsParams {
	var ()
	return &GetMeOrderOrderIDDetailsParams{
		HTTPClient: client,
	}
}

/*GetMeOrderOrderIDDetailsParams contains all the parameters to send to the API endpoint
for the get me order order ID details operation typically these are written to a http.Request
*/
type GetMeOrderOrderIDDetailsParams struct {

	/*OrderID*/
	OrderID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me order order ID details params
func (o *GetMeOrderOrderIDDetailsParams) WithTimeout(timeout time.Duration) *GetMeOrderOrderIDDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me order order ID details params
func (o *GetMeOrderOrderIDDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me order order ID details params
func (o *GetMeOrderOrderIDDetailsParams) WithContext(ctx context.Context) *GetMeOrderOrderIDDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me order order ID details params
func (o *GetMeOrderOrderIDDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me order order ID details params
func (o *GetMeOrderOrderIDDetailsParams) WithHTTPClient(client *http.Client) *GetMeOrderOrderIDDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me order order ID details params
func (o *GetMeOrderOrderIDDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderID adds the orderID to the get me order order ID details params
func (o *GetMeOrderOrderIDDetailsParams) WithOrderID(orderID int64) *GetMeOrderOrderIDDetailsParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the get me order order ID details params
func (o *GetMeOrderOrderIDDetailsParams) SetOrderID(orderID int64) {
	o.OrderID = orderID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeOrderOrderIDDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orderId
	if err := r.SetPathParam("orderId", swag.FormatInt64(o.OrderID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
