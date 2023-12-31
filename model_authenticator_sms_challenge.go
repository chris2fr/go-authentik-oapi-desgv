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

// checks if the AuthenticatorSMSChallenge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorSMSChallenge{}

// AuthenticatorSMSChallenge SMS Setup challenge
type AuthenticatorSMSChallenge struct {
	Type ChallengeChoices `json:"type"`
	FlowInfo *ContextualFlowInfo `json:"flow_info,omitempty"`
	Component *string `json:"component,omitempty"`
	ResponseErrors *map[string][]ErrorDetail `json:"response_errors,omitempty"`
	PendingUser string `json:"pending_user"`
	PendingUserAvatar string `json:"pending_user_avatar"`
	PhoneNumberRequired *bool `json:"phone_number_required,omitempty"`
}

// NewAuthenticatorSMSChallenge instantiates a new AuthenticatorSMSChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorSMSChallenge(type_ ChallengeChoices, pendingUser string, pendingUserAvatar string) *AuthenticatorSMSChallenge {
	this := AuthenticatorSMSChallenge{}
	this.Type = type_
	var component string = "ak-stage-authenticator-sms"
	this.Component = &component
	this.PendingUser = pendingUser
	this.PendingUserAvatar = pendingUserAvatar
	var phoneNumberRequired bool = true
	this.PhoneNumberRequired = &phoneNumberRequired
	return &this
}

// NewAuthenticatorSMSChallengeWithDefaults instantiates a new AuthenticatorSMSChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorSMSChallengeWithDefaults() *AuthenticatorSMSChallenge {
	this := AuthenticatorSMSChallenge{}
	var component string = "ak-stage-authenticator-sms"
	this.Component = &component
	var phoneNumberRequired bool = true
	this.PhoneNumberRequired = &phoneNumberRequired
	return &this
}

// GetType returns the Type field value
func (o *AuthenticatorSMSChallenge) GetType() ChallengeChoices {
	if o == nil {
		var ret ChallengeChoices
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSChallenge) GetTypeOk() (*ChallengeChoices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuthenticatorSMSChallenge) SetType(v ChallengeChoices) {
	o.Type = v
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *AuthenticatorSMSChallenge) GetFlowInfo() ContextualFlowInfo {
	if o == nil || IsNil(o.FlowInfo) {
		var ret ContextualFlowInfo
		return ret
	}
	return *o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool) {
	if o == nil || IsNil(o.FlowInfo) {
		return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *AuthenticatorSMSChallenge) HasFlowInfo() bool {
	if o != nil && !IsNil(o.FlowInfo) {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given ContextualFlowInfo and assigns it to the FlowInfo field.
func (o *AuthenticatorSMSChallenge) SetFlowInfo(v ContextualFlowInfo) {
	o.FlowInfo = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AuthenticatorSMSChallenge) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSChallenge) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AuthenticatorSMSChallenge) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AuthenticatorSMSChallenge) SetComponent(v string) {
	o.Component = &v
}

// GetResponseErrors returns the ResponseErrors field value if set, zero value otherwise.
func (o *AuthenticatorSMSChallenge) GetResponseErrors() map[string][]ErrorDetail {
	if o == nil || IsNil(o.ResponseErrors) {
		var ret map[string][]ErrorDetail
		return ret
	}
	return *o.ResponseErrors
}

// GetResponseErrorsOk returns a tuple with the ResponseErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool) {
	if o == nil || IsNil(o.ResponseErrors) {
		return nil, false
	}
	return o.ResponseErrors, true
}

// HasResponseErrors returns a boolean if a field has been set.
func (o *AuthenticatorSMSChallenge) HasResponseErrors() bool {
	if o != nil && !IsNil(o.ResponseErrors) {
		return true
	}

	return false
}

// SetResponseErrors gets a reference to the given map[string][]ErrorDetail and assigns it to the ResponseErrors field.
func (o *AuthenticatorSMSChallenge) SetResponseErrors(v map[string][]ErrorDetail) {
	o.ResponseErrors = &v
}

// GetPendingUser returns the PendingUser field value
func (o *AuthenticatorSMSChallenge) GetPendingUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUser
}

// GetPendingUserOk returns a tuple with the PendingUser field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSChallenge) GetPendingUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUser, true
}

// SetPendingUser sets field value
func (o *AuthenticatorSMSChallenge) SetPendingUser(v string) {
	o.PendingUser = v
}

// GetPendingUserAvatar returns the PendingUserAvatar field value
func (o *AuthenticatorSMSChallenge) GetPendingUserAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUserAvatar
}

// GetPendingUserAvatarOk returns a tuple with the PendingUserAvatar field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSChallenge) GetPendingUserAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUserAvatar, true
}

// SetPendingUserAvatar sets field value
func (o *AuthenticatorSMSChallenge) SetPendingUserAvatar(v string) {
	o.PendingUserAvatar = v
}

// GetPhoneNumberRequired returns the PhoneNumberRequired field value if set, zero value otherwise.
func (o *AuthenticatorSMSChallenge) GetPhoneNumberRequired() bool {
	if o == nil || IsNil(o.PhoneNumberRequired) {
		var ret bool
		return ret
	}
	return *o.PhoneNumberRequired
}

// GetPhoneNumberRequiredOk returns a tuple with the PhoneNumberRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSChallenge) GetPhoneNumberRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.PhoneNumberRequired) {
		return nil, false
	}
	return o.PhoneNumberRequired, true
}

// HasPhoneNumberRequired returns a boolean if a field has been set.
func (o *AuthenticatorSMSChallenge) HasPhoneNumberRequired() bool {
	if o != nil && !IsNil(o.PhoneNumberRequired) {
		return true
	}

	return false
}

// SetPhoneNumberRequired gets a reference to the given bool and assigns it to the PhoneNumberRequired field.
func (o *AuthenticatorSMSChallenge) SetPhoneNumberRequired(v bool) {
	o.PhoneNumberRequired = &v
}

func (o AuthenticatorSMSChallenge) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorSMSChallenge) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.PhoneNumberRequired) {
		toSerialize["phone_number_required"] = o.PhoneNumberRequired
	}
	return toSerialize, nil
}

type NullableAuthenticatorSMSChallenge struct {
	value *AuthenticatorSMSChallenge
	isSet bool
}

func (v NullableAuthenticatorSMSChallenge) Get() *AuthenticatorSMSChallenge {
	return v.value
}

func (v *NullableAuthenticatorSMSChallenge) Set(val *AuthenticatorSMSChallenge) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorSMSChallenge) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorSMSChallenge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorSMSChallenge(val *AuthenticatorSMSChallenge) *NullableAuthenticatorSMSChallenge {
	return &NullableAuthenticatorSMSChallenge{value: val, isSet: true}
}

func (v NullableAuthenticatorSMSChallenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorSMSChallenge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


