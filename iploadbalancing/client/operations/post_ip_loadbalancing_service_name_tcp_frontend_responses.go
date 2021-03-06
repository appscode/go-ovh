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

// PostIPLoadbalancingServiceNameTCPFrontendReader is a Reader for the PostIPLoadbalancingServiceNameTCPFrontend structure.
type PostIPLoadbalancingServiceNameTCPFrontendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIPLoadbalancingServiceNameTCPFrontendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostIPLoadbalancingServiceNameTCPFrontendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostIPLoadbalancingServiceNameTCPFrontendDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIPLoadbalancingServiceNameTCPFrontendOK creates a PostIPLoadbalancingServiceNameTCPFrontendOK with default headers values
func NewPostIPLoadbalancingServiceNameTCPFrontendOK() *PostIPLoadbalancingServiceNameTCPFrontendOK {
	return &PostIPLoadbalancingServiceNameTCPFrontendOK{}
}

/*PostIPLoadbalancingServiceNameTCPFrontendOK handles this case with default header values.

description of 'iplb.FrontendTcp.FrontendTcp' response
*/
type PostIPLoadbalancingServiceNameTCPFrontendOK struct {
	Payload *models.IPLBFrontendTCPFrontendTCP
}

func (o *PostIPLoadbalancingServiceNameTCPFrontendOK) Error() string {
	return fmt.Sprintf("[POST /ipLoadbalancing/{serviceName}/tcp/frontend][%d] postIpLoadbalancingServiceNameTcpFrontendOK  %+v", 200, o.Payload)
}

func (o *PostIPLoadbalancingServiceNameTCPFrontendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPLBFrontendTCPFrontendTCP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIPLoadbalancingServiceNameTCPFrontendDefault creates a PostIPLoadbalancingServiceNameTCPFrontendDefault with default headers values
func NewPostIPLoadbalancingServiceNameTCPFrontendDefault(code int) *PostIPLoadbalancingServiceNameTCPFrontendDefault {
	return &PostIPLoadbalancingServiceNameTCPFrontendDefault{
		_statusCode: code,
	}
}

/*PostIPLoadbalancingServiceNameTCPFrontendDefault handles this case with default header values.

Unexpected error
*/
type PostIPLoadbalancingServiceNameTCPFrontendDefault struct {
	_statusCode int

	Payload *models.PostIPLoadbalancingServiceNameTCPFrontendDefaultBody
}

// Code gets the status code for the post IP loadbalancing service name TCP frontend default response
func (o *PostIPLoadbalancingServiceNameTCPFrontendDefault) Code() int {
	return o._statusCode
}

func (o *PostIPLoadbalancingServiceNameTCPFrontendDefault) Error() string {
	return fmt.Sprintf("[POST /ipLoadbalancing/{serviceName}/tcp/frontend][%d] PostIPLoadbalancingServiceNameTCPFrontend default  %+v", o._statusCode, o.Payload)
}

func (o *PostIPLoadbalancingServiceNameTCPFrontendDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostIPLoadbalancingServiceNameTCPFrontendDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
