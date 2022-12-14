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

// GlobalSimAppletPLMNRecord struct for GlobalSimAppletPLMNRecord
type GlobalSimAppletPLMNRecord struct {
	ContainerId int64 `json:"containerId"`
	Mcc string `json:"mcc"`
	Mnc *string `json:"mnc,omitempty"`
}

// NewGlobalSimAppletPLMNRecord instantiates a new GlobalSimAppletPLMNRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalSimAppletPLMNRecord(containerId int64, mcc string) *GlobalSimAppletPLMNRecord {
	this := GlobalSimAppletPLMNRecord{}
	this.ContainerId = containerId
	this.Mcc = mcc
	return &this
}

// NewGlobalSimAppletPLMNRecordWithDefaults instantiates a new GlobalSimAppletPLMNRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalSimAppletPLMNRecordWithDefaults() *GlobalSimAppletPLMNRecord {
	this := GlobalSimAppletPLMNRecord{}
	return &this
}

// GetContainerId returns the ContainerId field value
func (o *GlobalSimAppletPLMNRecord) GetContainerId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value
// and a boolean to check if the value has been set.
func (o *GlobalSimAppletPLMNRecord) GetContainerIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerId, true
}

// SetContainerId sets field value
func (o *GlobalSimAppletPLMNRecord) SetContainerId(v int64) {
	o.ContainerId = v
}

// GetMcc returns the Mcc field value
func (o *GlobalSimAppletPLMNRecord) GetMcc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value
// and a boolean to check if the value has been set.
func (o *GlobalSimAppletPLMNRecord) GetMccOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mcc, true
}

// SetMcc sets field value
func (o *GlobalSimAppletPLMNRecord) SetMcc(v string) {
	o.Mcc = v
}

// GetMnc returns the Mnc field value if set, zero value otherwise.
func (o *GlobalSimAppletPLMNRecord) GetMnc() string {
	if o == nil || o.Mnc == nil {
		var ret string
		return ret
	}
	return *o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalSimAppletPLMNRecord) GetMncOk() (*string, bool) {
	if o == nil || o.Mnc == nil {
		return nil, false
	}
	return o.Mnc, true
}

// HasMnc returns a boolean if a field has been set.
func (o *GlobalSimAppletPLMNRecord) HasMnc() bool {
	if o != nil && o.Mnc != nil {
		return true
	}

	return false
}

// SetMnc gets a reference to the given string and assigns it to the Mnc field.
func (o *GlobalSimAppletPLMNRecord) SetMnc(v string) {
	o.Mnc = &v
}

func (o GlobalSimAppletPLMNRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["containerId"] = o.ContainerId
	}
	if true {
		toSerialize["mcc"] = o.Mcc
	}
	if o.Mnc != nil {
		toSerialize["mnc"] = o.Mnc
	}
	return json.Marshal(toSerialize)
}

type NullableGlobalSimAppletPLMNRecord struct {
	value *GlobalSimAppletPLMNRecord
	isSet bool
}

func (v NullableGlobalSimAppletPLMNRecord) Get() *GlobalSimAppletPLMNRecord {
	return v.value
}

func (v *NullableGlobalSimAppletPLMNRecord) Set(val *GlobalSimAppletPLMNRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalSimAppletPLMNRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalSimAppletPLMNRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalSimAppletPLMNRecord(val *GlobalSimAppletPLMNRecord) *NullableGlobalSimAppletPLMNRecord {
	return &NullableGlobalSimAppletPLMNRecord{value: val, isSet: true}
}

func (v NullableGlobalSimAppletPLMNRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalSimAppletPLMNRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


