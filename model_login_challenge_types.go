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

// LoginChallengeTypes - struct for LoginChallengeTypes
type LoginChallengeTypes struct {
	AppleLoginChallenge *AppleLoginChallenge
	PlexAuthenticationChallenge *PlexAuthenticationChallenge
	RedirectChallenge *RedirectChallenge
}

// AppleLoginChallengeAsLoginChallengeTypes is a convenience function that returns AppleLoginChallenge wrapped in LoginChallengeTypes
func AppleLoginChallengeAsLoginChallengeTypes(v *AppleLoginChallenge) LoginChallengeTypes {
	return LoginChallengeTypes{
		AppleLoginChallenge: v,
	}
}

// PlexAuthenticationChallengeAsLoginChallengeTypes is a convenience function that returns PlexAuthenticationChallenge wrapped in LoginChallengeTypes
func PlexAuthenticationChallengeAsLoginChallengeTypes(v *PlexAuthenticationChallenge) LoginChallengeTypes {
	return LoginChallengeTypes{
		PlexAuthenticationChallenge: v,
	}
}

// RedirectChallengeAsLoginChallengeTypes is a convenience function that returns RedirectChallenge wrapped in LoginChallengeTypes
func RedirectChallengeAsLoginChallengeTypes(v *RedirectChallenge) LoginChallengeTypes {
	return LoginChallengeTypes{
		RedirectChallenge: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *LoginChallengeTypes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AppleLoginChallenge
	err = newStrictDecoder(data).Decode(&dst.AppleLoginChallenge)
	if err == nil {
		jsonAppleLoginChallenge, _ := json.Marshal(dst.AppleLoginChallenge)
		if string(jsonAppleLoginChallenge) == "{}" { // empty struct
			dst.AppleLoginChallenge = nil
		} else {
			match++
		}
	} else {
		dst.AppleLoginChallenge = nil
	}

	// try to unmarshal data into PlexAuthenticationChallenge
	err = newStrictDecoder(data).Decode(&dst.PlexAuthenticationChallenge)
	if err == nil {
		jsonPlexAuthenticationChallenge, _ := json.Marshal(dst.PlexAuthenticationChallenge)
		if string(jsonPlexAuthenticationChallenge) == "{}" { // empty struct
			dst.PlexAuthenticationChallenge = nil
		} else {
			match++
		}
	} else {
		dst.PlexAuthenticationChallenge = nil
	}

	// try to unmarshal data into RedirectChallenge
	err = newStrictDecoder(data).Decode(&dst.RedirectChallenge)
	if err == nil {
		jsonRedirectChallenge, _ := json.Marshal(dst.RedirectChallenge)
		if string(jsonRedirectChallenge) == "{}" { // empty struct
			dst.RedirectChallenge = nil
		} else {
			match++
		}
	} else {
		dst.RedirectChallenge = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AppleLoginChallenge = nil
		dst.PlexAuthenticationChallenge = nil
		dst.RedirectChallenge = nil

		return fmt.Errorf("data matches more than one schema in oneOf(LoginChallengeTypes)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(LoginChallengeTypes)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LoginChallengeTypes) MarshalJSON() ([]byte, error) {
	if src.AppleLoginChallenge != nil {
		return json.Marshal(&src.AppleLoginChallenge)
	}

	if src.PlexAuthenticationChallenge != nil {
		return json.Marshal(&src.PlexAuthenticationChallenge)
	}

	if src.RedirectChallenge != nil {
		return json.Marshal(&src.RedirectChallenge)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LoginChallengeTypes) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AppleLoginChallenge != nil {
		return obj.AppleLoginChallenge
	}

	if obj.PlexAuthenticationChallenge != nil {
		return obj.PlexAuthenticationChallenge
	}

	if obj.RedirectChallenge != nil {
		return obj.RedirectChallenge
	}

	// all schemas are nil
	return nil
}

type NullableLoginChallengeTypes struct {
	value *LoginChallengeTypes
	isSet bool
}

func (v NullableLoginChallengeTypes) Get() *LoginChallengeTypes {
	return v.value
}

func (v *NullableLoginChallengeTypes) Set(val *LoginChallengeTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginChallengeTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginChallengeTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginChallengeTypes(val *LoginChallengeTypes) *NullableLoginChallengeTypes {
	return &NullableLoginChallengeTypes{value: val, isSet: true}
}

func (v NullableLoginChallengeTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginChallengeTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


