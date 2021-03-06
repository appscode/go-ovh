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

	"github.com/appscode/go-ovh/ip/models"
)

// NewPostIPIPMitigationParams creates a new PostIPIPMitigationParams object
// with the default values initialized.
func NewPostIPIPMitigationParams() *PostIPIPMitigationParams {
	var ()
	return &PostIPIPMitigationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIPIPMitigationParamsWithTimeout creates a new PostIPIPMitigationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIPIPMitigationParamsWithTimeout(timeout time.Duration) *PostIPIPMitigationParams {
	var ()
	return &PostIPIPMitigationParams{

		timeout: timeout,
	}
}

// NewPostIPIPMitigationParamsWithContext creates a new PostIPIPMitigationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIPIPMitigationParamsWithContext(ctx context.Context) *PostIPIPMitigationParams {
	var ()
	return &PostIPIPMitigationParams{

		Context: ctx,
	}
}

// NewPostIPIPMitigationParamsWithHTTPClient creates a new PostIPIPMitigationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIPIPMitigationParamsWithHTTPClient(client *http.Client) *PostIPIPMitigationParams {
	var ()
	return &PostIPIPMitigationParams{
		HTTPClient: client,
	}
}

/*PostIPIPMitigationParams contains all the parameters to send to the API endpoint
for the post IP IP mitigation operation typically these are written to a http.Request
*/
type PostIPIPMitigationParams struct {

	/*IP*/
	IP string
	/*IPMitigationPost*/
	IPMitigationPost *models.PostIPIPMitigationParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post IP IP mitigation params
func (o *PostIPIPMitigationParams) WithTimeout(timeout time.Duration) *PostIPIPMitigationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post IP IP mitigation params
func (o *PostIPIPMitigationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post IP IP mitigation params
func (o *PostIPIPMitigationParams) WithContext(ctx context.Context) *PostIPIPMitigationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post IP IP mitigation params
func (o *PostIPIPMitigationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post IP IP mitigation params
func (o *PostIPIPMitigationParams) WithHTTPClient(client *http.Client) *PostIPIPMitigationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post IP IP mitigation params
func (o *PostIPIPMitigationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIP adds the ip to the post IP IP mitigation params
func (o *PostIPIPMitigationParams) WithIP(ip string) *PostIPIPMitigationParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the post IP IP mitigation params
func (o *PostIPIPMitigationParams) SetIP(ip string) {
	o.IP = ip
}

// WithIPMitigationPost adds the iPMitigationPost to the post IP IP mitigation params
func (o *PostIPIPMitigationParams) WithIPMitigationPost(iPMitigationPost *models.PostIPIPMitigationParamsBody) *PostIPIPMitigationParams {
	o.SetIPMitigationPost(iPMitigationPost)
	return o
}

// SetIPMitigationPost adds the ipMitigationPost to the post IP IP mitigation params
func (o *PostIPIPMitigationParams) SetIPMitigationPost(iPMitigationPost *models.PostIPIPMitigationParamsBody) {
	o.IPMitigationPost = iPMitigationPost
}

// WriteToRequest writes these params to a swagger request
func (o *PostIPIPMitigationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ip
	if err := r.SetPathParam("ip", o.IP); err != nil {
		return err
	}

	if o.IPMitigationPost != nil {
		if err := r.SetBodyParam(o.IPMitigationPost); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
