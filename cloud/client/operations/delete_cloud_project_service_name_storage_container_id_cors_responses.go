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

// DeleteCloudProjectServiceNameStorageContainerIDCorsReader is a Reader for the DeleteCloudProjectServiceNameStorageContainerIDCors structure.
type DeleteCloudProjectServiceNameStorageContainerIDCorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCloudProjectServiceNameStorageContainerIDCorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteCloudProjectServiceNameStorageContainerIDCorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteCloudProjectServiceNameStorageContainerIDCorsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCloudProjectServiceNameStorageContainerIDCorsOK creates a DeleteCloudProjectServiceNameStorageContainerIDCorsOK with default headers values
func NewDeleteCloudProjectServiceNameStorageContainerIDCorsOK() *DeleteCloudProjectServiceNameStorageContainerIDCorsOK {
	return &DeleteCloudProjectServiceNameStorageContainerIDCorsOK{}
}

/*DeleteCloudProjectServiceNameStorageContainerIDCorsOK handles this case with default header values.

return 'void'
*/
type DeleteCloudProjectServiceNameStorageContainerIDCorsOK struct {
}

func (o *DeleteCloudProjectServiceNameStorageContainerIDCorsOK) Error() string {
	return fmt.Sprintf("[DELETE /cloud/project/{serviceName}/storage/{containerId}/cors][%d] deleteCloudProjectServiceNameStorageContainerIdCorsOK ", 200)
}

func (o *DeleteCloudProjectServiceNameStorageContainerIDCorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCloudProjectServiceNameStorageContainerIDCorsDefault creates a DeleteCloudProjectServiceNameStorageContainerIDCorsDefault with default headers values
func NewDeleteCloudProjectServiceNameStorageContainerIDCorsDefault(code int) *DeleteCloudProjectServiceNameStorageContainerIDCorsDefault {
	return &DeleteCloudProjectServiceNameStorageContainerIDCorsDefault{
		_statusCode: code,
	}
}

/*DeleteCloudProjectServiceNameStorageContainerIDCorsDefault handles this case with default header values.

Unexpected error
*/
type DeleteCloudProjectServiceNameStorageContainerIDCorsDefault struct {
	_statusCode int

	Payload *models.DeleteCloudProjectServiceNameStorageContainerIDCorsDefaultBody
}

// Code gets the status code for the delete cloud project service name storage container ID cors default response
func (o *DeleteCloudProjectServiceNameStorageContainerIDCorsDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCloudProjectServiceNameStorageContainerIDCorsDefault) Error() string {
	return fmt.Sprintf("[DELETE /cloud/project/{serviceName}/storage/{containerId}/cors][%d] DeleteCloudProjectServiceNameStorageContainerIDCors default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCloudProjectServiceNameStorageContainerIDCorsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteCloudProjectServiceNameStorageContainerIDCorsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}