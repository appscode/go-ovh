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

	"github.com/appscode/go-ovh/ip/models"
)

// GetIPIPDelegationReader is a Reader for the GetIPIPDelegation structure.
type GetIPIPDelegationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPIPDelegationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIPIPDelegationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetIPIPDelegationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIPIPDelegationOK creates a GetIPIPDelegationOK with default headers values
func NewGetIPIPDelegationOK() *GetIPIPDelegationOK {
	return &GetIPIPDelegationOK{}
}

/*GetIPIPDelegationOK handles this case with default header values.

return value
*/
type GetIPIPDelegationOK struct {
	Payload []string
}

func (o *GetIPIPDelegationOK) Error() string {
	return fmt.Sprintf("[GET /ip/{ip}/delegation][%d] getIpIpDelegationOK  %+v", 200, o.Payload)
}

func (o *GetIPIPDelegationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPIPDelegationDefault creates a GetIPIPDelegationDefault with default headers values
func NewGetIPIPDelegationDefault(code int) *GetIPIPDelegationDefault {
	return &GetIPIPDelegationDefault{
		_statusCode: code,
	}
}

/*GetIPIPDelegationDefault handles this case with default header values.

Unexpected error
*/
type GetIPIPDelegationDefault struct {
	_statusCode int

	Payload *models.GetIPIPDelegationDefaultBody
}

// Code gets the status code for the get IP IP delegation default response
func (o *GetIPIPDelegationDefault) Code() int {
	return o._statusCode
}

func (o *GetIPIPDelegationDefault) Error() string {
	return fmt.Sprintf("[GET /ip/{ip}/delegation][%d] GetIPIPDelegation default  %+v", o._statusCode, o.Payload)
}

func (o *GetIPIPDelegationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetIPIPDelegationDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}