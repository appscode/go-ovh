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

// NewPostCloudProjectServiceNameUserUserIDRegeneratePasswordParams creates a new PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams object
// with the default values initialized.
func NewPostCloudProjectServiceNameUserUserIDRegeneratePasswordParams() *PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams {
	var ()
	return &PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCloudProjectServiceNameUserUserIDRegeneratePasswordParamsWithTimeout creates a new PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCloudProjectServiceNameUserUserIDRegeneratePasswordParamsWithTimeout(timeout time.Duration) *PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams {
	var ()
	return &PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams{

		timeout: timeout,
	}
}

// NewPostCloudProjectServiceNameUserUserIDRegeneratePasswordParamsWithContext creates a new PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCloudProjectServiceNameUserUserIDRegeneratePasswordParamsWithContext(ctx context.Context) *PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams {
	var ()
	return &PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams{

		Context: ctx,
	}
}

// NewPostCloudProjectServiceNameUserUserIDRegeneratePasswordParamsWithHTTPClient creates a new PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCloudProjectServiceNameUserUserIDRegeneratePasswordParamsWithHTTPClient(client *http.Client) *PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams {
	var ()
	return &PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams{
		HTTPClient: client,
	}
}

/*PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams contains all the parameters to send to the API endpoint
for the post cloud project service name user user ID regenerate password operation typically these are written to a http.Request
*/
type PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams struct {

	/*ServiceName*/
	ServiceName string
	/*UserID*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post cloud project service name user user ID regenerate password params
func (o *PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams) WithTimeout(timeout time.Duration) *PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cloud project service name user user ID regenerate password params
func (o *PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cloud project service name user user ID regenerate password params
func (o *PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams) WithContext(ctx context.Context) *PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cloud project service name user user ID regenerate password params
func (o *PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cloud project service name user user ID regenerate password params
func (o *PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams) WithHTTPClient(client *http.Client) *PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cloud project service name user user ID regenerate password params
func (o *PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceName adds the serviceName to the post cloud project service name user user ID regenerate password params
func (o *PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams) WithServiceName(serviceName string) *PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post cloud project service name user user ID regenerate password params
func (o *PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WithUserID adds the userID to the post cloud project service name user user ID regenerate password params
func (o *PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams) WithUserID(userID int64) *PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the post cloud project service name user user ID regenerate password params
func (o *PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PostCloudProjectServiceNameUserUserIDRegeneratePasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
