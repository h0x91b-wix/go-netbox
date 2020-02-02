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
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PowerOutlet power outlet
// swagger:model PowerOutlet
type PowerOutlet struct {

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Connected endpoint
	//
	//
	//         Return the appropriate serializer for the type of connected object.
	//
	// Read Only: true
	ConnectedEndpoint map[string]string `json:"connected_endpoint,omitempty"`

	// Connected endpoint type
	// Read Only: true
	ConnectedEndpointType string `json:"connected_endpoint_type,omitempty"`

	// connection status
	ConnectionStatus *PowerOutletConnectionStatus `json:"connection_status,omitempty"`

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// device
	// Required: true
	Device *NestedDevice `json:"device"`

	// feed leg
	FeedLeg *PowerOutletFeedLeg `json:"feed_leg,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// power port
	PowerPort *NestedPowerPort `json:"power_port,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// type
	Type *PowerOutletType `json:"type,omitempty"`
}

// Validate validates this power outlet
func (m *PowerOutlet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeedLeg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PowerOutlet) validateCable(formats strfmt.Registry) error {

	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *PowerOutlet) validateConnectionStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionStatus) { // not required
		return nil
	}

	if m.ConnectionStatus != nil {
		if err := m.ConnectionStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection_status")
			}
			return err
		}
	}

	return nil
}

func (m *PowerOutlet) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutlet) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *PowerOutlet) validateFeedLeg(formats strfmt.Registry) error {

	if swag.IsZero(m.FeedLeg) { // not required
		return nil
	}

	if m.FeedLeg != nil {
		if err := m.FeedLeg.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("feed_leg")
			}
			return err
		}
	}

	return nil
}

func (m *PowerOutlet) validateName(formats strfmt.Registry) error {

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

func (m *PowerOutlet) validatePowerPort(formats strfmt.Registry) error {

	if swag.IsZero(m.PowerPort) { // not required
		return nil
	}

	if m.PowerPort != nil {
		if err := m.PowerPort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("power_port")
			}
			return err
		}
	}

	return nil
}

func (m *PowerOutlet) validateTags(formats strfmt.Registry) error {

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

func (m *PowerOutlet) validateType(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *PowerOutlet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerOutlet) UnmarshalBinary(b []byte) error {
	var res PowerOutlet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerOutletConnectionStatus Connection status
// swagger:model PowerOutletConnectionStatus
type PowerOutletConnectionStatus struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *bool `json:"value"`
}

// Validate validates this power outlet connection status
func (m *PowerOutletConnectionStatus) Validate(formats strfmt.Registry) error {
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

func (m *PowerOutletConnectionStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("connection_status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutletConnectionStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("connection_status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerOutletConnectionStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerOutletConnectionStatus) UnmarshalBinary(b []byte) error {
	var res PowerOutletConnectionStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerOutletFeedLeg Feed leg
// swagger:model PowerOutletFeedLeg
type PowerOutletFeedLeg struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this power outlet feed leg
func (m *PowerOutletFeedLeg) Validate(formats strfmt.Registry) error {
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

func (m *PowerOutletFeedLeg) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("feed_leg"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutletFeedLeg) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("feed_leg"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerOutletFeedLeg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerOutletFeedLeg) UnmarshalBinary(b []byte) error {
	var res PowerOutletFeedLeg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerOutletType Type
// swagger:model PowerOutletType
type PowerOutletType struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this power outlet type
func (m *PowerOutletType) Validate(formats strfmt.Registry) error {
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

func (m *PowerOutletType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutletType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerOutletType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerOutletType) UnmarshalBinary(b []byte) error {
	var res PowerOutletType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
