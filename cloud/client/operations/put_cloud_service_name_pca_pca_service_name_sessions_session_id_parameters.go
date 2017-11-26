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

// NewPutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams creates a new PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams object
// with the default values initialized.
func NewPutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams() *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams {
	var ()
	return &PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParamsWithTimeout creates a new PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParamsWithTimeout(timeout time.Duration) *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams {
	var ()
	return &PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams{

		timeout: timeout,
	}
}

// NewPutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParamsWithContext creates a new PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParamsWithContext(ctx context.Context) *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams {
	var ()
	return &PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams{

		Context: ctx,
	}
}

// NewPutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParamsWithHTTPClient creates a new PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParamsWithHTTPClient(client *http.Client) *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams {
	var ()
	return &PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams{
		HTTPClient: client,
	}
}

/*PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams contains all the parameters to send to the API endpoint
for the put cloud service name pca pca service name sessions session ID operation typically these are written to a http.Request
*/
type PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams struct {

	/*CloudPcaSessionsPut*/
	CloudPcaSessionsPut *models.PcaSession
	/*PcaServiceName*/
	PcaServiceName string
	/*ServiceName*/
	ServiceName string
	/*SessionID*/
	SessionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put cloud service name pca pca service name sessions session ID params
func (o *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams) WithTimeout(timeout time.Duration) *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put cloud service name pca pca service name sessions session ID params
func (o *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put cloud service name pca pca service name sessions session ID params
func (o *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams) WithContext(ctx context.Context) *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put cloud service name pca pca service name sessions session ID params
func (o *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put cloud service name pca pca service name sessions session ID params
func (o *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams) WithHTTPClient(client *http.Client) *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put cloud service name pca pca service name sessions session ID params
func (o *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudPcaSessionsPut adds the cloudPcaSessionsPut to the put cloud service name pca pca service name sessions session ID params
func (o *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams) WithCloudPcaSessionsPut(cloudPcaSessionsPut *models.PcaSession) *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams {
	o.SetCloudPcaSessionsPut(cloudPcaSessionsPut)
	return o
}

// SetCloudPcaSessionsPut adds the cloudPcaSessionsPut to the put cloud service name pca pca service name sessions session ID params
func (o *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams) SetCloudPcaSessionsPut(cloudPcaSessionsPut *models.PcaSession) {
	o.CloudPcaSessionsPut = cloudPcaSessionsPut
}

// WithPcaServiceName adds the pcaServiceName to the put cloud service name pca pca service name sessions session ID params
func (o *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams) WithPcaServiceName(pcaServiceName string) *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams {
	o.SetPcaServiceName(pcaServiceName)
	return o
}

// SetPcaServiceName adds the pcaServiceName to the put cloud service name pca pca service name sessions session ID params
func (o *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams) SetPcaServiceName(pcaServiceName string) {
	o.PcaServiceName = pcaServiceName
}

// WithServiceName adds the serviceName to the put cloud service name pca pca service name sessions session ID params
func (o *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams) WithServiceName(serviceName string) *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the put cloud service name pca pca service name sessions session ID params
func (o *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WithSessionID adds the sessionID to the put cloud service name pca pca service name sessions session ID params
func (o *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams) WithSessionID(sessionID string) *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams {
	o.SetSessionID(sessionID)
	return o
}

// SetSessionID adds the sessionId to the put cloud service name pca pca service name sessions session ID params
func (o *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams) SetSessionID(sessionID string) {
	o.SessionID = sessionID
}

// WriteToRequest writes these params to a swagger request
func (o *PutCloudServiceNamePcaPcaServiceNameSessionsSessionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CloudPcaSessionsPut != nil {
		if err := r.SetBodyParam(o.CloudPcaSessionsPut); err != nil {
			return err
		}
	}

	// path param pcaServiceName
	if err := r.SetPathParam("pcaServiceName", o.PcaServiceName); err != nil {
		return err
	}

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	// path param sessionId
	if err := r.SetPathParam("sessionId", o.SessionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}