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

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Secret secret
// swagger:model Secret
type Secret struct {

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// device
	// Required: true
	Device *NestedDevice `json:"device"`

	// Hash
	// Read Only: true
	// Min Length: 1
	Hash string `json:"hash,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Name
	// Max Length: 100
	Name string `json:"name,omitempty"`

	// Plaintext
	// Required: true
	// Min Length: 1
	Plaintext *string `json:"plaintext"`

	// role
	// Required: true
	Role *NestedSecretRole `json:"role"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this secret
func (m *Secret) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlaintext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
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

func (m *Secret) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Secret) validateDevice(formats strfmt.Registry) error {

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

func (m *Secret) validateHash(formats strfmt.Registry) error {

	if swag.IsZero(m.Hash) { // not required
		return nil
	}

	if err := validate.MinLength("hash", "body", string(m.Hash), 1); err != nil {
		return err
	}

	return nil
}

func (m *Secret) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Secret) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 100); err != nil {
		return err
	}

	return nil
}

func (m *Secret) validatePlaintext(formats strfmt.Registry) error {

	if err := validate.Required("plaintext", "body", m.Plaintext); err != nil {
		return err
	}

	if err := validate.MinLength("plaintext", "body", string(*m.Plaintext), 1); err != nil {
		return err
	}

	return nil
}

func (m *Secret) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *Secret) validateTags(formats strfmt.Registry) error {

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
func (m *Secret) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secret) UnmarshalBinary(b []byte) error {
	var res Secret
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
