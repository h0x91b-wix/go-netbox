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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/netbox-community/go-netbox/models"
)

// NewVirtualizationClusterTypesCreateParams creates a new VirtualizationClusterTypesCreateParams object
// with the default values initialized.
func NewVirtualizationClusterTypesCreateParams() *VirtualizationClusterTypesCreateParams {
	var ()
	return &VirtualizationClusterTypesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClusterTypesCreateParamsWithTimeout creates a new VirtualizationClusterTypesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationClusterTypesCreateParamsWithTimeout(timeout time.Duration) *VirtualizationClusterTypesCreateParams {
	var ()
	return &VirtualizationClusterTypesCreateParams{

		timeout: timeout,
	}
}

// NewVirtualizationClusterTypesCreateParamsWithContext creates a new VirtualizationClusterTypesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationClusterTypesCreateParamsWithContext(ctx context.Context) *VirtualizationClusterTypesCreateParams {
	var ()
	return &VirtualizationClusterTypesCreateParams{

		Context: ctx,
	}
}

// NewVirtualizationClusterTypesCreateParamsWithHTTPClient creates a new VirtualizationClusterTypesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationClusterTypesCreateParamsWithHTTPClient(client *http.Client) *VirtualizationClusterTypesCreateParams {
	var ()
	return &VirtualizationClusterTypesCreateParams{
		HTTPClient: client,
	}
}

/*VirtualizationClusterTypesCreateParams contains all the parameters to send to the API endpoint
for the virtualization cluster types create operation typically these are written to a http.Request
*/
type VirtualizationClusterTypesCreateParams struct {

	/*Data*/
	Data *models.ClusterType

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization cluster types create params
func (o *VirtualizationClusterTypesCreateParams) WithTimeout(timeout time.Duration) *VirtualizationClusterTypesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization cluster types create params
func (o *VirtualizationClusterTypesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization cluster types create params
func (o *VirtualizationClusterTypesCreateParams) WithContext(ctx context.Context) *VirtualizationClusterTypesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization cluster types create params
func (o *VirtualizationClusterTypesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization cluster types create params
func (o *VirtualizationClusterTypesCreateParams) WithHTTPClient(client *http.Client) *VirtualizationClusterTypesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization cluster types create params
func (o *VirtualizationClusterTypesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization cluster types create params
func (o *VirtualizationClusterTypesCreateParams) WithData(data *models.ClusterType) *VirtualizationClusterTypesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization cluster types create params
func (o *VirtualizationClusterTypesCreateParams) SetData(data *models.ClusterType) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClusterTypesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}