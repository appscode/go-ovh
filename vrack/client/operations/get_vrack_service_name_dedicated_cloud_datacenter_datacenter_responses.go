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

	"github.com/appscode/go-ovh/vrack/models"
)

// GetVrackServiceNameDedicatedCloudDatacenterDatacenterReader is a Reader for the GetVrackServiceNameDedicatedCloudDatacenterDatacenter structure.
type GetVrackServiceNameDedicatedCloudDatacenterDatacenterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVrackServiceNameDedicatedCloudDatacenterDatacenterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVrackServiceNameDedicatedCloudDatacenterDatacenterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetVrackServiceNameDedicatedCloudDatacenterDatacenterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVrackServiceNameDedicatedCloudDatacenterDatacenterOK creates a GetVrackServiceNameDedicatedCloudDatacenterDatacenterOK with default headers values
func NewGetVrackServiceNameDedicatedCloudDatacenterDatacenterOK() *GetVrackServiceNameDedicatedCloudDatacenterDatacenterOK {
	return &GetVrackServiceNameDedicatedCloudDatacenterDatacenterOK{}
}

/*GetVrackServiceNameDedicatedCloudDatacenterDatacenterOK handles this case with default header values.

description of 'vrack.PccDatacenter' response
*/
type GetVrackServiceNameDedicatedCloudDatacenterDatacenterOK struct {
	Payload *models.VrackPccDatacenter
}

func (o *GetVrackServiceNameDedicatedCloudDatacenterDatacenterOK) Error() string {
	return fmt.Sprintf("[GET /vrack/{serviceName}/dedicatedCloudDatacenter/{datacenter}][%d] getVrackServiceNameDedicatedCloudDatacenterDatacenterOK  %+v", 200, o.Payload)
}

func (o *GetVrackServiceNameDedicatedCloudDatacenterDatacenterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VrackPccDatacenter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVrackServiceNameDedicatedCloudDatacenterDatacenterDefault creates a GetVrackServiceNameDedicatedCloudDatacenterDatacenterDefault with default headers values
func NewGetVrackServiceNameDedicatedCloudDatacenterDatacenterDefault(code int) *GetVrackServiceNameDedicatedCloudDatacenterDatacenterDefault {
	return &GetVrackServiceNameDedicatedCloudDatacenterDatacenterDefault{
		_statusCode: code,
	}
}

/*GetVrackServiceNameDedicatedCloudDatacenterDatacenterDefault handles this case with default header values.

Unexpected error
*/
type GetVrackServiceNameDedicatedCloudDatacenterDatacenterDefault struct {
	_statusCode int

	Payload *models.GetVrackServiceNameDedicatedCloudDatacenterDatacenterDefaultBody
}

// Code gets the status code for the get vrack service name dedicated cloud datacenter datacenter default response
func (o *GetVrackServiceNameDedicatedCloudDatacenterDatacenterDefault) Code() int {
	return o._statusCode
}

func (o *GetVrackServiceNameDedicatedCloudDatacenterDatacenterDefault) Error() string {
	return fmt.Sprintf("[GET /vrack/{serviceName}/dedicatedCloudDatacenter/{datacenter}][%d] GetVrackServiceNameDedicatedCloudDatacenterDatacenter default  %+v", o._statusCode, o.Payload)
}

func (o *GetVrackServiceNameDedicatedCloudDatacenterDatacenterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetVrackServiceNameDedicatedCloudDatacenterDatacenterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
