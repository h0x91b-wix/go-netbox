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

// NewDcimFrontPortsCreateParams creates a new DcimFrontPortsCreateParams object
// with the default values initialized.
func NewDcimFrontPortsCreateParams() *DcimFrontPortsCreateParams {
	var ()
	return &DcimFrontPortsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortsCreateParamsWithTimeout creates a new DcimFrontPortsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimFrontPortsCreateParamsWithTimeout(timeout time.Duration) *DcimFrontPortsCreateParams {
	var ()
	return &DcimFrontPortsCreateParams{

		timeout: timeout,
	}
}

// NewDcimFrontPortsCreateParamsWithContext creates a new DcimFrontPortsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimFrontPortsCreateParamsWithContext(ctx context.Context) *DcimFrontPortsCreateParams {
	var ()
	return &DcimFrontPortsCreateParams{

		Context: ctx,
	}
}

// NewDcimFrontPortsCreateParamsWithHTTPClient creates a new DcimFrontPortsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimFrontPortsCreateParamsWithHTTPClient(client *http.Client) *DcimFrontPortsCreateParams {
	var ()
	return &DcimFrontPortsCreateParams{
		HTTPClient: client,
	}
}

/*DcimFrontPortsCreateParams contains all the parameters to send to the API endpoint
for the dcim front ports create operation typically these are written to a http.Request
*/
type DcimFrontPortsCreateParams struct {

	/*Data*/
	Data *models.WritableFrontPort

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim front ports create params
func (o *DcimFrontPortsCreateParams) WithTimeout(timeout time.Duration) *DcimFrontPortsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front ports create params
func (o *DcimFrontPortsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front ports create params
func (o *DcimFrontPortsCreateParams) WithContext(ctx context.Context) *DcimFrontPortsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front ports create params
func (o *DcimFrontPortsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front ports create params
func (o *DcimFrontPortsCreateParams) WithHTTPClient(client *http.Client) *DcimFrontPortsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front ports create params
func (o *DcimFrontPortsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim front ports create params
func (o *DcimFrontPortsCreateParams) WithData(data *models.WritableFrontPort) *DcimFrontPortsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim front ports create params
func (o *DcimFrontPortsCreateParams) SetData(data *models.WritableFrontPort) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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