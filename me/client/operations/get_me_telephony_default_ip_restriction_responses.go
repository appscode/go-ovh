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

// GetMeTelephonyDefaultIPRestrictionReader is a Reader for the GetMeTelephonyDefaultIPRestriction structure.
type GetMeTelephonyDefaultIPRestrictionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeTelephonyDefaultIPRestrictionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeTelephonyDefaultIPRestrictionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeTelephonyDefaultIPRestrictionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeTelephonyDefaultIPRestrictionOK creates a GetMeTelephonyDefaultIPRestrictionOK with default headers values
func NewGetMeTelephonyDefaultIPRestrictionOK() *GetMeTelephonyDefaultIPRestrictionOK {
	return &GetMeTelephonyDefaultIPRestrictionOK{}
}

/*GetMeTelephonyDefaultIPRestrictionOK handles this case with default header values.

return value
*/
type GetMeTelephonyDefaultIPRestrictionOK struct {
	Payload []int64
}

func (o *GetMeTelephonyDefaultIPRestrictionOK) Error() string {
	return fmt.Sprintf("[GET /me/telephony/defaultIpRestriction][%d] getMeTelephonyDefaultIpRestrictionOK  %+v", 200, o.Payload)
}

func (o *GetMeTelephonyDefaultIPRestrictionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeTelephonyDefaultIPRestrictionDefault creates a GetMeTelephonyDefaultIPRestrictionDefault with default headers values
func NewGetMeTelephonyDefaultIPRestrictionDefault(code int) *GetMeTelephonyDefaultIPRestrictionDefault {
	return &GetMeTelephonyDefaultIPRestrictionDefault{
		_statusCode: code,
	}
}

/*GetMeTelephonyDefaultIPRestrictionDefault handles this case with default header values.

Unexpected error
*/
type GetMeTelephonyDefaultIPRestrictionDefault struct {
	_statusCode int

	Payload *models.GetMeTelephonyDefaultIPRestrictionDefaultBody
}

// Code gets the status code for the get me telephony default IP restriction default response
func (o *GetMeTelephonyDefaultIPRestrictionDefault) Code() int {
	return o._statusCode
}

func (o *GetMeTelephonyDefaultIPRestrictionDefault) Error() string {
	return fmt.Sprintf("[GET /me/telephony/defaultIpRestriction][%d] GetMeTelephonyDefaultIPRestriction default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeTelephonyDefaultIPRestrictionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeTelephonyDefaultIPRestrictionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
