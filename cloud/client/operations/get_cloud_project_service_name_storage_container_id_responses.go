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

// GetCloudProjectServiceNameStorageContainerIDReader is a Reader for the GetCloudProjectServiceNameStorageContainerID structure.
type GetCloudProjectServiceNameStorageContainerIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudProjectServiceNameStorageContainerIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCloudProjectServiceNameStorageContainerIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetCloudProjectServiceNameStorageContainerIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCloudProjectServiceNameStorageContainerIDOK creates a GetCloudProjectServiceNameStorageContainerIDOK with default headers values
func NewGetCloudProjectServiceNameStorageContainerIDOK() *GetCloudProjectServiceNameStorageContainerIDOK {
	return &GetCloudProjectServiceNameStorageContainerIDOK{}
}

/*GetCloudProjectServiceNameStorageContainerIDOK handles this case with default header values.

description of 'cloud.Storage.ContainerDetail' response
*/
type GetCloudProjectServiceNameStorageContainerIDOK struct {
	Payload *models.CloudStorageContainerDetail
}

func (o *GetCloudProjectServiceNameStorageContainerIDOK) Error() string {
	return fmt.Sprintf("[GET /cloud/project/{serviceName}/storage/{containerId}][%d] getCloudProjectServiceNameStorageContainerIdOK  %+v", 200, o.Payload)
}

func (o *GetCloudProjectServiceNameStorageContainerIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudStorageContainerDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudProjectServiceNameStorageContainerIDDefault creates a GetCloudProjectServiceNameStorageContainerIDDefault with default headers values
func NewGetCloudProjectServiceNameStorageContainerIDDefault(code int) *GetCloudProjectServiceNameStorageContainerIDDefault {
	return &GetCloudProjectServiceNameStorageContainerIDDefault{
		_statusCode: code,
	}
}

/*GetCloudProjectServiceNameStorageContainerIDDefault handles this case with default header values.

Unexpected error
*/
type GetCloudProjectServiceNameStorageContainerIDDefault struct {
	_statusCode int

	Payload *models.GetCloudProjectServiceNameStorageContainerIDDefaultBody
}

// Code gets the status code for the get cloud project service name storage container ID default response
func (o *GetCloudProjectServiceNameStorageContainerIDDefault) Code() int {
	return o._statusCode
}

func (o *GetCloudProjectServiceNameStorageContainerIDDefault) Error() string {
	return fmt.Sprintf("[GET /cloud/project/{serviceName}/storage/{containerId}][%d] GetCloudProjectServiceNameStorageContainerID default  %+v", o._statusCode, o.Payload)
}

func (o *GetCloudProjectServiceNameStorageContainerIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCloudProjectServiceNameStorageContainerIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
