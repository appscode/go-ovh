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

// NewDeleteDomainZoneZoneNameRecordIDParams creates a new DeleteDomainZoneZoneNameRecordIDParams object
// with the default values initialized.
func NewDeleteDomainZoneZoneNameRecordIDParams() *DeleteDomainZoneZoneNameRecordIDParams {
	var ()
	return &DeleteDomainZoneZoneNameRecordIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDomainZoneZoneNameRecordIDParamsWithTimeout creates a new DeleteDomainZoneZoneNameRecordIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDomainZoneZoneNameRecordIDParamsWithTimeout(timeout time.Duration) *DeleteDomainZoneZoneNameRecordIDParams {
	var ()
	return &DeleteDomainZoneZoneNameRecordIDParams{

		timeout: timeout,
	}
}

// NewDeleteDomainZoneZoneNameRecordIDParamsWithContext creates a new DeleteDomainZoneZoneNameRecordIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDomainZoneZoneNameRecordIDParamsWithContext(ctx context.Context) *DeleteDomainZoneZoneNameRecordIDParams {
	var ()
	return &DeleteDomainZoneZoneNameRecordIDParams{

		Context: ctx,
	}
}

// NewDeleteDomainZoneZoneNameRecordIDParamsWithHTTPClient creates a new DeleteDomainZoneZoneNameRecordIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDomainZoneZoneNameRecordIDParamsWithHTTPClient(client *http.Client) *DeleteDomainZoneZoneNameRecordIDParams {
	var ()
	return &DeleteDomainZoneZoneNameRecordIDParams{
		HTTPClient: client,
	}
}

/*DeleteDomainZoneZoneNameRecordIDParams contains all the parameters to send to the API endpoint
for the delete domain zone zone name record ID operation typically these are written to a http.Request
*/
type DeleteDomainZoneZoneNameRecordIDParams struct {

	/*ID*/
	ID int64
	/*ZoneName*/
	ZoneName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete domain zone zone name record ID params
func (o *DeleteDomainZoneZoneNameRecordIDParams) WithTimeout(timeout time.Duration) *DeleteDomainZoneZoneNameRecordIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete domain zone zone name record ID params
func (o *DeleteDomainZoneZoneNameRecordIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete domain zone zone name record ID params
func (o *DeleteDomainZoneZoneNameRecordIDParams) WithContext(ctx context.Context) *DeleteDomainZoneZoneNameRecordIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete domain zone zone name record ID params
func (o *DeleteDomainZoneZoneNameRecordIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete domain zone zone name record ID params
func (o *DeleteDomainZoneZoneNameRecordIDParams) WithHTTPClient(client *http.Client) *DeleteDomainZoneZoneNameRecordIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete domain zone zone name record ID params
func (o *DeleteDomainZoneZoneNameRecordIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete domain zone zone name record ID params
func (o *DeleteDomainZoneZoneNameRecordIDParams) WithID(id int64) *DeleteDomainZoneZoneNameRecordIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete domain zone zone name record ID params
func (o *DeleteDomainZoneZoneNameRecordIDParams) SetID(id int64) {
	o.ID = id
}

// WithZoneName adds the zoneName to the delete domain zone zone name record ID params
func (o *DeleteDomainZoneZoneNameRecordIDParams) WithZoneName(zoneName string) *DeleteDomainZoneZoneNameRecordIDParams {
	o.SetZoneName(zoneName)
	return o
}

// SetZoneName adds the zoneName to the delete domain zone zone name record ID params
func (o *DeleteDomainZoneZoneNameRecordIDParams) SetZoneName(zoneName string) {
	o.ZoneName = zoneName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDomainZoneZoneNameRecordIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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