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

// GetMeAccessRestrictionSmsIDReader is a Reader for the GetMeAccessRestrictionSmsID structure.
type GetMeAccessRestrictionSmsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeAccessRestrictionSmsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeAccessRestrictionSmsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeAccessRestrictionSmsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeAccessRestrictionSmsIDOK creates a GetMeAccessRestrictionSmsIDOK with default headers values
func NewGetMeAccessRestrictionSmsIDOK() *GetMeAccessRestrictionSmsIDOK {
	return &GetMeAccessRestrictionSmsIDOK{}
}

/*GetMeAccessRestrictionSmsIDOK handles this case with default header values.

description of 'nichandle.AccessRestriction.SmsAccount' response
*/
type GetMeAccessRestrictionSmsIDOK struct {
	Payload *models.NichandleAccessRestrictionSmsAccount
}

func (o *GetMeAccessRestrictionSmsIDOK) Error() string {
	return fmt.Sprintf("[GET /me/accessRestriction/sms/{id}][%d] getMeAccessRestrictionSmsIdOK  %+v", 200, o.Payload)
}

func (o *GetMeAccessRestrictionSmsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NichandleAccessRestrictionSmsAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeAccessRestrictionSmsIDDefault creates a GetMeAccessRestrictionSmsIDDefault with default headers values
func NewGetMeAccessRestrictionSmsIDDefault(code int) *GetMeAccessRestrictionSmsIDDefault {
	return &GetMeAccessRestrictionSmsIDDefault{
		_statusCode: code,
	}
}

/*GetMeAccessRestrictionSmsIDDefault handles this case with default header values.

Unexpected error
*/
type GetMeAccessRestrictionSmsIDDefault struct {
	_statusCode int

	Payload *models.GetMeAccessRestrictionSmsIDDefaultBody
}

// Code gets the status code for the get me access restriction sms ID default response
func (o *GetMeAccessRestrictionSmsIDDefault) Code() int {
	return o._statusCode
}

func (o *GetMeAccessRestrictionSmsIDDefault) Error() string {
	return fmt.Sprintf("[GET /me/accessRestriction/sms/{id}][%d] GetMeAccessRestrictionSmsID default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeAccessRestrictionSmsIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeAccessRestrictionSmsIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
