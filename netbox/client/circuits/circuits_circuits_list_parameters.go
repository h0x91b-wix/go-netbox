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

// NewCircuitsCircuitsListParams creates a new CircuitsCircuitsListParams object
// with the default values initialized.
func NewCircuitsCircuitsListParams() *CircuitsCircuitsListParams {
	var ()
	return &CircuitsCircuitsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitsListParamsWithTimeout creates a new CircuitsCircuitsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCircuitsCircuitsListParamsWithTimeout(timeout time.Duration) *CircuitsCircuitsListParams {
	var ()
	return &CircuitsCircuitsListParams{

		timeout: timeout,
	}
}

// NewCircuitsCircuitsListParamsWithContext creates a new CircuitsCircuitsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewCircuitsCircuitsListParamsWithContext(ctx context.Context) *CircuitsCircuitsListParams {
	var ()
	return &CircuitsCircuitsListParams{

		Context: ctx,
	}
}

// NewCircuitsCircuitsListParamsWithHTTPClient creates a new CircuitsCircuitsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCircuitsCircuitsListParamsWithHTTPClient(client *http.Client) *CircuitsCircuitsListParams {
	var ()
	return &CircuitsCircuitsListParams{
		HTTPClient: client,
	}
}

/*CircuitsCircuitsListParams contains all the parameters to send to the API endpoint
for the circuits circuits list operation typically these are written to a http.Request
*/
type CircuitsCircuitsListParams struct {

	/*Cid*/
	Cid *string
	/*CommitRate*/
	CommitRate *string
	/*Created*/
	Created *string
	/*CreatedGte*/
	CreatedGte *string
	/*CreatedLte*/
	CreatedLte *string
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *string
	/*InstallDate*/
	InstallDate *string
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
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Provider*/
	Provider *string
	/*ProviderID*/
	ProviderID *string
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
	/*Status*/
	Status *string
	/*Tag*/
	Tag *string
	/*Tenant*/
	Tenant *string
	/*TenantGroup*/
	TenantGroup *string
	/*TenantGroupID*/
	TenantGroupID *string
	/*TenantID*/
	TenantID *string
	/*Type*/
	Type *string
	/*TypeID*/
	TypeID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTimeout(timeout time.Duration) *CircuitsCircuitsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithContext(ctx context.Context) *CircuitsCircuitsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithHTTPClient(client *http.Client) *CircuitsCircuitsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCid adds the cid to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCid(cid *string) *CircuitsCircuitsListParams {
	o.SetCid(cid)
	return o
}

// SetCid adds the cid to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCid(cid *string) {
	o.Cid = cid
}

// WithCommitRate adds the commitRate to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCommitRate(commitRate *string) *CircuitsCircuitsListParams {
	o.SetCommitRate(commitRate)
	return o
}

// SetCommitRate adds the commitRate to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCommitRate(commitRate *string) {
	o.CommitRate = commitRate
}

// WithCreated adds the created to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCreated(created *string) *CircuitsCircuitsListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCreatedGte(createdGte *string) *CircuitsCircuitsListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCreatedLte(createdLte *string) *CircuitsCircuitsListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithIDIn adds the iDIn to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithIDIn(iDIn *string) *CircuitsCircuitsListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithInstallDate adds the installDate to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithInstallDate(installDate *string) *CircuitsCircuitsListParams {
	o.SetInstallDate(installDate)
	return o
}

// SetInstallDate adds the installDate to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetInstallDate(installDate *string) {
	o.InstallDate = installDate
}

// WithLastUpdated adds the lastUpdated to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithLastUpdated(lastUpdated *string) *CircuitsCircuitsListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithLastUpdatedGte(lastUpdatedGte *string) *CircuitsCircuitsListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithLastUpdatedLte(lastUpdatedLte *string) *CircuitsCircuitsListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithLimit(limit *int64) *CircuitsCircuitsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithOffset(offset *int64) *CircuitsCircuitsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithProvider adds the provider to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithProvider(provider *string) *CircuitsCircuitsListParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetProvider(provider *string) {
	o.Provider = provider
}

// WithProviderID adds the providerID to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithProviderID(providerID *string) *CircuitsCircuitsListParams {
	o.SetProviderID(providerID)
	return o
}

// SetProviderID adds the providerId to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetProviderID(providerID *string) {
	o.ProviderID = providerID
}

// WithQ adds the q to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithQ(q *string) *CircuitsCircuitsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithRegion(region *string) *CircuitsCircuitsListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionID adds the regionID to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithRegionID(regionID *string) *CircuitsCircuitsListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithSite adds the site to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithSite(site *string) *CircuitsCircuitsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithSiteID(siteID *string) *CircuitsCircuitsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithStatus adds the status to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithStatus(status *string) *CircuitsCircuitsListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetStatus(status *string) {
	o.Status = status
}

// WithTag adds the tag to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTag(tag *string) *CircuitsCircuitsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTenant adds the tenant to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTenant(tenant *string) *CircuitsCircuitsListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantGroup adds the tenantGroup to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTenantGroup(tenantGroup *string) *CircuitsCircuitsListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupID adds the tenantGroupID to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTenantGroupID(tenantGroupID *string) *CircuitsCircuitsListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantID adds the tenantID to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTenantID(tenantID *string) *CircuitsCircuitsListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithType adds the typeVar to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithType(typeVar *string) *CircuitsCircuitsListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithTypeID adds the typeID to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTypeID(typeID *string) *CircuitsCircuitsListParams {
	o.SetTypeID(typeID)
	return o
}

// SetTypeID adds the typeId to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTypeID(typeID *string) {
	o.TypeID = typeID
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cid != nil {

		// query param cid
		var qrCid string
		if o.Cid != nil {
			qrCid = *o.Cid
		}
		qCid := qrCid
		if qCid != "" {
			if err := r.SetQueryParam("cid", qCid); err != nil {
				return err
			}
		}

	}

	if o.CommitRate != nil {

		// query param commit_rate
		var qrCommitRate string
		if o.CommitRate != nil {
			qrCommitRate = *o.CommitRate
		}
		qCommitRate := qrCommitRate
		if qCommitRate != "" {
			if err := r.SetQueryParam("commit_rate", qCommitRate); err != nil {
				return err
			}
		}

	}

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

	if o.InstallDate != nil {

		// query param install_date
		var qrInstallDate string
		if o.InstallDate != nil {
			qrInstallDate = *o.InstallDate
		}
		qInstallDate := qrInstallDate
		if qInstallDate != "" {
			if err := r.SetQueryParam("install_date", qInstallDate); err != nil {
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

	if o.Provider != nil {

		// query param provider
		var qrProvider string
		if o.Provider != nil {
			qrProvider = *o.Provider
		}
		qProvider := qrProvider
		if qProvider != "" {
			if err := r.SetQueryParam("provider", qProvider); err != nil {
				return err
			}
		}

	}

	if o.ProviderID != nil {

		// query param provider_id
		var qrProviderID string
		if o.ProviderID != nil {
			qrProviderID = *o.ProviderID
		}
		qProviderID := qrProviderID
		if qProviderID != "" {
			if err := r.SetQueryParam("provider_id", qProviderID); err != nil {
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

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
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

	if o.TenantGroup != nil {

		// query param tenant_group
		var qrTenantGroup string
		if o.TenantGroup != nil {
			qrTenantGroup = *o.TenantGroup
		}
		qTenantGroup := qrTenantGroup
		if qTenantGroup != "" {
			if err := r.SetQueryParam("tenant_group", qTenantGroup); err != nil {
				return err
			}
		}

	}

	if o.TenantGroupID != nil {

		// query param tenant_group_id
		var qrTenantGroupID string
		if o.TenantGroupID != nil {
			qrTenantGroupID = *o.TenantGroupID
		}
		qTenantGroupID := qrTenantGroupID
		if qTenantGroupID != "" {
			if err := r.SetQueryParam("tenant_group_id", qTenantGroupID); err != nil {
				return err
			}
		}

	}

	if o.TenantID != nil {

		// query param tenant_id
		var qrTenantID string
		if o.TenantID != nil {
			qrTenantID = *o.TenantID
		}
		qTenantID := qrTenantID
		if qTenantID != "" {
			if err := r.SetQueryParam("tenant_id", qTenantID); err != nil {
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
