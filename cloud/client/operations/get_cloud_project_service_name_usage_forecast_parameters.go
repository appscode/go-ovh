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

// NewGetCloudProjectServiceNameUsageForecastParams creates a new GetCloudProjectServiceNameUsageForecastParams object
// with the default values initialized.
func NewGetCloudProjectServiceNameUsageForecastParams() *GetCloudProjectServiceNameUsageForecastParams {
	var ()
	return &GetCloudProjectServiceNameUsageForecastParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCloudProjectServiceNameUsageForecastParamsWithTimeout creates a new GetCloudProjectServiceNameUsageForecastParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCloudProjectServiceNameUsageForecastParamsWithTimeout(timeout time.Duration) *GetCloudProjectServiceNameUsageForecastParams {
	var ()
	return &GetCloudProjectServiceNameUsageForecastParams{

		timeout: timeout,
	}
}

// NewGetCloudProjectServiceNameUsageForecastParamsWithContext creates a new GetCloudProjectServiceNameUsageForecastParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCloudProjectServiceNameUsageForecastParamsWithContext(ctx context.Context) *GetCloudProjectServiceNameUsageForecastParams {
	var ()
	return &GetCloudProjectServiceNameUsageForecastParams{

		Context: ctx,
	}
}

// NewGetCloudProjectServiceNameUsageForecastParamsWithHTTPClient creates a new GetCloudProjectServiceNameUsageForecastParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCloudProjectServiceNameUsageForecastParamsWithHTTPClient(client *http.Client) *GetCloudProjectServiceNameUsageForecastParams {
	var ()
	return &GetCloudProjectServiceNameUsageForecastParams{
		HTTPClient: client,
	}
}

/*GetCloudProjectServiceNameUsageForecastParams contains all the parameters to send to the API endpoint
for the get cloud project service name usage forecast operation typically these are written to a http.Request
*/
type GetCloudProjectServiceNameUsageForecastParams struct {

	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cloud project service name usage forecast params
func (o *GetCloudProjectServiceNameUsageForecastParams) WithTimeout(timeout time.Duration) *GetCloudProjectServiceNameUsageForecastParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cloud project service name usage forecast params
func (o *GetCloudProjectServiceNameUsageForecastParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cloud project service name usage forecast params
func (o *GetCloudProjectServiceNameUsageForecastParams) WithContext(ctx context.Context) *GetCloudProjectServiceNameUsageForecastParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cloud project service name usage forecast params
func (o *GetCloudProjectServiceNameUsageForecastParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cloud project service name usage forecast params
func (o *GetCloudProjectServiceNameUsageForecastParams) WithHTTPClient(client *http.Client) *GetCloudProjectServiceNameUsageForecastParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cloud project service name usage forecast params
func (o *GetCloudProjectServiceNameUsageForecastParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceName adds the serviceName to the get cloud project service name usage forecast params
func (o *GetCloudProjectServiceNameUsageForecastParams) WithServiceName(serviceName string) *GetCloudProjectServiceNameUsageForecastParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get cloud project service name usage forecast params
func (o *GetCloudProjectServiceNameUsageForecastParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetCloudProjectServiceNameUsageForecastParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
