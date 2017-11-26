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

// NewGetCloudServiceNamePcaPcaServiceNameBillingParams creates a new GetCloudServiceNamePcaPcaServiceNameBillingParams object
// with the default values initialized.
func NewGetCloudServiceNamePcaPcaServiceNameBillingParams() *GetCloudServiceNamePcaPcaServiceNameBillingParams {
	var (
		billedDefault = bool(false)
	)
	return &GetCloudServiceNamePcaPcaServiceNameBillingParams{
		Billed: &billedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCloudServiceNamePcaPcaServiceNameBillingParamsWithTimeout creates a new GetCloudServiceNamePcaPcaServiceNameBillingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCloudServiceNamePcaPcaServiceNameBillingParamsWithTimeout(timeout time.Duration) *GetCloudServiceNamePcaPcaServiceNameBillingParams {
	var (
		billedDefault = bool(false)
	)
	return &GetCloudServiceNamePcaPcaServiceNameBillingParams{
		Billed: &billedDefault,

		timeout: timeout,
	}
}

// NewGetCloudServiceNamePcaPcaServiceNameBillingParamsWithContext creates a new GetCloudServiceNamePcaPcaServiceNameBillingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCloudServiceNamePcaPcaServiceNameBillingParamsWithContext(ctx context.Context) *GetCloudServiceNamePcaPcaServiceNameBillingParams {
	var (
		billedDefault = bool(false)
	)
	return &GetCloudServiceNamePcaPcaServiceNameBillingParams{
		Billed: &billedDefault,

		Context: ctx,
	}
}

// NewGetCloudServiceNamePcaPcaServiceNameBillingParamsWithHTTPClient creates a new GetCloudServiceNamePcaPcaServiceNameBillingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCloudServiceNamePcaPcaServiceNameBillingParamsWithHTTPClient(client *http.Client) *GetCloudServiceNamePcaPcaServiceNameBillingParams {
	var (
		billedDefault = bool(false)
	)
	return &GetCloudServiceNamePcaPcaServiceNameBillingParams{
		Billed:     &billedDefault,
		HTTPClient: client,
	}
}

/*GetCloudServiceNamePcaPcaServiceNameBillingParams contains all the parameters to send to the API endpoint
for the get cloud service name pca pca service name billing operation typically these are written to a http.Request
*/
type GetCloudServiceNamePcaPcaServiceNameBillingParams struct {

	/*Billed*/
	Billed *bool
	/*DateFrom*/
	DateFrom *strfmt.DateTime
	/*DateTo*/
	DateTo *strfmt.DateTime
	/*PcaServiceName*/
	PcaServiceName string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cloud service name pca pca service name billing params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingParams) WithTimeout(timeout time.Duration) *GetCloudServiceNamePcaPcaServiceNameBillingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cloud service name pca pca service name billing params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cloud service name pca pca service name billing params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingParams) WithContext(ctx context.Context) *GetCloudServiceNamePcaPcaServiceNameBillingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cloud service name pca pca service name billing params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cloud service name pca pca service name billing params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingParams) WithHTTPClient(client *http.Client) *GetCloudServiceNamePcaPcaServiceNameBillingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cloud service name pca pca service name billing params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBilled adds the billed to the get cloud service name pca pca service name billing params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingParams) WithBilled(billed *bool) *GetCloudServiceNamePcaPcaServiceNameBillingParams {
	o.SetBilled(billed)
	return o
}

// SetBilled adds the billed to the get cloud service name pca pca service name billing params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingParams) SetBilled(billed *bool) {
	o.Billed = billed
}

// WithDateFrom adds the dateFrom to the get cloud service name pca pca service name billing params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingParams) WithDateFrom(dateFrom *strfmt.DateTime) *GetCloudServiceNamePcaPcaServiceNameBillingParams {
	o.SetDateFrom(dateFrom)
	return o
}

// SetDateFrom adds the dateFrom to the get cloud service name pca pca service name billing params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingParams) SetDateFrom(dateFrom *strfmt.DateTime) {
	o.DateFrom = dateFrom
}

// WithDateTo adds the dateTo to the get cloud service name pca pca service name billing params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingParams) WithDateTo(dateTo *strfmt.DateTime) *GetCloudServiceNamePcaPcaServiceNameBillingParams {
	o.SetDateTo(dateTo)
	return o
}

// SetDateTo adds the dateTo to the get cloud service name pca pca service name billing params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingParams) SetDateTo(dateTo *strfmt.DateTime) {
	o.DateTo = dateTo
}

// WithPcaServiceName adds the pcaServiceName to the get cloud service name pca pca service name billing params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingParams) WithPcaServiceName(pcaServiceName string) *GetCloudServiceNamePcaPcaServiceNameBillingParams {
	o.SetPcaServiceName(pcaServiceName)
	return o
}

// SetPcaServiceName adds the pcaServiceName to the get cloud service name pca pca service name billing params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingParams) SetPcaServiceName(pcaServiceName string) {
	o.PcaServiceName = pcaServiceName
}

// WithServiceName adds the serviceName to the get cloud service name pca pca service name billing params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingParams) WithServiceName(serviceName string) *GetCloudServiceNamePcaPcaServiceNameBillingParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get cloud service name pca pca service name billing params
func (o *GetCloudServiceNamePcaPcaServiceNameBillingParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetCloudServiceNamePcaPcaServiceNameBillingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Billed != nil {

		// query param billed
		var qrBilled bool
		if o.Billed != nil {
			qrBilled = *o.Billed
		}
		qBilled := swag.FormatBool(qrBilled)
		if qBilled != "" {
			if err := r.SetQueryParam("billed", qBilled); err != nil {
				return err
			}
		}

	}

	if o.DateFrom != nil {

		// query param date.from
		var qrDateFrom strfmt.DateTime
		if o.DateFrom != nil {
			qrDateFrom = *o.DateFrom
		}
		qDateFrom := qrDateFrom.String()
		if qDateFrom != "" {
			if err := r.SetQueryParam("date.from", qDateFrom); err != nil {
				return err
			}
		}

	}

	if o.DateTo != nil {

		// query param date.to
		var qrDateTo strfmt.DateTime
		if o.DateTo != nil {
			qrDateTo = *o.DateTo
		}
		qDateTo := qrDateTo.String()
		if qDateTo != "" {
			if err := r.SetQueryParam("date.to", qDateTo); err != nil {
				return err
			}
		}

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