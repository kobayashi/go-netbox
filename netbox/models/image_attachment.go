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
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ImageAttachment image attachment
// swagger:model ImageAttachment
type ImageAttachment struct {

	// Content type
	// Required: true
	ContentType *string `json:"content_type"`

	// Created
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Image
	// Read Only: true
	// Format: uri
	Image strfmt.URI `json:"image,omitempty"`

	// Image height
	// Required: true
	// Maximum: 32767
	// Minimum: 0
	ImageHeight *int64 `json:"image_height"`

	// Image width
	// Required: true
	// Maximum: 32767
	// Minimum: 0
	ImageWidth *int64 `json:"image_width"`

	// Name
	// Max Length: 50
	Name string `json:"name,omitempty"`

	// Object id
	// Required: true
	// Maximum: 2.147483647e+09
	// Minimum: 0
	ObjectID *int64 `json:"object_id"`

	// Parent
	// Read Only: true
	Parent map[string]string `json:"parent,omitempty"`
}

// Validate validates this image attachment
func (m *ImageAttachment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageWidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageAttachment) validateContentType(formats strfmt.Registry) error {

	if err := validate.Required("content_type", "body", m.ContentType); err != nil {
		return err
	}

	return nil
}

func (m *ImageAttachment) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ImageAttachment) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if err := validate.FormatOf("image", "body", "uri", m.Image.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ImageAttachment) validateImageHeight(formats strfmt.Registry) error {

	if err := validate.Required("image_height", "body", m.ImageHeight); err != nil {
		return err
	}

	if err := validate.MinimumInt("image_height", "body", int64(*m.ImageHeight), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("image_height", "body", int64(*m.ImageHeight), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *ImageAttachment) validateImageWidth(formats strfmt.Registry) error {

	if err := validate.Required("image_width", "body", m.ImageWidth); err != nil {
		return err
	}

	if err := validate.MinimumInt("image_width", "body", int64(*m.ImageWidth), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("image_width", "body", int64(*m.ImageWidth), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *ImageAttachment) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 50); err != nil {
		return err
	}

	return nil
}

func (m *ImageAttachment) validateObjectID(formats strfmt.Registry) error {

	if err := validate.Required("object_id", "body", m.ObjectID); err != nil {
		return err
	}

	if err := validate.MinimumInt("object_id", "body", int64(*m.ObjectID), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("object_id", "body", int64(*m.ObjectID), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageAttachment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageAttachment) UnmarshalBinary(b []byte) error {
	var res ImageAttachment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
