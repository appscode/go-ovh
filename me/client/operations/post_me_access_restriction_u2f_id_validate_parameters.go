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

// NewPostMeAccessRestrictionU2fIDValidateParams creates a new PostMeAccessRestrictionU2fIDValidateParams object
// with the default values initialized.
func NewPostMeAccessRestrictionU2fIDValidateParams() *PostMeAccessRestrictionU2fIDValidateParams {
	var ()
	return &PostMeAccessRestrictionU2fIDValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeAccessRestrictionU2fIDValidateParamsWithTimeout creates a new PostMeAccessRestrictionU2fIDValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeAccessRestrictionU2fIDValidateParamsWithTimeout(timeout time.Duration) *PostMeAccessRestrictionU2fIDValidateParams {
	var ()
	return &PostMeAccessRestrictionU2fIDValidateParams{

		timeout: timeout,
	}
}

// NewPostMeAccessRestrictionU2fIDValidateParamsWithContext creates a new PostMeAccessRestrictionU2fIDValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeAccessRestrictionU2fIDValidateParamsWithContext(ctx context.Context) *PostMeAccessRestrictionU2fIDValidateParams {
	var ()
	return &PostMeAccessRestrictionU2fIDValidateParams{

		Context: ctx,
	}
}

// NewPostMeAccessRestrictionU2fIDValidateParamsWithHTTPClient creates a new PostMeAccessRestrictionU2fIDValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeAccessRestrictionU2fIDValidateParamsWithHTTPClient(client *http.Client) *PostMeAccessRestrictionU2fIDValidateParams {
	var ()
	return &PostMeAccessRestrictionU2fIDValidateParams{
		HTTPClient: client,
	}
}

/*PostMeAccessRestrictionU2fIDValidateParams contains all the parameters to send to the API endpoint
for the post me access restriction u2f ID validate operation typically these are written to a http.Request
*/
type PostMeAccessRestrictionU2fIDValidateParams struct {

	/*ID*/
	ID int64
	/*MeAccessRestrictionU2fValidatePost*/
	MeAccessRestrictionU2fValidatePost *models.PostMeAccessRestrictionU2fIDValidateParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me access restriction u2f ID validate params
func (o *PostMeAccessRestrictionU2fIDValidateParams) WithTimeout(timeout time.Duration) *PostMeAccessRestrictionU2fIDValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me access restriction u2f ID validate params
func (o *PostMeAccessRestrictionU2fIDValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me access restriction u2f ID validate params
func (o *PostMeAccessRestrictionU2fIDValidateParams) WithContext(ctx context.Context) *PostMeAccessRestrictionU2fIDValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me access restriction u2f ID validate params
func (o *PostMeAccessRestrictionU2fIDValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me access restriction u2f ID validate params
func (o *PostMeAccessRestrictionU2fIDValidateParams) WithHTTPClient(client *http.Client) *PostMeAccessRestrictionU2fIDValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me access restriction u2f ID validate params
func (o *PostMeAccessRestrictionU2fIDValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post me access restriction u2f ID validate params
func (o *PostMeAccessRestrictionU2fIDValidateParams) WithID(id int64) *PostMeAccessRestrictionU2fIDValidateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post me access restriction u2f ID validate params
func (o *PostMeAccessRestrictionU2fIDValidateParams) SetID(id int64) {
	o.ID = id
}

// WithMeAccessRestrictionU2fValidatePost adds the meAccessRestrictionU2fValidatePost to the post me access restriction u2f ID validate params
func (o *PostMeAccessRestrictionU2fIDValidateParams) WithMeAccessRestrictionU2fValidatePost(meAccessRestrictionU2fValidatePost *models.PostMeAccessRestrictionU2fIDValidateParamsBody) *PostMeAccessRestrictionU2fIDValidateParams {
	o.SetMeAccessRestrictionU2fValidatePost(meAccessRestrictionU2fValidatePost)
	return o
}

// SetMeAccessRestrictionU2fValidatePost adds the meAccessRestrictionU2fValidatePost to the post me access restriction u2f ID validate params
func (o *PostMeAccessRestrictionU2fIDValidateParams) SetMeAccessRestrictionU2fValidatePost(meAccessRestrictionU2fValidatePost *models.PostMeAccessRestrictionU2fIDValidateParamsBody) {
	o.MeAccessRestrictionU2fValidatePost = meAccessRestrictionU2fValidatePost
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeAccessRestrictionU2fIDValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.MeAccessRestrictionU2fValidatePost != nil {
		if err := r.SetBodyParam(o.MeAccessRestrictionU2fValidatePost); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
