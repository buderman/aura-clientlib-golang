/*
content-delivery

Aura LCDN Content Delivery API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package content_delivery

import (
	"encoding/json"
	"fmt"
)

// TlsBufferSizeType Sets the size of the TLS buffer used for sending data (KB).
type TlsBufferSizeType int32

// List of tlsBufferSizeType
const (
	TLSBUFFERSIZETYPE__4 TlsBufferSizeType = 4
	TLSBUFFERSIZETYPE__16 TlsBufferSizeType = 16
	TLSBUFFERSIZETYPE__32 TlsBufferSizeType = 32
)

// All allowed values of TlsBufferSizeType enum
var AllowedTlsBufferSizeTypeEnumValues = []TlsBufferSizeType{
	4,
	16,
	32,
}

func (v *TlsBufferSizeType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TlsBufferSizeType(value)
	for _, existing := range AllowedTlsBufferSizeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TlsBufferSizeType", value)
}

// NewTlsBufferSizeTypeFromValue returns a pointer to a valid TlsBufferSizeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTlsBufferSizeTypeFromValue(v int32) (*TlsBufferSizeType, error) {
	ev := TlsBufferSizeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TlsBufferSizeType: valid values are %v", v, AllowedTlsBufferSizeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TlsBufferSizeType) IsValid() bool {
	for _, existing := range AllowedTlsBufferSizeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tlsBufferSizeType value
func (v TlsBufferSizeType) Ptr() *TlsBufferSizeType {
	return &v
}

type NullableTlsBufferSizeType struct {
	value *TlsBufferSizeType
	isSet bool
}

func (v NullableTlsBufferSizeType) Get() *TlsBufferSizeType {
	return v.value
}

func (v *NullableTlsBufferSizeType) Set(val *TlsBufferSizeType) {
	v.value = val
	v.isSet = true
}

func (v NullableTlsBufferSizeType) IsSet() bool {
	return v.isSet
}

func (v *NullableTlsBufferSizeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTlsBufferSizeType(val *TlsBufferSizeType) *NullableTlsBufferSizeType {
	return &NullableTlsBufferSizeType{value: val, isSet: true}
}

func (v NullableTlsBufferSizeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTlsBufferSizeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

