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

// NewGetVrackServiceNameCloudProjectParams creates a new GetVrackServiceNameCloudProjectParams object
// with the default values initialized.
func NewGetVrackServiceNameCloudProjectParams() *GetVrackServiceNameCloudProjectParams {
	var ()
	return &GetVrackServiceNameCloudProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVrackServiceNameCloudProjectParamsWithTimeout creates a new GetVrackServiceNameCloudProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVrackServiceNameCloudProjectParamsWithTimeout(timeout time.Duration) *GetVrackServiceNameCloudProjectParams {
	var ()
	return &GetVrackServiceNameCloudProjectParams{

		timeout: timeout,
	}
}

// NewGetVrackServiceNameCloudProjectParamsWithContext creates a new GetVrackServiceNameCloudProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVrackServiceNameCloudProjectParamsWithContext(ctx context.Context) *GetVrackServiceNameCloudProjectParams {
	var ()
	return &GetVrackServiceNameCloudProjectParams{

		Context: ctx,
	}
}

// NewGetVrackServiceNameCloudProjectParamsWithHTTPClient creates a new GetVrackServiceNameCloudProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVrackServiceNameCloudProjectParamsWithHTTPClient(client *http.Client) *GetVrackServiceNameCloudProjectParams {
	var ()
	return &GetVrackServiceNameCloudProjectParams{
		HTTPClient: client,
	}
}

/*GetVrackServiceNameCloudProjectParams contains all the parameters to send to the API endpoint
for the get vrack service name cloud project operation typically these are written to a http.Request
*/
type GetVrackServiceNameCloudProjectParams struct {

	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vrack service name cloud project params
func (o *GetVrackServiceNameCloudProjectParams) WithTimeout(timeout time.Duration) *GetVrackServiceNameCloudProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vrack service name cloud project params
func (o *GetVrackServiceNameCloudProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vrack service name cloud project params
func (o *GetVrackServiceNameCloudProjectParams) WithContext(ctx context.Context) *GetVrackServiceNameCloudProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vrack service name cloud project params
func (o *GetVrackServiceNameCloudProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vrack service name cloud project params
func (o *GetVrackServiceNameCloudProjectParams) WithHTTPClient(client *http.Client) *GetVrackServiceNameCloudProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vrack service name cloud project params
func (o *GetVrackServiceNameCloudProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceName adds the serviceName to the get vrack service name cloud project params
func (o *GetVrackServiceNameCloudProjectParams) WithServiceName(serviceName string) *GetVrackServiceNameCloudProjectParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get vrack service name cloud project params
func (o *GetVrackServiceNameCloudProjectParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetVrackServiceNameCloudProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
