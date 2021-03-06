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

// NewDeleteCloudProjectServiceNameInstanceGroupGroupIDParams creates a new DeleteCloudProjectServiceNameInstanceGroupGroupIDParams object
// with the default values initialized.
func NewDeleteCloudProjectServiceNameInstanceGroupGroupIDParams() *DeleteCloudProjectServiceNameInstanceGroupGroupIDParams {
	var ()
	return &DeleteCloudProjectServiceNameInstanceGroupGroupIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCloudProjectServiceNameInstanceGroupGroupIDParamsWithTimeout creates a new DeleteCloudProjectServiceNameInstanceGroupGroupIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCloudProjectServiceNameInstanceGroupGroupIDParamsWithTimeout(timeout time.Duration) *DeleteCloudProjectServiceNameInstanceGroupGroupIDParams {
	var ()
	return &DeleteCloudProjectServiceNameInstanceGroupGroupIDParams{

		timeout: timeout,
	}
}

// NewDeleteCloudProjectServiceNameInstanceGroupGroupIDParamsWithContext creates a new DeleteCloudProjectServiceNameInstanceGroupGroupIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCloudProjectServiceNameInstanceGroupGroupIDParamsWithContext(ctx context.Context) *DeleteCloudProjectServiceNameInstanceGroupGroupIDParams {
	var ()
	return &DeleteCloudProjectServiceNameInstanceGroupGroupIDParams{

		Context: ctx,
	}
}

// NewDeleteCloudProjectServiceNameInstanceGroupGroupIDParamsWithHTTPClient creates a new DeleteCloudProjectServiceNameInstanceGroupGroupIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCloudProjectServiceNameInstanceGroupGroupIDParamsWithHTTPClient(client *http.Client) *DeleteCloudProjectServiceNameInstanceGroupGroupIDParams {
	var ()
	return &DeleteCloudProjectServiceNameInstanceGroupGroupIDParams{
		HTTPClient: client,
	}
}

/*DeleteCloudProjectServiceNameInstanceGroupGroupIDParams contains all the parameters to send to the API endpoint
for the delete cloud project service name instance group group ID operation typically these are written to a http.Request
*/
type DeleteCloudProjectServiceNameInstanceGroupGroupIDParams struct {

	/*GroupID*/
	GroupID string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cloud project service name instance group group ID params
func (o *DeleteCloudProjectServiceNameInstanceGroupGroupIDParams) WithTimeout(timeout time.Duration) *DeleteCloudProjectServiceNameInstanceGroupGroupIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cloud project service name instance group group ID params
func (o *DeleteCloudProjectServiceNameInstanceGroupGroupIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cloud project service name instance group group ID params
func (o *DeleteCloudProjectServiceNameInstanceGroupGroupIDParams) WithContext(ctx context.Context) *DeleteCloudProjectServiceNameInstanceGroupGroupIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cloud project service name instance group group ID params
func (o *DeleteCloudProjectServiceNameInstanceGroupGroupIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cloud project service name instance group group ID params
func (o *DeleteCloudProjectServiceNameInstanceGroupGroupIDParams) WithHTTPClient(client *http.Client) *DeleteCloudProjectServiceNameInstanceGroupGroupIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cloud project service name instance group group ID params
func (o *DeleteCloudProjectServiceNameInstanceGroupGroupIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the delete cloud project service name instance group group ID params
func (o *DeleteCloudProjectServiceNameInstanceGroupGroupIDParams) WithGroupID(groupID string) *DeleteCloudProjectServiceNameInstanceGroupGroupIDParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the delete cloud project service name instance group group ID params
func (o *DeleteCloudProjectServiceNameInstanceGroupGroupIDParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithServiceName adds the serviceName to the delete cloud project service name instance group group ID params
func (o *DeleteCloudProjectServiceNameInstanceGroupGroupIDParams) WithServiceName(serviceName string) *DeleteCloudProjectServiceNameInstanceGroupGroupIDParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the delete cloud project service name instance group group ID params
func (o *DeleteCloudProjectServiceNameInstanceGroupGroupIDParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCloudProjectServiceNameInstanceGroupGroupIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
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
