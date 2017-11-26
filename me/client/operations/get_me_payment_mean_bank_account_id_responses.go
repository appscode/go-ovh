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

// GetMePaymentMeanBankAccountIDReader is a Reader for the GetMePaymentMeanBankAccountID structure.
type GetMePaymentMeanBankAccountIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMePaymentMeanBankAccountIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMePaymentMeanBankAccountIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMePaymentMeanBankAccountIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMePaymentMeanBankAccountIDOK creates a GetMePaymentMeanBankAccountIDOK with default headers values
func NewGetMePaymentMeanBankAccountIDOK() *GetMePaymentMeanBankAccountIDOK {
	return &GetMePaymentMeanBankAccountIDOK{}
}

/*GetMePaymentMeanBankAccountIDOK handles this case with default header values.

description of 'billing.BankAccount' response
*/
type GetMePaymentMeanBankAccountIDOK struct {
	Payload *models.BillingBankAccount
}

func (o *GetMePaymentMeanBankAccountIDOK) Error() string {
	return fmt.Sprintf("[GET /me/paymentMean/bankAccount/{id}][%d] getMePaymentMeanBankAccountIdOK  %+v", 200, o.Payload)
}

func (o *GetMePaymentMeanBankAccountIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingBankAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMePaymentMeanBankAccountIDDefault creates a GetMePaymentMeanBankAccountIDDefault with default headers values
func NewGetMePaymentMeanBankAccountIDDefault(code int) *GetMePaymentMeanBankAccountIDDefault {
	return &GetMePaymentMeanBankAccountIDDefault{
		_statusCode: code,
	}
}

/*GetMePaymentMeanBankAccountIDDefault handles this case with default header values.

Unexpected error
*/
type GetMePaymentMeanBankAccountIDDefault struct {
	_statusCode int

	Payload *models.GetMePaymentMeanBankAccountIDDefaultBody
}

// Code gets the status code for the get me payment mean bank account ID default response
func (o *GetMePaymentMeanBankAccountIDDefault) Code() int {
	return o._statusCode
}

func (o *GetMePaymentMeanBankAccountIDDefault) Error() string {
	return fmt.Sprintf("[GET /me/paymentMean/bankAccount/{id}][%d] GetMePaymentMeanBankAccountID default  %+v", o._statusCode, o.Payload)
}

func (o *GetMePaymentMeanBankAccountIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMePaymentMeanBankAccountIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
