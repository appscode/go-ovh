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

// GetVpsServiceNameAutomatedBackupAttachedBackupReader is a Reader for the GetVpsServiceNameAutomatedBackupAttachedBackup structure.
type GetVpsServiceNameAutomatedBackupAttachedBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVpsServiceNameAutomatedBackupAttachedBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVpsServiceNameAutomatedBackupAttachedBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetVpsServiceNameAutomatedBackupAttachedBackupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVpsServiceNameAutomatedBackupAttachedBackupOK creates a GetVpsServiceNameAutomatedBackupAttachedBackupOK with default headers values
func NewGetVpsServiceNameAutomatedBackupAttachedBackupOK() *GetVpsServiceNameAutomatedBackupAttachedBackupOK {
	return &GetVpsServiceNameAutomatedBackupAttachedBackupOK{}
}

/*GetVpsServiceNameAutomatedBackupAttachedBackupOK handles this case with default header values.

return value
*/
type GetVpsServiceNameAutomatedBackupAttachedBackupOK struct {
	Payload models.GetVpsServiceNameAutomatedBackupAttachedBackupOKBody
}

func (o *GetVpsServiceNameAutomatedBackupAttachedBackupOK) Error() string {
	return fmt.Sprintf("[GET /vps/{serviceName}/automatedBackup/attachedBackup][%d] getVpsServiceNameAutomatedBackupAttachedBackupOK  %+v", 200, o.Payload)
}

func (o *GetVpsServiceNameAutomatedBackupAttachedBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVpsServiceNameAutomatedBackupAttachedBackupDefault creates a GetVpsServiceNameAutomatedBackupAttachedBackupDefault with default headers values
func NewGetVpsServiceNameAutomatedBackupAttachedBackupDefault(code int) *GetVpsServiceNameAutomatedBackupAttachedBackupDefault {
	return &GetVpsServiceNameAutomatedBackupAttachedBackupDefault{
		_statusCode: code,
	}
}

/*GetVpsServiceNameAutomatedBackupAttachedBackupDefault handles this case with default header values.

Unexpected error
*/
type GetVpsServiceNameAutomatedBackupAttachedBackupDefault struct {
	_statusCode int

	Payload *models.GetVpsServiceNameAutomatedBackupAttachedBackupDefaultBody
}

// Code gets the status code for the get vps service name automated backup attached backup default response
func (o *GetVpsServiceNameAutomatedBackupAttachedBackupDefault) Code() int {
	return o._statusCode
}

func (o *GetVpsServiceNameAutomatedBackupAttachedBackupDefault) Error() string {
	return fmt.Sprintf("[GET /vps/{serviceName}/automatedBackup/attachedBackup][%d] GetVpsServiceNameAutomatedBackupAttachedBackup default  %+v", o._statusCode, o.Payload)
}

func (o *GetVpsServiceNameAutomatedBackupAttachedBackupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetVpsServiceNameAutomatedBackupAttachedBackupDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}