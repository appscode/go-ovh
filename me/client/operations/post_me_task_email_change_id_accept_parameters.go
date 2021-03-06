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

	"github.com/appscode/go-ovh/me/models"
)

// NewPostMeTaskEmailChangeIDAcceptParams creates a new PostMeTaskEmailChangeIDAcceptParams object
// with the default values initialized.
func NewPostMeTaskEmailChangeIDAcceptParams() *PostMeTaskEmailChangeIDAcceptParams {
	var ()
	return &PostMeTaskEmailChangeIDAcceptParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeTaskEmailChangeIDAcceptParamsWithTimeout creates a new PostMeTaskEmailChangeIDAcceptParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeTaskEmailChangeIDAcceptParamsWithTimeout(timeout time.Duration) *PostMeTaskEmailChangeIDAcceptParams {
	var ()
	return &PostMeTaskEmailChangeIDAcceptParams{

		timeout: timeout,
	}
}

// NewPostMeTaskEmailChangeIDAcceptParamsWithContext creates a new PostMeTaskEmailChangeIDAcceptParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeTaskEmailChangeIDAcceptParamsWithContext(ctx context.Context) *PostMeTaskEmailChangeIDAcceptParams {
	var ()
	return &PostMeTaskEmailChangeIDAcceptParams{

		Context: ctx,
	}
}

// NewPostMeTaskEmailChangeIDAcceptParamsWithHTTPClient creates a new PostMeTaskEmailChangeIDAcceptParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeTaskEmailChangeIDAcceptParamsWithHTTPClient(client *http.Client) *PostMeTaskEmailChangeIDAcceptParams {
	var ()
	return &PostMeTaskEmailChangeIDAcceptParams{
		HTTPClient: client,
	}
}

/*PostMeTaskEmailChangeIDAcceptParams contains all the parameters to send to the API endpoint
for the post me task email change ID accept operation typically these are written to a http.Request
*/
type PostMeTaskEmailChangeIDAcceptParams struct {

	/*ID*/
	ID int64
	/*MeTaskEmailChangeAcceptPost*/
	MeTaskEmailChangeAcceptPost *models.PostMeTaskEmailChangeIDAcceptParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me task email change ID accept params
func (o *PostMeTaskEmailChangeIDAcceptParams) WithTimeout(timeout time.Duration) *PostMeTaskEmailChangeIDAcceptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me task email change ID accept params
func (o *PostMeTaskEmailChangeIDAcceptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me task email change ID accept params
func (o *PostMeTaskEmailChangeIDAcceptParams) WithContext(ctx context.Context) *PostMeTaskEmailChangeIDAcceptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me task email change ID accept params
func (o *PostMeTaskEmailChangeIDAcceptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me task email change ID accept params
func (o *PostMeTaskEmailChangeIDAcceptParams) WithHTTPClient(client *http.Client) *PostMeTaskEmailChangeIDAcceptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me task email change ID accept params
func (o *PostMeTaskEmailChangeIDAcceptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post me task email change ID accept params
func (o *PostMeTaskEmailChangeIDAcceptParams) WithID(id int64) *PostMeTaskEmailChangeIDAcceptParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post me task email change ID accept params
func (o *PostMeTaskEmailChangeIDAcceptParams) SetID(id int64) {
	o.ID = id
}

// WithMeTaskEmailChangeAcceptPost adds the meTaskEmailChangeAcceptPost to the post me task email change ID accept params
func (o *PostMeTaskEmailChangeIDAcceptParams) WithMeTaskEmailChangeAcceptPost(meTaskEmailChangeAcceptPost *models.PostMeTaskEmailChangeIDAcceptParamsBody) *PostMeTaskEmailChangeIDAcceptParams {
	o.SetMeTaskEmailChangeAcceptPost(meTaskEmailChangeAcceptPost)
	return o
}

// SetMeTaskEmailChangeAcceptPost adds the meTaskEmailChangeAcceptPost to the post me task email change ID accept params
func (o *PostMeTaskEmailChangeIDAcceptParams) SetMeTaskEmailChangeAcceptPost(meTaskEmailChangeAcceptPost *models.PostMeTaskEmailChangeIDAcceptParamsBody) {
	o.MeTaskEmailChangeAcceptPost = meTaskEmailChangeAcceptPost
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeTaskEmailChangeIDAcceptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.MeTaskEmailChangeAcceptPost != nil {
		if err := r.SetBodyParam(o.MeTaskEmailChangeAcceptPost); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
