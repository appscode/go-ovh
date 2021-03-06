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

// PostIPIPFirewallReader is a Reader for the PostIPIPFirewall structure.
type PostIPIPFirewallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIPIPFirewallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostIPIPFirewallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostIPIPFirewallDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIPIPFirewallOK creates a PostIPIPFirewallOK with default headers values
func NewPostIPIPFirewallOK() *PostIPIPFirewallOK {
	return &PostIPIPFirewallOK{}
}

/*PostIPIPFirewallOK handles this case with default header values.

description of 'ip.FirewallIp' response
*/
type PostIPIPFirewallOK struct {
	Payload *models.IPFirewallIP
}

func (o *PostIPIPFirewallOK) Error() string {
	return fmt.Sprintf("[POST /ip/{ip}/firewall][%d] postIpIpFirewallOK  %+v", 200, o.Payload)
}

func (o *PostIPIPFirewallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPFirewallIP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIPIPFirewallDefault creates a PostIPIPFirewallDefault with default headers values
func NewPostIPIPFirewallDefault(code int) *PostIPIPFirewallDefault {
	return &PostIPIPFirewallDefault{
		_statusCode: code,
	}
}

/*PostIPIPFirewallDefault handles this case with default header values.

Unexpected error
*/
type PostIPIPFirewallDefault struct {
	_statusCode int

	Payload *models.PostIPIPFirewallDefaultBody
}

// Code gets the status code for the post IP IP firewall default response
func (o *PostIPIPFirewallDefault) Code() int {
	return o._statusCode
}

func (o *PostIPIPFirewallDefault) Error() string {
	return fmt.Sprintf("[POST /ip/{ip}/firewall][%d] PostIPIPFirewall default  %+v", o._statusCode, o.Payload)
}

func (o *PostIPIPFirewallDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostIPIPFirewallDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
