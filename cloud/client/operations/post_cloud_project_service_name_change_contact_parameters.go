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

	"github.com/appscode/go-ovh/cloud/models"
)

// NewPostCloudProjectServiceNameChangeContactParams creates a new PostCloudProjectServiceNameChangeContactParams object
// with the default values initialized.
func NewPostCloudProjectServiceNameChangeContactParams() *PostCloudProjectServiceNameChangeContactParams {
	var ()
	return &PostCloudProjectServiceNameChangeContactParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCloudProjectServiceNameChangeContactParamsWithTimeout creates a new PostCloudProjectServiceNameChangeContactParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCloudProjectServiceNameChangeContactParamsWithTimeout(timeout time.Duration) *PostCloudProjectServiceNameChangeContactParams {
	var ()
	return &PostCloudProjectServiceNameChangeContactParams{

		timeout: timeout,
	}
}

// NewPostCloudProjectServiceNameChangeContactParamsWithContext creates a new PostCloudProjectServiceNameChangeContactParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCloudProjectServiceNameChangeContactParamsWithContext(ctx context.Context) *PostCloudProjectServiceNameChangeContactParams {
	var ()
	return &PostCloudProjectServiceNameChangeContactParams{

		Context: ctx,
	}
}

// NewPostCloudProjectServiceNameChangeContactParamsWithHTTPClient creates a new PostCloudProjectServiceNameChangeContactParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCloudProjectServiceNameChangeContactParamsWithHTTPClient(client *http.Client) *PostCloudProjectServiceNameChangeContactParams {
	var ()
	return &PostCloudProjectServiceNameChangeContactParams{
		HTTPClient: client,
	}
}

/*PostCloudProjectServiceNameChangeContactParams contains all the parameters to send to the API endpoint
for the post cloud project service name change contact operation typically these are written to a http.Request
*/
type PostCloudProjectServiceNameChangeContactParams struct {

	/*CloudProjectChangeContactPost*/
	CloudProjectChangeContactPost *models.PostCloudProjectServiceNameChangeContactParamsBody
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post cloud project service name change contact params
func (o *PostCloudProjectServiceNameChangeContactParams) WithTimeout(timeout time.Duration) *PostCloudProjectServiceNameChangeContactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cloud project service name change contact params
func (o *PostCloudProjectServiceNameChangeContactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cloud project service name change contact params
func (o *PostCloudProjectServiceNameChangeContactParams) WithContext(ctx context.Context) *PostCloudProjectServiceNameChangeContactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cloud project service name change contact params
func (o *PostCloudProjectServiceNameChangeContactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cloud project service name change contact params
func (o *PostCloudProjectServiceNameChangeContactParams) WithHTTPClient(client *http.Client) *PostCloudProjectServiceNameChangeContactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cloud project service name change contact params
func (o *PostCloudProjectServiceNameChangeContactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudProjectChangeContactPost adds the cloudProjectChangeContactPost to the post cloud project service name change contact params
func (o *PostCloudProjectServiceNameChangeContactParams) WithCloudProjectChangeContactPost(cloudProjectChangeContactPost *models.PostCloudProjectServiceNameChangeContactParamsBody) *PostCloudProjectServiceNameChangeContactParams {
	o.SetCloudProjectChangeContactPost(cloudProjectChangeContactPost)
	return o
}

// SetCloudProjectChangeContactPost adds the cloudProjectChangeContactPost to the post cloud project service name change contact params
func (o *PostCloudProjectServiceNameChangeContactParams) SetCloudProjectChangeContactPost(cloudProjectChangeContactPost *models.PostCloudProjectServiceNameChangeContactParamsBody) {
	o.CloudProjectChangeContactPost = cloudProjectChangeContactPost
}

// WithServiceName adds the serviceName to the post cloud project service name change contact params
func (o *PostCloudProjectServiceNameChangeContactParams) WithServiceName(serviceName string) *PostCloudProjectServiceNameChangeContactParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post cloud project service name change contact params
func (o *PostCloudProjectServiceNameChangeContactParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostCloudProjectServiceNameChangeContactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CloudProjectChangeContactPost != nil {
		if err := r.SetBodyParam(o.CloudProjectChangeContactPost); err != nil {
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
