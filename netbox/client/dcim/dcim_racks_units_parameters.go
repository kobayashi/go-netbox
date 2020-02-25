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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDcimRacksUnitsParams creates a new DcimRacksUnitsParams object
// with the default values initialized.
func NewDcimRacksUnitsParams() *DcimRacksUnitsParams {
	var ()
	return &DcimRacksUnitsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRacksUnitsParamsWithTimeout creates a new DcimRacksUnitsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRacksUnitsParamsWithTimeout(timeout time.Duration) *DcimRacksUnitsParams {
	var ()
	return &DcimRacksUnitsParams{

		timeout: timeout,
	}
}

// NewDcimRacksUnitsParamsWithContext creates a new DcimRacksUnitsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRacksUnitsParamsWithContext(ctx context.Context) *DcimRacksUnitsParams {
	var ()
	return &DcimRacksUnitsParams{

		Context: ctx,
	}
}

// NewDcimRacksUnitsParamsWithHTTPClient creates a new DcimRacksUnitsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRacksUnitsParamsWithHTTPClient(client *http.Client) *DcimRacksUnitsParams {
	var ()
	return &DcimRacksUnitsParams{
		HTTPClient: client,
	}
}

/*DcimRacksUnitsParams contains all the parameters to send to the API endpoint
for the dcim racks units operation typically these are written to a http.Request
*/
type DcimRacksUnitsParams struct {

	/*ID
	  A unique integer value identifying this rack.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim racks units params
func (o *DcimRacksUnitsParams) WithTimeout(timeout time.Duration) *DcimRacksUnitsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim racks units params
func (o *DcimRacksUnitsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim racks units params
func (o *DcimRacksUnitsParams) WithContext(ctx context.Context) *DcimRacksUnitsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim racks units params
func (o *DcimRacksUnitsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim racks units params
func (o *DcimRacksUnitsParams) WithHTTPClient(client *http.Client) *DcimRacksUnitsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim racks units params
func (o *DcimRacksUnitsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim racks units params
func (o *DcimRacksUnitsParams) WithID(id int64) *DcimRacksUnitsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim racks units params
func (o *DcimRacksUnitsParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRacksUnitsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
