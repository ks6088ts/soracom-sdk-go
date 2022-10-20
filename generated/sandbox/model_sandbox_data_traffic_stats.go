/*
SORACOM SANDBOX API

SORACOM SANDBOX API v1

API version: 20160218
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sandbox

import (
	"encoding/json"
)

// SandboxDataTrafficStats struct for SandboxDataTrafficStats
type SandboxDataTrafficStats struct {
	DownloadByteSizeTotal *int64 `json:"downloadByteSizeTotal,omitempty"`
	DownloadPacketSizeTotal *int64 `json:"downloadPacketSizeTotal,omitempty"`
	UploadByteSizeTotal *int64 `json:"uploadByteSizeTotal,omitempty"`
	UploadPacketSizeTotal *int64 `json:"uploadPacketSizeTotal,omitempty"`
}

// NewSandboxDataTrafficStats instantiates a new SandboxDataTrafficStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxDataTrafficStats() *SandboxDataTrafficStats {
	this := SandboxDataTrafficStats{}
	return &this
}

// NewSandboxDataTrafficStatsWithDefaults instantiates a new SandboxDataTrafficStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxDataTrafficStatsWithDefaults() *SandboxDataTrafficStats {
	this := SandboxDataTrafficStats{}
	return &this
}

// GetDownloadByteSizeTotal returns the DownloadByteSizeTotal field value if set, zero value otherwise.
func (o *SandboxDataTrafficStats) GetDownloadByteSizeTotal() int64 {
	if o == nil || o.DownloadByteSizeTotal == nil {
		var ret int64
		return ret
	}
	return *o.DownloadByteSizeTotal
}

// GetDownloadByteSizeTotalOk returns a tuple with the DownloadByteSizeTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxDataTrafficStats) GetDownloadByteSizeTotalOk() (*int64, bool) {
	if o == nil || o.DownloadByteSizeTotal == nil {
		return nil, false
	}
	return o.DownloadByteSizeTotal, true
}

// HasDownloadByteSizeTotal returns a boolean if a field has been set.
func (o *SandboxDataTrafficStats) HasDownloadByteSizeTotal() bool {
	if o != nil && o.DownloadByteSizeTotal != nil {
		return true
	}

	return false
}

// SetDownloadByteSizeTotal gets a reference to the given int64 and assigns it to the DownloadByteSizeTotal field.
func (o *SandboxDataTrafficStats) SetDownloadByteSizeTotal(v int64) {
	o.DownloadByteSizeTotal = &v
}

// GetDownloadPacketSizeTotal returns the DownloadPacketSizeTotal field value if set, zero value otherwise.
func (o *SandboxDataTrafficStats) GetDownloadPacketSizeTotal() int64 {
	if o == nil || o.DownloadPacketSizeTotal == nil {
		var ret int64
		return ret
	}
	return *o.DownloadPacketSizeTotal
}

// GetDownloadPacketSizeTotalOk returns a tuple with the DownloadPacketSizeTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxDataTrafficStats) GetDownloadPacketSizeTotalOk() (*int64, bool) {
	if o == nil || o.DownloadPacketSizeTotal == nil {
		return nil, false
	}
	return o.DownloadPacketSizeTotal, true
}

// HasDownloadPacketSizeTotal returns a boolean if a field has been set.
func (o *SandboxDataTrafficStats) HasDownloadPacketSizeTotal() bool {
	if o != nil && o.DownloadPacketSizeTotal != nil {
		return true
	}

	return false
}

// SetDownloadPacketSizeTotal gets a reference to the given int64 and assigns it to the DownloadPacketSizeTotal field.
func (o *SandboxDataTrafficStats) SetDownloadPacketSizeTotal(v int64) {
	o.DownloadPacketSizeTotal = &v
}

// GetUploadByteSizeTotal returns the UploadByteSizeTotal field value if set, zero value otherwise.
func (o *SandboxDataTrafficStats) GetUploadByteSizeTotal() int64 {
	if o == nil || o.UploadByteSizeTotal == nil {
		var ret int64
		return ret
	}
	return *o.UploadByteSizeTotal
}

// GetUploadByteSizeTotalOk returns a tuple with the UploadByteSizeTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxDataTrafficStats) GetUploadByteSizeTotalOk() (*int64, bool) {
	if o == nil || o.UploadByteSizeTotal == nil {
		return nil, false
	}
	return o.UploadByteSizeTotal, true
}

// HasUploadByteSizeTotal returns a boolean if a field has been set.
func (o *SandboxDataTrafficStats) HasUploadByteSizeTotal() bool {
	if o != nil && o.UploadByteSizeTotal != nil {
		return true
	}

	return false
}

// SetUploadByteSizeTotal gets a reference to the given int64 and assigns it to the UploadByteSizeTotal field.
func (o *SandboxDataTrafficStats) SetUploadByteSizeTotal(v int64) {
	o.UploadByteSizeTotal = &v
}

// GetUploadPacketSizeTotal returns the UploadPacketSizeTotal field value if set, zero value otherwise.
func (o *SandboxDataTrafficStats) GetUploadPacketSizeTotal() int64 {
	if o == nil || o.UploadPacketSizeTotal == nil {
		var ret int64
		return ret
	}
	return *o.UploadPacketSizeTotal
}

// GetUploadPacketSizeTotalOk returns a tuple with the UploadPacketSizeTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxDataTrafficStats) GetUploadPacketSizeTotalOk() (*int64, bool) {
	if o == nil || o.UploadPacketSizeTotal == nil {
		return nil, false
	}
	return o.UploadPacketSizeTotal, true
}

// HasUploadPacketSizeTotal returns a boolean if a field has been set.
func (o *SandboxDataTrafficStats) HasUploadPacketSizeTotal() bool {
	if o != nil && o.UploadPacketSizeTotal != nil {
		return true
	}

	return false
}

// SetUploadPacketSizeTotal gets a reference to the given int64 and assigns it to the UploadPacketSizeTotal field.
func (o *SandboxDataTrafficStats) SetUploadPacketSizeTotal(v int64) {
	o.UploadPacketSizeTotal = &v
}

func (o SandboxDataTrafficStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DownloadByteSizeTotal != nil {
		toSerialize["downloadByteSizeTotal"] = o.DownloadByteSizeTotal
	}
	if o.DownloadPacketSizeTotal != nil {
		toSerialize["downloadPacketSizeTotal"] = o.DownloadPacketSizeTotal
	}
	if o.UploadByteSizeTotal != nil {
		toSerialize["uploadByteSizeTotal"] = o.UploadByteSizeTotal
	}
	if o.UploadPacketSizeTotal != nil {
		toSerialize["uploadPacketSizeTotal"] = o.UploadPacketSizeTotal
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxDataTrafficStats struct {
	value *SandboxDataTrafficStats
	isSet bool
}

func (v NullableSandboxDataTrafficStats) Get() *SandboxDataTrafficStats {
	return v.value
}

func (v *NullableSandboxDataTrafficStats) Set(val *SandboxDataTrafficStats) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxDataTrafficStats) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxDataTrafficStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxDataTrafficStats(val *SandboxDataTrafficStats) *NullableSandboxDataTrafficStats {
	return &NullableSandboxDataTrafficStats{value: val, isSet: true}
}

func (v NullableSandboxDataTrafficStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxDataTrafficStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

