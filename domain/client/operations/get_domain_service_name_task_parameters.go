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

// NewGetDomainServiceNameTaskParams creates a new GetDomainServiceNameTaskParams object
// with the default values initialized.
func NewGetDomainServiceNameTaskParams() *GetDomainServiceNameTaskParams {
	var ()
	return &GetDomainServiceNameTaskParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomainServiceNameTaskParamsWithTimeout creates a new GetDomainServiceNameTaskParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDomainServiceNameTaskParamsWithTimeout(timeout time.Duration) *GetDomainServiceNameTaskParams {
	var ()
	return &GetDomainServiceNameTaskParams{

		timeout: timeout,
	}
}

// NewGetDomainServiceNameTaskParamsWithContext creates a new GetDomainServiceNameTaskParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDomainServiceNameTaskParamsWithContext(ctx context.Context) *GetDomainServiceNameTaskParams {
	var ()
	return &GetDomainServiceNameTaskParams{

		Context: ctx,
	}
}

// NewGetDomainServiceNameTaskParamsWithHTTPClient creates a new GetDomainServiceNameTaskParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDomainServiceNameTaskParamsWithHTTPClient(client *http.Client) *GetDomainServiceNameTaskParams {
	var ()
	return &GetDomainServiceNameTaskParams{
		HTTPClient: client,
	}
}

/*GetDomainServiceNameTaskParams contains all the parameters to send to the API endpoint
for the get domain service name task operation typically these are written to a http.Request
*/
type GetDomainServiceNameTaskParams struct {

	/*Function*/
	Function *string
	/*ServiceName*/
	ServiceName string
	/*Status*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get domain service name task params
func (o *GetDomainServiceNameTaskParams) WithTimeout(timeout time.Duration) *GetDomainServiceNameTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domain service name task params
func (o *GetDomainServiceNameTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domain service name task params
func (o *GetDomainServiceNameTaskParams) WithContext(ctx context.Context) *GetDomainServiceNameTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domain service name task params
func (o *GetDomainServiceNameTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domain service name task params
func (o *GetDomainServiceNameTaskParams) WithHTTPClient(client *http.Client) *GetDomainServiceNameTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domain service name task params
func (o *GetDomainServiceNameTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFunction adds the function to the get domain service name task params
func (o *GetDomainServiceNameTaskParams) WithFunction(function *string) *GetDomainServiceNameTaskParams {
	o.SetFunction(function)
	return o
}

// SetFunction adds the function to the get domain service name task params
func (o *GetDomainServiceNameTaskParams) SetFunction(function *string) {
	o.Function = function
}

// WithServiceName adds the serviceName to the get domain service name task params
func (o *GetDomainServiceNameTaskParams) WithServiceName(serviceName string) *GetDomainServiceNameTaskParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get domain service name task params
func (o *GetDomainServiceNameTaskParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WithStatus adds the status to the get domain service name task params
func (o *GetDomainServiceNameTaskParams) WithStatus(status *string) *GetDomainServiceNameTaskParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get domain service name task params
func (o *GetDomainServiceNameTaskParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomainServiceNameTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Function != nil {

		// query param function
		var qrFunction string
		if o.Function != nil {
			qrFunction = *o.Function
		}
		qFunction := qrFunction
		if qFunction != "" {
			if err := r.SetQueryParam("function", qFunction); err != nil {
				return err
			}
		}

	}

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
