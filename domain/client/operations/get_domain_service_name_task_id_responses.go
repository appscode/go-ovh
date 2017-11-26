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

	"github.com/appscode/go-ovh/domain/models"
)

// GetDomainServiceNameTaskIDReader is a Reader for the GetDomainServiceNameTaskID structure.
type GetDomainServiceNameTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomainServiceNameTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDomainServiceNameTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDomainServiceNameTaskIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDomainServiceNameTaskIDOK creates a GetDomainServiceNameTaskIDOK with default headers values
func NewGetDomainServiceNameTaskIDOK() *GetDomainServiceNameTaskIDOK {
	return &GetDomainServiceNameTaskIDOK{}
}

/*GetDomainServiceNameTaskIDOK handles this case with default header values.

description of 'domain.Task' response
*/
type GetDomainServiceNameTaskIDOK struct {
	Payload *models.DomainTask
}

func (o *GetDomainServiceNameTaskIDOK) Error() string {
	return fmt.Sprintf("[GET /domain/{serviceName}/task/{id}][%d] getDomainServiceNameTaskIdOK  %+v", 200, o.Payload)
}

func (o *GetDomainServiceNameTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainServiceNameTaskIDDefault creates a GetDomainServiceNameTaskIDDefault with default headers values
func NewGetDomainServiceNameTaskIDDefault(code int) *GetDomainServiceNameTaskIDDefault {
	return &GetDomainServiceNameTaskIDDefault{
		_statusCode: code,
	}
}

/*GetDomainServiceNameTaskIDDefault handles this case with default header values.

Unexpected error
*/
type GetDomainServiceNameTaskIDDefault struct {
	_statusCode int

	Payload *models.GetDomainServiceNameTaskIDDefaultBody
}

// Code gets the status code for the get domain service name task ID default response
func (o *GetDomainServiceNameTaskIDDefault) Code() int {
	return o._statusCode
}

func (o *GetDomainServiceNameTaskIDDefault) Error() string {
	return fmt.Sprintf("[GET /domain/{serviceName}/task/{id}][%d] GetDomainServiceNameTaskID default  %+v", o._statusCode, o.Payload)
}

func (o *GetDomainServiceNameTaskIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetDomainServiceNameTaskIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}