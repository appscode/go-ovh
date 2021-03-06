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

// NewPostIPIPSpamIPSpammingUnblockParams creates a new PostIPIPSpamIPSpammingUnblockParams object
// with the default values initialized.
func NewPostIPIPSpamIPSpammingUnblockParams() *PostIPIPSpamIPSpammingUnblockParams {
	var ()
	return &PostIPIPSpamIPSpammingUnblockParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIPIPSpamIPSpammingUnblockParamsWithTimeout creates a new PostIPIPSpamIPSpammingUnblockParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIPIPSpamIPSpammingUnblockParamsWithTimeout(timeout time.Duration) *PostIPIPSpamIPSpammingUnblockParams {
	var ()
	return &PostIPIPSpamIPSpammingUnblockParams{

		timeout: timeout,
	}
}

// NewPostIPIPSpamIPSpammingUnblockParamsWithContext creates a new PostIPIPSpamIPSpammingUnblockParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIPIPSpamIPSpammingUnblockParamsWithContext(ctx context.Context) *PostIPIPSpamIPSpammingUnblockParams {
	var ()
	return &PostIPIPSpamIPSpammingUnblockParams{

		Context: ctx,
	}
}

// NewPostIPIPSpamIPSpammingUnblockParamsWithHTTPClient creates a new PostIPIPSpamIPSpammingUnblockParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIPIPSpamIPSpammingUnblockParamsWithHTTPClient(client *http.Client) *PostIPIPSpamIPSpammingUnblockParams {
	var ()
	return &PostIPIPSpamIPSpammingUnblockParams{
		HTTPClient: client,
	}
}

/*PostIPIPSpamIPSpammingUnblockParams contains all the parameters to send to the API endpoint
for the post IP IP spam IP spamming unblock operation typically these are written to a http.Request
*/
type PostIPIPSpamIPSpammingUnblockParams struct {

	/*IP*/
	IP string
	/*IPSpamming*/
	IPSpamming string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post IP IP spam IP spamming unblock params
func (o *PostIPIPSpamIPSpammingUnblockParams) WithTimeout(timeout time.Duration) *PostIPIPSpamIPSpammingUnblockParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post IP IP spam IP spamming unblock params
func (o *PostIPIPSpamIPSpammingUnblockParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post IP IP spam IP spamming unblock params
func (o *PostIPIPSpamIPSpammingUnblockParams) WithContext(ctx context.Context) *PostIPIPSpamIPSpammingUnblockParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post IP IP spam IP spamming unblock params
func (o *PostIPIPSpamIPSpammingUnblockParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post IP IP spam IP spamming unblock params
func (o *PostIPIPSpamIPSpammingUnblockParams) WithHTTPClient(client *http.Client) *PostIPIPSpamIPSpammingUnblockParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post IP IP spam IP spamming unblock params
func (o *PostIPIPSpamIPSpammingUnblockParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIP adds the ip to the post IP IP spam IP spamming unblock params
func (o *PostIPIPSpamIPSpammingUnblockParams) WithIP(ip string) *PostIPIPSpamIPSpammingUnblockParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the post IP IP spam IP spamming unblock params
func (o *PostIPIPSpamIPSpammingUnblockParams) SetIP(ip string) {
	o.IP = ip
}

// WithIPSpamming adds the iPSpamming to the post IP IP spam IP spamming unblock params
func (o *PostIPIPSpamIPSpammingUnblockParams) WithIPSpamming(iPSpamming string) *PostIPIPSpamIPSpammingUnblockParams {
	o.SetIPSpamming(iPSpamming)
	return o
}

// SetIPSpamming adds the ipSpamming to the post IP IP spam IP spamming unblock params
func (o *PostIPIPSpamIPSpammingUnblockParams) SetIPSpamming(iPSpamming string) {
	o.IPSpamming = iPSpamming
}

// WriteToRequest writes these params to a swagger request
func (o *PostIPIPSpamIPSpammingUnblockParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ip
	if err := r.SetPathParam("ip", o.IP); err != nil {
		return err
	}

	// path param ipSpamming
	if err := r.SetPathParam("ipSpamming", o.IPSpamming); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
