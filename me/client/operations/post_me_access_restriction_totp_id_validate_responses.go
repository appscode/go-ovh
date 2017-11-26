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

// PostMeAccessRestrictionTotpIDValidateReader is a Reader for the PostMeAccessRestrictionTotpIDValidate structure.
type PostMeAccessRestrictionTotpIDValidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMeAccessRestrictionTotpIDValidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMeAccessRestrictionTotpIDValidateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostMeAccessRestrictionTotpIDValidateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMeAccessRestrictionTotpIDValidateOK creates a PostMeAccessRestrictionTotpIDValidateOK with default headers values
func NewPostMeAccessRestrictionTotpIDValidateOK() *PostMeAccessRestrictionTotpIDValidateOK {
	return &PostMeAccessRestrictionTotpIDValidateOK{}
}

/*PostMeAccessRestrictionTotpIDValidateOK handles this case with default header values.

return 'void'
*/
type PostMeAccessRestrictionTotpIDValidateOK struct {
}

func (o *PostMeAccessRestrictionTotpIDValidateOK) Error() string {
	return fmt.Sprintf("[POST /me/accessRestriction/totp/{id}/validate][%d] postMeAccessRestrictionTotpIdValidateOK ", 200)
}

func (o *PostMeAccessRestrictionTotpIDValidateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMeAccessRestrictionTotpIDValidateDefault creates a PostMeAccessRestrictionTotpIDValidateDefault with default headers values
func NewPostMeAccessRestrictionTotpIDValidateDefault(code int) *PostMeAccessRestrictionTotpIDValidateDefault {
	return &PostMeAccessRestrictionTotpIDValidateDefault{
		_statusCode: code,
	}
}

/*PostMeAccessRestrictionTotpIDValidateDefault handles this case with default header values.

Unexpected error
*/
type PostMeAccessRestrictionTotpIDValidateDefault struct {
	_statusCode int

	Payload *models.PostMeAccessRestrictionTotpIDValidateDefaultBody
}

// Code gets the status code for the post me access restriction totp ID validate default response
func (o *PostMeAccessRestrictionTotpIDValidateDefault) Code() int {
	return o._statusCode
}

func (o *PostMeAccessRestrictionTotpIDValidateDefault) Error() string {
	return fmt.Sprintf("[POST /me/accessRestriction/totp/{id}/validate][%d] PostMeAccessRestrictionTotpIDValidate default  %+v", o._statusCode, o.Payload)
}

func (o *PostMeAccessRestrictionTotpIDValidateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMeAccessRestrictionTotpIDValidateDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
