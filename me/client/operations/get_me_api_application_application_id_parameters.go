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

// NewGetMeAPIApplicationApplicationIDParams creates a new GetMeAPIApplicationApplicationIDParams object
// with the default values initialized.
func NewGetMeAPIApplicationApplicationIDParams() *GetMeAPIApplicationApplicationIDParams {
	var ()
	return &GetMeAPIApplicationApplicationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeAPIApplicationApplicationIDParamsWithTimeout creates a new GetMeAPIApplicationApplicationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeAPIApplicationApplicationIDParamsWithTimeout(timeout time.Duration) *GetMeAPIApplicationApplicationIDParams {
	var ()
	return &GetMeAPIApplicationApplicationIDParams{

		timeout: timeout,
	}
}

// NewGetMeAPIApplicationApplicationIDParamsWithContext creates a new GetMeAPIApplicationApplicationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeAPIApplicationApplicationIDParamsWithContext(ctx context.Context) *GetMeAPIApplicationApplicationIDParams {
	var ()
	return &GetMeAPIApplicationApplicationIDParams{

		Context: ctx,
	}
}

// NewGetMeAPIApplicationApplicationIDParamsWithHTTPClient creates a new GetMeAPIApplicationApplicationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeAPIApplicationApplicationIDParamsWithHTTPClient(client *http.Client) *GetMeAPIApplicationApplicationIDParams {
	var ()
	return &GetMeAPIApplicationApplicationIDParams{
		HTTPClient: client,
	}
}

/*GetMeAPIApplicationApplicationIDParams contains all the parameters to send to the API endpoint
for the get me API application application ID operation typically these are written to a http.Request
*/
type GetMeAPIApplicationApplicationIDParams struct {

	/*ApplicationID*/
	ApplicationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me API application application ID params
func (o *GetMeAPIApplicationApplicationIDParams) WithTimeout(timeout time.Duration) *GetMeAPIApplicationApplicationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me API application application ID params
func (o *GetMeAPIApplicationApplicationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me API application application ID params
func (o *GetMeAPIApplicationApplicationIDParams) WithContext(ctx context.Context) *GetMeAPIApplicationApplicationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me API application application ID params
func (o *GetMeAPIApplicationApplicationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me API application application ID params
func (o *GetMeAPIApplicationApplicationIDParams) WithHTTPClient(client *http.Client) *GetMeAPIApplicationApplicationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me API application application ID params
func (o *GetMeAPIApplicationApplicationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationID adds the applicationID to the get me API application application ID params
func (o *GetMeAPIApplicationApplicationIDParams) WithApplicationID(applicationID int64) *GetMeAPIApplicationApplicationIDParams {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the get me API application application ID params
func (o *GetMeAPIApplicationApplicationIDParams) SetApplicationID(applicationID int64) {
	o.ApplicationID = applicationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeAPIApplicationApplicationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param applicationId
	if err := r.SetPathParam("applicationId", swag.FormatInt64(o.ApplicationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}