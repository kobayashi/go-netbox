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

// NewCircuitsCircuitTerminationsListParams creates a new CircuitsCircuitTerminationsListParams object
// with the default values initialized.
func NewCircuitsCircuitTerminationsListParams() *CircuitsCircuitTerminationsListParams {
	var ()
	return &CircuitsCircuitTerminationsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTerminationsListParamsWithTimeout creates a new CircuitsCircuitTerminationsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCircuitsCircuitTerminationsListParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTerminationsListParams {
	var ()
	return &CircuitsCircuitTerminationsListParams{

		timeout: timeout,
	}
}

// NewCircuitsCircuitTerminationsListParamsWithContext creates a new CircuitsCircuitTerminationsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewCircuitsCircuitTerminationsListParamsWithContext(ctx context.Context) *CircuitsCircuitTerminationsListParams {
	var ()
	return &CircuitsCircuitTerminationsListParams{

		Context: ctx,
	}
}

// NewCircuitsCircuitTerminationsListParamsWithHTTPClient creates a new CircuitsCircuitTerminationsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCircuitsCircuitTerminationsListParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTerminationsListParams {
	var ()
	return &CircuitsCircuitTerminationsListParams{
		HTTPClient: client,
	}
}

/*CircuitsCircuitTerminationsListParams contains all the parameters to send to the API endpoint
for the circuits circuit terminations list operation typically these are written to a http.Request
*/
type CircuitsCircuitTerminationsListParams struct {

	/*CircuitID*/
	CircuitID *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*PortSpeed*/
	PortSpeed *string
	/*Q*/
	Q *string
	/*Site*/
	Site *string
	/*SiteID*/
	SiteID *string
	/*TermSide*/
	TermSide *string
	/*UpstreamSpeed*/
	UpstreamSpeed *string
	/*XconnectID*/
	XconnectID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTerminationsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithContext(ctx context.Context) *CircuitsCircuitTerminationsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTerminationsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCircuitID adds the circuitID to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithCircuitID(circuitID *string) *CircuitsCircuitTerminationsListParams {
	o.SetCircuitID(circuitID)
	return o
}

// SetCircuitID adds the circuitId to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetCircuitID(circuitID *string) {
	o.CircuitID = circuitID
}

// WithLimit adds the limit to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithLimit(limit *int64) *CircuitsCircuitTerminationsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithOffset(offset *int64) *CircuitsCircuitTerminationsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPortSpeed adds the portSpeed to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithPortSpeed(portSpeed *string) *CircuitsCircuitTerminationsListParams {
	o.SetPortSpeed(portSpeed)
	return o
}

// SetPortSpeed adds the portSpeed to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetPortSpeed(portSpeed *string) {
	o.PortSpeed = portSpeed
}

// WithQ adds the q to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithQ(q *string) *CircuitsCircuitTerminationsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetQ(q *string) {
	o.Q = q
}

// WithSite adds the site to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithSite(site *string) *CircuitsCircuitTerminationsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithSiteID(siteID *string) *CircuitsCircuitTerminationsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithTermSide adds the termSide to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithTermSide(termSide *string) *CircuitsCircuitTerminationsListParams {
	o.SetTermSide(termSide)
	return o
}

// SetTermSide adds the termSide to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetTermSide(termSide *string) {
	o.TermSide = termSide
}

// WithUpstreamSpeed adds the upstreamSpeed to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithUpstreamSpeed(upstreamSpeed *string) *CircuitsCircuitTerminationsListParams {
	o.SetUpstreamSpeed(upstreamSpeed)
	return o
}

// SetUpstreamSpeed adds the upstreamSpeed to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetUpstreamSpeed(upstreamSpeed *string) {
	o.UpstreamSpeed = upstreamSpeed
}

// WithXconnectID adds the xconnectID to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) WithXconnectID(xconnectID *string) *CircuitsCircuitTerminationsListParams {
	o.SetXconnectID(xconnectID)
	return o
}

// SetXconnectID adds the xconnectId to the circuits circuit terminations list params
func (o *CircuitsCircuitTerminationsListParams) SetXconnectID(xconnectID *string) {
	o.XconnectID = xconnectID
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTerminationsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CircuitID != nil {

		// query param circuit_id
		var qrCircuitID string
		if o.CircuitID != nil {
			qrCircuitID = *o.CircuitID
		}
		qCircuitID := qrCircuitID
		if qCircuitID != "" {
			if err := r.SetQueryParam("circuit_id", qCircuitID); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.PortSpeed != nil {

		// query param port_speed
		var qrPortSpeed string
		if o.PortSpeed != nil {
			qrPortSpeed = *o.PortSpeed
		}
		qPortSpeed := qrPortSpeed
		if qPortSpeed != "" {
			if err := r.SetQueryParam("port_speed", qPortSpeed); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Site != nil {

		// query param site
		var qrSite string
		if o.Site != nil {
			qrSite = *o.Site
		}
		qSite := qrSite
		if qSite != "" {
			if err := r.SetQueryParam("site", qSite); err != nil {
				return err
			}
		}

	}

	if o.SiteID != nil {

		// query param site_id
		var qrSiteID string
		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := qrSiteID
		if qSiteID != "" {
			if err := r.SetQueryParam("site_id", qSiteID); err != nil {
				return err
			}
		}

	}

	if o.TermSide != nil {

		// query param term_side
		var qrTermSide string
		if o.TermSide != nil {
			qrTermSide = *o.TermSide
		}
		qTermSide := qrTermSide
		if qTermSide != "" {
			if err := r.SetQueryParam("term_side", qTermSide); err != nil {
				return err
			}
		}

	}

	if o.UpstreamSpeed != nil {

		// query param upstream_speed
		var qrUpstreamSpeed string
		if o.UpstreamSpeed != nil {
			qrUpstreamSpeed = *o.UpstreamSpeed
		}
		qUpstreamSpeed := qrUpstreamSpeed
		if qUpstreamSpeed != "" {
			if err := r.SetQueryParam("upstream_speed", qUpstreamSpeed); err != nil {
				return err
			}
		}

	}

	if o.XconnectID != nil {

		// query param xconnect_id
		var qrXconnectID string
		if o.XconnectID != nil {
			qrXconnectID = *o.XconnectID
		}
		qXconnectID := qrXconnectID
		if qXconnectID != "" {
			if err := r.SetQueryParam("xconnect_id", qXconnectID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
