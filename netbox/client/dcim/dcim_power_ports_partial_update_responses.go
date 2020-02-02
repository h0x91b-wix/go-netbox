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

	models "github.com/h0x91b-wix/go-netbox/netbox/models"
)

// DcimPowerPortsPartialUpdateReader is a Reader for the DcimPowerPortsPartialUpdate structure.
type DcimPowerPortsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPowerPortsPartialUpdateOK creates a DcimPowerPortsPartialUpdateOK with default headers values
func NewDcimPowerPortsPartialUpdateOK() *DcimPowerPortsPartialUpdateOK {
	return &DcimPowerPortsPartialUpdateOK{}
}

/*DcimPowerPortsPartialUpdateOK handles this case with default header values.

DcimPowerPortsPartialUpdateOK dcim power ports partial update o k
*/
type DcimPowerPortsPartialUpdateOK struct {
	Payload *models.PowerPort
}

func (o *DcimPowerPortsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-ports/{id}/][%d] dcimPowerPortsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortsPartialUpdateOK) GetPayload() *models.PowerPort {
	return o.Payload
}

func (o *DcimPowerPortsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
