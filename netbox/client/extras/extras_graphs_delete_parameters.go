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

package extras

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

// NewExtrasGraphsDeleteParams creates a new ExtrasGraphsDeleteParams object
// with the default values initialized.
func NewExtrasGraphsDeleteParams() *ExtrasGraphsDeleteParams {
	var ()
	return &ExtrasGraphsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGraphsDeleteParamsWithTimeout creates a new ExtrasGraphsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasGraphsDeleteParamsWithTimeout(timeout time.Duration) *ExtrasGraphsDeleteParams {
	var ()
	return &ExtrasGraphsDeleteParams{

		timeout: timeout,
	}
}

// NewExtrasGraphsDeleteParamsWithContext creates a new ExtrasGraphsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasGraphsDeleteParamsWithContext(ctx context.Context) *ExtrasGraphsDeleteParams {
	var ()
	return &ExtrasGraphsDeleteParams{

		Context: ctx,
	}
}

// NewExtrasGraphsDeleteParamsWithHTTPClient creates a new ExtrasGraphsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasGraphsDeleteParamsWithHTTPClient(client *http.Client) *ExtrasGraphsDeleteParams {
	var ()
	return &ExtrasGraphsDeleteParams{
		HTTPClient: client,
	}
}

/*ExtrasGraphsDeleteParams contains all the parameters to send to the API endpoint
for the extras graphs delete operation typically these are written to a http.Request
*/
type ExtrasGraphsDeleteParams struct {

	/*ID
	  A unique integer value identifying this graph.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras graphs delete params
func (o *ExtrasGraphsDeleteParams) WithTimeout(timeout time.Duration) *ExtrasGraphsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras graphs delete params
func (o *ExtrasGraphsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras graphs delete params
func (o *ExtrasGraphsDeleteParams) WithContext(ctx context.Context) *ExtrasGraphsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras graphs delete params
func (o *ExtrasGraphsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras graphs delete params
func (o *ExtrasGraphsDeleteParams) WithHTTPClient(client *http.Client) *ExtrasGraphsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras graphs delete params
func (o *ExtrasGraphsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras graphs delete params
func (o *ExtrasGraphsDeleteParams) WithID(id int64) *ExtrasGraphsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras graphs delete params
func (o *ExtrasGraphsDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGraphsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
