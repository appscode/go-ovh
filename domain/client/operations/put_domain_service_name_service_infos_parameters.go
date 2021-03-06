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

	"github.com/appscode/go-ovh/domain/models"
)

// NewPutDomainServiceNameServiceInfosParams creates a new PutDomainServiceNameServiceInfosParams object
// with the default values initialized.
func NewPutDomainServiceNameServiceInfosParams() *PutDomainServiceNameServiceInfosParams {
	var ()
	return &PutDomainServiceNameServiceInfosParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutDomainServiceNameServiceInfosParamsWithTimeout creates a new PutDomainServiceNameServiceInfosParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutDomainServiceNameServiceInfosParamsWithTimeout(timeout time.Duration) *PutDomainServiceNameServiceInfosParams {
	var ()
	return &PutDomainServiceNameServiceInfosParams{

		timeout: timeout,
	}
}

// NewPutDomainServiceNameServiceInfosParamsWithContext creates a new PutDomainServiceNameServiceInfosParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutDomainServiceNameServiceInfosParamsWithContext(ctx context.Context) *PutDomainServiceNameServiceInfosParams {
	var ()
	return &PutDomainServiceNameServiceInfosParams{

		Context: ctx,
	}
}

// NewPutDomainServiceNameServiceInfosParamsWithHTTPClient creates a new PutDomainServiceNameServiceInfosParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutDomainServiceNameServiceInfosParamsWithHTTPClient(client *http.Client) *PutDomainServiceNameServiceInfosParams {
	var ()
	return &PutDomainServiceNameServiceInfosParams{
		HTTPClient: client,
	}
}

/*PutDomainServiceNameServiceInfosParams contains all the parameters to send to the API endpoint
for the put domain service name service infos operation typically these are written to a http.Request
*/
type PutDomainServiceNameServiceInfosParams struct {

	/*DomainServiceInfosPut*/
	DomainServiceInfosPut *models.ServicesService
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put domain service name service infos params
func (o *PutDomainServiceNameServiceInfosParams) WithTimeout(timeout time.Duration) *PutDomainServiceNameServiceInfosParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put domain service name service infos params
func (o *PutDomainServiceNameServiceInfosParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put domain service name service infos params
func (o *PutDomainServiceNameServiceInfosParams) WithContext(ctx context.Context) *PutDomainServiceNameServiceInfosParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put domain service name service infos params
func (o *PutDomainServiceNameServiceInfosParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put domain service name service infos params
func (o *PutDomainServiceNameServiceInfosParams) WithHTTPClient(client *http.Client) *PutDomainServiceNameServiceInfosParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put domain service name service infos params
func (o *PutDomainServiceNameServiceInfosParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainServiceInfosPut adds the domainServiceInfosPut to the put domain service name service infos params
func (o *PutDomainServiceNameServiceInfosParams) WithDomainServiceInfosPut(domainServiceInfosPut *models.ServicesService) *PutDomainServiceNameServiceInfosParams {
	o.SetDomainServiceInfosPut(domainServiceInfosPut)
	return o
}

// SetDomainServiceInfosPut adds the domainServiceInfosPut to the put domain service name service infos params
func (o *PutDomainServiceNameServiceInfosParams) SetDomainServiceInfosPut(domainServiceInfosPut *models.ServicesService) {
	o.DomainServiceInfosPut = domainServiceInfosPut
}

// WithServiceName adds the serviceName to the put domain service name service infos params
func (o *PutDomainServiceNameServiceInfosParams) WithServiceName(serviceName string) *PutDomainServiceNameServiceInfosParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the put domain service name service infos params
func (o *PutDomainServiceNameServiceInfosParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PutDomainServiceNameServiceInfosParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DomainServiceInfosPut != nil {
		if err := r.SetBodyParam(o.DomainServiceInfosPut); err != nil {
			return err
		}
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
