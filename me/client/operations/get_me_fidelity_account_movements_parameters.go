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

// NewGetMeFidelityAccountMovementsParams creates a new GetMeFidelityAccountMovementsParams object
// with the default values initialized.
func NewGetMeFidelityAccountMovementsParams() *GetMeFidelityAccountMovementsParams {
	var ()
	return &GetMeFidelityAccountMovementsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeFidelityAccountMovementsParamsWithTimeout creates a new GetMeFidelityAccountMovementsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeFidelityAccountMovementsParamsWithTimeout(timeout time.Duration) *GetMeFidelityAccountMovementsParams {
	var ()
	return &GetMeFidelityAccountMovementsParams{

		timeout: timeout,
	}
}

// NewGetMeFidelityAccountMovementsParamsWithContext creates a new GetMeFidelityAccountMovementsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeFidelityAccountMovementsParamsWithContext(ctx context.Context) *GetMeFidelityAccountMovementsParams {
	var ()
	return &GetMeFidelityAccountMovementsParams{

		Context: ctx,
	}
}

// NewGetMeFidelityAccountMovementsParamsWithHTTPClient creates a new GetMeFidelityAccountMovementsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeFidelityAccountMovementsParamsWithHTTPClient(client *http.Client) *GetMeFidelityAccountMovementsParams {
	var ()
	return &GetMeFidelityAccountMovementsParams{
		HTTPClient: client,
	}
}

/*GetMeFidelityAccountMovementsParams contains all the parameters to send to the API endpoint
for the get me fidelity account movements operation typically these are written to a http.Request
*/
type GetMeFidelityAccountMovementsParams struct {

	/*DateFrom*/
	DateFrom *strfmt.DateTime
	/*DateTo*/
	DateTo *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me fidelity account movements params
func (o *GetMeFidelityAccountMovementsParams) WithTimeout(timeout time.Duration) *GetMeFidelityAccountMovementsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me fidelity account movements params
func (o *GetMeFidelityAccountMovementsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me fidelity account movements params
func (o *GetMeFidelityAccountMovementsParams) WithContext(ctx context.Context) *GetMeFidelityAccountMovementsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me fidelity account movements params
func (o *GetMeFidelityAccountMovementsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me fidelity account movements params
func (o *GetMeFidelityAccountMovementsParams) WithHTTPClient(client *http.Client) *GetMeFidelityAccountMovementsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me fidelity account movements params
func (o *GetMeFidelityAccountMovementsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDateFrom adds the dateFrom to the get me fidelity account movements params
func (o *GetMeFidelityAccountMovementsParams) WithDateFrom(dateFrom *strfmt.DateTime) *GetMeFidelityAccountMovementsParams {
	o.SetDateFrom(dateFrom)
	return o
}

// SetDateFrom adds the dateFrom to the get me fidelity account movements params
func (o *GetMeFidelityAccountMovementsParams) SetDateFrom(dateFrom *strfmt.DateTime) {
	o.DateFrom = dateFrom
}

// WithDateTo adds the dateTo to the get me fidelity account movements params
func (o *GetMeFidelityAccountMovementsParams) WithDateTo(dateTo *strfmt.DateTime) *GetMeFidelityAccountMovementsParams {
	o.SetDateTo(dateTo)
	return o
}

// SetDateTo adds the dateTo to the get me fidelity account movements params
func (o *GetMeFidelityAccountMovementsParams) SetDateTo(dateTo *strfmt.DateTime) {
	o.DateTo = dateTo
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeFidelityAccountMovementsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DateFrom != nil {

		// query param date.from
		var qrDateFrom strfmt.DateTime
		if o.DateFrom != nil {
			qrDateFrom = *o.DateFrom
		}
		qDateFrom := qrDateFrom.String()
		if qDateFrom != "" {
			if err := r.SetQueryParam("date.from", qDateFrom); err != nil {
				return err
			}
		}

	}

	if o.DateTo != nil {

		// query param date.to
		var qrDateTo strfmt.DateTime
		if o.DateTo != nil {
			qrDateTo = *o.DateTo
		}
		qDateTo := qrDateTo.String()
		if qDateTo != "" {
			if err := r.SetQueryParam("date.to", qDateTo); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}