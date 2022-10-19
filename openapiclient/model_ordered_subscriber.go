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

// OrderedSubscriber struct for OrderedSubscriber
type OrderedSubscriber struct {
	// IMSI
	Imsi *string `json:"imsi,omitempty"`
	// MSISDN
	Msisdn *string `json:"msisdn,omitempty"`
	// serialNumber
	SerialNumber *string `json:"serialNumber,omitempty"`
}

// NewOrderedSubscriber instantiates a new OrderedSubscriber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderedSubscriber() *OrderedSubscriber {
	this := OrderedSubscriber{}
	return &this
}

// NewOrderedSubscriberWithDefaults instantiates a new OrderedSubscriber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderedSubscriberWithDefaults() *OrderedSubscriber {
	this := OrderedSubscriber{}
	return &this
}

// GetImsi returns the Imsi field value if set, zero value otherwise.
func (o *OrderedSubscriber) GetImsi() string {
	if o == nil || o.Imsi == nil {
		var ret string
		return ret
	}
	return *o.Imsi
}

// GetImsiOk returns a tuple with the Imsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderedSubscriber) GetImsiOk() (*string, bool) {
	if o == nil || o.Imsi == nil {
		return nil, false
	}
	return o.Imsi, true
}

// HasImsi returns a boolean if a field has been set.
func (o *OrderedSubscriber) HasImsi() bool {
	if o != nil && o.Imsi != nil {
		return true
	}

	return false
}

// SetImsi gets a reference to the given string and assigns it to the Imsi field.
func (o *OrderedSubscriber) SetImsi(v string) {
	o.Imsi = &v
}

// GetMsisdn returns the Msisdn field value if set, zero value otherwise.
func (o *OrderedSubscriber) GetMsisdn() string {
	if o == nil || o.Msisdn == nil {
		var ret string
		return ret
	}
	return *o.Msisdn
}

// GetMsisdnOk returns a tuple with the Msisdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderedSubscriber) GetMsisdnOk() (*string, bool) {
	if o == nil || o.Msisdn == nil {
		return nil, false
	}
	return o.Msisdn, true
}

// HasMsisdn returns a boolean if a field has been set.
func (o *OrderedSubscriber) HasMsisdn() bool {
	if o != nil && o.Msisdn != nil {
		return true
	}

	return false
}

// SetMsisdn gets a reference to the given string and assigns it to the Msisdn field.
func (o *OrderedSubscriber) SetMsisdn(v string) {
	o.Msisdn = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *OrderedSubscriber) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderedSubscriber) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *OrderedSubscriber) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *OrderedSubscriber) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

func (o OrderedSubscriber) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Imsi != nil {
		toSerialize["imsi"] = o.Imsi
	}
	if o.Msisdn != nil {
		toSerialize["msisdn"] = o.Msisdn
	}
	if o.SerialNumber != nil {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	return json.Marshal(toSerialize)
}

type NullableOrderedSubscriber struct {
	value *OrderedSubscriber
	isSet bool
}

func (v NullableOrderedSubscriber) Get() *OrderedSubscriber {
	return v.value
}

func (v *NullableOrderedSubscriber) Set(val *OrderedSubscriber) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderedSubscriber) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderedSubscriber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderedSubscriber(val *OrderedSubscriber) *NullableOrderedSubscriber {
	return &NullableOrderedSubscriber{value: val, isSet: true}
}

func (v NullableOrderedSubscriber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderedSubscriber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


