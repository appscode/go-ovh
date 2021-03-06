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

	"github.com/appscode/go-ovh/domain/models"
)

// PutDomainZoneZoneNameDynHostRecordIDReader is a Reader for the PutDomainZoneZoneNameDynHostRecordID structure.
type PutDomainZoneZoneNameDynHostRecordIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDomainZoneZoneNameDynHostRecordIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutDomainZoneZoneNameDynHostRecordIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutDomainZoneZoneNameDynHostRecordIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutDomainZoneZoneNameDynHostRecordIDOK creates a PutDomainZoneZoneNameDynHostRecordIDOK with default headers values
func NewPutDomainZoneZoneNameDynHostRecordIDOK() *PutDomainZoneZoneNameDynHostRecordIDOK {
	return &PutDomainZoneZoneNameDynHostRecordIDOK{}
}

/*PutDomainZoneZoneNameDynHostRecordIDOK handles this case with default header values.

return 'void'
*/
type PutDomainZoneZoneNameDynHostRecordIDOK struct {
}

func (o *PutDomainZoneZoneNameDynHostRecordIDOK) Error() string {
	return fmt.Sprintf("[PUT /domain/zone/{zoneName}/dynHost/record/{id}][%d] putDomainZoneZoneNameDynHostRecordIdOK ", 200)
}

func (o *PutDomainZoneZoneNameDynHostRecordIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutDomainZoneZoneNameDynHostRecordIDDefault creates a PutDomainZoneZoneNameDynHostRecordIDDefault with default headers values
func NewPutDomainZoneZoneNameDynHostRecordIDDefault(code int) *PutDomainZoneZoneNameDynHostRecordIDDefault {
	return &PutDomainZoneZoneNameDynHostRecordIDDefault{
		_statusCode: code,
	}
}

/*PutDomainZoneZoneNameDynHostRecordIDDefault handles this case with default header values.

Unexpected error
*/
type PutDomainZoneZoneNameDynHostRecordIDDefault struct {
	_statusCode int

	Payload *models.PutDomainZoneZoneNameDynHostRecordIDDefaultBody
}

// Code gets the status code for the put domain zone zone name dyn host record ID default response
func (o *PutDomainZoneZoneNameDynHostRecordIDDefault) Code() int {
	return o._statusCode
}

func (o *PutDomainZoneZoneNameDynHostRecordIDDefault) Error() string {
	return fmt.Sprintf("[PUT /domain/zone/{zoneName}/dynHost/record/{id}][%d] PutDomainZoneZoneNameDynHostRecordID default  %+v", o._statusCode, o.Payload)
}

func (o *PutDomainZoneZoneNameDynHostRecordIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutDomainZoneZoneNameDynHostRecordIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
