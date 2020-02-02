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

// NewDcimInterfacesGraphsParams creates a new DcimInterfacesGraphsParams object
// with the default values initialized.
func NewDcimInterfacesGraphsParams() *DcimInterfacesGraphsParams {
	var ()
	return &DcimInterfacesGraphsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfacesGraphsParamsWithTimeout creates a new DcimInterfacesGraphsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimInterfacesGraphsParamsWithTimeout(timeout time.Duration) *DcimInterfacesGraphsParams {
	var ()
	return &DcimInterfacesGraphsParams{

		timeout: timeout,
	}
}

// NewDcimInterfacesGraphsParamsWithContext creates a new DcimInterfacesGraphsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimInterfacesGraphsParamsWithContext(ctx context.Context) *DcimInterfacesGraphsParams {
	var ()
	return &DcimInterfacesGraphsParams{

		Context: ctx,
	}
}

// NewDcimInterfacesGraphsParamsWithHTTPClient creates a new DcimInterfacesGraphsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimInterfacesGraphsParamsWithHTTPClient(client *http.Client) *DcimInterfacesGraphsParams {
	var ()
	return &DcimInterfacesGraphsParams{
		HTTPClient: client,
	}
}

/*DcimInterfacesGraphsParams contains all the parameters to send to the API endpoint
for the dcim interfaces graphs operation typically these are written to a http.Request
*/
type DcimInterfacesGraphsParams struct {

	/*ID
	  A unique integer value identifying this interface.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim interfaces graphs params
func (o *DcimInterfacesGraphsParams) WithTimeout(timeout time.Duration) *DcimInterfacesGraphsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interfaces graphs params
func (o *DcimInterfacesGraphsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interfaces graphs params
func (o *DcimInterfacesGraphsParams) WithContext(ctx context.Context) *DcimInterfacesGraphsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interfaces graphs params
func (o *DcimInterfacesGraphsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interfaces graphs params
func (o *DcimInterfacesGraphsParams) WithHTTPClient(client *http.Client) *DcimInterfacesGraphsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interfaces graphs params
func (o *DcimInterfacesGraphsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim interfaces graphs params
func (o *DcimInterfacesGraphsParams) WithID(id int64) *DcimInterfacesGraphsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim interfaces graphs params
func (o *DcimInterfacesGraphsParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfacesGraphsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
