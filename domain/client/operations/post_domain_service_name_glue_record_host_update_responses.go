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

// PostDomainServiceNameGlueRecordHostUpdateReader is a Reader for the PostDomainServiceNameGlueRecordHostUpdate structure.
type PostDomainServiceNameGlueRecordHostUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDomainServiceNameGlueRecordHostUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostDomainServiceNameGlueRecordHostUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostDomainServiceNameGlueRecordHostUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostDomainServiceNameGlueRecordHostUpdateOK creates a PostDomainServiceNameGlueRecordHostUpdateOK with default headers values
func NewPostDomainServiceNameGlueRecordHostUpdateOK() *PostDomainServiceNameGlueRecordHostUpdateOK {
	return &PostDomainServiceNameGlueRecordHostUpdateOK{}
}

/*PostDomainServiceNameGlueRecordHostUpdateOK handles this case with default header values.

description of 'domain.Task' response
*/
type PostDomainServiceNameGlueRecordHostUpdateOK struct {
	Payload *models.DomainTask
}

func (o *PostDomainServiceNameGlueRecordHostUpdateOK) Error() string {
	return fmt.Sprintf("[POST /domain/{serviceName}/glueRecord/{host}/update][%d] postDomainServiceNameGlueRecordHostUpdateOK  %+v", 200, o.Payload)
}

func (o *PostDomainServiceNameGlueRecordHostUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDomainServiceNameGlueRecordHostUpdateDefault creates a PostDomainServiceNameGlueRecordHostUpdateDefault with default headers values
func NewPostDomainServiceNameGlueRecordHostUpdateDefault(code int) *PostDomainServiceNameGlueRecordHostUpdateDefault {
	return &PostDomainServiceNameGlueRecordHostUpdateDefault{
		_statusCode: code,
	}
}

/*PostDomainServiceNameGlueRecordHostUpdateDefault handles this case with default header values.

Unexpected error
*/
type PostDomainServiceNameGlueRecordHostUpdateDefault struct {
	_statusCode int

	Payload *models.PostDomainServiceNameGlueRecordHostUpdateDefaultBody
}

// Code gets the status code for the post domain service name glue record host update default response
func (o *PostDomainServiceNameGlueRecordHostUpdateDefault) Code() int {
	return o._statusCode
}

func (o *PostDomainServiceNameGlueRecordHostUpdateDefault) Error() string {
	return fmt.Sprintf("[POST /domain/{serviceName}/glueRecord/{host}/update][%d] PostDomainServiceNameGlueRecordHostUpdate default  %+v", o._statusCode, o.Payload)
}

func (o *PostDomainServiceNameGlueRecordHostUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostDomainServiceNameGlueRecordHostUpdateDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
