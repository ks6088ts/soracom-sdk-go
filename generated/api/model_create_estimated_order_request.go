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

// CreateEstimatedOrderRequest struct for CreateEstimatedOrderRequest
type CreateEstimatedOrderRequest struct {
	// 発注商品リスト
	OrderItemList []OrderItemModel `json:"orderItemList,omitempty"`
	// 商品発送先 ID
	ShippingAddressId *string `json:"shippingAddressId,omitempty"`
}

// NewCreateEstimatedOrderRequest instantiates a new CreateEstimatedOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEstimatedOrderRequest() *CreateEstimatedOrderRequest {
	this := CreateEstimatedOrderRequest{}
	return &this
}

// NewCreateEstimatedOrderRequestWithDefaults instantiates a new CreateEstimatedOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEstimatedOrderRequestWithDefaults() *CreateEstimatedOrderRequest {
	this := CreateEstimatedOrderRequest{}
	return &this
}

// GetOrderItemList returns the OrderItemList field value if set, zero value otherwise.
func (o *CreateEstimatedOrderRequest) GetOrderItemList() []OrderItemModel {
	if o == nil || o.OrderItemList == nil {
		var ret []OrderItemModel
		return ret
	}
	return o.OrderItemList
}

// GetOrderItemListOk returns a tuple with the OrderItemList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEstimatedOrderRequest) GetOrderItemListOk() ([]OrderItemModel, bool) {
	if o == nil || o.OrderItemList == nil {
		return nil, false
	}
	return o.OrderItemList, true
}

// HasOrderItemList returns a boolean if a field has been set.
func (o *CreateEstimatedOrderRequest) HasOrderItemList() bool {
	if o != nil && o.OrderItemList != nil {
		return true
	}

	return false
}

// SetOrderItemList gets a reference to the given []OrderItemModel and assigns it to the OrderItemList field.
func (o *CreateEstimatedOrderRequest) SetOrderItemList(v []OrderItemModel) {
	o.OrderItemList = v
}

// GetShippingAddressId returns the ShippingAddressId field value if set, zero value otherwise.
func (o *CreateEstimatedOrderRequest) GetShippingAddressId() string {
	if o == nil || o.ShippingAddressId == nil {
		var ret string
		return ret
	}
	return *o.ShippingAddressId
}

// GetShippingAddressIdOk returns a tuple with the ShippingAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEstimatedOrderRequest) GetShippingAddressIdOk() (*string, bool) {
	if o == nil || o.ShippingAddressId == nil {
		return nil, false
	}
	return o.ShippingAddressId, true
}

// HasShippingAddressId returns a boolean if a field has been set.
func (o *CreateEstimatedOrderRequest) HasShippingAddressId() bool {
	if o != nil && o.ShippingAddressId != nil {
		return true
	}

	return false
}

// SetShippingAddressId gets a reference to the given string and assigns it to the ShippingAddressId field.
func (o *CreateEstimatedOrderRequest) SetShippingAddressId(v string) {
	o.ShippingAddressId = &v
}

func (o CreateEstimatedOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OrderItemList != nil {
		toSerialize["orderItemList"] = o.OrderItemList
	}
	if o.ShippingAddressId != nil {
		toSerialize["shippingAddressId"] = o.ShippingAddressId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateEstimatedOrderRequest struct {
	value *CreateEstimatedOrderRequest
	isSet bool
}

func (v NullableCreateEstimatedOrderRequest) Get() *CreateEstimatedOrderRequest {
	return v.value
}

func (v *NullableCreateEstimatedOrderRequest) Set(val *CreateEstimatedOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEstimatedOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEstimatedOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEstimatedOrderRequest(val *CreateEstimatedOrderRequest) *NullableCreateEstimatedOrderRequest {
	return &NullableCreateEstimatedOrderRequest{value: val, isSet: true}
}

func (v NullableCreateEstimatedOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEstimatedOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


