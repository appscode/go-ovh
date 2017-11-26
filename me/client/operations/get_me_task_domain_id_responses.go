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

	"github.com/appscode/go-ovh/me/models"
)

// GetMeTaskDomainIDReader is a Reader for the GetMeTaskDomainID structure.
type GetMeTaskDomainIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeTaskDomainIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeTaskDomainIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeTaskDomainIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeTaskDomainIDOK creates a GetMeTaskDomainIDOK with default headers values
func NewGetMeTaskDomainIDOK() *GetMeTaskDomainIDOK {
	return &GetMeTaskDomainIDOK{}
}

/*GetMeTaskDomainIDOK handles this case with default header values.

description of 'nichandle.DomainTask' response
*/
type GetMeTaskDomainIDOK struct {
	Payload *models.NichandleDomainTask
}

func (o *GetMeTaskDomainIDOK) Error() string {
	return fmt.Sprintf("[GET /me/task/domain/{id}][%d] getMeTaskDomainIdOK  %+v", 200, o.Payload)
}

func (o *GetMeTaskDomainIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NichandleDomainTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeTaskDomainIDDefault creates a GetMeTaskDomainIDDefault with default headers values
func NewGetMeTaskDomainIDDefault(code int) *GetMeTaskDomainIDDefault {
	return &GetMeTaskDomainIDDefault{
		_statusCode: code,
	}
}

/*GetMeTaskDomainIDDefault handles this case with default header values.

Unexpected error
*/
type GetMeTaskDomainIDDefault struct {
	_statusCode int

	Payload *models.GetMeTaskDomainIDDefaultBody
}

// Code gets the status code for the get me task domain ID default response
func (o *GetMeTaskDomainIDDefault) Code() int {
	return o._statusCode
}

func (o *GetMeTaskDomainIDDefault) Error() string {
	return fmt.Sprintf("[GET /me/task/domain/{id}][%d] GetMeTaskDomainID default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeTaskDomainIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeTaskDomainIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
