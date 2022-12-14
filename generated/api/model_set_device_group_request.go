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

// SetDeviceGroupRequest struct for SetDeviceGroupRequest
type SetDeviceGroupRequest struct {
	GroupId *string `json:"groupId,omitempty"`
}

// NewSetDeviceGroupRequest instantiates a new SetDeviceGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetDeviceGroupRequest() *SetDeviceGroupRequest {
	this := SetDeviceGroupRequest{}
	return &this
}

// NewSetDeviceGroupRequestWithDefaults instantiates a new SetDeviceGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetDeviceGroupRequestWithDefaults() *SetDeviceGroupRequest {
	this := SetDeviceGroupRequest{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *SetDeviceGroupRequest) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetDeviceGroupRequest) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *SetDeviceGroupRequest) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *SetDeviceGroupRequest) SetGroupId(v string) {
	o.GroupId = &v
}

func (o SetDeviceGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	return json.Marshal(toSerialize)
}

type NullableSetDeviceGroupRequest struct {
	value *SetDeviceGroupRequest
	isSet bool
}

func (v NullableSetDeviceGroupRequest) Get() *SetDeviceGroupRequest {
	return v.value
}

func (v *NullableSetDeviceGroupRequest) Set(val *SetDeviceGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetDeviceGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetDeviceGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetDeviceGroupRequest(val *SetDeviceGroupRequest) *NullableSetDeviceGroupRequest {
	return &NullableSetDeviceGroupRequest{value: val, isSet: true}
}

func (v NullableSetDeviceGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetDeviceGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


