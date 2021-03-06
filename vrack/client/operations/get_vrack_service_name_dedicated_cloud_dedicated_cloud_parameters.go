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

// NewGetVrackServiceNameDedicatedCloudDedicatedCloudParams creates a new GetVrackServiceNameDedicatedCloudDedicatedCloudParams object
// with the default values initialized.
func NewGetVrackServiceNameDedicatedCloudDedicatedCloudParams() *GetVrackServiceNameDedicatedCloudDedicatedCloudParams {
	var ()
	return &GetVrackServiceNameDedicatedCloudDedicatedCloudParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVrackServiceNameDedicatedCloudDedicatedCloudParamsWithTimeout creates a new GetVrackServiceNameDedicatedCloudDedicatedCloudParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVrackServiceNameDedicatedCloudDedicatedCloudParamsWithTimeout(timeout time.Duration) *GetVrackServiceNameDedicatedCloudDedicatedCloudParams {
	var ()
	return &GetVrackServiceNameDedicatedCloudDedicatedCloudParams{

		timeout: timeout,
	}
}

// NewGetVrackServiceNameDedicatedCloudDedicatedCloudParamsWithContext creates a new GetVrackServiceNameDedicatedCloudDedicatedCloudParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVrackServiceNameDedicatedCloudDedicatedCloudParamsWithContext(ctx context.Context) *GetVrackServiceNameDedicatedCloudDedicatedCloudParams {
	var ()
	return &GetVrackServiceNameDedicatedCloudDedicatedCloudParams{

		Context: ctx,
	}
}

// NewGetVrackServiceNameDedicatedCloudDedicatedCloudParamsWithHTTPClient creates a new GetVrackServiceNameDedicatedCloudDedicatedCloudParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVrackServiceNameDedicatedCloudDedicatedCloudParamsWithHTTPClient(client *http.Client) *GetVrackServiceNameDedicatedCloudDedicatedCloudParams {
	var ()
	return &GetVrackServiceNameDedicatedCloudDedicatedCloudParams{
		HTTPClient: client,
	}
}

/*GetVrackServiceNameDedicatedCloudDedicatedCloudParams contains all the parameters to send to the API endpoint
for the get vrack service name dedicated cloud dedicated cloud operation typically these are written to a http.Request
*/
type GetVrackServiceNameDedicatedCloudDedicatedCloudParams struct {

	/*DedicatedCloud*/
	DedicatedCloud string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vrack service name dedicated cloud dedicated cloud params
func (o *GetVrackServiceNameDedicatedCloudDedicatedCloudParams) WithTimeout(timeout time.Duration) *GetVrackServiceNameDedicatedCloudDedicatedCloudParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vrack service name dedicated cloud dedicated cloud params
func (o *GetVrackServiceNameDedicatedCloudDedicatedCloudParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vrack service name dedicated cloud dedicated cloud params
func (o *GetVrackServiceNameDedicatedCloudDedicatedCloudParams) WithContext(ctx context.Context) *GetVrackServiceNameDedicatedCloudDedicatedCloudParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vrack service name dedicated cloud dedicated cloud params
func (o *GetVrackServiceNameDedicatedCloudDedicatedCloudParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vrack service name dedicated cloud dedicated cloud params
func (o *GetVrackServiceNameDedicatedCloudDedicatedCloudParams) WithHTTPClient(client *http.Client) *GetVrackServiceNameDedicatedCloudDedicatedCloudParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vrack service name dedicated cloud dedicated cloud params
func (o *GetVrackServiceNameDedicatedCloudDedicatedCloudParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDedicatedCloud adds the dedicatedCloud to the get vrack service name dedicated cloud dedicated cloud params
func (o *GetVrackServiceNameDedicatedCloudDedicatedCloudParams) WithDedicatedCloud(dedicatedCloud string) *GetVrackServiceNameDedicatedCloudDedicatedCloudParams {
	o.SetDedicatedCloud(dedicatedCloud)
	return o
}

// SetDedicatedCloud adds the dedicatedCloud to the get vrack service name dedicated cloud dedicated cloud params
func (o *GetVrackServiceNameDedicatedCloudDedicatedCloudParams) SetDedicatedCloud(dedicatedCloud string) {
	o.DedicatedCloud = dedicatedCloud
}

// WithServiceName adds the serviceName to the get vrack service name dedicated cloud dedicated cloud params
func (o *GetVrackServiceNameDedicatedCloudDedicatedCloudParams) WithServiceName(serviceName string) *GetVrackServiceNameDedicatedCloudDedicatedCloudParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get vrack service name dedicated cloud dedicated cloud params
func (o *GetVrackServiceNameDedicatedCloudDedicatedCloudParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetVrackServiceNameDedicatedCloudDedicatedCloudParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dedicatedCloud
	if err := r.SetPathParam("dedicatedCloud", o.DedicatedCloud); err != nil {
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
