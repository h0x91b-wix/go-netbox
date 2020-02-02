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

package secrets

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

// NewSecretsSecretRolesReadParams creates a new SecretsSecretRolesReadParams object
// with the default values initialized.
func NewSecretsSecretRolesReadParams() *SecretsSecretRolesReadParams {
	var ()
	return &SecretsSecretRolesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSecretsSecretRolesReadParamsWithTimeout creates a new SecretsSecretRolesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSecretsSecretRolesReadParamsWithTimeout(timeout time.Duration) *SecretsSecretRolesReadParams {
	var ()
	return &SecretsSecretRolesReadParams{

		timeout: timeout,
	}
}

// NewSecretsSecretRolesReadParamsWithContext creates a new SecretsSecretRolesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewSecretsSecretRolesReadParamsWithContext(ctx context.Context) *SecretsSecretRolesReadParams {
	var ()
	return &SecretsSecretRolesReadParams{

		Context: ctx,
	}
}

// NewSecretsSecretRolesReadParamsWithHTTPClient creates a new SecretsSecretRolesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSecretsSecretRolesReadParamsWithHTTPClient(client *http.Client) *SecretsSecretRolesReadParams {
	var ()
	return &SecretsSecretRolesReadParams{
		HTTPClient: client,
	}
}

/*SecretsSecretRolesReadParams contains all the parameters to send to the API endpoint
for the secrets secret roles read operation typically these are written to a http.Request
*/
type SecretsSecretRolesReadParams struct {

	/*ID
	  A unique integer value identifying this secret role.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the secrets secret roles read params
func (o *SecretsSecretRolesReadParams) WithTimeout(timeout time.Duration) *SecretsSecretRolesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secrets secret roles read params
func (o *SecretsSecretRolesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secrets secret roles read params
func (o *SecretsSecretRolesReadParams) WithContext(ctx context.Context) *SecretsSecretRolesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secrets secret roles read params
func (o *SecretsSecretRolesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secrets secret roles read params
func (o *SecretsSecretRolesReadParams) WithHTTPClient(client *http.Client) *SecretsSecretRolesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secrets secret roles read params
func (o *SecretsSecretRolesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the secrets secret roles read params
func (o *SecretsSecretRolesReadParams) WithID(id int64) *SecretsSecretRolesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the secrets secret roles read params
func (o *SecretsSecretRolesReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SecretsSecretRolesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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