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

// PostMeAccessRestrictionTotpIDEnableReader is a Reader for the PostMeAccessRestrictionTotpIDEnable structure.
type PostMeAccessRestrictionTotpIDEnableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMeAccessRestrictionTotpIDEnableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMeAccessRestrictionTotpIDEnableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostMeAccessRestrictionTotpIDEnableDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMeAccessRestrictionTotpIDEnableOK creates a PostMeAccessRestrictionTotpIDEnableOK with default headers values
func NewPostMeAccessRestrictionTotpIDEnableOK() *PostMeAccessRestrictionTotpIDEnableOK {
	return &PostMeAccessRestrictionTotpIDEnableOK{}
}

/*PostMeAccessRestrictionTotpIDEnableOK handles this case with default header values.

return 'void'
*/
type PostMeAccessRestrictionTotpIDEnableOK struct {
}

func (o *PostMeAccessRestrictionTotpIDEnableOK) Error() string {
	return fmt.Sprintf("[POST /me/accessRestriction/totp/{id}/enable][%d] postMeAccessRestrictionTotpIdEnableOK ", 200)
}

func (o *PostMeAccessRestrictionTotpIDEnableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMeAccessRestrictionTotpIDEnableDefault creates a PostMeAccessRestrictionTotpIDEnableDefault with default headers values
func NewPostMeAccessRestrictionTotpIDEnableDefault(code int) *PostMeAccessRestrictionTotpIDEnableDefault {
	return &PostMeAccessRestrictionTotpIDEnableDefault{
		_statusCode: code,
	}
}

/*PostMeAccessRestrictionTotpIDEnableDefault handles this case with default header values.

Unexpected error
*/
type PostMeAccessRestrictionTotpIDEnableDefault struct {
	_statusCode int

	Payload *models.PostMeAccessRestrictionTotpIDEnableDefaultBody
}

// Code gets the status code for the post me access restriction totp ID enable default response
func (o *PostMeAccessRestrictionTotpIDEnableDefault) Code() int {
	return o._statusCode
}

func (o *PostMeAccessRestrictionTotpIDEnableDefault) Error() string {
	return fmt.Sprintf("[POST /me/accessRestriction/totp/{id}/enable][%d] PostMeAccessRestrictionTotpIDEnable default  %+v", o._statusCode, o.Payload)
}

func (o *PostMeAccessRestrictionTotpIDEnableDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMeAccessRestrictionTotpIDEnableDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
