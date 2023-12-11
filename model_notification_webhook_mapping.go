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

// checks if the NotificationWebhookMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationWebhookMapping{}

// NotificationWebhookMapping NotificationWebhookMapping Serializer
type NotificationWebhookMapping struct {
	Pk string `json:"pk"`
	Name string `json:"name"`
	Expression string `json:"expression"`
}

// NewNotificationWebhookMapping instantiates a new NotificationWebhookMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationWebhookMapping(pk string, name string, expression string) *NotificationWebhookMapping {
	this := NotificationWebhookMapping{}
	this.Pk = pk
	this.Name = name
	this.Expression = expression
	return &this
}

// NewNotificationWebhookMappingWithDefaults instantiates a new NotificationWebhookMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationWebhookMappingWithDefaults() *NotificationWebhookMapping {
	this := NotificationWebhookMapping{}
	return &this
}

// GetPk returns the Pk field value
func (o *NotificationWebhookMapping) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *NotificationWebhookMapping) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *NotificationWebhookMapping) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *NotificationWebhookMapping) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NotificationWebhookMapping) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NotificationWebhookMapping) SetName(v string) {
	o.Name = v
}

// GetExpression returns the Expression field value
func (o *NotificationWebhookMapping) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *NotificationWebhookMapping) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *NotificationWebhookMapping) SetExpression(v string) {
	o.Expression = v
}

func (o NotificationWebhookMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationWebhookMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pk"] = o.Pk
	toSerialize["name"] = o.Name
	toSerialize["expression"] = o.Expression
	return toSerialize, nil
}

type NullableNotificationWebhookMapping struct {
	value *NotificationWebhookMapping
	isSet bool
}

func (v NullableNotificationWebhookMapping) Get() *NotificationWebhookMapping {
	return v.value
}

func (v *NullableNotificationWebhookMapping) Set(val *NotificationWebhookMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationWebhookMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationWebhookMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationWebhookMapping(val *NotificationWebhookMapping) *NullableNotificationWebhookMapping {
	return &NullableNotificationWebhookMapping{value: val, isSet: true}
}

func (v NullableNotificationWebhookMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationWebhookMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


