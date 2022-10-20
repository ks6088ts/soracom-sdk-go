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

// JunctionMirroringPeer struct for JunctionMirroringPeer
type JunctionMirroringPeer struct {
	Description *string `json:"description,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	IpAddress *string `json:"ipAddress,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
}

// NewJunctionMirroringPeer instantiates a new JunctionMirroringPeer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJunctionMirroringPeer() *JunctionMirroringPeer {
	this := JunctionMirroringPeer{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewJunctionMirroringPeerWithDefaults instantiates a new JunctionMirroringPeer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJunctionMirroringPeerWithDefaults() *JunctionMirroringPeer {
	this := JunctionMirroringPeer{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JunctionMirroringPeer) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunctionMirroringPeer) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JunctionMirroringPeer) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JunctionMirroringPeer) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *JunctionMirroringPeer) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunctionMirroringPeer) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *JunctionMirroringPeer) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *JunctionMirroringPeer) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *JunctionMirroringPeer) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunctionMirroringPeer) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *JunctionMirroringPeer) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *JunctionMirroringPeer) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *JunctionMirroringPeer) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunctionMirroringPeer) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *JunctionMirroringPeer) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *JunctionMirroringPeer) SetProtocol(v string) {
	o.Protocol = &v
}

func (o JunctionMirroringPeer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.IpAddress != nil {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	return json.Marshal(toSerialize)
}

type NullableJunctionMirroringPeer struct {
	value *JunctionMirroringPeer
	isSet bool
}

func (v NullableJunctionMirroringPeer) Get() *JunctionMirroringPeer {
	return v.value
}

func (v *NullableJunctionMirroringPeer) Set(val *JunctionMirroringPeer) {
	v.value = val
	v.isSet = true
}

func (v NullableJunctionMirroringPeer) IsSet() bool {
	return v.isSet
}

func (v *NullableJunctionMirroringPeer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJunctionMirroringPeer(val *JunctionMirroringPeer) *NullableJunctionMirroringPeer {
	return &NullableJunctionMirroringPeer{value: val, isSet: true}
}

func (v NullableJunctionMirroringPeer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJunctionMirroringPeer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

