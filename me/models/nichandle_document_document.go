// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017 The go-ovh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NichandleDocumentDocument List of documents added on your account
// swagger:model nichandle.Document.Document
type NichandleDocumentDocument struct {

	// Document creation
	// Required: true
	// Read Only: true
	CreationDate strfmt.DateTime `json:"creationDate"`

	// Document expiration
	ExpirationDate strfmt.DateTime `json:"expirationDate,omitempty"`

	// URL used to get document
	// Required: true
	// Read Only: true
	GetURL string `json:"getUrl"`

	// Document id
	// Required: true
	// Read Only: true
	ID string `json:"id"`

	// Document name
	// Required: true
	// Read Only: true
	Name string `json:"name"`

	// URL used to put document
	// Required: true
	// Read Only: true
	PutURL string `json:"putUrl"`

	// Document size (in bytes)
	// Required: true
	// Read Only: true
	Size int64 `json:"size"`

	// tags
	// Required: true
	Tags NichandleDocumentDocumentTags `json:"tags"`

	// Document validation
	// Read Only: true
	ValidationDate strfmt.DateTime `json:"validationDate,omitempty"`
}

// Validate validates this nichandle document document
func (m *NichandleDocumentDocument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGetURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePutURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NichandleDocumentDocument) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", strfmt.DateTime(m.CreationDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NichandleDocumentDocument) validateGetURL(formats strfmt.Registry) error {

	if err := validate.RequiredString("getUrl", "body", string(m.GetURL)); err != nil {
		return err
	}

	return nil
}

func (m *NichandleDocumentDocument) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *NichandleDocumentDocument) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *NichandleDocumentDocument) validatePutURL(formats strfmt.Registry) error {

	if err := validate.RequiredString("putUrl", "body", string(m.PutURL)); err != nil {
		return err
	}

	return nil
}

func (m *NichandleDocumentDocument) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", int64(m.Size)); err != nil {
		return err
	}

	return nil
}

func (m *NichandleDocumentDocument) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags); err != nil {
		return err
	}

	if err := m.Tags.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tags")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NichandleDocumentDocument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NichandleDocumentDocument) UnmarshalBinary(b []byte) error {
	var res NichandleDocumentDocument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}