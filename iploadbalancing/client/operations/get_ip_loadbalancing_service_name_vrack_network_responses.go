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

// GetIPLoadbalancingServiceNameVrackNetworkReader is a Reader for the GetIPLoadbalancingServiceNameVrackNetwork structure.
type GetIPLoadbalancingServiceNameVrackNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPLoadbalancingServiceNameVrackNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIPLoadbalancingServiceNameVrackNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetIPLoadbalancingServiceNameVrackNetworkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIPLoadbalancingServiceNameVrackNetworkOK creates a GetIPLoadbalancingServiceNameVrackNetworkOK with default headers values
func NewGetIPLoadbalancingServiceNameVrackNetworkOK() *GetIPLoadbalancingServiceNameVrackNetworkOK {
	return &GetIPLoadbalancingServiceNameVrackNetworkOK{}
}

/*GetIPLoadbalancingServiceNameVrackNetworkOK handles this case with default header values.

return value
*/
type GetIPLoadbalancingServiceNameVrackNetworkOK struct {
	Payload []int64
}

func (o *GetIPLoadbalancingServiceNameVrackNetworkOK) Error() string {
	return fmt.Sprintf("[GET /ipLoadbalancing/{serviceName}/vrack/network][%d] getIpLoadbalancingServiceNameVrackNetworkOK  %+v", 200, o.Payload)
}

func (o *GetIPLoadbalancingServiceNameVrackNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPLoadbalancingServiceNameVrackNetworkDefault creates a GetIPLoadbalancingServiceNameVrackNetworkDefault with default headers values
func NewGetIPLoadbalancingServiceNameVrackNetworkDefault(code int) *GetIPLoadbalancingServiceNameVrackNetworkDefault {
	return &GetIPLoadbalancingServiceNameVrackNetworkDefault{
		_statusCode: code,
	}
}

/*GetIPLoadbalancingServiceNameVrackNetworkDefault handles this case with default header values.

Unexpected error
*/
type GetIPLoadbalancingServiceNameVrackNetworkDefault struct {
	_statusCode int

	Payload *models.GetIPLoadbalancingServiceNameVrackNetworkDefaultBody
}

// Code gets the status code for the get IP loadbalancing service name vrack network default response
func (o *GetIPLoadbalancingServiceNameVrackNetworkDefault) Code() int {
	return o._statusCode
}

func (o *GetIPLoadbalancingServiceNameVrackNetworkDefault) Error() string {
	return fmt.Sprintf("[GET /ipLoadbalancing/{serviceName}/vrack/network][%d] GetIPLoadbalancingServiceNameVrackNetwork default  %+v", o._statusCode, o.Payload)
}

func (o *GetIPLoadbalancingServiceNameVrackNetworkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetIPLoadbalancingServiceNameVrackNetworkDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}