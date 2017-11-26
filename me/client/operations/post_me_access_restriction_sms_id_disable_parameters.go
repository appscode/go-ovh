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

// NewPostMeAccessRestrictionSmsIDDisableParams creates a new PostMeAccessRestrictionSmsIDDisableParams object
// with the default values initialized.
func NewPostMeAccessRestrictionSmsIDDisableParams() *PostMeAccessRestrictionSmsIDDisableParams {
	var ()
	return &PostMeAccessRestrictionSmsIDDisableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeAccessRestrictionSmsIDDisableParamsWithTimeout creates a new PostMeAccessRestrictionSmsIDDisableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeAccessRestrictionSmsIDDisableParamsWithTimeout(timeout time.Duration) *PostMeAccessRestrictionSmsIDDisableParams {
	var ()
	return &PostMeAccessRestrictionSmsIDDisableParams{

		timeout: timeout,
	}
}

// NewPostMeAccessRestrictionSmsIDDisableParamsWithContext creates a new PostMeAccessRestrictionSmsIDDisableParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeAccessRestrictionSmsIDDisableParamsWithContext(ctx context.Context) *PostMeAccessRestrictionSmsIDDisableParams {
	var ()
	return &PostMeAccessRestrictionSmsIDDisableParams{

		Context: ctx,
	}
}

// NewPostMeAccessRestrictionSmsIDDisableParamsWithHTTPClient creates a new PostMeAccessRestrictionSmsIDDisableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeAccessRestrictionSmsIDDisableParamsWithHTTPClient(client *http.Client) *PostMeAccessRestrictionSmsIDDisableParams {
	var ()
	return &PostMeAccessRestrictionSmsIDDisableParams{
		HTTPClient: client,
	}
}

/*PostMeAccessRestrictionSmsIDDisableParams contains all the parameters to send to the API endpoint
for the post me access restriction sms ID disable operation typically these are written to a http.Request
*/
type PostMeAccessRestrictionSmsIDDisableParams struct {

	/*ID*/
	ID int64
	/*MeAccessRestrictionSmsDisablePost*/
	MeAccessRestrictionSmsDisablePost *models.PostMeAccessRestrictionSmsIDDisableParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me access restriction sms ID disable params
func (o *PostMeAccessRestrictionSmsIDDisableParams) WithTimeout(timeout time.Duration) *PostMeAccessRestrictionSmsIDDisableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me access restriction sms ID disable params
func (o *PostMeAccessRestrictionSmsIDDisableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me access restriction sms ID disable params
func (o *PostMeAccessRestrictionSmsIDDisableParams) WithContext(ctx context.Context) *PostMeAccessRestrictionSmsIDDisableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me access restriction sms ID disable params
func (o *PostMeAccessRestrictionSmsIDDisableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me access restriction sms ID disable params
func (o *PostMeAccessRestrictionSmsIDDisableParams) WithHTTPClient(client *http.Client) *PostMeAccessRestrictionSmsIDDisableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me access restriction sms ID disable params
func (o *PostMeAccessRestrictionSmsIDDisableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post me access restriction sms ID disable params
func (o *PostMeAccessRestrictionSmsIDDisableParams) WithID(id int64) *PostMeAccessRestrictionSmsIDDisableParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post me access restriction sms ID disable params
func (o *PostMeAccessRestrictionSmsIDDisableParams) SetID(id int64) {
	o.ID = id
}

// WithMeAccessRestrictionSmsDisablePost adds the meAccessRestrictionSmsDisablePost to the post me access restriction sms ID disable params
func (o *PostMeAccessRestrictionSmsIDDisableParams) WithMeAccessRestrictionSmsDisablePost(meAccessRestrictionSmsDisablePost *models.PostMeAccessRestrictionSmsIDDisableParamsBody) *PostMeAccessRestrictionSmsIDDisableParams {
	o.SetMeAccessRestrictionSmsDisablePost(meAccessRestrictionSmsDisablePost)
	return o
}

// SetMeAccessRestrictionSmsDisablePost adds the meAccessRestrictionSmsDisablePost to the post me access restriction sms ID disable params
func (o *PostMeAccessRestrictionSmsIDDisableParams) SetMeAccessRestrictionSmsDisablePost(meAccessRestrictionSmsDisablePost *models.PostMeAccessRestrictionSmsIDDisableParamsBody) {
	o.MeAccessRestrictionSmsDisablePost = meAccessRestrictionSmsDisablePost
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeAccessRestrictionSmsIDDisableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.MeAccessRestrictionSmsDisablePost != nil {
		if err := r.SetBodyParam(o.MeAccessRestrictionSmsDisablePost); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}