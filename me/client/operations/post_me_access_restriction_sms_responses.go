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

// PostMeAccessRestrictionSmsReader is a Reader for the PostMeAccessRestrictionSms structure.
type PostMeAccessRestrictionSmsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMeAccessRestrictionSmsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMeAccessRestrictionSmsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostMeAccessRestrictionSmsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMeAccessRestrictionSmsOK creates a PostMeAccessRestrictionSmsOK with default headers values
func NewPostMeAccessRestrictionSmsOK() *PostMeAccessRestrictionSmsOK {
	return &PostMeAccessRestrictionSmsOK{}
}

/*PostMeAccessRestrictionSmsOK handles this case with default header values.

description of 'nichandle.AccessRestriction.SmsSecret' response
*/
type PostMeAccessRestrictionSmsOK struct {
	Payload *models.NichandleAccessRestrictionSmsSecret
}

func (o *PostMeAccessRestrictionSmsOK) Error() string {
	return fmt.Sprintf("[POST /me/accessRestriction/sms][%d] postMeAccessRestrictionSmsOK  %+v", 200, o.Payload)
}

func (o *PostMeAccessRestrictionSmsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NichandleAccessRestrictionSmsSecret)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMeAccessRestrictionSmsDefault creates a PostMeAccessRestrictionSmsDefault with default headers values
func NewPostMeAccessRestrictionSmsDefault(code int) *PostMeAccessRestrictionSmsDefault {
	return &PostMeAccessRestrictionSmsDefault{
		_statusCode: code,
	}
}

/*PostMeAccessRestrictionSmsDefault handles this case with default header values.

Unexpected error
*/
type PostMeAccessRestrictionSmsDefault struct {
	_statusCode int

	Payload *models.PostMeAccessRestrictionSmsDefaultBody
}

// Code gets the status code for the post me access restriction sms default response
func (o *PostMeAccessRestrictionSmsDefault) Code() int {
	return o._statusCode
}

func (o *PostMeAccessRestrictionSmsDefault) Error() string {
	return fmt.Sprintf("[POST /me/accessRestriction/sms][%d] PostMeAccessRestrictionSms default  %+v", o._statusCode, o.Payload)
}

func (o *PostMeAccessRestrictionSmsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMeAccessRestrictionSmsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
