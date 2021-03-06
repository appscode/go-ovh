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

// GetVpsServiceNameActiveOptionsReader is a Reader for the GetVpsServiceNameActiveOptions structure.
type GetVpsServiceNameActiveOptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVpsServiceNameActiveOptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVpsServiceNameActiveOptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetVpsServiceNameActiveOptionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVpsServiceNameActiveOptionsOK creates a GetVpsServiceNameActiveOptionsOK with default headers values
func NewGetVpsServiceNameActiveOptionsOK() *GetVpsServiceNameActiveOptionsOK {
	return &GetVpsServiceNameActiveOptionsOK{}
}

/*GetVpsServiceNameActiveOptionsOK handles this case with default header values.

return value
*/
type GetVpsServiceNameActiveOptionsOK struct {
	Payload []string
}

func (o *GetVpsServiceNameActiveOptionsOK) Error() string {
	return fmt.Sprintf("[GET /vps/{serviceName}/activeOptions][%d] getVpsServiceNameActiveOptionsOK  %+v", 200, o.Payload)
}

func (o *GetVpsServiceNameActiveOptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVpsServiceNameActiveOptionsDefault creates a GetVpsServiceNameActiveOptionsDefault with default headers values
func NewGetVpsServiceNameActiveOptionsDefault(code int) *GetVpsServiceNameActiveOptionsDefault {
	return &GetVpsServiceNameActiveOptionsDefault{
		_statusCode: code,
	}
}

/*GetVpsServiceNameActiveOptionsDefault handles this case with default header values.

Unexpected error
*/
type GetVpsServiceNameActiveOptionsDefault struct {
	_statusCode int

	Payload *models.GetVpsServiceNameActiveOptionsDefaultBody
}

// Code gets the status code for the get vps service name active options default response
func (o *GetVpsServiceNameActiveOptionsDefault) Code() int {
	return o._statusCode
}

func (o *GetVpsServiceNameActiveOptionsDefault) Error() string {
	return fmt.Sprintf("[GET /vps/{serviceName}/activeOptions][%d] GetVpsServiceNameActiveOptions default  %+v", o._statusCode, o.Payload)
}

func (o *GetVpsServiceNameActiveOptionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetVpsServiceNameActiveOptionsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
