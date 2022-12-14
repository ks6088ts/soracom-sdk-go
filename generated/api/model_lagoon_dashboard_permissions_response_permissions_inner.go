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

// LagoonDashboardPermissionsResponsePermissionsInner struct for LagoonDashboardPermissionsResponsePermissionsInner
type LagoonDashboardPermissionsResponsePermissionsInner struct {
	DashboardId *float32 `json:"dashboardId,omitempty"`
	DashboardTitle *string `json:"dashboardTitle,omitempty"`
	PermissionName *string `json:"permissionName,omitempty"`
	UserEmail *string `json:"userEmail,omitempty"`
	UserId *float32 `json:"userId,omitempty"`
}

// NewLagoonDashboardPermissionsResponsePermissionsInner instantiates a new LagoonDashboardPermissionsResponsePermissionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLagoonDashboardPermissionsResponsePermissionsInner() *LagoonDashboardPermissionsResponsePermissionsInner {
	this := LagoonDashboardPermissionsResponsePermissionsInner{}
	return &this
}

// NewLagoonDashboardPermissionsResponsePermissionsInnerWithDefaults instantiates a new LagoonDashboardPermissionsResponsePermissionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLagoonDashboardPermissionsResponsePermissionsInnerWithDefaults() *LagoonDashboardPermissionsResponsePermissionsInner {
	this := LagoonDashboardPermissionsResponsePermissionsInner{}
	return &this
}

// GetDashboardId returns the DashboardId field value if set, zero value otherwise.
func (o *LagoonDashboardPermissionsResponsePermissionsInner) GetDashboardId() float32 {
	if o == nil || o.DashboardId == nil {
		var ret float32
		return ret
	}
	return *o.DashboardId
}

// GetDashboardIdOk returns a tuple with the DashboardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LagoonDashboardPermissionsResponsePermissionsInner) GetDashboardIdOk() (*float32, bool) {
	if o == nil || o.DashboardId == nil {
		return nil, false
	}
	return o.DashboardId, true
}

// HasDashboardId returns a boolean if a field has been set.
func (o *LagoonDashboardPermissionsResponsePermissionsInner) HasDashboardId() bool {
	if o != nil && o.DashboardId != nil {
		return true
	}

	return false
}

// SetDashboardId gets a reference to the given float32 and assigns it to the DashboardId field.
func (o *LagoonDashboardPermissionsResponsePermissionsInner) SetDashboardId(v float32) {
	o.DashboardId = &v
}

// GetDashboardTitle returns the DashboardTitle field value if set, zero value otherwise.
func (o *LagoonDashboardPermissionsResponsePermissionsInner) GetDashboardTitle() string {
	if o == nil || o.DashboardTitle == nil {
		var ret string
		return ret
	}
	return *o.DashboardTitle
}

// GetDashboardTitleOk returns a tuple with the DashboardTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LagoonDashboardPermissionsResponsePermissionsInner) GetDashboardTitleOk() (*string, bool) {
	if o == nil || o.DashboardTitle == nil {
		return nil, false
	}
	return o.DashboardTitle, true
}

// HasDashboardTitle returns a boolean if a field has been set.
func (o *LagoonDashboardPermissionsResponsePermissionsInner) HasDashboardTitle() bool {
	if o != nil && o.DashboardTitle != nil {
		return true
	}

	return false
}

// SetDashboardTitle gets a reference to the given string and assigns it to the DashboardTitle field.
func (o *LagoonDashboardPermissionsResponsePermissionsInner) SetDashboardTitle(v string) {
	o.DashboardTitle = &v
}

// GetPermissionName returns the PermissionName field value if set, zero value otherwise.
func (o *LagoonDashboardPermissionsResponsePermissionsInner) GetPermissionName() string {
	if o == nil || o.PermissionName == nil {
		var ret string
		return ret
	}
	return *o.PermissionName
}

// GetPermissionNameOk returns a tuple with the PermissionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LagoonDashboardPermissionsResponsePermissionsInner) GetPermissionNameOk() (*string, bool) {
	if o == nil || o.PermissionName == nil {
		return nil, false
	}
	return o.PermissionName, true
}

// HasPermissionName returns a boolean if a field has been set.
func (o *LagoonDashboardPermissionsResponsePermissionsInner) HasPermissionName() bool {
	if o != nil && o.PermissionName != nil {
		return true
	}

	return false
}

// SetPermissionName gets a reference to the given string and assigns it to the PermissionName field.
func (o *LagoonDashboardPermissionsResponsePermissionsInner) SetPermissionName(v string) {
	o.PermissionName = &v
}

// GetUserEmail returns the UserEmail field value if set, zero value otherwise.
func (o *LagoonDashboardPermissionsResponsePermissionsInner) GetUserEmail() string {
	if o == nil || o.UserEmail == nil {
		var ret string
		return ret
	}
	return *o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LagoonDashboardPermissionsResponsePermissionsInner) GetUserEmailOk() (*string, bool) {
	if o == nil || o.UserEmail == nil {
		return nil, false
	}
	return o.UserEmail, true
}

// HasUserEmail returns a boolean if a field has been set.
func (o *LagoonDashboardPermissionsResponsePermissionsInner) HasUserEmail() bool {
	if o != nil && o.UserEmail != nil {
		return true
	}

	return false
}

// SetUserEmail gets a reference to the given string and assigns it to the UserEmail field.
func (o *LagoonDashboardPermissionsResponsePermissionsInner) SetUserEmail(v string) {
	o.UserEmail = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *LagoonDashboardPermissionsResponsePermissionsInner) GetUserId() float32 {
	if o == nil || o.UserId == nil {
		var ret float32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LagoonDashboardPermissionsResponsePermissionsInner) GetUserIdOk() (*float32, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *LagoonDashboardPermissionsResponsePermissionsInner) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given float32 and assigns it to the UserId field.
func (o *LagoonDashboardPermissionsResponsePermissionsInner) SetUserId(v float32) {
	o.UserId = &v
}

func (o LagoonDashboardPermissionsResponsePermissionsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DashboardId != nil {
		toSerialize["dashboardId"] = o.DashboardId
	}
	if o.DashboardTitle != nil {
		toSerialize["dashboardTitle"] = o.DashboardTitle
	}
	if o.PermissionName != nil {
		toSerialize["permissionName"] = o.PermissionName
	}
	if o.UserEmail != nil {
		toSerialize["userEmail"] = o.UserEmail
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableLagoonDashboardPermissionsResponsePermissionsInner struct {
	value *LagoonDashboardPermissionsResponsePermissionsInner
	isSet bool
}

func (v NullableLagoonDashboardPermissionsResponsePermissionsInner) Get() *LagoonDashboardPermissionsResponsePermissionsInner {
	return v.value
}

func (v *NullableLagoonDashboardPermissionsResponsePermissionsInner) Set(val *LagoonDashboardPermissionsResponsePermissionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableLagoonDashboardPermissionsResponsePermissionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableLagoonDashboardPermissionsResponsePermissionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLagoonDashboardPermissionsResponsePermissionsInner(val *LagoonDashboardPermissionsResponsePermissionsInner) *NullableLagoonDashboardPermissionsResponsePermissionsInner {
	return &NullableLagoonDashboardPermissionsResponsePermissionsInner{value: val, isSet: true}
}

func (v NullableLagoonDashboardPermissionsResponsePermissionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLagoonDashboardPermissionsResponsePermissionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


