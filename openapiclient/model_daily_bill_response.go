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

// DailyBillResponse struct for DailyBillResponse
type DailyBillResponse struct {
	// 日ごとの利用料金
	BillList []DailyBill `json:"billList,omitempty"`
}

// NewDailyBillResponse instantiates a new DailyBillResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailyBillResponse() *DailyBillResponse {
	this := DailyBillResponse{}
	return &this
}

// NewDailyBillResponseWithDefaults instantiates a new DailyBillResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailyBillResponseWithDefaults() *DailyBillResponse {
	this := DailyBillResponse{}
	return &this
}

// GetBillList returns the BillList field value if set, zero value otherwise.
func (o *DailyBillResponse) GetBillList() []DailyBill {
	if o == nil || o.BillList == nil {
		var ret []DailyBill
		return ret
	}
	return o.BillList
}

// GetBillListOk returns a tuple with the BillList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyBillResponse) GetBillListOk() ([]DailyBill, bool) {
	if o == nil || o.BillList == nil {
		return nil, false
	}
	return o.BillList, true
}

// HasBillList returns a boolean if a field has been set.
func (o *DailyBillResponse) HasBillList() bool {
	if o != nil && o.BillList != nil {
		return true
	}

	return false
}

// SetBillList gets a reference to the given []DailyBill and assigns it to the BillList field.
func (o *DailyBillResponse) SetBillList(v []DailyBill) {
	o.BillList = v
}

func (o DailyBillResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BillList != nil {
		toSerialize["billList"] = o.BillList
	}
	return json.Marshal(toSerialize)
}

type NullableDailyBillResponse struct {
	value *DailyBillResponse
	isSet bool
}

func (v NullableDailyBillResponse) Get() *DailyBillResponse {
	return v.value
}

func (v *NullableDailyBillResponse) Set(val *DailyBillResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyBillResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyBillResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyBillResponse(val *DailyBillResponse) *NullableDailyBillResponse {
	return &NullableDailyBillResponse{value: val, isSet: true}
}

func (v NullableDailyBillResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailyBillResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

