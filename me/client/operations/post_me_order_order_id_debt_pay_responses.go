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

// PostMeOrderOrderIDDebtPayReader is a Reader for the PostMeOrderOrderIDDebtPay structure.
type PostMeOrderOrderIDDebtPayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMeOrderOrderIDDebtPayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMeOrderOrderIDDebtPayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostMeOrderOrderIDDebtPayDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMeOrderOrderIDDebtPayOK creates a PostMeOrderOrderIDDebtPayOK with default headers values
func NewPostMeOrderOrderIDDebtPayOK() *PostMeOrderOrderIDDebtPayOK {
	return &PostMeOrderOrderIDDebtPayOK{}
}

/*PostMeOrderOrderIDDebtPayOK handles this case with default header values.

description of 'billing.Order' response
*/
type PostMeOrderOrderIDDebtPayOK struct {
	Payload *models.BillingOrder
}

func (o *PostMeOrderOrderIDDebtPayOK) Error() string {
	return fmt.Sprintf("[POST /me/order/{orderId}/debt/pay][%d] postMeOrderOrderIdDebtPayOK  %+v", 200, o.Payload)
}

func (o *PostMeOrderOrderIDDebtPayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingOrder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMeOrderOrderIDDebtPayDefault creates a PostMeOrderOrderIDDebtPayDefault with default headers values
func NewPostMeOrderOrderIDDebtPayDefault(code int) *PostMeOrderOrderIDDebtPayDefault {
	return &PostMeOrderOrderIDDebtPayDefault{
		_statusCode: code,
	}
}

/*PostMeOrderOrderIDDebtPayDefault handles this case with default header values.

Unexpected error
*/
type PostMeOrderOrderIDDebtPayDefault struct {
	_statusCode int

	Payload *models.PostMeOrderOrderIDDebtPayDefaultBody
}

// Code gets the status code for the post me order order ID debt pay default response
func (o *PostMeOrderOrderIDDebtPayDefault) Code() int {
	return o._statusCode
}

func (o *PostMeOrderOrderIDDebtPayDefault) Error() string {
	return fmt.Sprintf("[POST /me/order/{orderId}/debt/pay][%d] PostMeOrderOrderIDDebtPay default  %+v", o._statusCode, o.Payload)
}

func (o *PostMeOrderOrderIDDebtPayDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMeOrderOrderIDDebtPayDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
