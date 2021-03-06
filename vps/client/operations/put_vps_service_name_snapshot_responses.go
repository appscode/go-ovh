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

// PutVpsServiceNameSnapshotReader is a Reader for the PutVpsServiceNameSnapshot structure.
type PutVpsServiceNameSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutVpsServiceNameSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutVpsServiceNameSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutVpsServiceNameSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutVpsServiceNameSnapshotOK creates a PutVpsServiceNameSnapshotOK with default headers values
func NewPutVpsServiceNameSnapshotOK() *PutVpsServiceNameSnapshotOK {
	return &PutVpsServiceNameSnapshotOK{}
}

/*PutVpsServiceNameSnapshotOK handles this case with default header values.

return 'void'
*/
type PutVpsServiceNameSnapshotOK struct {
}

func (o *PutVpsServiceNameSnapshotOK) Error() string {
	return fmt.Sprintf("[PUT /vps/{serviceName}/snapshot][%d] putVpsServiceNameSnapshotOK ", 200)
}

func (o *PutVpsServiceNameSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutVpsServiceNameSnapshotDefault creates a PutVpsServiceNameSnapshotDefault with default headers values
func NewPutVpsServiceNameSnapshotDefault(code int) *PutVpsServiceNameSnapshotDefault {
	return &PutVpsServiceNameSnapshotDefault{
		_statusCode: code,
	}
}

/*PutVpsServiceNameSnapshotDefault handles this case with default header values.

Unexpected error
*/
type PutVpsServiceNameSnapshotDefault struct {
	_statusCode int

	Payload *models.PutVpsServiceNameSnapshotDefaultBody
}

// Code gets the status code for the put vps service name snapshot default response
func (o *PutVpsServiceNameSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *PutVpsServiceNameSnapshotDefault) Error() string {
	return fmt.Sprintf("[PUT /vps/{serviceName}/snapshot][%d] PutVpsServiceNameSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *PutVpsServiceNameSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutVpsServiceNameSnapshotDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
