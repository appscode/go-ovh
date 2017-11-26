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

// PostMeSubAccountReader is a Reader for the PostMeSubAccount structure.
type PostMeSubAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMeSubAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMeSubAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostMeSubAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMeSubAccountOK creates a PostMeSubAccountOK with default headers values
func NewPostMeSubAccountOK() *PostMeSubAccountOK {
	return &PostMeSubAccountOK{}
}

/*PostMeSubAccountOK handles this case with default header values.

return value
*/
type PostMeSubAccountOK struct {
	Payload int64
}

func (o *PostMeSubAccountOK) Error() string {
	return fmt.Sprintf("[POST /me/subAccount][%d] postMeSubAccountOK  %+v", 200, o.Payload)
}

func (o *PostMeSubAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMeSubAccountDefault creates a PostMeSubAccountDefault with default headers values
func NewPostMeSubAccountDefault(code int) *PostMeSubAccountDefault {
	return &PostMeSubAccountDefault{
		_statusCode: code,
	}
}

/*PostMeSubAccountDefault handles this case with default header values.

Unexpected error
*/
type PostMeSubAccountDefault struct {
	_statusCode int

	Payload *models.PostMeSubAccountDefaultBody
}

// Code gets the status code for the post me sub account default response
func (o *PostMeSubAccountDefault) Code() int {
	return o._statusCode
}

func (o *PostMeSubAccountDefault) Error() string {
	return fmt.Sprintf("[POST /me/subAccount][%d] PostMeSubAccount default  %+v", o._statusCode, o.Payload)
}

func (o *PostMeSubAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMeSubAccountDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
