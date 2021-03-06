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

// NewPostCloudServiceNamePcaPcaServiceNameTasksParams creates a new PostCloudServiceNamePcaPcaServiceNameTasksParams object
// with the default values initialized.
func NewPostCloudServiceNamePcaPcaServiceNameTasksParams() *PostCloudServiceNamePcaPcaServiceNameTasksParams {
	var ()
	return &PostCloudServiceNamePcaPcaServiceNameTasksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCloudServiceNamePcaPcaServiceNameTasksParamsWithTimeout creates a new PostCloudServiceNamePcaPcaServiceNameTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCloudServiceNamePcaPcaServiceNameTasksParamsWithTimeout(timeout time.Duration) *PostCloudServiceNamePcaPcaServiceNameTasksParams {
	var ()
	return &PostCloudServiceNamePcaPcaServiceNameTasksParams{

		timeout: timeout,
	}
}

// NewPostCloudServiceNamePcaPcaServiceNameTasksParamsWithContext creates a new PostCloudServiceNamePcaPcaServiceNameTasksParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCloudServiceNamePcaPcaServiceNameTasksParamsWithContext(ctx context.Context) *PostCloudServiceNamePcaPcaServiceNameTasksParams {
	var ()
	return &PostCloudServiceNamePcaPcaServiceNameTasksParams{

		Context: ctx,
	}
}

// NewPostCloudServiceNamePcaPcaServiceNameTasksParamsWithHTTPClient creates a new PostCloudServiceNamePcaPcaServiceNameTasksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCloudServiceNamePcaPcaServiceNameTasksParamsWithHTTPClient(client *http.Client) *PostCloudServiceNamePcaPcaServiceNameTasksParams {
	var ()
	return &PostCloudServiceNamePcaPcaServiceNameTasksParams{
		HTTPClient: client,
	}
}

/*PostCloudServiceNamePcaPcaServiceNameTasksParams contains all the parameters to send to the API endpoint
for the post cloud service name pca pca service name tasks operation typically these are written to a http.Request
*/
type PostCloudServiceNamePcaPcaServiceNameTasksParams struct {

	/*CloudPcaTasksPost*/
	CloudPcaTasksPost *models.PostCloudServiceNamePcaPcaServiceNameTasksParamsBody
	/*PcaServiceName*/
	PcaServiceName string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post cloud service name pca pca service name tasks params
func (o *PostCloudServiceNamePcaPcaServiceNameTasksParams) WithTimeout(timeout time.Duration) *PostCloudServiceNamePcaPcaServiceNameTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cloud service name pca pca service name tasks params
func (o *PostCloudServiceNamePcaPcaServiceNameTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cloud service name pca pca service name tasks params
func (o *PostCloudServiceNamePcaPcaServiceNameTasksParams) WithContext(ctx context.Context) *PostCloudServiceNamePcaPcaServiceNameTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cloud service name pca pca service name tasks params
func (o *PostCloudServiceNamePcaPcaServiceNameTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cloud service name pca pca service name tasks params
func (o *PostCloudServiceNamePcaPcaServiceNameTasksParams) WithHTTPClient(client *http.Client) *PostCloudServiceNamePcaPcaServiceNameTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cloud service name pca pca service name tasks params
func (o *PostCloudServiceNamePcaPcaServiceNameTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudPcaTasksPost adds the cloudPcaTasksPost to the post cloud service name pca pca service name tasks params
func (o *PostCloudServiceNamePcaPcaServiceNameTasksParams) WithCloudPcaTasksPost(cloudPcaTasksPost *models.PostCloudServiceNamePcaPcaServiceNameTasksParamsBody) *PostCloudServiceNamePcaPcaServiceNameTasksParams {
	o.SetCloudPcaTasksPost(cloudPcaTasksPost)
	return o
}

// SetCloudPcaTasksPost adds the cloudPcaTasksPost to the post cloud service name pca pca service name tasks params
func (o *PostCloudServiceNamePcaPcaServiceNameTasksParams) SetCloudPcaTasksPost(cloudPcaTasksPost *models.PostCloudServiceNamePcaPcaServiceNameTasksParamsBody) {
	o.CloudPcaTasksPost = cloudPcaTasksPost
}

// WithPcaServiceName adds the pcaServiceName to the post cloud service name pca pca service name tasks params
func (o *PostCloudServiceNamePcaPcaServiceNameTasksParams) WithPcaServiceName(pcaServiceName string) *PostCloudServiceNamePcaPcaServiceNameTasksParams {
	o.SetPcaServiceName(pcaServiceName)
	return o
}

// SetPcaServiceName adds the pcaServiceName to the post cloud service name pca pca service name tasks params
func (o *PostCloudServiceNamePcaPcaServiceNameTasksParams) SetPcaServiceName(pcaServiceName string) {
	o.PcaServiceName = pcaServiceName
}

// WithServiceName adds the serviceName to the post cloud service name pca pca service name tasks params
func (o *PostCloudServiceNamePcaPcaServiceNameTasksParams) WithServiceName(serviceName string) *PostCloudServiceNamePcaPcaServiceNameTasksParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post cloud service name pca pca service name tasks params
func (o *PostCloudServiceNamePcaPcaServiceNameTasksParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostCloudServiceNamePcaPcaServiceNameTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CloudPcaTasksPost != nil {
		if err := r.SetBodyParam(o.CloudPcaTasksPost); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
