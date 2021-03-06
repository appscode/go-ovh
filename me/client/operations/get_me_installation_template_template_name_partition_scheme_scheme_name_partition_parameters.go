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

// NewGetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams creates a new GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams object
// with the default values initialized.
func NewGetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams() *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams {
	var ()
	return &GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParamsWithTimeout creates a new GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParamsWithTimeout(timeout time.Duration) *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams {
	var ()
	return &GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams{

		timeout: timeout,
	}
}

// NewGetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParamsWithContext creates a new GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParamsWithContext(ctx context.Context) *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams {
	var ()
	return &GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams{

		Context: ctx,
	}
}

// NewGetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParamsWithHTTPClient creates a new GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParamsWithHTTPClient(client *http.Client) *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams {
	var ()
	return &GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams{
		HTTPClient: client,
	}
}

/*GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams contains all the parameters to send to the API endpoint
for the get me installation template template name partition scheme scheme name partition operation typically these are written to a http.Request
*/
type GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams struct {

	/*SchemeName*/
	SchemeName string
	/*TemplateName*/
	TemplateName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me installation template template name partition scheme scheme name partition params
func (o *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams) WithTimeout(timeout time.Duration) *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me installation template template name partition scheme scheme name partition params
func (o *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me installation template template name partition scheme scheme name partition params
func (o *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams) WithContext(ctx context.Context) *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me installation template template name partition scheme scheme name partition params
func (o *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me installation template template name partition scheme scheme name partition params
func (o *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams) WithHTTPClient(client *http.Client) *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me installation template template name partition scheme scheme name partition params
func (o *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSchemeName adds the schemeName to the get me installation template template name partition scheme scheme name partition params
func (o *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams) WithSchemeName(schemeName string) *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams {
	o.SetSchemeName(schemeName)
	return o
}

// SetSchemeName adds the schemeName to the get me installation template template name partition scheme scheme name partition params
func (o *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams) SetSchemeName(schemeName string) {
	o.SchemeName = schemeName
}

// WithTemplateName adds the templateName to the get me installation template template name partition scheme scheme name partition params
func (o *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams) WithTemplateName(templateName string) *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams {
	o.SetTemplateName(templateName)
	return o
}

// SetTemplateName adds the templateName to the get me installation template template name partition scheme scheme name partition params
func (o *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams) SetTemplateName(templateName string) {
	o.TemplateName = templateName
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param schemeName
	if err := r.SetPathParam("schemeName", o.SchemeName); err != nil {
		return err
	}

	// path param templateName
	if err := r.SetPathParam("templateName", o.TemplateName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
