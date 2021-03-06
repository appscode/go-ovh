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

// PostDomainZoneZoneNameTerminateReader is a Reader for the PostDomainZoneZoneNameTerminate structure.
type PostDomainZoneZoneNameTerminateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDomainZoneZoneNameTerminateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostDomainZoneZoneNameTerminateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostDomainZoneZoneNameTerminateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostDomainZoneZoneNameTerminateOK creates a PostDomainZoneZoneNameTerminateOK with default headers values
func NewPostDomainZoneZoneNameTerminateOK() *PostDomainZoneZoneNameTerminateOK {
	return &PostDomainZoneZoneNameTerminateOK{}
}

/*PostDomainZoneZoneNameTerminateOK handles this case with default header values.

return value
*/
type PostDomainZoneZoneNameTerminateOK struct {
	Payload string
}

func (o *PostDomainZoneZoneNameTerminateOK) Error() string {
	return fmt.Sprintf("[POST /domain/zone/{zoneName}/terminate][%d] postDomainZoneZoneNameTerminateOK  %+v", 200, o.Payload)
}

func (o *PostDomainZoneZoneNameTerminateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDomainZoneZoneNameTerminateDefault creates a PostDomainZoneZoneNameTerminateDefault with default headers values
func NewPostDomainZoneZoneNameTerminateDefault(code int) *PostDomainZoneZoneNameTerminateDefault {
	return &PostDomainZoneZoneNameTerminateDefault{
		_statusCode: code,
	}
}

/*PostDomainZoneZoneNameTerminateDefault handles this case with default header values.

Unexpected error
*/
type PostDomainZoneZoneNameTerminateDefault struct {
	_statusCode int

	Payload *models.PostDomainZoneZoneNameTerminateDefaultBody
}

// Code gets the status code for the post domain zone zone name terminate default response
func (o *PostDomainZoneZoneNameTerminateDefault) Code() int {
	return o._statusCode
}

func (o *PostDomainZoneZoneNameTerminateDefault) Error() string {
	return fmt.Sprintf("[POST /domain/zone/{zoneName}/terminate][%d] PostDomainZoneZoneNameTerminate default  %+v", o._statusCode, o.Payload)
}

func (o *PostDomainZoneZoneNameTerminateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostDomainZoneZoneNameTerminateDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
