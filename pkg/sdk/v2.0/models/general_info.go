// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GeneralInfo general info
//
// swagger:model GeneralInfo
type GeneralInfo struct {

	// The auth mode of current Harbor instance.
	AuthMode *string `json:"auth_mode,omitempty"`

	// The setting of auth proxy this is only available when Harbor relies on authproxy for authentication.
	AuthproxySettings *AuthproxySetting `json:"authproxy_settings,omitempty"`

	// The banner message for the UI. It is the stringified result of the banner message object.
	// Example: {\"closable\":true,\"message\":\"your banner message content\",\"type\":\"warning\",\"fromDate\":\"06/19/2023\",\"toDate\":\"06/21/2023\"}
	BannerMessage *string `json:"banner_message,omitempty"`

	// The current time of the server.
	// Format: date-time
	CurrentTime *strfmt.DateTime `json:"current_time,omitempty"`

	// The external URL of Harbor, with protocol.
	ExternalURL *string `json:"external_url,omitempty"`

	// The build version of Harbor.
	HarborVersion *string `json:"harbor_version,omitempty"`

	// Indicate whether there is a ca root cert file ready for download in the file system.
	HasCaRoot *bool `json:"has_ca_root,omitempty"`

	// The flag to indicate whether notification mechanism is enabled on Harbor instance.
	NotificationEnable *bool `json:"notification_enable,omitempty"`

	// The flag to indicate whether the current auth mode should consider as a primary one.
	PrimaryAuthMode *bool `json:"primary_auth_mode,omitempty"`

	// Indicate who can create projects, it could be 'adminonly' or 'everyone'.
	ProjectCreationRestriction *string `json:"project_creation_restriction,omitempty"`

	// The flag to indicate whether Harbor is in readonly mode.
	ReadOnly *bool `json:"read_only,omitempty"`

	// The storage provider's name of Harbor registry
	RegistryStorageProviderName *string `json:"registry_storage_provider_name,omitempty"`

	// The url of registry against which the docker command should be issued.
	RegistryURL *string `json:"registry_url,omitempty"`

	// Indicate whether the Harbor instance enable user to register himself.
	SelfRegistration *bool `json:"self_registration,omitempty"`
}

// Validate validates this general info
func (m *GeneralInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthproxySettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GeneralInfo) validateAuthproxySettings(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthproxySettings) { // not required
		return nil
	}

	if m.AuthproxySettings != nil {
		if err := m.AuthproxySettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authproxy_settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authproxy_settings")
			}
			return err
		}
	}

	return nil
}

func (m *GeneralInfo) validateCurrentTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CurrentTime) { // not required
		return nil
	}

	if err := validate.FormatOf("current_time", "body", "date-time", m.CurrentTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this general info based on the context it is used
func (m *GeneralInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthproxySettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GeneralInfo) contextValidateAuthproxySettings(ctx context.Context, formats strfmt.Registry) error {

	if m.AuthproxySettings != nil {
		if err := m.AuthproxySettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authproxy_settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authproxy_settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GeneralInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GeneralInfo) UnmarshalBinary(b []byte) error {
	var res GeneralInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
