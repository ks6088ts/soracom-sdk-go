/*
SORACOM API

SORACOM API v1

API version: VERSION_PLACEHOLDER
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"encoding/json"
)

// SoracomFunkStats struct for SoracomFunkStats
type SoracomFunkStats struct {
	Count *int64 `json:"count,omitempty"`
}

// NewSoracomFunkStats instantiates a new SoracomFunkStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoracomFunkStats() *SoracomFunkStats {
	this := SoracomFunkStats{}
	return &this
}

// NewSoracomFunkStatsWithDefaults instantiates a new SoracomFunkStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoracomFunkStatsWithDefaults() *SoracomFunkStats {
	this := SoracomFunkStats{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SoracomFunkStats) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoracomFunkStats) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SoracomFunkStats) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *SoracomFunkStats) SetCount(v int64) {
	o.Count = &v
}

func (o SoracomFunkStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableSoracomFunkStats struct {
	value *SoracomFunkStats
	isSet bool
}

func (v NullableSoracomFunkStats) Get() *SoracomFunkStats {
	return v.value
}

func (v *NullableSoracomFunkStats) Set(val *SoracomFunkStats) {
	v.value = val
	v.isSet = true
}

func (v NullableSoracomFunkStats) IsSet() bool {
	return v.isSet
}

func (v *NullableSoracomFunkStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoracomFunkStats(val *SoracomFunkStats) *NullableSoracomFunkStats {
	return &NullableSoracomFunkStats{value: val, isSet: true}
}

func (v NullableSoracomFunkStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoracomFunkStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

