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

	"github.com/appscode/go-ovh/domain/models"
)

// NewPostDomainZoneZoneNameChangeContactParams creates a new PostDomainZoneZoneNameChangeContactParams object
// with the default values initialized.
func NewPostDomainZoneZoneNameChangeContactParams() *PostDomainZoneZoneNameChangeContactParams {
	var ()
	return &PostDomainZoneZoneNameChangeContactParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDomainZoneZoneNameChangeContactParamsWithTimeout creates a new PostDomainZoneZoneNameChangeContactParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDomainZoneZoneNameChangeContactParamsWithTimeout(timeout time.Duration) *PostDomainZoneZoneNameChangeContactParams {
	var ()
	return &PostDomainZoneZoneNameChangeContactParams{

		timeout: timeout,
	}
}

// NewPostDomainZoneZoneNameChangeContactParamsWithContext creates a new PostDomainZoneZoneNameChangeContactParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDomainZoneZoneNameChangeContactParamsWithContext(ctx context.Context) *PostDomainZoneZoneNameChangeContactParams {
	var ()
	return &PostDomainZoneZoneNameChangeContactParams{

		Context: ctx,
	}
}

// NewPostDomainZoneZoneNameChangeContactParamsWithHTTPClient creates a new PostDomainZoneZoneNameChangeContactParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDomainZoneZoneNameChangeContactParamsWithHTTPClient(client *http.Client) *PostDomainZoneZoneNameChangeContactParams {
	var ()
	return &PostDomainZoneZoneNameChangeContactParams{
		HTTPClient: client,
	}
}

/*PostDomainZoneZoneNameChangeContactParams contains all the parameters to send to the API endpoint
for the post domain zone zone name change contact operation typically these are written to a http.Request
*/
type PostDomainZoneZoneNameChangeContactParams struct {

	/*DomainZoneChangeContactPost*/
	DomainZoneChangeContactPost *models.PostDomainZoneZoneNameChangeContactParamsBody
	/*ZoneName*/
	ZoneName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post domain zone zone name change contact params
func (o *PostDomainZoneZoneNameChangeContactParams) WithTimeout(timeout time.Duration) *PostDomainZoneZoneNameChangeContactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post domain zone zone name change contact params
func (o *PostDomainZoneZoneNameChangeContactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post domain zone zone name change contact params
func (o *PostDomainZoneZoneNameChangeContactParams) WithContext(ctx context.Context) *PostDomainZoneZoneNameChangeContactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post domain zone zone name change contact params
func (o *PostDomainZoneZoneNameChangeContactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post domain zone zone name change contact params
func (o *PostDomainZoneZoneNameChangeContactParams) WithHTTPClient(client *http.Client) *PostDomainZoneZoneNameChangeContactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post domain zone zone name change contact params
func (o *PostDomainZoneZoneNameChangeContactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainZoneChangeContactPost adds the domainZoneChangeContactPost to the post domain zone zone name change contact params
func (o *PostDomainZoneZoneNameChangeContactParams) WithDomainZoneChangeContactPost(domainZoneChangeContactPost *models.PostDomainZoneZoneNameChangeContactParamsBody) *PostDomainZoneZoneNameChangeContactParams {
	o.SetDomainZoneChangeContactPost(domainZoneChangeContactPost)
	return o
}

// SetDomainZoneChangeContactPost adds the domainZoneChangeContactPost to the post domain zone zone name change contact params
func (o *PostDomainZoneZoneNameChangeContactParams) SetDomainZoneChangeContactPost(domainZoneChangeContactPost *models.PostDomainZoneZoneNameChangeContactParamsBody) {
	o.DomainZoneChangeContactPost = domainZoneChangeContactPost
}

// WithZoneName adds the zoneName to the post domain zone zone name change contact params
func (o *PostDomainZoneZoneNameChangeContactParams) WithZoneName(zoneName string) *PostDomainZoneZoneNameChangeContactParams {
	o.SetZoneName(zoneName)
	return o
}

// SetZoneName adds the zoneName to the post domain zone zone name change contact params
func (o *PostDomainZoneZoneNameChangeContactParams) SetZoneName(zoneName string) {
	o.ZoneName = zoneName
}

// WriteToRequest writes these params to a swagger request
func (o *PostDomainZoneZoneNameChangeContactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DomainZoneChangeContactPost != nil {
		if err := r.SetBodyParam(o.DomainZoneChangeContactPost); err != nil {
			return err
		}
	}

	// path param zoneName
	if err := r.SetPathParam("zoneName", o.ZoneName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
