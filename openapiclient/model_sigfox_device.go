/*
SORACOM API

SORACOM API v1

API version: VERSION_PLACEHOLDER
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"encoding/json"
	"time"
)

// SigfoxDevice struct for SigfoxDevice
type SigfoxDevice struct {
	DeviceId *string `json:"device_id,omitempty"`
	GroupId *string `json:"groupId,omitempty"`
	LastModifiedTime *time.Time `json:"lastModifiedTime,omitempty"`
	LastSeen *LastSeen `json:"lastSeen,omitempty"`
	OperatorId *string `json:"operatorId,omitempty"`
	Status *string `json:"status,omitempty"`
	Tags *map[string]string `json:"tags,omitempty"`
	TerminationEnabled *bool `json:"terminationEnabled,omitempty"`
}

// NewSigfoxDevice instantiates a new SigfoxDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSigfoxDevice() *SigfoxDevice {
	this := SigfoxDevice{}
	var terminationEnabled bool = false
	this.TerminationEnabled = &terminationEnabled
	return &this
}

// NewSigfoxDeviceWithDefaults instantiates a new SigfoxDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSigfoxDeviceWithDefaults() *SigfoxDevice {
	this := SigfoxDevice{}
	var terminationEnabled bool = false
	this.TerminationEnabled = &terminationEnabled
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *SigfoxDevice) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigfoxDevice) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *SigfoxDevice) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *SigfoxDevice) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *SigfoxDevice) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigfoxDevice) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *SigfoxDevice) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *SigfoxDevice) SetGroupId(v string) {
	o.GroupId = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *SigfoxDevice) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigfoxDevice) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *SigfoxDevice) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *SigfoxDevice) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *SigfoxDevice) GetLastSeen() LastSeen {
	if o == nil || o.LastSeen == nil {
		var ret LastSeen
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigfoxDevice) GetLastSeenOk() (*LastSeen, bool) {
	if o == nil || o.LastSeen == nil {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *SigfoxDevice) HasLastSeen() bool {
	if o != nil && o.LastSeen != nil {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given LastSeen and assigns it to the LastSeen field.
func (o *SigfoxDevice) SetLastSeen(v LastSeen) {
	o.LastSeen = &v
}

// GetOperatorId returns the OperatorId field value if set, zero value otherwise.
func (o *SigfoxDevice) GetOperatorId() string {
	if o == nil || o.OperatorId == nil {
		var ret string
		return ret
	}
	return *o.OperatorId
}

// GetOperatorIdOk returns a tuple with the OperatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigfoxDevice) GetOperatorIdOk() (*string, bool) {
	if o == nil || o.OperatorId == nil {
		return nil, false
	}
	return o.OperatorId, true
}

// HasOperatorId returns a boolean if a field has been set.
func (o *SigfoxDevice) HasOperatorId() bool {
	if o != nil && o.OperatorId != nil {
		return true
	}

	return false
}

// SetOperatorId gets a reference to the given string and assigns it to the OperatorId field.
func (o *SigfoxDevice) SetOperatorId(v string) {
	o.OperatorId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SigfoxDevice) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigfoxDevice) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SigfoxDevice) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SigfoxDevice) SetStatus(v string) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SigfoxDevice) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigfoxDevice) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SigfoxDevice) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *SigfoxDevice) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetTerminationEnabled returns the TerminationEnabled field value if set, zero value otherwise.
func (o *SigfoxDevice) GetTerminationEnabled() bool {
	if o == nil || o.TerminationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TerminationEnabled
}

// GetTerminationEnabledOk returns a tuple with the TerminationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigfoxDevice) GetTerminationEnabledOk() (*bool, bool) {
	if o == nil || o.TerminationEnabled == nil {
		return nil, false
	}
	return o.TerminationEnabled, true
}

// HasTerminationEnabled returns a boolean if a field has been set.
func (o *SigfoxDevice) HasTerminationEnabled() bool {
	if o != nil && o.TerminationEnabled != nil {
		return true
	}

	return false
}

// SetTerminationEnabled gets a reference to the given bool and assigns it to the TerminationEnabled field.
func (o *SigfoxDevice) SetTerminationEnabled(v bool) {
	o.TerminationEnabled = &v
}

func (o SigfoxDevice) MarshalJSON() ([]byte, error) {
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

type NullableSigfoxDevice struct {
	value *SigfoxDevice
	isSet bool
}

func (v NullableSigfoxDevice) Get() *SigfoxDevice {
	return v.value
}

func (v *NullableSigfoxDevice) Set(val *SigfoxDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableSigfoxDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableSigfoxDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSigfoxDevice(val *SigfoxDevice) *NullableSigfoxDevice {
	return &NullableSigfoxDevice{value: val, isSet: true}
}

func (v NullableSigfoxDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSigfoxDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


