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

// PostMeTaskContactChangeIDResendEmailReader is a Reader for the PostMeTaskContactChangeIDResendEmail structure.
type PostMeTaskContactChangeIDResendEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMeTaskContactChangeIDResendEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMeTaskContactChangeIDResendEmailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostMeTaskContactChangeIDResendEmailDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMeTaskContactChangeIDResendEmailOK creates a PostMeTaskContactChangeIDResendEmailOK with default headers values
func NewPostMeTaskContactChangeIDResendEmailOK() *PostMeTaskContactChangeIDResendEmailOK {
	return &PostMeTaskContactChangeIDResendEmailOK{}
}

/*PostMeTaskContactChangeIDResendEmailOK handles this case with default header values.

return 'void'
*/
type PostMeTaskContactChangeIDResendEmailOK struct {
}

func (o *PostMeTaskContactChangeIDResendEmailOK) Error() string {
	return fmt.Sprintf("[POST /me/task/contactChange/{id}/resendEmail][%d] postMeTaskContactChangeIdResendEmailOK ", 200)
}

func (o *PostMeTaskContactChangeIDResendEmailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMeTaskContactChangeIDResendEmailDefault creates a PostMeTaskContactChangeIDResendEmailDefault with default headers values
func NewPostMeTaskContactChangeIDResendEmailDefault(code int) *PostMeTaskContactChangeIDResendEmailDefault {
	return &PostMeTaskContactChangeIDResendEmailDefault{
		_statusCode: code,
	}
}

/*PostMeTaskContactChangeIDResendEmailDefault handles this case with default header values.

Unexpected error
*/
type PostMeTaskContactChangeIDResendEmailDefault struct {
	_statusCode int

	Payload *models.PostMeTaskContactChangeIDResendEmailDefaultBody
}

// Code gets the status code for the post me task contact change ID resend email default response
func (o *PostMeTaskContactChangeIDResendEmailDefault) Code() int {
	return o._statusCode
}

func (o *PostMeTaskContactChangeIDResendEmailDefault) Error() string {
	return fmt.Sprintf("[POST /me/task/contactChange/{id}/resendEmail][%d] PostMeTaskContactChangeIDResendEmail default  %+v", o._statusCode, o.Payload)
}

func (o *PostMeTaskContactChangeIDResendEmailDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMeTaskContactChangeIDResendEmailDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
