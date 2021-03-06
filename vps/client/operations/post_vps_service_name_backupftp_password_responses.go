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

// PostVpsServiceNameBackupftpPasswordReader is a Reader for the PostVpsServiceNameBackupftpPassword structure.
type PostVpsServiceNameBackupftpPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVpsServiceNameBackupftpPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostVpsServiceNameBackupftpPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostVpsServiceNameBackupftpPasswordDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostVpsServiceNameBackupftpPasswordOK creates a PostVpsServiceNameBackupftpPasswordOK with default headers values
func NewPostVpsServiceNameBackupftpPasswordOK() *PostVpsServiceNameBackupftpPasswordOK {
	return &PostVpsServiceNameBackupftpPasswordOK{}
}

/*PostVpsServiceNameBackupftpPasswordOK handles this case with default header values.

description of 'dedicated.Server.Task' response
*/
type PostVpsServiceNameBackupftpPasswordOK struct {
	Payload *models.DedicatedServerTask
}

func (o *PostVpsServiceNameBackupftpPasswordOK) Error() string {
	return fmt.Sprintf("[POST /vps/{serviceName}/backupftp/password][%d] postVpsServiceNameBackupftpPasswordOK  %+v", 200, o.Payload)
}

func (o *PostVpsServiceNameBackupftpPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DedicatedServerTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVpsServiceNameBackupftpPasswordDefault creates a PostVpsServiceNameBackupftpPasswordDefault with default headers values
func NewPostVpsServiceNameBackupftpPasswordDefault(code int) *PostVpsServiceNameBackupftpPasswordDefault {
	return &PostVpsServiceNameBackupftpPasswordDefault{
		_statusCode: code,
	}
}

/*PostVpsServiceNameBackupftpPasswordDefault handles this case with default header values.

Unexpected error
*/
type PostVpsServiceNameBackupftpPasswordDefault struct {
	_statusCode int

	Payload *models.PostVpsServiceNameBackupftpPasswordDefaultBody
}

// Code gets the status code for the post vps service name backupftp password default response
func (o *PostVpsServiceNameBackupftpPasswordDefault) Code() int {
	return o._statusCode
}

func (o *PostVpsServiceNameBackupftpPasswordDefault) Error() string {
	return fmt.Sprintf("[POST /vps/{serviceName}/backupftp/password][%d] PostVpsServiceNameBackupftpPassword default  %+v", o._statusCode, o.Payload)
}

func (o *PostVpsServiceNameBackupftpPasswordDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostVpsServiceNameBackupftpPasswordDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
