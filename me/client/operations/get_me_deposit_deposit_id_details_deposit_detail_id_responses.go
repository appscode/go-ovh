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

// GetMeDepositDepositIDDetailsDepositDetailIDReader is a Reader for the GetMeDepositDepositIDDetailsDepositDetailID structure.
type GetMeDepositDepositIDDetailsDepositDetailIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeDepositDepositIDDetailsDepositDetailIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeDepositDepositIDDetailsDepositDetailIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeDepositDepositIDDetailsDepositDetailIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeDepositDepositIDDetailsDepositDetailIDOK creates a GetMeDepositDepositIDDetailsDepositDetailIDOK with default headers values
func NewGetMeDepositDepositIDDetailsDepositDetailIDOK() *GetMeDepositDepositIDDetailsDepositDetailIDOK {
	return &GetMeDepositDepositIDDetailsDepositDetailIDOK{}
}

/*GetMeDepositDepositIDDetailsDepositDetailIDOK handles this case with default header values.

description of 'billing.DepositDetail' response
*/
type GetMeDepositDepositIDDetailsDepositDetailIDOK struct {
	Payload *models.BillingDepositDetail
}

func (o *GetMeDepositDepositIDDetailsDepositDetailIDOK) Error() string {
	return fmt.Sprintf("[GET /me/deposit/{depositId}/details/{depositDetailId}][%d] getMeDepositDepositIdDetailsDepositDetailIdOK  %+v", 200, o.Payload)
}

func (o *GetMeDepositDepositIDDetailsDepositDetailIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingDepositDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeDepositDepositIDDetailsDepositDetailIDDefault creates a GetMeDepositDepositIDDetailsDepositDetailIDDefault with default headers values
func NewGetMeDepositDepositIDDetailsDepositDetailIDDefault(code int) *GetMeDepositDepositIDDetailsDepositDetailIDDefault {
	return &GetMeDepositDepositIDDetailsDepositDetailIDDefault{
		_statusCode: code,
	}
}

/*GetMeDepositDepositIDDetailsDepositDetailIDDefault handles this case with default header values.

Unexpected error
*/
type GetMeDepositDepositIDDetailsDepositDetailIDDefault struct {
	_statusCode int

	Payload *models.GetMeDepositDepositIDDetailsDepositDetailIDDefaultBody
}

// Code gets the status code for the get me deposit deposit ID details deposit detail ID default response
func (o *GetMeDepositDepositIDDetailsDepositDetailIDDefault) Code() int {
	return o._statusCode
}

func (o *GetMeDepositDepositIDDetailsDepositDetailIDDefault) Error() string {
	return fmt.Sprintf("[GET /me/deposit/{depositId}/details/{depositDetailId}][%d] GetMeDepositDepositIDDetailsDepositDetailID default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeDepositDepositIDDetailsDepositDetailIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeDepositDepositIDDetailsDepositDetailIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
