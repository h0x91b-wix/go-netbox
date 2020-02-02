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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/h0x91b-wix/go-netbox/netbox/models"
)

// NewDcimManufacturersCreateParams creates a new DcimManufacturersCreateParams object
// with the default values initialized.
func NewDcimManufacturersCreateParams() *DcimManufacturersCreateParams {
	var ()
	return &DcimManufacturersCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimManufacturersCreateParamsWithTimeout creates a new DcimManufacturersCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimManufacturersCreateParamsWithTimeout(timeout time.Duration) *DcimManufacturersCreateParams {
	var ()
	return &DcimManufacturersCreateParams{

		timeout: timeout,
	}
}

// NewDcimManufacturersCreateParamsWithContext creates a new DcimManufacturersCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimManufacturersCreateParamsWithContext(ctx context.Context) *DcimManufacturersCreateParams {
	var ()
	return &DcimManufacturersCreateParams{

		Context: ctx,
	}
}

// NewDcimManufacturersCreateParamsWithHTTPClient creates a new DcimManufacturersCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimManufacturersCreateParamsWithHTTPClient(client *http.Client) *DcimManufacturersCreateParams {
	var ()
	return &DcimManufacturersCreateParams{
		HTTPClient: client,
	}
}

/*DcimManufacturersCreateParams contains all the parameters to send to the API endpoint
for the dcim manufacturers create operation typically these are written to a http.Request
*/
type DcimManufacturersCreateParams struct {

	/*Data*/
	Data *models.Manufacturer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim manufacturers create params
func (o *DcimManufacturersCreateParams) WithTimeout(timeout time.Duration) *DcimManufacturersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim manufacturers create params
func (o *DcimManufacturersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim manufacturers create params
func (o *DcimManufacturersCreateParams) WithContext(ctx context.Context) *DcimManufacturersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim manufacturers create params
func (o *DcimManufacturersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim manufacturers create params
func (o *DcimManufacturersCreateParams) WithHTTPClient(client *http.Client) *DcimManufacturersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim manufacturers create params
func (o *DcimManufacturersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim manufacturers create params
func (o *DcimManufacturersCreateParams) WithData(data *models.Manufacturer) *DcimManufacturersCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim manufacturers create params
func (o *DcimManufacturersCreateParams) SetData(data *models.Manufacturer) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimManufacturersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
