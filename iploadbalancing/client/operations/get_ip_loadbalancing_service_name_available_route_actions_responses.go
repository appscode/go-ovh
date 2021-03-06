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

// GetIPLoadbalancingServiceNameAvailableRouteActionsReader is a Reader for the GetIPLoadbalancingServiceNameAvailableRouteActions structure.
type GetIPLoadbalancingServiceNameAvailableRouteActionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPLoadbalancingServiceNameAvailableRouteActionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIPLoadbalancingServiceNameAvailableRouteActionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetIPLoadbalancingServiceNameAvailableRouteActionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIPLoadbalancingServiceNameAvailableRouteActionsOK creates a GetIPLoadbalancingServiceNameAvailableRouteActionsOK with default headers values
func NewGetIPLoadbalancingServiceNameAvailableRouteActionsOK() *GetIPLoadbalancingServiceNameAvailableRouteActionsOK {
	return &GetIPLoadbalancingServiceNameAvailableRouteActionsOK{}
}

/*GetIPLoadbalancingServiceNameAvailableRouteActionsOK handles this case with default header values.

return value
*/
type GetIPLoadbalancingServiceNameAvailableRouteActionsOK struct {
	Payload models.GetIPLoadbalancingServiceNameAvailableRouteActionsOKBody
}

func (o *GetIPLoadbalancingServiceNameAvailableRouteActionsOK) Error() string {
	return fmt.Sprintf("[GET /ipLoadbalancing/{serviceName}/availableRouteActions][%d] getIpLoadbalancingServiceNameAvailableRouteActionsOK  %+v", 200, o.Payload)
}

func (o *GetIPLoadbalancingServiceNameAvailableRouteActionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPLoadbalancingServiceNameAvailableRouteActionsDefault creates a GetIPLoadbalancingServiceNameAvailableRouteActionsDefault with default headers values
func NewGetIPLoadbalancingServiceNameAvailableRouteActionsDefault(code int) *GetIPLoadbalancingServiceNameAvailableRouteActionsDefault {
	return &GetIPLoadbalancingServiceNameAvailableRouteActionsDefault{
		_statusCode: code,
	}
}

/*GetIPLoadbalancingServiceNameAvailableRouteActionsDefault handles this case with default header values.

Unexpected error
*/
type GetIPLoadbalancingServiceNameAvailableRouteActionsDefault struct {
	_statusCode int

	Payload *models.GetIPLoadbalancingServiceNameAvailableRouteActionsDefaultBody
}

// Code gets the status code for the get IP loadbalancing service name available route actions default response
func (o *GetIPLoadbalancingServiceNameAvailableRouteActionsDefault) Code() int {
	return o._statusCode
}

func (o *GetIPLoadbalancingServiceNameAvailableRouteActionsDefault) Error() string {
	return fmt.Sprintf("[GET /ipLoadbalancing/{serviceName}/availableRouteActions][%d] GetIPLoadbalancingServiceNameAvailableRouteActions default  %+v", o._statusCode, o.Payload)
}

func (o *GetIPLoadbalancingServiceNameAvailableRouteActionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetIPLoadbalancingServiceNameAvailableRouteActionsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
