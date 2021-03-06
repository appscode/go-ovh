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

// PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanReader is a Reader for the PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMean structure.
type PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanOK creates a PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanOK with default headers values
func NewPostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanOK() *PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanOK {
	return &PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanOK{}
}

/*PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanOK handles this case with default header values.

return 'void'
*/
type PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanOK struct {
}

func (o *PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanOK) Error() string {
	return fmt.Sprintf("[POST /me/paymentMean/creditCard/{id}/chooseAsDefaultPaymentMean][%d] postMePaymentMeanCreditCardIdChooseAsDefaultPaymentMeanOK ", 200)
}

func (o *PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanDefault creates a PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanDefault with default headers values
func NewPostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanDefault(code int) *PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanDefault {
	return &PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanDefault{
		_statusCode: code,
	}
}

/*PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanDefault handles this case with default header values.

Unexpected error
*/
type PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanDefault struct {
	_statusCode int

	Payload *models.PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanDefaultBody
}

// Code gets the status code for the post me payment mean credit card ID choose as default payment mean default response
func (o *PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanDefault) Code() int {
	return o._statusCode
}

func (o *PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanDefault) Error() string {
	return fmt.Sprintf("[POST /me/paymentMean/creditCard/{id}/chooseAsDefaultPaymentMean][%d] PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMean default  %+v", o._statusCode, o.Payload)
}

func (o *PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMePaymentMeanCreditCardIDChooseAsDefaultPaymentMeanDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
