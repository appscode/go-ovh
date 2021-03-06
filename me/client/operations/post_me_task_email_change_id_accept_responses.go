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

// PostMeTaskEmailChangeIDAcceptReader is a Reader for the PostMeTaskEmailChangeIDAccept structure.
type PostMeTaskEmailChangeIDAcceptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMeTaskEmailChangeIDAcceptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMeTaskEmailChangeIDAcceptOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostMeTaskEmailChangeIDAcceptDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMeTaskEmailChangeIDAcceptOK creates a PostMeTaskEmailChangeIDAcceptOK with default headers values
func NewPostMeTaskEmailChangeIDAcceptOK() *PostMeTaskEmailChangeIDAcceptOK {
	return &PostMeTaskEmailChangeIDAcceptOK{}
}

/*PostMeTaskEmailChangeIDAcceptOK handles this case with default header values.

return 'void'
*/
type PostMeTaskEmailChangeIDAcceptOK struct {
}

func (o *PostMeTaskEmailChangeIDAcceptOK) Error() string {
	return fmt.Sprintf("[POST /me/task/emailChange/{id}/accept][%d] postMeTaskEmailChangeIdAcceptOK ", 200)
}

func (o *PostMeTaskEmailChangeIDAcceptOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMeTaskEmailChangeIDAcceptDefault creates a PostMeTaskEmailChangeIDAcceptDefault with default headers values
func NewPostMeTaskEmailChangeIDAcceptDefault(code int) *PostMeTaskEmailChangeIDAcceptDefault {
	return &PostMeTaskEmailChangeIDAcceptDefault{
		_statusCode: code,
	}
}

/*PostMeTaskEmailChangeIDAcceptDefault handles this case with default header values.

Unexpected error
*/
type PostMeTaskEmailChangeIDAcceptDefault struct {
	_statusCode int

	Payload *models.PostMeTaskEmailChangeIDAcceptDefaultBody
}

// Code gets the status code for the post me task email change ID accept default response
func (o *PostMeTaskEmailChangeIDAcceptDefault) Code() int {
	return o._statusCode
}

func (o *PostMeTaskEmailChangeIDAcceptDefault) Error() string {
	return fmt.Sprintf("[POST /me/task/emailChange/{id}/accept][%d] PostMeTaskEmailChangeIDAccept default  %+v", o._statusCode, o.Payload)
}

func (o *PostMeTaskEmailChangeIDAcceptDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMeTaskEmailChangeIDAcceptDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
