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

// GetMeNotificationEmailHistoryIDReader is a Reader for the GetMeNotificationEmailHistoryID structure.
type GetMeNotificationEmailHistoryIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeNotificationEmailHistoryIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeNotificationEmailHistoryIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeNotificationEmailHistoryIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeNotificationEmailHistoryIDOK creates a GetMeNotificationEmailHistoryIDOK with default headers values
func NewGetMeNotificationEmailHistoryIDOK() *GetMeNotificationEmailHistoryIDOK {
	return &GetMeNotificationEmailHistoryIDOK{}
}

/*GetMeNotificationEmailHistoryIDOK handles this case with default header values.

description of 'nichandle.EmailNotification' response
*/
type GetMeNotificationEmailHistoryIDOK struct {
	Payload *models.NichandleEmailNotification
}

func (o *GetMeNotificationEmailHistoryIDOK) Error() string {
	return fmt.Sprintf("[GET /me/notification/email/history/{id}][%d] getMeNotificationEmailHistoryIdOK  %+v", 200, o.Payload)
}

func (o *GetMeNotificationEmailHistoryIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NichandleEmailNotification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeNotificationEmailHistoryIDDefault creates a GetMeNotificationEmailHistoryIDDefault with default headers values
func NewGetMeNotificationEmailHistoryIDDefault(code int) *GetMeNotificationEmailHistoryIDDefault {
	return &GetMeNotificationEmailHistoryIDDefault{
		_statusCode: code,
	}
}

/*GetMeNotificationEmailHistoryIDDefault handles this case with default header values.

Unexpected error
*/
type GetMeNotificationEmailHistoryIDDefault struct {
	_statusCode int

	Payload *models.GetMeNotificationEmailHistoryIDDefaultBody
}

// Code gets the status code for the get me notification email history ID default response
func (o *GetMeNotificationEmailHistoryIDDefault) Code() int {
	return o._statusCode
}

func (o *GetMeNotificationEmailHistoryIDDefault) Error() string {
	return fmt.Sprintf("[GET /me/notification/email/history/{id}][%d] GetMeNotificationEmailHistoryID default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeNotificationEmailHistoryIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeNotificationEmailHistoryIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
