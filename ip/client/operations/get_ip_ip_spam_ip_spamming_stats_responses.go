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

// GetIPIPSpamIPSpammingStatsReader is a Reader for the GetIPIPSpamIPSpammingStats structure.
type GetIPIPSpamIPSpammingStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPIPSpamIPSpammingStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIPIPSpamIPSpammingStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetIPIPSpamIPSpammingStatsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIPIPSpamIPSpammingStatsOK creates a GetIPIPSpamIPSpammingStatsOK with default headers values
func NewGetIPIPSpamIPSpammingStatsOK() *GetIPIPSpamIPSpammingStatsOK {
	return &GetIPIPSpamIPSpammingStatsOK{}
}

/*GetIPIPSpamIPSpammingStatsOK handles this case with default header values.

return value
*/
type GetIPIPSpamIPSpammingStatsOK struct {
	Payload models.GetIPIPSpamIPSpammingStatsOKBody
}

func (o *GetIPIPSpamIPSpammingStatsOK) Error() string {
	return fmt.Sprintf("[GET /ip/{ip}/spam/{ipSpamming}/stats][%d] getIpIpSpamIpSpammingStatsOK  %+v", 200, o.Payload)
}

func (o *GetIPIPSpamIPSpammingStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPIPSpamIPSpammingStatsDefault creates a GetIPIPSpamIPSpammingStatsDefault with default headers values
func NewGetIPIPSpamIPSpammingStatsDefault(code int) *GetIPIPSpamIPSpammingStatsDefault {
	return &GetIPIPSpamIPSpammingStatsDefault{
		_statusCode: code,
	}
}

/*GetIPIPSpamIPSpammingStatsDefault handles this case with default header values.

Unexpected error
*/
type GetIPIPSpamIPSpammingStatsDefault struct {
	_statusCode int

	Payload *models.GetIPIPSpamIPSpammingStatsDefaultBody
}

// Code gets the status code for the get IP IP spam IP spamming stats default response
func (o *GetIPIPSpamIPSpammingStatsDefault) Code() int {
	return o._statusCode
}

func (o *GetIPIPSpamIPSpammingStatsDefault) Error() string {
	return fmt.Sprintf("[GET /ip/{ip}/spam/{ipSpamming}/stats][%d] GetIPIPSpamIPSpammingStats default  %+v", o._statusCode, o.Payload)
}

func (o *GetIPIPSpamIPSpammingStatsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetIPIPSpamIPSpammingStatsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}