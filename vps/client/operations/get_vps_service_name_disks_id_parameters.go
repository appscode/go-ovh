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

// NewGetVpsServiceNameDisksIDParams creates a new GetVpsServiceNameDisksIDParams object
// with the default values initialized.
func NewGetVpsServiceNameDisksIDParams() *GetVpsServiceNameDisksIDParams {
	var ()
	return &GetVpsServiceNameDisksIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVpsServiceNameDisksIDParamsWithTimeout creates a new GetVpsServiceNameDisksIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVpsServiceNameDisksIDParamsWithTimeout(timeout time.Duration) *GetVpsServiceNameDisksIDParams {
	var ()
	return &GetVpsServiceNameDisksIDParams{

		timeout: timeout,
	}
}

// NewGetVpsServiceNameDisksIDParamsWithContext creates a new GetVpsServiceNameDisksIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVpsServiceNameDisksIDParamsWithContext(ctx context.Context) *GetVpsServiceNameDisksIDParams {
	var ()
	return &GetVpsServiceNameDisksIDParams{

		Context: ctx,
	}
}

// NewGetVpsServiceNameDisksIDParamsWithHTTPClient creates a new GetVpsServiceNameDisksIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVpsServiceNameDisksIDParamsWithHTTPClient(client *http.Client) *GetVpsServiceNameDisksIDParams {
	var ()
	return &GetVpsServiceNameDisksIDParams{
		HTTPClient: client,
	}
}

/*GetVpsServiceNameDisksIDParams contains all the parameters to send to the API endpoint
for the get vps service name disks ID operation typically these are written to a http.Request
*/
type GetVpsServiceNameDisksIDParams struct {

	/*ID*/
	ID int64
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vps service name disks ID params
func (o *GetVpsServiceNameDisksIDParams) WithTimeout(timeout time.Duration) *GetVpsServiceNameDisksIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vps service name disks ID params
func (o *GetVpsServiceNameDisksIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vps service name disks ID params
func (o *GetVpsServiceNameDisksIDParams) WithContext(ctx context.Context) *GetVpsServiceNameDisksIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vps service name disks ID params
func (o *GetVpsServiceNameDisksIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vps service name disks ID params
func (o *GetVpsServiceNameDisksIDParams) WithHTTPClient(client *http.Client) *GetVpsServiceNameDisksIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vps service name disks ID params
func (o *GetVpsServiceNameDisksIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get vps service name disks ID params
func (o *GetVpsServiceNameDisksIDParams) WithID(id int64) *GetVpsServiceNameDisksIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get vps service name disks ID params
func (o *GetVpsServiceNameDisksIDParams) SetID(id int64) {
	o.ID = id
}

// WithServiceName adds the serviceName to the get vps service name disks ID params
func (o *GetVpsServiceNameDisksIDParams) WithServiceName(serviceName string) *GetVpsServiceNameDisksIDParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get vps service name disks ID params
func (o *GetVpsServiceNameDisksIDParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetVpsServiceNameDisksIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
