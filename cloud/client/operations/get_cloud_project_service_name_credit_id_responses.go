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

// GetCloudProjectServiceNameCreditIDReader is a Reader for the GetCloudProjectServiceNameCreditID structure.
type GetCloudProjectServiceNameCreditIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudProjectServiceNameCreditIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCloudProjectServiceNameCreditIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetCloudProjectServiceNameCreditIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCloudProjectServiceNameCreditIDOK creates a GetCloudProjectServiceNameCreditIDOK with default headers values
func NewGetCloudProjectServiceNameCreditIDOK() *GetCloudProjectServiceNameCreditIDOK {
	return &GetCloudProjectServiceNameCreditIDOK{}
}

/*GetCloudProjectServiceNameCreditIDOK handles this case with default header values.

description of 'cloud.Credit' response
*/
type GetCloudProjectServiceNameCreditIDOK struct {
	Payload *models.CloudCredit
}

func (o *GetCloudProjectServiceNameCreditIDOK) Error() string {
	return fmt.Sprintf("[GET /cloud/project/{serviceName}/credit/{id}][%d] getCloudProjectServiceNameCreditIdOK  %+v", 200, o.Payload)
}

func (o *GetCloudProjectServiceNameCreditIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudCredit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudProjectServiceNameCreditIDDefault creates a GetCloudProjectServiceNameCreditIDDefault with default headers values
func NewGetCloudProjectServiceNameCreditIDDefault(code int) *GetCloudProjectServiceNameCreditIDDefault {
	return &GetCloudProjectServiceNameCreditIDDefault{
		_statusCode: code,
	}
}

/*GetCloudProjectServiceNameCreditIDDefault handles this case with default header values.

Unexpected error
*/
type GetCloudProjectServiceNameCreditIDDefault struct {
	_statusCode int

	Payload *models.GetCloudProjectServiceNameCreditIDDefaultBody
}

// Code gets the status code for the get cloud project service name credit ID default response
func (o *GetCloudProjectServiceNameCreditIDDefault) Code() int {
	return o._statusCode
}

func (o *GetCloudProjectServiceNameCreditIDDefault) Error() string {
	return fmt.Sprintf("[GET /cloud/project/{serviceName}/credit/{id}][%d] GetCloudProjectServiceNameCreditID default  %+v", o._statusCode, o.Payload)
}

func (o *GetCloudProjectServiceNameCreditIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCloudProjectServiceNameCreditIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
