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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams creates a new DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams object
// with the default values initialized.
func NewDeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams() *DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams {
	var ()
	return &DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParamsWithTimeout creates a new DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParamsWithTimeout(timeout time.Duration) *DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams {
	var ()
	return &DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams{

		timeout: timeout,
	}
}

// NewDeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParamsWithContext creates a new DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParamsWithContext(ctx context.Context) *DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams {
	var ()
	return &DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams{

		Context: ctx,
	}
}

// NewDeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParamsWithHTTPClient creates a new DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParamsWithHTTPClient(client *http.Client) *DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams {
	var ()
	return &DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams{
		HTTPClient: client,
	}
}

/*DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams contains all the parameters to send to the API endpoint
for the delete IP loadbalancing service name HTTP frontend frontend ID operation typically these are written to a http.Request
*/
type DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams struct {

	/*FrontendID*/
	FrontendID int64
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete IP loadbalancing service name HTTP frontend frontend ID params
func (o *DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) WithTimeout(timeout time.Duration) *DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete IP loadbalancing service name HTTP frontend frontend ID params
func (o *DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete IP loadbalancing service name HTTP frontend frontend ID params
func (o *DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) WithContext(ctx context.Context) *DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete IP loadbalancing service name HTTP frontend frontend ID params
func (o *DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete IP loadbalancing service name HTTP frontend frontend ID params
func (o *DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) WithHTTPClient(client *http.Client) *DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete IP loadbalancing service name HTTP frontend frontend ID params
func (o *DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrontendID adds the frontendID to the delete IP loadbalancing service name HTTP frontend frontend ID params
func (o *DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) WithFrontendID(frontendID int64) *DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams {
	o.SetFrontendID(frontendID)
	return o
}

// SetFrontendID adds the frontendId to the delete IP loadbalancing service name HTTP frontend frontend ID params
func (o *DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) SetFrontendID(frontendID int64) {
	o.FrontendID = frontendID
}

// WithServiceName adds the serviceName to the delete IP loadbalancing service name HTTP frontend frontend ID params
func (o *DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) WithServiceName(serviceName string) *DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the delete IP loadbalancing service name HTTP frontend frontend ID params
func (o *DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIPLoadbalancingServiceNameHTTPFrontendFrontendIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param frontendId
	if err := r.SetPathParam("frontendId", swag.FormatInt64(o.FrontendID)); err != nil {
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
