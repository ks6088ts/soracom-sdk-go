/*
SORACOM API

SORACOM API v1

API version: VERSION_PLACEHOLDER
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// LoraDevice struct for LoraDevice
type LoraDevice struct {
	DeviceId *string `json:"device_id,omitempty"`
	GroupId *string `json:"groupId,omitempty"`
	LastModifiedTime *time.Time `json:"lastModifiedTime,omitempty"`
	LastSeen *LastSeen `json:"lastSeen,omitempty"`
	OperatorId *string `json:"operatorId,omitempty"`
	Status *string `json:"status,omitempty"`
	Tags *map[string]string `json:"tags,omitempty"`
	TerminationEnabled *bool `json:"terminationEnabled,omitempty"`
}

// NewLoraDevice instantiates a new LoraDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoraDevice() *LoraDevice {
	this := LoraDevice{}
	var terminationEnabled bool = false
	this.TerminationEnabled = &terminationEnabled
	return &this
}

// NewLoraDeviceWithDefaults instantiates a new LoraDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoraDeviceWithDefaults() *LoraDevice {
	this := LoraDevice{}
	var terminationEnabled bool = false
	this.TerminationEnabled = &terminationEnabled
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *LoraDevice) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraDevice) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *LoraDevice) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *LoraDevice) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *LoraDevice) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraDevice) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *LoraDevice) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *LoraDevice) SetGroupId(v string) {
	o.GroupId = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *LoraDevice) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraDevice) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *LoraDevice) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *LoraDevice) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *LoraDevice) GetLastSeen() LastSeen {
	if o == nil || o.LastSeen == nil {
		var ret LastSeen
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraDevice) GetLastSeenOk() (*LastSeen, bool) {
	if o == nil || o.LastSeen == nil {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *LoraDevice) HasLastSeen() bool {
	if o != nil && o.LastSeen != nil {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given LastSeen and assigns it to the LastSeen field.
func (o *LoraDevice) SetLastSeen(v LastSeen) {
	o.LastSeen = &v
}

// GetOperatorId returns the OperatorId field value if set, zero value otherwise.
func (o *LoraDevice) GetOperatorId() string {
	if o == nil || o.OperatorId == nil {
		var ret string
		return ret
	}
	return *o.OperatorId
}

// GetOperatorIdOk returns a tuple with the OperatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraDevice) GetOperatorIdOk() (*string, bool) {
	if o == nil || o.OperatorId == nil {
		return nil, false
	}
	return o.OperatorId, true
}

// HasOperatorId returns a boolean if a field has been set.
func (o *LoraDevice) HasOperatorId() bool {
	if o != nil && o.OperatorId != nil {
		return true
	}

	return false
}

// SetOperatorId gets a reference to the given string and assigns it to the OperatorId field.
func (o *LoraDevice) SetOperatorId(v string) {
	o.OperatorId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LoraDevice) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraDevice) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LoraDevice) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LoraDevice) SetStatus(v string) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LoraDevice) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraDevice) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LoraDevice) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *LoraDevice) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetTerminationEnabled returns the TerminationEnabled field value if set, zero value otherwise.
func (o *LoraDevice) GetTerminationEnabled() bool {
	if o == nil || o.TerminationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TerminationEnabled
}

// GetTerminationEnabledOk returns a tuple with the TerminationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraDevice) GetTerminationEnabledOk() (*bool, bool) {
	if o == nil || o.TerminationEnabled == nil {
		return nil, false
	}
	return o.TerminationEnabled, true
}

// HasTerminationEnabled returns a boolean if a field has been set.
func (o *LoraDevice) HasTerminationEnabled() bool {
	if o != nil && o.TerminationEnabled != nil {
		return true
	}

	return false
}

// SetTerminationEnabled gets a reference to the given bool and assigns it to the TerminationEnabled field.
func (o *LoraDevice) SetTerminationEnabled(v bool) {
	o.TerminationEnabled = &v
}

func (o LoraDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceId != nil {
		toSerialize["device_id"] = o.DeviceId
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.LastModifiedTime != nil {
		toSerialize["lastModifiedTime"] = o.LastModifiedTime
	}
	if o.LastSeen != nil {
		toSerialize["lastSeen"] = o.LastSeen
	}
	if o.OperatorId != nil {
		toSerialize["operatorId"] = o.OperatorId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TerminationEnabled != nil {
		toSerialize["terminationEnabled"] = o.TerminationEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableLoraDevice struct {
	value *LoraDevice
	isSet bool
}

func (v NullableLoraDevice) Get() *LoraDevice {
	return v.value
}

func (v *NullableLoraDevice) Set(val *LoraDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableLoraDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableLoraDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoraDevice(val *LoraDevice) *NullableLoraDevice {
	return &NullableLoraDevice{value: val, isSet: true}
}

func (v NullableLoraDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoraDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

