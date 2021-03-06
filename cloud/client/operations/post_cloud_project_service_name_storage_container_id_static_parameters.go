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

// NewPostCloudProjectServiceNameStorageContainerIDStaticParams creates a new PostCloudProjectServiceNameStorageContainerIDStaticParams object
// with the default values initialized.
func NewPostCloudProjectServiceNameStorageContainerIDStaticParams() *PostCloudProjectServiceNameStorageContainerIDStaticParams {
	var ()
	return &PostCloudProjectServiceNameStorageContainerIDStaticParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCloudProjectServiceNameStorageContainerIDStaticParamsWithTimeout creates a new PostCloudProjectServiceNameStorageContainerIDStaticParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCloudProjectServiceNameStorageContainerIDStaticParamsWithTimeout(timeout time.Duration) *PostCloudProjectServiceNameStorageContainerIDStaticParams {
	var ()
	return &PostCloudProjectServiceNameStorageContainerIDStaticParams{

		timeout: timeout,
	}
}

// NewPostCloudProjectServiceNameStorageContainerIDStaticParamsWithContext creates a new PostCloudProjectServiceNameStorageContainerIDStaticParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCloudProjectServiceNameStorageContainerIDStaticParamsWithContext(ctx context.Context) *PostCloudProjectServiceNameStorageContainerIDStaticParams {
	var ()
	return &PostCloudProjectServiceNameStorageContainerIDStaticParams{

		Context: ctx,
	}
}

// NewPostCloudProjectServiceNameStorageContainerIDStaticParamsWithHTTPClient creates a new PostCloudProjectServiceNameStorageContainerIDStaticParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCloudProjectServiceNameStorageContainerIDStaticParamsWithHTTPClient(client *http.Client) *PostCloudProjectServiceNameStorageContainerIDStaticParams {
	var ()
	return &PostCloudProjectServiceNameStorageContainerIDStaticParams{
		HTTPClient: client,
	}
}

/*PostCloudProjectServiceNameStorageContainerIDStaticParams contains all the parameters to send to the API endpoint
for the post cloud project service name storage container ID static operation typically these are written to a http.Request
*/
type PostCloudProjectServiceNameStorageContainerIDStaticParams struct {

	/*ContainerID*/
	ContainerID string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post cloud project service name storage container ID static params
func (o *PostCloudProjectServiceNameStorageContainerIDStaticParams) WithTimeout(timeout time.Duration) *PostCloudProjectServiceNameStorageContainerIDStaticParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cloud project service name storage container ID static params
func (o *PostCloudProjectServiceNameStorageContainerIDStaticParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cloud project service name storage container ID static params
func (o *PostCloudProjectServiceNameStorageContainerIDStaticParams) WithContext(ctx context.Context) *PostCloudProjectServiceNameStorageContainerIDStaticParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cloud project service name storage container ID static params
func (o *PostCloudProjectServiceNameStorageContainerIDStaticParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cloud project service name storage container ID static params
func (o *PostCloudProjectServiceNameStorageContainerIDStaticParams) WithHTTPClient(client *http.Client) *PostCloudProjectServiceNameStorageContainerIDStaticParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cloud project service name storage container ID static params
func (o *PostCloudProjectServiceNameStorageContainerIDStaticParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContainerID adds the containerID to the post cloud project service name storage container ID static params
func (o *PostCloudProjectServiceNameStorageContainerIDStaticParams) WithContainerID(containerID string) *PostCloudProjectServiceNameStorageContainerIDStaticParams {
	o.SetContainerID(containerID)
	return o
}

// SetContainerID adds the containerId to the post cloud project service name storage container ID static params
func (o *PostCloudProjectServiceNameStorageContainerIDStaticParams) SetContainerID(containerID string) {
	o.ContainerID = containerID
}

// WithServiceName adds the serviceName to the post cloud project service name storage container ID static params
func (o *PostCloudProjectServiceNameStorageContainerIDStaticParams) WithServiceName(serviceName string) *PostCloudProjectServiceNameStorageContainerIDStaticParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post cloud project service name storage container ID static params
func (o *PostCloudProjectServiceNameStorageContainerIDStaticParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostCloudProjectServiceNameStorageContainerIDStaticParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param containerId
	if err := r.SetPathParam("containerId", o.ContainerID); err != nil {
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
