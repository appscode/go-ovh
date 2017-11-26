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

// NewGetCloudProjectServiceNameSnapshotSnapshotIDParams creates a new GetCloudProjectServiceNameSnapshotSnapshotIDParams object
// with the default values initialized.
func NewGetCloudProjectServiceNameSnapshotSnapshotIDParams() *GetCloudProjectServiceNameSnapshotSnapshotIDParams {
	var ()
	return &GetCloudProjectServiceNameSnapshotSnapshotIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCloudProjectServiceNameSnapshotSnapshotIDParamsWithTimeout creates a new GetCloudProjectServiceNameSnapshotSnapshotIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCloudProjectServiceNameSnapshotSnapshotIDParamsWithTimeout(timeout time.Duration) *GetCloudProjectServiceNameSnapshotSnapshotIDParams {
	var ()
	return &GetCloudProjectServiceNameSnapshotSnapshotIDParams{

		timeout: timeout,
	}
}

// NewGetCloudProjectServiceNameSnapshotSnapshotIDParamsWithContext creates a new GetCloudProjectServiceNameSnapshotSnapshotIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCloudProjectServiceNameSnapshotSnapshotIDParamsWithContext(ctx context.Context) *GetCloudProjectServiceNameSnapshotSnapshotIDParams {
	var ()
	return &GetCloudProjectServiceNameSnapshotSnapshotIDParams{

		Context: ctx,
	}
}

// NewGetCloudProjectServiceNameSnapshotSnapshotIDParamsWithHTTPClient creates a new GetCloudProjectServiceNameSnapshotSnapshotIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCloudProjectServiceNameSnapshotSnapshotIDParamsWithHTTPClient(client *http.Client) *GetCloudProjectServiceNameSnapshotSnapshotIDParams {
	var ()
	return &GetCloudProjectServiceNameSnapshotSnapshotIDParams{
		HTTPClient: client,
	}
}

/*GetCloudProjectServiceNameSnapshotSnapshotIDParams contains all the parameters to send to the API endpoint
for the get cloud project service name snapshot snapshot ID operation typically these are written to a http.Request
*/
type GetCloudProjectServiceNameSnapshotSnapshotIDParams struct {

	/*ServiceName*/
	ServiceName string
	/*SnapshotID*/
	SnapshotID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cloud project service name snapshot snapshot ID params
func (o *GetCloudProjectServiceNameSnapshotSnapshotIDParams) WithTimeout(timeout time.Duration) *GetCloudProjectServiceNameSnapshotSnapshotIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cloud project service name snapshot snapshot ID params
func (o *GetCloudProjectServiceNameSnapshotSnapshotIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cloud project service name snapshot snapshot ID params
func (o *GetCloudProjectServiceNameSnapshotSnapshotIDParams) WithContext(ctx context.Context) *GetCloudProjectServiceNameSnapshotSnapshotIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cloud project service name snapshot snapshot ID params
func (o *GetCloudProjectServiceNameSnapshotSnapshotIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cloud project service name snapshot snapshot ID params
func (o *GetCloudProjectServiceNameSnapshotSnapshotIDParams) WithHTTPClient(client *http.Client) *GetCloudProjectServiceNameSnapshotSnapshotIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cloud project service name snapshot snapshot ID params
func (o *GetCloudProjectServiceNameSnapshotSnapshotIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceName adds the serviceName to the get cloud project service name snapshot snapshot ID params
func (o *GetCloudProjectServiceNameSnapshotSnapshotIDParams) WithServiceName(serviceName string) *GetCloudProjectServiceNameSnapshotSnapshotIDParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get cloud project service name snapshot snapshot ID params
func (o *GetCloudProjectServiceNameSnapshotSnapshotIDParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WithSnapshotID adds the snapshotID to the get cloud project service name snapshot snapshot ID params
func (o *GetCloudProjectServiceNameSnapshotSnapshotIDParams) WithSnapshotID(snapshotID string) *GetCloudProjectServiceNameSnapshotSnapshotIDParams {
	o.SetSnapshotID(snapshotID)
	return o
}

// SetSnapshotID adds the snapshotId to the get cloud project service name snapshot snapshot ID params
func (o *GetCloudProjectServiceNameSnapshotSnapshotIDParams) SetSnapshotID(snapshotID string) {
	o.SnapshotID = snapshotID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCloudProjectServiceNameSnapshotSnapshotIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	// path param snapshotId
	if err := r.SetPathParam("snapshotId", o.SnapshotID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}