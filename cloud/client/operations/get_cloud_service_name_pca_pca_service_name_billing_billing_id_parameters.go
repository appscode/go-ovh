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

// NewGetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams creates a new GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams object
// with the default values initialized.
func NewGetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams() *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams {
	var ()
	return &GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCloudServiceNamePcaPcaServiceNameBillingBillingIDParamsWithTimeout creates a new GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCloudServiceNamePcaPcaServiceNameBillingBillingIDParamsWithTimeout(timeout time.Duration) *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams {
	var ()
	return &GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams{

		timeout: timeout,
	}
}

// NewGetCloudServiceNamePcaPcaServiceNameBillingBillingIDParamsWithContext creates a new GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCloudServiceNamePcaPcaServiceNameBillingBillingIDParamsWithContext(ctx context.Context) *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams {
	var ()
	return &GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams{

		Context: ctx,
	}
}

// NewGetCloudServiceNamePcaPcaServiceNameBillingBillingIDParamsWithHTTPClient creates a new GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCloudServiceNamePcaPcaServiceNameBillingBillingIDParamsWithHTTPClient(client *http.Client) *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams {
	var ()
	return &GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams{
		HTTPClient: client,
	}
}

/*GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams contains all the parameters to send to the API endpoint
for the get cloud service name pca pca service name billing billing ID operation typically these are written to a http.Request
*/
type GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams struct {

	/*BillingID*/
	BillingID int64
	/*PcaServiceName*/
	PcaServiceName string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cloud service name pca pca service name billing billing ID params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams) WithTimeout(timeout time.Duration) *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cloud service name pca pca service name billing billing ID params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cloud service name pca pca service name billing billing ID params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams) WithContext(ctx context.Context) *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cloud service name pca pca service name billing billing ID params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cloud service name pca pca service name billing billing ID params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams) WithHTTPClient(client *http.Client) *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cloud service name pca pca service name billing billing ID params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBillingID adds the billingID to the get cloud service name pca pca service name billing billing ID params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams) WithBillingID(billingID int64) *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams {
	o.SetBillingID(billingID)
	return o
}

// SetBillingID adds the billingId to the get cloud service name pca pca service name billing billing ID params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams) SetBillingID(billingID int64) {
	o.BillingID = billingID
}

// WithPcaServiceName adds the pcaServiceName to the get cloud service name pca pca service name billing billing ID params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams) WithPcaServiceName(pcaServiceName string) *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams {
	o.SetPcaServiceName(pcaServiceName)
	return o
}

// SetPcaServiceName adds the pcaServiceName to the get cloud service name pca pca service name billing billing ID params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams) SetPcaServiceName(pcaServiceName string) {
	o.PcaServiceName = pcaServiceName
}

// WithServiceName adds the serviceName to the get cloud service name pca pca service name billing billing ID params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams) WithServiceName(serviceName string) *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get cloud service name pca pca service name billing billing ID params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetCloudServiceNamePcaPcaServiceNameBillingBillingIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param billingId
	if err := r.SetPathParam("billingId", swag.FormatInt64(o.BillingID)); err != nil {
		return err
	}

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
