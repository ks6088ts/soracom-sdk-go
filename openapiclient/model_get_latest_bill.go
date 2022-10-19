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

// GetLatestBill struct for GetLatestBill
type GetLatestBill struct {
	// 金額
	Amount *int64 `json:"amount,omitempty"`
	// 通貨
	Currency *string `json:"currency,omitempty"`
	// 利用額計算を実施した時間
	LastEvaluatedTime *string `json:"lastEvaluatedTime,omitempty"`
}

// NewGetLatestBill instantiates a new GetLatestBill object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLatestBill() *GetLatestBill {
	this := GetLatestBill{}
	return &this
}

// NewGetLatestBillWithDefaults instantiates a new GetLatestBill object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLatestBillWithDefaults() *GetLatestBill {
	this := GetLatestBill{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetLatestBill) GetAmount() int64 {
	if o == nil || o.Amount == nil {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLatestBill) GetAmountOk() (*int64, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetLatestBill) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *GetLatestBill) SetAmount(v int64) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *GetLatestBill) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLatestBill) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *GetLatestBill) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *GetLatestBill) SetCurrency(v string) {
	o.Currency = &v
}

// GetLastEvaluatedTime returns the LastEvaluatedTime field value if set, zero value otherwise.
func (o *GetLatestBill) GetLastEvaluatedTime() string {
	if o == nil || o.LastEvaluatedTime == nil {
		var ret string
		return ret
	}
	return *o.LastEvaluatedTime
}

// GetLastEvaluatedTimeOk returns a tuple with the LastEvaluatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLatestBill) GetLastEvaluatedTimeOk() (*string, bool) {
	if o == nil || o.LastEvaluatedTime == nil {
		return nil, false
	}
	return o.LastEvaluatedTime, true
}

// HasLastEvaluatedTime returns a boolean if a field has been set.
func (o *GetLatestBill) HasLastEvaluatedTime() bool {
	if o != nil && o.LastEvaluatedTime != nil {
		return true
	}

	return false
}

// SetLastEvaluatedTime gets a reference to the given string and assigns it to the LastEvaluatedTime field.
func (o *GetLatestBill) SetLastEvaluatedTime(v string) {
	o.LastEvaluatedTime = &v
}

func (o GetLatestBill) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.LastEvaluatedTime != nil {
		toSerialize["lastEvaluatedTime"] = o.LastEvaluatedTime
	}
	return json.Marshal(toSerialize)
}

type NullableGetLatestBill struct {
	value *GetLatestBill
	isSet bool
}

func (v NullableGetLatestBill) Get() *GetLatestBill {
	return v.value
}

func (v *NullableGetLatestBill) Set(val *GetLatestBill) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLatestBill) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLatestBill) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLatestBill(val *GetLatestBill) *NullableGetLatestBill {
	return &NullableGetLatestBill{value: val, isSet: true}
}

func (v NullableGetLatestBill) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLatestBill) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


