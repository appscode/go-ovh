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

// NewPutCloudProjectServiceNameVolumeVolumeIDParams creates a new PutCloudProjectServiceNameVolumeVolumeIDParams object
// with the default values initialized.
func NewPutCloudProjectServiceNameVolumeVolumeIDParams() *PutCloudProjectServiceNameVolumeVolumeIDParams {
	var ()
	return &PutCloudProjectServiceNameVolumeVolumeIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCloudProjectServiceNameVolumeVolumeIDParamsWithTimeout creates a new PutCloudProjectServiceNameVolumeVolumeIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCloudProjectServiceNameVolumeVolumeIDParamsWithTimeout(timeout time.Duration) *PutCloudProjectServiceNameVolumeVolumeIDParams {
	var ()
	return &PutCloudProjectServiceNameVolumeVolumeIDParams{

		timeout: timeout,
	}
}

// NewPutCloudProjectServiceNameVolumeVolumeIDParamsWithContext creates a new PutCloudProjectServiceNameVolumeVolumeIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCloudProjectServiceNameVolumeVolumeIDParamsWithContext(ctx context.Context) *PutCloudProjectServiceNameVolumeVolumeIDParams {
	var ()
	return &PutCloudProjectServiceNameVolumeVolumeIDParams{

		Context: ctx,
	}
}

// NewPutCloudProjectServiceNameVolumeVolumeIDParamsWithHTTPClient creates a new PutCloudProjectServiceNameVolumeVolumeIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCloudProjectServiceNameVolumeVolumeIDParamsWithHTTPClient(client *http.Client) *PutCloudProjectServiceNameVolumeVolumeIDParams {
	var ()
	return &PutCloudProjectServiceNameVolumeVolumeIDParams{
		HTTPClient: client,
	}
}

/*PutCloudProjectServiceNameVolumeVolumeIDParams contains all the parameters to send to the API endpoint
for the put cloud project service name volume volume ID operation typically these are written to a http.Request
*/
type PutCloudProjectServiceNameVolumeVolumeIDParams struct {

	/*CloudProjectVolumePut*/
	CloudProjectVolumePut *models.PutCloudProjectServiceNameVolumeVolumeIDParamsBody
	/*ServiceName*/
	ServiceName string
	/*VolumeID*/
	VolumeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put cloud project service name volume volume ID params
func (o *PutCloudProjectServiceNameVolumeVolumeIDParams) WithTimeout(timeout time.Duration) *PutCloudProjectServiceNameVolumeVolumeIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put cloud project service name volume volume ID params
func (o *PutCloudProjectServiceNameVolumeVolumeIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put cloud project service name volume volume ID params
func (o *PutCloudProjectServiceNameVolumeVolumeIDParams) WithContext(ctx context.Context) *PutCloudProjectServiceNameVolumeVolumeIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put cloud project service name volume volume ID params
func (o *PutCloudProjectServiceNameVolumeVolumeIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put cloud project service name volume volume ID params
func (o *PutCloudProjectServiceNameVolumeVolumeIDParams) WithHTTPClient(client *http.Client) *PutCloudProjectServiceNameVolumeVolumeIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put cloud project service name volume volume ID params
func (o *PutCloudProjectServiceNameVolumeVolumeIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudProjectVolumePut adds the cloudProjectVolumePut to the put cloud project service name volume volume ID params
func (o *PutCloudProjectServiceNameVolumeVolumeIDParams) WithCloudProjectVolumePut(cloudProjectVolumePut *models.PutCloudProjectServiceNameVolumeVolumeIDParamsBody) *PutCloudProjectServiceNameVolumeVolumeIDParams {
	o.SetCloudProjectVolumePut(cloudProjectVolumePut)
	return o
}

// SetCloudProjectVolumePut adds the cloudProjectVolumePut to the put cloud project service name volume volume ID params
func (o *PutCloudProjectServiceNameVolumeVolumeIDParams) SetCloudProjectVolumePut(cloudProjectVolumePut *models.PutCloudProjectServiceNameVolumeVolumeIDParamsBody) {
	o.CloudProjectVolumePut = cloudProjectVolumePut
}

// WithServiceName adds the serviceName to the put cloud project service name volume volume ID params
func (o *PutCloudProjectServiceNameVolumeVolumeIDParams) WithServiceName(serviceName string) *PutCloudProjectServiceNameVolumeVolumeIDParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the put cloud project service name volume volume ID params
func (o *PutCloudProjectServiceNameVolumeVolumeIDParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WithVolumeID adds the volumeID to the put cloud project service name volume volume ID params
func (o *PutCloudProjectServiceNameVolumeVolumeIDParams) WithVolumeID(volumeID string) *PutCloudProjectServiceNameVolumeVolumeIDParams {
	o.SetVolumeID(volumeID)
	return o
}

// SetVolumeID adds the volumeId to the put cloud project service name volume volume ID params
func (o *PutCloudProjectServiceNameVolumeVolumeIDParams) SetVolumeID(volumeID string) {
	o.VolumeID = volumeID
}

// WriteToRequest writes these params to a swagger request
func (o *PutCloudProjectServiceNameVolumeVolumeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CloudProjectVolumePut != nil {
		if err := r.SetBodyParam(o.CloudProjectVolumePut); err != nil {
			return err
		}
	}

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	// path param volumeId
	if err := r.SetPathParam("volumeId", o.VolumeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
