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

// GetCloudProjectServiceNameVolumeSnapshotSnapshotIDReader is a Reader for the GetCloudProjectServiceNameVolumeSnapshotSnapshotID structure.
type GetCloudProjectServiceNameVolumeSnapshotSnapshotIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudProjectServiceNameVolumeSnapshotSnapshotIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCloudProjectServiceNameVolumeSnapshotSnapshotIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCloudProjectServiceNameVolumeSnapshotSnapshotIDOK creates a GetCloudProjectServiceNameVolumeSnapshotSnapshotIDOK with default headers values
func NewGetCloudProjectServiceNameVolumeSnapshotSnapshotIDOK() *GetCloudProjectServiceNameVolumeSnapshotSnapshotIDOK {
	return &GetCloudProjectServiceNameVolumeSnapshotSnapshotIDOK{}
}

/*GetCloudProjectServiceNameVolumeSnapshotSnapshotIDOK handles this case with default header values.

description of 'cloud.Volume.Snapshot' response
*/
type GetCloudProjectServiceNameVolumeSnapshotSnapshotIDOK struct {
	Payload *models.CloudVolumeSnapshot
}

func (o *GetCloudProjectServiceNameVolumeSnapshotSnapshotIDOK) Error() string {
	return fmt.Sprintf("[GET /cloud/project/{serviceName}/volume/snapshot/{snapshotId}][%d] getCloudProjectServiceNameVolumeSnapshotSnapshotIdOK  %+v", 200, o.Payload)
}

func (o *GetCloudProjectServiceNameVolumeSnapshotSnapshotIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudVolumeSnapshot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault creates a GetCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault with default headers values
func NewGetCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault(code int) *GetCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault {
	return &GetCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault{
		_statusCode: code,
	}
}

/*GetCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault handles this case with default header values.

Unexpected error
*/
type GetCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault struct {
	_statusCode int

	Payload *models.GetCloudProjectServiceNameVolumeSnapshotSnapshotIDDefaultBody
}

// Code gets the status code for the get cloud project service name volume snapshot snapshot ID default response
func (o *GetCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault) Code() int {
	return o._statusCode
}

func (o *GetCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault) Error() string {
	return fmt.Sprintf("[GET /cloud/project/{serviceName}/volume/snapshot/{snapshotId}][%d] GetCloudProjectServiceNameVolumeSnapshotSnapshotID default  %+v", o._statusCode, o.Payload)
}

func (o *GetCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCloudProjectServiceNameVolumeSnapshotSnapshotIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}