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

// UpdatePermissionRequest struct for UpdatePermissionRequest
type UpdatePermissionRequest struct {
	OperatorId *string `json:"operatorId,omitempty"`
}

// NewUpdatePermissionRequest instantiates a new UpdatePermissionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePermissionRequest() *UpdatePermissionRequest {
	this := UpdatePermissionRequest{}
	return &this
}

// NewUpdatePermissionRequestWithDefaults instantiates a new UpdatePermissionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePermissionRequestWithDefaults() *UpdatePermissionRequest {
	this := UpdatePermissionRequest{}
	return &this
}

// GetOperatorId returns the OperatorId field value if set, zero value otherwise.
func (o *UpdatePermissionRequest) GetOperatorId() string {
	if o == nil || o.OperatorId == nil {
		var ret string
		return ret
	}
	return *o.OperatorId
}

// GetOperatorIdOk returns a tuple with the OperatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePermissionRequest) GetOperatorIdOk() (*string, bool) {
	if o == nil || o.OperatorId == nil {
		return nil, false
	}
	return o.OperatorId, true
}

// HasOperatorId returns a boolean if a field has been set.
func (o *UpdatePermissionRequest) HasOperatorId() bool {
	if o != nil && o.OperatorId != nil {
		return true
	}

	return false
}

// SetOperatorId gets a reference to the given string and assigns it to the OperatorId field.
func (o *UpdatePermissionRequest) SetOperatorId(v string) {
	o.OperatorId = &v
}

func (o UpdatePermissionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OperatorId != nil {
		toSerialize["operatorId"] = o.OperatorId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdatePermissionRequest struct {
	value *UpdatePermissionRequest
	isSet bool
}

func (v NullableUpdatePermissionRequest) Get() *UpdatePermissionRequest {
	return v.value
}

func (v *NullableUpdatePermissionRequest) Set(val *UpdatePermissionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePermissionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePermissionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePermissionRequest(val *UpdatePermissionRequest) *NullableUpdatePermissionRequest {
	return &NullableUpdatePermissionRequest{value: val, isSet: true}
}

func (v NullableUpdatePermissionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePermissionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


