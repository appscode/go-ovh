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

// NewPostCloudProjectServiceNameInstanceInstanceIDSnapshotParams creates a new PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams object
// with the default values initialized.
func NewPostCloudProjectServiceNameInstanceInstanceIDSnapshotParams() *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams {
	var ()
	return &PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCloudProjectServiceNameInstanceInstanceIDSnapshotParamsWithTimeout creates a new PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCloudProjectServiceNameInstanceInstanceIDSnapshotParamsWithTimeout(timeout time.Duration) *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams {
	var ()
	return &PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams{

		timeout: timeout,
	}
}

// NewPostCloudProjectServiceNameInstanceInstanceIDSnapshotParamsWithContext creates a new PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCloudProjectServiceNameInstanceInstanceIDSnapshotParamsWithContext(ctx context.Context) *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams {
	var ()
	return &PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams{

		Context: ctx,
	}
}

// NewPostCloudProjectServiceNameInstanceInstanceIDSnapshotParamsWithHTTPClient creates a new PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCloudProjectServiceNameInstanceInstanceIDSnapshotParamsWithHTTPClient(client *http.Client) *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams {
	var ()
	return &PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams{
		HTTPClient: client,
	}
}

/*PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams contains all the parameters to send to the API endpoint
for the post cloud project service name instance instance ID snapshot operation typically these are written to a http.Request
*/
type PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams struct {

	/*CloudProjectInstanceSnapshotPost*/
	CloudProjectInstanceSnapshotPost *models.PostCloudProjectServiceNameInstanceInstanceIDSnapshotParamsBody
	/*InstanceID*/
	InstanceID string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post cloud project service name instance instance ID snapshot params
func (o *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams) WithTimeout(timeout time.Duration) *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cloud project service name instance instance ID snapshot params
func (o *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cloud project service name instance instance ID snapshot params
func (o *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams) WithContext(ctx context.Context) *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cloud project service name instance instance ID snapshot params
func (o *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cloud project service name instance instance ID snapshot params
func (o *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams) WithHTTPClient(client *http.Client) *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cloud project service name instance instance ID snapshot params
func (o *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudProjectInstanceSnapshotPost adds the cloudProjectInstanceSnapshotPost to the post cloud project service name instance instance ID snapshot params
func (o *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams) WithCloudProjectInstanceSnapshotPost(cloudProjectInstanceSnapshotPost *models.PostCloudProjectServiceNameInstanceInstanceIDSnapshotParamsBody) *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams {
	o.SetCloudProjectInstanceSnapshotPost(cloudProjectInstanceSnapshotPost)
	return o
}

// SetCloudProjectInstanceSnapshotPost adds the cloudProjectInstanceSnapshotPost to the post cloud project service name instance instance ID snapshot params
func (o *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams) SetCloudProjectInstanceSnapshotPost(cloudProjectInstanceSnapshotPost *models.PostCloudProjectServiceNameInstanceInstanceIDSnapshotParamsBody) {
	o.CloudProjectInstanceSnapshotPost = cloudProjectInstanceSnapshotPost
}

// WithInstanceID adds the instanceID to the post cloud project service name instance instance ID snapshot params
func (o *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams) WithInstanceID(instanceID string) *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the post cloud project service name instance instance ID snapshot params
func (o *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithServiceName adds the serviceName to the post cloud project service name instance instance ID snapshot params
func (o *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams) WithServiceName(serviceName string) *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post cloud project service name instance instance ID snapshot params
func (o *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostCloudProjectServiceNameInstanceInstanceIDSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CloudProjectInstanceSnapshotPost != nil {
		if err := r.SetBodyParam(o.CloudProjectInstanceSnapshotPost); err != nil {
			return err
		}
	}

	// path param instanceId
	if err := r.SetPathParam("instanceId", o.InstanceID); err != nil {
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