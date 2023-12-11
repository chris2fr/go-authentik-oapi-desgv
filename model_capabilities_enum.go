/*
authentik

Making authentication simple.

API version: 2023.10.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// CapabilitiesEnum * `can_save_media` - Can Save Media * `can_geo_ip` - Can Geo Ip * `can_impersonate` - Can Impersonate * `can_debug` - Can Debug * `is_enterprise` - Is Enterprise
type CapabilitiesEnum string

// List of CapabilitiesEnum
const (
	CAN_SAVE_MEDIA CapabilitiesEnum = "can_save_media"
	CAN_GEO_IP CapabilitiesEnum = "can_geo_ip"
	CAN_IMPERSONATE CapabilitiesEnum = "can_impersonate"
	CAN_DEBUG CapabilitiesEnum = "can_debug"
	IS_ENTERPRISE CapabilitiesEnum = "is_enterprise"
)

// All allowed values of CapabilitiesEnum enum
var AllowedCapabilitiesEnumEnumValues = []CapabilitiesEnum{
	"can_save_media",
	"can_geo_ip",
	"can_impersonate",
	"can_debug",
	"is_enterprise",
}

func (v *CapabilitiesEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CapabilitiesEnum(value)
	for _, existing := range AllowedCapabilitiesEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CapabilitiesEnum", value)
}

// NewCapabilitiesEnumFromValue returns a pointer to a valid CapabilitiesEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCapabilitiesEnumFromValue(v string) (*CapabilitiesEnum, error) {
	ev := CapabilitiesEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CapabilitiesEnum: valid values are %v", v, AllowedCapabilitiesEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CapabilitiesEnum) IsValid() bool {
	for _, existing := range AllowedCapabilitiesEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CapabilitiesEnum value
func (v CapabilitiesEnum) Ptr() *CapabilitiesEnum {
	return &v
}

type NullableCapabilitiesEnum struct {
	value *CapabilitiesEnum
	isSet bool
}

func (v NullableCapabilitiesEnum) Get() *CapabilitiesEnum {
	return v.value
}

func (v *NullableCapabilitiesEnum) Set(val *CapabilitiesEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitiesEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitiesEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitiesEnum(val *CapabilitiesEnum) *NullableCapabilitiesEnum {
	return &NullableCapabilitiesEnum{value: val, isSet: true}
}

func (v NullableCapabilitiesEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitiesEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

