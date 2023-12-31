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
)

// checks if the AppleChallengeResponseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppleChallengeResponseRequest{}

// AppleChallengeResponseRequest Pseudo class for plex response
type AppleChallengeResponseRequest struct {
	Component *string `json:"component,omitempty"`
}

// NewAppleChallengeResponseRequest instantiates a new AppleChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppleChallengeResponseRequest() *AppleChallengeResponseRequest {
	this := AppleChallengeResponseRequest{}
	var component string = "ak-source-oauth-apple"
	this.Component = &component
	return &this
}

// NewAppleChallengeResponseRequestWithDefaults instantiates a new AppleChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppleChallengeResponseRequestWithDefaults() *AppleChallengeResponseRequest {
	this := AppleChallengeResponseRequest{}
	var component string = "ak-source-oauth-apple"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AppleChallengeResponseRequest) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AppleChallengeResponseRequest) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AppleChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

func (o AppleChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppleChallengeResponseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	return toSerialize, nil
}

type NullableAppleChallengeResponseRequest struct {
	value *AppleChallengeResponseRequest
	isSet bool
}

func (v NullableAppleChallengeResponseRequest) Get() *AppleChallengeResponseRequest {
	return v.value
}

func (v *NullableAppleChallengeResponseRequest) Set(val *AppleChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppleChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppleChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppleChallengeResponseRequest(val *AppleChallengeResponseRequest) *NullableAppleChallengeResponseRequest {
	return &NullableAppleChallengeResponseRequest{value: val, isSet: true}
}

func (v NullableAppleChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppleChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


