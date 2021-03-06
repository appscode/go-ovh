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

// PostIPLoadbalancingServiceNameUDPFarmReader is a Reader for the PostIPLoadbalancingServiceNameUDPFarm structure.
type PostIPLoadbalancingServiceNameUDPFarmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIPLoadbalancingServiceNameUDPFarmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostIPLoadbalancingServiceNameUDPFarmOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostIPLoadbalancingServiceNameUDPFarmDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIPLoadbalancingServiceNameUDPFarmOK creates a PostIPLoadbalancingServiceNameUDPFarmOK with default headers values
func NewPostIPLoadbalancingServiceNameUDPFarmOK() *PostIPLoadbalancingServiceNameUDPFarmOK {
	return &PostIPLoadbalancingServiceNameUDPFarmOK{}
}

/*PostIPLoadbalancingServiceNameUDPFarmOK handles this case with default header values.

description of 'iplb.BackendUdp.BackendUdp' response
*/
type PostIPLoadbalancingServiceNameUDPFarmOK struct {
	Payload *models.IPLBBackendUDPBackendUDP
}

func (o *PostIPLoadbalancingServiceNameUDPFarmOK) Error() string {
	return fmt.Sprintf("[POST /ipLoadbalancing/{serviceName}/udp/farm][%d] postIpLoadbalancingServiceNameUdpFarmOK  %+v", 200, o.Payload)
}

func (o *PostIPLoadbalancingServiceNameUDPFarmOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPLBBackendUDPBackendUDP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIPLoadbalancingServiceNameUDPFarmDefault creates a PostIPLoadbalancingServiceNameUDPFarmDefault with default headers values
func NewPostIPLoadbalancingServiceNameUDPFarmDefault(code int) *PostIPLoadbalancingServiceNameUDPFarmDefault {
	return &PostIPLoadbalancingServiceNameUDPFarmDefault{
		_statusCode: code,
	}
}

/*PostIPLoadbalancingServiceNameUDPFarmDefault handles this case with default header values.

Unexpected error
*/
type PostIPLoadbalancingServiceNameUDPFarmDefault struct {
	_statusCode int

	Payload *models.PostIPLoadbalancingServiceNameUDPFarmDefaultBody
}

// Code gets the status code for the post IP loadbalancing service name UDP farm default response
func (o *PostIPLoadbalancingServiceNameUDPFarmDefault) Code() int {
	return o._statusCode
}

func (o *PostIPLoadbalancingServiceNameUDPFarmDefault) Error() string {
	return fmt.Sprintf("[POST /ipLoadbalancing/{serviceName}/udp/farm][%d] PostIPLoadbalancingServiceNameUDPFarm default  %+v", o._statusCode, o.Payload)
}

func (o *PostIPLoadbalancingServiceNameUDPFarmDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostIPLoadbalancingServiceNameUDPFarmDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
