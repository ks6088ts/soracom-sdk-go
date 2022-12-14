/*
SORACOM API

SORACOM API v1

API version: VERSION_PLACEHOLDER
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SubscriptionContainerStatus struct for SubscriptionContainerStatus
type SubscriptionContainerStatus struct {
	Containers []SubscriptionContainer `json:"containers,omitempty"`
	CountryMapping *map[string]SubscriptionContainerStatusCountryMappingValue `json:"countryMapping,omitempty"`
}

// NewSubscriptionContainerStatus instantiates a new SubscriptionContainerStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionContainerStatus() *SubscriptionContainerStatus {
	this := SubscriptionContainerStatus{}
	return &this
}

// NewSubscriptionContainerStatusWithDefaults instantiates a new SubscriptionContainerStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionContainerStatusWithDefaults() *SubscriptionContainerStatus {
	this := SubscriptionContainerStatus{}
	return &this
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *SubscriptionContainerStatus) GetContainers() []SubscriptionContainer {
	if o == nil || o.Containers == nil {
		var ret []SubscriptionContainer
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionContainerStatus) GetContainersOk() ([]SubscriptionContainer, bool) {
	if o == nil || o.Containers == nil {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *SubscriptionContainerStatus) HasContainers() bool {
	if o != nil && o.Containers != nil {
		return true
	}

	return false
}

// SetContainers gets a reference to the given []SubscriptionContainer and assigns it to the Containers field.
func (o *SubscriptionContainerStatus) SetContainers(v []SubscriptionContainer) {
	o.Containers = v
}

// GetCountryMapping returns the CountryMapping field value if set, zero value otherwise.
func (o *SubscriptionContainerStatus) GetCountryMapping() map[string]SubscriptionContainerStatusCountryMappingValue {
	if o == nil || o.CountryMapping == nil {
		var ret map[string]SubscriptionContainerStatusCountryMappingValue
		return ret
	}
	return *o.CountryMapping
}

// GetCountryMappingOk returns a tuple with the CountryMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionContainerStatus) GetCountryMappingOk() (*map[string]SubscriptionContainerStatusCountryMappingValue, bool) {
	if o == nil || o.CountryMapping == nil {
		return nil, false
	}
	return o.CountryMapping, true
}

// HasCountryMapping returns a boolean if a field has been set.
func (o *SubscriptionContainerStatus) HasCountryMapping() bool {
	if o != nil && o.CountryMapping != nil {
		return true
	}

	return false
}

// SetCountryMapping gets a reference to the given map[string]SubscriptionContainerStatusCountryMappingValue and assigns it to the CountryMapping field.
func (o *SubscriptionContainerStatus) SetCountryMapping(v map[string]SubscriptionContainerStatusCountryMappingValue) {
	o.CountryMapping = &v
}

func (o SubscriptionContainerStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Containers != nil {
		toSerialize["containers"] = o.Containers
	}
	if o.CountryMapping != nil {
		toSerialize["countryMapping"] = o.CountryMapping
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionContainerStatus struct {
	value *SubscriptionContainerStatus
	isSet bool
}

func (v NullableSubscriptionContainerStatus) Get() *SubscriptionContainerStatus {
	return v.value
}

func (v *NullableSubscriptionContainerStatus) Set(val *SubscriptionContainerStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionContainerStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionContainerStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionContainerStatus(val *SubscriptionContainerStatus) *NullableSubscriptionContainerStatus {
	return &NullableSubscriptionContainerStatus{value: val, isSet: true}
}

func (v NullableSubscriptionContainerStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionContainerStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


