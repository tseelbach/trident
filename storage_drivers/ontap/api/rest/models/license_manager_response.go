// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LicenseManagerResponse Information on license manager instances associated with the cluster.
//
// swagger:model license_manager_response
type LicenseManagerResponse struct {

	// links
	Links *CollectionLinks `json:"_links,omitempty"`

	// Number of records
	NumRecords int64 `json:"num_records,omitempty"`

	// records
	Records []*LicenseManagerResponseRecordsItems0 `json:"records,omitempty"`
}

// Validate validates this license manager response
func (m *LicenseManagerResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LicenseManagerResponse) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *LicenseManagerResponse) validateRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.Records) { // not required
		return nil
	}

	for i := 0; i < len(m.Records); i++ {
		if swag.IsZero(m.Records[i]) { // not required
			continue
		}

		if m.Records[i] != nil {
			if err := m.Records[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this license manager response based on the context it is used
func (m *LicenseManagerResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LicenseManagerResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *LicenseManagerResponse) contextValidateRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Records); i++ {

		if m.Records[i] != nil {
			if err := m.Records[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LicenseManagerResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicenseManagerResponse) UnmarshalBinary(b []byte) error {
	var res LicenseManagerResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LicenseManagerResponseRecordsItems0 Information on a license manager instance associated with the cluster.
//
// swagger:model LicenseManagerResponseRecordsItems0
type LicenseManagerResponseRecordsItems0 struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// Flag that indicates whether it's the default license manager instance used by the cluster.'
	// When a capacity pool is created and if the license manager field is omitted, it is assumed that the license of the capacity pool is installed on the default license manager instance.
	//
	// Read Only: true
	Default *bool `json:"default,omitempty"`

	// uri
	URI *LicenseManagerResponseRecordsItems0URI `json:"uri,omitempty"`

	// uuid
	// Example: 4ea7a442-86d1-11e0-ae1c-112233445566
	// Read Only: true
	// Format: uuid
	UUID strfmt.UUID `json:"uuid,omitempty"`
}

// Validate validates this license manager response records items0
func (m *LicenseManagerResponseRecordsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LicenseManagerResponseRecordsItems0) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *LicenseManagerResponseRecordsItems0) validateURI(formats strfmt.Registry) error {
	if swag.IsZero(m.URI) { // not required
		return nil
	}

	if m.URI != nil {
		if err := m.URI.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uri")
			}
			return err
		}
	}

	return nil
}

func (m *LicenseManagerResponseRecordsItems0) validateUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.FormatOf("uuid", "body", "uuid", m.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this license manager response records items0 based on the context it is used
func (m *LicenseManagerResponseRecordsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefault(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LicenseManagerResponseRecordsItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *LicenseManagerResponseRecordsItems0) contextValidateDefault(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "default", "body", m.Default); err != nil {
		return err
	}

	return nil
}

func (m *LicenseManagerResponseRecordsItems0) contextValidateURI(ctx context.Context, formats strfmt.Registry) error {

	if m.URI != nil {
		if err := m.URI.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uri")
			}
			return err
		}
	}

	return nil
}

func (m *LicenseManagerResponseRecordsItems0) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", strfmt.UUID(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LicenseManagerResponseRecordsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicenseManagerResponseRecordsItems0) UnmarshalBinary(b []byte) error {
	var res LicenseManagerResponseRecordsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LicenseManagerResponseRecordsItems0URI License manager URI.
//
// swagger:model LicenseManagerResponseRecordsItems0URI
type LicenseManagerResponseRecordsItems0URI struct {

	// License manager host name, IPv4 or IPv6 address.
	// Example: 10.1.1.1
	Host string `json:"host,omitempty"`
}

// Validate validates this license manager response records items0 URI
func (m *LicenseManagerResponseRecordsItems0URI) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this license manager response records items0 URI based on context it is used
func (m *LicenseManagerResponseRecordsItems0URI) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LicenseManagerResponseRecordsItems0URI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicenseManagerResponseRecordsItems0URI) UnmarshalBinary(b []byte) error {
	var res LicenseManagerResponseRecordsItems0URI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}