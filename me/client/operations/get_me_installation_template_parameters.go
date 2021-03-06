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

// NewGetMeInstallationTemplateParams creates a new GetMeInstallationTemplateParams object
// with the default values initialized.
func NewGetMeInstallationTemplateParams() *GetMeInstallationTemplateParams {

	return &GetMeInstallationTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeInstallationTemplateParamsWithTimeout creates a new GetMeInstallationTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeInstallationTemplateParamsWithTimeout(timeout time.Duration) *GetMeInstallationTemplateParams {

	return &GetMeInstallationTemplateParams{

		timeout: timeout,
	}
}

// NewGetMeInstallationTemplateParamsWithContext creates a new GetMeInstallationTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeInstallationTemplateParamsWithContext(ctx context.Context) *GetMeInstallationTemplateParams {

	return &GetMeInstallationTemplateParams{

		Context: ctx,
	}
}

// NewGetMeInstallationTemplateParamsWithHTTPClient creates a new GetMeInstallationTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeInstallationTemplateParamsWithHTTPClient(client *http.Client) *GetMeInstallationTemplateParams {

	return &GetMeInstallationTemplateParams{
		HTTPClient: client,
	}
}

/*GetMeInstallationTemplateParams contains all the parameters to send to the API endpoint
for the get me installation template operation typically these are written to a http.Request
*/
type GetMeInstallationTemplateParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me installation template params
func (o *GetMeInstallationTemplateParams) WithTimeout(timeout time.Duration) *GetMeInstallationTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me installation template params
func (o *GetMeInstallationTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me installation template params
func (o *GetMeInstallationTemplateParams) WithContext(ctx context.Context) *GetMeInstallationTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me installation template params
func (o *GetMeInstallationTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me installation template params
func (o *GetMeInstallationTemplateParams) WithHTTPClient(client *http.Client) *GetMeInstallationTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me installation template params
func (o *GetMeInstallationTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeInstallationTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
