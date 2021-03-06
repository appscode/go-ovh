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

// NewPostDomainServiceNameNameServersUpdateParams creates a new PostDomainServiceNameNameServersUpdateParams object
// with the default values initialized.
func NewPostDomainServiceNameNameServersUpdateParams() *PostDomainServiceNameNameServersUpdateParams {
	var ()
	return &PostDomainServiceNameNameServersUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDomainServiceNameNameServersUpdateParamsWithTimeout creates a new PostDomainServiceNameNameServersUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDomainServiceNameNameServersUpdateParamsWithTimeout(timeout time.Duration) *PostDomainServiceNameNameServersUpdateParams {
	var ()
	return &PostDomainServiceNameNameServersUpdateParams{

		timeout: timeout,
	}
}

// NewPostDomainServiceNameNameServersUpdateParamsWithContext creates a new PostDomainServiceNameNameServersUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDomainServiceNameNameServersUpdateParamsWithContext(ctx context.Context) *PostDomainServiceNameNameServersUpdateParams {
	var ()
	return &PostDomainServiceNameNameServersUpdateParams{

		Context: ctx,
	}
}

// NewPostDomainServiceNameNameServersUpdateParamsWithHTTPClient creates a new PostDomainServiceNameNameServersUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDomainServiceNameNameServersUpdateParamsWithHTTPClient(client *http.Client) *PostDomainServiceNameNameServersUpdateParams {
	var ()
	return &PostDomainServiceNameNameServersUpdateParams{
		HTTPClient: client,
	}
}

/*PostDomainServiceNameNameServersUpdateParams contains all the parameters to send to the API endpoint
for the post domain service name name servers update operation typically these are written to a http.Request
*/
type PostDomainServiceNameNameServersUpdateParams struct {

	/*DomainNameServersUpdatePost*/
	DomainNameServersUpdatePost *models.PostDomainServiceNameNameServersUpdateParamsBody
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post domain service name name servers update params
func (o *PostDomainServiceNameNameServersUpdateParams) WithTimeout(timeout time.Duration) *PostDomainServiceNameNameServersUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post domain service name name servers update params
func (o *PostDomainServiceNameNameServersUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post domain service name name servers update params
func (o *PostDomainServiceNameNameServersUpdateParams) WithContext(ctx context.Context) *PostDomainServiceNameNameServersUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post domain service name name servers update params
func (o *PostDomainServiceNameNameServersUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post domain service name name servers update params
func (o *PostDomainServiceNameNameServersUpdateParams) WithHTTPClient(client *http.Client) *PostDomainServiceNameNameServersUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post domain service name name servers update params
func (o *PostDomainServiceNameNameServersUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainNameServersUpdatePost adds the domainNameServersUpdatePost to the post domain service name name servers update params
func (o *PostDomainServiceNameNameServersUpdateParams) WithDomainNameServersUpdatePost(domainNameServersUpdatePost *models.PostDomainServiceNameNameServersUpdateParamsBody) *PostDomainServiceNameNameServersUpdateParams {
	o.SetDomainNameServersUpdatePost(domainNameServersUpdatePost)
	return o
}

// SetDomainNameServersUpdatePost adds the domainNameServersUpdatePost to the post domain service name name servers update params
func (o *PostDomainServiceNameNameServersUpdateParams) SetDomainNameServersUpdatePost(domainNameServersUpdatePost *models.PostDomainServiceNameNameServersUpdateParamsBody) {
	o.DomainNameServersUpdatePost = domainNameServersUpdatePost
}

// WithServiceName adds the serviceName to the post domain service name name servers update params
func (o *PostDomainServiceNameNameServersUpdateParams) WithServiceName(serviceName string) *PostDomainServiceNameNameServersUpdateParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post domain service name name servers update params
func (o *PostDomainServiceNameNameServersUpdateParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostDomainServiceNameNameServersUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DomainNameServersUpdatePost != nil {
		if err := r.SetBodyParam(o.DomainNameServersUpdatePost); err != nil {
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
