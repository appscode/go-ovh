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

// GetVrackServiceNameDedicatedServerReader is a Reader for the GetVrackServiceNameDedicatedServer structure.
type GetVrackServiceNameDedicatedServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVrackServiceNameDedicatedServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVrackServiceNameDedicatedServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetVrackServiceNameDedicatedServerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVrackServiceNameDedicatedServerOK creates a GetVrackServiceNameDedicatedServerOK with default headers values
func NewGetVrackServiceNameDedicatedServerOK() *GetVrackServiceNameDedicatedServerOK {
	return &GetVrackServiceNameDedicatedServerOK{}
}

/*GetVrackServiceNameDedicatedServerOK handles this case with default header values.

return value
*/
type GetVrackServiceNameDedicatedServerOK struct {
	Payload []string
}

func (o *GetVrackServiceNameDedicatedServerOK) Error() string {
	return fmt.Sprintf("[GET /vrack/{serviceName}/dedicatedServer][%d] getVrackServiceNameDedicatedServerOK  %+v", 200, o.Payload)
}

func (o *GetVrackServiceNameDedicatedServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVrackServiceNameDedicatedServerDefault creates a GetVrackServiceNameDedicatedServerDefault with default headers values
func NewGetVrackServiceNameDedicatedServerDefault(code int) *GetVrackServiceNameDedicatedServerDefault {
	return &GetVrackServiceNameDedicatedServerDefault{
		_statusCode: code,
	}
}

/*GetVrackServiceNameDedicatedServerDefault handles this case with default header values.

Unexpected error
*/
type GetVrackServiceNameDedicatedServerDefault struct {
	_statusCode int

	Payload *models.GetVrackServiceNameDedicatedServerDefaultBody
}

// Code gets the status code for the get vrack service name dedicated server default response
func (o *GetVrackServiceNameDedicatedServerDefault) Code() int {
	return o._statusCode
}

func (o *GetVrackServiceNameDedicatedServerDefault) Error() string {
	return fmt.Sprintf("[GET /vrack/{serviceName}/dedicatedServer][%d] GetVrackServiceNameDedicatedServer default  %+v", o._statusCode, o.Payload)
}

func (o *GetVrackServiceNameDedicatedServerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetVrackServiceNameDedicatedServerDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
