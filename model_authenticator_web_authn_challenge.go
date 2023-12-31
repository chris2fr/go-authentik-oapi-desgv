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

// checks if the AuthenticatorWebAuthnChallenge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorWebAuthnChallenge{}

// AuthenticatorWebAuthnChallenge WebAuthn Challenge
type AuthenticatorWebAuthnChallenge struct {
	Type ChallengeChoices `json:"type"`
	FlowInfo *ContextualFlowInfo `json:"flow_info,omitempty"`
	Component *string `json:"component,omitempty"`
	ResponseErrors *map[string][]ErrorDetail `json:"response_errors,omitempty"`
	PendingUser string `json:"pending_user"`
	PendingUserAvatar string `json:"pending_user_avatar"`
	Registration map[string]interface{} `json:"registration"`
}

// NewAuthenticatorWebAuthnChallenge instantiates a new AuthenticatorWebAuthnChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorWebAuthnChallenge(type_ ChallengeChoices, pendingUser string, pendingUserAvatar string, registration map[string]interface{}) *AuthenticatorWebAuthnChallenge {
	this := AuthenticatorWebAuthnChallenge{}
	this.Type = type_
	var component string = "ak-stage-authenticator-webauthn"
	this.Component = &component
	this.PendingUser = pendingUser
	this.PendingUserAvatar = pendingUserAvatar
	this.Registration = registration
	return &this
}

// NewAuthenticatorWebAuthnChallengeWithDefaults instantiates a new AuthenticatorWebAuthnChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorWebAuthnChallengeWithDefaults() *AuthenticatorWebAuthnChallenge {
	this := AuthenticatorWebAuthnChallenge{}
	var component string = "ak-stage-authenticator-webauthn"
	this.Component = &component
	return &this
}

// GetType returns the Type field value
func (o *AuthenticatorWebAuthnChallenge) GetType() ChallengeChoices {
	if o == nil {
		var ret ChallengeChoices
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorWebAuthnChallenge) GetTypeOk() (*ChallengeChoices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuthenticatorWebAuthnChallenge) SetType(v ChallengeChoices) {
	o.Type = v
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *AuthenticatorWebAuthnChallenge) GetFlowInfo() ContextualFlowInfo {
	if o == nil || IsNil(o.FlowInfo) {
		var ret ContextualFlowInfo
		return ret
	}
	return *o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorWebAuthnChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool) {
	if o == nil || IsNil(o.FlowInfo) {
		return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *AuthenticatorWebAuthnChallenge) HasFlowInfo() bool {
	if o != nil && !IsNil(o.FlowInfo) {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given ContextualFlowInfo and assigns it to the FlowInfo field.
func (o *AuthenticatorWebAuthnChallenge) SetFlowInfo(v ContextualFlowInfo) {
	o.FlowInfo = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AuthenticatorWebAuthnChallenge) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorWebAuthnChallenge) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AuthenticatorWebAuthnChallenge) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AuthenticatorWebAuthnChallenge) SetComponent(v string) {
	o.Component = &v
}

// GetResponseErrors returns the ResponseErrors field value if set, zero value otherwise.
func (o *AuthenticatorWebAuthnChallenge) GetResponseErrors() map[string][]ErrorDetail {
	if o == nil || IsNil(o.ResponseErrors) {
		var ret map[string][]ErrorDetail
		return ret
	}
	return *o.ResponseErrors
}

// GetResponseErrorsOk returns a tuple with the ResponseErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorWebAuthnChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool) {
	if o == nil || IsNil(o.ResponseErrors) {
		return nil, false
	}
	return o.ResponseErrors, true
}

// HasResponseErrors returns a boolean if a field has been set.
func (o *AuthenticatorWebAuthnChallenge) HasResponseErrors() bool {
	if o != nil && !IsNil(o.ResponseErrors) {
		return true
	}

	return false
}

// SetResponseErrors gets a reference to the given map[string][]ErrorDetail and assigns it to the ResponseErrors field.
func (o *AuthenticatorWebAuthnChallenge) SetResponseErrors(v map[string][]ErrorDetail) {
	o.ResponseErrors = &v
}

// GetPendingUser returns the PendingUser field value
func (o *AuthenticatorWebAuthnChallenge) GetPendingUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUser
}

// GetPendingUserOk returns a tuple with the PendingUser field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorWebAuthnChallenge) GetPendingUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUser, true
}

// SetPendingUser sets field value
func (o *AuthenticatorWebAuthnChallenge) SetPendingUser(v string) {
	o.PendingUser = v
}

// GetPendingUserAvatar returns the PendingUserAvatar field value
func (o *AuthenticatorWebAuthnChallenge) GetPendingUserAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUserAvatar
}

// GetPendingUserAvatarOk returns a tuple with the PendingUserAvatar field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorWebAuthnChallenge) GetPendingUserAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUserAvatar, true
}

// SetPendingUserAvatar sets field value
func (o *AuthenticatorWebAuthnChallenge) SetPendingUserAvatar(v string) {
	o.PendingUserAvatar = v
}

// GetRegistration returns the Registration field value
func (o *AuthenticatorWebAuthnChallenge) GetRegistration() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Registration
}

// GetRegistrationOk returns a tuple with the Registration field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorWebAuthnChallenge) GetRegistrationOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Registration, true
}

// SetRegistration sets field value
func (o *AuthenticatorWebAuthnChallenge) SetRegistration(v map[string]interface{}) {
	o.Registration = v
}

func (o AuthenticatorWebAuthnChallenge) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorWebAuthnChallenge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.FlowInfo) {
		toSerialize["flow_info"] = o.FlowInfo
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if !IsNil(o.ResponseErrors) {
		toSerialize["response_errors"] = o.ResponseErrors
	}
	toSerialize["pending_user"] = o.PendingUser
	toSerialize["pending_user_avatar"] = o.PendingUserAvatar
	toSerialize["registration"] = o.Registration
	return toSerialize, nil
}

type NullableAuthenticatorWebAuthnChallenge struct {
	value *AuthenticatorWebAuthnChallenge
	isSet bool
}

func (v NullableAuthenticatorWebAuthnChallenge) Get() *AuthenticatorWebAuthnChallenge {
	return v.value
}

func (v *NullableAuthenticatorWebAuthnChallenge) Set(val *AuthenticatorWebAuthnChallenge) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorWebAuthnChallenge) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorWebAuthnChallenge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorWebAuthnChallenge(val *AuthenticatorWebAuthnChallenge) *NullableAuthenticatorWebAuthnChallenge {
	return &NullableAuthenticatorWebAuthnChallenge{value: val, isSet: true}
}

func (v NullableAuthenticatorWebAuthnChallenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorWebAuthnChallenge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


