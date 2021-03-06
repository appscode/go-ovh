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

// PostIPIPSpamIPSpammingUnblockReader is a Reader for the PostIPIPSpamIPSpammingUnblock structure.
type PostIPIPSpamIPSpammingUnblockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIPIPSpamIPSpammingUnblockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostIPIPSpamIPSpammingUnblockOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostIPIPSpamIPSpammingUnblockDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIPIPSpamIPSpammingUnblockOK creates a PostIPIPSpamIPSpammingUnblockOK with default headers values
func NewPostIPIPSpamIPSpammingUnblockOK() *PostIPIPSpamIPSpammingUnblockOK {
	return &PostIPIPSpamIPSpammingUnblockOK{}
}

/*PostIPIPSpamIPSpammingUnblockOK handles this case with default header values.

description of 'ip.SpamIp' response
*/
type PostIPIPSpamIPSpammingUnblockOK struct {
	Payload *models.IPSpamIP
}

func (o *PostIPIPSpamIPSpammingUnblockOK) Error() string {
	return fmt.Sprintf("[POST /ip/{ip}/spam/{ipSpamming}/unblock][%d] postIpIpSpamIpSpammingUnblockOK  %+v", 200, o.Payload)
}

func (o *PostIPIPSpamIPSpammingUnblockOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPSpamIP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIPIPSpamIPSpammingUnblockDefault creates a PostIPIPSpamIPSpammingUnblockDefault with default headers values
func NewPostIPIPSpamIPSpammingUnblockDefault(code int) *PostIPIPSpamIPSpammingUnblockDefault {
	return &PostIPIPSpamIPSpammingUnblockDefault{
		_statusCode: code,
	}
}

/*PostIPIPSpamIPSpammingUnblockDefault handles this case with default header values.

Unexpected error
*/
type PostIPIPSpamIPSpammingUnblockDefault struct {
	_statusCode int

	Payload *models.PostIPIPSpamIPSpammingUnblockDefaultBody
}

// Code gets the status code for the post IP IP spam IP spamming unblock default response
func (o *PostIPIPSpamIPSpammingUnblockDefault) Code() int {
	return o._statusCode
}

func (o *PostIPIPSpamIPSpammingUnblockDefault) Error() string {
	return fmt.Sprintf("[POST /ip/{ip}/spam/{ipSpamming}/unblock][%d] PostIPIPSpamIPSpammingUnblock default  %+v", o._statusCode, o.Payload)
}

func (o *PostIPIPSpamIPSpammingUnblockDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostIPIPSpamIPSpammingUnblockDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
