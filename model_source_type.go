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

// checks if the SourceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceType{}

// SourceType Serializer for SourceType
type SourceType struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	UrlsCustomizable bool `json:"urls_customizable"`
	RequestTokenUrl NullableString `json:"request_token_url"`
	AuthorizationUrl NullableString `json:"authorization_url"`
	AccessTokenUrl NullableString `json:"access_token_url"`
	ProfileUrl NullableString `json:"profile_url"`
	OidcWellKnownUrl NullableString `json:"oidc_well_known_url"`
	OidcJwksUrl NullableString `json:"oidc_jwks_url"`
}

// NewSourceType instantiates a new SourceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceType(name string, slug string, urlsCustomizable bool, requestTokenUrl NullableString, authorizationUrl NullableString, accessTokenUrl NullableString, profileUrl NullableString, oidcWellKnownUrl NullableString, oidcJwksUrl NullableString) *SourceType {
	this := SourceType{}
	this.Name = name
	this.Slug = slug
	this.UrlsCustomizable = urlsCustomizable
	this.RequestTokenUrl = requestTokenUrl
	this.AuthorizationUrl = authorizationUrl
	this.AccessTokenUrl = accessTokenUrl
	this.ProfileUrl = profileUrl
	this.OidcWellKnownUrl = oidcWellKnownUrl
	this.OidcJwksUrl = oidcJwksUrl
	return &this
}

// NewSourceTypeWithDefaults instantiates a new SourceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceTypeWithDefaults() *SourceType {
	this := SourceType{}
	return &this
}

// GetName returns the Name field value
func (o *SourceType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SourceType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SourceType) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *SourceType) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *SourceType) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *SourceType) SetSlug(v string) {
	o.Slug = v
}

// GetUrlsCustomizable returns the UrlsCustomizable field value
func (o *SourceType) GetUrlsCustomizable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UrlsCustomizable
}

// GetUrlsCustomizableOk returns a tuple with the UrlsCustomizable field value
// and a boolean to check if the value has been set.
func (o *SourceType) GetUrlsCustomizableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UrlsCustomizable, true
}

// SetUrlsCustomizable sets field value
func (o *SourceType) SetUrlsCustomizable(v bool) {
	o.UrlsCustomizable = v
}

// GetRequestTokenUrl returns the RequestTokenUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SourceType) GetRequestTokenUrl() string {
	if o == nil || o.RequestTokenUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.RequestTokenUrl.Get()
}

// GetRequestTokenUrlOk returns a tuple with the RequestTokenUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceType) GetRequestTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestTokenUrl.Get(), o.RequestTokenUrl.IsSet()
}

// SetRequestTokenUrl sets field value
func (o *SourceType) SetRequestTokenUrl(v string) {
	o.RequestTokenUrl.Set(&v)
}

// GetAuthorizationUrl returns the AuthorizationUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SourceType) GetAuthorizationUrl() string {
	if o == nil || o.AuthorizationUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthorizationUrl.Get()
}

// GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceType) GetAuthorizationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizationUrl.Get(), o.AuthorizationUrl.IsSet()
}

// SetAuthorizationUrl sets field value
func (o *SourceType) SetAuthorizationUrl(v string) {
	o.AuthorizationUrl.Set(&v)
}

// GetAccessTokenUrl returns the AccessTokenUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SourceType) GetAccessTokenUrl() string {
	if o == nil || o.AccessTokenUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccessTokenUrl.Get()
}

// GetAccessTokenUrlOk returns a tuple with the AccessTokenUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceType) GetAccessTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessTokenUrl.Get(), o.AccessTokenUrl.IsSet()
}

// SetAccessTokenUrl sets field value
func (o *SourceType) SetAccessTokenUrl(v string) {
	o.AccessTokenUrl.Set(&v)
}

// GetProfileUrl returns the ProfileUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SourceType) GetProfileUrl() string {
	if o == nil || o.ProfileUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProfileUrl.Get()
}

// GetProfileUrlOk returns a tuple with the ProfileUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceType) GetProfileUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileUrl.Get(), o.ProfileUrl.IsSet()
}

// SetProfileUrl sets field value
func (o *SourceType) SetProfileUrl(v string) {
	o.ProfileUrl.Set(&v)
}

// GetOidcWellKnownUrl returns the OidcWellKnownUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SourceType) GetOidcWellKnownUrl() string {
	if o == nil || o.OidcWellKnownUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.OidcWellKnownUrl.Get()
}

// GetOidcWellKnownUrlOk returns a tuple with the OidcWellKnownUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceType) GetOidcWellKnownUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OidcWellKnownUrl.Get(), o.OidcWellKnownUrl.IsSet()
}

// SetOidcWellKnownUrl sets field value
func (o *SourceType) SetOidcWellKnownUrl(v string) {
	o.OidcWellKnownUrl.Set(&v)
}

// GetOidcJwksUrl returns the OidcJwksUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SourceType) GetOidcJwksUrl() string {
	if o == nil || o.OidcJwksUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.OidcJwksUrl.Get()
}

// GetOidcJwksUrlOk returns a tuple with the OidcJwksUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceType) GetOidcJwksUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OidcJwksUrl.Get(), o.OidcJwksUrl.IsSet()
}

// SetOidcJwksUrl sets field value
func (o *SourceType) SetOidcJwksUrl(v string) {
	o.OidcJwksUrl.Set(&v)
}

func (o SourceType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	toSerialize["urls_customizable"] = o.UrlsCustomizable
	toSerialize["request_token_url"] = o.RequestTokenUrl.Get()
	toSerialize["authorization_url"] = o.AuthorizationUrl.Get()
	toSerialize["access_token_url"] = o.AccessTokenUrl.Get()
	toSerialize["profile_url"] = o.ProfileUrl.Get()
	toSerialize["oidc_well_known_url"] = o.OidcWellKnownUrl.Get()
	toSerialize["oidc_jwks_url"] = o.OidcJwksUrl.Get()
	return toSerialize, nil
}

type NullableSourceType struct {
	value *SourceType
	isSet bool
}

func (v NullableSourceType) Get() *SourceType {
	return v.value
}

func (v *NullableSourceType) Set(val *SourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceType(val *SourceType) *NullableSourceType {
	return &NullableSourceType{value: val, isSet: true}
}

func (v NullableSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


