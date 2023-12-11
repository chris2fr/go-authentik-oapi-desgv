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
	"time"
)

// checks if the Invitation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Invitation{}

// Invitation Invitation Serializer
type Invitation struct {
	Pk string `json:"pk"`
	Name string `json:"name"`
	Expires *time.Time `json:"expires,omitempty"`
	FixedData map[string]interface{} `json:"fixed_data,omitempty"`
	CreatedBy GroupMember `json:"created_by"`
	// When enabled, the invitation will be deleted after usage.
	SingleUse *bool `json:"single_use,omitempty"`
	// When set, only the configured flow can use this invitation.
	Flow NullableString `json:"flow,omitempty"`
	FlowObj Flow `json:"flow_obj"`
}

// NewInvitation instantiates a new Invitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitation(pk string, name string, createdBy GroupMember, flowObj Flow) *Invitation {
	this := Invitation{}
	this.Pk = pk
	this.Name = name
	this.CreatedBy = createdBy
	this.FlowObj = flowObj
	return &this
}

// NewInvitationWithDefaults instantiates a new Invitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationWithDefaults() *Invitation {
	this := Invitation{}
	return &this
}

// GetPk returns the Pk field value
func (o *Invitation) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *Invitation) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *Invitation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Invitation) SetName(v string) {
	o.Name = v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *Invitation) GetExpires() time.Time {
	if o == nil || IsNil(o.Expires) {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetExpiresOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expires) {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *Invitation) HasExpires() bool {
	if o != nil && !IsNil(o.Expires) {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *Invitation) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetFixedData returns the FixedData field value if set, zero value otherwise.
func (o *Invitation) GetFixedData() map[string]interface{} {
	if o == nil || IsNil(o.FixedData) {
		var ret map[string]interface{}
		return ret
	}
	return o.FixedData
}

// GetFixedDataOk returns a tuple with the FixedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetFixedDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.FixedData) {
		return map[string]interface{}{}, false
	}
	return o.FixedData, true
}

// HasFixedData returns a boolean if a field has been set.
func (o *Invitation) HasFixedData() bool {
	if o != nil && !IsNil(o.FixedData) {
		return true
	}

	return false
}

// SetFixedData gets a reference to the given map[string]interface{} and assigns it to the FixedData field.
func (o *Invitation) SetFixedData(v map[string]interface{}) {
	o.FixedData = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *Invitation) GetCreatedBy() GroupMember {
	if o == nil {
		var ret GroupMember
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetCreatedByOk() (*GroupMember, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *Invitation) SetCreatedBy(v GroupMember) {
	o.CreatedBy = v
}

// GetSingleUse returns the SingleUse field value if set, zero value otherwise.
func (o *Invitation) GetSingleUse() bool {
	if o == nil || IsNil(o.SingleUse) {
		var ret bool
		return ret
	}
	return *o.SingleUse
}

// GetSingleUseOk returns a tuple with the SingleUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetSingleUseOk() (*bool, bool) {
	if o == nil || IsNil(o.SingleUse) {
		return nil, false
	}
	return o.SingleUse, true
}

// HasSingleUse returns a boolean if a field has been set.
func (o *Invitation) HasSingleUse() bool {
	if o != nil && !IsNil(o.SingleUse) {
		return true
	}

	return false
}

// SetSingleUse gets a reference to the given bool and assigns it to the SingleUse field.
func (o *Invitation) SetSingleUse(v bool) {
	o.SingleUse = &v
}

// GetFlow returns the Flow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Invitation) GetFlow() string {
	if o == nil || IsNil(o.Flow.Get()) {
		var ret string
		return ret
	}
	return *o.Flow.Get()
}

// GetFlowOk returns a tuple with the Flow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Invitation) GetFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Flow.Get(), o.Flow.IsSet()
}

// HasFlow returns a boolean if a field has been set.
func (o *Invitation) HasFlow() bool {
	if o != nil && o.Flow.IsSet() {
		return true
	}

	return false
}

// SetFlow gets a reference to the given NullableString and assigns it to the Flow field.
func (o *Invitation) SetFlow(v string) {
	o.Flow.Set(&v)
}
// SetFlowNil sets the value for Flow to be an explicit nil
func (o *Invitation) SetFlowNil() {
	o.Flow.Set(nil)
}

// UnsetFlow ensures that no value is present for Flow, not even an explicit nil
func (o *Invitation) UnsetFlow() {
	o.Flow.Unset()
}

// GetFlowObj returns the FlowObj field value
func (o *Invitation) GetFlowObj() Flow {
	if o == nil {
		var ret Flow
		return ret
	}

	return o.FlowObj
}

// GetFlowObjOk returns a tuple with the FlowObj field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetFlowObjOk() (*Flow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlowObj, true
}

// SetFlowObj sets field value
func (o *Invitation) SetFlowObj(v Flow) {
	o.FlowObj = v
}

func (o Invitation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Invitation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pk"] = o.Pk
	toSerialize["name"] = o.Name
	if !IsNil(o.Expires) {
		toSerialize["expires"] = o.Expires
	}
	if !IsNil(o.FixedData) {
		toSerialize["fixed_data"] = o.FixedData
	}
	toSerialize["created_by"] = o.CreatedBy
	if !IsNil(o.SingleUse) {
		toSerialize["single_use"] = o.SingleUse
	}
	if o.Flow.IsSet() {
		toSerialize["flow"] = o.Flow.Get()
	}
	toSerialize["flow_obj"] = o.FlowObj
	return toSerialize, nil
}

type NullableInvitation struct {
	value *Invitation
	isSet bool
}

func (v NullableInvitation) Get() *Invitation {
	return v.value
}

func (v *NullableInvitation) Set(val *Invitation) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitation(val *Invitation) *NullableInvitation {
	return &NullableInvitation{value: val, isSet: true}
}

func (v NullableInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

