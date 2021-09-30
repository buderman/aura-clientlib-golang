/*
infrastructure

Infrastructure API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infrastructure

import (
	"encoding/json"
	"fmt"
)

// BootStateType When booting, the action the node should take `BOOT`, `REINSTALL`, or `UPGRADE`. In `BOOT` state, the node completes a normal boot process and runs in production mode. `REINSTALL` causes the node to reinstall and overwrite the system partition and boot partition on all disks and the associated data, but does not update the kernel partition. Do not use `UPGRADE` to upgrade the software on a service node; use the upgrade_step.py script. For detailed information about upgrading a service node using the upgrade_step.py script, see the \"Service Node Upgrade with upgrade_step.py\" in the  [Aura LCDN/LMS Upgrade Guide](https://control.akamai.com/wh/CUSTOMER/AKAMAI/en-US/WEBHELP/portal-guides/aura-licensed-cdn/GUID-867C3390-F48F-4845-9EDD-BDD5F917B71A.html).
type BootStateType string

// List of bootStateType
const (
	BOOTSTATETYPE_BOOT BootStateType = "BOOT"
	BOOTSTATETYPE_REINSTALL BootStateType = "REINSTALL"
	BOOTSTATETYPE_UPGRADE BootStateType = "UPGRADE"
)

var allowedBootStateTypeEnumValues = []BootStateType{
	"BOOT",
	"REINSTALL",
	"UPGRADE",
}

func (v *BootStateType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BootStateType(value)
	for _, existing := range allowedBootStateTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BootStateType", value)
}

// NewBootStateTypeFromValue returns a pointer to a valid BootStateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBootStateTypeFromValue(v string) (*BootStateType, error) {
	ev := BootStateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BootStateType: valid values are %v", v, allowedBootStateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BootStateType) IsValid() bool {
	for _, existing := range allowedBootStateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to bootStateType value
func (v BootStateType) Ptr() *BootStateType {
	return &v
}

type NullableBootStateType struct {
	value *BootStateType
	isSet bool
}

func (v NullableBootStateType) Get() *BootStateType {
	return v.value
}

func (v *NullableBootStateType) Set(val *BootStateType) {
	v.value = val
	v.isSet = true
}

func (v NullableBootStateType) IsSet() bool {
	return v.isSet
}

func (v *NullableBootStateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootStateType(val *BootStateType) *NullableBootStateType {
	return &NullableBootStateType{value: val, isSet: true}
}

func (v NullableBootStateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootStateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
