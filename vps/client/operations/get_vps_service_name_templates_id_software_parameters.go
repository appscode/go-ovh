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

// NewGetVpsServiceNameTemplatesIDSoftwareParams creates a new GetVpsServiceNameTemplatesIDSoftwareParams object
// with the default values initialized.
func NewGetVpsServiceNameTemplatesIDSoftwareParams() *GetVpsServiceNameTemplatesIDSoftwareParams {
	var ()
	return &GetVpsServiceNameTemplatesIDSoftwareParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVpsServiceNameTemplatesIDSoftwareParamsWithTimeout creates a new GetVpsServiceNameTemplatesIDSoftwareParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVpsServiceNameTemplatesIDSoftwareParamsWithTimeout(timeout time.Duration) *GetVpsServiceNameTemplatesIDSoftwareParams {
	var ()
	return &GetVpsServiceNameTemplatesIDSoftwareParams{

		timeout: timeout,
	}
}

// NewGetVpsServiceNameTemplatesIDSoftwareParamsWithContext creates a new GetVpsServiceNameTemplatesIDSoftwareParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVpsServiceNameTemplatesIDSoftwareParamsWithContext(ctx context.Context) *GetVpsServiceNameTemplatesIDSoftwareParams {
	var ()
	return &GetVpsServiceNameTemplatesIDSoftwareParams{

		Context: ctx,
	}
}

// NewGetVpsServiceNameTemplatesIDSoftwareParamsWithHTTPClient creates a new GetVpsServiceNameTemplatesIDSoftwareParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVpsServiceNameTemplatesIDSoftwareParamsWithHTTPClient(client *http.Client) *GetVpsServiceNameTemplatesIDSoftwareParams {
	var ()
	return &GetVpsServiceNameTemplatesIDSoftwareParams{
		HTTPClient: client,
	}
}

/*GetVpsServiceNameTemplatesIDSoftwareParams contains all the parameters to send to the API endpoint
for the get vps service name templates ID software operation typically these are written to a http.Request
*/
type GetVpsServiceNameTemplatesIDSoftwareParams struct {

	/*ID*/
	ID int64
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vps service name templates ID software params
func (o *GetVpsServiceNameTemplatesIDSoftwareParams) WithTimeout(timeout time.Duration) *GetVpsServiceNameTemplatesIDSoftwareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vps service name templates ID software params
func (o *GetVpsServiceNameTemplatesIDSoftwareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vps service name templates ID software params
func (o *GetVpsServiceNameTemplatesIDSoftwareParams) WithContext(ctx context.Context) *GetVpsServiceNameTemplatesIDSoftwareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vps service name templates ID software params
func (o *GetVpsServiceNameTemplatesIDSoftwareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vps service name templates ID software params
func (o *GetVpsServiceNameTemplatesIDSoftwareParams) WithHTTPClient(client *http.Client) *GetVpsServiceNameTemplatesIDSoftwareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vps service name templates ID software params
func (o *GetVpsServiceNameTemplatesIDSoftwareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get vps service name templates ID software params
func (o *GetVpsServiceNameTemplatesIDSoftwareParams) WithID(id int64) *GetVpsServiceNameTemplatesIDSoftwareParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get vps service name templates ID software params
func (o *GetVpsServiceNameTemplatesIDSoftwareParams) SetID(id int64) {
	o.ID = id
}

// WithServiceName adds the serviceName to the get vps service name templates ID software params
func (o *GetVpsServiceNameTemplatesIDSoftwareParams) WithServiceName(serviceName string) *GetVpsServiceNameTemplatesIDSoftwareParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get vps service name templates ID software params
func (o *GetVpsServiceNameTemplatesIDSoftwareParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetVpsServiceNameTemplatesIDSoftwareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
