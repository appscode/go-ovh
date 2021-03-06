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

// PutMePaymentMeanDeferredPaymentAccountIDReader is a Reader for the PutMePaymentMeanDeferredPaymentAccountID structure.
type PutMePaymentMeanDeferredPaymentAccountIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutMePaymentMeanDeferredPaymentAccountIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutMePaymentMeanDeferredPaymentAccountIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutMePaymentMeanDeferredPaymentAccountIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutMePaymentMeanDeferredPaymentAccountIDOK creates a PutMePaymentMeanDeferredPaymentAccountIDOK with default headers values
func NewPutMePaymentMeanDeferredPaymentAccountIDOK() *PutMePaymentMeanDeferredPaymentAccountIDOK {
	return &PutMePaymentMeanDeferredPaymentAccountIDOK{}
}

/*PutMePaymentMeanDeferredPaymentAccountIDOK handles this case with default header values.

return 'void'
*/
type PutMePaymentMeanDeferredPaymentAccountIDOK struct {
}

func (o *PutMePaymentMeanDeferredPaymentAccountIDOK) Error() string {
	return fmt.Sprintf("[PUT /me/paymentMean/deferredPaymentAccount/{id}][%d] putMePaymentMeanDeferredPaymentAccountIdOK ", 200)
}

func (o *PutMePaymentMeanDeferredPaymentAccountIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutMePaymentMeanDeferredPaymentAccountIDDefault creates a PutMePaymentMeanDeferredPaymentAccountIDDefault with default headers values
func NewPutMePaymentMeanDeferredPaymentAccountIDDefault(code int) *PutMePaymentMeanDeferredPaymentAccountIDDefault {
	return &PutMePaymentMeanDeferredPaymentAccountIDDefault{
		_statusCode: code,
	}
}

/*PutMePaymentMeanDeferredPaymentAccountIDDefault handles this case with default header values.

Unexpected error
*/
type PutMePaymentMeanDeferredPaymentAccountIDDefault struct {
	_statusCode int

	Payload *models.PutMePaymentMeanDeferredPaymentAccountIDDefaultBody
}

// Code gets the status code for the put me payment mean deferred payment account ID default response
func (o *PutMePaymentMeanDeferredPaymentAccountIDDefault) Code() int {
	return o._statusCode
}

func (o *PutMePaymentMeanDeferredPaymentAccountIDDefault) Error() string {
	return fmt.Sprintf("[PUT /me/paymentMean/deferredPaymentAccount/{id}][%d] PutMePaymentMeanDeferredPaymentAccountID default  %+v", o._statusCode, o.Payload)
}

func (o *PutMePaymentMeanDeferredPaymentAccountIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutMePaymentMeanDeferredPaymentAccountIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
