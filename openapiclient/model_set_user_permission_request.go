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

// SetUserPermissionRequest struct for SetUserPermissionRequest
type SetUserPermissionRequest struct {
	Description *string `json:"description,omitempty"`
	// Permission の JSON
	Permission string `json:"permission"`
}

// NewSetUserPermissionRequest instantiates a new SetUserPermissionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetUserPermissionRequest(permission string) *SetUserPermissionRequest {
	this := SetUserPermissionRequest{}
	this.Permission = permission
	return &this
}

// NewSetUserPermissionRequestWithDefaults instantiates a new SetUserPermissionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetUserPermissionRequestWithDefaults() *SetUserPermissionRequest {
	this := SetUserPermissionRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SetUserPermissionRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetUserPermissionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SetUserPermissionRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SetUserPermissionRequest) SetDescription(v string) {
	o.Description = &v
}

// GetPermission returns the Permission field value
func (o *SetUserPermissionRequest) GetPermission() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *SetUserPermissionRequest) GetPermissionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *SetUserPermissionRequest) SetPermission(v string) {
	o.Permission = v
}

func (o SetUserPermissionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["permission"] = o.Permission
	}
	return json.Marshal(toSerialize)
}

type NullableSetUserPermissionRequest struct {
	value *SetUserPermissionRequest
	isSet bool
}

func (v NullableSetUserPermissionRequest) Get() *SetUserPermissionRequest {
	return v.value
}

func (v *NullableSetUserPermissionRequest) Set(val *SetUserPermissionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetUserPermissionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetUserPermissionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetUserPermissionRequest(val *SetUserPermissionRequest) *NullableSetUserPermissionRequest {
	return &NullableSetUserPermissionRequest{value: val, isSet: true}
}

func (v NullableSetUserPermissionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetUserPermissionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

