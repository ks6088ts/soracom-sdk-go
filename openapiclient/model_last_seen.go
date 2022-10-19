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

// LastSeen struct for LastSeen
type LastSeen struct {
	Rssi *int32 `json:"rssi,omitempty"`
	Snr *int32 `json:"snr,omitempty"`
	Time *time.Time `json:"time,omitempty"`
}

// NewLastSeen instantiates a new LastSeen object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLastSeen() *LastSeen {
	this := LastSeen{}
	return &this
}

// NewLastSeenWithDefaults instantiates a new LastSeen object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLastSeenWithDefaults() *LastSeen {
	this := LastSeen{}
	return &this
}

// GetRssi returns the Rssi field value if set, zero value otherwise.
func (o *LastSeen) GetRssi() int32 {
	if o == nil || o.Rssi == nil {
		var ret int32
		return ret
	}
	return *o.Rssi
}

// GetRssiOk returns a tuple with the Rssi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastSeen) GetRssiOk() (*int32, bool) {
	if o == nil || o.Rssi == nil {
		return nil, false
	}
	return o.Rssi, true
}

// HasRssi returns a boolean if a field has been set.
func (o *LastSeen) HasRssi() bool {
	if o != nil && o.Rssi != nil {
		return true
	}

	return false
}

// SetRssi gets a reference to the given int32 and assigns it to the Rssi field.
func (o *LastSeen) SetRssi(v int32) {
	o.Rssi = &v
}

// GetSnr returns the Snr field value if set, zero value otherwise.
func (o *LastSeen) GetSnr() int32 {
	if o == nil || o.Snr == nil {
		var ret int32
		return ret
	}
	return *o.Snr
}

// GetSnrOk returns a tuple with the Snr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastSeen) GetSnrOk() (*int32, bool) {
	if o == nil || o.Snr == nil {
		return nil, false
	}
	return o.Snr, true
}

// HasSnr returns a boolean if a field has been set.
func (o *LastSeen) HasSnr() bool {
	if o != nil && o.Snr != nil {
		return true
	}

	return false
}

// SetSnr gets a reference to the given int32 and assigns it to the Snr field.
func (o *LastSeen) SetSnr(v int32) {
	o.Snr = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *LastSeen) GetTime() time.Time {
	if o == nil || o.Time == nil {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastSeen) GetTimeOk() (*time.Time, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *LastSeen) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *LastSeen) SetTime(v time.Time) {
	o.Time = &v
}

func (o LastSeen) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rssi != nil {
		toSerialize["rssi"] = o.Rssi
	}
	if o.Snr != nil {
		toSerialize["snr"] = o.Snr
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableLastSeen struct {
	value *LastSeen
	isSet bool
}

func (v NullableLastSeen) Get() *LastSeen {
	return v.value
}

func (v *NullableLastSeen) Set(val *LastSeen) {
	v.value = val
	v.isSet = true
}

func (v NullableLastSeen) IsSet() bool {
	return v.isSet
}

func (v *NullableLastSeen) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLastSeen(val *LastSeen) *NullableLastSeen {
	return &NullableLastSeen{value: val, isSet: true}
}

func (v NullableLastSeen) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLastSeen) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


