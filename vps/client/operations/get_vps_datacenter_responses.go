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

// GetVpsDatacenterReader is a Reader for the GetVpsDatacenter structure.
type GetVpsDatacenterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVpsDatacenterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVpsDatacenterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetVpsDatacenterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVpsDatacenterOK creates a GetVpsDatacenterOK with default headers values
func NewGetVpsDatacenterOK() *GetVpsDatacenterOK {
	return &GetVpsDatacenterOK{}
}

/*GetVpsDatacenterOK handles this case with default header values.

return value
*/
type GetVpsDatacenterOK struct {
	Payload []string
}

func (o *GetVpsDatacenterOK) Error() string {
	return fmt.Sprintf("[GET /vps/datacenter][%d] getVpsDatacenterOK  %+v", 200, o.Payload)
}

func (o *GetVpsDatacenterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVpsDatacenterDefault creates a GetVpsDatacenterDefault with default headers values
func NewGetVpsDatacenterDefault(code int) *GetVpsDatacenterDefault {
	return &GetVpsDatacenterDefault{
		_statusCode: code,
	}
}

/*GetVpsDatacenterDefault handles this case with default header values.

Unexpected error
*/
type GetVpsDatacenterDefault struct {
	_statusCode int

	Payload *models.GetVpsDatacenterDefaultBody
}

// Code gets the status code for the get vps datacenter default response
func (o *GetVpsDatacenterDefault) Code() int {
	return o._statusCode
}

func (o *GetVpsDatacenterDefault) Error() string {
	return fmt.Sprintf("[GET /vps/datacenter][%d] GetVpsDatacenter default  %+v", o._statusCode, o.Payload)
}

func (o *GetVpsDatacenterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetVpsDatacenterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
