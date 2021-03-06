/*
infrastructure

Aura Infrastructure API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infrastructure

import (
	"encoding/json"
	"fmt"
)

// InterfaceType - Base interface type
type InterfaceType struct {
	BondInterfaceType *BondInterfaceType
	EthernetInterfaceType *EthernetInterfaceType
}

// BondInterfaceTypeAsInterfaceType is a convenience function that returns BondInterfaceType wrapped in InterfaceType
func BondInterfaceTypeAsInterfaceType(v *BondInterfaceType) InterfaceType {
	return InterfaceType{ BondInterfaceType: v}
}

// EthernetInterfaceTypeAsInterfaceType is a convenience function that returns EthernetInterfaceType wrapped in InterfaceType
func EthernetInterfaceTypeAsInterfaceType(v *EthernetInterfaceType) InterfaceType {
	return InterfaceType{ EthernetInterfaceType: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *InterfaceType) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BOND'
	if jsonDict["type"] == "BOND" {
		// try to unmarshal JSON data into BondInterfaceType
		err = json.Unmarshal(data, &dst.BondInterfaceType)
		if err == nil {
			return nil // data stored in dst.BondInterfaceType, return on the first match
		} else {
			dst.BondInterfaceType = nil
			return fmt.Errorf("Failed to unmarshal InterfaceType as BondInterfaceType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ETHERNET'
	if jsonDict["type"] == "ETHERNET" {
		// try to unmarshal JSON data into EthernetInterfaceType
		err = json.Unmarshal(data, &dst.EthernetInterfaceType)
		if err == nil {
			return nil // data stored in dst.EthernetInterfaceType, return on the first match
		} else {
			dst.EthernetInterfaceType = nil
			return fmt.Errorf("Failed to unmarshal InterfaceType as EthernetInterfaceType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'bondInterfaceType'
	if jsonDict["type"] == "bondInterfaceType" {
		// try to unmarshal JSON data into BondInterfaceType
		err = json.Unmarshal(data, &dst.BondInterfaceType)
		if err == nil {
			return nil // data stored in dst.BondInterfaceType, return on the first match
		} else {
			dst.BondInterfaceType = nil
			return fmt.Errorf("Failed to unmarshal InterfaceType as BondInterfaceType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ethernetInterfaceType'
	if jsonDict["type"] == "ethernetInterfaceType" {
		// try to unmarshal JSON data into EthernetInterfaceType
		err = json.Unmarshal(data, &dst.EthernetInterfaceType)
		if err == nil {
			return nil // data stored in dst.EthernetInterfaceType, return on the first match
		} else {
			dst.EthernetInterfaceType = nil
			return fmt.Errorf("Failed to unmarshal InterfaceType as EthernetInterfaceType: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src InterfaceType) MarshalJSON() ([]byte, error) {
	if src.BondInterfaceType != nil {
		return json.Marshal(&src.BondInterfaceType)
	}

	if src.EthernetInterfaceType != nil {
		return json.Marshal(&src.EthernetInterfaceType)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *InterfaceType) GetActualInstance() (interface{}) {
	if obj.BondInterfaceType != nil {
		return obj.BondInterfaceType
	}

	if obj.EthernetInterfaceType != nil {
		return obj.EthernetInterfaceType
	}

	// all schemas are nil
	return nil
}

type NullableInterfaceType struct {
	value *InterfaceType
	isSet bool
}

func (v NullableInterfaceType) Get() *InterfaceType {
	return v.value
}

func (v *NullableInterfaceType) Set(val *InterfaceType) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceType) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceType(val *InterfaceType) *NullableInterfaceType {
	return &NullableInterfaceType{value: val, isSet: true}
}

func (v NullableInterfaceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


