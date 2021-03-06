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

// GetDomainServiceNameGlueRecordReader is a Reader for the GetDomainServiceNameGlueRecord structure.
type GetDomainServiceNameGlueRecordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomainServiceNameGlueRecordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDomainServiceNameGlueRecordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDomainServiceNameGlueRecordDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDomainServiceNameGlueRecordOK creates a GetDomainServiceNameGlueRecordOK with default headers values
func NewGetDomainServiceNameGlueRecordOK() *GetDomainServiceNameGlueRecordOK {
	return &GetDomainServiceNameGlueRecordOK{}
}

/*GetDomainServiceNameGlueRecordOK handles this case with default header values.

return value
*/
type GetDomainServiceNameGlueRecordOK struct {
	Payload []string
}

func (o *GetDomainServiceNameGlueRecordOK) Error() string {
	return fmt.Sprintf("[GET /domain/{serviceName}/glueRecord][%d] getDomainServiceNameGlueRecordOK  %+v", 200, o.Payload)
}

func (o *GetDomainServiceNameGlueRecordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainServiceNameGlueRecordDefault creates a GetDomainServiceNameGlueRecordDefault with default headers values
func NewGetDomainServiceNameGlueRecordDefault(code int) *GetDomainServiceNameGlueRecordDefault {
	return &GetDomainServiceNameGlueRecordDefault{
		_statusCode: code,
	}
}

/*GetDomainServiceNameGlueRecordDefault handles this case with default header values.

Unexpected error
*/
type GetDomainServiceNameGlueRecordDefault struct {
	_statusCode int

	Payload *models.GetDomainServiceNameGlueRecordDefaultBody
}

// Code gets the status code for the get domain service name glue record default response
func (o *GetDomainServiceNameGlueRecordDefault) Code() int {
	return o._statusCode
}

func (o *GetDomainServiceNameGlueRecordDefault) Error() string {
	return fmt.Sprintf("[GET /domain/{serviceName}/glueRecord][%d] GetDomainServiceNameGlueRecord default  %+v", o._statusCode, o.Payload)
}

func (o *GetDomainServiceNameGlueRecordDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetDomainServiceNameGlueRecordDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
