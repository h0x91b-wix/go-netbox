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

	models "github.com/h0x91b-wix/go-netbox/netbox/models"
)

// NewDcimConsolePortsUpdateParams creates a new DcimConsolePortsUpdateParams object
// with the default values initialized.
func NewDcimConsolePortsUpdateParams() *DcimConsolePortsUpdateParams {
	var ()
	return &DcimConsolePortsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortsUpdateParamsWithTimeout creates a new DcimConsolePortsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimConsolePortsUpdateParamsWithTimeout(timeout time.Duration) *DcimConsolePortsUpdateParams {
	var ()
	return &DcimConsolePortsUpdateParams{

		timeout: timeout,
	}
}

// NewDcimConsolePortsUpdateParamsWithContext creates a new DcimConsolePortsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimConsolePortsUpdateParamsWithContext(ctx context.Context) *DcimConsolePortsUpdateParams {
	var ()
	return &DcimConsolePortsUpdateParams{

		Context: ctx,
	}
}

// NewDcimConsolePortsUpdateParamsWithHTTPClient creates a new DcimConsolePortsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimConsolePortsUpdateParamsWithHTTPClient(client *http.Client) *DcimConsolePortsUpdateParams {
	var ()
	return &DcimConsolePortsUpdateParams{
		HTTPClient: client,
	}
}

/*DcimConsolePortsUpdateParams contains all the parameters to send to the API endpoint
for the dcim console ports update operation typically these are written to a http.Request
*/
type DcimConsolePortsUpdateParams struct {

	/*Data*/
	Data *models.WritableConsolePort
	/*ID
	  A unique integer value identifying this console port.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim console ports update params
func (o *DcimConsolePortsUpdateParams) WithTimeout(timeout time.Duration) *DcimConsolePortsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console ports update params
func (o *DcimConsolePortsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console ports update params
func (o *DcimConsolePortsUpdateParams) WithContext(ctx context.Context) *DcimConsolePortsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console ports update params
func (o *DcimConsolePortsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console ports update params
func (o *DcimConsolePortsUpdateParams) WithHTTPClient(client *http.Client) *DcimConsolePortsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console ports update params
func (o *DcimConsolePortsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console ports update params
func (o *DcimConsolePortsUpdateParams) WithData(data *models.WritableConsolePort) *DcimConsolePortsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console ports update params
func (o *DcimConsolePortsUpdateParams) SetData(data *models.WritableConsolePort) {
	o.Data = data
}

// WithID adds the id to the dcim console ports update params
func (o *DcimConsolePortsUpdateParams) WithID(id int64) *DcimConsolePortsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console ports update params
func (o *DcimConsolePortsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
