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

// checks if the PatchedUserLoginStageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedUserLoginStageRequest{}

// PatchedUserLoginStageRequest UserLoginStage Serializer
type PatchedUserLoginStageRequest struct {
	Name *string `json:"name,omitempty"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	// Determines how long a session lasts. Default of 0 means that the sessions lasts until the browser is closed. (Format: hours=-1;minutes=-2;seconds=-3)
	SessionDuration *string `json:"session_duration,omitempty"`
	// Terminate all other sessions of the user logging in.
	TerminateOtherSessions *bool `json:"terminate_other_sessions,omitempty"`
	// Offset the session will be extended by when the user picks the remember me option. Default of 0 means that the remember me option will not be shown. (Format: hours=-1;minutes=-2;seconds=-3)
	RememberMeOffset *string `json:"remember_me_offset,omitempty"`
}

// NewPatchedUserLoginStageRequest instantiates a new PatchedUserLoginStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedUserLoginStageRequest() *PatchedUserLoginStageRequest {
	this := PatchedUserLoginStageRequest{}
	return &this
}

// NewPatchedUserLoginStageRequestWithDefaults instantiates a new PatchedUserLoginStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedUserLoginStageRequestWithDefaults() *PatchedUserLoginStageRequest {
	this := PatchedUserLoginStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedUserLoginStageRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserLoginStageRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedUserLoginStageRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedUserLoginStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedUserLoginStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || IsNil(o.FlowSet) {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserLoginStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || IsNil(o.FlowSet) {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedUserLoginStageRequest) HasFlowSet() bool {
	if o != nil && !IsNil(o.FlowSet) {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PatchedUserLoginStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetSessionDuration returns the SessionDuration field value if set, zero value otherwise.
func (o *PatchedUserLoginStageRequest) GetSessionDuration() string {
	if o == nil || IsNil(o.SessionDuration) {
		var ret string
		return ret
	}
	return *o.SessionDuration
}

// GetSessionDurationOk returns a tuple with the SessionDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserLoginStageRequest) GetSessionDurationOk() (*string, bool) {
	if o == nil || IsNil(o.SessionDuration) {
		return nil, false
	}
	return o.SessionDuration, true
}

// HasSessionDuration returns a boolean if a field has been set.
func (o *PatchedUserLoginStageRequest) HasSessionDuration() bool {
	if o != nil && !IsNil(o.SessionDuration) {
		return true
	}

	return false
}

// SetSessionDuration gets a reference to the given string and assigns it to the SessionDuration field.
func (o *PatchedUserLoginStageRequest) SetSessionDuration(v string) {
	o.SessionDuration = &v
}

// GetTerminateOtherSessions returns the TerminateOtherSessions field value if set, zero value otherwise.
func (o *PatchedUserLoginStageRequest) GetTerminateOtherSessions() bool {
	if o == nil || IsNil(o.TerminateOtherSessions) {
		var ret bool
		return ret
	}
	return *o.TerminateOtherSessions
}

// GetTerminateOtherSessionsOk returns a tuple with the TerminateOtherSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserLoginStageRequest) GetTerminateOtherSessionsOk() (*bool, bool) {
	if o == nil || IsNil(o.TerminateOtherSessions) {
		return nil, false
	}
	return o.TerminateOtherSessions, true
}

// HasTerminateOtherSessions returns a boolean if a field has been set.
func (o *PatchedUserLoginStageRequest) HasTerminateOtherSessions() bool {
	if o != nil && !IsNil(o.TerminateOtherSessions) {
		return true
	}

	return false
}

// SetTerminateOtherSessions gets a reference to the given bool and assigns it to the TerminateOtherSessions field.
func (o *PatchedUserLoginStageRequest) SetTerminateOtherSessions(v bool) {
	o.TerminateOtherSessions = &v
}

// GetRememberMeOffset returns the RememberMeOffset field value if set, zero value otherwise.
func (o *PatchedUserLoginStageRequest) GetRememberMeOffset() string {
	if o == nil || IsNil(o.RememberMeOffset) {
		var ret string
		return ret
	}
	return *o.RememberMeOffset
}

// GetRememberMeOffsetOk returns a tuple with the RememberMeOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserLoginStageRequest) GetRememberMeOffsetOk() (*string, bool) {
	if o == nil || IsNil(o.RememberMeOffset) {
		return nil, false
	}
	return o.RememberMeOffset, true
}

// HasRememberMeOffset returns a boolean if a field has been set.
func (o *PatchedUserLoginStageRequest) HasRememberMeOffset() bool {
	if o != nil && !IsNil(o.RememberMeOffset) {
		return true
	}

	return false
}

// SetRememberMeOffset gets a reference to the given string and assigns it to the RememberMeOffset field.
func (o *PatchedUserLoginStageRequest) SetRememberMeOffset(v string) {
	o.RememberMeOffset = &v
}

func (o PatchedUserLoginStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedUserLoginStageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.FlowSet) {
		toSerialize["flow_set"] = o.FlowSet
	}
	if !IsNil(o.SessionDuration) {
		toSerialize["session_duration"] = o.SessionDuration
	}
	if !IsNil(o.TerminateOtherSessions) {
		toSerialize["terminate_other_sessions"] = o.TerminateOtherSessions
	}
	if !IsNil(o.RememberMeOffset) {
		toSerialize["remember_me_offset"] = o.RememberMeOffset
	}
	return toSerialize, nil
}

type NullablePatchedUserLoginStageRequest struct {
	value *PatchedUserLoginStageRequest
	isSet bool
}

func (v NullablePatchedUserLoginStageRequest) Get() *PatchedUserLoginStageRequest {
	return v.value
}

func (v *NullablePatchedUserLoginStageRequest) Set(val *PatchedUserLoginStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedUserLoginStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedUserLoginStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedUserLoginStageRequest(val *PatchedUserLoginStageRequest) *NullablePatchedUserLoginStageRequest {
	return &NullablePatchedUserLoginStageRequest{value: val, isSet: true}
}

func (v NullablePatchedUserLoginStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedUserLoginStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

