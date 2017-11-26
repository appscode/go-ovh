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

// NewDeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams creates a new DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams object
// with the default values initialized.
func NewDeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams() *DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams {
	var ()
	return &DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParamsWithTimeout creates a new DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParamsWithTimeout(timeout time.Duration) *DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams {
	var ()
	return &DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams{

		timeout: timeout,
	}
}

// NewDeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParamsWithContext creates a new DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParamsWithContext(ctx context.Context) *DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams {
	var ()
	return &DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams{

		Context: ctx,
	}
}

// NewDeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParamsWithHTTPClient creates a new DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParamsWithHTTPClient(client *http.Client) *DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams {
	var ()
	return &DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams{
		HTTPClient: client,
	}
}

/*DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams contains all the parameters to send to the API endpoint
for the delete me installation template template name partition scheme scheme name operation typically these are written to a http.Request
*/
type DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams struct {

	/*SchemeName*/
	SchemeName string
	/*TemplateName*/
	TemplateName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete me installation template template name partition scheme scheme name params
func (o *DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams) WithTimeout(timeout time.Duration) *DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete me installation template template name partition scheme scheme name params
func (o *DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete me installation template template name partition scheme scheme name params
func (o *DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams) WithContext(ctx context.Context) *DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete me installation template template name partition scheme scheme name params
func (o *DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete me installation template template name partition scheme scheme name params
func (o *DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams) WithHTTPClient(client *http.Client) *DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete me installation template template name partition scheme scheme name params
func (o *DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSchemeName adds the schemeName to the delete me installation template template name partition scheme scheme name params
func (o *DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams) WithSchemeName(schemeName string) *DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams {
	o.SetSchemeName(schemeName)
	return o
}

// SetSchemeName adds the schemeName to the delete me installation template template name partition scheme scheme name params
func (o *DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams) SetSchemeName(schemeName string) {
	o.SchemeName = schemeName
}

// WithTemplateName adds the templateName to the delete me installation template template name partition scheme scheme name params
func (o *DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams) WithTemplateName(templateName string) *DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams {
	o.SetTemplateName(templateName)
	return o
}

// SetTemplateName adds the templateName to the delete me installation template template name partition scheme scheme name params
func (o *DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams) SetTemplateName(templateName string) {
	o.TemplateName = templateName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
