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

// checks if the PatchedDenyStageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedDenyStageRequest{}

// PatchedDenyStageRequest DenyStage Serializer
type PatchedDenyStageRequest struct {
	Name *string `json:"name,omitempty"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	DenyMessage *string `json:"deny_message,omitempty"`
}

// NewPatchedDenyStageRequest instantiates a new PatchedDenyStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedDenyStageRequest() *PatchedDenyStageRequest {
	this := PatchedDenyStageRequest{}
	return &this
}

// NewPatchedDenyStageRequestWithDefaults instantiates a new PatchedDenyStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedDenyStageRequestWithDefaults() *PatchedDenyStageRequest {
	this := PatchedDenyStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedDenyStageRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDenyStageRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedDenyStageRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedDenyStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedDenyStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || IsNil(o.FlowSet) {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDenyStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || IsNil(o.FlowSet) {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedDenyStageRequest) HasFlowSet() bool {
	if o != nil && !IsNil(o.FlowSet) {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PatchedDenyStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetDenyMessage returns the DenyMessage field value if set, zero value otherwise.
func (o *PatchedDenyStageRequest) GetDenyMessage() string {
	if o == nil || IsNil(o.DenyMessage) {
		var ret string
		return ret
	}
	return *o.DenyMessage
}

// GetDenyMessageOk returns a tuple with the DenyMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDenyStageRequest) GetDenyMessageOk() (*string, bool) {
	if o == nil || IsNil(o.DenyMessage) {
		return nil, false
	}
	return o.DenyMessage, true
}

// HasDenyMessage returns a boolean if a field has been set.
func (o *PatchedDenyStageRequest) HasDenyMessage() bool {
	if o != nil && !IsNil(o.DenyMessage) {
		return true
	}

	return false
}

// SetDenyMessage gets a reference to the given string and assigns it to the DenyMessage field.
func (o *PatchedDenyStageRequest) SetDenyMessage(v string) {
	o.DenyMessage = &v
}

func (o PatchedDenyStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedDenyStageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.FlowSet) {
		toSerialize["flow_set"] = o.FlowSet
	}
	if !IsNil(o.DenyMessage) {
		toSerialize["deny_message"] = o.DenyMessage
	}
	return toSerialize, nil
}

type NullablePatchedDenyStageRequest struct {
	value *PatchedDenyStageRequest
	isSet bool
}

func (v NullablePatchedDenyStageRequest) Get() *PatchedDenyStageRequest {
	return v.value
}

func (v *NullablePatchedDenyStageRequest) Set(val *PatchedDenyStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedDenyStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedDenyStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedDenyStageRequest(val *PatchedDenyStageRequest) *NullablePatchedDenyStageRequest {
	return &NullablePatchedDenyStageRequest{value: val, isSet: true}
}

func (v NullablePatchedDenyStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedDenyStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


