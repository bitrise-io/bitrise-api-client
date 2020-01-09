// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V0AppleAPICredentialResponseItem v0 apple API credential response item
// swagger:model v0.AppleAPICredentialResponseItem
type V0AppleAPICredentialResponseItem struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// issuer id
	IssuerID string `json:"issuer_id,omitempty"`

	// key id
	KeyID string `json:"key_id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// private key
	PrivateKey string `json:"private_key,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this v0 apple API credential response item
func (m *V0AppleAPICredentialResponseItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0AppleAPICredentialResponseItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0AppleAPICredentialResponseItem) UnmarshalBinary(b []byte) error {
	var res V0AppleAPICredentialResponseItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
