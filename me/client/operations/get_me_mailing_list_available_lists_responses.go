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

// GetMeMailingListAvailableListsReader is a Reader for the GetMeMailingListAvailableLists structure.
type GetMeMailingListAvailableListsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeMailingListAvailableListsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeMailingListAvailableListsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeMailingListAvailableListsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeMailingListAvailableListsOK creates a GetMeMailingListAvailableListsOK with default headers values
func NewGetMeMailingListAvailableListsOK() *GetMeMailingListAvailableListsOK {
	return &GetMeMailingListAvailableListsOK{}
}

/*GetMeMailingListAvailableListsOK handles this case with default header values.

return value
*/
type GetMeMailingListAvailableListsOK struct {
	Payload []string
}

func (o *GetMeMailingListAvailableListsOK) Error() string {
	return fmt.Sprintf("[GET /me/mailingList/availableLists][%d] getMeMailingListAvailableListsOK  %+v", 200, o.Payload)
}

func (o *GetMeMailingListAvailableListsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeMailingListAvailableListsDefault creates a GetMeMailingListAvailableListsDefault with default headers values
func NewGetMeMailingListAvailableListsDefault(code int) *GetMeMailingListAvailableListsDefault {
	return &GetMeMailingListAvailableListsDefault{
		_statusCode: code,
	}
}

/*GetMeMailingListAvailableListsDefault handles this case with default header values.

Unexpected error
*/
type GetMeMailingListAvailableListsDefault struct {
	_statusCode int

	Payload *models.GetMeMailingListAvailableListsDefaultBody
}

// Code gets the status code for the get me mailing list available lists default response
func (o *GetMeMailingListAvailableListsDefault) Code() int {
	return o._statusCode
}

func (o *GetMeMailingListAvailableListsDefault) Error() string {
	return fmt.Sprintf("[GET /me/mailingList/availableLists][%d] GetMeMailingListAvailableLists default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeMailingListAvailableListsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeMailingListAvailableListsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
