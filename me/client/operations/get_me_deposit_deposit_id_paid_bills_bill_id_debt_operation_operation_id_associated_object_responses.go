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

// GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectReader is a Reader for the GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObject structure.
type GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectOK creates a GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectOK with default headers values
func NewGetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectOK() *GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectOK {
	return &GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectOK{}
}

/*GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectOK handles this case with default header values.

description of 'debt.Entry.AssociatedObject' response
*/
type GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectOK struct {
	Payload *models.DebtEntryAssociatedObject
}

func (o *GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectOK) Error() string {
	return fmt.Sprintf("[GET /me/deposit/{depositId}/paidBills/{billId}/debt/operation/{operationId}/associatedObject][%d] getMeDepositDepositIdPaidBillsBillIdDebtOperationOperationIdAssociatedObjectOK  %+v", 200, o.Payload)
}

func (o *GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DebtEntryAssociatedObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectDefault creates a GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectDefault with default headers values
func NewGetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectDefault(code int) *GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectDefault {
	return &GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectDefault{
		_statusCode: code,
	}
}

/*GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectDefault handles this case with default header values.

Unexpected error
*/
type GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectDefault struct {
	_statusCode int

	Payload *models.GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectDefaultBody
}

// Code gets the status code for the get me deposit deposit ID paid bills bill ID debt operation operation ID associated object default response
func (o *GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectDefault) Code() int {
	return o._statusCode
}

func (o *GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectDefault) Error() string {
	return fmt.Sprintf("[GET /me/deposit/{depositId}/paidBills/{billId}/debt/operation/{operationId}/associatedObject][%d] GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObject default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeDepositDepositIDPaidBillsBillIDDebtOperationOperationIDAssociatedObjectDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
