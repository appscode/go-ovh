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

// NewGetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams creates a new GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams object
// with the default values initialized.
func NewGetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams() *GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams {
	var ()
	return &GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParamsWithTimeout creates a new GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParamsWithTimeout(timeout time.Duration) *GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams {
	var ()
	return &GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams{

		timeout: timeout,
	}
}

// NewGetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParamsWithContext creates a new GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParamsWithContext(ctx context.Context) *GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams {
	var ()
	return &GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams{

		Context: ctx,
	}
}

// NewGetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParamsWithHTTPClient creates a new GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParamsWithHTTPClient(client *http.Client) *GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams {
	var ()
	return &GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams{
		HTTPClient: client,
	}
}

/*GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams contains all the parameters to send to the API endpoint
for the get vrack service name dedicated server interface dedicated server interface operation typically these are written to a http.Request
*/
type GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams struct {

	/*DedicatedServerInterface*/
	DedicatedServerInterface string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vrack service name dedicated server interface dedicated server interface params
func (o *GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams) WithTimeout(timeout time.Duration) *GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vrack service name dedicated server interface dedicated server interface params
func (o *GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vrack service name dedicated server interface dedicated server interface params
func (o *GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams) WithContext(ctx context.Context) *GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vrack service name dedicated server interface dedicated server interface params
func (o *GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vrack service name dedicated server interface dedicated server interface params
func (o *GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams) WithHTTPClient(client *http.Client) *GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vrack service name dedicated server interface dedicated server interface params
func (o *GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDedicatedServerInterface adds the dedicatedServerInterface to the get vrack service name dedicated server interface dedicated server interface params
func (o *GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams) WithDedicatedServerInterface(dedicatedServerInterface string) *GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams {
	o.SetDedicatedServerInterface(dedicatedServerInterface)
	return o
}

// SetDedicatedServerInterface adds the dedicatedServerInterface to the get vrack service name dedicated server interface dedicated server interface params
func (o *GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams) SetDedicatedServerInterface(dedicatedServerInterface string) {
	o.DedicatedServerInterface = dedicatedServerInterface
}

// WithServiceName adds the serviceName to the get vrack service name dedicated server interface dedicated server interface params
func (o *GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams) WithServiceName(serviceName string) *GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get vrack service name dedicated server interface dedicated server interface params
func (o *GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dedicatedServerInterface
	if err := r.SetPathParam("dedicatedServerInterface", o.DedicatedServerInterface); err != nil {
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