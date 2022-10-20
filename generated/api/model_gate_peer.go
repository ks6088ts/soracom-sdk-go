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

// GatePeer struct for GatePeer
type GatePeer struct {
	InnerIpAddress *string `json:"innerIpAddress,omitempty"`
	Netmask *string `json:"netmask,omitempty"`
	OuterIpAddress *string `json:"outerIpAddress,omitempty"`
	OwnedByCustomer *bool `json:"ownedByCustomer,omitempty"`
}

// NewGatePeer instantiates a new GatePeer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatePeer() *GatePeer {
	this := GatePeer{}
	return &this
}

// NewGatePeerWithDefaults instantiates a new GatePeer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatePeerWithDefaults() *GatePeer {
	this := GatePeer{}
	return &this
}

// GetInnerIpAddress returns the InnerIpAddress field value if set, zero value otherwise.
func (o *GatePeer) GetInnerIpAddress() string {
	if o == nil || o.InnerIpAddress == nil {
		var ret string
		return ret
	}
	return *o.InnerIpAddress
}

// GetInnerIpAddressOk returns a tuple with the InnerIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatePeer) GetInnerIpAddressOk() (*string, bool) {
	if o == nil || o.InnerIpAddress == nil {
		return nil, false
	}
	return o.InnerIpAddress, true
}

// HasInnerIpAddress returns a boolean if a field has been set.
func (o *GatePeer) HasInnerIpAddress() bool {
	if o != nil && o.InnerIpAddress != nil {
		return true
	}

	return false
}

// SetInnerIpAddress gets a reference to the given string and assigns it to the InnerIpAddress field.
func (o *GatePeer) SetInnerIpAddress(v string) {
	o.InnerIpAddress = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *GatePeer) GetNetmask() string {
	if o == nil || o.Netmask == nil {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatePeer) GetNetmaskOk() (*string, bool) {
	if o == nil || o.Netmask == nil {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *GatePeer) HasNetmask() bool {
	if o != nil && o.Netmask != nil {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *GatePeer) SetNetmask(v string) {
	o.Netmask = &v
}

// GetOuterIpAddress returns the OuterIpAddress field value if set, zero value otherwise.
func (o *GatePeer) GetOuterIpAddress() string {
	if o == nil || o.OuterIpAddress == nil {
		var ret string
		return ret
	}
	return *o.OuterIpAddress
}

// GetOuterIpAddressOk returns a tuple with the OuterIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatePeer) GetOuterIpAddressOk() (*string, bool) {
	if o == nil || o.OuterIpAddress == nil {
		return nil, false
	}
	return o.OuterIpAddress, true
}

// HasOuterIpAddress returns a boolean if a field has been set.
func (o *GatePeer) HasOuterIpAddress() bool {
	if o != nil && o.OuterIpAddress != nil {
		return true
	}

	return false
}

// SetOuterIpAddress gets a reference to the given string and assigns it to the OuterIpAddress field.
func (o *GatePeer) SetOuterIpAddress(v string) {
	o.OuterIpAddress = &v
}

// GetOwnedByCustomer returns the OwnedByCustomer field value if set, zero value otherwise.
func (o *GatePeer) GetOwnedByCustomer() bool {
	if o == nil || o.OwnedByCustomer == nil {
		var ret bool
		return ret
	}
	return *o.OwnedByCustomer
}

// GetOwnedByCustomerOk returns a tuple with the OwnedByCustomer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatePeer) GetOwnedByCustomerOk() (*bool, bool) {
	if o == nil || o.OwnedByCustomer == nil {
		return nil, false
	}
	return o.OwnedByCustomer, true
}

// HasOwnedByCustomer returns a boolean if a field has been set.
func (o *GatePeer) HasOwnedByCustomer() bool {
	if o != nil && o.OwnedByCustomer != nil {
		return true
	}

	return false
}

// SetOwnedByCustomer gets a reference to the given bool and assigns it to the OwnedByCustomer field.
func (o *GatePeer) SetOwnedByCustomer(v bool) {
	o.OwnedByCustomer = &v
}

func (o GatePeer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InnerIpAddress != nil {
		toSerialize["innerIpAddress"] = o.InnerIpAddress
	}
	if o.Netmask != nil {
		toSerialize["netmask"] = o.Netmask
	}
	if o.OuterIpAddress != nil {
		toSerialize["outerIpAddress"] = o.OuterIpAddress
	}
	if o.OwnedByCustomer != nil {
		toSerialize["ownedByCustomer"] = o.OwnedByCustomer
	}
	return json.Marshal(toSerialize)
}

type NullableGatePeer struct {
	value *GatePeer
	isSet bool
}

func (v NullableGatePeer) Get() *GatePeer {
	return v.value
}

func (v *NullableGatePeer) Set(val *GatePeer) {
	v.value = val
	v.isSet = true
}

func (v NullableGatePeer) IsSet() bool {
	return v.isSet
}

func (v *NullableGatePeer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatePeer(val *GatePeer) *NullableGatePeer {
	return &NullableGatePeer{value: val, isSet: true}
}

func (v NullableGatePeer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatePeer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

