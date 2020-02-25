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

	"github.com/kobayashi/go-netbox/netbox/models"
)

// NewDcimVirtualChassisUpdateParams creates a new DcimVirtualChassisUpdateParams object
// with the default values initialized.
func NewDcimVirtualChassisUpdateParams() *DcimVirtualChassisUpdateParams {
	var ()
	return &DcimVirtualChassisUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimVirtualChassisUpdateParamsWithTimeout creates a new DcimVirtualChassisUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimVirtualChassisUpdateParamsWithTimeout(timeout time.Duration) *DcimVirtualChassisUpdateParams {
	var ()
	return &DcimVirtualChassisUpdateParams{

		timeout: timeout,
	}
}

// NewDcimVirtualChassisUpdateParamsWithContext creates a new DcimVirtualChassisUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimVirtualChassisUpdateParamsWithContext(ctx context.Context) *DcimVirtualChassisUpdateParams {
	var ()
	return &DcimVirtualChassisUpdateParams{

		Context: ctx,
	}
}

// NewDcimVirtualChassisUpdateParamsWithHTTPClient creates a new DcimVirtualChassisUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimVirtualChassisUpdateParamsWithHTTPClient(client *http.Client) *DcimVirtualChassisUpdateParams {
	var ()
	return &DcimVirtualChassisUpdateParams{
		HTTPClient: client,
	}
}

/*DcimVirtualChassisUpdateParams contains all the parameters to send to the API endpoint
for the dcim virtual chassis update operation typically these are written to a http.Request
*/
type DcimVirtualChassisUpdateParams struct {

	/*Data*/
	Data *models.WritableVirtualChassis
	/*ID
	  A unique integer value identifying this virtual chassis.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim virtual chassis update params
func (o *DcimVirtualChassisUpdateParams) WithTimeout(timeout time.Duration) *DcimVirtualChassisUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim virtual chassis update params
func (o *DcimVirtualChassisUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim virtual chassis update params
func (o *DcimVirtualChassisUpdateParams) WithContext(ctx context.Context) *DcimVirtualChassisUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim virtual chassis update params
func (o *DcimVirtualChassisUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim virtual chassis update params
func (o *DcimVirtualChassisUpdateParams) WithHTTPClient(client *http.Client) *DcimVirtualChassisUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim virtual chassis update params
func (o *DcimVirtualChassisUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim virtual chassis update params
func (o *DcimVirtualChassisUpdateParams) WithData(data *models.WritableVirtualChassis) *DcimVirtualChassisUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim virtual chassis update params
func (o *DcimVirtualChassisUpdateParams) SetData(data *models.WritableVirtualChassis) {
	o.Data = data
}

// WithID adds the id to the dcim virtual chassis update params
func (o *DcimVirtualChassisUpdateParams) WithID(id int64) *DcimVirtualChassisUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim virtual chassis update params
func (o *DcimVirtualChassisUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimVirtualChassisUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
