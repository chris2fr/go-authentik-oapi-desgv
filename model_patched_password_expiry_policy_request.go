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

// checks if the PatchedPasswordExpiryPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedPasswordExpiryPolicyRequest{}

// PatchedPasswordExpiryPolicyRequest Password Expiry Policy Serializer
type PatchedPasswordExpiryPolicyRequest struct {
	Name *string `json:"name,omitempty"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool `json:"execution_logging,omitempty"`
	Days *int32 `json:"days,omitempty"`
	DenyOnly *bool `json:"deny_only,omitempty"`
}

// NewPatchedPasswordExpiryPolicyRequest instantiates a new PatchedPasswordExpiryPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedPasswordExpiryPolicyRequest() *PatchedPasswordExpiryPolicyRequest {
	this := PatchedPasswordExpiryPolicyRequest{}
	return &this
}

// NewPatchedPasswordExpiryPolicyRequestWithDefaults instantiates a new PatchedPasswordExpiryPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedPasswordExpiryPolicyRequestWithDefaults() *PatchedPasswordExpiryPolicyRequest {
	this := PatchedPasswordExpiryPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedPasswordExpiryPolicyRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordExpiryPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedPasswordExpiryPolicyRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedPasswordExpiryPolicyRequest) SetName(v string) {
	o.Name = &v
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *PatchedPasswordExpiryPolicyRequest) GetExecutionLogging() bool {
	if o == nil || IsNil(o.ExecutionLogging) {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordExpiryPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || IsNil(o.ExecutionLogging) {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *PatchedPasswordExpiryPolicyRequest) HasExecutionLogging() bool {
	if o != nil && !IsNil(o.ExecutionLogging) {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *PatchedPasswordExpiryPolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *PatchedPasswordExpiryPolicyRequest) GetDays() int32 {
	if o == nil || IsNil(o.Days) {
		var ret int32
		return ret
	}
	return *o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordExpiryPolicyRequest) GetDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.Days) {
		return nil, false
	}
	return o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *PatchedPasswordExpiryPolicyRequest) HasDays() bool {
	if o != nil && !IsNil(o.Days) {
		return true
	}

	return false
}

// SetDays gets a reference to the given int32 and assigns it to the Days field.
func (o *PatchedPasswordExpiryPolicyRequest) SetDays(v int32) {
	o.Days = &v
}

// GetDenyOnly returns the DenyOnly field value if set, zero value otherwise.
func (o *PatchedPasswordExpiryPolicyRequest) GetDenyOnly() bool {
	if o == nil || IsNil(o.DenyOnly) {
		var ret bool
		return ret
	}
	return *o.DenyOnly
}

// GetDenyOnlyOk returns a tuple with the DenyOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordExpiryPolicyRequest) GetDenyOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.DenyOnly) {
		return nil, false
	}
	return o.DenyOnly, true
}

// HasDenyOnly returns a boolean if a field has been set.
func (o *PatchedPasswordExpiryPolicyRequest) HasDenyOnly() bool {
	if o != nil && !IsNil(o.DenyOnly) {
		return true
	}

	return false
}

// SetDenyOnly gets a reference to the given bool and assigns it to the DenyOnly field.
func (o *PatchedPasswordExpiryPolicyRequest) SetDenyOnly(v bool) {
	o.DenyOnly = &v
}

func (o PatchedPasswordExpiryPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedPasswordExpiryPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ExecutionLogging) {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if !IsNil(o.Days) {
		toSerialize["days"] = o.Days
	}
	if !IsNil(o.DenyOnly) {
		toSerialize["deny_only"] = o.DenyOnly
	}
	return toSerialize, nil
}

type NullablePatchedPasswordExpiryPolicyRequest struct {
	value *PatchedPasswordExpiryPolicyRequest
	isSet bool
}

func (v NullablePatchedPasswordExpiryPolicyRequest) Get() *PatchedPasswordExpiryPolicyRequest {
	return v.value
}

func (v *NullablePatchedPasswordExpiryPolicyRequest) Set(val *PatchedPasswordExpiryPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedPasswordExpiryPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedPasswordExpiryPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedPasswordExpiryPolicyRequest(val *PatchedPasswordExpiryPolicyRequest) *NullablePatchedPasswordExpiryPolicyRequest {
	return &NullablePatchedPasswordExpiryPolicyRequest{value: val, isSet: true}
}

func (v NullablePatchedPasswordExpiryPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedPasswordExpiryPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


