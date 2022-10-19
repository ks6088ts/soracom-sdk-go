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

// CreateEstimatedVolumeDiscountRequest struct for CreateEstimatedVolumeDiscountRequest
type CreateEstimatedVolumeDiscountRequest struct {
	// 契約月数
	ContractTermMonth int32 `json:"contractTermMonth"`
	// 数量
	Quantity int32 `json:"quantity"`
	// 適用開始日
	StartDate *string `json:"startDate,omitempty"`
	// お支払い方法
	VolumeDiscountPaymentType string `json:"volumeDiscountPaymentType"`
	// 料金計算方法
	VolumeDiscountType string `json:"volumeDiscountType"`
}

// NewCreateEstimatedVolumeDiscountRequest instantiates a new CreateEstimatedVolumeDiscountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEstimatedVolumeDiscountRequest(contractTermMonth int32, quantity int32, volumeDiscountPaymentType string, volumeDiscountType string) *CreateEstimatedVolumeDiscountRequest {
	this := CreateEstimatedVolumeDiscountRequest{}
	this.ContractTermMonth = contractTermMonth
	this.Quantity = quantity
	this.VolumeDiscountPaymentType = volumeDiscountPaymentType
	this.VolumeDiscountType = volumeDiscountType
	return &this
}

// NewCreateEstimatedVolumeDiscountRequestWithDefaults instantiates a new CreateEstimatedVolumeDiscountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEstimatedVolumeDiscountRequestWithDefaults() *CreateEstimatedVolumeDiscountRequest {
	this := CreateEstimatedVolumeDiscountRequest{}
	var contractTermMonth int32 = 12
	this.ContractTermMonth = contractTermMonth
	return &this
}

// GetContractTermMonth returns the ContractTermMonth field value
func (o *CreateEstimatedVolumeDiscountRequest) GetContractTermMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ContractTermMonth
}

// GetContractTermMonthOk returns a tuple with the ContractTermMonth field value
// and a boolean to check if the value has been set.
func (o *CreateEstimatedVolumeDiscountRequest) GetContractTermMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractTermMonth, true
}

// SetContractTermMonth sets field value
func (o *CreateEstimatedVolumeDiscountRequest) SetContractTermMonth(v int32) {
	o.ContractTermMonth = v
}

// GetQuantity returns the Quantity field value
func (o *CreateEstimatedVolumeDiscountRequest) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *CreateEstimatedVolumeDiscountRequest) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *CreateEstimatedVolumeDiscountRequest) SetQuantity(v int32) {
	o.Quantity = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CreateEstimatedVolumeDiscountRequest) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEstimatedVolumeDiscountRequest) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CreateEstimatedVolumeDiscountRequest) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *CreateEstimatedVolumeDiscountRequest) SetStartDate(v string) {
	o.StartDate = &v
}

// GetVolumeDiscountPaymentType returns the VolumeDiscountPaymentType field value
func (o *CreateEstimatedVolumeDiscountRequest) GetVolumeDiscountPaymentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VolumeDiscountPaymentType
}

// GetVolumeDiscountPaymentTypeOk returns a tuple with the VolumeDiscountPaymentType field value
// and a boolean to check if the value has been set.
func (o *CreateEstimatedVolumeDiscountRequest) GetVolumeDiscountPaymentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeDiscountPaymentType, true
}

// SetVolumeDiscountPaymentType sets field value
func (o *CreateEstimatedVolumeDiscountRequest) SetVolumeDiscountPaymentType(v string) {
	o.VolumeDiscountPaymentType = v
}

// GetVolumeDiscountType returns the VolumeDiscountType field value
func (o *CreateEstimatedVolumeDiscountRequest) GetVolumeDiscountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VolumeDiscountType
}

// GetVolumeDiscountTypeOk returns a tuple with the VolumeDiscountType field value
// and a boolean to check if the value has been set.
func (o *CreateEstimatedVolumeDiscountRequest) GetVolumeDiscountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeDiscountType, true
}

// SetVolumeDiscountType sets field value
func (o *CreateEstimatedVolumeDiscountRequest) SetVolumeDiscountType(v string) {
	o.VolumeDiscountType = v
}

func (o CreateEstimatedVolumeDiscountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contractTermMonth"] = o.ContractTermMonth
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if true {
		toSerialize["volumeDiscountPaymentType"] = o.VolumeDiscountPaymentType
	}
	if true {
		toSerialize["volumeDiscountType"] = o.VolumeDiscountType
	}
	return json.Marshal(toSerialize)
}

type NullableCreateEstimatedVolumeDiscountRequest struct {
	value *CreateEstimatedVolumeDiscountRequest
	isSet bool
}

func (v NullableCreateEstimatedVolumeDiscountRequest) Get() *CreateEstimatedVolumeDiscountRequest {
	return v.value
}

func (v *NullableCreateEstimatedVolumeDiscountRequest) Set(val *CreateEstimatedVolumeDiscountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEstimatedVolumeDiscountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEstimatedVolumeDiscountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEstimatedVolumeDiscountRequest(val *CreateEstimatedVolumeDiscountRequest) *NullableCreateEstimatedVolumeDiscountRequest {
	return &NullableCreateEstimatedVolumeDiscountRequest{value: val, isSet: true}
}

func (v NullableCreateEstimatedVolumeDiscountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEstimatedVolumeDiscountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


