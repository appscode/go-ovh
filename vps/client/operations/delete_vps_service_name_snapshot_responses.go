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

// DeleteVpsServiceNameSnapshotReader is a Reader for the DeleteVpsServiceNameSnapshot structure.
type DeleteVpsServiceNameSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVpsServiceNameSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteVpsServiceNameSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteVpsServiceNameSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteVpsServiceNameSnapshotOK creates a DeleteVpsServiceNameSnapshotOK with default headers values
func NewDeleteVpsServiceNameSnapshotOK() *DeleteVpsServiceNameSnapshotOK {
	return &DeleteVpsServiceNameSnapshotOK{}
}

/*DeleteVpsServiceNameSnapshotOK handles this case with default header values.

description of 'vps.Task' response
*/
type DeleteVpsServiceNameSnapshotOK struct {
	Payload *models.VpsTask
}

func (o *DeleteVpsServiceNameSnapshotOK) Error() string {
	return fmt.Sprintf("[DELETE /vps/{serviceName}/snapshot][%d] deleteVpsServiceNameSnapshotOK  %+v", 200, o.Payload)
}

func (o *DeleteVpsServiceNameSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VpsTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVpsServiceNameSnapshotDefault creates a DeleteVpsServiceNameSnapshotDefault with default headers values
func NewDeleteVpsServiceNameSnapshotDefault(code int) *DeleteVpsServiceNameSnapshotDefault {
	return &DeleteVpsServiceNameSnapshotDefault{
		_statusCode: code,
	}
}

/*DeleteVpsServiceNameSnapshotDefault handles this case with default header values.

Unexpected error
*/
type DeleteVpsServiceNameSnapshotDefault struct {
	_statusCode int

	Payload *models.DeleteVpsServiceNameSnapshotDefaultBody
}

// Code gets the status code for the delete vps service name snapshot default response
func (o *DeleteVpsServiceNameSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *DeleteVpsServiceNameSnapshotDefault) Error() string {
	return fmt.Sprintf("[DELETE /vps/{serviceName}/snapshot][%d] DeleteVpsServiceNameSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteVpsServiceNameSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteVpsServiceNameSnapshotDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
