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

// NewDeleteDomainZoneZoneNameRedirectionIDParams creates a new DeleteDomainZoneZoneNameRedirectionIDParams object
// with the default values initialized.
func NewDeleteDomainZoneZoneNameRedirectionIDParams() *DeleteDomainZoneZoneNameRedirectionIDParams {
	var ()
	return &DeleteDomainZoneZoneNameRedirectionIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDomainZoneZoneNameRedirectionIDParamsWithTimeout creates a new DeleteDomainZoneZoneNameRedirectionIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDomainZoneZoneNameRedirectionIDParamsWithTimeout(timeout time.Duration) *DeleteDomainZoneZoneNameRedirectionIDParams {
	var ()
	return &DeleteDomainZoneZoneNameRedirectionIDParams{

		timeout: timeout,
	}
}

// NewDeleteDomainZoneZoneNameRedirectionIDParamsWithContext creates a new DeleteDomainZoneZoneNameRedirectionIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDomainZoneZoneNameRedirectionIDParamsWithContext(ctx context.Context) *DeleteDomainZoneZoneNameRedirectionIDParams {
	var ()
	return &DeleteDomainZoneZoneNameRedirectionIDParams{

		Context: ctx,
	}
}

// NewDeleteDomainZoneZoneNameRedirectionIDParamsWithHTTPClient creates a new DeleteDomainZoneZoneNameRedirectionIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDomainZoneZoneNameRedirectionIDParamsWithHTTPClient(client *http.Client) *DeleteDomainZoneZoneNameRedirectionIDParams {
	var ()
	return &DeleteDomainZoneZoneNameRedirectionIDParams{
		HTTPClient: client,
	}
}

/*DeleteDomainZoneZoneNameRedirectionIDParams contains all the parameters to send to the API endpoint
for the delete domain zone zone name redirection ID operation typically these are written to a http.Request
*/
type DeleteDomainZoneZoneNameRedirectionIDParams struct {

	/*ID*/
	ID int64
	/*ZoneName*/
	ZoneName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete domain zone zone name redirection ID params
func (o *DeleteDomainZoneZoneNameRedirectionIDParams) WithTimeout(timeout time.Duration) *DeleteDomainZoneZoneNameRedirectionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete domain zone zone name redirection ID params
func (o *DeleteDomainZoneZoneNameRedirectionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete domain zone zone name redirection ID params
func (o *DeleteDomainZoneZoneNameRedirectionIDParams) WithContext(ctx context.Context) *DeleteDomainZoneZoneNameRedirectionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete domain zone zone name redirection ID params
func (o *DeleteDomainZoneZoneNameRedirectionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete domain zone zone name redirection ID params
func (o *DeleteDomainZoneZoneNameRedirectionIDParams) WithHTTPClient(client *http.Client) *DeleteDomainZoneZoneNameRedirectionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete domain zone zone name redirection ID params
func (o *DeleteDomainZoneZoneNameRedirectionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete domain zone zone name redirection ID params
func (o *DeleteDomainZoneZoneNameRedirectionIDParams) WithID(id int64) *DeleteDomainZoneZoneNameRedirectionIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete domain zone zone name redirection ID params
func (o *DeleteDomainZoneZoneNameRedirectionIDParams) SetID(id int64) {
	o.ID = id
}

// WithZoneName adds the zoneName to the delete domain zone zone name redirection ID params
func (o *DeleteDomainZoneZoneNameRedirectionIDParams) WithZoneName(zoneName string) *DeleteDomainZoneZoneNameRedirectionIDParams {
	o.SetZoneName(zoneName)
	return o
}

// SetZoneName adds the zoneName to the delete domain zone zone name redirection ID params
func (o *DeleteDomainZoneZoneNameRedirectionIDParams) SetZoneName(zoneName string) {
	o.ZoneName = zoneName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDomainZoneZoneNameRedirectionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
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