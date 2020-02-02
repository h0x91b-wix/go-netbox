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

	strfmt "github.com/go-openapi/strfmt"
)

// NewExtrasScriptsReadParams creates a new ExtrasScriptsReadParams object
// with the default values initialized.
func NewExtrasScriptsReadParams() *ExtrasScriptsReadParams {
	var ()
	return &ExtrasScriptsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasScriptsReadParamsWithTimeout creates a new ExtrasScriptsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasScriptsReadParamsWithTimeout(timeout time.Duration) *ExtrasScriptsReadParams {
	var ()
	return &ExtrasScriptsReadParams{

		timeout: timeout,
	}
}

// NewExtrasScriptsReadParamsWithContext creates a new ExtrasScriptsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasScriptsReadParamsWithContext(ctx context.Context) *ExtrasScriptsReadParams {
	var ()
	return &ExtrasScriptsReadParams{

		Context: ctx,
	}
}

// NewExtrasScriptsReadParamsWithHTTPClient creates a new ExtrasScriptsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasScriptsReadParamsWithHTTPClient(client *http.Client) *ExtrasScriptsReadParams {
	var ()
	return &ExtrasScriptsReadParams{
		HTTPClient: client,
	}
}

/*ExtrasScriptsReadParams contains all the parameters to send to the API endpoint
for the extras scripts read operation typically these are written to a http.Request
*/
type ExtrasScriptsReadParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras scripts read params
func (o *ExtrasScriptsReadParams) WithTimeout(timeout time.Duration) *ExtrasScriptsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras scripts read params
func (o *ExtrasScriptsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras scripts read params
func (o *ExtrasScriptsReadParams) WithContext(ctx context.Context) *ExtrasScriptsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras scripts read params
func (o *ExtrasScriptsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras scripts read params
func (o *ExtrasScriptsReadParams) WithHTTPClient(client *http.Client) *ExtrasScriptsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras scripts read params
func (o *ExtrasScriptsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras scripts read params
func (o *ExtrasScriptsReadParams) WithID(id string) *ExtrasScriptsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras scripts read params
func (o *ExtrasScriptsReadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasScriptsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
