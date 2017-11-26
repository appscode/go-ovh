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

// GetMeWithdrawalWithdrawalIDPaymentReader is a Reader for the GetMeWithdrawalWithdrawalIDPayment structure.
type GetMeWithdrawalWithdrawalIDPaymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeWithdrawalWithdrawalIDPaymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeWithdrawalWithdrawalIDPaymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeWithdrawalWithdrawalIDPaymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeWithdrawalWithdrawalIDPaymentOK creates a GetMeWithdrawalWithdrawalIDPaymentOK with default headers values
func NewGetMeWithdrawalWithdrawalIDPaymentOK() *GetMeWithdrawalWithdrawalIDPaymentOK {
	return &GetMeWithdrawalWithdrawalIDPaymentOK{}
}

/*GetMeWithdrawalWithdrawalIDPaymentOK handles this case with default header values.

description of 'billing.Payment' response
*/
type GetMeWithdrawalWithdrawalIDPaymentOK struct {
	Payload *models.BillingPayment
}

func (o *GetMeWithdrawalWithdrawalIDPaymentOK) Error() string {
	return fmt.Sprintf("[GET /me/withdrawal/{withdrawalId}/payment][%d] getMeWithdrawalWithdrawalIdPaymentOK  %+v", 200, o.Payload)
}

func (o *GetMeWithdrawalWithdrawalIDPaymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingPayment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeWithdrawalWithdrawalIDPaymentDefault creates a GetMeWithdrawalWithdrawalIDPaymentDefault with default headers values
func NewGetMeWithdrawalWithdrawalIDPaymentDefault(code int) *GetMeWithdrawalWithdrawalIDPaymentDefault {
	return &GetMeWithdrawalWithdrawalIDPaymentDefault{
		_statusCode: code,
	}
}

/*GetMeWithdrawalWithdrawalIDPaymentDefault handles this case with default header values.

Unexpected error
*/
type GetMeWithdrawalWithdrawalIDPaymentDefault struct {
	_statusCode int

	Payload *models.GetMeWithdrawalWithdrawalIDPaymentDefaultBody
}

// Code gets the status code for the get me withdrawal withdrawal ID payment default response
func (o *GetMeWithdrawalWithdrawalIDPaymentDefault) Code() int {
	return o._statusCode
}

func (o *GetMeWithdrawalWithdrawalIDPaymentDefault) Error() string {
	return fmt.Sprintf("[GET /me/withdrawal/{withdrawalId}/payment][%d] GetMeWithdrawalWithdrawalIDPayment default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeWithdrawalWithdrawalIDPaymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeWithdrawalWithdrawalIDPaymentDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
