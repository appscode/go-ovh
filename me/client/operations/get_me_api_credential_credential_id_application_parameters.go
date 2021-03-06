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

// NewGetMeAPICredentialCredentialIDApplicationParams creates a new GetMeAPICredentialCredentialIDApplicationParams object
// with the default values initialized.
func NewGetMeAPICredentialCredentialIDApplicationParams() *GetMeAPICredentialCredentialIDApplicationParams {
	var ()
	return &GetMeAPICredentialCredentialIDApplicationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeAPICredentialCredentialIDApplicationParamsWithTimeout creates a new GetMeAPICredentialCredentialIDApplicationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeAPICredentialCredentialIDApplicationParamsWithTimeout(timeout time.Duration) *GetMeAPICredentialCredentialIDApplicationParams {
	var ()
	return &GetMeAPICredentialCredentialIDApplicationParams{

		timeout: timeout,
	}
}

// NewGetMeAPICredentialCredentialIDApplicationParamsWithContext creates a new GetMeAPICredentialCredentialIDApplicationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeAPICredentialCredentialIDApplicationParamsWithContext(ctx context.Context) *GetMeAPICredentialCredentialIDApplicationParams {
	var ()
	return &GetMeAPICredentialCredentialIDApplicationParams{

		Context: ctx,
	}
}

// NewGetMeAPICredentialCredentialIDApplicationParamsWithHTTPClient creates a new GetMeAPICredentialCredentialIDApplicationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeAPICredentialCredentialIDApplicationParamsWithHTTPClient(client *http.Client) *GetMeAPICredentialCredentialIDApplicationParams {
	var ()
	return &GetMeAPICredentialCredentialIDApplicationParams{
		HTTPClient: client,
	}
}

/*GetMeAPICredentialCredentialIDApplicationParams contains all the parameters to send to the API endpoint
for the get me API credential credential ID application operation typically these are written to a http.Request
*/
type GetMeAPICredentialCredentialIDApplicationParams struct {

	/*CredentialID*/
	CredentialID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me API credential credential ID application params
func (o *GetMeAPICredentialCredentialIDApplicationParams) WithTimeout(timeout time.Duration) *GetMeAPICredentialCredentialIDApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me API credential credential ID application params
func (o *GetMeAPICredentialCredentialIDApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me API credential credential ID application params
func (o *GetMeAPICredentialCredentialIDApplicationParams) WithContext(ctx context.Context) *GetMeAPICredentialCredentialIDApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me API credential credential ID application params
func (o *GetMeAPICredentialCredentialIDApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me API credential credential ID application params
func (o *GetMeAPICredentialCredentialIDApplicationParams) WithHTTPClient(client *http.Client) *GetMeAPICredentialCredentialIDApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me API credential credential ID application params
func (o *GetMeAPICredentialCredentialIDApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredentialID adds the credentialID to the get me API credential credential ID application params
func (o *GetMeAPICredentialCredentialIDApplicationParams) WithCredentialID(credentialID int64) *GetMeAPICredentialCredentialIDApplicationParams {
	o.SetCredentialID(credentialID)
	return o
}

// SetCredentialID adds the credentialId to the get me API credential credential ID application params
func (o *GetMeAPICredentialCredentialIDApplicationParams) SetCredentialID(credentialID int64) {
	o.CredentialID = credentialID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeAPICredentialCredentialIDApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param credentialId
	if err := r.SetPathParam("credentialId", swag.FormatInt64(o.CredentialID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
