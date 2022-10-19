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

// PacketCaptureSessionRequest struct for PacketCaptureSessionRequest
type PacketCaptureSessionRequest struct {
	Duration int32 `json:"duration"`
	Prefix *string `json:"prefix,omitempty"`
}

// NewPacketCaptureSessionRequest instantiates a new PacketCaptureSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPacketCaptureSessionRequest(duration int32) *PacketCaptureSessionRequest {
	this := PacketCaptureSessionRequest{}
	this.Duration = duration
	return &this
}

// NewPacketCaptureSessionRequestWithDefaults instantiates a new PacketCaptureSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPacketCaptureSessionRequestWithDefaults() *PacketCaptureSessionRequest {
	this := PacketCaptureSessionRequest{}
	return &this
}

// GetDuration returns the Duration field value
func (o *PacketCaptureSessionRequest) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *PacketCaptureSessionRequest) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *PacketCaptureSessionRequest) SetDuration(v int32) {
	o.Duration = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *PacketCaptureSessionRequest) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PacketCaptureSessionRequest) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *PacketCaptureSessionRequest) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *PacketCaptureSessionRequest) SetPrefix(v string) {
	o.Prefix = &v
}

func (o PacketCaptureSessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["duration"] = o.Duration
	}
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}
	return json.Marshal(toSerialize)
}

type NullablePacketCaptureSessionRequest struct {
	value *PacketCaptureSessionRequest
	isSet bool
}

func (v NullablePacketCaptureSessionRequest) Get() *PacketCaptureSessionRequest {
	return v.value
}

func (v *NullablePacketCaptureSessionRequest) Set(val *PacketCaptureSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePacketCaptureSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePacketCaptureSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePacketCaptureSessionRequest(val *PacketCaptureSessionRequest) *NullablePacketCaptureSessionRequest {
	return &NullablePacketCaptureSessionRequest{value: val, isSet: true}
}

func (v NullablePacketCaptureSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePacketCaptureSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


