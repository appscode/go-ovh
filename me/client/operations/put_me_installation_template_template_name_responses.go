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

// PutMeInstallationTemplateTemplateNameReader is a Reader for the PutMeInstallationTemplateTemplateName structure.
type PutMeInstallationTemplateTemplateNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutMeInstallationTemplateTemplateNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutMeInstallationTemplateTemplateNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutMeInstallationTemplateTemplateNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutMeInstallationTemplateTemplateNameOK creates a PutMeInstallationTemplateTemplateNameOK with default headers values
func NewPutMeInstallationTemplateTemplateNameOK() *PutMeInstallationTemplateTemplateNameOK {
	return &PutMeInstallationTemplateTemplateNameOK{}
}

/*PutMeInstallationTemplateTemplateNameOK handles this case with default header values.

return 'void'
*/
type PutMeInstallationTemplateTemplateNameOK struct {
}

func (o *PutMeInstallationTemplateTemplateNameOK) Error() string {
	return fmt.Sprintf("[PUT /me/installationTemplate/{templateName}][%d] putMeInstallationTemplateTemplateNameOK ", 200)
}

func (o *PutMeInstallationTemplateTemplateNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutMeInstallationTemplateTemplateNameDefault creates a PutMeInstallationTemplateTemplateNameDefault with default headers values
func NewPutMeInstallationTemplateTemplateNameDefault(code int) *PutMeInstallationTemplateTemplateNameDefault {
	return &PutMeInstallationTemplateTemplateNameDefault{
		_statusCode: code,
	}
}

/*PutMeInstallationTemplateTemplateNameDefault handles this case with default header values.

Unexpected error
*/
type PutMeInstallationTemplateTemplateNameDefault struct {
	_statusCode int

	Payload *models.PutMeInstallationTemplateTemplateNameDefaultBody
}

// Code gets the status code for the put me installation template template name default response
func (o *PutMeInstallationTemplateTemplateNameDefault) Code() int {
	return o._statusCode
}

func (o *PutMeInstallationTemplateTemplateNameDefault) Error() string {
	return fmt.Sprintf("[PUT /me/installationTemplate/{templateName}][%d] PutMeInstallationTemplateTemplateName default  %+v", o._statusCode, o.Payload)
}

func (o *PutMeInstallationTemplateTemplateNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutMeInstallationTemplateTemplateNameDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
