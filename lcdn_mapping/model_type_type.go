/*
lcdn-mapping

Aura LCDN Mapping API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lcdn_mapping

import (
	"encoding/json"
	"fmt"
)

// TypeType Type of GeoIP data. Specifies the address type used in the GeoIP configuration file. `IPV4` indicates the GeoIP configuration file uses an IPv4 address. `IPV6` indicates the GeoIP configuration file uses an IPv6 address.
type TypeType string

// List of typeType
const (
	TYPETYPE_IPV4 TypeType = "IPV4"
	TYPETYPE_IPV6 TypeType = "IPV6"
)

// All allowed values of TypeType enum
var AllowedTypeTypeEnumValues = []TypeType{
	"IPV4",
	"IPV6",
}

func (v *TypeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeType(value)
	for _, existing := range AllowedTypeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeType", value)
}

// NewTypeTypeFromValue returns a pointer to a valid TypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeTypeFromValue(v string) (*TypeType, error) {
	ev := TypeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypeType: valid values are %v", v, AllowedTypeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeType) IsValid() bool {
	for _, existing := range AllowedTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to typeType value
func (v TypeType) Ptr() *TypeType {
	return &v
}

type NullableTypeType struct {
	value *TypeType
	isSet bool
}

func (v NullableTypeType) Get() *TypeType {
	return v.value
}

func (v *NullableTypeType) Set(val *TypeType) {
	v.value = val
	v.isSet = true
}

func (v NullableTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullableTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypeType(val *TypeType) *NullableTypeType {
	return &NullableTypeType{value: val, isSet: true}
}

func (v NullableTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

