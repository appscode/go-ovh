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

// PostCloudProjectServiceNameVolumeVolumeIDSnapshotReader is a Reader for the PostCloudProjectServiceNameVolumeVolumeIDSnapshot structure.
type PostCloudProjectServiceNameVolumeVolumeIDSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCloudProjectServiceNameVolumeVolumeIDSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostCloudProjectServiceNameVolumeVolumeIDSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostCloudProjectServiceNameVolumeVolumeIDSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCloudProjectServiceNameVolumeVolumeIDSnapshotOK creates a PostCloudProjectServiceNameVolumeVolumeIDSnapshotOK with default headers values
func NewPostCloudProjectServiceNameVolumeVolumeIDSnapshotOK() *PostCloudProjectServiceNameVolumeVolumeIDSnapshotOK {
	return &PostCloudProjectServiceNameVolumeVolumeIDSnapshotOK{}
}

/*PostCloudProjectServiceNameVolumeVolumeIDSnapshotOK handles this case with default header values.

description of 'cloud.Volume.Snapshot' response
*/
type PostCloudProjectServiceNameVolumeVolumeIDSnapshotOK struct {
	Payload *models.CloudVolumeSnapshot
}

func (o *PostCloudProjectServiceNameVolumeVolumeIDSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /cloud/project/{serviceName}/volume/{volumeId}/snapshot][%d] postCloudProjectServiceNameVolumeVolumeIdSnapshotOK  %+v", 200, o.Payload)
}

func (o *PostCloudProjectServiceNameVolumeVolumeIDSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudVolumeSnapshot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCloudProjectServiceNameVolumeVolumeIDSnapshotDefault creates a PostCloudProjectServiceNameVolumeVolumeIDSnapshotDefault with default headers values
func NewPostCloudProjectServiceNameVolumeVolumeIDSnapshotDefault(code int) *PostCloudProjectServiceNameVolumeVolumeIDSnapshotDefault {
	return &PostCloudProjectServiceNameVolumeVolumeIDSnapshotDefault{
		_statusCode: code,
	}
}

/*PostCloudProjectServiceNameVolumeVolumeIDSnapshotDefault handles this case with default header values.

Unexpected error
*/
type PostCloudProjectServiceNameVolumeVolumeIDSnapshotDefault struct {
	_statusCode int

	Payload *models.PostCloudProjectServiceNameVolumeVolumeIDSnapshotDefaultBody
}

// Code gets the status code for the post cloud project service name volume volume ID snapshot default response
func (o *PostCloudProjectServiceNameVolumeVolumeIDSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *PostCloudProjectServiceNameVolumeVolumeIDSnapshotDefault) Error() string {
	return fmt.Sprintf("[POST /cloud/project/{serviceName}/volume/{volumeId}/snapshot][%d] PostCloudProjectServiceNameVolumeVolumeIDSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *PostCloudProjectServiceNameVolumeVolumeIDSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostCloudProjectServiceNameVolumeVolumeIDSnapshotDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
