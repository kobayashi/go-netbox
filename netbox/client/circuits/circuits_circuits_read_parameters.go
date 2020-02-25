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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCircuitsCircuitsReadParams creates a new CircuitsCircuitsReadParams object
// with the default values initialized.
func NewCircuitsCircuitsReadParams() *CircuitsCircuitsReadParams {
	var ()
	return &CircuitsCircuitsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitsReadParamsWithTimeout creates a new CircuitsCircuitsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCircuitsCircuitsReadParamsWithTimeout(timeout time.Duration) *CircuitsCircuitsReadParams {
	var ()
	return &CircuitsCircuitsReadParams{

		timeout: timeout,
	}
}

// NewCircuitsCircuitsReadParamsWithContext creates a new CircuitsCircuitsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewCircuitsCircuitsReadParamsWithContext(ctx context.Context) *CircuitsCircuitsReadParams {
	var ()
	return &CircuitsCircuitsReadParams{

		Context: ctx,
	}
}

// NewCircuitsCircuitsReadParamsWithHTTPClient creates a new CircuitsCircuitsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCircuitsCircuitsReadParamsWithHTTPClient(client *http.Client) *CircuitsCircuitsReadParams {
	var ()
	return &CircuitsCircuitsReadParams{
		HTTPClient: client,
	}
}

/*CircuitsCircuitsReadParams contains all the parameters to send to the API endpoint
for the circuits circuits read operation typically these are written to a http.Request
*/
type CircuitsCircuitsReadParams struct {

	/*ID
	  A unique integer value identifying this circuit.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the circuits circuits read params
func (o *CircuitsCircuitsReadParams) WithTimeout(timeout time.Duration) *CircuitsCircuitsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuits read params
func (o *CircuitsCircuitsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuits read params
func (o *CircuitsCircuitsReadParams) WithContext(ctx context.Context) *CircuitsCircuitsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuits read params
func (o *CircuitsCircuitsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuits read params
func (o *CircuitsCircuitsReadParams) WithHTTPClient(client *http.Client) *CircuitsCircuitsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuits read params
func (o *CircuitsCircuitsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the circuits circuits read params
func (o *CircuitsCircuitsReadParams) WithID(id int64) *CircuitsCircuitsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits circuits read params
func (o *CircuitsCircuitsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
