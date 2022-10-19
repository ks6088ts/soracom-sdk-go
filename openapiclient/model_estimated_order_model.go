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

// EstimatedOrderModel struct for EstimatedOrderModel
type EstimatedOrderModel struct {
	// Applied shipping options
	AppliedShippingOptions []map[string]interface{} `json:"appliedShippingOptions,omitempty"`
	// Balance due
	BalanceDue *float64 `json:"balanceDue,omitempty"`
	// Currency
	Currency *string `json:"currency,omitempty"`
	// メールアドレス
	Email *string `json:"email,omitempty"`
	// Order channel
	OrderChannel *string `json:"orderChannel,omitempty"`
	// 発注 ID
	OrderId *string `json:"orderId,omitempty"`
	// 発注商品リスト
	OrderItemList []EstimatedOrderItemModel `json:"orderItemList,omitempty"`
	// Preferred delivery date
	PreferredDeliveryDate *string `json:"preferredDeliveryDate,omitempty"`
	// Purchase order number
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
	ShippingAddress *ShippingAddressModel `json:"shippingAddress,omitempty"`
	// 商品発送先 ID
	ShippingAddressId *string `json:"shippingAddressId,omitempty"`
	// 配送料
	ShippingCost *float64 `json:"shippingCost,omitempty"`
	// 消費税
	TaxAmount *float64 `json:"taxAmount,omitempty"`
	// Tax included in shipping cost
	TaxIncludedInShippingCost *float64 `json:"taxIncludedInShippingCost,omitempty"`
	// Whether shipping cost is taxable
	TaxOnShippingCost *bool `json:"taxOnShippingCost,omitempty"`
	// 合計金額
	TotalAmount *float64 `json:"totalAmount,omitempty"`
	// Withholding tax amount
	WithholdingTaxAmount *float64 `json:"withholdingTaxAmount,omitempty"`
}

// NewEstimatedOrderModel instantiates a new EstimatedOrderModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimatedOrderModel() *EstimatedOrderModel {
	this := EstimatedOrderModel{}
	return &this
}

// NewEstimatedOrderModelWithDefaults instantiates a new EstimatedOrderModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimatedOrderModelWithDefaults() *EstimatedOrderModel {
	this := EstimatedOrderModel{}
	return &this
}

// GetAppliedShippingOptions returns the AppliedShippingOptions field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetAppliedShippingOptions() []map[string]interface{} {
	if o == nil || o.AppliedShippingOptions == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.AppliedShippingOptions
}

// GetAppliedShippingOptionsOk returns a tuple with the AppliedShippingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetAppliedShippingOptionsOk() ([]map[string]interface{}, bool) {
	if o == nil || o.AppliedShippingOptions == nil {
		return nil, false
	}
	return o.AppliedShippingOptions, true
}

// HasAppliedShippingOptions returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasAppliedShippingOptions() bool {
	if o != nil && o.AppliedShippingOptions != nil {
		return true
	}

	return false
}

// SetAppliedShippingOptions gets a reference to the given []map[string]interface{} and assigns it to the AppliedShippingOptions field.
func (o *EstimatedOrderModel) SetAppliedShippingOptions(v []map[string]interface{}) {
	o.AppliedShippingOptions = v
}

// GetBalanceDue returns the BalanceDue field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetBalanceDue() float64 {
	if o == nil || o.BalanceDue == nil {
		var ret float64
		return ret
	}
	return *o.BalanceDue
}

// GetBalanceDueOk returns a tuple with the BalanceDue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetBalanceDueOk() (*float64, bool) {
	if o == nil || o.BalanceDue == nil {
		return nil, false
	}
	return o.BalanceDue, true
}

// HasBalanceDue returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasBalanceDue() bool {
	if o != nil && o.BalanceDue != nil {
		return true
	}

	return false
}

// SetBalanceDue gets a reference to the given float64 and assigns it to the BalanceDue field.
func (o *EstimatedOrderModel) SetBalanceDue(v float64) {
	o.BalanceDue = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *EstimatedOrderModel) SetCurrency(v string) {
	o.Currency = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *EstimatedOrderModel) SetEmail(v string) {
	o.Email = &v
}

// GetOrderChannel returns the OrderChannel field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetOrderChannel() string {
	if o == nil || o.OrderChannel == nil {
		var ret string
		return ret
	}
	return *o.OrderChannel
}

// GetOrderChannelOk returns a tuple with the OrderChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetOrderChannelOk() (*string, bool) {
	if o == nil || o.OrderChannel == nil {
		return nil, false
	}
	return o.OrderChannel, true
}

// HasOrderChannel returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasOrderChannel() bool {
	if o != nil && o.OrderChannel != nil {
		return true
	}

	return false
}

// SetOrderChannel gets a reference to the given string and assigns it to the OrderChannel field.
func (o *EstimatedOrderModel) SetOrderChannel(v string) {
	o.OrderChannel = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetOrderId() string {
	if o == nil || o.OrderId == nil {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetOrderIdOk() (*string, bool) {
	if o == nil || o.OrderId == nil {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasOrderId() bool {
	if o != nil && o.OrderId != nil {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *EstimatedOrderModel) SetOrderId(v string) {
	o.OrderId = &v
}

// GetOrderItemList returns the OrderItemList field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetOrderItemList() []EstimatedOrderItemModel {
	if o == nil || o.OrderItemList == nil {
		var ret []EstimatedOrderItemModel
		return ret
	}
	return o.OrderItemList
}

// GetOrderItemListOk returns a tuple with the OrderItemList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetOrderItemListOk() ([]EstimatedOrderItemModel, bool) {
	if o == nil || o.OrderItemList == nil {
		return nil, false
	}
	return o.OrderItemList, true
}

// HasOrderItemList returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasOrderItemList() bool {
	if o != nil && o.OrderItemList != nil {
		return true
	}

	return false
}

// SetOrderItemList gets a reference to the given []EstimatedOrderItemModel and assigns it to the OrderItemList field.
func (o *EstimatedOrderModel) SetOrderItemList(v []EstimatedOrderItemModel) {
	o.OrderItemList = v
}

// GetPreferredDeliveryDate returns the PreferredDeliveryDate field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetPreferredDeliveryDate() string {
	if o == nil || o.PreferredDeliveryDate == nil {
		var ret string
		return ret
	}
	return *o.PreferredDeliveryDate
}

// GetPreferredDeliveryDateOk returns a tuple with the PreferredDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetPreferredDeliveryDateOk() (*string, bool) {
	if o == nil || o.PreferredDeliveryDate == nil {
		return nil, false
	}
	return o.PreferredDeliveryDate, true
}

// HasPreferredDeliveryDate returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasPreferredDeliveryDate() bool {
	if o != nil && o.PreferredDeliveryDate != nil {
		return true
	}

	return false
}

// SetPreferredDeliveryDate gets a reference to the given string and assigns it to the PreferredDeliveryDate field.
func (o *EstimatedOrderModel) SetPreferredDeliveryDate(v string) {
	o.PreferredDeliveryDate = &v
}

// GetPurchaseOrderNumber returns the PurchaseOrderNumber field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetPurchaseOrderNumber() string {
	if o == nil || o.PurchaseOrderNumber == nil {
		var ret string
		return ret
	}
	return *o.PurchaseOrderNumber
}

// GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetPurchaseOrderNumberOk() (*string, bool) {
	if o == nil || o.PurchaseOrderNumber == nil {
		return nil, false
	}
	return o.PurchaseOrderNumber, true
}

// HasPurchaseOrderNumber returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasPurchaseOrderNumber() bool {
	if o != nil && o.PurchaseOrderNumber != nil {
		return true
	}

	return false
}

// SetPurchaseOrderNumber gets a reference to the given string and assigns it to the PurchaseOrderNumber field.
func (o *EstimatedOrderModel) SetPurchaseOrderNumber(v string) {
	o.PurchaseOrderNumber = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetShippingAddress() ShippingAddressModel {
	if o == nil || o.ShippingAddress == nil {
		var ret ShippingAddressModel
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetShippingAddressOk() (*ShippingAddressModel, bool) {
	if o == nil || o.ShippingAddress == nil {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasShippingAddress() bool {
	if o != nil && o.ShippingAddress != nil {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given ShippingAddressModel and assigns it to the ShippingAddress field.
func (o *EstimatedOrderModel) SetShippingAddress(v ShippingAddressModel) {
	o.ShippingAddress = &v
}

// GetShippingAddressId returns the ShippingAddressId field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetShippingAddressId() string {
	if o == nil || o.ShippingAddressId == nil {
		var ret string
		return ret
	}
	return *o.ShippingAddressId
}

// GetShippingAddressIdOk returns a tuple with the ShippingAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetShippingAddressIdOk() (*string, bool) {
	if o == nil || o.ShippingAddressId == nil {
		return nil, false
	}
	return o.ShippingAddressId, true
}

// HasShippingAddressId returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasShippingAddressId() bool {
	if o != nil && o.ShippingAddressId != nil {
		return true
	}

	return false
}

// SetShippingAddressId gets a reference to the given string and assigns it to the ShippingAddressId field.
func (o *EstimatedOrderModel) SetShippingAddressId(v string) {
	o.ShippingAddressId = &v
}

// GetShippingCost returns the ShippingCost field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetShippingCost() float64 {
	if o == nil || o.ShippingCost == nil {
		var ret float64
		return ret
	}
	return *o.ShippingCost
}

// GetShippingCostOk returns a tuple with the ShippingCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetShippingCostOk() (*float64, bool) {
	if o == nil || o.ShippingCost == nil {
		return nil, false
	}
	return o.ShippingCost, true
}

// HasShippingCost returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasShippingCost() bool {
	if o != nil && o.ShippingCost != nil {
		return true
	}

	return false
}

// SetShippingCost gets a reference to the given float64 and assigns it to the ShippingCost field.
func (o *EstimatedOrderModel) SetShippingCost(v float64) {
	o.ShippingCost = &v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetTaxAmount() float64 {
	if o == nil || o.TaxAmount == nil {
		var ret float64
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetTaxAmountOk() (*float64, bool) {
	if o == nil || o.TaxAmount == nil {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasTaxAmount() bool {
	if o != nil && o.TaxAmount != nil {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given float64 and assigns it to the TaxAmount field.
func (o *EstimatedOrderModel) SetTaxAmount(v float64) {
	o.TaxAmount = &v
}

// GetTaxIncludedInShippingCost returns the TaxIncludedInShippingCost field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetTaxIncludedInShippingCost() float64 {
	if o == nil || o.TaxIncludedInShippingCost == nil {
		var ret float64
		return ret
	}
	return *o.TaxIncludedInShippingCost
}

// GetTaxIncludedInShippingCostOk returns a tuple with the TaxIncludedInShippingCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetTaxIncludedInShippingCostOk() (*float64, bool) {
	if o == nil || o.TaxIncludedInShippingCost == nil {
		return nil, false
	}
	return o.TaxIncludedInShippingCost, true
}

// HasTaxIncludedInShippingCost returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasTaxIncludedInShippingCost() bool {
	if o != nil && o.TaxIncludedInShippingCost != nil {
		return true
	}

	return false
}

// SetTaxIncludedInShippingCost gets a reference to the given float64 and assigns it to the TaxIncludedInShippingCost field.
func (o *EstimatedOrderModel) SetTaxIncludedInShippingCost(v float64) {
	o.TaxIncludedInShippingCost = &v
}

// GetTaxOnShippingCost returns the TaxOnShippingCost field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetTaxOnShippingCost() bool {
	if o == nil || o.TaxOnShippingCost == nil {
		var ret bool
		return ret
	}
	return *o.TaxOnShippingCost
}

// GetTaxOnShippingCostOk returns a tuple with the TaxOnShippingCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetTaxOnShippingCostOk() (*bool, bool) {
	if o == nil || o.TaxOnShippingCost == nil {
		return nil, false
	}
	return o.TaxOnShippingCost, true
}

// HasTaxOnShippingCost returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasTaxOnShippingCost() bool {
	if o != nil && o.TaxOnShippingCost != nil {
		return true
	}

	return false
}

// SetTaxOnShippingCost gets a reference to the given bool and assigns it to the TaxOnShippingCost field.
func (o *EstimatedOrderModel) SetTaxOnShippingCost(v bool) {
	o.TaxOnShippingCost = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetTotalAmount() float64 {
	if o == nil || o.TotalAmount == nil {
		var ret float64
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetTotalAmountOk() (*float64, bool) {
	if o == nil || o.TotalAmount == nil {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasTotalAmount() bool {
	if o != nil && o.TotalAmount != nil {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given float64 and assigns it to the TotalAmount field.
func (o *EstimatedOrderModel) SetTotalAmount(v float64) {
	o.TotalAmount = &v
}

// GetWithholdingTaxAmount returns the WithholdingTaxAmount field value if set, zero value otherwise.
func (o *EstimatedOrderModel) GetWithholdingTaxAmount() float64 {
	if o == nil || o.WithholdingTaxAmount == nil {
		var ret float64
		return ret
	}
	return *o.WithholdingTaxAmount
}

// GetWithholdingTaxAmountOk returns a tuple with the WithholdingTaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedOrderModel) GetWithholdingTaxAmountOk() (*float64, bool) {
	if o == nil || o.WithholdingTaxAmount == nil {
		return nil, false
	}
	return o.WithholdingTaxAmount, true
}

// HasWithholdingTaxAmount returns a boolean if a field has been set.
func (o *EstimatedOrderModel) HasWithholdingTaxAmount() bool {
	if o != nil && o.WithholdingTaxAmount != nil {
		return true
	}

	return false
}

// SetWithholdingTaxAmount gets a reference to the given float64 and assigns it to the WithholdingTaxAmount field.
func (o *EstimatedOrderModel) SetWithholdingTaxAmount(v float64) {
	o.WithholdingTaxAmount = &v
}

func (o EstimatedOrderModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppliedShippingOptions != nil {
		toSerialize["appliedShippingOptions"] = o.AppliedShippingOptions
	}
	if o.BalanceDue != nil {
		toSerialize["balanceDue"] = o.BalanceDue
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.OrderChannel != nil {
		toSerialize["orderChannel"] = o.OrderChannel
	}
	if o.OrderId != nil {
		toSerialize["orderId"] = o.OrderId
	}
	if o.OrderItemList != nil {
		toSerialize["orderItemList"] = o.OrderItemList
	}
	if o.PreferredDeliveryDate != nil {
		toSerialize["preferredDeliveryDate"] = o.PreferredDeliveryDate
	}
	if o.PurchaseOrderNumber != nil {
		toSerialize["purchaseOrderNumber"] = o.PurchaseOrderNumber
	}
	if o.ShippingAddress != nil {
		toSerialize["shippingAddress"] = o.ShippingAddress
	}
	if o.ShippingAddressId != nil {
		toSerialize["shippingAddressId"] = o.ShippingAddressId
	}
	if o.ShippingCost != nil {
		toSerialize["shippingCost"] = o.ShippingCost
	}
	if o.TaxAmount != nil {
		toSerialize["taxAmount"] = o.TaxAmount
	}
	if o.TaxIncludedInShippingCost != nil {
		toSerialize["taxIncludedInShippingCost"] = o.TaxIncludedInShippingCost
	}
	if o.TaxOnShippingCost != nil {
		toSerialize["taxOnShippingCost"] = o.TaxOnShippingCost
	}
	if o.TotalAmount != nil {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if o.WithholdingTaxAmount != nil {
		toSerialize["withholdingTaxAmount"] = o.WithholdingTaxAmount
	}
	return json.Marshal(toSerialize)
}

type NullableEstimatedOrderModel struct {
	value *EstimatedOrderModel
	isSet bool
}

func (v NullableEstimatedOrderModel) Get() *EstimatedOrderModel {
	return v.value
}

func (v *NullableEstimatedOrderModel) Set(val *EstimatedOrderModel) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimatedOrderModel) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimatedOrderModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimatedOrderModel(val *EstimatedOrderModel) *NullableEstimatedOrderModel {
	return &NullableEstimatedOrderModel{value: val, isSet: true}
}

func (v NullableEstimatedOrderModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimatedOrderModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


