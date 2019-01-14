// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V0PlanDataModel v0 plan data model
// swagger:model v0.PlanDataModel
type V0PlanDataModel struct {

	// container count
	ContainerCount string `json:"container_count,omitempty"`

	// expires at
	ExpiresAt string `json:"expires_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// price
	Price int64 `json:"price,omitempty"`
}

// Validate validates this v0 plan data model
func (m *V0PlanDataModel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0PlanDataModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0PlanDataModel) UnmarshalBinary(b []byte) error {
	var res V0PlanDataModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
