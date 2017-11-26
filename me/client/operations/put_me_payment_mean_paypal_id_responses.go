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

// PutMePaymentMeanPaypalIDReader is a Reader for the PutMePaymentMeanPaypalID structure.
type PutMePaymentMeanPaypalIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutMePaymentMeanPaypalIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutMePaymentMeanPaypalIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutMePaymentMeanPaypalIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutMePaymentMeanPaypalIDOK creates a PutMePaymentMeanPaypalIDOK with default headers values
func NewPutMePaymentMeanPaypalIDOK() *PutMePaymentMeanPaypalIDOK {
	return &PutMePaymentMeanPaypalIDOK{}
}

/*PutMePaymentMeanPaypalIDOK handles this case with default header values.

return 'void'
*/
type PutMePaymentMeanPaypalIDOK struct {
}

func (o *PutMePaymentMeanPaypalIDOK) Error() string {
	return fmt.Sprintf("[PUT /me/paymentMean/paypal/{id}][%d] putMePaymentMeanPaypalIdOK ", 200)
}

func (o *PutMePaymentMeanPaypalIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutMePaymentMeanPaypalIDDefault creates a PutMePaymentMeanPaypalIDDefault with default headers values
func NewPutMePaymentMeanPaypalIDDefault(code int) *PutMePaymentMeanPaypalIDDefault {
	return &PutMePaymentMeanPaypalIDDefault{
		_statusCode: code,
	}
}

/*PutMePaymentMeanPaypalIDDefault handles this case with default header values.

Unexpected error
*/
type PutMePaymentMeanPaypalIDDefault struct {
	_statusCode int

	Payload *models.PutMePaymentMeanPaypalIDDefaultBody
}

// Code gets the status code for the put me payment mean paypal ID default response
func (o *PutMePaymentMeanPaypalIDDefault) Code() int {
	return o._statusCode
}

func (o *PutMePaymentMeanPaypalIDDefault) Error() string {
	return fmt.Sprintf("[PUT /me/paymentMean/paypal/{id}][%d] PutMePaymentMeanPaypalID default  %+v", o._statusCode, o.Payload)
}

func (o *PutMePaymentMeanPaypalIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutMePaymentMeanPaypalIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
