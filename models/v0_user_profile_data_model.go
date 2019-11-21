// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V0UserProfileDataModel v0 user profile data model
// swagger:model v0.UserProfileDataModel
type V0UserProfileDataModel struct {

	// avatar url
	AvatarURL string `json:"avatar_url,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// is chat available
	IsChatAvailable bool `json:"is_chat_available,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// unconfirmed email
	UnconfirmedEmail string `json:"unconfirmed_email,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this v0 user profile data model
func (m *V0UserProfileDataModel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0UserProfileDataModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0UserProfileDataModel) UnmarshalBinary(b []byte) error {
	var res V0UserProfileDataModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
