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

// GetMeSubscriptionReader is a Reader for the GetMeSubscription structure.
type GetMeSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeSubscriptionOK creates a GetMeSubscriptionOK with default headers values
func NewGetMeSubscriptionOK() *GetMeSubscriptionOK {
	return &GetMeSubscriptionOK{}
}

/*GetMeSubscriptionOK handles this case with default header values.

return value
*/
type GetMeSubscriptionOK struct {
	Payload []string
}

func (o *GetMeSubscriptionOK) Error() string {
	return fmt.Sprintf("[GET /me/subscription][%d] getMeSubscriptionOK  %+v", 200, o.Payload)
}

func (o *GetMeSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeSubscriptionDefault creates a GetMeSubscriptionDefault with default headers values
func NewGetMeSubscriptionDefault(code int) *GetMeSubscriptionDefault {
	return &GetMeSubscriptionDefault{
		_statusCode: code,
	}
}

/*GetMeSubscriptionDefault handles this case with default header values.

Unexpected error
*/
type GetMeSubscriptionDefault struct {
	_statusCode int

	Payload *models.GetMeSubscriptionDefaultBody
}

// Code gets the status code for the get me subscription default response
func (o *GetMeSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *GetMeSubscriptionDefault) Error() string {
	return fmt.Sprintf("[GET /me/subscription][%d] GetMeSubscription default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeSubscriptionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
