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

	"github.com/appscode/go-ovh/me/models"
)

// GetMeSLAIDReader is a Reader for the GetMeSLAID structure.
type GetMeSLAIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeSLAIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeSLAIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeSLAIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeSLAIDOK creates a GetMeSLAIDOK with default headers values
func NewGetMeSLAIDOK() *GetMeSLAIDOK {
	return &GetMeSLAIDOK{}
}

/*GetMeSLAIDOK handles this case with default header values.

description of 'billing.SlaOperation' response
*/
type GetMeSLAIDOK struct {
	Payload *models.BillingSLAOperation
}

func (o *GetMeSLAIDOK) Error() string {
	return fmt.Sprintf("[GET /me/sla/{id}][%d] getMeSlaIdOK  %+v", 200, o.Payload)
}

func (o *GetMeSLAIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingSLAOperation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeSLAIDDefault creates a GetMeSLAIDDefault with default headers values
func NewGetMeSLAIDDefault(code int) *GetMeSLAIDDefault {
	return &GetMeSLAIDDefault{
		_statusCode: code,
	}
}

/*GetMeSLAIDDefault handles this case with default header values.

Unexpected error
*/
type GetMeSLAIDDefault struct {
	_statusCode int

	Payload *models.GetMeSLAIDDefaultBody
}

// Code gets the status code for the get me SLA ID default response
func (o *GetMeSLAIDDefault) Code() int {
	return o._statusCode
}

func (o *GetMeSLAIDDefault) Error() string {
	return fmt.Sprintf("[GET /me/sla/{id}][%d] GetMeSLAID default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeSLAIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeSLAIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
