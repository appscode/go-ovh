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

// PostMeAccessRestrictionTotpIDDisableReader is a Reader for the PostMeAccessRestrictionTotpIDDisable structure.
type PostMeAccessRestrictionTotpIDDisableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMeAccessRestrictionTotpIDDisableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMeAccessRestrictionTotpIDDisableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostMeAccessRestrictionTotpIDDisableDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMeAccessRestrictionTotpIDDisableOK creates a PostMeAccessRestrictionTotpIDDisableOK with default headers values
func NewPostMeAccessRestrictionTotpIDDisableOK() *PostMeAccessRestrictionTotpIDDisableOK {
	return &PostMeAccessRestrictionTotpIDDisableOK{}
}

/*PostMeAccessRestrictionTotpIDDisableOK handles this case with default header values.

return 'void'
*/
type PostMeAccessRestrictionTotpIDDisableOK struct {
}

func (o *PostMeAccessRestrictionTotpIDDisableOK) Error() string {
	return fmt.Sprintf("[POST /me/accessRestriction/totp/{id}/disable][%d] postMeAccessRestrictionTotpIdDisableOK ", 200)
}

func (o *PostMeAccessRestrictionTotpIDDisableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMeAccessRestrictionTotpIDDisableDefault creates a PostMeAccessRestrictionTotpIDDisableDefault with default headers values
func NewPostMeAccessRestrictionTotpIDDisableDefault(code int) *PostMeAccessRestrictionTotpIDDisableDefault {
	return &PostMeAccessRestrictionTotpIDDisableDefault{
		_statusCode: code,
	}
}

/*PostMeAccessRestrictionTotpIDDisableDefault handles this case with default header values.

Unexpected error
*/
type PostMeAccessRestrictionTotpIDDisableDefault struct {
	_statusCode int

	Payload *models.PostMeAccessRestrictionTotpIDDisableDefaultBody
}

// Code gets the status code for the post me access restriction totp ID disable default response
func (o *PostMeAccessRestrictionTotpIDDisableDefault) Code() int {
	return o._statusCode
}

func (o *PostMeAccessRestrictionTotpIDDisableDefault) Error() string {
	return fmt.Sprintf("[POST /me/accessRestriction/totp/{id}/disable][%d] PostMeAccessRestrictionTotpIDDisable default  %+v", o._statusCode, o.Payload)
}

func (o *PostMeAccessRestrictionTotpIDDisableDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMeAccessRestrictionTotpIDDisableDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
