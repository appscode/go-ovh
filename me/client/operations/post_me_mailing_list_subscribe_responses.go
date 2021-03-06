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

// PostMeMailingListSubscribeReader is a Reader for the PostMeMailingListSubscribe structure.
type PostMeMailingListSubscribeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMeMailingListSubscribeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMeMailingListSubscribeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostMeMailingListSubscribeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMeMailingListSubscribeOK creates a PostMeMailingListSubscribeOK with default headers values
func NewPostMeMailingListSubscribeOK() *PostMeMailingListSubscribeOK {
	return &PostMeMailingListSubscribeOK{}
}

/*PostMeMailingListSubscribeOK handles this case with default header values.

return 'void'
*/
type PostMeMailingListSubscribeOK struct {
}

func (o *PostMeMailingListSubscribeOK) Error() string {
	return fmt.Sprintf("[POST /me/mailingList/subscribe][%d] postMeMailingListSubscribeOK ", 200)
}

func (o *PostMeMailingListSubscribeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMeMailingListSubscribeDefault creates a PostMeMailingListSubscribeDefault with default headers values
func NewPostMeMailingListSubscribeDefault(code int) *PostMeMailingListSubscribeDefault {
	return &PostMeMailingListSubscribeDefault{
		_statusCode: code,
	}
}

/*PostMeMailingListSubscribeDefault handles this case with default header values.

Unexpected error
*/
type PostMeMailingListSubscribeDefault struct {
	_statusCode int

	Payload *models.PostMeMailingListSubscribeDefaultBody
}

// Code gets the status code for the post me mailing list subscribe default response
func (o *PostMeMailingListSubscribeDefault) Code() int {
	return o._statusCode
}

func (o *PostMeMailingListSubscribeDefault) Error() string {
	return fmt.Sprintf("[POST /me/mailingList/subscribe][%d] PostMeMailingListSubscribe default  %+v", o._statusCode, o.Payload)
}

func (o *PostMeMailingListSubscribeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMeMailingListSubscribeDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
