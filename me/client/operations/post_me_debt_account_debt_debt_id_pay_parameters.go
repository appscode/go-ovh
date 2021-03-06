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

// NewPostMeDebtAccountDebtDebtIDPayParams creates a new PostMeDebtAccountDebtDebtIDPayParams object
// with the default values initialized.
func NewPostMeDebtAccountDebtDebtIDPayParams() *PostMeDebtAccountDebtDebtIDPayParams {
	var ()
	return &PostMeDebtAccountDebtDebtIDPayParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeDebtAccountDebtDebtIDPayParamsWithTimeout creates a new PostMeDebtAccountDebtDebtIDPayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeDebtAccountDebtDebtIDPayParamsWithTimeout(timeout time.Duration) *PostMeDebtAccountDebtDebtIDPayParams {
	var ()
	return &PostMeDebtAccountDebtDebtIDPayParams{

		timeout: timeout,
	}
}

// NewPostMeDebtAccountDebtDebtIDPayParamsWithContext creates a new PostMeDebtAccountDebtDebtIDPayParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeDebtAccountDebtDebtIDPayParamsWithContext(ctx context.Context) *PostMeDebtAccountDebtDebtIDPayParams {
	var ()
	return &PostMeDebtAccountDebtDebtIDPayParams{

		Context: ctx,
	}
}

// NewPostMeDebtAccountDebtDebtIDPayParamsWithHTTPClient creates a new PostMeDebtAccountDebtDebtIDPayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeDebtAccountDebtDebtIDPayParamsWithHTTPClient(client *http.Client) *PostMeDebtAccountDebtDebtIDPayParams {
	var ()
	return &PostMeDebtAccountDebtDebtIDPayParams{
		HTTPClient: client,
	}
}

/*PostMeDebtAccountDebtDebtIDPayParams contains all the parameters to send to the API endpoint
for the post me debt account debt debt ID pay operation typically these are written to a http.Request
*/
type PostMeDebtAccountDebtDebtIDPayParams struct {

	/*DebtID*/
	DebtID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me debt account debt debt ID pay params
func (o *PostMeDebtAccountDebtDebtIDPayParams) WithTimeout(timeout time.Duration) *PostMeDebtAccountDebtDebtIDPayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me debt account debt debt ID pay params
func (o *PostMeDebtAccountDebtDebtIDPayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me debt account debt debt ID pay params
func (o *PostMeDebtAccountDebtDebtIDPayParams) WithContext(ctx context.Context) *PostMeDebtAccountDebtDebtIDPayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me debt account debt debt ID pay params
func (o *PostMeDebtAccountDebtDebtIDPayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me debt account debt debt ID pay params
func (o *PostMeDebtAccountDebtDebtIDPayParams) WithHTTPClient(client *http.Client) *PostMeDebtAccountDebtDebtIDPayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me debt account debt debt ID pay params
func (o *PostMeDebtAccountDebtDebtIDPayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDebtID adds the debtID to the post me debt account debt debt ID pay params
func (o *PostMeDebtAccountDebtDebtIDPayParams) WithDebtID(debtID int64) *PostMeDebtAccountDebtDebtIDPayParams {
	o.SetDebtID(debtID)
	return o
}

// SetDebtID adds the debtId to the post me debt account debt debt ID pay params
func (o *PostMeDebtAccountDebtDebtIDPayParams) SetDebtID(debtID int64) {
	o.DebtID = debtID
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeDebtAccountDebtDebtIDPayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param debtId
	if err := r.SetPathParam("debtId", swag.FormatInt64(o.DebtID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
