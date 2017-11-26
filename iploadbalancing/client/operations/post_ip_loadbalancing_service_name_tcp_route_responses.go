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

// PostIPLoadbalancingServiceNameTCPRouteReader is a Reader for the PostIPLoadbalancingServiceNameTCPRoute structure.
type PostIPLoadbalancingServiceNameTCPRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIPLoadbalancingServiceNameTCPRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostIPLoadbalancingServiceNameTCPRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostIPLoadbalancingServiceNameTCPRouteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIPLoadbalancingServiceNameTCPRouteOK creates a PostIPLoadbalancingServiceNameTCPRouteOK with default headers values
func NewPostIPLoadbalancingServiceNameTCPRouteOK() *PostIPLoadbalancingServiceNameTCPRouteOK {
	return &PostIPLoadbalancingServiceNameTCPRouteOK{}
}

/*PostIPLoadbalancingServiceNameTCPRouteOK handles this case with default header values.

description of 'iplb.RouteTcp.RouteTcp' response
*/
type PostIPLoadbalancingServiceNameTCPRouteOK struct {
	Payload *models.IPLBRouteTCPRouteTCP
}

func (o *PostIPLoadbalancingServiceNameTCPRouteOK) Error() string {
	return fmt.Sprintf("[POST /ipLoadbalancing/{serviceName}/tcp/route][%d] postIpLoadbalancingServiceNameTcpRouteOK  %+v", 200, o.Payload)
}

func (o *PostIPLoadbalancingServiceNameTCPRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPLBRouteTCPRouteTCP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIPLoadbalancingServiceNameTCPRouteDefault creates a PostIPLoadbalancingServiceNameTCPRouteDefault with default headers values
func NewPostIPLoadbalancingServiceNameTCPRouteDefault(code int) *PostIPLoadbalancingServiceNameTCPRouteDefault {
	return &PostIPLoadbalancingServiceNameTCPRouteDefault{
		_statusCode: code,
	}
}

/*PostIPLoadbalancingServiceNameTCPRouteDefault handles this case with default header values.

Unexpected error
*/
type PostIPLoadbalancingServiceNameTCPRouteDefault struct {
	_statusCode int

	Payload *models.PostIPLoadbalancingServiceNameTCPRouteDefaultBody
}

// Code gets the status code for the post IP loadbalancing service name TCP route default response
func (o *PostIPLoadbalancingServiceNameTCPRouteDefault) Code() int {
	return o._statusCode
}

func (o *PostIPLoadbalancingServiceNameTCPRouteDefault) Error() string {
	return fmt.Sprintf("[POST /ipLoadbalancing/{serviceName}/tcp/route][%d] PostIPLoadbalancingServiceNameTCPRoute default  %+v", o._statusCode, o.Payload)
}

func (o *PostIPLoadbalancingServiceNameTCPRouteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostIPLoadbalancingServiceNameTCPRouteDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}