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

// checks if the AuthenticatorSMSChallengeResponseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorSMSChallengeResponseRequest{}

// AuthenticatorSMSChallengeResponseRequest SMS Challenge response, device is set by get_response_instance
type AuthenticatorSMSChallengeResponseRequest struct {
	Component *string `json:"component,omitempty"`
	Code *int32 `json:"code,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
}

// NewAuthenticatorSMSChallengeResponseRequest instantiates a new AuthenticatorSMSChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorSMSChallengeResponseRequest() *AuthenticatorSMSChallengeResponseRequest {
	this := AuthenticatorSMSChallengeResponseRequest{}
	var component string = "ak-stage-authenticator-sms"
	this.Component = &component
	return &this
}

// NewAuthenticatorSMSChallengeResponseRequestWithDefaults instantiates a new AuthenticatorSMSChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorSMSChallengeResponseRequestWithDefaults() *AuthenticatorSMSChallengeResponseRequest {
	this := AuthenticatorSMSChallengeResponseRequest{}
	var component string = "ak-stage-authenticator-sms"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AuthenticatorSMSChallengeResponseRequest) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AuthenticatorSMSChallengeResponseRequest) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AuthenticatorSMSChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AuthenticatorSMSChallengeResponseRequest) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSChallengeResponseRequest) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AuthenticatorSMSChallengeResponseRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *AuthenticatorSMSChallengeResponseRequest) SetCode(v int32) {
	o.Code = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *AuthenticatorSMSChallengeResponseRequest) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSChallengeResponseRequest) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *AuthenticatorSMSChallengeResponseRequest) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *AuthenticatorSMSChallengeResponseRequest) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

func (o AuthenticatorSMSChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorSMSChallengeResponseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	return toSerialize, nil
}

type NullableAuthenticatorSMSChallengeResponseRequest struct {
	value *AuthenticatorSMSChallengeResponseRequest
	isSet bool
}

func (v NullableAuthenticatorSMSChallengeResponseRequest) Get() *AuthenticatorSMSChallengeResponseRequest {
	return v.value
}

func (v *NullableAuthenticatorSMSChallengeResponseRequest) Set(val *AuthenticatorSMSChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorSMSChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorSMSChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorSMSChallengeResponseRequest(val *AuthenticatorSMSChallengeResponseRequest) *NullableAuthenticatorSMSChallengeResponseRequest {
	return &NullableAuthenticatorSMSChallengeResponseRequest{value: val, isSet: true}
}

func (v NullableAuthenticatorSMSChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorSMSChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


