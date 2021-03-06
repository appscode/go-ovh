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

	"github.com/appscode/go-ovh/me/models"
)

// NewPostMeInstallationTemplateParams creates a new PostMeInstallationTemplateParams object
// with the default values initialized.
func NewPostMeInstallationTemplateParams() *PostMeInstallationTemplateParams {
	var ()
	return &PostMeInstallationTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeInstallationTemplateParamsWithTimeout creates a new PostMeInstallationTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeInstallationTemplateParamsWithTimeout(timeout time.Duration) *PostMeInstallationTemplateParams {
	var ()
	return &PostMeInstallationTemplateParams{

		timeout: timeout,
	}
}

// NewPostMeInstallationTemplateParamsWithContext creates a new PostMeInstallationTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeInstallationTemplateParamsWithContext(ctx context.Context) *PostMeInstallationTemplateParams {
	var ()
	return &PostMeInstallationTemplateParams{

		Context: ctx,
	}
}

// NewPostMeInstallationTemplateParamsWithHTTPClient creates a new PostMeInstallationTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeInstallationTemplateParamsWithHTTPClient(client *http.Client) *PostMeInstallationTemplateParams {
	var ()
	return &PostMeInstallationTemplateParams{
		HTTPClient: client,
	}
}

/*PostMeInstallationTemplateParams contains all the parameters to send to the API endpoint
for the post me installation template operation typically these are written to a http.Request
*/
type PostMeInstallationTemplateParams struct {

	/*MeInstallationTemplatePost*/
	MeInstallationTemplatePost *models.PostMeInstallationTemplateParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me installation template params
func (o *PostMeInstallationTemplateParams) WithTimeout(timeout time.Duration) *PostMeInstallationTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me installation template params
func (o *PostMeInstallationTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me installation template params
func (o *PostMeInstallationTemplateParams) WithContext(ctx context.Context) *PostMeInstallationTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me installation template params
func (o *PostMeInstallationTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me installation template params
func (o *PostMeInstallationTemplateParams) WithHTTPClient(client *http.Client) *PostMeInstallationTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me installation template params
func (o *PostMeInstallationTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMeInstallationTemplatePost adds the meInstallationTemplatePost to the post me installation template params
func (o *PostMeInstallationTemplateParams) WithMeInstallationTemplatePost(meInstallationTemplatePost *models.PostMeInstallationTemplateParamsBody) *PostMeInstallationTemplateParams {
	o.SetMeInstallationTemplatePost(meInstallationTemplatePost)
	return o
}

// SetMeInstallationTemplatePost adds the meInstallationTemplatePost to the post me installation template params
func (o *PostMeInstallationTemplateParams) SetMeInstallationTemplatePost(meInstallationTemplatePost *models.PostMeInstallationTemplateParamsBody) {
	o.MeInstallationTemplatePost = meInstallationTemplatePost
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeInstallationTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MeInstallationTemplatePost != nil {
		if err := r.SetBodyParam(o.MeInstallationTemplatePost); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
