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

// NewPostVpsServiceNameTerminateParams creates a new PostVpsServiceNameTerminateParams object
// with the default values initialized.
func NewPostVpsServiceNameTerminateParams() *PostVpsServiceNameTerminateParams {
	var ()
	return &PostVpsServiceNameTerminateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostVpsServiceNameTerminateParamsWithTimeout creates a new PostVpsServiceNameTerminateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostVpsServiceNameTerminateParamsWithTimeout(timeout time.Duration) *PostVpsServiceNameTerminateParams {
	var ()
	return &PostVpsServiceNameTerminateParams{

		timeout: timeout,
	}
}

// NewPostVpsServiceNameTerminateParamsWithContext creates a new PostVpsServiceNameTerminateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostVpsServiceNameTerminateParamsWithContext(ctx context.Context) *PostVpsServiceNameTerminateParams {
	var ()
	return &PostVpsServiceNameTerminateParams{

		Context: ctx,
	}
}

// NewPostVpsServiceNameTerminateParamsWithHTTPClient creates a new PostVpsServiceNameTerminateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostVpsServiceNameTerminateParamsWithHTTPClient(client *http.Client) *PostVpsServiceNameTerminateParams {
	var ()
	return &PostVpsServiceNameTerminateParams{
		HTTPClient: client,
	}
}

/*PostVpsServiceNameTerminateParams contains all the parameters to send to the API endpoint
for the post vps service name terminate operation typically these are written to a http.Request
*/
type PostVpsServiceNameTerminateParams struct {

	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post vps service name terminate params
func (o *PostVpsServiceNameTerminateParams) WithTimeout(timeout time.Duration) *PostVpsServiceNameTerminateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post vps service name terminate params
func (o *PostVpsServiceNameTerminateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post vps service name terminate params
func (o *PostVpsServiceNameTerminateParams) WithContext(ctx context.Context) *PostVpsServiceNameTerminateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post vps service name terminate params
func (o *PostVpsServiceNameTerminateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post vps service name terminate params
func (o *PostVpsServiceNameTerminateParams) WithHTTPClient(client *http.Client) *PostVpsServiceNameTerminateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post vps service name terminate params
func (o *PostVpsServiceNameTerminateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceName adds the serviceName to the post vps service name terminate params
func (o *PostVpsServiceNameTerminateParams) WithServiceName(serviceName string) *PostVpsServiceNameTerminateParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post vps service name terminate params
func (o *PostVpsServiceNameTerminateParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostVpsServiceNameTerminateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
