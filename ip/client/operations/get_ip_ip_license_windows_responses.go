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

// GetIPIPLicenseWindowsReader is a Reader for the GetIPIPLicenseWindows structure.
type GetIPIPLicenseWindowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPIPLicenseWindowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIPIPLicenseWindowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetIPIPLicenseWindowsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIPIPLicenseWindowsOK creates a GetIPIPLicenseWindowsOK with default headers values
func NewGetIPIPLicenseWindowsOK() *GetIPIPLicenseWindowsOK {
	return &GetIPIPLicenseWindowsOK{}
}

/*GetIPIPLicenseWindowsOK handles this case with default header values.

return value
*/
type GetIPIPLicenseWindowsOK struct {
	Payload []string
}

func (o *GetIPIPLicenseWindowsOK) Error() string {
	return fmt.Sprintf("[GET /ip/{ip}/license/windows][%d] getIpIpLicenseWindowsOK  %+v", 200, o.Payload)
}

func (o *GetIPIPLicenseWindowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPIPLicenseWindowsDefault creates a GetIPIPLicenseWindowsDefault with default headers values
func NewGetIPIPLicenseWindowsDefault(code int) *GetIPIPLicenseWindowsDefault {
	return &GetIPIPLicenseWindowsDefault{
		_statusCode: code,
	}
}

/*GetIPIPLicenseWindowsDefault handles this case with default header values.

Unexpected error
*/
type GetIPIPLicenseWindowsDefault struct {
	_statusCode int

	Payload *models.GetIPIPLicenseWindowsDefaultBody
}

// Code gets the status code for the get IP IP license windows default response
func (o *GetIPIPLicenseWindowsDefault) Code() int {
	return o._statusCode
}

func (o *GetIPIPLicenseWindowsDefault) Error() string {
	return fmt.Sprintf("[GET /ip/{ip}/license/windows][%d] GetIPIPLicenseWindows default  %+v", o._statusCode, o.Payload)
}

func (o *GetIPIPLicenseWindowsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetIPIPLicenseWindowsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
