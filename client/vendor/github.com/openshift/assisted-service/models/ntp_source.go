// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NtpSource ntp source
//
// swagger:model ntp_source
type NtpSource struct {

	// NTP source name or IP.
	SourceName string `json:"source_name,omitempty"`

	// Indication of state of an NTP source.
	SourceState SourceState `json:"source_state,omitempty"`
}

// Validate validates this ntp source
func (m *NtpSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSourceState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NtpSource) validateSourceState(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceState) { // not required
		return nil
	}

	if err := m.SourceState.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("source_state")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("source_state")
		}
		return err
	}

	return nil
}

// ContextValidate validate this ntp source based on the context it is used
func (m *NtpSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSourceState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NtpSource) contextValidateSourceState(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SourceState.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("source_state")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("source_state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NtpSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NtpSource) UnmarshalBinary(b []byte) error {
	var res NtpSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
