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

	"github.com/appscode/go-ovh/vrack/models"
)

// GetVrackServiceNameServiceInfosReader is a Reader for the GetVrackServiceNameServiceInfos structure.
type GetVrackServiceNameServiceInfosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVrackServiceNameServiceInfosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVrackServiceNameServiceInfosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetVrackServiceNameServiceInfosDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVrackServiceNameServiceInfosOK creates a GetVrackServiceNameServiceInfosOK with default headers values
func NewGetVrackServiceNameServiceInfosOK() *GetVrackServiceNameServiceInfosOK {
	return &GetVrackServiceNameServiceInfosOK{}
}

/*GetVrackServiceNameServiceInfosOK handles this case with default header values.

description of 'services.NonExpiringService' response
*/
type GetVrackServiceNameServiceInfosOK struct {
	Payload *models.ServicesNonExpiringService
}

func (o *GetVrackServiceNameServiceInfosOK) Error() string {
	return fmt.Sprintf("[GET /vrack/{serviceName}/serviceInfos][%d] getVrackServiceNameServiceInfosOK  %+v", 200, o.Payload)
}

func (o *GetVrackServiceNameServiceInfosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServicesNonExpiringService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVrackServiceNameServiceInfosDefault creates a GetVrackServiceNameServiceInfosDefault with default headers values
func NewGetVrackServiceNameServiceInfosDefault(code int) *GetVrackServiceNameServiceInfosDefault {
	return &GetVrackServiceNameServiceInfosDefault{
		_statusCode: code,
	}
}

/*GetVrackServiceNameServiceInfosDefault handles this case with default header values.

Unexpected error
*/
type GetVrackServiceNameServiceInfosDefault struct {
	_statusCode int

	Payload *models.GetVrackServiceNameServiceInfosDefaultBody
}

// Code gets the status code for the get vrack service name service infos default response
func (o *GetVrackServiceNameServiceInfosDefault) Code() int {
	return o._statusCode
}

func (o *GetVrackServiceNameServiceInfosDefault) Error() string {
	return fmt.Sprintf("[GET /vrack/{serviceName}/serviceInfos][%d] GetVrackServiceNameServiceInfos default  %+v", o._statusCode, o.Payload)
}

func (o *GetVrackServiceNameServiceInfosDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetVrackServiceNameServiceInfosDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
