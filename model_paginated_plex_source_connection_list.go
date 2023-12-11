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

// checks if the PaginatedPlexSourceConnectionList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedPlexSourceConnectionList{}

// PaginatedPlexSourceConnectionList struct for PaginatedPlexSourceConnectionList
type PaginatedPlexSourceConnectionList struct {
	Pagination Pagination `json:"pagination"`
	Results []PlexSourceConnection `json:"results"`
}

// NewPaginatedPlexSourceConnectionList instantiates a new PaginatedPlexSourceConnectionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedPlexSourceConnectionList(pagination Pagination, results []PlexSourceConnection) *PaginatedPlexSourceConnectionList {
	this := PaginatedPlexSourceConnectionList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedPlexSourceConnectionListWithDefaults instantiates a new PaginatedPlexSourceConnectionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedPlexSourceConnectionListWithDefaults() *PaginatedPlexSourceConnectionList {
	this := PaginatedPlexSourceConnectionList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedPlexSourceConnectionList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedPlexSourceConnectionList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedPlexSourceConnectionList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedPlexSourceConnectionList) GetResults() []PlexSourceConnection {
	if o == nil {
		var ret []PlexSourceConnection
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedPlexSourceConnectionList) GetResultsOk() ([]PlexSourceConnection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedPlexSourceConnectionList) SetResults(v []PlexSourceConnection) {
	o.Results = v
}

func (o PaginatedPlexSourceConnectionList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedPlexSourceConnectionList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pagination"] = o.Pagination
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullablePaginatedPlexSourceConnectionList struct {
	value *PaginatedPlexSourceConnectionList
	isSet bool
}

func (v NullablePaginatedPlexSourceConnectionList) Get() *PaginatedPlexSourceConnectionList {
	return v.value
}

func (v *NullablePaginatedPlexSourceConnectionList) Set(val *PaginatedPlexSourceConnectionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedPlexSourceConnectionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedPlexSourceConnectionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedPlexSourceConnectionList(val *PaginatedPlexSourceConnectionList) *NullablePaginatedPlexSourceConnectionList {
	return &NullablePaginatedPlexSourceConnectionList{value: val, isSet: true}
}

func (v NullablePaginatedPlexSourceConnectionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedPlexSourceConnectionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


