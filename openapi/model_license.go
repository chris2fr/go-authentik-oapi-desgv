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

// checks if the License type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &License{}

// License License Serializer
type License struct {
	LicenseUuid string `json:"license_uuid"`
	Name string `json:"name"`
	Key string `json:"key"`
	Expiry time.Time `json:"expiry"`
	InternalUsers int32 `json:"internal_users"`
	ExternalUsers int32 `json:"external_users"`
}

// NewLicense instantiates a new License object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicense(licenseUuid string, name string, key string, expiry time.Time, internalUsers int32, externalUsers int32) *License {
	this := License{}
	this.LicenseUuid = licenseUuid
	this.Name = name
	this.Key = key
	this.Expiry = expiry
	this.InternalUsers = internalUsers
	this.ExternalUsers = externalUsers
	return &this
}

// NewLicenseWithDefaults instantiates a new License object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseWithDefaults() *License {
	this := License{}
	return &this
}

// GetLicenseUuid returns the LicenseUuid field value
func (o *License) GetLicenseUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LicenseUuid
}

// GetLicenseUuidOk returns a tuple with the LicenseUuid field value
// and a boolean to check if the value has been set.
func (o *License) GetLicenseUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseUuid, true
}

// SetLicenseUuid sets field value
func (o *License) SetLicenseUuid(v string) {
	o.LicenseUuid = v
}

// GetName returns the Name field value
func (o *License) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *License) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *License) SetName(v string) {
	o.Name = v
}

// GetKey returns the Key field value
func (o *License) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *License) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *License) SetKey(v string) {
	o.Key = v
}

// GetExpiry returns the Expiry field value
func (o *License) GetExpiry() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value
// and a boolean to check if the value has been set.
func (o *License) GetExpiryOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiry, true
}

// SetExpiry sets field value
func (o *License) SetExpiry(v time.Time) {
	o.Expiry = v
}

// GetInternalUsers returns the InternalUsers field value
func (o *License) GetInternalUsers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InternalUsers
}

// GetInternalUsersOk returns a tuple with the InternalUsers field value
// and a boolean to check if the value has been set.
func (o *License) GetInternalUsersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InternalUsers, true
}

// SetInternalUsers sets field value
func (o *License) SetInternalUsers(v int32) {
	o.InternalUsers = v
}

// GetExternalUsers returns the ExternalUsers field value
func (o *License) GetExternalUsers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExternalUsers
}

// GetExternalUsersOk returns a tuple with the ExternalUsers field value
// and a boolean to check if the value has been set.
func (o *License) GetExternalUsersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUsers, true
}

// SetExternalUsers sets field value
func (o *License) SetExternalUsers(v int32) {
	o.ExternalUsers = v
}

func (o License) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o License) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["license_uuid"] = o.LicenseUuid
	toSerialize["name"] = o.Name
	toSerialize["key"] = o.Key
	toSerialize["expiry"] = o.Expiry
	toSerialize["internal_users"] = o.InternalUsers
	toSerialize["external_users"] = o.ExternalUsers
	return toSerialize, nil
}

type NullableLicense struct {
	value *License
	isSet bool
}

func (v NullableLicense) Get() *License {
	return v.value
}

func (v *NullableLicense) Set(val *License) {
	v.value = val
	v.isSet = true
}

func (v NullableLicense) IsSet() bool {
	return v.isSet
}

func (v *NullableLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicense(val *License) *NullableLicense {
	return &NullableLicense{value: val, isSet: true}
}

func (v NullableLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

