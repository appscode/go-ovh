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

// PostCloudProjectServiceNameUnleashReader is a Reader for the PostCloudProjectServiceNameUnleash structure.
type PostCloudProjectServiceNameUnleashReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCloudProjectServiceNameUnleashReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostCloudProjectServiceNameUnleashOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostCloudProjectServiceNameUnleashDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCloudProjectServiceNameUnleashOK creates a PostCloudProjectServiceNameUnleashOK with default headers values
func NewPostCloudProjectServiceNameUnleashOK() *PostCloudProjectServiceNameUnleashOK {
	return &PostCloudProjectServiceNameUnleashOK{}
}

/*PostCloudProjectServiceNameUnleashOK handles this case with default header values.

return 'void'
*/
type PostCloudProjectServiceNameUnleashOK struct {
}

func (o *PostCloudProjectServiceNameUnleashOK) Error() string {
	return fmt.Sprintf("[POST /cloud/project/{serviceName}/unleash][%d] postCloudProjectServiceNameUnleashOK ", 200)
}

func (o *PostCloudProjectServiceNameUnleashOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCloudProjectServiceNameUnleashDefault creates a PostCloudProjectServiceNameUnleashDefault with default headers values
func NewPostCloudProjectServiceNameUnleashDefault(code int) *PostCloudProjectServiceNameUnleashDefault {
	return &PostCloudProjectServiceNameUnleashDefault{
		_statusCode: code,
	}
}

/*PostCloudProjectServiceNameUnleashDefault handles this case with default header values.

Unexpected error
*/
type PostCloudProjectServiceNameUnleashDefault struct {
	_statusCode int

	Payload *models.PostCloudProjectServiceNameUnleashDefaultBody
}

// Code gets the status code for the post cloud project service name unleash default response
func (o *PostCloudProjectServiceNameUnleashDefault) Code() int {
	return o._statusCode
}

func (o *PostCloudProjectServiceNameUnleashDefault) Error() string {
	return fmt.Sprintf("[POST /cloud/project/{serviceName}/unleash][%d] PostCloudProjectServiceNameUnleash default  %+v", o._statusCode, o.Payload)
}

func (o *PostCloudProjectServiceNameUnleashDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostCloudProjectServiceNameUnleashDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
