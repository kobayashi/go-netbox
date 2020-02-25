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

package virtualization

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

// NewVirtualizationClustersListParams creates a new VirtualizationClustersListParams object
// with the default values initialized.
func NewVirtualizationClustersListParams() *VirtualizationClustersListParams {
	var ()
	return &VirtualizationClustersListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClustersListParamsWithTimeout creates a new VirtualizationClustersListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationClustersListParamsWithTimeout(timeout time.Duration) *VirtualizationClustersListParams {
	var ()
	return &VirtualizationClustersListParams{

		timeout: timeout,
	}
}

// NewVirtualizationClustersListParamsWithContext creates a new VirtualizationClustersListParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationClustersListParamsWithContext(ctx context.Context) *VirtualizationClustersListParams {
	var ()
	return &VirtualizationClustersListParams{

		Context: ctx,
	}
}

// NewVirtualizationClustersListParamsWithHTTPClient creates a new VirtualizationClustersListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationClustersListParamsWithHTTPClient(client *http.Client) *VirtualizationClustersListParams {
	var ()
	return &VirtualizationClustersListParams{
		HTTPClient: client,
	}
}

/*VirtualizationClustersListParams contains all the parameters to send to the API endpoint
for the virtualization clusters list operation typically these are written to a http.Request
*/
type VirtualizationClustersListParams struct {

	/*Created*/
	Created *string
	/*CreatedGte*/
	CreatedGte *string
	/*CreatedLte*/
	CreatedLte *string
	/*Group*/
	Group *string
	/*GroupID*/
	GroupID *string
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *string
	/*LastUpdated*/
	LastUpdated *string
	/*LastUpdatedGte*/
	LastUpdatedGte *string
	/*LastUpdatedLte*/
	LastUpdatedLte *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*Region*/
	Region *string
	/*RegionID*/
	RegionID *string
	/*Site*/
	Site *string
	/*SiteID*/
	SiteID *string
	/*Tag*/
	Tag *string
	/*Tenant*/
	Tenant *string
	/*Type*/
	Type *string
	/*TypeID*/
	TypeID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithTimeout(timeout time.Duration) *VirtualizationClustersListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithContext(ctx context.Context) *VirtualizationClustersListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithHTTPClient(client *http.Client) *VirtualizationClustersListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithCreated(created *string) *VirtualizationClustersListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithCreatedGte(createdGte *string) *VirtualizationClustersListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithCreatedLte(createdLte *string) *VirtualizationClustersListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithGroup adds the group to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithGroup(group *string) *VirtualizationClustersListParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetGroup(group *string) {
	o.Group = group
}

// WithGroupID adds the groupID to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithGroupID(groupID *string) *VirtualizationClustersListParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetGroupID(groupID *string) {
	o.GroupID = groupID
}

// WithIDIn adds the iDIn to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithIDIn(iDIn *string) *VirtualizationClustersListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithLastUpdated adds the lastUpdated to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithLastUpdated(lastUpdated *string) *VirtualizationClustersListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithLastUpdatedGte(lastUpdatedGte *string) *VirtualizationClustersListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithLastUpdatedLte(lastUpdatedLte *string) *VirtualizationClustersListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithLimit(limit *int64) *VirtualizationClustersListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithName(name *string) *VirtualizationClustersListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithOffset(offset *int64) *VirtualizationClustersListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithQ(q *string) *VirtualizationClustersListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithRegion(region *string) *VirtualizationClustersListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionID adds the regionID to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithRegionID(regionID *string) *VirtualizationClustersListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithSite adds the site to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithSite(site *string) *VirtualizationClustersListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithSiteID(siteID *string) *VirtualizationClustersListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithTag adds the tag to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithTag(tag *string) *VirtualizationClustersListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTenant adds the tenant to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithTenant(tenant *string) *VirtualizationClustersListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithType adds the typeVar to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithType(typeVar *string) *VirtualizationClustersListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithTypeID adds the typeID to the virtualization clusters list params
func (o *VirtualizationClustersListParams) WithTypeID(typeID *string) *VirtualizationClustersListParams {
	o.SetTypeID(typeID)
	return o
}

// SetTypeID adds the typeId to the virtualization clusters list params
func (o *VirtualizationClustersListParams) SetTypeID(typeID *string) {
	o.TypeID = typeID
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClustersListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Created != nil {

		// query param created
		var qrCreated string
		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {
			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}

	}

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string
		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {
			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}

	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string
		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {
			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
				return err
			}
		}

	}

	if o.Group != nil {

		// query param group
		var qrGroup string
		if o.Group != nil {
			qrGroup = *o.Group
		}
		qGroup := qrGroup
		if qGroup != "" {
			if err := r.SetQueryParam("group", qGroup); err != nil {
				return err
			}
		}

	}

	if o.GroupID != nil {

		// query param group_id
		var qrGroupID string
		if o.GroupID != nil {
			qrGroupID = *o.GroupID
		}
		qGroupID := qrGroupID
		if qGroupID != "" {
			if err := r.SetQueryParam("group_id", qGroupID); err != nil {
				return err
			}
		}

	}

	if o.IDIn != nil {

		// query param id__in
		var qrIDIn string
		if o.IDIn != nil {
			qrIDIn = *o.IDIn
		}
		qIDIn := qrIDIn
		if qIDIn != "" {
			if err := r.SetQueryParam("id__in", qIDIn); err != nil {
				return err
			}
		}

	}

	if o.LastUpdated != nil {

		// query param last_updated
		var qrLastUpdated string
		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := qrLastUpdated
		if qLastUpdated != "" {
			if err := r.SetQueryParam("last_updated", qLastUpdated); err != nil {
				return err
			}
		}

	}

	if o.LastUpdatedGte != nil {

		// query param last_updated__gte
		var qrLastUpdatedGte string
		if o.LastUpdatedGte != nil {
			qrLastUpdatedGte = *o.LastUpdatedGte
		}
		qLastUpdatedGte := qrLastUpdatedGte
		if qLastUpdatedGte != "" {
			if err := r.SetQueryParam("last_updated__gte", qLastUpdatedGte); err != nil {
				return err
			}
		}

	}

	if o.LastUpdatedLte != nil {

		// query param last_updated__lte
		var qrLastUpdatedLte string
		if o.LastUpdatedLte != nil {
			qrLastUpdatedLte = *o.LastUpdatedLte
		}
		qLastUpdatedLte := qrLastUpdatedLte
		if qLastUpdatedLte != "" {
			if err := r.SetQueryParam("last_updated__lte", qLastUpdatedLte); err != nil {
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

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
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

	if o.Region != nil {

		// query param region
		var qrRegion string
		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {
			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
		}

	}

	if o.RegionID != nil {

		// query param region_id
		var qrRegionID string
		if o.RegionID != nil {
			qrRegionID = *o.RegionID
		}
		qRegionID := qrRegionID
		if qRegionID != "" {
			if err := r.SetQueryParam("region_id", qRegionID); err != nil {
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

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if o.Tenant != nil {

		// query param tenant
		var qrTenant string
		if o.Tenant != nil {
			qrTenant = *o.Tenant
		}
		qTenant := qrTenant
		if qTenant != "" {
			if err := r.SetQueryParam("tenant", qTenant); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if o.TypeID != nil {

		// query param type_id
		var qrTypeID string
		if o.TypeID != nil {
			qrTypeID = *o.TypeID
		}
		qTypeID := qrTypeID
		if qTypeID != "" {
			if err := r.SetQueryParam("type_id", qTypeID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
