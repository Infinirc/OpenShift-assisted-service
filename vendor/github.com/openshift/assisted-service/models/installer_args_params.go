// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InstallerArgsParams installer args params
//
// swagger:model installer-args-params
type InstallerArgsParams struct {

	// List of additional arguments passed to coreos-installer
	// Example: ["--append-karg","ip=192.0.2.2::192.0.2.254:255.255.255.0:core0.example.com:enp1s0:none","--save-partindex","1","-n"]
	Args []string `json:"args"`
}

// Validate validates this installer args params
func (m *InstallerArgsParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this installer args params based on context it is used
func (m *InstallerArgsParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InstallerArgsParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstallerArgsParams) UnmarshalBinary(b []byte) error {
	var res InstallerArgsParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
