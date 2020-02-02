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

	models "github.com/h0x91b-wix/go-netbox/netbox/models"
)

// NewSecretsSecretRolesPartialUpdateParams creates a new SecretsSecretRolesPartialUpdateParams object
// with the default values initialized.
func NewSecretsSecretRolesPartialUpdateParams() *SecretsSecretRolesPartialUpdateParams {
	var ()
	return &SecretsSecretRolesPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSecretsSecretRolesPartialUpdateParamsWithTimeout creates a new SecretsSecretRolesPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSecretsSecretRolesPartialUpdateParamsWithTimeout(timeout time.Duration) *SecretsSecretRolesPartialUpdateParams {
	var ()
	return &SecretsSecretRolesPartialUpdateParams{

		timeout: timeout,
	}
}

// NewSecretsSecretRolesPartialUpdateParamsWithContext creates a new SecretsSecretRolesPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSecretsSecretRolesPartialUpdateParamsWithContext(ctx context.Context) *SecretsSecretRolesPartialUpdateParams {
	var ()
	return &SecretsSecretRolesPartialUpdateParams{

		Context: ctx,
	}
}

// NewSecretsSecretRolesPartialUpdateParamsWithHTTPClient creates a new SecretsSecretRolesPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSecretsSecretRolesPartialUpdateParamsWithHTTPClient(client *http.Client) *SecretsSecretRolesPartialUpdateParams {
	var ()
	return &SecretsSecretRolesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*SecretsSecretRolesPartialUpdateParams contains all the parameters to send to the API endpoint
for the secrets secret roles partial update operation typically these are written to a http.Request
*/
type SecretsSecretRolesPartialUpdateParams struct {

	/*Data*/
	Data *models.SecretRole
	/*ID
	  A unique integer value identifying this secret role.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the secrets secret roles partial update params
func (o *SecretsSecretRolesPartialUpdateParams) WithTimeout(timeout time.Duration) *SecretsSecretRolesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secrets secret roles partial update params
func (o *SecretsSecretRolesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secrets secret roles partial update params
func (o *SecretsSecretRolesPartialUpdateParams) WithContext(ctx context.Context) *SecretsSecretRolesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secrets secret roles partial update params
func (o *SecretsSecretRolesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secrets secret roles partial update params
func (o *SecretsSecretRolesPartialUpdateParams) WithHTTPClient(client *http.Client) *SecretsSecretRolesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secrets secret roles partial update params
func (o *SecretsSecretRolesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the secrets secret roles partial update params
func (o *SecretsSecretRolesPartialUpdateParams) WithData(data *models.SecretRole) *SecretsSecretRolesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the secrets secret roles partial update params
func (o *SecretsSecretRolesPartialUpdateParams) SetData(data *models.SecretRole) {
	o.Data = data
}

// WithID adds the id to the secrets secret roles partial update params
func (o *SecretsSecretRolesPartialUpdateParams) WithID(id int64) *SecretsSecretRolesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the secrets secret roles partial update params
func (o *SecretsSecretRolesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SecretsSecretRolesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
