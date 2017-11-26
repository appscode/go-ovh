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

// NewPostMeSubAccountIDCreateConsumerKeyParams creates a new PostMeSubAccountIDCreateConsumerKeyParams object
// with the default values initialized.
func NewPostMeSubAccountIDCreateConsumerKeyParams() *PostMeSubAccountIDCreateConsumerKeyParams {
	var ()
	return &PostMeSubAccountIDCreateConsumerKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeSubAccountIDCreateConsumerKeyParamsWithTimeout creates a new PostMeSubAccountIDCreateConsumerKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeSubAccountIDCreateConsumerKeyParamsWithTimeout(timeout time.Duration) *PostMeSubAccountIDCreateConsumerKeyParams {
	var ()
	return &PostMeSubAccountIDCreateConsumerKeyParams{

		timeout: timeout,
	}
}

// NewPostMeSubAccountIDCreateConsumerKeyParamsWithContext creates a new PostMeSubAccountIDCreateConsumerKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeSubAccountIDCreateConsumerKeyParamsWithContext(ctx context.Context) *PostMeSubAccountIDCreateConsumerKeyParams {
	var ()
	return &PostMeSubAccountIDCreateConsumerKeyParams{

		Context: ctx,
	}
}

// NewPostMeSubAccountIDCreateConsumerKeyParamsWithHTTPClient creates a new PostMeSubAccountIDCreateConsumerKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeSubAccountIDCreateConsumerKeyParamsWithHTTPClient(client *http.Client) *PostMeSubAccountIDCreateConsumerKeyParams {
	var ()
	return &PostMeSubAccountIDCreateConsumerKeyParams{
		HTTPClient: client,
	}
}

/*PostMeSubAccountIDCreateConsumerKeyParams contains all the parameters to send to the API endpoint
for the post me sub account ID create consumer key operation typically these are written to a http.Request
*/
type PostMeSubAccountIDCreateConsumerKeyParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me sub account ID create consumer key params
func (o *PostMeSubAccountIDCreateConsumerKeyParams) WithTimeout(timeout time.Duration) *PostMeSubAccountIDCreateConsumerKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me sub account ID create consumer key params
func (o *PostMeSubAccountIDCreateConsumerKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me sub account ID create consumer key params
func (o *PostMeSubAccountIDCreateConsumerKeyParams) WithContext(ctx context.Context) *PostMeSubAccountIDCreateConsumerKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me sub account ID create consumer key params
func (o *PostMeSubAccountIDCreateConsumerKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me sub account ID create consumer key params
func (o *PostMeSubAccountIDCreateConsumerKeyParams) WithHTTPClient(client *http.Client) *PostMeSubAccountIDCreateConsumerKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me sub account ID create consumer key params
func (o *PostMeSubAccountIDCreateConsumerKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post me sub account ID create consumer key params
func (o *PostMeSubAccountIDCreateConsumerKeyParams) WithID(id int64) *PostMeSubAccountIDCreateConsumerKeyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post me sub account ID create consumer key params
func (o *PostMeSubAccountIDCreateConsumerKeyParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeSubAccountIDCreateConsumerKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}