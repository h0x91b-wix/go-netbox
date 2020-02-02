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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableAggregate writable aggregate
// swagger:model WritableAggregate
type WritableAggregate struct {

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Date added
	// Format: date
	DateAdded *strfmt.Date `json:"date_added,omitempty"`

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// Family
	// Read Only: true
	// Enum: [4 6]
	Family int64 `json:"family,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Prefix
	// Required: true
	Prefix *string `json:"prefix"`

	// RIR
	// Required: true
	Rir *int64 `json:"rir"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this writable aggregate
func (m *WritableAggregate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateAdded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFamily(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRir(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableAggregate) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableAggregate) validateDateAdded(formats strfmt.Registry) error {

	if swag.IsZero(m.DateAdded) { // not required
		return nil
	}

	if err := validate.FormatOf("date_added", "body", "date", m.DateAdded.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableAggregate) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

var writableAggregateTypeFamilyPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[4,6]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableAggregateTypeFamilyPropEnum = append(writableAggregateTypeFamilyPropEnum, v)
	}
}

// prop value enum
func (m *WritableAggregate) validateFamilyEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, writableAggregateTypeFamilyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableAggregate) validateFamily(formats strfmt.Registry) error {

	if swag.IsZero(m.Family) { // not required
		return nil
	}

	// value enum
	if err := m.validateFamilyEnum("family", "body", m.Family); err != nil {
		return err
	}

	return nil
}

func (m *WritableAggregate) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableAggregate) validatePrefix(formats strfmt.Registry) error {

	if err := validate.Required("prefix", "body", m.Prefix); err != nil {
		return err
	}

	return nil
}

func (m *WritableAggregate) validateRir(formats strfmt.Registry) error {

	if err := validate.Required("rir", "body", m.Rir); err != nil {
		return err
	}

	return nil
}

func (m *WritableAggregate) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.MinLength("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableAggregate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableAggregate) UnmarshalBinary(b []byte) error {
	var res WritableAggregate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
