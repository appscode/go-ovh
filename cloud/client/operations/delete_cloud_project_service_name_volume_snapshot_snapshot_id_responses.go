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

// DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDReader is a Reader for the DeleteCloudProjectServiceNameVolumeSnapshotSnapshotID structure.
type DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDOK creates a DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDOK with default headers values
func NewDeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDOK() *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDOK {
	return &DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDOK{}
}

/*DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDOK handles this case with default header values.

return 'void'
*/
type DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDOK struct {
}

func (o *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDOK) Error() string {
	return fmt.Sprintf("[DELETE /cloud/project/{serviceName}/volume/snapshot/{snapshotId}][%d] deleteCloudProjectServiceNameVolumeSnapshotSnapshotIdOK ", 200)
}

func (o *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault creates a DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault with default headers values
func NewDeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault(code int) *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault {
	return &DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault{
		_statusCode: code,
	}
}

/*DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault handles this case with default header values.

Unexpected error
*/
type DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault struct {
	_statusCode int

	Payload *models.DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDDefaultBody
}

// Code gets the status code for the delete cloud project service name volume snapshot snapshot ID default response
func (o *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /cloud/project/{serviceName}/volume/snapshot/{snapshotId}][%d] DeleteCloudProjectServiceNameVolumeSnapshotSnapshotID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}