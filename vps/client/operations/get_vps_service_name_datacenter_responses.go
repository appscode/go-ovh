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

	"github.com/appscode/go-ovh/vps/models"
)

// GetVpsServiceNameDatacenterReader is a Reader for the GetVpsServiceNameDatacenter structure.
type GetVpsServiceNameDatacenterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVpsServiceNameDatacenterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVpsServiceNameDatacenterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetVpsServiceNameDatacenterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVpsServiceNameDatacenterOK creates a GetVpsServiceNameDatacenterOK with default headers values
func NewGetVpsServiceNameDatacenterOK() *GetVpsServiceNameDatacenterOK {
	return &GetVpsServiceNameDatacenterOK{}
}

/*GetVpsServiceNameDatacenterOK handles this case with default header values.

description of 'vps.Datacenter' response
*/
type GetVpsServiceNameDatacenterOK struct {
	Payload *models.VpsDatacenter
}

func (o *GetVpsServiceNameDatacenterOK) Error() string {
	return fmt.Sprintf("[GET /vps/{serviceName}/datacenter][%d] getVpsServiceNameDatacenterOK  %+v", 200, o.Payload)
}

func (o *GetVpsServiceNameDatacenterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VpsDatacenter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVpsServiceNameDatacenterDefault creates a GetVpsServiceNameDatacenterDefault with default headers values
func NewGetVpsServiceNameDatacenterDefault(code int) *GetVpsServiceNameDatacenterDefault {
	return &GetVpsServiceNameDatacenterDefault{
		_statusCode: code,
	}
}

/*GetVpsServiceNameDatacenterDefault handles this case with default header values.

Unexpected error
*/
type GetVpsServiceNameDatacenterDefault struct {
	_statusCode int

	Payload *models.GetVpsServiceNameDatacenterDefaultBody
}

// Code gets the status code for the get vps service name datacenter default response
func (o *GetVpsServiceNameDatacenterDefault) Code() int {
	return o._statusCode
}

func (o *GetVpsServiceNameDatacenterDefault) Error() string {
	return fmt.Sprintf("[GET /vps/{serviceName}/datacenter][%d] GetVpsServiceNameDatacenter default  %+v", o._statusCode, o.Payload)
}

func (o *GetVpsServiceNameDatacenterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetVpsServiceNameDatacenterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
