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

// GetIPLoadbalancingServiceNameVrackNetworkCreationRulesReader is a Reader for the GetIPLoadbalancingServiceNameVrackNetworkCreationRules structure.
type GetIPLoadbalancingServiceNameVrackNetworkCreationRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPLoadbalancingServiceNameVrackNetworkCreationRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIPLoadbalancingServiceNameVrackNetworkCreationRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetIPLoadbalancingServiceNameVrackNetworkCreationRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIPLoadbalancingServiceNameVrackNetworkCreationRulesOK creates a GetIPLoadbalancingServiceNameVrackNetworkCreationRulesOK with default headers values
func NewGetIPLoadbalancingServiceNameVrackNetworkCreationRulesOK() *GetIPLoadbalancingServiceNameVrackNetworkCreationRulesOK {
	return &GetIPLoadbalancingServiceNameVrackNetworkCreationRulesOK{}
}

/*GetIPLoadbalancingServiceNameVrackNetworkCreationRulesOK handles this case with default header values.

description of 'iplb.VrackNetworkCreationRules' response
*/
type GetIPLoadbalancingServiceNameVrackNetworkCreationRulesOK struct {
	Payload *models.IPLBVrackNetworkCreationRules
}

func (o *GetIPLoadbalancingServiceNameVrackNetworkCreationRulesOK) Error() string {
	return fmt.Sprintf("[GET /ipLoadbalancing/{serviceName}/vrack/networkCreationRules][%d] getIpLoadbalancingServiceNameVrackNetworkCreationRulesOK  %+v", 200, o.Payload)
}

func (o *GetIPLoadbalancingServiceNameVrackNetworkCreationRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPLBVrackNetworkCreationRules)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPLoadbalancingServiceNameVrackNetworkCreationRulesDefault creates a GetIPLoadbalancingServiceNameVrackNetworkCreationRulesDefault with default headers values
func NewGetIPLoadbalancingServiceNameVrackNetworkCreationRulesDefault(code int) *GetIPLoadbalancingServiceNameVrackNetworkCreationRulesDefault {
	return &GetIPLoadbalancingServiceNameVrackNetworkCreationRulesDefault{
		_statusCode: code,
	}
}

/*GetIPLoadbalancingServiceNameVrackNetworkCreationRulesDefault handles this case with default header values.

Unexpected error
*/
type GetIPLoadbalancingServiceNameVrackNetworkCreationRulesDefault struct {
	_statusCode int

	Payload *models.GetIPLoadbalancingServiceNameVrackNetworkCreationRulesDefaultBody
}

// Code gets the status code for the get IP loadbalancing service name vrack network creation rules default response
func (o *GetIPLoadbalancingServiceNameVrackNetworkCreationRulesDefault) Code() int {
	return o._statusCode
}

func (o *GetIPLoadbalancingServiceNameVrackNetworkCreationRulesDefault) Error() string {
	return fmt.Sprintf("[GET /ipLoadbalancing/{serviceName}/vrack/networkCreationRules][%d] GetIPLoadbalancingServiceNameVrackNetworkCreationRules default  %+v", o._statusCode, o.Payload)
}

func (o *GetIPLoadbalancingServiceNameVrackNetworkCreationRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetIPLoadbalancingServiceNameVrackNetworkCreationRulesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}