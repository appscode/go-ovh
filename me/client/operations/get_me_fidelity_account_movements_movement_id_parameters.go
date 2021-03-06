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

// NewGetMeFidelityAccountMovementsMovementIDParams creates a new GetMeFidelityAccountMovementsMovementIDParams object
// with the default values initialized.
func NewGetMeFidelityAccountMovementsMovementIDParams() *GetMeFidelityAccountMovementsMovementIDParams {
	var ()
	return &GetMeFidelityAccountMovementsMovementIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeFidelityAccountMovementsMovementIDParamsWithTimeout creates a new GetMeFidelityAccountMovementsMovementIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeFidelityAccountMovementsMovementIDParamsWithTimeout(timeout time.Duration) *GetMeFidelityAccountMovementsMovementIDParams {
	var ()
	return &GetMeFidelityAccountMovementsMovementIDParams{

		timeout: timeout,
	}
}

// NewGetMeFidelityAccountMovementsMovementIDParamsWithContext creates a new GetMeFidelityAccountMovementsMovementIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeFidelityAccountMovementsMovementIDParamsWithContext(ctx context.Context) *GetMeFidelityAccountMovementsMovementIDParams {
	var ()
	return &GetMeFidelityAccountMovementsMovementIDParams{

		Context: ctx,
	}
}

// NewGetMeFidelityAccountMovementsMovementIDParamsWithHTTPClient creates a new GetMeFidelityAccountMovementsMovementIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeFidelityAccountMovementsMovementIDParamsWithHTTPClient(client *http.Client) *GetMeFidelityAccountMovementsMovementIDParams {
	var ()
	return &GetMeFidelityAccountMovementsMovementIDParams{
		HTTPClient: client,
	}
}

/*GetMeFidelityAccountMovementsMovementIDParams contains all the parameters to send to the API endpoint
for the get me fidelity account movements movement ID operation typically these are written to a http.Request
*/
type GetMeFidelityAccountMovementsMovementIDParams struct {

	/*MovementID*/
	MovementID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me fidelity account movements movement ID params
func (o *GetMeFidelityAccountMovementsMovementIDParams) WithTimeout(timeout time.Duration) *GetMeFidelityAccountMovementsMovementIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me fidelity account movements movement ID params
func (o *GetMeFidelityAccountMovementsMovementIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me fidelity account movements movement ID params
func (o *GetMeFidelityAccountMovementsMovementIDParams) WithContext(ctx context.Context) *GetMeFidelityAccountMovementsMovementIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me fidelity account movements movement ID params
func (o *GetMeFidelityAccountMovementsMovementIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me fidelity account movements movement ID params
func (o *GetMeFidelityAccountMovementsMovementIDParams) WithHTTPClient(client *http.Client) *GetMeFidelityAccountMovementsMovementIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me fidelity account movements movement ID params
func (o *GetMeFidelityAccountMovementsMovementIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMovementID adds the movementID to the get me fidelity account movements movement ID params
func (o *GetMeFidelityAccountMovementsMovementIDParams) WithMovementID(movementID int64) *GetMeFidelityAccountMovementsMovementIDParams {
	o.SetMovementID(movementID)
	return o
}

// SetMovementID adds the movementId to the get me fidelity account movements movement ID params
func (o *GetMeFidelityAccountMovementsMovementIDParams) SetMovementID(movementID int64) {
	o.MovementID = movementID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeFidelityAccountMovementsMovementIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param movementId
	if err := r.SetPathParam("movementId", swag.FormatInt64(o.MovementID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
