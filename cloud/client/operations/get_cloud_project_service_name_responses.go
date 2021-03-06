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

// GetCloudProjectServiceNameReader is a Reader for the GetCloudProjectServiceName structure.
type GetCloudProjectServiceNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudProjectServiceNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCloudProjectServiceNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetCloudProjectServiceNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCloudProjectServiceNameOK creates a GetCloudProjectServiceNameOK with default headers values
func NewGetCloudProjectServiceNameOK() *GetCloudProjectServiceNameOK {
	return &GetCloudProjectServiceNameOK{}
}

/*GetCloudProjectServiceNameOK handles this case with default header values.

description of 'cloud.Project' response
*/
type GetCloudProjectServiceNameOK struct {
	Payload *models.CloudProject
}

func (o *GetCloudProjectServiceNameOK) Error() string {
	return fmt.Sprintf("[GET /cloud/project/{serviceName}][%d] getCloudProjectServiceNameOK  %+v", 200, o.Payload)
}

func (o *GetCloudProjectServiceNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudProject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudProjectServiceNameDefault creates a GetCloudProjectServiceNameDefault with default headers values
func NewGetCloudProjectServiceNameDefault(code int) *GetCloudProjectServiceNameDefault {
	return &GetCloudProjectServiceNameDefault{
		_statusCode: code,
	}
}

/*GetCloudProjectServiceNameDefault handles this case with default header values.

Unexpected error
*/
type GetCloudProjectServiceNameDefault struct {
	_statusCode int

	Payload *models.GetCloudProjectServiceNameDefaultBody
}

// Code gets the status code for the get cloud project service name default response
func (o *GetCloudProjectServiceNameDefault) Code() int {
	return o._statusCode
}

func (o *GetCloudProjectServiceNameDefault) Error() string {
	return fmt.Sprintf("[GET /cloud/project/{serviceName}][%d] GetCloudProjectServiceName default  %+v", o._statusCode, o.Payload)
}

func (o *GetCloudProjectServiceNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCloudProjectServiceNameDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
