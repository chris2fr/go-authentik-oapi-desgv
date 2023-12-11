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

// checks if the OAuth2ProviderSetupURLs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2ProviderSetupURLs{}

// OAuth2ProviderSetupURLs OAuth2 Provider Metadata serializer
type OAuth2ProviderSetupURLs struct {
	Issuer string `json:"issuer"`
	Authorize string `json:"authorize"`
	Token string `json:"token"`
	UserInfo string `json:"user_info"`
	ProviderInfo string `json:"provider_info"`
	Logout string `json:"logout"`
	Jwks string `json:"jwks"`
}

// NewOAuth2ProviderSetupURLs instantiates a new OAuth2ProviderSetupURLs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ProviderSetupURLs(issuer string, authorize string, token string, userInfo string, providerInfo string, logout string, jwks string) *OAuth2ProviderSetupURLs {
	this := OAuth2ProviderSetupURLs{}
	this.Issuer = issuer
	this.Authorize = authorize
	this.Token = token
	this.UserInfo = userInfo
	this.ProviderInfo = providerInfo
	this.Logout = logout
	this.Jwks = jwks
	return &this
}

// NewOAuth2ProviderSetupURLsWithDefaults instantiates a new OAuth2ProviderSetupURLs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ProviderSetupURLsWithDefaults() *OAuth2ProviderSetupURLs {
	this := OAuth2ProviderSetupURLs{}
	return &this
}

// GetIssuer returns the Issuer field value
func (o *OAuth2ProviderSetupURLs) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderSetupURLs) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *OAuth2ProviderSetupURLs) SetIssuer(v string) {
	o.Issuer = v
}

// GetAuthorize returns the Authorize field value
func (o *OAuth2ProviderSetupURLs) GetAuthorize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Authorize
}

// GetAuthorizeOk returns a tuple with the Authorize field value
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderSetupURLs) GetAuthorizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authorize, true
}

// SetAuthorize sets field value
func (o *OAuth2ProviderSetupURLs) SetAuthorize(v string) {
	o.Authorize = v
}

// GetToken returns the Token field value
func (o *OAuth2ProviderSetupURLs) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderSetupURLs) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *OAuth2ProviderSetupURLs) SetToken(v string) {
	o.Token = v
}

// GetUserInfo returns the UserInfo field value
func (o *OAuth2ProviderSetupURLs) GetUserInfo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserInfo
}

// GetUserInfoOk returns a tuple with the UserInfo field value
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderSetupURLs) GetUserInfoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserInfo, true
}

// SetUserInfo sets field value
func (o *OAuth2ProviderSetupURLs) SetUserInfo(v string) {
	o.UserInfo = v
}

// GetProviderInfo returns the ProviderInfo field value
func (o *OAuth2ProviderSetupURLs) GetProviderInfo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderInfo
}

// GetProviderInfoOk returns a tuple with the ProviderInfo field value
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderSetupURLs) GetProviderInfoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderInfo, true
}

// SetProviderInfo sets field value
func (o *OAuth2ProviderSetupURLs) SetProviderInfo(v string) {
	o.ProviderInfo = v
}

// GetLogout returns the Logout field value
func (o *OAuth2ProviderSetupURLs) GetLogout() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Logout
}

// GetLogoutOk returns a tuple with the Logout field value
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderSetupURLs) GetLogoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logout, true
}

// SetLogout sets field value
func (o *OAuth2ProviderSetupURLs) SetLogout(v string) {
	o.Logout = v
}

// GetJwks returns the Jwks field value
func (o *OAuth2ProviderSetupURLs) GetJwks() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Jwks
}

// GetJwksOk returns a tuple with the Jwks field value
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderSetupURLs) GetJwksOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Jwks, true
}

// SetJwks sets field value
func (o *OAuth2ProviderSetupURLs) SetJwks(v string) {
	o.Jwks = v
}

func (o OAuth2ProviderSetupURLs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2ProviderSetupURLs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["issuer"] = o.Issuer
	toSerialize["authorize"] = o.Authorize
	toSerialize["token"] = o.Token
	toSerialize["user_info"] = o.UserInfo
	toSerialize["provider_info"] = o.ProviderInfo
	toSerialize["logout"] = o.Logout
	toSerialize["jwks"] = o.Jwks
	return toSerialize, nil
}

type NullableOAuth2ProviderSetupURLs struct {
	value *OAuth2ProviderSetupURLs
	isSet bool
}

func (v NullableOAuth2ProviderSetupURLs) Get() *OAuth2ProviderSetupURLs {
	return v.value
}

func (v *NullableOAuth2ProviderSetupURLs) Set(val *OAuth2ProviderSetupURLs) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ProviderSetupURLs) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ProviderSetupURLs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ProviderSetupURLs(val *OAuth2ProviderSetupURLs) *NullableOAuth2ProviderSetupURLs {
	return &NullableOAuth2ProviderSetupURLs{value: val, isSet: true}
}

func (v NullableOAuth2ProviderSetupURLs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ProviderSetupURLs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


