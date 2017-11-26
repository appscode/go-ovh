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

// PostCloudProjectServiceNameVolumeVolumeIDDetachReader is a Reader for the PostCloudProjectServiceNameVolumeVolumeIDDetach structure.
type PostCloudProjectServiceNameVolumeVolumeIDDetachReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCloudProjectServiceNameVolumeVolumeIDDetachReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostCloudProjectServiceNameVolumeVolumeIDDetachOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostCloudProjectServiceNameVolumeVolumeIDDetachDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCloudProjectServiceNameVolumeVolumeIDDetachOK creates a PostCloudProjectServiceNameVolumeVolumeIDDetachOK with default headers values
func NewPostCloudProjectServiceNameVolumeVolumeIDDetachOK() *PostCloudProjectServiceNameVolumeVolumeIDDetachOK {
	return &PostCloudProjectServiceNameVolumeVolumeIDDetachOK{}
}

/*PostCloudProjectServiceNameVolumeVolumeIDDetachOK handles this case with default header values.

description of 'cloud.Volume.Volume' response
*/
type PostCloudProjectServiceNameVolumeVolumeIDDetachOK struct {
	Payload *models.CloudVolumeVolume
}

func (o *PostCloudProjectServiceNameVolumeVolumeIDDetachOK) Error() string {
	return fmt.Sprintf("[POST /cloud/project/{serviceName}/volume/{volumeId}/detach][%d] postCloudProjectServiceNameVolumeVolumeIdDetachOK  %+v", 200, o.Payload)
}

func (o *PostCloudProjectServiceNameVolumeVolumeIDDetachOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudVolumeVolume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCloudProjectServiceNameVolumeVolumeIDDetachDefault creates a PostCloudProjectServiceNameVolumeVolumeIDDetachDefault with default headers values
func NewPostCloudProjectServiceNameVolumeVolumeIDDetachDefault(code int) *PostCloudProjectServiceNameVolumeVolumeIDDetachDefault {
	return &PostCloudProjectServiceNameVolumeVolumeIDDetachDefault{
		_statusCode: code,
	}
}

/*PostCloudProjectServiceNameVolumeVolumeIDDetachDefault handles this case with default header values.

Unexpected error
*/
type PostCloudProjectServiceNameVolumeVolumeIDDetachDefault struct {
	_statusCode int

	Payload *models.PostCloudProjectServiceNameVolumeVolumeIDDetachDefaultBody
}

// Code gets the status code for the post cloud project service name volume volume ID detach default response
func (o *PostCloudProjectServiceNameVolumeVolumeIDDetachDefault) Code() int {
	return o._statusCode
}

func (o *PostCloudProjectServiceNameVolumeVolumeIDDetachDefault) Error() string {
	return fmt.Sprintf("[POST /cloud/project/{serviceName}/volume/{volumeId}/detach][%d] PostCloudProjectServiceNameVolumeVolumeIDDetach default  %+v", o._statusCode, o.Payload)
}

func (o *PostCloudProjectServiceNameVolumeVolumeIDDetachDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostCloudProjectServiceNameVolumeVolumeIDDetachDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}