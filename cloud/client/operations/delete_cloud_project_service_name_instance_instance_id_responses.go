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

	"github.com/appscode/go-ovh/cloud/models"
)

// DeleteCloudProjectServiceNameInstanceInstanceIDReader is a Reader for the DeleteCloudProjectServiceNameInstanceInstanceID structure.
type DeleteCloudProjectServiceNameInstanceInstanceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCloudProjectServiceNameInstanceInstanceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteCloudProjectServiceNameInstanceInstanceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteCloudProjectServiceNameInstanceInstanceIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCloudProjectServiceNameInstanceInstanceIDOK creates a DeleteCloudProjectServiceNameInstanceInstanceIDOK with default headers values
func NewDeleteCloudProjectServiceNameInstanceInstanceIDOK() *DeleteCloudProjectServiceNameInstanceInstanceIDOK {
	return &DeleteCloudProjectServiceNameInstanceInstanceIDOK{}
}

/*DeleteCloudProjectServiceNameInstanceInstanceIDOK handles this case with default header values.

return 'void'
*/
type DeleteCloudProjectServiceNameInstanceInstanceIDOK struct {
}

func (o *DeleteCloudProjectServiceNameInstanceInstanceIDOK) Error() string {
	return fmt.Sprintf("[DELETE /cloud/project/{serviceName}/instance/{instanceId}][%d] deleteCloudProjectServiceNameInstanceInstanceIdOK ", 200)
}

func (o *DeleteCloudProjectServiceNameInstanceInstanceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCloudProjectServiceNameInstanceInstanceIDDefault creates a DeleteCloudProjectServiceNameInstanceInstanceIDDefault with default headers values
func NewDeleteCloudProjectServiceNameInstanceInstanceIDDefault(code int) *DeleteCloudProjectServiceNameInstanceInstanceIDDefault {
	return &DeleteCloudProjectServiceNameInstanceInstanceIDDefault{
		_statusCode: code,
	}
}

/*DeleteCloudProjectServiceNameInstanceInstanceIDDefault handles this case with default header values.

Unexpected error
*/
type DeleteCloudProjectServiceNameInstanceInstanceIDDefault struct {
	_statusCode int

	Payload *models.DeleteCloudProjectServiceNameInstanceInstanceIDDefaultBody
}

// Code gets the status code for the delete cloud project service name instance instance ID default response
func (o *DeleteCloudProjectServiceNameInstanceInstanceIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCloudProjectServiceNameInstanceInstanceIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /cloud/project/{serviceName}/instance/{instanceId}][%d] DeleteCloudProjectServiceNameInstanceInstanceID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCloudProjectServiceNameInstanceInstanceIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteCloudProjectServiceNameInstanceInstanceIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
