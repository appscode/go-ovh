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

// NewPostMeOrderOrderIDRetractionParams creates a new PostMeOrderOrderIDRetractionParams object
// with the default values initialized.
func NewPostMeOrderOrderIDRetractionParams() *PostMeOrderOrderIDRetractionParams {
	var ()
	return &PostMeOrderOrderIDRetractionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeOrderOrderIDRetractionParamsWithTimeout creates a new PostMeOrderOrderIDRetractionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeOrderOrderIDRetractionParamsWithTimeout(timeout time.Duration) *PostMeOrderOrderIDRetractionParams {
	var ()
	return &PostMeOrderOrderIDRetractionParams{

		timeout: timeout,
	}
}

// NewPostMeOrderOrderIDRetractionParamsWithContext creates a new PostMeOrderOrderIDRetractionParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeOrderOrderIDRetractionParamsWithContext(ctx context.Context) *PostMeOrderOrderIDRetractionParams {
	var ()
	return &PostMeOrderOrderIDRetractionParams{

		Context: ctx,
	}
}

// NewPostMeOrderOrderIDRetractionParamsWithHTTPClient creates a new PostMeOrderOrderIDRetractionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeOrderOrderIDRetractionParamsWithHTTPClient(client *http.Client) *PostMeOrderOrderIDRetractionParams {
	var ()
	return &PostMeOrderOrderIDRetractionParams{
		HTTPClient: client,
	}
}

/*PostMeOrderOrderIDRetractionParams contains all the parameters to send to the API endpoint
for the post me order order ID retraction operation typically these are written to a http.Request
*/
type PostMeOrderOrderIDRetractionParams struct {

	/*MeOrderRetractionPost*/
	MeOrderRetractionPost *models.PostMeOrderOrderIDRetractionParamsBody
	/*OrderID*/
	OrderID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me order order ID retraction params
func (o *PostMeOrderOrderIDRetractionParams) WithTimeout(timeout time.Duration) *PostMeOrderOrderIDRetractionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me order order ID retraction params
func (o *PostMeOrderOrderIDRetractionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me order order ID retraction params
func (o *PostMeOrderOrderIDRetractionParams) WithContext(ctx context.Context) *PostMeOrderOrderIDRetractionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me order order ID retraction params
func (o *PostMeOrderOrderIDRetractionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me order order ID retraction params
func (o *PostMeOrderOrderIDRetractionParams) WithHTTPClient(client *http.Client) *PostMeOrderOrderIDRetractionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me order order ID retraction params
func (o *PostMeOrderOrderIDRetractionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMeOrderRetractionPost adds the meOrderRetractionPost to the post me order order ID retraction params
func (o *PostMeOrderOrderIDRetractionParams) WithMeOrderRetractionPost(meOrderRetractionPost *models.PostMeOrderOrderIDRetractionParamsBody) *PostMeOrderOrderIDRetractionParams {
	o.SetMeOrderRetractionPost(meOrderRetractionPost)
	return o
}

// SetMeOrderRetractionPost adds the meOrderRetractionPost to the post me order order ID retraction params
func (o *PostMeOrderOrderIDRetractionParams) SetMeOrderRetractionPost(meOrderRetractionPost *models.PostMeOrderOrderIDRetractionParamsBody) {
	o.MeOrderRetractionPost = meOrderRetractionPost
}

// WithOrderID adds the orderID to the post me order order ID retraction params
func (o *PostMeOrderOrderIDRetractionParams) WithOrderID(orderID int64) *PostMeOrderOrderIDRetractionParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the post me order order ID retraction params
func (o *PostMeOrderOrderIDRetractionParams) SetOrderID(orderID int64) {
	o.OrderID = orderID
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeOrderOrderIDRetractionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MeOrderRetractionPost != nil {
		if err := r.SetBodyParam(o.MeOrderRetractionPost); err != nil {
			return err
		}
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
