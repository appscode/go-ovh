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

// PostMeInstallationTemplateReader is a Reader for the PostMeInstallationTemplate structure.
type PostMeInstallationTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMeInstallationTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMeInstallationTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostMeInstallationTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMeInstallationTemplateOK creates a PostMeInstallationTemplateOK with default headers values
func NewPostMeInstallationTemplateOK() *PostMeInstallationTemplateOK {
	return &PostMeInstallationTemplateOK{}
}

/*PostMeInstallationTemplateOK handles this case with default header values.

return 'void'
*/
type PostMeInstallationTemplateOK struct {
}

func (o *PostMeInstallationTemplateOK) Error() string {
	return fmt.Sprintf("[POST /me/installationTemplate][%d] postMeInstallationTemplateOK ", 200)
}

func (o *PostMeInstallationTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMeInstallationTemplateDefault creates a PostMeInstallationTemplateDefault with default headers values
func NewPostMeInstallationTemplateDefault(code int) *PostMeInstallationTemplateDefault {
	return &PostMeInstallationTemplateDefault{
		_statusCode: code,
	}
}

/*PostMeInstallationTemplateDefault handles this case with default header values.

Unexpected error
*/
type PostMeInstallationTemplateDefault struct {
	_statusCode int

	Payload *models.PostMeInstallationTemplateDefaultBody
}

// Code gets the status code for the post me installation template default response
func (o *PostMeInstallationTemplateDefault) Code() int {
	return o._statusCode
}

func (o *PostMeInstallationTemplateDefault) Error() string {
	return fmt.Sprintf("[POST /me/installationTemplate][%d] PostMeInstallationTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *PostMeInstallationTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMeInstallationTemplateDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
