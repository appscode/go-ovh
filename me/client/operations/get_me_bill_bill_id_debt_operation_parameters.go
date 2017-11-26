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

// NewGetMeBillBillIDDebtOperationParams creates a new GetMeBillBillIDDebtOperationParams object
// with the default values initialized.
func NewGetMeBillBillIDDebtOperationParams() *GetMeBillBillIDDebtOperationParams {
	var ()
	return &GetMeBillBillIDDebtOperationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeBillBillIDDebtOperationParamsWithTimeout creates a new GetMeBillBillIDDebtOperationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeBillBillIDDebtOperationParamsWithTimeout(timeout time.Duration) *GetMeBillBillIDDebtOperationParams {
	var ()
	return &GetMeBillBillIDDebtOperationParams{

		timeout: timeout,
	}
}

// NewGetMeBillBillIDDebtOperationParamsWithContext creates a new GetMeBillBillIDDebtOperationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeBillBillIDDebtOperationParamsWithContext(ctx context.Context) *GetMeBillBillIDDebtOperationParams {
	var ()
	return &GetMeBillBillIDDebtOperationParams{

		Context: ctx,
	}
}

// NewGetMeBillBillIDDebtOperationParamsWithHTTPClient creates a new GetMeBillBillIDDebtOperationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeBillBillIDDebtOperationParamsWithHTTPClient(client *http.Client) *GetMeBillBillIDDebtOperationParams {
	var ()
	return &GetMeBillBillIDDebtOperationParams{
		HTTPClient: client,
	}
}

/*GetMeBillBillIDDebtOperationParams contains all the parameters to send to the API endpoint
for the get me bill bill ID debt operation operation typically these are written to a http.Request
*/
type GetMeBillBillIDDebtOperationParams struct {

	/*BillID*/
	BillID string
	/*DepositOrderID*/
	DepositOrderID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me bill bill ID debt operation params
func (o *GetMeBillBillIDDebtOperationParams) WithTimeout(timeout time.Duration) *GetMeBillBillIDDebtOperationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me bill bill ID debt operation params
func (o *GetMeBillBillIDDebtOperationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me bill bill ID debt operation params
func (o *GetMeBillBillIDDebtOperationParams) WithContext(ctx context.Context) *GetMeBillBillIDDebtOperationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me bill bill ID debt operation params
func (o *GetMeBillBillIDDebtOperationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me bill bill ID debt operation params
func (o *GetMeBillBillIDDebtOperationParams) WithHTTPClient(client *http.Client) *GetMeBillBillIDDebtOperationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me bill bill ID debt operation params
func (o *GetMeBillBillIDDebtOperationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBillID adds the billID to the get me bill bill ID debt operation params
func (o *GetMeBillBillIDDebtOperationParams) WithBillID(billID string) *GetMeBillBillIDDebtOperationParams {
	o.SetBillID(billID)
	return o
}

// SetBillID adds the billId to the get me bill bill ID debt operation params
func (o *GetMeBillBillIDDebtOperationParams) SetBillID(billID string) {
	o.BillID = billID
}

// WithDepositOrderID adds the depositOrderID to the get me bill bill ID debt operation params
func (o *GetMeBillBillIDDebtOperationParams) WithDepositOrderID(depositOrderID *int64) *GetMeBillBillIDDebtOperationParams {
	o.SetDepositOrderID(depositOrderID)
	return o
}

// SetDepositOrderID adds the depositOrderId to the get me bill bill ID debt operation params
func (o *GetMeBillBillIDDebtOperationParams) SetDepositOrderID(depositOrderID *int64) {
	o.DepositOrderID = depositOrderID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeBillBillIDDebtOperationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param billId
	if err := r.SetPathParam("billId", o.BillID); err != nil {
		return err
	}

	if o.DepositOrderID != nil {

		// query param depositOrderId
		var qrDepositOrderID int64
		if o.DepositOrderID != nil {
			qrDepositOrderID = *o.DepositOrderID
		}
		qDepositOrderID := swag.FormatInt64(qrDepositOrderID)
		if qDepositOrderID != "" {
			if err := r.SetQueryParam("depositOrderId", qDepositOrderID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
