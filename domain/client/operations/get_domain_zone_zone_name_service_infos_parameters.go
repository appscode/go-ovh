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
)

// NewGetDomainZoneZoneNameServiceInfosParams creates a new GetDomainZoneZoneNameServiceInfosParams object
// with the default values initialized.
func NewGetDomainZoneZoneNameServiceInfosParams() *GetDomainZoneZoneNameServiceInfosParams {
	var ()
	return &GetDomainZoneZoneNameServiceInfosParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomainZoneZoneNameServiceInfosParamsWithTimeout creates a new GetDomainZoneZoneNameServiceInfosParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDomainZoneZoneNameServiceInfosParamsWithTimeout(timeout time.Duration) *GetDomainZoneZoneNameServiceInfosParams {
	var ()
	return &GetDomainZoneZoneNameServiceInfosParams{

		timeout: timeout,
	}
}

// NewGetDomainZoneZoneNameServiceInfosParamsWithContext creates a new GetDomainZoneZoneNameServiceInfosParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDomainZoneZoneNameServiceInfosParamsWithContext(ctx context.Context) *GetDomainZoneZoneNameServiceInfosParams {
	var ()
	return &GetDomainZoneZoneNameServiceInfosParams{

		Context: ctx,
	}
}

// NewGetDomainZoneZoneNameServiceInfosParamsWithHTTPClient creates a new GetDomainZoneZoneNameServiceInfosParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDomainZoneZoneNameServiceInfosParamsWithHTTPClient(client *http.Client) *GetDomainZoneZoneNameServiceInfosParams {
	var ()
	return &GetDomainZoneZoneNameServiceInfosParams{
		HTTPClient: client,
	}
}

/*GetDomainZoneZoneNameServiceInfosParams contains all the parameters to send to the API endpoint
for the get domain zone zone name service infos operation typically these are written to a http.Request
*/
type GetDomainZoneZoneNameServiceInfosParams struct {

	/*ZoneName*/
	ZoneName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get domain zone zone name service infos params
func (o *GetDomainZoneZoneNameServiceInfosParams) WithTimeout(timeout time.Duration) *GetDomainZoneZoneNameServiceInfosParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domain zone zone name service infos params
func (o *GetDomainZoneZoneNameServiceInfosParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domain zone zone name service infos params
func (o *GetDomainZoneZoneNameServiceInfosParams) WithContext(ctx context.Context) *GetDomainZoneZoneNameServiceInfosParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domain zone zone name service infos params
func (o *GetDomainZoneZoneNameServiceInfosParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domain zone zone name service infos params
func (o *GetDomainZoneZoneNameServiceInfosParams) WithHTTPClient(client *http.Client) *GetDomainZoneZoneNameServiceInfosParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domain zone zone name service infos params
func (o *GetDomainZoneZoneNameServiceInfosParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithZoneName adds the zoneName to the get domain zone zone name service infos params
func (o *GetDomainZoneZoneNameServiceInfosParams) WithZoneName(zoneName string) *GetDomainZoneZoneNameServiceInfosParams {
	o.SetZoneName(zoneName)
	return o
}

// SetZoneName adds the zoneName to the get domain zone zone name service infos params
func (o *GetDomainZoneZoneNameServiceInfosParams) SetZoneName(zoneName string) {
	o.ZoneName = zoneName
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomainZoneZoneNameServiceInfosParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param zoneName
	if err := r.SetPathParam("zoneName", o.ZoneName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
