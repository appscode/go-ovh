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

// NewGetCloudServiceNamePcaPcaServiceNameParams creates a new GetCloudServiceNamePcaPcaServiceNameParams object
// with the default values initialized.
func NewGetCloudServiceNamePcaPcaServiceNameParams() *GetCloudServiceNamePcaPcaServiceNameParams {
	var ()
	return &GetCloudServiceNamePcaPcaServiceNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCloudServiceNamePcaPcaServiceNameParamsWithTimeout creates a new GetCloudServiceNamePcaPcaServiceNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCloudServiceNamePcaPcaServiceNameParamsWithTimeout(timeout time.Duration) *GetCloudServiceNamePcaPcaServiceNameParams {
	var ()
	return &GetCloudServiceNamePcaPcaServiceNameParams{

		timeout: timeout,
	}
}

// NewGetCloudServiceNamePcaPcaServiceNameParamsWithContext creates a new GetCloudServiceNamePcaPcaServiceNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCloudServiceNamePcaPcaServiceNameParamsWithContext(ctx context.Context) *GetCloudServiceNamePcaPcaServiceNameParams {
	var ()
	return &GetCloudServiceNamePcaPcaServiceNameParams{

		Context: ctx,
	}
}

// NewGetCloudServiceNamePcaPcaServiceNameParamsWithHTTPClient creates a new GetCloudServiceNamePcaPcaServiceNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCloudServiceNamePcaPcaServiceNameParamsWithHTTPClient(client *http.Client) *GetCloudServiceNamePcaPcaServiceNameParams {
	var ()
	return &GetCloudServiceNamePcaPcaServiceNameParams{
		HTTPClient: client,
	}
}

/*GetCloudServiceNamePcaPcaServiceNameParams contains all the parameters to send to the API endpoint
for the get cloud service name pca pca service name operation typically these are written to a http.Request
*/
type GetCloudServiceNamePcaPcaServiceNameParams struct {

	/*PcaServiceName*/
	PcaServiceName string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cloud service name pca pca service name params
func (o *GetCloudServiceNamePcaPcaServiceNameParams) WithTimeout(timeout time.Duration) *GetCloudServiceNamePcaPcaServiceNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cloud service name pca pca service name params
func (o *GetCloudServiceNamePcaPcaServiceNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cloud service name pca pca service name params
func (o *GetCloudServiceNamePcaPcaServiceNameParams) WithContext(ctx context.Context) *GetCloudServiceNamePcaPcaServiceNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cloud service name pca pca service name params
func (o *GetCloudServiceNamePcaPcaServiceNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cloud service name pca pca service name params
func (o *GetCloudServiceNamePcaPcaServiceNameParams) WithHTTPClient(client *http.Client) *GetCloudServiceNamePcaPcaServiceNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cloud service name pca pca service name params
func (o *GetCloudServiceNamePcaPcaServiceNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPcaServiceName adds the pcaServiceName to the get cloud service name pca pca service name params
func (o *GetCloudServiceNamePcaPcaServiceNameParams) WithPcaServiceName(pcaServiceName string) *GetCloudServiceNamePcaPcaServiceNameParams {
	o.SetPcaServiceName(pcaServiceName)
	return o
}

// SetPcaServiceName adds the pcaServiceName to the get cloud service name pca pca service name params
func (o *GetCloudServiceNamePcaPcaServiceNameParams) SetPcaServiceName(pcaServiceName string) {
	o.PcaServiceName = pcaServiceName
}

// WithServiceName adds the serviceName to the get cloud service name pca pca service name params
func (o *GetCloudServiceNamePcaPcaServiceNameParams) WithServiceName(serviceName string) *GetCloudServiceNamePcaPcaServiceNameParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get cloud service name pca pca service name params
func (o *GetCloudServiceNamePcaPcaServiceNameParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetCloudServiceNamePcaPcaServiceNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param pcaServiceName
	if err := r.SetPathParam("pcaServiceName", o.PcaServiceName); err != nil {
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
