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

// PostCloudProjectServiceNameStorageContainerIDPublicURLReader is a Reader for the PostCloudProjectServiceNameStorageContainerIDPublicURL structure.
type PostCloudProjectServiceNameStorageContainerIDPublicURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCloudProjectServiceNameStorageContainerIDPublicURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostCloudProjectServiceNameStorageContainerIDPublicURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostCloudProjectServiceNameStorageContainerIDPublicURLDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCloudProjectServiceNameStorageContainerIDPublicURLOK creates a PostCloudProjectServiceNameStorageContainerIDPublicURLOK with default headers values
func NewPostCloudProjectServiceNameStorageContainerIDPublicURLOK() *PostCloudProjectServiceNameStorageContainerIDPublicURLOK {
	return &PostCloudProjectServiceNameStorageContainerIDPublicURLOK{}
}

/*PostCloudProjectServiceNameStorageContainerIDPublicURLOK handles this case with default header values.

description of 'cloud.Storage.ContainerObjectTempURL' response
*/
type PostCloudProjectServiceNameStorageContainerIDPublicURLOK struct {
	Payload *models.CloudStorageContainerObjectTempURL
}

func (o *PostCloudProjectServiceNameStorageContainerIDPublicURLOK) Error() string {
	return fmt.Sprintf("[POST /cloud/project/{serviceName}/storage/{containerId}/publicUrl][%d] postCloudProjectServiceNameStorageContainerIdPublicUrlOK  %+v", 200, o.Payload)
}

func (o *PostCloudProjectServiceNameStorageContainerIDPublicURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudStorageContainerObjectTempURL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCloudProjectServiceNameStorageContainerIDPublicURLDefault creates a PostCloudProjectServiceNameStorageContainerIDPublicURLDefault with default headers values
func NewPostCloudProjectServiceNameStorageContainerIDPublicURLDefault(code int) *PostCloudProjectServiceNameStorageContainerIDPublicURLDefault {
	return &PostCloudProjectServiceNameStorageContainerIDPublicURLDefault{
		_statusCode: code,
	}
}

/*PostCloudProjectServiceNameStorageContainerIDPublicURLDefault handles this case with default header values.

Unexpected error
*/
type PostCloudProjectServiceNameStorageContainerIDPublicURLDefault struct {
	_statusCode int

	Payload *models.PostCloudProjectServiceNameStorageContainerIDPublicURLDefaultBody
}

// Code gets the status code for the post cloud project service name storage container ID public URL default response
func (o *PostCloudProjectServiceNameStorageContainerIDPublicURLDefault) Code() int {
	return o._statusCode
}

func (o *PostCloudProjectServiceNameStorageContainerIDPublicURLDefault) Error() string {
	return fmt.Sprintf("[POST /cloud/project/{serviceName}/storage/{containerId}/publicUrl][%d] PostCloudProjectServiceNameStorageContainerIDPublicURL default  %+v", o._statusCode, o.Payload)
}

func (o *PostCloudProjectServiceNameStorageContainerIDPublicURLDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostCloudProjectServiceNameStorageContainerIDPublicURLDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
