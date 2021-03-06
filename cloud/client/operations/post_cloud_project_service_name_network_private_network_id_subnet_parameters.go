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

// NewPostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams creates a new PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams object
// with the default values initialized.
func NewPostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams() *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams {
	var ()
	return &PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsWithTimeout creates a new PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsWithTimeout(timeout time.Duration) *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams {
	var ()
	return &PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams{

		timeout: timeout,
	}
}

// NewPostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsWithContext creates a new PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsWithContext(ctx context.Context) *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams {
	var ()
	return &PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams{

		Context: ctx,
	}
}

// NewPostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsWithHTTPClient creates a new PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsWithHTTPClient(client *http.Client) *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams {
	var ()
	return &PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams{
		HTTPClient: client,
	}
}

/*PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams contains all the parameters to send to the API endpoint
for the post cloud project service name network private network ID subnet operation typically these are written to a http.Request
*/
type PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams struct {

	/*CloudProjectNetworkPrivateSubnetPost*/
	CloudProjectNetworkPrivateSubnetPost *models.PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsBody
	/*NetworkID*/
	NetworkID string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post cloud project service name network private network ID subnet params
func (o *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams) WithTimeout(timeout time.Duration) *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cloud project service name network private network ID subnet params
func (o *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cloud project service name network private network ID subnet params
func (o *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams) WithContext(ctx context.Context) *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cloud project service name network private network ID subnet params
func (o *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cloud project service name network private network ID subnet params
func (o *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams) WithHTTPClient(client *http.Client) *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cloud project service name network private network ID subnet params
func (o *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudProjectNetworkPrivateSubnetPost adds the cloudProjectNetworkPrivateSubnetPost to the post cloud project service name network private network ID subnet params
func (o *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams) WithCloudProjectNetworkPrivateSubnetPost(cloudProjectNetworkPrivateSubnetPost *models.PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsBody) *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams {
	o.SetCloudProjectNetworkPrivateSubnetPost(cloudProjectNetworkPrivateSubnetPost)
	return o
}

// SetCloudProjectNetworkPrivateSubnetPost adds the cloudProjectNetworkPrivateSubnetPost to the post cloud project service name network private network ID subnet params
func (o *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams) SetCloudProjectNetworkPrivateSubnetPost(cloudProjectNetworkPrivateSubnetPost *models.PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsBody) {
	o.CloudProjectNetworkPrivateSubnetPost = cloudProjectNetworkPrivateSubnetPost
}

// WithNetworkID adds the networkID to the post cloud project service name network private network ID subnet params
func (o *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams) WithNetworkID(networkID string) *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post cloud project service name network private network ID subnet params
func (o *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithServiceName adds the serviceName to the post cloud project service name network private network ID subnet params
func (o *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams) WithServiceName(serviceName string) *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post cloud project service name network private network ID subnet params
func (o *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CloudProjectNetworkPrivateSubnetPost != nil {
		if err := r.SetBodyParam(o.CloudProjectNetworkPrivateSubnetPost); err != nil {
			return err
		}
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
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
