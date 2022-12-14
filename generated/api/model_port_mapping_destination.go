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

// PortMappingDestination struct for PortMappingDestination
type PortMappingDestination struct {
	// The target IMSI of the subscriber.
	Imsi string `json:"imsi"`
	// The port on your device used for access.
	Port float32 `json:"port"`
}

// NewPortMappingDestination instantiates a new PortMappingDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortMappingDestination(imsi string, port float32) *PortMappingDestination {
	this := PortMappingDestination{}
	this.Imsi = imsi
	this.Port = port
	return &this
}

// NewPortMappingDestinationWithDefaults instantiates a new PortMappingDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortMappingDestinationWithDefaults() *PortMappingDestination {
	this := PortMappingDestination{}
	return &this
}

// GetImsi returns the Imsi field value
func (o *PortMappingDestination) GetImsi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Imsi
}

// GetImsiOk returns a tuple with the Imsi field value
// and a boolean to check if the value has been set.
func (o *PortMappingDestination) GetImsiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Imsi, true
}

// SetImsi sets field value
func (o *PortMappingDestination) SetImsi(v string) {
	o.Imsi = v
}

// GetPort returns the Port field value
func (o *PortMappingDestination) GetPort() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *PortMappingDestination) GetPortOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *PortMappingDestination) SetPort(v float32) {
	o.Port = v
}

func (o PortMappingDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["imsi"] = o.Imsi
	}
	if true {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullablePortMappingDestination struct {
	value *PortMappingDestination
	isSet bool
}

func (v NullablePortMappingDestination) Get() *PortMappingDestination {
	return v.value
}

func (v *NullablePortMappingDestination) Set(val *PortMappingDestination) {
	v.value = val
	v.isSet = true
}

func (v NullablePortMappingDestination) IsSet() bool {
	return v.isSet
}

func (v *NullablePortMappingDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortMappingDestination(val *PortMappingDestination) *NullablePortMappingDestination {
	return &NullablePortMappingDestination{value: val, isSet: true}
}

func (v NullablePortMappingDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortMappingDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


