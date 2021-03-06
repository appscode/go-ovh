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

// GetDomainZoneReader is a Reader for the GetDomainZone structure.
type GetDomainZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomainZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDomainZoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDomainZoneDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDomainZoneOK creates a GetDomainZoneOK with default headers values
func NewGetDomainZoneOK() *GetDomainZoneOK {
	return &GetDomainZoneOK{}
}

/*GetDomainZoneOK handles this case with default header values.

return value
*/
type GetDomainZoneOK struct {
	Payload []string
}

func (o *GetDomainZoneOK) Error() string {
	return fmt.Sprintf("[GET /domain/zone][%d] getDomainZoneOK  %+v", 200, o.Payload)
}

func (o *GetDomainZoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainZoneDefault creates a GetDomainZoneDefault with default headers values
func NewGetDomainZoneDefault(code int) *GetDomainZoneDefault {
	return &GetDomainZoneDefault{
		_statusCode: code,
	}
}

/*GetDomainZoneDefault handles this case with default header values.

Unexpected error
*/
type GetDomainZoneDefault struct {
	_statusCode int

	Payload *models.GetDomainZoneDefaultBody
}

// Code gets the status code for the get domain zone default response
func (o *GetDomainZoneDefault) Code() int {
	return o._statusCode
}

func (o *GetDomainZoneDefault) Error() string {
	return fmt.Sprintf("[GET /domain/zone][%d] GetDomainZone default  %+v", o._statusCode, o.Payload)
}

func (o *GetDomainZoneDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetDomainZoneDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
