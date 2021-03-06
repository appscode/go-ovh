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

// PostMeInstallationTemplateTemplateNamePartitionSchemeReader is a Reader for the PostMeInstallationTemplateTemplateNamePartitionScheme structure.
type PostMeInstallationTemplateTemplateNamePartitionSchemeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMeInstallationTemplateTemplateNamePartitionSchemeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMeInstallationTemplateTemplateNamePartitionSchemeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostMeInstallationTemplateTemplateNamePartitionSchemeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMeInstallationTemplateTemplateNamePartitionSchemeOK creates a PostMeInstallationTemplateTemplateNamePartitionSchemeOK with default headers values
func NewPostMeInstallationTemplateTemplateNamePartitionSchemeOK() *PostMeInstallationTemplateTemplateNamePartitionSchemeOK {
	return &PostMeInstallationTemplateTemplateNamePartitionSchemeOK{}
}

/*PostMeInstallationTemplateTemplateNamePartitionSchemeOK handles this case with default header values.

return 'void'
*/
type PostMeInstallationTemplateTemplateNamePartitionSchemeOK struct {
}

func (o *PostMeInstallationTemplateTemplateNamePartitionSchemeOK) Error() string {
	return fmt.Sprintf("[POST /me/installationTemplate/{templateName}/partitionScheme][%d] postMeInstallationTemplateTemplateNamePartitionSchemeOK ", 200)
}

func (o *PostMeInstallationTemplateTemplateNamePartitionSchemeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMeInstallationTemplateTemplateNamePartitionSchemeDefault creates a PostMeInstallationTemplateTemplateNamePartitionSchemeDefault with default headers values
func NewPostMeInstallationTemplateTemplateNamePartitionSchemeDefault(code int) *PostMeInstallationTemplateTemplateNamePartitionSchemeDefault {
	return &PostMeInstallationTemplateTemplateNamePartitionSchemeDefault{
		_statusCode: code,
	}
}

/*PostMeInstallationTemplateTemplateNamePartitionSchemeDefault handles this case with default header values.

Unexpected error
*/
type PostMeInstallationTemplateTemplateNamePartitionSchemeDefault struct {
	_statusCode int

	Payload *models.PostMeInstallationTemplateTemplateNamePartitionSchemeDefaultBody
}

// Code gets the status code for the post me installation template template name partition scheme default response
func (o *PostMeInstallationTemplateTemplateNamePartitionSchemeDefault) Code() int {
	return o._statusCode
}

func (o *PostMeInstallationTemplateTemplateNamePartitionSchemeDefault) Error() string {
	return fmt.Sprintf("[POST /me/installationTemplate/{templateName}/partitionScheme][%d] PostMeInstallationTemplateTemplateNamePartitionScheme default  %+v", o._statusCode, o.Payload)
}

func (o *PostMeInstallationTemplateTemplateNamePartitionSchemeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMeInstallationTemplateTemplateNamePartitionSchemeDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
