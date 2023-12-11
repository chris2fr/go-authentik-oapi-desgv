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

// checks if the PasswordStage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordStage{}

// PasswordStage PasswordStage Serializer
type PasswordStage struct {
	Pk string `json:"pk"`
	Name string `json:"name"`
	// Get object type so that we know how to edit the object
	Component string `json:"component"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName string `json:"meta_model_name"`
	FlowSet []FlowSet `json:"flow_set,omitempty"`
	// Selection of backends to test the password against.
	Backends []BackendsEnum `json:"backends"`
	// Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage.
	ConfigureFlow NullableString `json:"configure_flow,omitempty"`
	// How many attempts a user has before the flow is canceled. To lock the user out, use a reputation policy and a user_write stage.
	FailedAttemptsBeforeCancel *int32 `json:"failed_attempts_before_cancel,omitempty"`
}

// NewPasswordStage instantiates a new PasswordStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, backends []BackendsEnum) *PasswordStage {
	this := PasswordStage{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.Backends = backends
	return &this
}

// NewPasswordStageWithDefaults instantiates a new PasswordStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordStageWithDefaults() *PasswordStage {
	this := PasswordStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *PasswordStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *PasswordStage) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *PasswordStage) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *PasswordStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PasswordStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PasswordStage) SetName(v string) {
	o.Name = v
}

// GetComponent returns the Component field value
func (o *PasswordStage) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *PasswordStage) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *PasswordStage) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *PasswordStage) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *PasswordStage) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *PasswordStage) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *PasswordStage) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *PasswordStage) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *PasswordStage) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *PasswordStage) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *PasswordStage) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *PasswordStage) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PasswordStage) GetFlowSet() []FlowSet {
	if o == nil || IsNil(o.FlowSet) {
		var ret []FlowSet
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordStage) GetFlowSetOk() ([]FlowSet, bool) {
	if o == nil || IsNil(o.FlowSet) {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PasswordStage) HasFlowSet() bool {
	if o != nil && !IsNil(o.FlowSet) {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSet and assigns it to the FlowSet field.
func (o *PasswordStage) SetFlowSet(v []FlowSet) {
	o.FlowSet = v
}

// GetBackends returns the Backends field value
func (o *PasswordStage) GetBackends() []BackendsEnum {
	if o == nil {
		var ret []BackendsEnum
		return ret
	}

	return o.Backends
}

// GetBackendsOk returns a tuple with the Backends field value
// and a boolean to check if the value has been set.
func (o *PasswordStage) GetBackendsOk() ([]BackendsEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Backends, true
}

// SetBackends sets field value
func (o *PasswordStage) SetBackends(v []BackendsEnum) {
	o.Backends = v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PasswordStage) GetConfigureFlow() string {
	if o == nil || IsNil(o.ConfigureFlow.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordStage) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *PasswordStage) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *PasswordStage) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}
// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *PasswordStage) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *PasswordStage) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetFailedAttemptsBeforeCancel returns the FailedAttemptsBeforeCancel field value if set, zero value otherwise.
func (o *PasswordStage) GetFailedAttemptsBeforeCancel() int32 {
	if o == nil || IsNil(o.FailedAttemptsBeforeCancel) {
		var ret int32
		return ret
	}
	return *o.FailedAttemptsBeforeCancel
}

// GetFailedAttemptsBeforeCancelOk returns a tuple with the FailedAttemptsBeforeCancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordStage) GetFailedAttemptsBeforeCancelOk() (*int32, bool) {
	if o == nil || IsNil(o.FailedAttemptsBeforeCancel) {
		return nil, false
	}
	return o.FailedAttemptsBeforeCancel, true
}

// HasFailedAttemptsBeforeCancel returns a boolean if a field has been set.
func (o *PasswordStage) HasFailedAttemptsBeforeCancel() bool {
	if o != nil && !IsNil(o.FailedAttemptsBeforeCancel) {
		return true
	}

	return false
}

// SetFailedAttemptsBeforeCancel gets a reference to the given int32 and assigns it to the FailedAttemptsBeforeCancel field.
func (o *PasswordStage) SetFailedAttemptsBeforeCancel(v int32) {
	o.FailedAttemptsBeforeCancel = &v
}

func (o PasswordStage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordStage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pk"] = o.Pk
	toSerialize["name"] = o.Name
	toSerialize["component"] = o.Component
	toSerialize["verbose_name"] = o.VerboseName
	toSerialize["verbose_name_plural"] = o.VerboseNamePlural
	toSerialize["meta_model_name"] = o.MetaModelName
	if !IsNil(o.FlowSet) {
		toSerialize["flow_set"] = o.FlowSet
	}
	toSerialize["backends"] = o.Backends
	if o.ConfigureFlow.IsSet() {
		toSerialize["configure_flow"] = o.ConfigureFlow.Get()
	}
	if !IsNil(o.FailedAttemptsBeforeCancel) {
		toSerialize["failed_attempts_before_cancel"] = o.FailedAttemptsBeforeCancel
	}
	return toSerialize, nil
}

type NullablePasswordStage struct {
	value *PasswordStage
	isSet bool
}

func (v NullablePasswordStage) Get() *PasswordStage {
	return v.value
}

func (v *NullablePasswordStage) Set(val *PasswordStage) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordStage) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordStage(val *PasswordStage) *NullablePasswordStage {
	return &NullablePasswordStage{value: val, isSet: true}
}

func (v NullablePasswordStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


