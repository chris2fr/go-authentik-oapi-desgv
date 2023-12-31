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

// checks if the AuthenticatedSessionGeoIp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatedSessionGeoIp{}

// AuthenticatedSessionGeoIp Get parsed user agent
type AuthenticatedSessionGeoIp struct {
	Continent string `json:"continent"`
	Country string `json:"country"`
	Lat float64 `json:"lat"`
	Long float64 `json:"long"`
	City string `json:"city"`
}

// NewAuthenticatedSessionGeoIp instantiates a new AuthenticatedSessionGeoIp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatedSessionGeoIp(continent string, country string, lat float64, long float64, city string) *AuthenticatedSessionGeoIp {
	this := AuthenticatedSessionGeoIp{}
	this.Continent = continent
	this.Country = country
	this.Lat = lat
	this.Long = long
	this.City = city
	return &this
}

// NewAuthenticatedSessionGeoIpWithDefaults instantiates a new AuthenticatedSessionGeoIp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatedSessionGeoIpWithDefaults() *AuthenticatedSessionGeoIp {
	this := AuthenticatedSessionGeoIp{}
	return &this
}

// GetContinent returns the Continent field value
func (o *AuthenticatedSessionGeoIp) GetContinent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Continent
}

// GetContinentOk returns a tuple with the Continent field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionGeoIp) GetContinentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Continent, true
}

// SetContinent sets field value
func (o *AuthenticatedSessionGeoIp) SetContinent(v string) {
	o.Continent = v
}

// GetCountry returns the Country field value
func (o *AuthenticatedSessionGeoIp) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionGeoIp) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *AuthenticatedSessionGeoIp) SetCountry(v string) {
	o.Country = v
}

// GetLat returns the Lat field value
func (o *AuthenticatedSessionGeoIp) GetLat() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Lat
}

// GetLatOk returns a tuple with the Lat field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionGeoIp) GetLatOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lat, true
}

// SetLat sets field value
func (o *AuthenticatedSessionGeoIp) SetLat(v float64) {
	o.Lat = v
}

// GetLong returns the Long field value
func (o *AuthenticatedSessionGeoIp) GetLong() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Long
}

// GetLongOk returns a tuple with the Long field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionGeoIp) GetLongOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Long, true
}

// SetLong sets field value
func (o *AuthenticatedSessionGeoIp) SetLong(v float64) {
	o.Long = v
}

// GetCity returns the City field value
func (o *AuthenticatedSessionGeoIp) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionGeoIp) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *AuthenticatedSessionGeoIp) SetCity(v string) {
	o.City = v
}

func (o AuthenticatedSessionGeoIp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatedSessionGeoIp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["continent"] = o.Continent
	toSerialize["country"] = o.Country
	toSerialize["lat"] = o.Lat
	toSerialize["long"] = o.Long
	toSerialize["city"] = o.City
	return toSerialize, nil
}

type NullableAuthenticatedSessionGeoIp struct {
	value *AuthenticatedSessionGeoIp
	isSet bool
}

func (v NullableAuthenticatedSessionGeoIp) Get() *AuthenticatedSessionGeoIp {
	return v.value
}

func (v *NullableAuthenticatedSessionGeoIp) Set(val *AuthenticatedSessionGeoIp) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatedSessionGeoIp) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatedSessionGeoIp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatedSessionGeoIp(val *AuthenticatedSessionGeoIp) *NullableAuthenticatedSessionGeoIp {
	return &NullableAuthenticatedSessionGeoIp{value: val, isSet: true}
}

func (v NullableAuthenticatedSessionGeoIp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatedSessionGeoIp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


