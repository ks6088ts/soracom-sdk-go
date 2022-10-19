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

// ImeiLock struct for ImeiLock
type ImeiLock struct {
	Imei *string `json:"imei,omitempty"`
}

// NewImeiLock instantiates a new ImeiLock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImeiLock() *ImeiLock {
	this := ImeiLock{}
	return &this
}

// NewImeiLockWithDefaults instantiates a new ImeiLock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImeiLockWithDefaults() *ImeiLock {
	this := ImeiLock{}
	return &this
}

// GetImei returns the Imei field value if set, zero value otherwise.
func (o *ImeiLock) GetImei() string {
	if o == nil || o.Imei == nil {
		var ret string
		return ret
	}
	return *o.Imei
}

// GetImeiOk returns a tuple with the Imei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImeiLock) GetImeiOk() (*string, bool) {
	if o == nil || o.Imei == nil {
		return nil, false
	}
	return o.Imei, true
}

// HasImei returns a boolean if a field has been set.
func (o *ImeiLock) HasImei() bool {
	if o != nil && o.Imei != nil {
		return true
	}

	return false
}

// SetImei gets a reference to the given string and assigns it to the Imei field.
func (o *ImeiLock) SetImei(v string) {
	o.Imei = &v
}

func (o ImeiLock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Imei != nil {
		toSerialize["imei"] = o.Imei
	}
	return json.Marshal(toSerialize)
}

type NullableImeiLock struct {
	value *ImeiLock
	isSet bool
}

func (v NullableImeiLock) Get() *ImeiLock {
	return v.value
}

func (v *NullableImeiLock) Set(val *ImeiLock) {
	v.value = val
	v.isSet = true
}

func (v NullableImeiLock) IsSet() bool {
	return v.isSet
}

func (v *NullableImeiLock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImeiLock(val *ImeiLock) *NullableImeiLock {
	return &NullableImeiLock{value: val, isSet: true}
}

func (v NullableImeiLock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImeiLock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


