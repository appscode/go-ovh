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

// DeleteVpsServiceNameBackupftpAccessIPBlockReader is a Reader for the DeleteVpsServiceNameBackupftpAccessIPBlock structure.
type DeleteVpsServiceNameBackupftpAccessIPBlockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVpsServiceNameBackupftpAccessIPBlockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteVpsServiceNameBackupftpAccessIPBlockOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteVpsServiceNameBackupftpAccessIPBlockDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteVpsServiceNameBackupftpAccessIPBlockOK creates a DeleteVpsServiceNameBackupftpAccessIPBlockOK with default headers values
func NewDeleteVpsServiceNameBackupftpAccessIPBlockOK() *DeleteVpsServiceNameBackupftpAccessIPBlockOK {
	return &DeleteVpsServiceNameBackupftpAccessIPBlockOK{}
}

/*DeleteVpsServiceNameBackupftpAccessIPBlockOK handles this case with default header values.

description of 'dedicated.Server.Task' response
*/
type DeleteVpsServiceNameBackupftpAccessIPBlockOK struct {
	Payload *models.DedicatedServerTask
}

func (o *DeleteVpsServiceNameBackupftpAccessIPBlockOK) Error() string {
	return fmt.Sprintf("[DELETE /vps/{serviceName}/backupftp/access/{ipBlock}][%d] deleteVpsServiceNameBackupftpAccessIpBlockOK  %+v", 200, o.Payload)
}

func (o *DeleteVpsServiceNameBackupftpAccessIPBlockOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DedicatedServerTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVpsServiceNameBackupftpAccessIPBlockDefault creates a DeleteVpsServiceNameBackupftpAccessIPBlockDefault with default headers values
func NewDeleteVpsServiceNameBackupftpAccessIPBlockDefault(code int) *DeleteVpsServiceNameBackupftpAccessIPBlockDefault {
	return &DeleteVpsServiceNameBackupftpAccessIPBlockDefault{
		_statusCode: code,
	}
}

/*DeleteVpsServiceNameBackupftpAccessIPBlockDefault handles this case with default header values.

Unexpected error
*/
type DeleteVpsServiceNameBackupftpAccessIPBlockDefault struct {
	_statusCode int

	Payload *models.DeleteVpsServiceNameBackupftpAccessIPBlockDefaultBody
}

// Code gets the status code for the delete vps service name backupftp access IP block default response
func (o *DeleteVpsServiceNameBackupftpAccessIPBlockDefault) Code() int {
	return o._statusCode
}

func (o *DeleteVpsServiceNameBackupftpAccessIPBlockDefault) Error() string {
	return fmt.Sprintf("[DELETE /vps/{serviceName}/backupftp/access/{ipBlock}][%d] DeleteVpsServiceNameBackupftpAccessIPBlock default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteVpsServiceNameBackupftpAccessIPBlockDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteVpsServiceNameBackupftpAccessIPBlockDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
