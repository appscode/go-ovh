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

// PostCloudProjectServiceNameRetainReader is a Reader for the PostCloudProjectServiceNameRetain structure.
type PostCloudProjectServiceNameRetainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCloudProjectServiceNameRetainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostCloudProjectServiceNameRetainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostCloudProjectServiceNameRetainDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCloudProjectServiceNameRetainOK creates a PostCloudProjectServiceNameRetainOK with default headers values
func NewPostCloudProjectServiceNameRetainOK() *PostCloudProjectServiceNameRetainOK {
	return &PostCloudProjectServiceNameRetainOK{}
}

/*PostCloudProjectServiceNameRetainOK handles this case with default header values.

return 'void'
*/
type PostCloudProjectServiceNameRetainOK struct {
}

func (o *PostCloudProjectServiceNameRetainOK) Error() string {
	return fmt.Sprintf("[POST /cloud/project/{serviceName}/retain][%d] postCloudProjectServiceNameRetainOK ", 200)
}

func (o *PostCloudProjectServiceNameRetainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCloudProjectServiceNameRetainDefault creates a PostCloudProjectServiceNameRetainDefault with default headers values
func NewPostCloudProjectServiceNameRetainDefault(code int) *PostCloudProjectServiceNameRetainDefault {
	return &PostCloudProjectServiceNameRetainDefault{
		_statusCode: code,
	}
}

/*PostCloudProjectServiceNameRetainDefault handles this case with default header values.

Unexpected error
*/
type PostCloudProjectServiceNameRetainDefault struct {
	_statusCode int

	Payload *models.PostCloudProjectServiceNameRetainDefaultBody
}

// Code gets the status code for the post cloud project service name retain default response
func (o *PostCloudProjectServiceNameRetainDefault) Code() int {
	return o._statusCode
}

func (o *PostCloudProjectServiceNameRetainDefault) Error() string {
	return fmt.Sprintf("[POST /cloud/project/{serviceName}/retain][%d] PostCloudProjectServiceNameRetain default  %+v", o._statusCode, o.Payload)
}

func (o *PostCloudProjectServiceNameRetainDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostCloudProjectServiceNameRetainDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
