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

// PowerFeed power feed
// swagger:model PowerFeed
type PowerFeed struct {

	// Amperage
	// Maximum: 32767
	// Minimum: 1
	Amperage int64 `json:"amperage,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Max utilization
	//
	// Maximum permissible draw (percentage)
	// Maximum: 100
	// Minimum: 1
	MaxUtilization int64 `json:"max_utilization,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// phase
	Phase *PowerFeedPhase `json:"phase,omitempty"`

	// power panel
	// Required: true
	PowerPanel *NestedPowerPanel `json:"power_panel"`

	// rack
	Rack *NestedRack `json:"rack,omitempty"`

	// status
	Status *PowerFeedStatus `json:"status,omitempty"`

	// supply
	Supply *PowerFeedSupply `json:"supply,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// type
	Type *PowerFeedType `json:"type,omitempty"`

	// Voltage
	// Maximum: 32767
	// Minimum: 1
	Voltage int64 `json:"voltage,omitempty"`
}

// Validate validates this power feed
func (m *PowerFeed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmperage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxUtilization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerPanel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupply(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoltage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PowerFeed) validateAmperage(formats strfmt.Registry) error {

	if swag.IsZero(m.Amperage) { // not required
		return nil
	}

	if err := validate.MinimumInt("amperage", "body", int64(m.Amperage), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("amperage", "body", int64(m.Amperage), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) validateMaxUtilization(formats strfmt.Registry) error {

	if swag.IsZero(m.MaxUtilization) { // not required
		return nil
	}

	if err := validate.MinimumInt("max_utilization", "body", int64(m.MaxUtilization), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("max_utilization", "body", int64(m.MaxUtilization), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) validatePhase(formats strfmt.Registry) error {

	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	if m.Phase != nil {
		if err := m.Phase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) validatePowerPanel(formats strfmt.Registry) error {

	if err := validate.Required("power_panel", "body", m.PowerPanel); err != nil {
		return err
	}

	if m.PowerPanel != nil {
		if err := m.PowerPanel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("power_panel")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) validateRack(formats strfmt.Registry) error {

	if swag.IsZero(m.Rack) { // not required
		return nil
	}

	if m.Rack != nil {
		if err := m.Rack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rack")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) validateSupply(formats strfmt.Registry) error {

	if swag.IsZero(m.Supply) { // not required
		return nil
	}

	if m.Supply != nil {
		if err := m.Supply.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("supply")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) validateTags(formats strfmt.Registry) error {

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

func (m *PowerFeed) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) validateVoltage(formats strfmt.Registry) error {

	if swag.IsZero(m.Voltage) { // not required
		return nil
	}

	if err := validate.MinimumInt("voltage", "body", int64(m.Voltage), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("voltage", "body", int64(m.Voltage), 32767, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerFeed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerFeed) UnmarshalBinary(b []byte) error {
	var res PowerFeed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerFeedPhase Phase
// swagger:model PowerFeedPhase
type PowerFeedPhase struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *string `json:"value"`
}

func (m *PowerFeedPhase) UnmarshalJSON(b []byte) error {
	type PowerFeedPhaseAlias PowerFeedPhase
	var t PowerFeedPhaseAlias
	if err := json.Unmarshal([]byte("{\"id\":1,\"label\":\"Single phase\",\"value\":\"single-phase\"}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = PowerFeedPhase(t)
	return nil
}

// Validate validates this power feed phase
func (m *PowerFeedPhase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PowerFeedPhase) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("phase"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeedPhase) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("phase"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerFeedPhase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerFeedPhase) UnmarshalBinary(b []byte) error {
	var res PowerFeedPhase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerFeedStatus Status
// swagger:model PowerFeedStatus
type PowerFeedStatus struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *string `json:"value"`
}

func (m *PowerFeedStatus) UnmarshalJSON(b []byte) error {
	type PowerFeedStatusAlias PowerFeedStatus
	var t PowerFeedStatusAlias
	if err := json.Unmarshal([]byte("{\"id\":1,\"label\":\"Active\",\"value\":\"active\"}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = PowerFeedStatus(t)
	return nil
}

// Validate validates this power feed status
func (m *PowerFeedStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PowerFeedStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeedStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerFeedStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerFeedStatus) UnmarshalBinary(b []byte) error {
	var res PowerFeedStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerFeedSupply Supply
// swagger:model PowerFeedSupply
type PowerFeedSupply struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *string `json:"value"`
}

func (m *PowerFeedSupply) UnmarshalJSON(b []byte) error {
	type PowerFeedSupplyAlias PowerFeedSupply
	var t PowerFeedSupplyAlias
	if err := json.Unmarshal([]byte("{\"id\":1,\"label\":\"AC\",\"value\":\"ac\"}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = PowerFeedSupply(t)
	return nil
}

// Validate validates this power feed supply
func (m *PowerFeedSupply) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PowerFeedSupply) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("supply"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeedSupply) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("supply"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerFeedSupply) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerFeedSupply) UnmarshalBinary(b []byte) error {
	var res PowerFeedSupply
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerFeedType Type
// swagger:model PowerFeedType
type PowerFeedType struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *string `json:"value"`
}

func (m *PowerFeedType) UnmarshalJSON(b []byte) error {
	type PowerFeedTypeAlias PowerFeedType
	var t PowerFeedTypeAlias
	if err := json.Unmarshal([]byte("{\"id\":1,\"label\":\"Primary\",\"value\":\"primary\"}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = PowerFeedType(t)
	return nil
}

// Validate validates this power feed type
func (m *PowerFeedType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PowerFeedType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeedType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerFeedType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerFeedType) UnmarshalBinary(b []byte) error {
	var res PowerFeedType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}