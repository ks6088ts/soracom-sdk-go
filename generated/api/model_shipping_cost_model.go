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

// ShippingCostModel 発送料
type ShippingCostModel struct {
	// 適用済み配送オプション
	AppliedShippingOptions []map[string]interface{} `json:"appliedShippingOptions,omitempty"`
	// 通貨
	Currency *string `json:"currency,omitempty"`
	// (日本カバレッジのみ) 発送先エリア
	ShippingArea *string `json:"shippingArea,omitempty"`
	// 発送先エリア名称
	ShippingAreaName *string `json:"shippingAreaName,omitempty"`
	// 送料
	ShippingCost *float64 `json:"shippingCost,omitempty"`
	// 発送サイズ
	Size *int32 `json:"size,omitempty"`
}

// NewShippingCostModel instantiates a new ShippingCostModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingCostModel() *ShippingCostModel {
	this := ShippingCostModel{}
	return &this
}

// NewShippingCostModelWithDefaults instantiates a new ShippingCostModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingCostModelWithDefaults() *ShippingCostModel {
	this := ShippingCostModel{}
	return &this
}

// GetAppliedShippingOptions returns the AppliedShippingOptions field value if set, zero value otherwise.
func (o *ShippingCostModel) GetAppliedShippingOptions() []map[string]interface{} {
	if o == nil || o.AppliedShippingOptions == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.AppliedShippingOptions
}

// GetAppliedShippingOptionsOk returns a tuple with the AppliedShippingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingCostModel) GetAppliedShippingOptionsOk() ([]map[string]interface{}, bool) {
	if o == nil || o.AppliedShippingOptions == nil {
		return nil, false
	}
	return o.AppliedShippingOptions, true
}

// HasAppliedShippingOptions returns a boolean if a field has been set.
func (o *ShippingCostModel) HasAppliedShippingOptions() bool {
	if o != nil && o.AppliedShippingOptions != nil {
		return true
	}

	return false
}

// SetAppliedShippingOptions gets a reference to the given []map[string]interface{} and assigns it to the AppliedShippingOptions field.
func (o *ShippingCostModel) SetAppliedShippingOptions(v []map[string]interface{}) {
	o.AppliedShippingOptions = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ShippingCostModel) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingCostModel) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ShippingCostModel) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ShippingCostModel) SetCurrency(v string) {
	o.Currency = &v
}

// GetShippingArea returns the ShippingArea field value if set, zero value otherwise.
func (o *ShippingCostModel) GetShippingArea() string {
	if o == nil || o.ShippingArea == nil {
		var ret string
		return ret
	}
	return *o.ShippingArea
}

// GetShippingAreaOk returns a tuple with the ShippingArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingCostModel) GetShippingAreaOk() (*string, bool) {
	if o == nil || o.ShippingArea == nil {
		return nil, false
	}
	return o.ShippingArea, true
}

// HasShippingArea returns a boolean if a field has been set.
func (o *ShippingCostModel) HasShippingArea() bool {
	if o != nil && o.ShippingArea != nil {
		return true
	}

	return false
}

// SetShippingArea gets a reference to the given string and assigns it to the ShippingArea field.
func (o *ShippingCostModel) SetShippingArea(v string) {
	o.ShippingArea = &v
}

// GetShippingAreaName returns the ShippingAreaName field value if set, zero value otherwise.
func (o *ShippingCostModel) GetShippingAreaName() string {
	if o == nil || o.ShippingAreaName == nil {
		var ret string
		return ret
	}
	return *o.ShippingAreaName
}

// GetShippingAreaNameOk returns a tuple with the ShippingAreaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingCostModel) GetShippingAreaNameOk() (*string, bool) {
	if o == nil || o.ShippingAreaName == nil {
		return nil, false
	}
	return o.ShippingAreaName, true
}

// HasShippingAreaName returns a boolean if a field has been set.
func (o *ShippingCostModel) HasShippingAreaName() bool {
	if o != nil && o.ShippingAreaName != nil {
		return true
	}

	return false
}

// SetShippingAreaName gets a reference to the given string and assigns it to the ShippingAreaName field.
func (o *ShippingCostModel) SetShippingAreaName(v string) {
	o.ShippingAreaName = &v
}

// GetShippingCost returns the ShippingCost field value if set, zero value otherwise.
func (o *ShippingCostModel) GetShippingCost() float64 {
	if o == nil || o.ShippingCost == nil {
		var ret float64
		return ret
	}
	return *o.ShippingCost
}

// GetShippingCostOk returns a tuple with the ShippingCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingCostModel) GetShippingCostOk() (*float64, bool) {
	if o == nil || o.ShippingCost == nil {
		return nil, false
	}
	return o.ShippingCost, true
}

// HasShippingCost returns a boolean if a field has been set.
func (o *ShippingCostModel) HasShippingCost() bool {
	if o != nil && o.ShippingCost != nil {
		return true
	}

	return false
}

// SetShippingCost gets a reference to the given float64 and assigns it to the ShippingCost field.
func (o *ShippingCostModel) SetShippingCost(v float64) {
	o.ShippingCost = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ShippingCostModel) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingCostModel) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ShippingCostModel) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *ShippingCostModel) SetSize(v int32) {
	o.Size = &v
}

func (o ShippingCostModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppliedShippingOptions != nil {
		toSerialize["appliedShippingOptions"] = o.AppliedShippingOptions
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.ShippingArea != nil {
		toSerialize["shippingArea"] = o.ShippingArea
	}
	if o.ShippingAreaName != nil {
		toSerialize["shippingAreaName"] = o.ShippingAreaName
	}
	if o.ShippingCost != nil {
		toSerialize["shippingCost"] = o.ShippingCost
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableShippingCostModel struct {
	value *ShippingCostModel
	isSet bool
}

func (v NullableShippingCostModel) Get() *ShippingCostModel {
	return v.value
}

func (v *NullableShippingCostModel) Set(val *ShippingCostModel) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingCostModel) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingCostModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingCostModel(val *ShippingCostModel) *NullableShippingCostModel {
	return &NullableShippingCostModel{value: val, isSet: true}
}

func (v NullableShippingCostModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingCostModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

