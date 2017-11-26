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

// NewGetMeOrderOrderIDDetailsOrderDetailIDExtensionParams creates a new GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams object
// with the default values initialized.
func NewGetMeOrderOrderIDDetailsOrderDetailIDExtensionParams() *GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams {
	var ()
	return &GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeOrderOrderIDDetailsOrderDetailIDExtensionParamsWithTimeout creates a new GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeOrderOrderIDDetailsOrderDetailIDExtensionParamsWithTimeout(timeout time.Duration) *GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams {
	var ()
	return &GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams{

		timeout: timeout,
	}
}

// NewGetMeOrderOrderIDDetailsOrderDetailIDExtensionParamsWithContext creates a new GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeOrderOrderIDDetailsOrderDetailIDExtensionParamsWithContext(ctx context.Context) *GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams {
	var ()
	return &GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams{

		Context: ctx,
	}
}

// NewGetMeOrderOrderIDDetailsOrderDetailIDExtensionParamsWithHTTPClient creates a new GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeOrderOrderIDDetailsOrderDetailIDExtensionParamsWithHTTPClient(client *http.Client) *GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams {
	var ()
	return &GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams{
		HTTPClient: client,
	}
}

/*GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams contains all the parameters to send to the API endpoint
for the get me order order ID details order detail ID extension operation typically these are written to a http.Request
*/
type GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams struct {

	/*OrderDetailID*/
	OrderDetailID int64
	/*OrderID*/
	OrderID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me order order ID details order detail ID extension params
func (o *GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams) WithTimeout(timeout time.Duration) *GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me order order ID details order detail ID extension params
func (o *GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me order order ID details order detail ID extension params
func (o *GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams) WithContext(ctx context.Context) *GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me order order ID details order detail ID extension params
func (o *GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me order order ID details order detail ID extension params
func (o *GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams) WithHTTPClient(client *http.Client) *GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me order order ID details order detail ID extension params
func (o *GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderDetailID adds the orderDetailID to the get me order order ID details order detail ID extension params
func (o *GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams) WithOrderDetailID(orderDetailID int64) *GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams {
	o.SetOrderDetailID(orderDetailID)
	return o
}

// SetOrderDetailID adds the orderDetailId to the get me order order ID details order detail ID extension params
func (o *GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams) SetOrderDetailID(orderDetailID int64) {
	o.OrderDetailID = orderDetailID
}

// WithOrderID adds the orderID to the get me order order ID details order detail ID extension params
func (o *GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams) WithOrderID(orderID int64) *GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the get me order order ID details order detail ID extension params
func (o *GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams) SetOrderID(orderID int64) {
	o.OrderID = orderID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeOrderOrderIDDetailsOrderDetailIDExtensionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orderDetailId
	if err := r.SetPathParam("orderDetailId", swag.FormatInt64(o.OrderDetailID)); err != nil {
		return err
	}

	// path param orderId
	if err := r.SetPathParam("orderId", swag.FormatInt64(o.OrderID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}