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

	"github.com/appscode/go-ovh/vps/models"
)

// NewPostVpsServiceNameCreateSnapshotParams creates a new PostVpsServiceNameCreateSnapshotParams object
// with the default values initialized.
func NewPostVpsServiceNameCreateSnapshotParams() *PostVpsServiceNameCreateSnapshotParams {
	var ()
	return &PostVpsServiceNameCreateSnapshotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostVpsServiceNameCreateSnapshotParamsWithTimeout creates a new PostVpsServiceNameCreateSnapshotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostVpsServiceNameCreateSnapshotParamsWithTimeout(timeout time.Duration) *PostVpsServiceNameCreateSnapshotParams {
	var ()
	return &PostVpsServiceNameCreateSnapshotParams{

		timeout: timeout,
	}
}

// NewPostVpsServiceNameCreateSnapshotParamsWithContext creates a new PostVpsServiceNameCreateSnapshotParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostVpsServiceNameCreateSnapshotParamsWithContext(ctx context.Context) *PostVpsServiceNameCreateSnapshotParams {
	var ()
	return &PostVpsServiceNameCreateSnapshotParams{

		Context: ctx,
	}
}

// NewPostVpsServiceNameCreateSnapshotParamsWithHTTPClient creates a new PostVpsServiceNameCreateSnapshotParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostVpsServiceNameCreateSnapshotParamsWithHTTPClient(client *http.Client) *PostVpsServiceNameCreateSnapshotParams {
	var ()
	return &PostVpsServiceNameCreateSnapshotParams{
		HTTPClient: client,
	}
}

/*PostVpsServiceNameCreateSnapshotParams contains all the parameters to send to the API endpoint
for the post vps service name create snapshot operation typically these are written to a http.Request
*/
type PostVpsServiceNameCreateSnapshotParams struct {

	/*ServiceName*/
	ServiceName string
	/*VpsCreateSnapshotPost*/
	VpsCreateSnapshotPost *models.PostVpsServiceNameCreateSnapshotParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post vps service name create snapshot params
func (o *PostVpsServiceNameCreateSnapshotParams) WithTimeout(timeout time.Duration) *PostVpsServiceNameCreateSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post vps service name create snapshot params
func (o *PostVpsServiceNameCreateSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post vps service name create snapshot params
func (o *PostVpsServiceNameCreateSnapshotParams) WithContext(ctx context.Context) *PostVpsServiceNameCreateSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post vps service name create snapshot params
func (o *PostVpsServiceNameCreateSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post vps service name create snapshot params
func (o *PostVpsServiceNameCreateSnapshotParams) WithHTTPClient(client *http.Client) *PostVpsServiceNameCreateSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post vps service name create snapshot params
func (o *PostVpsServiceNameCreateSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceName adds the serviceName to the post vps service name create snapshot params
func (o *PostVpsServiceNameCreateSnapshotParams) WithServiceName(serviceName string) *PostVpsServiceNameCreateSnapshotParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post vps service name create snapshot params
func (o *PostVpsServiceNameCreateSnapshotParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WithVpsCreateSnapshotPost adds the vpsCreateSnapshotPost to the post vps service name create snapshot params
func (o *PostVpsServiceNameCreateSnapshotParams) WithVpsCreateSnapshotPost(vpsCreateSnapshotPost *models.PostVpsServiceNameCreateSnapshotParamsBody) *PostVpsServiceNameCreateSnapshotParams {
	o.SetVpsCreateSnapshotPost(vpsCreateSnapshotPost)
	return o
}

// SetVpsCreateSnapshotPost adds the vpsCreateSnapshotPost to the post vps service name create snapshot params
func (o *PostVpsServiceNameCreateSnapshotParams) SetVpsCreateSnapshotPost(vpsCreateSnapshotPost *models.PostVpsServiceNameCreateSnapshotParamsBody) {
	o.VpsCreateSnapshotPost = vpsCreateSnapshotPost
}

// WriteToRequest writes these params to a swagger request
func (o *PostVpsServiceNameCreateSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	if o.VpsCreateSnapshotPost != nil {
		if err := r.SetBodyParam(o.VpsCreateSnapshotPost); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}