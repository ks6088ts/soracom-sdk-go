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

// DailyBill struct for DailyBill
type DailyBill struct {
	// 金額
	Amount *int64 `json:"amount,omitempty"`
	// 通貨
	Currency *string `json:"currency,omitempty"`
	// 年月日
	Date *string `json:"date,omitempty"`
}

// NewDailyBill instantiates a new DailyBill object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailyBill() *DailyBill {
	this := DailyBill{}
	return &this
}

// NewDailyBillWithDefaults instantiates a new DailyBill object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailyBillWithDefaults() *DailyBill {
	this := DailyBill{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *DailyBill) GetAmount() int64 {
	if o == nil || o.Amount == nil {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyBill) GetAmountOk() (*int64, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *DailyBill) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *DailyBill) SetAmount(v int64) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *DailyBill) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyBill) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *DailyBill) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *DailyBill) SetCurrency(v string) {
	o.Currency = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *DailyBill) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyBill) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *DailyBill) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *DailyBill) SetDate(v string) {
	o.Date = &v
}

func (o DailyBill) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	return json.Marshal(toSerialize)
}

type NullableDailyBill struct {
	value *DailyBill
	isSet bool
}

func (v NullableDailyBill) Get() *DailyBill {
	return v.value
}

func (v *NullableDailyBill) Set(val *DailyBill) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyBill) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyBill) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyBill(val *DailyBill) *NullableDailyBill {
	return &NullableDailyBill{value: val, isSet: true}
}

func (v NullableDailyBill) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailyBill) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

