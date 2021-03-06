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

// PostMePaymentMeanBankAccountReader is a Reader for the PostMePaymentMeanBankAccount structure.
type PostMePaymentMeanBankAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMePaymentMeanBankAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMePaymentMeanBankAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostMePaymentMeanBankAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMePaymentMeanBankAccountOK creates a PostMePaymentMeanBankAccountOK with default headers values
func NewPostMePaymentMeanBankAccountOK() *PostMePaymentMeanBankAccountOK {
	return &PostMePaymentMeanBankAccountOK{}
}

/*PostMePaymentMeanBankAccountOK handles this case with default header values.

description of 'billing.PaymentMeanValidation' response
*/
type PostMePaymentMeanBankAccountOK struct {
	Payload *models.BillingPaymentMeanValidation
}

func (o *PostMePaymentMeanBankAccountOK) Error() string {
	return fmt.Sprintf("[POST /me/paymentMean/bankAccount][%d] postMePaymentMeanBankAccountOK  %+v", 200, o.Payload)
}

func (o *PostMePaymentMeanBankAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingPaymentMeanValidation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMePaymentMeanBankAccountDefault creates a PostMePaymentMeanBankAccountDefault with default headers values
func NewPostMePaymentMeanBankAccountDefault(code int) *PostMePaymentMeanBankAccountDefault {
	return &PostMePaymentMeanBankAccountDefault{
		_statusCode: code,
	}
}

/*PostMePaymentMeanBankAccountDefault handles this case with default header values.

Unexpected error
*/
type PostMePaymentMeanBankAccountDefault struct {
	_statusCode int

	Payload *models.PostMePaymentMeanBankAccountDefaultBody
}

// Code gets the status code for the post me payment mean bank account default response
func (o *PostMePaymentMeanBankAccountDefault) Code() int {
	return o._statusCode
}

func (o *PostMePaymentMeanBankAccountDefault) Error() string {
	return fmt.Sprintf("[POST /me/paymentMean/bankAccount][%d] PostMePaymentMeanBankAccount default  %+v", o._statusCode, o.Payload)
}

func (o *PostMePaymentMeanBankAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMePaymentMeanBankAccountDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
