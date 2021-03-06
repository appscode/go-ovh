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

// NewPostDomainServiceNameGlueRecordHostUpdateParams creates a new PostDomainServiceNameGlueRecordHostUpdateParams object
// with the default values initialized.
func NewPostDomainServiceNameGlueRecordHostUpdateParams() *PostDomainServiceNameGlueRecordHostUpdateParams {
	var ()
	return &PostDomainServiceNameGlueRecordHostUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDomainServiceNameGlueRecordHostUpdateParamsWithTimeout creates a new PostDomainServiceNameGlueRecordHostUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDomainServiceNameGlueRecordHostUpdateParamsWithTimeout(timeout time.Duration) *PostDomainServiceNameGlueRecordHostUpdateParams {
	var ()
	return &PostDomainServiceNameGlueRecordHostUpdateParams{

		timeout: timeout,
	}
}

// NewPostDomainServiceNameGlueRecordHostUpdateParamsWithContext creates a new PostDomainServiceNameGlueRecordHostUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDomainServiceNameGlueRecordHostUpdateParamsWithContext(ctx context.Context) *PostDomainServiceNameGlueRecordHostUpdateParams {
	var ()
	return &PostDomainServiceNameGlueRecordHostUpdateParams{

		Context: ctx,
	}
}

// NewPostDomainServiceNameGlueRecordHostUpdateParamsWithHTTPClient creates a new PostDomainServiceNameGlueRecordHostUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDomainServiceNameGlueRecordHostUpdateParamsWithHTTPClient(client *http.Client) *PostDomainServiceNameGlueRecordHostUpdateParams {
	var ()
	return &PostDomainServiceNameGlueRecordHostUpdateParams{
		HTTPClient: client,
	}
}

/*PostDomainServiceNameGlueRecordHostUpdateParams contains all the parameters to send to the API endpoint
for the post domain service name glue record host update operation typically these are written to a http.Request
*/
type PostDomainServiceNameGlueRecordHostUpdateParams struct {

	/*DomainGlueRecordUpdatePost*/
	DomainGlueRecordUpdatePost *models.PostDomainServiceNameGlueRecordHostUpdateParamsBody
	/*Host*/
	Host string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post domain service name glue record host update params
func (o *PostDomainServiceNameGlueRecordHostUpdateParams) WithTimeout(timeout time.Duration) *PostDomainServiceNameGlueRecordHostUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post domain service name glue record host update params
func (o *PostDomainServiceNameGlueRecordHostUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post domain service name glue record host update params
func (o *PostDomainServiceNameGlueRecordHostUpdateParams) WithContext(ctx context.Context) *PostDomainServiceNameGlueRecordHostUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post domain service name glue record host update params
func (o *PostDomainServiceNameGlueRecordHostUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post domain service name glue record host update params
func (o *PostDomainServiceNameGlueRecordHostUpdateParams) WithHTTPClient(client *http.Client) *PostDomainServiceNameGlueRecordHostUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post domain service name glue record host update params
func (o *PostDomainServiceNameGlueRecordHostUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainGlueRecordUpdatePost adds the domainGlueRecordUpdatePost to the post domain service name glue record host update params
func (o *PostDomainServiceNameGlueRecordHostUpdateParams) WithDomainGlueRecordUpdatePost(domainGlueRecordUpdatePost *models.PostDomainServiceNameGlueRecordHostUpdateParamsBody) *PostDomainServiceNameGlueRecordHostUpdateParams {
	o.SetDomainGlueRecordUpdatePost(domainGlueRecordUpdatePost)
	return o
}

// SetDomainGlueRecordUpdatePost adds the domainGlueRecordUpdatePost to the post domain service name glue record host update params
func (o *PostDomainServiceNameGlueRecordHostUpdateParams) SetDomainGlueRecordUpdatePost(domainGlueRecordUpdatePost *models.PostDomainServiceNameGlueRecordHostUpdateParamsBody) {
	o.DomainGlueRecordUpdatePost = domainGlueRecordUpdatePost
}

// WithHost adds the host to the post domain service name glue record host update params
func (o *PostDomainServiceNameGlueRecordHostUpdateParams) WithHost(host string) *PostDomainServiceNameGlueRecordHostUpdateParams {
	o.SetHost(host)
	return o
}

// SetHost adds the host to the post domain service name glue record host update params
func (o *PostDomainServiceNameGlueRecordHostUpdateParams) SetHost(host string) {
	o.Host = host
}

// WithServiceName adds the serviceName to the post domain service name glue record host update params
func (o *PostDomainServiceNameGlueRecordHostUpdateParams) WithServiceName(serviceName string) *PostDomainServiceNameGlueRecordHostUpdateParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post domain service name glue record host update params
func (o *PostDomainServiceNameGlueRecordHostUpdateParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostDomainServiceNameGlueRecordHostUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DomainGlueRecordUpdatePost != nil {
		if err := r.SetBodyParam(o.DomainGlueRecordUpdatePost); err != nil {
			return err
		}
	}

	// path param host
	if err := r.SetPathParam("host", o.Host); err != nil {
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
