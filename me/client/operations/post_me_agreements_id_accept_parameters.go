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

// NewPostMeAgreementsIDAcceptParams creates a new PostMeAgreementsIDAcceptParams object
// with the default values initialized.
func NewPostMeAgreementsIDAcceptParams() *PostMeAgreementsIDAcceptParams {
	var ()
	return &PostMeAgreementsIDAcceptParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeAgreementsIDAcceptParamsWithTimeout creates a new PostMeAgreementsIDAcceptParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeAgreementsIDAcceptParamsWithTimeout(timeout time.Duration) *PostMeAgreementsIDAcceptParams {
	var ()
	return &PostMeAgreementsIDAcceptParams{

		timeout: timeout,
	}
}

// NewPostMeAgreementsIDAcceptParamsWithContext creates a new PostMeAgreementsIDAcceptParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeAgreementsIDAcceptParamsWithContext(ctx context.Context) *PostMeAgreementsIDAcceptParams {
	var ()
	return &PostMeAgreementsIDAcceptParams{

		Context: ctx,
	}
}

// NewPostMeAgreementsIDAcceptParamsWithHTTPClient creates a new PostMeAgreementsIDAcceptParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeAgreementsIDAcceptParamsWithHTTPClient(client *http.Client) *PostMeAgreementsIDAcceptParams {
	var ()
	return &PostMeAgreementsIDAcceptParams{
		HTTPClient: client,
	}
}

/*PostMeAgreementsIDAcceptParams contains all the parameters to send to the API endpoint
for the post me agreements ID accept operation typically these are written to a http.Request
*/
type PostMeAgreementsIDAcceptParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me agreements ID accept params
func (o *PostMeAgreementsIDAcceptParams) WithTimeout(timeout time.Duration) *PostMeAgreementsIDAcceptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me agreements ID accept params
func (o *PostMeAgreementsIDAcceptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me agreements ID accept params
func (o *PostMeAgreementsIDAcceptParams) WithContext(ctx context.Context) *PostMeAgreementsIDAcceptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me agreements ID accept params
func (o *PostMeAgreementsIDAcceptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me agreements ID accept params
func (o *PostMeAgreementsIDAcceptParams) WithHTTPClient(client *http.Client) *PostMeAgreementsIDAcceptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me agreements ID accept params
func (o *PostMeAgreementsIDAcceptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post me agreements ID accept params
func (o *PostMeAgreementsIDAcceptParams) WithID(id int64) *PostMeAgreementsIDAcceptParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post me agreements ID accept params
func (o *PostMeAgreementsIDAcceptParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeAgreementsIDAcceptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
