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

// DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDReader is a Reader for the DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetID structure.
type DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDOK creates a DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDOK with default headers values
func NewDeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDOK() *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDOK {
	return &DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDOK{}
}

/*DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDOK handles this case with default header values.

return 'void'
*/
type DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDOK struct {
}

func (o *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDOK) Error() string {
	return fmt.Sprintf("[DELETE /cloud/project/{serviceName}/network/private/{networkId}/subnet/{subnetId}][%d] deleteCloudProjectServiceNameNetworkPrivateNetworkIdSubnetSubnetIdOK ", 200)
}

func (o *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDDefault creates a DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDDefault with default headers values
func NewDeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDDefault(code int) *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDDefault {
	return &DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDDefault{
		_statusCode: code,
	}
}

/*DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDDefault handles this case with default header values.

Unexpected error
*/
type DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDDefault struct {
	_statusCode int

	Payload *models.DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDDefaultBody
}

// Code gets the status code for the delete cloud project service name network private network ID subnet subnet ID default response
func (o *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /cloud/project/{serviceName}/network/private/{networkId}/subnet/{subnetId}][%d] DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteCloudProjectServiceNameNetworkPrivateNetworkIDSubnetSubnetIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
