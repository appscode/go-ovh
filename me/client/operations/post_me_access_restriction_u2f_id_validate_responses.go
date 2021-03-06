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

// PostMeAccessRestrictionU2fIDValidateReader is a Reader for the PostMeAccessRestrictionU2fIDValidate structure.
type PostMeAccessRestrictionU2fIDValidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMeAccessRestrictionU2fIDValidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMeAccessRestrictionU2fIDValidateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostMeAccessRestrictionU2fIDValidateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMeAccessRestrictionU2fIDValidateOK creates a PostMeAccessRestrictionU2fIDValidateOK with default headers values
func NewPostMeAccessRestrictionU2fIDValidateOK() *PostMeAccessRestrictionU2fIDValidateOK {
	return &PostMeAccessRestrictionU2fIDValidateOK{}
}

/*PostMeAccessRestrictionU2fIDValidateOK handles this case with default header values.

return 'void'
*/
type PostMeAccessRestrictionU2fIDValidateOK struct {
}

func (o *PostMeAccessRestrictionU2fIDValidateOK) Error() string {
	return fmt.Sprintf("[POST /me/accessRestriction/u2f/{id}/validate][%d] postMeAccessRestrictionU2fIdValidateOK ", 200)
}

func (o *PostMeAccessRestrictionU2fIDValidateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMeAccessRestrictionU2fIDValidateDefault creates a PostMeAccessRestrictionU2fIDValidateDefault with default headers values
func NewPostMeAccessRestrictionU2fIDValidateDefault(code int) *PostMeAccessRestrictionU2fIDValidateDefault {
	return &PostMeAccessRestrictionU2fIDValidateDefault{
		_statusCode: code,
	}
}

/*PostMeAccessRestrictionU2fIDValidateDefault handles this case with default header values.

Unexpected error
*/
type PostMeAccessRestrictionU2fIDValidateDefault struct {
	_statusCode int

	Payload *models.PostMeAccessRestrictionU2fIDValidateDefaultBody
}

// Code gets the status code for the post me access restriction u2f ID validate default response
func (o *PostMeAccessRestrictionU2fIDValidateDefault) Code() int {
	return o._statusCode
}

func (o *PostMeAccessRestrictionU2fIDValidateDefault) Error() string {
	return fmt.Sprintf("[POST /me/accessRestriction/u2f/{id}/validate][%d] PostMeAccessRestrictionU2fIDValidate default  %+v", o._statusCode, o.Payload)
}

func (o *PostMeAccessRestrictionU2fIDValidateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMeAccessRestrictionU2fIDValidateDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
