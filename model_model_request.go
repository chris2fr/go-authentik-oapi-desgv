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

// ModelRequest - struct for ModelRequest
type ModelRequest struct {
	LDAPProviderRequest *LDAPProviderRequest
	OAuth2ProviderRequest *OAuth2ProviderRequest
	ProxyProviderRequest *ProxyProviderRequest
	RadiusProviderRequest *RadiusProviderRequest
	SAMLProviderRequest *SAMLProviderRequest
	SCIMProviderRequest *SCIMProviderRequest
}

// LDAPProviderRequestAsModelRequest is a convenience function that returns LDAPProviderRequest wrapped in ModelRequest
func LDAPProviderRequestAsModelRequest(v *LDAPProviderRequest) ModelRequest {
	return ModelRequest{
		LDAPProviderRequest: v,
	}
}

// OAuth2ProviderRequestAsModelRequest is a convenience function that returns OAuth2ProviderRequest wrapped in ModelRequest
func OAuth2ProviderRequestAsModelRequest(v *OAuth2ProviderRequest) ModelRequest {
	return ModelRequest{
		OAuth2ProviderRequest: v,
	}
}

// ProxyProviderRequestAsModelRequest is a convenience function that returns ProxyProviderRequest wrapped in ModelRequest
func ProxyProviderRequestAsModelRequest(v *ProxyProviderRequest) ModelRequest {
	return ModelRequest{
		ProxyProviderRequest: v,
	}
}

// RadiusProviderRequestAsModelRequest is a convenience function that returns RadiusProviderRequest wrapped in ModelRequest
func RadiusProviderRequestAsModelRequest(v *RadiusProviderRequest) ModelRequest {
	return ModelRequest{
		RadiusProviderRequest: v,
	}
}

// SAMLProviderRequestAsModelRequest is a convenience function that returns SAMLProviderRequest wrapped in ModelRequest
func SAMLProviderRequestAsModelRequest(v *SAMLProviderRequest) ModelRequest {
	return ModelRequest{
		SAMLProviderRequest: v,
	}
}

// SCIMProviderRequestAsModelRequest is a convenience function that returns SCIMProviderRequest wrapped in ModelRequest
func SCIMProviderRequestAsModelRequest(v *SCIMProviderRequest) ModelRequest {
	return ModelRequest{
		SCIMProviderRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ModelRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LDAPProviderRequest
	err = newStrictDecoder(data).Decode(&dst.LDAPProviderRequest)
	if err == nil {
		jsonLDAPProviderRequest, _ := json.Marshal(dst.LDAPProviderRequest)
		if string(jsonLDAPProviderRequest) == "{}" { // empty struct
			dst.LDAPProviderRequest = nil
		} else {
			match++
		}
	} else {
		dst.LDAPProviderRequest = nil
	}

	// try to unmarshal data into OAuth2ProviderRequest
	err = newStrictDecoder(data).Decode(&dst.OAuth2ProviderRequest)
	if err == nil {
		jsonOAuth2ProviderRequest, _ := json.Marshal(dst.OAuth2ProviderRequest)
		if string(jsonOAuth2ProviderRequest) == "{}" { // empty struct
			dst.OAuth2ProviderRequest = nil
		} else {
			match++
		}
	} else {
		dst.OAuth2ProviderRequest = nil
	}

	// try to unmarshal data into ProxyProviderRequest
	err = newStrictDecoder(data).Decode(&dst.ProxyProviderRequest)
	if err == nil {
		jsonProxyProviderRequest, _ := json.Marshal(dst.ProxyProviderRequest)
		if string(jsonProxyProviderRequest) == "{}" { // empty struct
			dst.ProxyProviderRequest = nil
		} else {
			match++
		}
	} else {
		dst.ProxyProviderRequest = nil
	}

	// try to unmarshal data into RadiusProviderRequest
	err = newStrictDecoder(data).Decode(&dst.RadiusProviderRequest)
	if err == nil {
		jsonRadiusProviderRequest, _ := json.Marshal(dst.RadiusProviderRequest)
		if string(jsonRadiusProviderRequest) == "{}" { // empty struct
			dst.RadiusProviderRequest = nil
		} else {
			match++
		}
	} else {
		dst.RadiusProviderRequest = nil
	}

	// try to unmarshal data into SAMLProviderRequest
	err = newStrictDecoder(data).Decode(&dst.SAMLProviderRequest)
	if err == nil {
		jsonSAMLProviderRequest, _ := json.Marshal(dst.SAMLProviderRequest)
		if string(jsonSAMLProviderRequest) == "{}" { // empty struct
			dst.SAMLProviderRequest = nil
		} else {
			match++
		}
	} else {
		dst.SAMLProviderRequest = nil
	}

	// try to unmarshal data into SCIMProviderRequest
	err = newStrictDecoder(data).Decode(&dst.SCIMProviderRequest)
	if err == nil {
		jsonSCIMProviderRequest, _ := json.Marshal(dst.SCIMProviderRequest)
		if string(jsonSCIMProviderRequest) == "{}" { // empty struct
			dst.SCIMProviderRequest = nil
		} else {
			match++
		}
	} else {
		dst.SCIMProviderRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.LDAPProviderRequest = nil
		dst.OAuth2ProviderRequest = nil
		dst.ProxyProviderRequest = nil
		dst.RadiusProviderRequest = nil
		dst.SAMLProviderRequest = nil
		dst.SCIMProviderRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ModelRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ModelRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ModelRequest) MarshalJSON() ([]byte, error) {
	if src.LDAPProviderRequest != nil {
		return json.Marshal(&src.LDAPProviderRequest)
	}

	if src.OAuth2ProviderRequest != nil {
		return json.Marshal(&src.OAuth2ProviderRequest)
	}

	if src.ProxyProviderRequest != nil {
		return json.Marshal(&src.ProxyProviderRequest)
	}

	if src.RadiusProviderRequest != nil {
		return json.Marshal(&src.RadiusProviderRequest)
	}

	if src.SAMLProviderRequest != nil {
		return json.Marshal(&src.SAMLProviderRequest)
	}

	if src.SCIMProviderRequest != nil {
		return json.Marshal(&src.SCIMProviderRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ModelRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.LDAPProviderRequest != nil {
		return obj.LDAPProviderRequest
	}

	if obj.OAuth2ProviderRequest != nil {
		return obj.OAuth2ProviderRequest
	}

	if obj.ProxyProviderRequest != nil {
		return obj.ProxyProviderRequest
	}

	if obj.RadiusProviderRequest != nil {
		return obj.RadiusProviderRequest
	}

	if obj.SAMLProviderRequest != nil {
		return obj.SAMLProviderRequest
	}

	if obj.SCIMProviderRequest != nil {
		return obj.SCIMProviderRequest
	}

	// all schemas are nil
	return nil
}

type NullableModelRequest struct {
	value *ModelRequest
	isSet bool
}

func (v NullableModelRequest) Get() *ModelRequest {
	return v.value
}

func (v *NullableModelRequest) Set(val *ModelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelRequest(val *ModelRequest) *NullableModelRequest {
	return &NullableModelRequest{value: val, isSet: true}
}

func (v NullableModelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


