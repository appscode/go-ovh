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

// GetIPIPGameIPOnGameRuleIDReader is a Reader for the GetIPIPGameIPOnGameRuleID structure.
type GetIPIPGameIPOnGameRuleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPIPGameIPOnGameRuleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIPIPGameIPOnGameRuleIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetIPIPGameIPOnGameRuleIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIPIPGameIPOnGameRuleIDOK creates a GetIPIPGameIPOnGameRuleIDOK with default headers values
func NewGetIPIPGameIPOnGameRuleIDOK() *GetIPIPGameIPOnGameRuleIDOK {
	return &GetIPIPGameIPOnGameRuleIDOK{}
}

/*GetIPIPGameIPOnGameRuleIDOK handles this case with default header values.

description of 'ip.GameMitigationRule' response
*/
type GetIPIPGameIPOnGameRuleIDOK struct {
	Payload *models.IPGameMitigationRule
}

func (o *GetIPIPGameIPOnGameRuleIDOK) Error() string {
	return fmt.Sprintf("[GET /ip/{ip}/game/{ipOnGame}/rule/{id}][%d] getIpIpGameIpOnGameRuleIdOK  %+v", 200, o.Payload)
}

func (o *GetIPIPGameIPOnGameRuleIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPGameMitigationRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPIPGameIPOnGameRuleIDDefault creates a GetIPIPGameIPOnGameRuleIDDefault with default headers values
func NewGetIPIPGameIPOnGameRuleIDDefault(code int) *GetIPIPGameIPOnGameRuleIDDefault {
	return &GetIPIPGameIPOnGameRuleIDDefault{
		_statusCode: code,
	}
}

/*GetIPIPGameIPOnGameRuleIDDefault handles this case with default header values.

Unexpected error
*/
type GetIPIPGameIPOnGameRuleIDDefault struct {
	_statusCode int

	Payload *models.GetIPIPGameIPOnGameRuleIDDefaultBody
}

// Code gets the status code for the get IP IP game IP on game rule ID default response
func (o *GetIPIPGameIPOnGameRuleIDDefault) Code() int {
	return o._statusCode
}

func (o *GetIPIPGameIPOnGameRuleIDDefault) Error() string {
	return fmt.Sprintf("[GET /ip/{ip}/game/{ipOnGame}/rule/{id}][%d] GetIPIPGameIPOnGameRuleID default  %+v", o._statusCode, o.Payload)
}

func (o *GetIPIPGameIPOnGameRuleIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetIPIPGameIPOnGameRuleIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}