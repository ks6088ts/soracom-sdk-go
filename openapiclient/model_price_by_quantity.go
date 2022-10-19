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

// PriceByQuantity struct for PriceByQuantity
type PriceByQuantity struct {
	// 最低数量
	MinQuantity *int32 `json:"minQuantity,omitempty"`
	// 販売価格
	Price *float64 `json:"price,omitempty"`
	// 税込販売価格
	TaxIncludedPrice *float64 `json:"taxIncludedPrice,omitempty"`
}

// NewPriceByQuantity instantiates a new PriceByQuantity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceByQuantity() *PriceByQuantity {
	this := PriceByQuantity{}
	return &this
}

// NewPriceByQuantityWithDefaults instantiates a new PriceByQuantity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceByQuantityWithDefaults() *PriceByQuantity {
	this := PriceByQuantity{}
	return &this
}

// GetMinQuantity returns the MinQuantity field value if set, zero value otherwise.
func (o *PriceByQuantity) GetMinQuantity() int32 {
	if o == nil || o.MinQuantity == nil {
		var ret int32
		return ret
	}
	return *o.MinQuantity
}

// GetMinQuantityOk returns a tuple with the MinQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceByQuantity) GetMinQuantityOk() (*int32, bool) {
	if o == nil || o.MinQuantity == nil {
		return nil, false
	}
	return o.MinQuantity, true
}

// HasMinQuantity returns a boolean if a field has been set.
func (o *PriceByQuantity) HasMinQuantity() bool {
	if o != nil && o.MinQuantity != nil {
		return true
	}

	return false
}

// SetMinQuantity gets a reference to the given int32 and assigns it to the MinQuantity field.
func (o *PriceByQuantity) SetMinQuantity(v int32) {
	o.MinQuantity = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PriceByQuantity) GetPrice() float64 {
	if o == nil || o.Price == nil {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceByQuantity) GetPriceOk() (*float64, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PriceByQuantity) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *PriceByQuantity) SetPrice(v float64) {
	o.Price = &v
}

// GetTaxIncludedPrice returns the TaxIncludedPrice field value if set, zero value otherwise.
func (o *PriceByQuantity) GetTaxIncludedPrice() float64 {
	if o == nil || o.TaxIncludedPrice == nil {
		var ret float64
		return ret
	}
	return *o.TaxIncludedPrice
}

// GetTaxIncludedPriceOk returns a tuple with the TaxIncludedPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceByQuantity) GetTaxIncludedPriceOk() (*float64, bool) {
	if o == nil || o.TaxIncludedPrice == nil {
		return nil, false
	}
	return o.TaxIncludedPrice, true
}

// HasTaxIncludedPrice returns a boolean if a field has been set.
func (o *PriceByQuantity) HasTaxIncludedPrice() bool {
	if o != nil && o.TaxIncludedPrice != nil {
		return true
	}

	return false
}

// SetTaxIncludedPrice gets a reference to the given float64 and assigns it to the TaxIncludedPrice field.
func (o *PriceByQuantity) SetTaxIncludedPrice(v float64) {
	o.TaxIncludedPrice = &v
}

func (o PriceByQuantity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MinQuantity != nil {
		toSerialize["minQuantity"] = o.MinQuantity
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.TaxIncludedPrice != nil {
		toSerialize["taxIncludedPrice"] = o.TaxIncludedPrice
	}
	return json.Marshal(toSerialize)
}

type NullablePriceByQuantity struct {
	value *PriceByQuantity
	isSet bool
}

func (v NullablePriceByQuantity) Get() *PriceByQuantity {
	return v.value
}

func (v *NullablePriceByQuantity) Set(val *PriceByQuantity) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceByQuantity) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceByQuantity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceByQuantity(val *PriceByQuantity) *NullablePriceByQuantity {
	return &NullablePriceByQuantity{value: val, isSet: true}
}

func (v NullablePriceByQuantity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceByQuantity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


