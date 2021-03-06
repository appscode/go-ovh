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

// GetCloudProjectServiceNameSshkeyReader is a Reader for the GetCloudProjectServiceNameSshkey structure.
type GetCloudProjectServiceNameSshkeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudProjectServiceNameSshkeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCloudProjectServiceNameSshkeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetCloudProjectServiceNameSshkeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCloudProjectServiceNameSshkeyOK creates a GetCloudProjectServiceNameSshkeyOK with default headers values
func NewGetCloudProjectServiceNameSshkeyOK() *GetCloudProjectServiceNameSshkeyOK {
	return &GetCloudProjectServiceNameSshkeyOK{}
}

/*GetCloudProjectServiceNameSshkeyOK handles this case with default header values.

return value
*/
type GetCloudProjectServiceNameSshkeyOK struct {
	Payload models.GetCloudProjectServiceNameSshkeyOKBody
}

func (o *GetCloudProjectServiceNameSshkeyOK) Error() string {
	return fmt.Sprintf("[GET /cloud/project/{serviceName}/sshkey][%d] getCloudProjectServiceNameSshkeyOK  %+v", 200, o.Payload)
}

func (o *GetCloudProjectServiceNameSshkeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudProjectServiceNameSshkeyDefault creates a GetCloudProjectServiceNameSshkeyDefault with default headers values
func NewGetCloudProjectServiceNameSshkeyDefault(code int) *GetCloudProjectServiceNameSshkeyDefault {
	return &GetCloudProjectServiceNameSshkeyDefault{
		_statusCode: code,
	}
}

/*GetCloudProjectServiceNameSshkeyDefault handles this case with default header values.

Unexpected error
*/
type GetCloudProjectServiceNameSshkeyDefault struct {
	_statusCode int

	Payload *models.GetCloudProjectServiceNameSshkeyDefaultBody
}

// Code gets the status code for the get cloud project service name sshkey default response
func (o *GetCloudProjectServiceNameSshkeyDefault) Code() int {
	return o._statusCode
}

func (o *GetCloudProjectServiceNameSshkeyDefault) Error() string {
	return fmt.Sprintf("[GET /cloud/project/{serviceName}/sshkey][%d] GetCloudProjectServiceNameSshkey default  %+v", o._statusCode, o.Payload)
}

func (o *GetCloudProjectServiceNameSshkeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCloudProjectServiceNameSshkeyDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
