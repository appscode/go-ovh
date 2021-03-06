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

// DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceReader is a Reader for the DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterface structure.
type DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceOK creates a DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceOK with default headers values
func NewDeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceOK() *DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceOK {
	return &DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceOK{}
}

/*DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceOK handles this case with default header values.

description of 'vrack.Task' response
*/
type DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceOK struct {
	Payload *models.VrackTask
}

func (o *DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceOK) Error() string {
	return fmt.Sprintf("[DELETE /vrack/{serviceName}/dedicatedServerInterface/{dedicatedServerInterface}][%d] deleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceOK  %+v", 200, o.Payload)
}

func (o *DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VrackTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceDefault creates a DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceDefault with default headers values
func NewDeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceDefault(code int) *DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceDefault {
	return &DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceDefault{
		_statusCode: code,
	}
}

/*DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceDefault handles this case with default header values.

Unexpected error
*/
type DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceDefault struct {
	_statusCode int

	Payload *models.DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceDefaultBody
}

// Code gets the status code for the delete vrack service name dedicated server interface dedicated server interface default response
func (o *DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceDefault) Code() int {
	return o._statusCode
}

func (o *DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceDefault) Error() string {
	return fmt.Sprintf("[DELETE /vrack/{serviceName}/dedicatedServerInterface/{dedicatedServerInterface}][%d] DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterface default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteVrackServiceNameDedicatedServerInterfaceDedicatedServerInterfaceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
