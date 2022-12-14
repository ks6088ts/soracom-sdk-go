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

// UpdateSpeedClassRequest struct for UpdateSpeedClassRequest
type UpdateSpeedClassRequest struct {
	// 速度クラス。以下のいずれかを指定します。ただし、サブスクリプションにあわせた速度クラスを指定してください。 - plan01s、plan01s - Low Data Volume、planP1、planX3 X3-5MB、plan-D の場合:     - `s1.minimum`     - `s1.slow`     - `s1.standard`     - `s1.fast`     - `s1.4xfast` - plan-KM1 の場合:     - `t1.standard` - plan-DU の場合:     - `u1.standard`     - `u1.slow` - バーチャル SIM/Subscriber の場合:     - `arc.standard` 
	SpeedClass string `json:"speedClass"`
}

// NewUpdateSpeedClassRequest instantiates a new UpdateSpeedClassRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSpeedClassRequest(speedClass string) *UpdateSpeedClassRequest {
	this := UpdateSpeedClassRequest{}
	this.SpeedClass = speedClass
	return &this
}

// NewUpdateSpeedClassRequestWithDefaults instantiates a new UpdateSpeedClassRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSpeedClassRequestWithDefaults() *UpdateSpeedClassRequest {
	this := UpdateSpeedClassRequest{}
	return &this
}

// GetSpeedClass returns the SpeedClass field value
func (o *UpdateSpeedClassRequest) GetSpeedClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpeedClass
}

// GetSpeedClassOk returns a tuple with the SpeedClass field value
// and a boolean to check if the value has been set.
func (o *UpdateSpeedClassRequest) GetSpeedClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpeedClass, true
}

// SetSpeedClass sets field value
func (o *UpdateSpeedClassRequest) SetSpeedClass(v string) {
	o.SpeedClass = v
}

func (o UpdateSpeedClassRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["speedClass"] = o.SpeedClass
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSpeedClassRequest struct {
	value *UpdateSpeedClassRequest
	isSet bool
}

func (v NullableUpdateSpeedClassRequest) Get() *UpdateSpeedClassRequest {
	return v.value
}

func (v *NullableUpdateSpeedClassRequest) Set(val *UpdateSpeedClassRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSpeedClassRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSpeedClassRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSpeedClassRequest(val *UpdateSpeedClassRequest) *NullableUpdateSpeedClassRequest {
	return &NullableUpdateSpeedClassRequest{value: val, isSet: true}
}

func (v NullableUpdateSpeedClassRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSpeedClassRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


