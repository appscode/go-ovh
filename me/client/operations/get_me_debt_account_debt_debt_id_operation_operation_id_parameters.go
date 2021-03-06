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

// NewGetMeDebtAccountDebtDebtIDOperationOperationIDParams creates a new GetMeDebtAccountDebtDebtIDOperationOperationIDParams object
// with the default values initialized.
func NewGetMeDebtAccountDebtDebtIDOperationOperationIDParams() *GetMeDebtAccountDebtDebtIDOperationOperationIDParams {
	var ()
	return &GetMeDebtAccountDebtDebtIDOperationOperationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeDebtAccountDebtDebtIDOperationOperationIDParamsWithTimeout creates a new GetMeDebtAccountDebtDebtIDOperationOperationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeDebtAccountDebtDebtIDOperationOperationIDParamsWithTimeout(timeout time.Duration) *GetMeDebtAccountDebtDebtIDOperationOperationIDParams {
	var ()
	return &GetMeDebtAccountDebtDebtIDOperationOperationIDParams{

		timeout: timeout,
	}
}

// NewGetMeDebtAccountDebtDebtIDOperationOperationIDParamsWithContext creates a new GetMeDebtAccountDebtDebtIDOperationOperationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeDebtAccountDebtDebtIDOperationOperationIDParamsWithContext(ctx context.Context) *GetMeDebtAccountDebtDebtIDOperationOperationIDParams {
	var ()
	return &GetMeDebtAccountDebtDebtIDOperationOperationIDParams{

		Context: ctx,
	}
}

// NewGetMeDebtAccountDebtDebtIDOperationOperationIDParamsWithHTTPClient creates a new GetMeDebtAccountDebtDebtIDOperationOperationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeDebtAccountDebtDebtIDOperationOperationIDParamsWithHTTPClient(client *http.Client) *GetMeDebtAccountDebtDebtIDOperationOperationIDParams {
	var ()
	return &GetMeDebtAccountDebtDebtIDOperationOperationIDParams{
		HTTPClient: client,
	}
}

/*GetMeDebtAccountDebtDebtIDOperationOperationIDParams contains all the parameters to send to the API endpoint
for the get me debt account debt debt ID operation operation ID operation typically these are written to a http.Request
*/
type GetMeDebtAccountDebtDebtIDOperationOperationIDParams struct {

	/*DebtID*/
	DebtID int64
	/*OperationID*/
	OperationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me debt account debt debt ID operation operation ID params
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDParams) WithTimeout(timeout time.Duration) *GetMeDebtAccountDebtDebtIDOperationOperationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me debt account debt debt ID operation operation ID params
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me debt account debt debt ID operation operation ID params
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDParams) WithContext(ctx context.Context) *GetMeDebtAccountDebtDebtIDOperationOperationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me debt account debt debt ID operation operation ID params
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me debt account debt debt ID operation operation ID params
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDParams) WithHTTPClient(client *http.Client) *GetMeDebtAccountDebtDebtIDOperationOperationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me debt account debt debt ID operation operation ID params
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDebtID adds the debtID to the get me debt account debt debt ID operation operation ID params
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDParams) WithDebtID(debtID int64) *GetMeDebtAccountDebtDebtIDOperationOperationIDParams {
	o.SetDebtID(debtID)
	return o
}

// SetDebtID adds the debtId to the get me debt account debt debt ID operation operation ID params
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDParams) SetDebtID(debtID int64) {
	o.DebtID = debtID
}

// WithOperationID adds the operationID to the get me debt account debt debt ID operation operation ID params
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDParams) WithOperationID(operationID int64) *GetMeDebtAccountDebtDebtIDOperationOperationIDParams {
	o.SetOperationID(operationID)
	return o
}

// SetOperationID adds the operationId to the get me debt account debt debt ID operation operation ID params
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDParams) SetOperationID(operationID int64) {
	o.OperationID = operationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param debtId
	if err := r.SetPathParam("debtId", swag.FormatInt64(o.DebtID)); err != nil {
		return err
	}

	// path param operationId
	if err := r.SetPathParam("operationId", swag.FormatInt64(o.OperationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
