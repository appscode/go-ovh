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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/appscode/go-ovh/iploadbalancing/models"
)

// PostIPLoadbalancingServiceNameRefreshReader is a Reader for the PostIPLoadbalancingServiceNameRefresh structure.
type PostIPLoadbalancingServiceNameRefreshReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIPLoadbalancingServiceNameRefreshReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostIPLoadbalancingServiceNameRefreshOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostIPLoadbalancingServiceNameRefreshDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIPLoadbalancingServiceNameRefreshOK creates a PostIPLoadbalancingServiceNameRefreshOK with default headers values
func NewPostIPLoadbalancingServiceNameRefreshOK() *PostIPLoadbalancingServiceNameRefreshOK {
	return &PostIPLoadbalancingServiceNameRefreshOK{}
}

/*PostIPLoadbalancingServiceNameRefreshOK handles this case with default header values.

description of 'iplb.Task.Task' response
*/
type PostIPLoadbalancingServiceNameRefreshOK struct {
	Payload *models.IPLBTaskTask
}

func (o *PostIPLoadbalancingServiceNameRefreshOK) Error() string {
	return fmt.Sprintf("[POST /ipLoadbalancing/{serviceName}/refresh][%d] postIpLoadbalancingServiceNameRefreshOK  %+v", 200, o.Payload)
}

func (o *PostIPLoadbalancingServiceNameRefreshOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPLBTaskTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIPLoadbalancingServiceNameRefreshDefault creates a PostIPLoadbalancingServiceNameRefreshDefault with default headers values
func NewPostIPLoadbalancingServiceNameRefreshDefault(code int) *PostIPLoadbalancingServiceNameRefreshDefault {
	return &PostIPLoadbalancingServiceNameRefreshDefault{
		_statusCode: code,
	}
}

/*PostIPLoadbalancingServiceNameRefreshDefault handles this case with default header values.

Unexpected error
*/
type PostIPLoadbalancingServiceNameRefreshDefault struct {
	_statusCode int

	Payload *models.PostIPLoadbalancingServiceNameRefreshDefaultBody
}

// Code gets the status code for the post IP loadbalancing service name refresh default response
func (o *PostIPLoadbalancingServiceNameRefreshDefault) Code() int {
	return o._statusCode
}

func (o *PostIPLoadbalancingServiceNameRefreshDefault) Error() string {
	return fmt.Sprintf("[POST /ipLoadbalancing/{serviceName}/refresh][%d] PostIPLoadbalancingServiceNameRefresh default  %+v", o._statusCode, o.Payload)
}

func (o *PostIPLoadbalancingServiceNameRefreshDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostIPLoadbalancingServiceNameRefreshDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
