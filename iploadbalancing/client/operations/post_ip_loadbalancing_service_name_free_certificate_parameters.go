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

	"github.com/appscode/go-ovh/iploadbalancing/models"
)

// NewPostIPLoadbalancingServiceNameFreeCertificateParams creates a new PostIPLoadbalancingServiceNameFreeCertificateParams object
// with the default values initialized.
func NewPostIPLoadbalancingServiceNameFreeCertificateParams() *PostIPLoadbalancingServiceNameFreeCertificateParams {
	var ()
	return &PostIPLoadbalancingServiceNameFreeCertificateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIPLoadbalancingServiceNameFreeCertificateParamsWithTimeout creates a new PostIPLoadbalancingServiceNameFreeCertificateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIPLoadbalancingServiceNameFreeCertificateParamsWithTimeout(timeout time.Duration) *PostIPLoadbalancingServiceNameFreeCertificateParams {
	var ()
	return &PostIPLoadbalancingServiceNameFreeCertificateParams{

		timeout: timeout,
	}
}

// NewPostIPLoadbalancingServiceNameFreeCertificateParamsWithContext creates a new PostIPLoadbalancingServiceNameFreeCertificateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIPLoadbalancingServiceNameFreeCertificateParamsWithContext(ctx context.Context) *PostIPLoadbalancingServiceNameFreeCertificateParams {
	var ()
	return &PostIPLoadbalancingServiceNameFreeCertificateParams{

		Context: ctx,
	}
}

// NewPostIPLoadbalancingServiceNameFreeCertificateParamsWithHTTPClient creates a new PostIPLoadbalancingServiceNameFreeCertificateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIPLoadbalancingServiceNameFreeCertificateParamsWithHTTPClient(client *http.Client) *PostIPLoadbalancingServiceNameFreeCertificateParams {
	var ()
	return &PostIPLoadbalancingServiceNameFreeCertificateParams{
		HTTPClient: client,
	}
}

/*PostIPLoadbalancingServiceNameFreeCertificateParams contains all the parameters to send to the API endpoint
for the post IP loadbalancing service name free certificate operation typically these are written to a http.Request
*/
type PostIPLoadbalancingServiceNameFreeCertificateParams struct {

	/*IPLBFreeCertificatePost*/
	IPLBFreeCertificatePost *models.PostIPLoadbalancingServiceNameFreeCertificateParamsBody
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post IP loadbalancing service name free certificate params
func (o *PostIPLoadbalancingServiceNameFreeCertificateParams) WithTimeout(timeout time.Duration) *PostIPLoadbalancingServiceNameFreeCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post IP loadbalancing service name free certificate params
func (o *PostIPLoadbalancingServiceNameFreeCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post IP loadbalancing service name free certificate params
func (o *PostIPLoadbalancingServiceNameFreeCertificateParams) WithContext(ctx context.Context) *PostIPLoadbalancingServiceNameFreeCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post IP loadbalancing service name free certificate params
func (o *PostIPLoadbalancingServiceNameFreeCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post IP loadbalancing service name free certificate params
func (o *PostIPLoadbalancingServiceNameFreeCertificateParams) WithHTTPClient(client *http.Client) *PostIPLoadbalancingServiceNameFreeCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post IP loadbalancing service name free certificate params
func (o *PostIPLoadbalancingServiceNameFreeCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIPLBFreeCertificatePost adds the iPLBFreeCertificatePost to the post IP loadbalancing service name free certificate params
func (o *PostIPLoadbalancingServiceNameFreeCertificateParams) WithIPLBFreeCertificatePost(iPLBFreeCertificatePost *models.PostIPLoadbalancingServiceNameFreeCertificateParamsBody) *PostIPLoadbalancingServiceNameFreeCertificateParams {
	o.SetIPLBFreeCertificatePost(iPLBFreeCertificatePost)
	return o
}

// SetIPLBFreeCertificatePost adds the iplbFreeCertificatePost to the post IP loadbalancing service name free certificate params
func (o *PostIPLoadbalancingServiceNameFreeCertificateParams) SetIPLBFreeCertificatePost(iPLBFreeCertificatePost *models.PostIPLoadbalancingServiceNameFreeCertificateParamsBody) {
	o.IPLBFreeCertificatePost = iPLBFreeCertificatePost
}

// WithServiceName adds the serviceName to the post IP loadbalancing service name free certificate params
func (o *PostIPLoadbalancingServiceNameFreeCertificateParams) WithServiceName(serviceName string) *PostIPLoadbalancingServiceNameFreeCertificateParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post IP loadbalancing service name free certificate params
func (o *PostIPLoadbalancingServiceNameFreeCertificateParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostIPLoadbalancingServiceNameFreeCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IPLBFreeCertificatePost != nil {
		if err := r.SetBodyParam(o.IPLBFreeCertificatePost); err != nil {
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
