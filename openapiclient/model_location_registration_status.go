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

// LocationRegistrationStatus struct for LocationRegistrationStatus
type LocationRegistrationStatus struct {
	Cs *LocationRegistrationStatusForNetwork `json:"cs,omitempty"`
	Eps *LocationRegistrationStatusForNetwork `json:"eps,omitempty"`
	Ps *LocationRegistrationStatusForNetwork `json:"ps,omitempty"`
}

// NewLocationRegistrationStatus instantiates a new LocationRegistrationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationRegistrationStatus() *LocationRegistrationStatus {
	this := LocationRegistrationStatus{}
	return &this
}

// NewLocationRegistrationStatusWithDefaults instantiates a new LocationRegistrationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationRegistrationStatusWithDefaults() *LocationRegistrationStatus {
	this := LocationRegistrationStatus{}
	return &this
}

// GetCs returns the Cs field value if set, zero value otherwise.
func (o *LocationRegistrationStatus) GetCs() LocationRegistrationStatusForNetwork {
	if o == nil || o.Cs == nil {
		var ret LocationRegistrationStatusForNetwork
		return ret
	}
	return *o.Cs
}

// GetCsOk returns a tuple with the Cs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRegistrationStatus) GetCsOk() (*LocationRegistrationStatusForNetwork, bool) {
	if o == nil || o.Cs == nil {
		return nil, false
	}
	return o.Cs, true
}

// HasCs returns a boolean if a field has been set.
func (o *LocationRegistrationStatus) HasCs() bool {
	if o != nil && o.Cs != nil {
		return true
	}

	return false
}

// SetCs gets a reference to the given LocationRegistrationStatusForNetwork and assigns it to the Cs field.
func (o *LocationRegistrationStatus) SetCs(v LocationRegistrationStatusForNetwork) {
	o.Cs = &v
}

// GetEps returns the Eps field value if set, zero value otherwise.
func (o *LocationRegistrationStatus) GetEps() LocationRegistrationStatusForNetwork {
	if o == nil || o.Eps == nil {
		var ret LocationRegistrationStatusForNetwork
		return ret
	}
	return *o.Eps
}

// GetEpsOk returns a tuple with the Eps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRegistrationStatus) GetEpsOk() (*LocationRegistrationStatusForNetwork, bool) {
	if o == nil || o.Eps == nil {
		return nil, false
	}
	return o.Eps, true
}

// HasEps returns a boolean if a field has been set.
func (o *LocationRegistrationStatus) HasEps() bool {
	if o != nil && o.Eps != nil {
		return true
	}

	return false
}

// SetEps gets a reference to the given LocationRegistrationStatusForNetwork and assigns it to the Eps field.
func (o *LocationRegistrationStatus) SetEps(v LocationRegistrationStatusForNetwork) {
	o.Eps = &v
}

// GetPs returns the Ps field value if set, zero value otherwise.
func (o *LocationRegistrationStatus) GetPs() LocationRegistrationStatusForNetwork {
	if o == nil || o.Ps == nil {
		var ret LocationRegistrationStatusForNetwork
		return ret
	}
	return *o.Ps
}

// GetPsOk returns a tuple with the Ps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRegistrationStatus) GetPsOk() (*LocationRegistrationStatusForNetwork, bool) {
	if o == nil || o.Ps == nil {
		return nil, false
	}
	return o.Ps, true
}

// HasPs returns a boolean if a field has been set.
func (o *LocationRegistrationStatus) HasPs() bool {
	if o != nil && o.Ps != nil {
		return true
	}

	return false
}

// SetPs gets a reference to the given LocationRegistrationStatusForNetwork and assigns it to the Ps field.
func (o *LocationRegistrationStatus) SetPs(v LocationRegistrationStatusForNetwork) {
	o.Ps = &v
}

func (o LocationRegistrationStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cs != nil {
		toSerialize["cs"] = o.Cs
	}
	if o.Eps != nil {
		toSerialize["eps"] = o.Eps
	}
	if o.Ps != nil {
		toSerialize["ps"] = o.Ps
	}
	return json.Marshal(toSerialize)
}

type NullableLocationRegistrationStatus struct {
	value *LocationRegistrationStatus
	isSet bool
}

func (v NullableLocationRegistrationStatus) Get() *LocationRegistrationStatus {
	return v.value
}

func (v *NullableLocationRegistrationStatus) Set(val *LocationRegistrationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationRegistrationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationRegistrationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationRegistrationStatus(val *LocationRegistrationStatus) *NullableLocationRegistrationStatus {
	return &NullableLocationRegistrationStatus{value: val, isSet: true}
}

func (v NullableLocationRegistrationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationRegistrationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


