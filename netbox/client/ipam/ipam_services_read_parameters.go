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

package ipam

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

// NewIpamServicesReadParams creates a new IpamServicesReadParams object
// with the default values initialized.
func NewIpamServicesReadParams() *IpamServicesReadParams {
	var ()
	return &IpamServicesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIpamServicesReadParamsWithTimeout creates a new IpamServicesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIpamServicesReadParamsWithTimeout(timeout time.Duration) *IpamServicesReadParams {
	var ()
	return &IpamServicesReadParams{

		timeout: timeout,
	}
}

// NewIpamServicesReadParamsWithContext creates a new IpamServicesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewIpamServicesReadParamsWithContext(ctx context.Context) *IpamServicesReadParams {
	var ()
	return &IpamServicesReadParams{

		Context: ctx,
	}
}

// NewIpamServicesReadParamsWithHTTPClient creates a new IpamServicesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIpamServicesReadParamsWithHTTPClient(client *http.Client) *IpamServicesReadParams {
	var ()
	return &IpamServicesReadParams{
		HTTPClient: client,
	}
}

/*IpamServicesReadParams contains all the parameters to send to the API endpoint
for the ipam services read operation typically these are written to a http.Request
*/
type IpamServicesReadParams struct {

	/*ID
	  A unique integer value identifying this service.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam services read params
func (o *IpamServicesReadParams) WithTimeout(timeout time.Duration) *IpamServicesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam services read params
func (o *IpamServicesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam services read params
func (o *IpamServicesReadParams) WithContext(ctx context.Context) *IpamServicesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam services read params
func (o *IpamServicesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam services read params
func (o *IpamServicesReadParams) WithHTTPClient(client *http.Client) *IpamServicesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam services read params
func (o *IpamServicesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam services read params
func (o *IpamServicesReadParams) WithID(id int64) *IpamServicesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam services read params
func (o *IpamServicesReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamServicesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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