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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritablePowerPortTemplate writable power port template
// swagger:model WritablePowerPortTemplate
type WritablePowerPortTemplate struct {

	// Allocated draw
	//
	// Allocated power draw (watts)
	// Maximum: 32767
	// Minimum: 1
	AllocatedDraw *int64 `json:"allocated_draw,omitempty"`

	// Device type
	// Required: true
	DeviceType *int64 `json:"device_type"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Maximum draw
	//
	// Maximum power draw (watts)
	// Maximum: 32767
	// Minimum: 1
	MaximumDraw *int64 `json:"maximum_draw,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// Type
	// Enum: [iec-60320-c6 iec-60320-c8 iec-60320-c14 iec-60320-c16 iec-60320-c20 iec-60309-p-n-e-4h iec-60309-p-n-e-6h iec-60309-p-n-e-9h iec-60309-2p-e-4h iec-60309-2p-e-6h iec-60309-2p-e-9h iec-60309-3p-e-4h iec-60309-3p-e-6h iec-60309-3p-e-9h iec-60309-3p-n-e-4h iec-60309-3p-n-e-6h iec-60309-3p-n-e-9h nema-5-15p nema-5-20p nema-5-30p nema-5-50p nema-6-15p nema-6-20p nema-6-30p nema-6-50p nema-l5-15p nema-l5-20p nema-l5-30p nema-l5-50p nema-l6-20p nema-l6-30p nema-l6-50p cs6361c cs6365c cs8165c cs8265c cs8365c cs8465c ita-e ita-f ita-ef ita-g ita-h ita-i ita-j ita-k ita-l ita-m ita-n ita-o]
	Type string `json:"type,omitempty"`
}

// Validate validates this writable power port template
func (m *WritablePowerPortTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocatedDraw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaximumDraw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *WritablePowerPortTemplate) validateAllocatedDraw(formats strfmt.Registry) error {

	if swag.IsZero(m.AllocatedDraw) { // not required
		return nil
	}

	if err := validate.MinimumInt("allocated_draw", "body", int64(*m.AllocatedDraw), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("allocated_draw", "body", int64(*m.AllocatedDraw), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPortTemplate) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPortTemplate) validateMaximumDraw(formats strfmt.Registry) error {

	if swag.IsZero(m.MaximumDraw) { // not required
		return nil
	}

	if err := validate.MinimumInt("maximum_draw", "body", int64(*m.MaximumDraw), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("maximum_draw", "body", int64(*m.MaximumDraw), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPortTemplate) validateName(formats strfmt.Registry) error {

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

var writablePowerPortTemplateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["iec-60320-c6","iec-60320-c8","iec-60320-c14","iec-60320-c16","iec-60320-c20","iec-60309-p-n-e-4h","iec-60309-p-n-e-6h","iec-60309-p-n-e-9h","iec-60309-2p-e-4h","iec-60309-2p-e-6h","iec-60309-2p-e-9h","iec-60309-3p-e-4h","iec-60309-3p-e-6h","iec-60309-3p-e-9h","iec-60309-3p-n-e-4h","iec-60309-3p-n-e-6h","iec-60309-3p-n-e-9h","nema-5-15p","nema-5-20p","nema-5-30p","nema-5-50p","nema-6-15p","nema-6-20p","nema-6-30p","nema-6-50p","nema-l5-15p","nema-l5-20p","nema-l5-30p","nema-l5-50p","nema-l6-20p","nema-l6-30p","nema-l6-50p","cs6361c","cs6365c","cs8165c","cs8265c","cs8365c","cs8465c","ita-e","ita-f","ita-ef","ita-g","ita-h","ita-i","ita-j","ita-k","ita-l","ita-m","ita-n","ita-o"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerPortTemplateTypeTypePropEnum = append(writablePowerPortTemplateTypeTypePropEnum, v)
	}
}

const (

	// WritablePowerPortTemplateTypeIec60320C6 captures enum value "iec-60320-c6"
	WritablePowerPortTemplateTypeIec60320C6 string = "iec-60320-c6"

	// WritablePowerPortTemplateTypeIec60320C8 captures enum value "iec-60320-c8"
	WritablePowerPortTemplateTypeIec60320C8 string = "iec-60320-c8"

	// WritablePowerPortTemplateTypeIec60320C14 captures enum value "iec-60320-c14"
	WritablePowerPortTemplateTypeIec60320C14 string = "iec-60320-c14"

	// WritablePowerPortTemplateTypeIec60320C16 captures enum value "iec-60320-c16"
	WritablePowerPortTemplateTypeIec60320C16 string = "iec-60320-c16"

	// WritablePowerPortTemplateTypeIec60320C20 captures enum value "iec-60320-c20"
	WritablePowerPortTemplateTypeIec60320C20 string = "iec-60320-c20"

	// WritablePowerPortTemplateTypeIec60309pne4h captures enum value "iec-60309-p-n-e-4h"
	WritablePowerPortTemplateTypeIec60309pne4h string = "iec-60309-p-n-e-4h"

	// WritablePowerPortTemplateTypeIec60309pne6h captures enum value "iec-60309-p-n-e-6h"
	WritablePowerPortTemplateTypeIec60309pne6h string = "iec-60309-p-n-e-6h"

	// WritablePowerPortTemplateTypeIec60309pne9h captures enum value "iec-60309-p-n-e-9h"
	WritablePowerPortTemplateTypeIec60309pne9h string = "iec-60309-p-n-e-9h"

	// WritablePowerPortTemplateTypeIec603092pe4h captures enum value "iec-60309-2p-e-4h"
	WritablePowerPortTemplateTypeIec603092pe4h string = "iec-60309-2p-e-4h"

	// WritablePowerPortTemplateTypeIec603092pe6h captures enum value "iec-60309-2p-e-6h"
	WritablePowerPortTemplateTypeIec603092pe6h string = "iec-60309-2p-e-6h"

	// WritablePowerPortTemplateTypeIec603092pe9h captures enum value "iec-60309-2p-e-9h"
	WritablePowerPortTemplateTypeIec603092pe9h string = "iec-60309-2p-e-9h"

	// WritablePowerPortTemplateTypeIec603093pe4h captures enum value "iec-60309-3p-e-4h"
	WritablePowerPortTemplateTypeIec603093pe4h string = "iec-60309-3p-e-4h"

	// WritablePowerPortTemplateTypeIec603093pe6h captures enum value "iec-60309-3p-e-6h"
	WritablePowerPortTemplateTypeIec603093pe6h string = "iec-60309-3p-e-6h"

	// WritablePowerPortTemplateTypeIec603093pe9h captures enum value "iec-60309-3p-e-9h"
	WritablePowerPortTemplateTypeIec603093pe9h string = "iec-60309-3p-e-9h"

	// WritablePowerPortTemplateTypeIec603093pne4h captures enum value "iec-60309-3p-n-e-4h"
	WritablePowerPortTemplateTypeIec603093pne4h string = "iec-60309-3p-n-e-4h"

	// WritablePowerPortTemplateTypeIec603093pne6h captures enum value "iec-60309-3p-n-e-6h"
	WritablePowerPortTemplateTypeIec603093pne6h string = "iec-60309-3p-n-e-6h"

	// WritablePowerPortTemplateTypeIec603093pne9h captures enum value "iec-60309-3p-n-e-9h"
	WritablePowerPortTemplateTypeIec603093pne9h string = "iec-60309-3p-n-e-9h"

	// WritablePowerPortTemplateTypeNema515p captures enum value "nema-5-15p"
	WritablePowerPortTemplateTypeNema515p string = "nema-5-15p"

	// WritablePowerPortTemplateTypeNema520p captures enum value "nema-5-20p"
	WritablePowerPortTemplateTypeNema520p string = "nema-5-20p"

	// WritablePowerPortTemplateTypeNema530p captures enum value "nema-5-30p"
	WritablePowerPortTemplateTypeNema530p string = "nema-5-30p"

	// WritablePowerPortTemplateTypeNema550p captures enum value "nema-5-50p"
	WritablePowerPortTemplateTypeNema550p string = "nema-5-50p"

	// WritablePowerPortTemplateTypeNema615p captures enum value "nema-6-15p"
	WritablePowerPortTemplateTypeNema615p string = "nema-6-15p"

	// WritablePowerPortTemplateTypeNema620p captures enum value "nema-6-20p"
	WritablePowerPortTemplateTypeNema620p string = "nema-6-20p"

	// WritablePowerPortTemplateTypeNema630p captures enum value "nema-6-30p"
	WritablePowerPortTemplateTypeNema630p string = "nema-6-30p"

	// WritablePowerPortTemplateTypeNema650p captures enum value "nema-6-50p"
	WritablePowerPortTemplateTypeNema650p string = "nema-6-50p"

	// WritablePowerPortTemplateTypeNemaL515p captures enum value "nema-l5-15p"
	WritablePowerPortTemplateTypeNemaL515p string = "nema-l5-15p"

	// WritablePowerPortTemplateTypeNemaL520p captures enum value "nema-l5-20p"
	WritablePowerPortTemplateTypeNemaL520p string = "nema-l5-20p"

	// WritablePowerPortTemplateTypeNemaL530p captures enum value "nema-l5-30p"
	WritablePowerPortTemplateTypeNemaL530p string = "nema-l5-30p"

	// WritablePowerPortTemplateTypeNemaL550p captures enum value "nema-l5-50p"
	WritablePowerPortTemplateTypeNemaL550p string = "nema-l5-50p"

	// WritablePowerPortTemplateTypeNemaL620p captures enum value "nema-l6-20p"
	WritablePowerPortTemplateTypeNemaL620p string = "nema-l6-20p"

	// WritablePowerPortTemplateTypeNemaL630p captures enum value "nema-l6-30p"
	WritablePowerPortTemplateTypeNemaL630p string = "nema-l6-30p"

	// WritablePowerPortTemplateTypeNemaL650p captures enum value "nema-l6-50p"
	WritablePowerPortTemplateTypeNemaL650p string = "nema-l6-50p"

	// WritablePowerPortTemplateTypeCs6361c captures enum value "cs6361c"
	WritablePowerPortTemplateTypeCs6361c string = "cs6361c"

	// WritablePowerPortTemplateTypeCs6365c captures enum value "cs6365c"
	WritablePowerPortTemplateTypeCs6365c string = "cs6365c"

	// WritablePowerPortTemplateTypeCs8165c captures enum value "cs8165c"
	WritablePowerPortTemplateTypeCs8165c string = "cs8165c"

	// WritablePowerPortTemplateTypeCs8265c captures enum value "cs8265c"
	WritablePowerPortTemplateTypeCs8265c string = "cs8265c"

	// WritablePowerPortTemplateTypeCs8365c captures enum value "cs8365c"
	WritablePowerPortTemplateTypeCs8365c string = "cs8365c"

	// WritablePowerPortTemplateTypeCs8465c captures enum value "cs8465c"
	WritablePowerPortTemplateTypeCs8465c string = "cs8465c"

	// WritablePowerPortTemplateTypeItae captures enum value "ita-e"
	WritablePowerPortTemplateTypeItae string = "ita-e"

	// WritablePowerPortTemplateTypeItaf captures enum value "ita-f"
	WritablePowerPortTemplateTypeItaf string = "ita-f"

	// WritablePowerPortTemplateTypeItaEf captures enum value "ita-ef"
	WritablePowerPortTemplateTypeItaEf string = "ita-ef"

	// WritablePowerPortTemplateTypeItag captures enum value "ita-g"
	WritablePowerPortTemplateTypeItag string = "ita-g"

	// WritablePowerPortTemplateTypeItah captures enum value "ita-h"
	WritablePowerPortTemplateTypeItah string = "ita-h"

	// WritablePowerPortTemplateTypeItai captures enum value "ita-i"
	WritablePowerPortTemplateTypeItai string = "ita-i"

	// WritablePowerPortTemplateTypeItaj captures enum value "ita-j"
	WritablePowerPortTemplateTypeItaj string = "ita-j"

	// WritablePowerPortTemplateTypeItak captures enum value "ita-k"
	WritablePowerPortTemplateTypeItak string = "ita-k"

	// WritablePowerPortTemplateTypeItal captures enum value "ita-l"
	WritablePowerPortTemplateTypeItal string = "ita-l"

	// WritablePowerPortTemplateTypeItam captures enum value "ita-m"
	WritablePowerPortTemplateTypeItam string = "ita-m"

	// WritablePowerPortTemplateTypeItan captures enum value "ita-n"
	WritablePowerPortTemplateTypeItan string = "ita-n"

	// WritablePowerPortTemplateTypeItao captures enum value "ita-o"
	WritablePowerPortTemplateTypeItao string = "ita-o"
)

// prop value enum
func (m *WritablePowerPortTemplate) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writablePowerPortTemplateTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerPortTemplate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritablePowerPortTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritablePowerPortTemplate) UnmarshalBinary(b []byte) error {
	var res WritablePowerPortTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
