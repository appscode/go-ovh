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

// NewPostDomainZoneZoneNameConfirmTerminationParams creates a new PostDomainZoneZoneNameConfirmTerminationParams object
// with the default values initialized.
func NewPostDomainZoneZoneNameConfirmTerminationParams() *PostDomainZoneZoneNameConfirmTerminationParams {
	var ()
	return &PostDomainZoneZoneNameConfirmTerminationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDomainZoneZoneNameConfirmTerminationParamsWithTimeout creates a new PostDomainZoneZoneNameConfirmTerminationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDomainZoneZoneNameConfirmTerminationParamsWithTimeout(timeout time.Duration) *PostDomainZoneZoneNameConfirmTerminationParams {
	var ()
	return &PostDomainZoneZoneNameConfirmTerminationParams{

		timeout: timeout,
	}
}

// NewPostDomainZoneZoneNameConfirmTerminationParamsWithContext creates a new PostDomainZoneZoneNameConfirmTerminationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDomainZoneZoneNameConfirmTerminationParamsWithContext(ctx context.Context) *PostDomainZoneZoneNameConfirmTerminationParams {
	var ()
	return &PostDomainZoneZoneNameConfirmTerminationParams{

		Context: ctx,
	}
}

// NewPostDomainZoneZoneNameConfirmTerminationParamsWithHTTPClient creates a new PostDomainZoneZoneNameConfirmTerminationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDomainZoneZoneNameConfirmTerminationParamsWithHTTPClient(client *http.Client) *PostDomainZoneZoneNameConfirmTerminationParams {
	var ()
	return &PostDomainZoneZoneNameConfirmTerminationParams{
		HTTPClient: client,
	}
}

/*PostDomainZoneZoneNameConfirmTerminationParams contains all the parameters to send to the API endpoint
for the post domain zone zone name confirm termination operation typically these are written to a http.Request
*/
type PostDomainZoneZoneNameConfirmTerminationParams struct {

	/*DomainZoneConfirmTerminationPost*/
	DomainZoneConfirmTerminationPost *models.PostDomainZoneZoneNameConfirmTerminationParamsBody
	/*ZoneName*/
	ZoneName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post domain zone zone name confirm termination params
func (o *PostDomainZoneZoneNameConfirmTerminationParams) WithTimeout(timeout time.Duration) *PostDomainZoneZoneNameConfirmTerminationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post domain zone zone name confirm termination params
func (o *PostDomainZoneZoneNameConfirmTerminationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post domain zone zone name confirm termination params
func (o *PostDomainZoneZoneNameConfirmTerminationParams) WithContext(ctx context.Context) *PostDomainZoneZoneNameConfirmTerminationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post domain zone zone name confirm termination params
func (o *PostDomainZoneZoneNameConfirmTerminationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post domain zone zone name confirm termination params
func (o *PostDomainZoneZoneNameConfirmTerminationParams) WithHTTPClient(client *http.Client) *PostDomainZoneZoneNameConfirmTerminationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post domain zone zone name confirm termination params
func (o *PostDomainZoneZoneNameConfirmTerminationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainZoneConfirmTerminationPost adds the domainZoneConfirmTerminationPost to the post domain zone zone name confirm termination params
func (o *PostDomainZoneZoneNameConfirmTerminationParams) WithDomainZoneConfirmTerminationPost(domainZoneConfirmTerminationPost *models.PostDomainZoneZoneNameConfirmTerminationParamsBody) *PostDomainZoneZoneNameConfirmTerminationParams {
	o.SetDomainZoneConfirmTerminationPost(domainZoneConfirmTerminationPost)
	return o
}

// SetDomainZoneConfirmTerminationPost adds the domainZoneConfirmTerminationPost to the post domain zone zone name confirm termination params
func (o *PostDomainZoneZoneNameConfirmTerminationParams) SetDomainZoneConfirmTerminationPost(domainZoneConfirmTerminationPost *models.PostDomainZoneZoneNameConfirmTerminationParamsBody) {
	o.DomainZoneConfirmTerminationPost = domainZoneConfirmTerminationPost
}

// WithZoneName adds the zoneName to the post domain zone zone name confirm termination params
func (o *PostDomainZoneZoneNameConfirmTerminationParams) WithZoneName(zoneName string) *PostDomainZoneZoneNameConfirmTerminationParams {
	o.SetZoneName(zoneName)
	return o
}

// SetZoneName adds the zoneName to the post domain zone zone name confirm termination params
func (o *PostDomainZoneZoneNameConfirmTerminationParams) SetZoneName(zoneName string) {
	o.ZoneName = zoneName
}

// WriteToRequest writes these params to a swagger request
func (o *PostDomainZoneZoneNameConfirmTerminationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DomainZoneConfirmTerminationPost != nil {
		if err := r.SetBodyParam(o.DomainZoneConfirmTerminationPost); err != nil {
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