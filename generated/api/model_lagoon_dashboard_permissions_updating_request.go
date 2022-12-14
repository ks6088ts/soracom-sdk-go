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

// LagoonDashboardPermissionsUpdatingRequest struct for LagoonDashboardPermissionsUpdatingRequest
type LagoonDashboardPermissionsUpdatingRequest struct {
	Permissions []LagoonDashboardPermissionsUpdatingRequestPermissionsInner `json:"permissions,omitempty"`
}

// NewLagoonDashboardPermissionsUpdatingRequest instantiates a new LagoonDashboardPermissionsUpdatingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLagoonDashboardPermissionsUpdatingRequest() *LagoonDashboardPermissionsUpdatingRequest {
	this := LagoonDashboardPermissionsUpdatingRequest{}
	return &this
}

// NewLagoonDashboardPermissionsUpdatingRequestWithDefaults instantiates a new LagoonDashboardPermissionsUpdatingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLagoonDashboardPermissionsUpdatingRequestWithDefaults() *LagoonDashboardPermissionsUpdatingRequest {
	this := LagoonDashboardPermissionsUpdatingRequest{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *LagoonDashboardPermissionsUpdatingRequest) GetPermissions() []LagoonDashboardPermissionsUpdatingRequestPermissionsInner {
	if o == nil || o.Permissions == nil {
		var ret []LagoonDashboardPermissionsUpdatingRequestPermissionsInner
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LagoonDashboardPermissionsUpdatingRequest) GetPermissionsOk() ([]LagoonDashboardPermissionsUpdatingRequestPermissionsInner, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *LagoonDashboardPermissionsUpdatingRequest) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []LagoonDashboardPermissionsUpdatingRequestPermissionsInner and assigns it to the Permissions field.
func (o *LagoonDashboardPermissionsUpdatingRequest) SetPermissions(v []LagoonDashboardPermissionsUpdatingRequestPermissionsInner) {
	o.Permissions = v
}

func (o LagoonDashboardPermissionsUpdatingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableLagoonDashboardPermissionsUpdatingRequest struct {
	value *LagoonDashboardPermissionsUpdatingRequest
	isSet bool
}

func (v NullableLagoonDashboardPermissionsUpdatingRequest) Get() *LagoonDashboardPermissionsUpdatingRequest {
	return v.value
}

func (v *NullableLagoonDashboardPermissionsUpdatingRequest) Set(val *LagoonDashboardPermissionsUpdatingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLagoonDashboardPermissionsUpdatingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLagoonDashboardPermissionsUpdatingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLagoonDashboardPermissionsUpdatingRequest(val *LagoonDashboardPermissionsUpdatingRequest) *NullableLagoonDashboardPermissionsUpdatingRequest {
	return &NullableLagoonDashboardPermissionsUpdatingRequest{value: val, isSet: true}
}

func (v NullableLagoonDashboardPermissionsUpdatingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLagoonDashboardPermissionsUpdatingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


