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

// NewPostIPLoadbalancingServiceNameConfirmTerminationParams creates a new PostIPLoadbalancingServiceNameConfirmTerminationParams object
// with the default values initialized.
func NewPostIPLoadbalancingServiceNameConfirmTerminationParams() *PostIPLoadbalancingServiceNameConfirmTerminationParams {
	var ()
	return &PostIPLoadbalancingServiceNameConfirmTerminationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIPLoadbalancingServiceNameConfirmTerminationParamsWithTimeout creates a new PostIPLoadbalancingServiceNameConfirmTerminationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIPLoadbalancingServiceNameConfirmTerminationParamsWithTimeout(timeout time.Duration) *PostIPLoadbalancingServiceNameConfirmTerminationParams {
	var ()
	return &PostIPLoadbalancingServiceNameConfirmTerminationParams{

		timeout: timeout,
	}
}

// NewPostIPLoadbalancingServiceNameConfirmTerminationParamsWithContext creates a new PostIPLoadbalancingServiceNameConfirmTerminationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIPLoadbalancingServiceNameConfirmTerminationParamsWithContext(ctx context.Context) *PostIPLoadbalancingServiceNameConfirmTerminationParams {
	var ()
	return &PostIPLoadbalancingServiceNameConfirmTerminationParams{

		Context: ctx,
	}
}

// NewPostIPLoadbalancingServiceNameConfirmTerminationParamsWithHTTPClient creates a new PostIPLoadbalancingServiceNameConfirmTerminationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIPLoadbalancingServiceNameConfirmTerminationParamsWithHTTPClient(client *http.Client) *PostIPLoadbalancingServiceNameConfirmTerminationParams {
	var ()
	return &PostIPLoadbalancingServiceNameConfirmTerminationParams{
		HTTPClient: client,
	}
}

/*PostIPLoadbalancingServiceNameConfirmTerminationParams contains all the parameters to send to the API endpoint
for the post IP loadbalancing service name confirm termination operation typically these are written to a http.Request
*/
type PostIPLoadbalancingServiceNameConfirmTerminationParams struct {

	/*IPLBConfirmTerminationPost*/
	IPLBConfirmTerminationPost *models.PostIPLoadbalancingServiceNameConfirmTerminationParamsBody
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post IP loadbalancing service name confirm termination params
func (o *PostIPLoadbalancingServiceNameConfirmTerminationParams) WithTimeout(timeout time.Duration) *PostIPLoadbalancingServiceNameConfirmTerminationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post IP loadbalancing service name confirm termination params
func (o *PostIPLoadbalancingServiceNameConfirmTerminationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post IP loadbalancing service name confirm termination params
func (o *PostIPLoadbalancingServiceNameConfirmTerminationParams) WithContext(ctx context.Context) *PostIPLoadbalancingServiceNameConfirmTerminationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post IP loadbalancing service name confirm termination params
func (o *PostIPLoadbalancingServiceNameConfirmTerminationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post IP loadbalancing service name confirm termination params
func (o *PostIPLoadbalancingServiceNameConfirmTerminationParams) WithHTTPClient(client *http.Client) *PostIPLoadbalancingServiceNameConfirmTerminationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post IP loadbalancing service name confirm termination params
func (o *PostIPLoadbalancingServiceNameConfirmTerminationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIPLBConfirmTerminationPost adds the iPLBConfirmTerminationPost to the post IP loadbalancing service name confirm termination params
func (o *PostIPLoadbalancingServiceNameConfirmTerminationParams) WithIPLBConfirmTerminationPost(iPLBConfirmTerminationPost *models.PostIPLoadbalancingServiceNameConfirmTerminationParamsBody) *PostIPLoadbalancingServiceNameConfirmTerminationParams {
	o.SetIPLBConfirmTerminationPost(iPLBConfirmTerminationPost)
	return o
}

// SetIPLBConfirmTerminationPost adds the iplbConfirmTerminationPost to the post IP loadbalancing service name confirm termination params
func (o *PostIPLoadbalancingServiceNameConfirmTerminationParams) SetIPLBConfirmTerminationPost(iPLBConfirmTerminationPost *models.PostIPLoadbalancingServiceNameConfirmTerminationParamsBody) {
	o.IPLBConfirmTerminationPost = iPLBConfirmTerminationPost
}

// WithServiceName adds the serviceName to the post IP loadbalancing service name confirm termination params
func (o *PostIPLoadbalancingServiceNameConfirmTerminationParams) WithServiceName(serviceName string) *PostIPLoadbalancingServiceNameConfirmTerminationParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post IP loadbalancing service name confirm termination params
func (o *PostIPLoadbalancingServiceNameConfirmTerminationParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostIPLoadbalancingServiceNameConfirmTerminationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IPLBConfirmTerminationPost != nil {
		if err := r.SetBodyParam(o.IPLBConfirmTerminationPost); err != nil {
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