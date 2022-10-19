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

// GroupConfigurationUpdateRequest struct for GroupConfigurationUpdateRequest
type GroupConfigurationUpdateRequest struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

// NewGroupConfigurationUpdateRequest instantiates a new GroupConfigurationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupConfigurationUpdateRequest(key string, value string) *GroupConfigurationUpdateRequest {
	this := GroupConfigurationUpdateRequest{}
	this.Key = key
	this.Value = value
	return &this
}

// NewGroupConfigurationUpdateRequestWithDefaults instantiates a new GroupConfigurationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupConfigurationUpdateRequestWithDefaults() *GroupConfigurationUpdateRequest {
	this := GroupConfigurationUpdateRequest{}
	return &this
}

// GetKey returns the Key field value
func (o *GroupConfigurationUpdateRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *GroupConfigurationUpdateRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *GroupConfigurationUpdateRequest) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *GroupConfigurationUpdateRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GroupConfigurationUpdateRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GroupConfigurationUpdateRequest) SetValue(v string) {
	o.Value = v
}

func (o GroupConfigurationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableGroupConfigurationUpdateRequest struct {
	value *GroupConfigurationUpdateRequest
	isSet bool
}

func (v NullableGroupConfigurationUpdateRequest) Get() *GroupConfigurationUpdateRequest {
	return v.value
}

func (v *NullableGroupConfigurationUpdateRequest) Set(val *GroupConfigurationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupConfigurationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupConfigurationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupConfigurationUpdateRequest(val *GroupConfigurationUpdateRequest) *NullableGroupConfigurationUpdateRequest {
	return &NullableGroupConfigurationUpdateRequest{value: val, isSet: true}
}

func (v NullableGroupConfigurationUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupConfigurationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


