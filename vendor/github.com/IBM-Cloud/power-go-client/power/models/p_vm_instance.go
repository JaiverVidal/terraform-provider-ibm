// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PVMInstance p VM instance
// swagger:model PVMInstance
type PVMInstance struct {

	// (deprecated - replaced by networks) The list of addresses and their network information
	Addresses []*PVMInstanceNetwork `json:"addresses"`

	// Date/Time of PVM creation
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// Size of allocated disk (in GB)
	// Required: true
	DiskSize *float64 `json:"diskSize"`

	// fault
	Fault *PVMInstanceFault `json:"fault,omitempty"`

	// health
	Health *PVMInstanceHealth `json:"health,omitempty"`

	// The ImageID used by the server
	// Required: true
	ImageID *string `json:"imageID"`

	// Maximum amount of memory that can be allocated (in GB, for resize)
	Maxmem float64 `json:"maxmem,omitempty"`

	// Maximum number of processors that can be allocated (for resize)
	Maxproc float64 `json:"maxproc,omitempty"`

	// Amount of memory allocated (in GB)
	// Required: true
	Memory *float64 `json:"memory"`

	// whether the instance can be migrated
	Migratable *bool `json:"migratable,omitempty"`

	// Minimum amount of memory that can be allocated (in GB, for resize)
	Minmem float64 `json:"minmem,omitempty"`

	// Minimum number of processors that can be allocated (for resize)
	Minproc float64 `json:"minproc,omitempty"`

	// (deprecated - replaced by networks) List of Network IDs
	// Required: true
	NetworkIds []string `json:"networkIDs"`

	// The pvm instance networks information
	Networks []*PVMInstanceNetwork `json:"networks"`

	// Processor type (dedicated, shared, capped)
	// Required: true
	// Enum: [dedicated shared capped]
	ProcType *string `json:"procType"`

	// Number of processors allocated
	// Required: true
	Processors *float64 `json:"processors"`

	// The progress of an operation
	Progress float64 `json:"progress,omitempty"`

	// PCloud PVM Instance ID
	// Required: true
	PvmInstanceID *string `json:"pvmInstanceID"`

	// Name of the server
	// Required: true
	ServerName *string `json:"serverName"`

	// The pvm instance Software Licenses
	SoftwareLicenses *SoftwareLicenses `json:"softwareLicenses,omitempty"`

	// The status of the instance
	// Required: true
	Status *string `json:"status"`

	// Storage type where server is deployed
	// Required: true
	StorageType *string `json:"storageType"`

	// System type used to host the instance
	SysType string `json:"sysType,omitempty"`

	// Date/Time of PVM last update
	// Format: date-time
	UpdatedDate strfmt.DateTime `json:"updatedDate,omitempty"`

	// List of volume IDs
	// Required: true
	VolumeIds []string `json:"volumeIDs"`
}

// Validate validates this p VM instance
func (m *PVMInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiskSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFault(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePvmInstanceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareLicenses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PVMInstance) validateAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.Addresses) { // not required
		return nil
	}

	for i := 0; i < len(m.Addresses); i++ {
		if swag.IsZero(m.Addresses[i]) { // not required
			continue
		}

		if m.Addresses[i] != nil {
			if err := m.Addresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PVMInstance) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateDiskSize(formats strfmt.Registry) error {

	if err := validate.Required("diskSize", "body", m.DiskSize); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateFault(formats strfmt.Registry) error {

	if swag.IsZero(m.Fault) { // not required
		return nil
	}

	if m.Fault != nil {
		if err := m.Fault.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fault")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstance) validateHealth(formats strfmt.Registry) error {

	if swag.IsZero(m.Health) { // not required
		return nil
	}

	if m.Health != nil {
		if err := m.Health.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstance) validateImageID(formats strfmt.Registry) error {

	if err := validate.Required("imageID", "body", m.ImageID); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateNetworkIds(formats strfmt.Registry) error {

	if err := validate.Required("networkIDs", "body", m.NetworkIds); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateNetworks(formats strfmt.Registry) error {

	if swag.IsZero(m.Networks) { // not required
		return nil
	}

	for i := 0; i < len(m.Networks); i++ {
		if swag.IsZero(m.Networks[i]) { // not required
			continue
		}

		if m.Networks[i] != nil {
			if err := m.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var pVmInstanceTypeProcTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dedicated","shared","capped"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pVmInstanceTypeProcTypePropEnum = append(pVmInstanceTypeProcTypePropEnum, v)
	}
}

const (

	// PVMInstanceProcTypeDedicated captures enum value "dedicated"
	PVMInstanceProcTypeDedicated string = "dedicated"

	// PVMInstanceProcTypeShared captures enum value "shared"
	PVMInstanceProcTypeShared string = "shared"

	// PVMInstanceProcTypeCapped captures enum value "capped"
	PVMInstanceProcTypeCapped string = "capped"
)

// prop value enum
func (m *PVMInstance) validateProcTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, pVmInstanceTypeProcTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PVMInstance) validateProcType(formats strfmt.Registry) error {

	if err := validate.Required("procType", "body", m.ProcType); err != nil {
		return err
	}

	// value enum
	if err := m.validateProcTypeEnum("procType", "body", *m.ProcType); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateProcessors(formats strfmt.Registry) error {

	if err := validate.Required("processors", "body", m.Processors); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validatePvmInstanceID(formats strfmt.Registry) error {

	if err := validate.Required("pvmInstanceID", "body", m.PvmInstanceID); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateServerName(formats strfmt.Registry) error {

	if err := validate.Required("serverName", "body", m.ServerName); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateSoftwareLicenses(formats strfmt.Registry) error {

	if swag.IsZero(m.SoftwareLicenses) { // not required
		return nil
	}

	if m.SoftwareLicenses != nil {
		if err := m.SoftwareLicenses.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("softwareLicenses")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstance) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateStorageType(formats strfmt.Registry) error {

	if err := validate.Required("storageType", "body", m.StorageType); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateUpdatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedDate", "body", "date-time", m.UpdatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateVolumeIds(formats strfmt.Registry) error {

	if err := validate.Required("volumeIDs", "body", m.VolumeIds); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PVMInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PVMInstance) UnmarshalBinary(b []byte) error {
	var res PVMInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
