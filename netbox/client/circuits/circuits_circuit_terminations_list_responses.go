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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/kobayashi/go-netbox/netbox/models"
)

// CircuitsCircuitTerminationsListReader is a Reader for the CircuitsCircuitTerminationsList structure.
type CircuitsCircuitTerminationsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTerminationsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCircuitsCircuitTerminationsListOK creates a CircuitsCircuitTerminationsListOK with default headers values
func NewCircuitsCircuitTerminationsListOK() *CircuitsCircuitTerminationsListOK {
	return &CircuitsCircuitTerminationsListOK{}
}

/*CircuitsCircuitTerminationsListOK handles this case with default header values.

CircuitsCircuitTerminationsListOK circuits circuit terminations list o k
*/
type CircuitsCircuitTerminationsListOK struct {
	Payload *CircuitsCircuitTerminationsListOKBody
}

func (o *CircuitsCircuitTerminationsListOK) Error() string {
	return fmt.Sprintf("[GET /circuits/circuit-terminations/][%d] circuitsCircuitTerminationsListOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTerminationsListOK) GetPayload() *CircuitsCircuitTerminationsListOKBody {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CircuitsCircuitTerminationsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CircuitsCircuitTerminationsListOKBody circuits circuit terminations list o k body
swagger:model CircuitsCircuitTerminationsListOKBody
*/
type CircuitsCircuitTerminationsListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.CircuitTermination `json:"results"`
}

// Validate validates this circuits circuit terminations list o k body
func (o *CircuitsCircuitTerminationsListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CircuitsCircuitTerminationsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("circuitsCircuitTerminationsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *CircuitsCircuitTerminationsListOKBody) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("circuitsCircuitTerminationsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *CircuitsCircuitTerminationsListOKBody) validatePrevious(formats strfmt.Registry) error {

	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("circuitsCircuitTerminationsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *CircuitsCircuitTerminationsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("circuitsCircuitTerminationsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("circuitsCircuitTerminationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CircuitsCircuitTerminationsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CircuitsCircuitTerminationsListOKBody) UnmarshalBinary(b []byte) error {
	var res CircuitsCircuitTerminationsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
