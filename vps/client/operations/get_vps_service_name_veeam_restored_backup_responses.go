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

// GetVpsServiceNameVeeamRestoredBackupReader is a Reader for the GetVpsServiceNameVeeamRestoredBackup structure.
type GetVpsServiceNameVeeamRestoredBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVpsServiceNameVeeamRestoredBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVpsServiceNameVeeamRestoredBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetVpsServiceNameVeeamRestoredBackupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVpsServiceNameVeeamRestoredBackupOK creates a GetVpsServiceNameVeeamRestoredBackupOK with default headers values
func NewGetVpsServiceNameVeeamRestoredBackupOK() *GetVpsServiceNameVeeamRestoredBackupOK {
	return &GetVpsServiceNameVeeamRestoredBackupOK{}
}

/*GetVpsServiceNameVeeamRestoredBackupOK handles this case with default header values.

description of 'vps.Veeam.RestoredBackup' response
*/
type GetVpsServiceNameVeeamRestoredBackupOK struct {
	Payload *models.VpsVeeamRestoredBackup
}

func (o *GetVpsServiceNameVeeamRestoredBackupOK) Error() string {
	return fmt.Sprintf("[GET /vps/{serviceName}/veeam/restoredBackup][%d] getVpsServiceNameVeeamRestoredBackupOK  %+v", 200, o.Payload)
}

func (o *GetVpsServiceNameVeeamRestoredBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VpsVeeamRestoredBackup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVpsServiceNameVeeamRestoredBackupDefault creates a GetVpsServiceNameVeeamRestoredBackupDefault with default headers values
func NewGetVpsServiceNameVeeamRestoredBackupDefault(code int) *GetVpsServiceNameVeeamRestoredBackupDefault {
	return &GetVpsServiceNameVeeamRestoredBackupDefault{
		_statusCode: code,
	}
}

/*GetVpsServiceNameVeeamRestoredBackupDefault handles this case with default header values.

Unexpected error
*/
type GetVpsServiceNameVeeamRestoredBackupDefault struct {
	_statusCode int

	Payload *models.GetVpsServiceNameVeeamRestoredBackupDefaultBody
}

// Code gets the status code for the get vps service name veeam restored backup default response
func (o *GetVpsServiceNameVeeamRestoredBackupDefault) Code() int {
	return o._statusCode
}

func (o *GetVpsServiceNameVeeamRestoredBackupDefault) Error() string {
	return fmt.Sprintf("[GET /vps/{serviceName}/veeam/restoredBackup][%d] GetVpsServiceNameVeeamRestoredBackup default  %+v", o._statusCode, o.Payload)
}

func (o *GetVpsServiceNameVeeamRestoredBackupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetVpsServiceNameVeeamRestoredBackupDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}