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

// NewPutCloudProjectServiceNameInstanceInstanceIDParams creates a new PutCloudProjectServiceNameInstanceInstanceIDParams object
// with the default values initialized.
func NewPutCloudProjectServiceNameInstanceInstanceIDParams() *PutCloudProjectServiceNameInstanceInstanceIDParams {
	var ()
	return &PutCloudProjectServiceNameInstanceInstanceIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCloudProjectServiceNameInstanceInstanceIDParamsWithTimeout creates a new PutCloudProjectServiceNameInstanceInstanceIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCloudProjectServiceNameInstanceInstanceIDParamsWithTimeout(timeout time.Duration) *PutCloudProjectServiceNameInstanceInstanceIDParams {
	var ()
	return &PutCloudProjectServiceNameInstanceInstanceIDParams{

		timeout: timeout,
	}
}

// NewPutCloudProjectServiceNameInstanceInstanceIDParamsWithContext creates a new PutCloudProjectServiceNameInstanceInstanceIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCloudProjectServiceNameInstanceInstanceIDParamsWithContext(ctx context.Context) *PutCloudProjectServiceNameInstanceInstanceIDParams {
	var ()
	return &PutCloudProjectServiceNameInstanceInstanceIDParams{

		Context: ctx,
	}
}

// NewPutCloudProjectServiceNameInstanceInstanceIDParamsWithHTTPClient creates a new PutCloudProjectServiceNameInstanceInstanceIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCloudProjectServiceNameInstanceInstanceIDParamsWithHTTPClient(client *http.Client) *PutCloudProjectServiceNameInstanceInstanceIDParams {
	var ()
	return &PutCloudProjectServiceNameInstanceInstanceIDParams{
		HTTPClient: client,
	}
}

/*PutCloudProjectServiceNameInstanceInstanceIDParams contains all the parameters to send to the API endpoint
for the put cloud project service name instance instance ID operation typically these are written to a http.Request
*/
type PutCloudProjectServiceNameInstanceInstanceIDParams struct {

	/*CloudProjectInstancePut*/
	CloudProjectInstancePut *models.PutCloudProjectServiceNameInstanceInstanceIDParamsBody
	/*InstanceID*/
	InstanceID string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put cloud project service name instance instance ID params
func (o *PutCloudProjectServiceNameInstanceInstanceIDParams) WithTimeout(timeout time.Duration) *PutCloudProjectServiceNameInstanceInstanceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put cloud project service name instance instance ID params
func (o *PutCloudProjectServiceNameInstanceInstanceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put cloud project service name instance instance ID params
func (o *PutCloudProjectServiceNameInstanceInstanceIDParams) WithContext(ctx context.Context) *PutCloudProjectServiceNameInstanceInstanceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put cloud project service name instance instance ID params
func (o *PutCloudProjectServiceNameInstanceInstanceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put cloud project service name instance instance ID params
func (o *PutCloudProjectServiceNameInstanceInstanceIDParams) WithHTTPClient(client *http.Client) *PutCloudProjectServiceNameInstanceInstanceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put cloud project service name instance instance ID params
func (o *PutCloudProjectServiceNameInstanceInstanceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudProjectInstancePut adds the cloudProjectInstancePut to the put cloud project service name instance instance ID params
func (o *PutCloudProjectServiceNameInstanceInstanceIDParams) WithCloudProjectInstancePut(cloudProjectInstancePut *models.PutCloudProjectServiceNameInstanceInstanceIDParamsBody) *PutCloudProjectServiceNameInstanceInstanceIDParams {
	o.SetCloudProjectInstancePut(cloudProjectInstancePut)
	return o
}

// SetCloudProjectInstancePut adds the cloudProjectInstancePut to the put cloud project service name instance instance ID params
func (o *PutCloudProjectServiceNameInstanceInstanceIDParams) SetCloudProjectInstancePut(cloudProjectInstancePut *models.PutCloudProjectServiceNameInstanceInstanceIDParamsBody) {
	o.CloudProjectInstancePut = cloudProjectInstancePut
}

// WithInstanceID adds the instanceID to the put cloud project service name instance instance ID params
func (o *PutCloudProjectServiceNameInstanceInstanceIDParams) WithInstanceID(instanceID string) *PutCloudProjectServiceNameInstanceInstanceIDParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the put cloud project service name instance instance ID params
func (o *PutCloudProjectServiceNameInstanceInstanceIDParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithServiceName adds the serviceName to the put cloud project service name instance instance ID params
func (o *PutCloudProjectServiceNameInstanceInstanceIDParams) WithServiceName(serviceName string) *PutCloudProjectServiceNameInstanceInstanceIDParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the put cloud project service name instance instance ID params
func (o *PutCloudProjectServiceNameInstanceInstanceIDParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PutCloudProjectServiceNameInstanceInstanceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CloudProjectInstancePut != nil {
		if err := r.SetBodyParam(o.CloudProjectInstancePut); err != nil {
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
