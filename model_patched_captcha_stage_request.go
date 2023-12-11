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

// checks if the PatchedCaptchaStageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedCaptchaStageRequest{}

// PatchedCaptchaStageRequest CaptchaStage Serializer
type PatchedCaptchaStageRequest struct {
	Name *string `json:"name,omitempty"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	// Public key, acquired your captcha Provider.
	PublicKey *string `json:"public_key,omitempty"`
	// Private key, acquired your captcha Provider.
	PrivateKey *string `json:"private_key,omitempty"`
	JsUrl *string `json:"js_url,omitempty"`
	ApiUrl *string `json:"api_url,omitempty"`
}

// NewPatchedCaptchaStageRequest instantiates a new PatchedCaptchaStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedCaptchaStageRequest() *PatchedCaptchaStageRequest {
	this := PatchedCaptchaStageRequest{}
	return &this
}

// NewPatchedCaptchaStageRequestWithDefaults instantiates a new PatchedCaptchaStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedCaptchaStageRequestWithDefaults() *PatchedCaptchaStageRequest {
	this := PatchedCaptchaStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedCaptchaStageRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCaptchaStageRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedCaptchaStageRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedCaptchaStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedCaptchaStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || IsNil(o.FlowSet) {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCaptchaStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || IsNil(o.FlowSet) {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedCaptchaStageRequest) HasFlowSet() bool {
	if o != nil && !IsNil(o.FlowSet) {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PatchedCaptchaStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *PatchedCaptchaStageRequest) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCaptchaStageRequest) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *PatchedCaptchaStageRequest) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *PatchedCaptchaStageRequest) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *PatchedCaptchaStageRequest) GetPrivateKey() string {
	if o == nil || IsNil(o.PrivateKey) {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCaptchaStageRequest) GetPrivateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKey) {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *PatchedCaptchaStageRequest) HasPrivateKey() bool {
	if o != nil && !IsNil(o.PrivateKey) {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *PatchedCaptchaStageRequest) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetJsUrl returns the JsUrl field value if set, zero value otherwise.
func (o *PatchedCaptchaStageRequest) GetJsUrl() string {
	if o == nil || IsNil(o.JsUrl) {
		var ret string
		return ret
	}
	return *o.JsUrl
}

// GetJsUrlOk returns a tuple with the JsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCaptchaStageRequest) GetJsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.JsUrl) {
		return nil, false
	}
	return o.JsUrl, true
}

// HasJsUrl returns a boolean if a field has been set.
func (o *PatchedCaptchaStageRequest) HasJsUrl() bool {
	if o != nil && !IsNil(o.JsUrl) {
		return true
	}

	return false
}

// SetJsUrl gets a reference to the given string and assigns it to the JsUrl field.
func (o *PatchedCaptchaStageRequest) SetJsUrl(v string) {
	o.JsUrl = &v
}

// GetApiUrl returns the ApiUrl field value if set, zero value otherwise.
func (o *PatchedCaptchaStageRequest) GetApiUrl() string {
	if o == nil || IsNil(o.ApiUrl) {
		var ret string
		return ret
	}
	return *o.ApiUrl
}

// GetApiUrlOk returns a tuple with the ApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCaptchaStageRequest) GetApiUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ApiUrl) {
		return nil, false
	}
	return o.ApiUrl, true
}

// HasApiUrl returns a boolean if a field has been set.
func (o *PatchedCaptchaStageRequest) HasApiUrl() bool {
	if o != nil && !IsNil(o.ApiUrl) {
		return true
	}

	return false
}

// SetApiUrl gets a reference to the given string and assigns it to the ApiUrl field.
func (o *PatchedCaptchaStageRequest) SetApiUrl(v string) {
	o.ApiUrl = &v
}

func (o PatchedCaptchaStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedCaptchaStageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.FlowSet) {
		toSerialize["flow_set"] = o.FlowSet
	}
	if !IsNil(o.PublicKey) {
		toSerialize["public_key"] = o.PublicKey
	}
	if !IsNil(o.PrivateKey) {
		toSerialize["private_key"] = o.PrivateKey
	}
	if !IsNil(o.JsUrl) {
		toSerialize["js_url"] = o.JsUrl
	}
	if !IsNil(o.ApiUrl) {
		toSerialize["api_url"] = o.ApiUrl
	}
	return toSerialize, nil
}

type NullablePatchedCaptchaStageRequest struct {
	value *PatchedCaptchaStageRequest
	isSet bool
}

func (v NullablePatchedCaptchaStageRequest) Get() *PatchedCaptchaStageRequest {
	return v.value
}

func (v *NullablePatchedCaptchaStageRequest) Set(val *PatchedCaptchaStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedCaptchaStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedCaptchaStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedCaptchaStageRequest(val *PatchedCaptchaStageRequest) *NullablePatchedCaptchaStageRequest {
	return &NullablePatchedCaptchaStageRequest{value: val, isSet: true}
}

func (v NullablePatchedCaptchaStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedCaptchaStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

