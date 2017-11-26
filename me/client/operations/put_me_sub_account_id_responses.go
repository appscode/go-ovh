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

// PutMeSubAccountIDReader is a Reader for the PutMeSubAccountID structure.
type PutMeSubAccountIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutMeSubAccountIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutMeSubAccountIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutMeSubAccountIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutMeSubAccountIDOK creates a PutMeSubAccountIDOK with default headers values
func NewPutMeSubAccountIDOK() *PutMeSubAccountIDOK {
	return &PutMeSubAccountIDOK{}
}

/*PutMeSubAccountIDOK handles this case with default header values.

return 'void'
*/
type PutMeSubAccountIDOK struct {
}

func (o *PutMeSubAccountIDOK) Error() string {
	return fmt.Sprintf("[PUT /me/subAccount/{id}][%d] putMeSubAccountIdOK ", 200)
}

func (o *PutMeSubAccountIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutMeSubAccountIDDefault creates a PutMeSubAccountIDDefault with default headers values
func NewPutMeSubAccountIDDefault(code int) *PutMeSubAccountIDDefault {
	return &PutMeSubAccountIDDefault{
		_statusCode: code,
	}
}

/*PutMeSubAccountIDDefault handles this case with default header values.

Unexpected error
*/
type PutMeSubAccountIDDefault struct {
	_statusCode int

	Payload *models.PutMeSubAccountIDDefaultBody
}

// Code gets the status code for the put me sub account ID default response
func (o *PutMeSubAccountIDDefault) Code() int {
	return o._statusCode
}

func (o *PutMeSubAccountIDDefault) Error() string {
	return fmt.Sprintf("[PUT /me/subAccount/{id}][%d] PutMeSubAccountID default  %+v", o._statusCode, o.Payload)
}

func (o *PutMeSubAccountIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutMeSubAccountIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
