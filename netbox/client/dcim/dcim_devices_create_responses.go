// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/kobayashi/go-netbox/netbox/models"
)

// DcimDevicesCreateReader is a Reader for the DcimDevicesCreate structure.
type DcimDevicesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimDevicesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimDevicesCreateCreated creates a DcimDevicesCreateCreated with default headers values
func NewDcimDevicesCreateCreated() *DcimDevicesCreateCreated {
	return &DcimDevicesCreateCreated{}
}

/*DcimDevicesCreateCreated handles this case with default header values.

DcimDevicesCreateCreated dcim devices create created
*/
type DcimDevicesCreateCreated struct {
	Payload *models.DeviceWithConfigContext
}

func (o *DcimDevicesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/devices/][%d] dcimDevicesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimDevicesCreateCreated) GetPayload() *models.DeviceWithConfigContext {
	return o.Payload
}

func (o *DcimDevicesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
