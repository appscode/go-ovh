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

// PostDomainServiceNameTaskIDRelaunchReader is a Reader for the PostDomainServiceNameTaskIDRelaunch structure.
type PostDomainServiceNameTaskIDRelaunchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDomainServiceNameTaskIDRelaunchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostDomainServiceNameTaskIDRelaunchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostDomainServiceNameTaskIDRelaunchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostDomainServiceNameTaskIDRelaunchOK creates a PostDomainServiceNameTaskIDRelaunchOK with default headers values
func NewPostDomainServiceNameTaskIDRelaunchOK() *PostDomainServiceNameTaskIDRelaunchOK {
	return &PostDomainServiceNameTaskIDRelaunchOK{}
}

/*PostDomainServiceNameTaskIDRelaunchOK handles this case with default header values.

return 'void'
*/
type PostDomainServiceNameTaskIDRelaunchOK struct {
}

func (o *PostDomainServiceNameTaskIDRelaunchOK) Error() string {
	return fmt.Sprintf("[POST /domain/{serviceName}/task/{id}/relaunch][%d] postDomainServiceNameTaskIdRelaunchOK ", 200)
}

func (o *PostDomainServiceNameTaskIDRelaunchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDomainServiceNameTaskIDRelaunchDefault creates a PostDomainServiceNameTaskIDRelaunchDefault with default headers values
func NewPostDomainServiceNameTaskIDRelaunchDefault(code int) *PostDomainServiceNameTaskIDRelaunchDefault {
	return &PostDomainServiceNameTaskIDRelaunchDefault{
		_statusCode: code,
	}
}

/*PostDomainServiceNameTaskIDRelaunchDefault handles this case with default header values.

Unexpected error
*/
type PostDomainServiceNameTaskIDRelaunchDefault struct {
	_statusCode int

	Payload *models.PostDomainServiceNameTaskIDRelaunchDefaultBody
}

// Code gets the status code for the post domain service name task ID relaunch default response
func (o *PostDomainServiceNameTaskIDRelaunchDefault) Code() int {
	return o._statusCode
}

func (o *PostDomainServiceNameTaskIDRelaunchDefault) Error() string {
	return fmt.Sprintf("[POST /domain/{serviceName}/task/{id}/relaunch][%d] PostDomainServiceNameTaskIDRelaunch default  %+v", o._statusCode, o.Payload)
}

func (o *PostDomainServiceNameTaskIDRelaunchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostDomainServiceNameTaskIDRelaunchDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}