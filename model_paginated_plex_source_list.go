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

// checks if the PaginatedPlexSourceList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedPlexSourceList{}

// PaginatedPlexSourceList struct for PaginatedPlexSourceList
type PaginatedPlexSourceList struct {
	Pagination Pagination `json:"pagination"`
	Results []PlexSource `json:"results"`
}

// NewPaginatedPlexSourceList instantiates a new PaginatedPlexSourceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedPlexSourceList(pagination Pagination, results []PlexSource) *PaginatedPlexSourceList {
	this := PaginatedPlexSourceList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedPlexSourceListWithDefaults instantiates a new PaginatedPlexSourceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedPlexSourceListWithDefaults() *PaginatedPlexSourceList {
	this := PaginatedPlexSourceList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedPlexSourceList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedPlexSourceList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedPlexSourceList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedPlexSourceList) GetResults() []PlexSource {
	if o == nil {
		var ret []PlexSource
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedPlexSourceList) GetResultsOk() ([]PlexSource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedPlexSourceList) SetResults(v []PlexSource) {
	o.Results = v
}

func (o PaginatedPlexSourceList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedPlexSourceList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pagination"] = o.Pagination
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullablePaginatedPlexSourceList struct {
	value *PaginatedPlexSourceList
	isSet bool
}

func (v NullablePaginatedPlexSourceList) Get() *PaginatedPlexSourceList {
	return v.value
}

func (v *NullablePaginatedPlexSourceList) Set(val *PaginatedPlexSourceList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedPlexSourceList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedPlexSourceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedPlexSourceList(val *PaginatedPlexSourceList) *NullablePaginatedPlexSourceList {
	return &NullablePaginatedPlexSourceList{value: val, isSet: true}
}

func (v NullablePaginatedPlexSourceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedPlexSourceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


