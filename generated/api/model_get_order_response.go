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

// GetOrderResponse struct for GetOrderResponse
type GetOrderResponse struct {
	// 適用済み配送オプション
	AppliedShippingOptions []map[string]interface{} `json:"appliedShippingOptions,omitempty"`
	// 差引請求額
	BalanceDue *float64 `json:"balanceDue,omitempty"`
	ContainsTaxIncludedAmounts *bool `json:"containsTaxIncludedAmounts,omitempty"`
	// 通貨
	Currency *string `json:"currency,omitempty"`
	// メールアドレス
	Email *string `json:"email,omitempty"`
	// 発注チャンネル
	OrderChannel *string `json:"orderChannel,omitempty"`
	// 発注日時 (yyyyMMddHHmmss)
	OrderDateTime *string `json:"orderDateTime,omitempty"`
	// 発注 ID
	OrderId *string `json:"orderId,omitempty"`
	// 発注商品リスト
	OrderItemList []EstimatedOrderItemModel `json:"orderItemList,omitempty"`
	// 発注ステータス
	OrderStatus *string `json:"orderStatus,omitempty"`
	// 希望納品日
	PreferredDeliveryDate *string `json:"preferredDeliveryDate,omitempty"`
	// 注文番号
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
	ShippingAddress *ShippingAddressModel `json:"shippingAddress,omitempty"`
	// 商品発送先 ID
	ShippingAddressId *string `json:"shippingAddressId,omitempty"`
	// 配送料
	ShippingCost *float64 `json:"shippingCost,omitempty"`
	// 宅配便送り状番号 (代表)
	ShippingLabelNumber *string `json:"shippingLabelNumber,omitempty"`
	// 宅配便送り状番号
	ShippingLabelNumbers []string `json:"shippingLabelNumbers,omitempty"`
	// 税額
	TaxAmount *float64 `json:"taxAmount,omitempty"`
	// 送料に含まれる税額
	TaxIncludedInShippingCost *float64 `json:"taxIncludedInShippingCost,omitempty"`
	// 配送料の課税有無
	TaxOnShippingCost *bool `json:"taxOnShippingCost,omitempty"`
	// 合計額 (税込)
	TotalAmount *float64 `json:"totalAmount,omitempty"`
	// 源泉徴収税額
	WithholdingTaxAmount *float64 `json:"withholdingTaxAmount,omitempty"`
}

// NewGetOrderResponse instantiates a new GetOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderResponse() *GetOrderResponse {
	this := GetOrderResponse{}
	return &this
}

// NewGetOrderResponseWithDefaults instantiates a new GetOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderResponseWithDefaults() *GetOrderResponse {
	this := GetOrderResponse{}
	return &this
}

// GetAppliedShippingOptions returns the AppliedShippingOptions field value if set, zero value otherwise.
func (o *GetOrderResponse) GetAppliedShippingOptions() []map[string]interface{} {
	if o == nil || o.AppliedShippingOptions == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.AppliedShippingOptions
}

// GetAppliedShippingOptionsOk returns a tuple with the AppliedShippingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetAppliedShippingOptionsOk() ([]map[string]interface{}, bool) {
	if o == nil || o.AppliedShippingOptions == nil {
		return nil, false
	}
	return o.AppliedShippingOptions, true
}

// HasAppliedShippingOptions returns a boolean if a field has been set.
func (o *GetOrderResponse) HasAppliedShippingOptions() bool {
	if o != nil && o.AppliedShippingOptions != nil {
		return true
	}

	return false
}

// SetAppliedShippingOptions gets a reference to the given []map[string]interface{} and assigns it to the AppliedShippingOptions field.
func (o *GetOrderResponse) SetAppliedShippingOptions(v []map[string]interface{}) {
	o.AppliedShippingOptions = v
}

// GetBalanceDue returns the BalanceDue field value if set, zero value otherwise.
func (o *GetOrderResponse) GetBalanceDue() float64 {
	if o == nil || o.BalanceDue == nil {
		var ret float64
		return ret
	}
	return *o.BalanceDue
}

// GetBalanceDueOk returns a tuple with the BalanceDue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetBalanceDueOk() (*float64, bool) {
	if o == nil || o.BalanceDue == nil {
		return nil, false
	}
	return o.BalanceDue, true
}

// HasBalanceDue returns a boolean if a field has been set.
func (o *GetOrderResponse) HasBalanceDue() bool {
	if o != nil && o.BalanceDue != nil {
		return true
	}

	return false
}

// SetBalanceDue gets a reference to the given float64 and assigns it to the BalanceDue field.
func (o *GetOrderResponse) SetBalanceDue(v float64) {
	o.BalanceDue = &v
}

// GetContainsTaxIncludedAmounts returns the ContainsTaxIncludedAmounts field value if set, zero value otherwise.
func (o *GetOrderResponse) GetContainsTaxIncludedAmounts() bool {
	if o == nil || o.ContainsTaxIncludedAmounts == nil {
		var ret bool
		return ret
	}
	return *o.ContainsTaxIncludedAmounts
}

// GetContainsTaxIncludedAmountsOk returns a tuple with the ContainsTaxIncludedAmounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetContainsTaxIncludedAmountsOk() (*bool, bool) {
	if o == nil || o.ContainsTaxIncludedAmounts == nil {
		return nil, false
	}
	return o.ContainsTaxIncludedAmounts, true
}

// HasContainsTaxIncludedAmounts returns a boolean if a field has been set.
func (o *GetOrderResponse) HasContainsTaxIncludedAmounts() bool {
	if o != nil && o.ContainsTaxIncludedAmounts != nil {
		return true
	}

	return false
}

// SetContainsTaxIncludedAmounts gets a reference to the given bool and assigns it to the ContainsTaxIncludedAmounts field.
func (o *GetOrderResponse) SetContainsTaxIncludedAmounts(v bool) {
	o.ContainsTaxIncludedAmounts = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *GetOrderResponse) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *GetOrderResponse) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *GetOrderResponse) SetCurrency(v string) {
	o.Currency = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetOrderResponse) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetOrderResponse) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetOrderResponse) SetEmail(v string) {
	o.Email = &v
}

// GetOrderChannel returns the OrderChannel field value if set, zero value otherwise.
func (o *GetOrderResponse) GetOrderChannel() string {
	if o == nil || o.OrderChannel == nil {
		var ret string
		return ret
	}
	return *o.OrderChannel
}

// GetOrderChannelOk returns a tuple with the OrderChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetOrderChannelOk() (*string, bool) {
	if o == nil || o.OrderChannel == nil {
		return nil, false
	}
	return o.OrderChannel, true
}

// HasOrderChannel returns a boolean if a field has been set.
func (o *GetOrderResponse) HasOrderChannel() bool {
	if o != nil && o.OrderChannel != nil {
		return true
	}

	return false
}

// SetOrderChannel gets a reference to the given string and assigns it to the OrderChannel field.
func (o *GetOrderResponse) SetOrderChannel(v string) {
	o.OrderChannel = &v
}

// GetOrderDateTime returns the OrderDateTime field value if set, zero value otherwise.
func (o *GetOrderResponse) GetOrderDateTime() string {
	if o == nil || o.OrderDateTime == nil {
		var ret string
		return ret
	}
	return *o.OrderDateTime
}

// GetOrderDateTimeOk returns a tuple with the OrderDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetOrderDateTimeOk() (*string, bool) {
	if o == nil || o.OrderDateTime == nil {
		return nil, false
	}
	return o.OrderDateTime, true
}

// HasOrderDateTime returns a boolean if a field has been set.
func (o *GetOrderResponse) HasOrderDateTime() bool {
	if o != nil && o.OrderDateTime != nil {
		return true
	}

	return false
}

// SetOrderDateTime gets a reference to the given string and assigns it to the OrderDateTime field.
func (o *GetOrderResponse) SetOrderDateTime(v string) {
	o.OrderDateTime = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetOrderResponse) GetOrderId() string {
	if o == nil || o.OrderId == nil {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetOrderIdOk() (*string, bool) {
	if o == nil || o.OrderId == nil {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetOrderResponse) HasOrderId() bool {
	if o != nil && o.OrderId != nil {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *GetOrderResponse) SetOrderId(v string) {
	o.OrderId = &v
}

// GetOrderItemList returns the OrderItemList field value if set, zero value otherwise.
func (o *GetOrderResponse) GetOrderItemList() []EstimatedOrderItemModel {
	if o == nil || o.OrderItemList == nil {
		var ret []EstimatedOrderItemModel
		return ret
	}
	return o.OrderItemList
}

// GetOrderItemListOk returns a tuple with the OrderItemList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetOrderItemListOk() ([]EstimatedOrderItemModel, bool) {
	if o == nil || o.OrderItemList == nil {
		return nil, false
	}
	return o.OrderItemList, true
}

// HasOrderItemList returns a boolean if a field has been set.
func (o *GetOrderResponse) HasOrderItemList() bool {
	if o != nil && o.OrderItemList != nil {
		return true
	}

	return false
}

// SetOrderItemList gets a reference to the given []EstimatedOrderItemModel and assigns it to the OrderItemList field.
func (o *GetOrderResponse) SetOrderItemList(v []EstimatedOrderItemModel) {
	o.OrderItemList = v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *GetOrderResponse) GetOrderStatus() string {
	if o == nil || o.OrderStatus == nil {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetOrderStatusOk() (*string, bool) {
	if o == nil || o.OrderStatus == nil {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *GetOrderResponse) HasOrderStatus() bool {
	if o != nil && o.OrderStatus != nil {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *GetOrderResponse) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

// GetPreferredDeliveryDate returns the PreferredDeliveryDate field value if set, zero value otherwise.
func (o *GetOrderResponse) GetPreferredDeliveryDate() string {
	if o == nil || o.PreferredDeliveryDate == nil {
		var ret string
		return ret
	}
	return *o.PreferredDeliveryDate
}

// GetPreferredDeliveryDateOk returns a tuple with the PreferredDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetPreferredDeliveryDateOk() (*string, bool) {
	if o == nil || o.PreferredDeliveryDate == nil {
		return nil, false
	}
	return o.PreferredDeliveryDate, true
}

// HasPreferredDeliveryDate returns a boolean if a field has been set.
func (o *GetOrderResponse) HasPreferredDeliveryDate() bool {
	if o != nil && o.PreferredDeliveryDate != nil {
		return true
	}

	return false
}

// SetPreferredDeliveryDate gets a reference to the given string and assigns it to the PreferredDeliveryDate field.
func (o *GetOrderResponse) SetPreferredDeliveryDate(v string) {
	o.PreferredDeliveryDate = &v
}

// GetPurchaseOrderNumber returns the PurchaseOrderNumber field value if set, zero value otherwise.
func (o *GetOrderResponse) GetPurchaseOrderNumber() string {
	if o == nil || o.PurchaseOrderNumber == nil {
		var ret string
		return ret
	}
	return *o.PurchaseOrderNumber
}

// GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetPurchaseOrderNumberOk() (*string, bool) {
	if o == nil || o.PurchaseOrderNumber == nil {
		return nil, false
	}
	return o.PurchaseOrderNumber, true
}

// HasPurchaseOrderNumber returns a boolean if a field has been set.
func (o *GetOrderResponse) HasPurchaseOrderNumber() bool {
	if o != nil && o.PurchaseOrderNumber != nil {
		return true
	}

	return false
}

// SetPurchaseOrderNumber gets a reference to the given string and assigns it to the PurchaseOrderNumber field.
func (o *GetOrderResponse) SetPurchaseOrderNumber(v string) {
	o.PurchaseOrderNumber = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *GetOrderResponse) GetShippingAddress() ShippingAddressModel {
	if o == nil || o.ShippingAddress == nil {
		var ret ShippingAddressModel
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetShippingAddressOk() (*ShippingAddressModel, bool) {
	if o == nil || o.ShippingAddress == nil {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *GetOrderResponse) HasShippingAddress() bool {
	if o != nil && o.ShippingAddress != nil {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given ShippingAddressModel and assigns it to the ShippingAddress field.
func (o *GetOrderResponse) SetShippingAddress(v ShippingAddressModel) {
	o.ShippingAddress = &v
}

// GetShippingAddressId returns the ShippingAddressId field value if set, zero value otherwise.
func (o *GetOrderResponse) GetShippingAddressId() string {
	if o == nil || o.ShippingAddressId == nil {
		var ret string
		return ret
	}
	return *o.ShippingAddressId
}

// GetShippingAddressIdOk returns a tuple with the ShippingAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetShippingAddressIdOk() (*string, bool) {
	if o == nil || o.ShippingAddressId == nil {
		return nil, false
	}
	return o.ShippingAddressId, true
}

// HasShippingAddressId returns a boolean if a field has been set.
func (o *GetOrderResponse) HasShippingAddressId() bool {
	if o != nil && o.ShippingAddressId != nil {
		return true
	}

	return false
}

// SetShippingAddressId gets a reference to the given string and assigns it to the ShippingAddressId field.
func (o *GetOrderResponse) SetShippingAddressId(v string) {
	o.ShippingAddressId = &v
}

// GetShippingCost returns the ShippingCost field value if set, zero value otherwise.
func (o *GetOrderResponse) GetShippingCost() float64 {
	if o == nil || o.ShippingCost == nil {
		var ret float64
		return ret
	}
	return *o.ShippingCost
}

// GetShippingCostOk returns a tuple with the ShippingCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetShippingCostOk() (*float64, bool) {
	if o == nil || o.ShippingCost == nil {
		return nil, false
	}
	return o.ShippingCost, true
}

// HasShippingCost returns a boolean if a field has been set.
func (o *GetOrderResponse) HasShippingCost() bool {
	if o != nil && o.ShippingCost != nil {
		return true
	}

	return false
}

// SetShippingCost gets a reference to the given float64 and assigns it to the ShippingCost field.
func (o *GetOrderResponse) SetShippingCost(v float64) {
	o.ShippingCost = &v
}

// GetShippingLabelNumber returns the ShippingLabelNumber field value if set, zero value otherwise.
func (o *GetOrderResponse) GetShippingLabelNumber() string {
	if o == nil || o.ShippingLabelNumber == nil {
		var ret string
		return ret
	}
	return *o.ShippingLabelNumber
}

// GetShippingLabelNumberOk returns a tuple with the ShippingLabelNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetShippingLabelNumberOk() (*string, bool) {
	if o == nil || o.ShippingLabelNumber == nil {
		return nil, false
	}
	return o.ShippingLabelNumber, true
}

// HasShippingLabelNumber returns a boolean if a field has been set.
func (o *GetOrderResponse) HasShippingLabelNumber() bool {
	if o != nil && o.ShippingLabelNumber != nil {
		return true
	}

	return false
}

// SetShippingLabelNumber gets a reference to the given string and assigns it to the ShippingLabelNumber field.
func (o *GetOrderResponse) SetShippingLabelNumber(v string) {
	o.ShippingLabelNumber = &v
}

// GetShippingLabelNumbers returns the ShippingLabelNumbers field value if set, zero value otherwise.
func (o *GetOrderResponse) GetShippingLabelNumbers() []string {
	if o == nil || o.ShippingLabelNumbers == nil {
		var ret []string
		return ret
	}
	return o.ShippingLabelNumbers
}

// GetShippingLabelNumbersOk returns a tuple with the ShippingLabelNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetShippingLabelNumbersOk() ([]string, bool) {
	if o == nil || o.ShippingLabelNumbers == nil {
		return nil, false
	}
	return o.ShippingLabelNumbers, true
}

// HasShippingLabelNumbers returns a boolean if a field has been set.
func (o *GetOrderResponse) HasShippingLabelNumbers() bool {
	if o != nil && o.ShippingLabelNumbers != nil {
		return true
	}

	return false
}

// SetShippingLabelNumbers gets a reference to the given []string and assigns it to the ShippingLabelNumbers field.
func (o *GetOrderResponse) SetShippingLabelNumbers(v []string) {
	o.ShippingLabelNumbers = v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *GetOrderResponse) GetTaxAmount() float64 {
	if o == nil || o.TaxAmount == nil {
		var ret float64
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetTaxAmountOk() (*float64, bool) {
	if o == nil || o.TaxAmount == nil {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *GetOrderResponse) HasTaxAmount() bool {
	if o != nil && o.TaxAmount != nil {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given float64 and assigns it to the TaxAmount field.
func (o *GetOrderResponse) SetTaxAmount(v float64) {
	o.TaxAmount = &v
}

// GetTaxIncludedInShippingCost returns the TaxIncludedInShippingCost field value if set, zero value otherwise.
func (o *GetOrderResponse) GetTaxIncludedInShippingCost() float64 {
	if o == nil || o.TaxIncludedInShippingCost == nil {
		var ret float64
		return ret
	}
	return *o.TaxIncludedInShippingCost
}

// GetTaxIncludedInShippingCostOk returns a tuple with the TaxIncludedInShippingCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetTaxIncludedInShippingCostOk() (*float64, bool) {
	if o == nil || o.TaxIncludedInShippingCost == nil {
		return nil, false
	}
	return o.TaxIncludedInShippingCost, true
}

// HasTaxIncludedInShippingCost returns a boolean if a field has been set.
func (o *GetOrderResponse) HasTaxIncludedInShippingCost() bool {
	if o != nil && o.TaxIncludedInShippingCost != nil {
		return true
	}

	return false
}

// SetTaxIncludedInShippingCost gets a reference to the given float64 and assigns it to the TaxIncludedInShippingCost field.
func (o *GetOrderResponse) SetTaxIncludedInShippingCost(v float64) {
	o.TaxIncludedInShippingCost = &v
}

// GetTaxOnShippingCost returns the TaxOnShippingCost field value if set, zero value otherwise.
func (o *GetOrderResponse) GetTaxOnShippingCost() bool {
	if o == nil || o.TaxOnShippingCost == nil {
		var ret bool
		return ret
	}
	return *o.TaxOnShippingCost
}

// GetTaxOnShippingCostOk returns a tuple with the TaxOnShippingCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetTaxOnShippingCostOk() (*bool, bool) {
	if o == nil || o.TaxOnShippingCost == nil {
		return nil, false
	}
	return o.TaxOnShippingCost, true
}

// HasTaxOnShippingCost returns a boolean if a field has been set.
func (o *GetOrderResponse) HasTaxOnShippingCost() bool {
	if o != nil && o.TaxOnShippingCost != nil {
		return true
	}

	return false
}

// SetTaxOnShippingCost gets a reference to the given bool and assigns it to the TaxOnShippingCost field.
func (o *GetOrderResponse) SetTaxOnShippingCost(v bool) {
	o.TaxOnShippingCost = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *GetOrderResponse) GetTotalAmount() float64 {
	if o == nil || o.TotalAmount == nil {
		var ret float64
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetTotalAmountOk() (*float64, bool) {
	if o == nil || o.TotalAmount == nil {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *GetOrderResponse) HasTotalAmount() bool {
	if o != nil && o.TotalAmount != nil {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given float64 and assigns it to the TotalAmount field.
func (o *GetOrderResponse) SetTotalAmount(v float64) {
	o.TotalAmount = &v
}

// GetWithholdingTaxAmount returns the WithholdingTaxAmount field value if set, zero value otherwise.
func (o *GetOrderResponse) GetWithholdingTaxAmount() float64 {
	if o == nil || o.WithholdingTaxAmount == nil {
		var ret float64
		return ret
	}
	return *o.WithholdingTaxAmount
}

// GetWithholdingTaxAmountOk returns a tuple with the WithholdingTaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderResponse) GetWithholdingTaxAmountOk() (*float64, bool) {
	if o == nil || o.WithholdingTaxAmount == nil {
		return nil, false
	}
	return o.WithholdingTaxAmount, true
}

// HasWithholdingTaxAmount returns a boolean if a field has been set.
func (o *GetOrderResponse) HasWithholdingTaxAmount() bool {
	if o != nil && o.WithholdingTaxAmount != nil {
		return true
	}

	return false
}

// SetWithholdingTaxAmount gets a reference to the given float64 and assigns it to the WithholdingTaxAmount field.
func (o *GetOrderResponse) SetWithholdingTaxAmount(v float64) {
	o.WithholdingTaxAmount = &v
}

func (o GetOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppliedShippingOptions != nil {
		toSerialize["appliedShippingOptions"] = o.AppliedShippingOptions
	}
	if o.BalanceDue != nil {
		toSerialize["balanceDue"] = o.BalanceDue
	}
	if o.ContainsTaxIncludedAmounts != nil {
		toSerialize["containsTaxIncludedAmounts"] = o.ContainsTaxIncludedAmounts
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
	if o.OrderDateTime != nil {
		toSerialize["orderDateTime"] = o.OrderDateTime
	}
	if o.OrderId != nil {
		toSerialize["orderId"] = o.OrderId
	}
	if o.OrderItemList != nil {
		toSerialize["orderItemList"] = o.OrderItemList
	}
	if o.OrderStatus != nil {
		toSerialize["orderStatus"] = o.OrderStatus
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
	if o.ShippingLabelNumber != nil {
		toSerialize["shippingLabelNumber"] = o.ShippingLabelNumber
	}
	if o.ShippingLabelNumbers != nil {
		toSerialize["shippingLabelNumbers"] = o.ShippingLabelNumbers
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

type NullableGetOrderResponse struct {
	value *GetOrderResponse
	isSet bool
}

func (v NullableGetOrderResponse) Get() *GetOrderResponse {
	return v.value
}

func (v *NullableGetOrderResponse) Set(val *GetOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrderResponse(val *GetOrderResponse) *NullableGetOrderResponse {
	return &NullableGetOrderResponse{value: val, isSet: true}
}

func (v NullableGetOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


