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

// NewDcimConsoleServerPortsListParams creates a new DcimConsoleServerPortsListParams object
// with the default values initialized.
func NewDcimConsoleServerPortsListParams() *DcimConsoleServerPortsListParams {
	var ()
	return &DcimConsoleServerPortsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsoleServerPortsListParamsWithTimeout creates a new DcimConsoleServerPortsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimConsoleServerPortsListParamsWithTimeout(timeout time.Duration) *DcimConsoleServerPortsListParams {
	var ()
	return &DcimConsoleServerPortsListParams{

		timeout: timeout,
	}
}

// NewDcimConsoleServerPortsListParamsWithContext creates a new DcimConsoleServerPortsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimConsoleServerPortsListParamsWithContext(ctx context.Context) *DcimConsoleServerPortsListParams {
	var ()
	return &DcimConsoleServerPortsListParams{

		Context: ctx,
	}
}

// NewDcimConsoleServerPortsListParamsWithHTTPClient creates a new DcimConsoleServerPortsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimConsoleServerPortsListParamsWithHTTPClient(client *http.Client) *DcimConsoleServerPortsListParams {
	var ()
	return &DcimConsoleServerPortsListParams{
		HTTPClient: client,
	}
}

/*DcimConsoleServerPortsListParams contains all the parameters to send to the API endpoint
for the dcim console server ports list operation typically these are written to a http.Request
*/
type DcimConsoleServerPortsListParams struct {

	/*Cabled*/
	Cabled *string
	/*ConnectionStatus*/
	ConnectionStatus *string
	/*Description*/
	Description *string
	/*Device*/
	Device *string
	/*DeviceID*/
	DeviceID *string
	/*ID*/
	ID *string
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
	/*Type*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithTimeout(timeout time.Duration) *DcimConsoleServerPortsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithContext(ctx context.Context) *DcimConsoleServerPortsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithHTTPClient(client *http.Client) *DcimConsoleServerPortsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCabled adds the cabled to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithCabled(cabled *string) *DcimConsoleServerPortsListParams {
	o.SetCabled(cabled)
	return o
}

// SetCabled adds the cabled to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetCabled(cabled *string) {
	o.Cabled = cabled
}

// WithConnectionStatus adds the connectionStatus to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithConnectionStatus(connectionStatus *string) *DcimConsoleServerPortsListParams {
	o.SetConnectionStatus(connectionStatus)
	return o
}

// SetConnectionStatus adds the connectionStatus to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetConnectionStatus(connectionStatus *string) {
	o.ConnectionStatus = connectionStatus
}

// WithDescription adds the description to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithDescription(description *string) *DcimConsoleServerPortsListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDevice adds the device to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithDevice(device *string) *DcimConsoleServerPortsListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithDeviceID(deviceID *string) *DcimConsoleServerPortsListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithID adds the id to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithID(id *string) *DcimConsoleServerPortsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetID(id *string) {
	o.ID = id
}

// WithLimit adds the limit to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithLimit(limit *int64) *DcimConsoleServerPortsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithName(name *string) *DcimConsoleServerPortsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithOffset(offset *int64) *DcimConsoleServerPortsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithQ(q *string) *DcimConsoleServerPortsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithRegion(region *string) *DcimConsoleServerPortsListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionID adds the regionID to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithRegionID(regionID *string) *DcimConsoleServerPortsListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithSite adds the site to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithSite(site *string) *DcimConsoleServerPortsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithSiteID(siteID *string) *DcimConsoleServerPortsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithTag adds the tag to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithTag(tag *string) *DcimConsoleServerPortsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithType adds the typeVar to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithType(typeVar *string) *DcimConsoleServerPortsListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsoleServerPortsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cabled != nil {

		// query param cabled
		var qrCabled string
		if o.Cabled != nil {
			qrCabled = *o.Cabled
		}
		qCabled := qrCabled
		if qCabled != "" {
			if err := r.SetQueryParam("cabled", qCabled); err != nil {
				return err
			}
		}

	}

	if o.ConnectionStatus != nil {

		// query param connection_status
		var qrConnectionStatus string
		if o.ConnectionStatus != nil {
			qrConnectionStatus = *o.ConnectionStatus
		}
		qConnectionStatus := qrConnectionStatus
		if qConnectionStatus != "" {
			if err := r.SetQueryParam("connection_status", qConnectionStatus); err != nil {
				return err
			}
		}

	}

	if o.Description != nil {

		// query param description
		var qrDescription string
		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {
			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}

	}

	if o.Device != nil {

		// query param device
		var qrDevice string
		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {
			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}

	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID string
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := qrDeviceID
		if qDeviceID != "" {
			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
				return err
			}
		}

	}

	if o.ID != nil {

		// query param id
		var qrID string
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
