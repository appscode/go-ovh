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

// PostVpsServiceNameStartReader is a Reader for the PostVpsServiceNameStart structure.
type PostVpsServiceNameStartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVpsServiceNameStartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostVpsServiceNameStartOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostVpsServiceNameStartDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostVpsServiceNameStartOK creates a PostVpsServiceNameStartOK with default headers values
func NewPostVpsServiceNameStartOK() *PostVpsServiceNameStartOK {
	return &PostVpsServiceNameStartOK{}
}

/*PostVpsServiceNameStartOK handles this case with default header values.

description of 'vps.Task' response
*/
type PostVpsServiceNameStartOK struct {
	Payload *models.VpsTask
}

func (o *PostVpsServiceNameStartOK) Error() string {
	return fmt.Sprintf("[POST /vps/{serviceName}/start][%d] postVpsServiceNameStartOK  %+v", 200, o.Payload)
}

func (o *PostVpsServiceNameStartOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VpsTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVpsServiceNameStartDefault creates a PostVpsServiceNameStartDefault with default headers values
func NewPostVpsServiceNameStartDefault(code int) *PostVpsServiceNameStartDefault {
	return &PostVpsServiceNameStartDefault{
		_statusCode: code,
	}
}

/*PostVpsServiceNameStartDefault handles this case with default header values.

Unexpected error
*/
type PostVpsServiceNameStartDefault struct {
	_statusCode int

	Payload *models.PostVpsServiceNameStartDefaultBody
}

// Code gets the status code for the post vps service name start default response
func (o *PostVpsServiceNameStartDefault) Code() int {
	return o._statusCode
}

func (o *PostVpsServiceNameStartDefault) Error() string {
	return fmt.Sprintf("[POST /vps/{serviceName}/start][%d] PostVpsServiceNameStart default  %+v", o._statusCode, o.Payload)
}

func (o *PostVpsServiceNameStartDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostVpsServiceNameStartDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}