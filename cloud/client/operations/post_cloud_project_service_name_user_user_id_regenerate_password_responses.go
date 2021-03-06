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

// PostCloudProjectServiceNameUserUserIDRegeneratePasswordReader is a Reader for the PostCloudProjectServiceNameUserUserIDRegeneratePassword structure.
type PostCloudProjectServiceNameUserUserIDRegeneratePasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCloudProjectServiceNameUserUserIDRegeneratePasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostCloudProjectServiceNameUserUserIDRegeneratePasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostCloudProjectServiceNameUserUserIDRegeneratePasswordDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCloudProjectServiceNameUserUserIDRegeneratePasswordOK creates a PostCloudProjectServiceNameUserUserIDRegeneratePasswordOK with default headers values
func NewPostCloudProjectServiceNameUserUserIDRegeneratePasswordOK() *PostCloudProjectServiceNameUserUserIDRegeneratePasswordOK {
	return &PostCloudProjectServiceNameUserUserIDRegeneratePasswordOK{}
}

/*PostCloudProjectServiceNameUserUserIDRegeneratePasswordOK handles this case with default header values.

description of 'cloud.User.UserDetail' response
*/
type PostCloudProjectServiceNameUserUserIDRegeneratePasswordOK struct {
	Payload *models.CloudUserUserDetail
}

func (o *PostCloudProjectServiceNameUserUserIDRegeneratePasswordOK) Error() string {
	return fmt.Sprintf("[POST /cloud/project/{serviceName}/user/{userId}/regeneratePassword][%d] postCloudProjectServiceNameUserUserIdRegeneratePasswordOK  %+v", 200, o.Payload)
}

func (o *PostCloudProjectServiceNameUserUserIDRegeneratePasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudUserUserDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCloudProjectServiceNameUserUserIDRegeneratePasswordDefault creates a PostCloudProjectServiceNameUserUserIDRegeneratePasswordDefault with default headers values
func NewPostCloudProjectServiceNameUserUserIDRegeneratePasswordDefault(code int) *PostCloudProjectServiceNameUserUserIDRegeneratePasswordDefault {
	return &PostCloudProjectServiceNameUserUserIDRegeneratePasswordDefault{
		_statusCode: code,
	}
}

/*PostCloudProjectServiceNameUserUserIDRegeneratePasswordDefault handles this case with default header values.

Unexpected error
*/
type PostCloudProjectServiceNameUserUserIDRegeneratePasswordDefault struct {
	_statusCode int

	Payload *models.PostCloudProjectServiceNameUserUserIDRegeneratePasswordDefaultBody
}

// Code gets the status code for the post cloud project service name user user ID regenerate password default response
func (o *PostCloudProjectServiceNameUserUserIDRegeneratePasswordDefault) Code() int {
	return o._statusCode
}

func (o *PostCloudProjectServiceNameUserUserIDRegeneratePasswordDefault) Error() string {
	return fmt.Sprintf("[POST /cloud/project/{serviceName}/user/{userId}/regeneratePassword][%d] PostCloudProjectServiceNameUserUserIDRegeneratePassword default  %+v", o._statusCode, o.Payload)
}

func (o *PostCloudProjectServiceNameUserUserIDRegeneratePasswordDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostCloudProjectServiceNameUserUserIDRegeneratePasswordDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
