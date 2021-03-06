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

// GetDomainZoneZoneNameStatusReader is a Reader for the GetDomainZoneZoneNameStatus structure.
type GetDomainZoneZoneNameStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomainZoneZoneNameStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDomainZoneZoneNameStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDomainZoneZoneNameStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDomainZoneZoneNameStatusOK creates a GetDomainZoneZoneNameStatusOK with default headers values
func NewGetDomainZoneZoneNameStatusOK() *GetDomainZoneZoneNameStatusOK {
	return &GetDomainZoneZoneNameStatusOK{}
}

/*GetDomainZoneZoneNameStatusOK handles this case with default header values.

description of 'zone.Status' response
*/
type GetDomainZoneZoneNameStatusOK struct {
	Payload *models.ZoneStatus
}

func (o *GetDomainZoneZoneNameStatusOK) Error() string {
	return fmt.Sprintf("[GET /domain/zone/{zoneName}/status][%d] getDomainZoneZoneNameStatusOK  %+v", 200, o.Payload)
}

func (o *GetDomainZoneZoneNameStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZoneStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainZoneZoneNameStatusDefault creates a GetDomainZoneZoneNameStatusDefault with default headers values
func NewGetDomainZoneZoneNameStatusDefault(code int) *GetDomainZoneZoneNameStatusDefault {
	return &GetDomainZoneZoneNameStatusDefault{
		_statusCode: code,
	}
}

/*GetDomainZoneZoneNameStatusDefault handles this case with default header values.

Unexpected error
*/
type GetDomainZoneZoneNameStatusDefault struct {
	_statusCode int

	Payload *models.GetDomainZoneZoneNameStatusDefaultBody
}

// Code gets the status code for the get domain zone zone name status default response
func (o *GetDomainZoneZoneNameStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetDomainZoneZoneNameStatusDefault) Error() string {
	return fmt.Sprintf("[GET /domain/zone/{zoneName}/status][%d] GetDomainZoneZoneNameStatus default  %+v", o._statusCode, o.Payload)
}

func (o *GetDomainZoneZoneNameStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetDomainZoneZoneNameStatusDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
