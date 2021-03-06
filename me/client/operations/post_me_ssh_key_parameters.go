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

// NewPostMeSSHKeyParams creates a new PostMeSSHKeyParams object
// with the default values initialized.
func NewPostMeSSHKeyParams() *PostMeSSHKeyParams {
	var ()
	return &PostMeSSHKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeSSHKeyParamsWithTimeout creates a new PostMeSSHKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeSSHKeyParamsWithTimeout(timeout time.Duration) *PostMeSSHKeyParams {
	var ()
	return &PostMeSSHKeyParams{

		timeout: timeout,
	}
}

// NewPostMeSSHKeyParamsWithContext creates a new PostMeSSHKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeSSHKeyParamsWithContext(ctx context.Context) *PostMeSSHKeyParams {
	var ()
	return &PostMeSSHKeyParams{

		Context: ctx,
	}
}

// NewPostMeSSHKeyParamsWithHTTPClient creates a new PostMeSSHKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeSSHKeyParamsWithHTTPClient(client *http.Client) *PostMeSSHKeyParams {
	var ()
	return &PostMeSSHKeyParams{
		HTTPClient: client,
	}
}

/*PostMeSSHKeyParams contains all the parameters to send to the API endpoint
for the post me SSH key operation typically these are written to a http.Request
*/
type PostMeSSHKeyParams struct {

	/*MeSSHKeyPost*/
	MeSSHKeyPost *models.PostMeSSHKeyParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me SSH key params
func (o *PostMeSSHKeyParams) WithTimeout(timeout time.Duration) *PostMeSSHKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me SSH key params
func (o *PostMeSSHKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me SSH key params
func (o *PostMeSSHKeyParams) WithContext(ctx context.Context) *PostMeSSHKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me SSH key params
func (o *PostMeSSHKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me SSH key params
func (o *PostMeSSHKeyParams) WithHTTPClient(client *http.Client) *PostMeSSHKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me SSH key params
func (o *PostMeSSHKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMeSSHKeyPost adds the meSSHKeyPost to the post me SSH key params
func (o *PostMeSSHKeyParams) WithMeSSHKeyPost(meSSHKeyPost *models.PostMeSSHKeyParamsBody) *PostMeSSHKeyParams {
	o.SetMeSSHKeyPost(meSSHKeyPost)
	return o
}

// SetMeSSHKeyPost adds the meSshKeyPost to the post me SSH key params
func (o *PostMeSSHKeyParams) SetMeSSHKeyPost(meSSHKeyPost *models.PostMeSSHKeyParamsBody) {
	o.MeSSHKeyPost = meSSHKeyPost
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeSSHKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MeSSHKeyPost != nil {
		if err := r.SetBodyParam(o.MeSSHKeyPost); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
