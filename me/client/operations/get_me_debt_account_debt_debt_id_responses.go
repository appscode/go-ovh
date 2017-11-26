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

// GetMeDebtAccountDebtDebtIDReader is a Reader for the GetMeDebtAccountDebtDebtID structure.
type GetMeDebtAccountDebtDebtIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeDebtAccountDebtDebtIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeDebtAccountDebtDebtIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeDebtAccountDebtDebtIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeDebtAccountDebtDebtIDOK creates a GetMeDebtAccountDebtDebtIDOK with default headers values
func NewGetMeDebtAccountDebtDebtIDOK() *GetMeDebtAccountDebtDebtIDOK {
	return &GetMeDebtAccountDebtDebtIDOK{}
}

/*GetMeDebtAccountDebtDebtIDOK handles this case with default header values.

description of 'debt.Debt' response
*/
type GetMeDebtAccountDebtDebtIDOK struct {
	Payload *models.DebtDebt
}

func (o *GetMeDebtAccountDebtDebtIDOK) Error() string {
	return fmt.Sprintf("[GET /me/debtAccount/debt/{debtId}][%d] getMeDebtAccountDebtDebtIdOK  %+v", 200, o.Payload)
}

func (o *GetMeDebtAccountDebtDebtIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DebtDebt)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeDebtAccountDebtDebtIDDefault creates a GetMeDebtAccountDebtDebtIDDefault with default headers values
func NewGetMeDebtAccountDebtDebtIDDefault(code int) *GetMeDebtAccountDebtDebtIDDefault {
	return &GetMeDebtAccountDebtDebtIDDefault{
		_statusCode: code,
	}
}

/*GetMeDebtAccountDebtDebtIDDefault handles this case with default header values.

Unexpected error
*/
type GetMeDebtAccountDebtDebtIDDefault struct {
	_statusCode int

	Payload *models.GetMeDebtAccountDebtDebtIDDefaultBody
}

// Code gets the status code for the get me debt account debt debt ID default response
func (o *GetMeDebtAccountDebtDebtIDDefault) Code() int {
	return o._statusCode
}

func (o *GetMeDebtAccountDebtDebtIDDefault) Error() string {
	return fmt.Sprintf("[GET /me/debtAccount/debt/{debtId}][%d] GetMeDebtAccountDebtDebtID default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeDebtAccountDebtDebtIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeDebtAccountDebtDebtIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
