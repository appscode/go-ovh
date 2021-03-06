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

	"github.com/appscode/go-ovh/cloud/models"
)

// GetCloudProjectServiceNameForecastReader is a Reader for the GetCloudProjectServiceNameForecast structure.
type GetCloudProjectServiceNameForecastReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudProjectServiceNameForecastReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCloudProjectServiceNameForecastOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetCloudProjectServiceNameForecastDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCloudProjectServiceNameForecastOK creates a GetCloudProjectServiceNameForecastOK with default headers values
func NewGetCloudProjectServiceNameForecastOK() *GetCloudProjectServiceNameForecastOK {
	return &GetCloudProjectServiceNameForecastOK{}
}

/*GetCloudProjectServiceNameForecastOK handles this case with default header values.

description of 'cloud.Forecast.ProjectForecast' response
*/
type GetCloudProjectServiceNameForecastOK struct {
	Payload *models.CloudForecastProjectForecast
}

func (o *GetCloudProjectServiceNameForecastOK) Error() string {
	return fmt.Sprintf("[GET /cloud/project/{serviceName}/forecast][%d] getCloudProjectServiceNameForecastOK  %+v", 200, o.Payload)
}

func (o *GetCloudProjectServiceNameForecastOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudForecastProjectForecast)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudProjectServiceNameForecastDefault creates a GetCloudProjectServiceNameForecastDefault with default headers values
func NewGetCloudProjectServiceNameForecastDefault(code int) *GetCloudProjectServiceNameForecastDefault {
	return &GetCloudProjectServiceNameForecastDefault{
		_statusCode: code,
	}
}

/*GetCloudProjectServiceNameForecastDefault handles this case with default header values.

Unexpected error
*/
type GetCloudProjectServiceNameForecastDefault struct {
	_statusCode int

	Payload *models.GetCloudProjectServiceNameForecastDefaultBody
}

// Code gets the status code for the get cloud project service name forecast default response
func (o *GetCloudProjectServiceNameForecastDefault) Code() int {
	return o._statusCode
}

func (o *GetCloudProjectServiceNameForecastDefault) Error() string {
	return fmt.Sprintf("[GET /cloud/project/{serviceName}/forecast][%d] GetCloudProjectServiceNameForecast default  %+v", o._statusCode, o.Payload)
}

func (o *GetCloudProjectServiceNameForecastDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCloudProjectServiceNameForecastDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
