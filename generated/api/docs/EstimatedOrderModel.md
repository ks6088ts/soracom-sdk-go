# EstimatedOrderModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedShippingOptions** | Pointer to **[]map[string]interface{}** | Applied shipping options | [optional] 
**BalanceDue** | Pointer to **float64** | Balance due | [optional] 
**Currency** | Pointer to **string** | Currency | [optional] 
**Email** | Pointer to **string** | メールアドレス | [optional] 
**OrderChannel** | Pointer to **string** | Order channel | [optional] 
**OrderId** | Pointer to **string** | 発注 ID | [optional] 
**OrderItemList** | Pointer to [**[]EstimatedOrderItemModel**](EstimatedOrderItemModel.md) | 発注商品リスト | [optional] 
**PreferredDeliveryDate** | Pointer to **string** | Preferred delivery date | [optional] 
**PurchaseOrderNumber** | Pointer to **string** | Purchase order number | [optional] 
**ShippingAddress** | Pointer to [**ShippingAddressModel**](ShippingAddressModel.md) |  | [optional] 
**ShippingAddressId** | Pointer to **string** | 商品発送先 ID | [optional] 
**ShippingCost** | Pointer to **float64** | 配送料 | [optional] 
**TaxAmount** | Pointer to **float64** | 消費税 | [optional] 
**TaxIncludedInShippingCost** | Pointer to **float64** | Tax included in shipping cost | [optional] 
**TaxOnShippingCost** | Pointer to **bool** | Whether shipping cost is taxable | [optional] 
**TotalAmount** | Pointer to **float64** | 合計金額 | [optional] 
**WithholdingTaxAmount** | Pointer to **float64** | Withholding tax amount | [optional] 

## Methods

### NewEstimatedOrderModel

`func NewEstimatedOrderModel() *EstimatedOrderModel`

NewEstimatedOrderModel instantiates a new EstimatedOrderModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimatedOrderModelWithDefaults

`func NewEstimatedOrderModelWithDefaults() *EstimatedOrderModel`

NewEstimatedOrderModelWithDefaults instantiates a new EstimatedOrderModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedShippingOptions

`func (o *EstimatedOrderModel) GetAppliedShippingOptions() []map[string]interface{}`

GetAppliedShippingOptions returns the AppliedShippingOptions field if non-nil, zero value otherwise.

### GetAppliedShippingOptionsOk

`func (o *EstimatedOrderModel) GetAppliedShippingOptionsOk() (*[]map[string]interface{}, bool)`

GetAppliedShippingOptionsOk returns a tuple with the AppliedShippingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedShippingOptions

`func (o *EstimatedOrderModel) SetAppliedShippingOptions(v []map[string]interface{})`

SetAppliedShippingOptions sets AppliedShippingOptions field to given value.

### HasAppliedShippingOptions

`func (o *EstimatedOrderModel) HasAppliedShippingOptions() bool`

HasAppliedShippingOptions returns a boolean if a field has been set.

### GetBalanceDue

`func (o *EstimatedOrderModel) GetBalanceDue() float64`

GetBalanceDue returns the BalanceDue field if non-nil, zero value otherwise.

### GetBalanceDueOk

`func (o *EstimatedOrderModel) GetBalanceDueOk() (*float64, bool)`

GetBalanceDueOk returns a tuple with the BalanceDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceDue

`func (o *EstimatedOrderModel) SetBalanceDue(v float64)`

SetBalanceDue sets BalanceDue field to given value.

### HasBalanceDue

`func (o *EstimatedOrderModel) HasBalanceDue() bool`

HasBalanceDue returns a boolean if a field has been set.

### GetCurrency

`func (o *EstimatedOrderModel) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *EstimatedOrderModel) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *EstimatedOrderModel) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *EstimatedOrderModel) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEmail

`func (o *EstimatedOrderModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EstimatedOrderModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EstimatedOrderModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EstimatedOrderModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetOrderChannel

`func (o *EstimatedOrderModel) GetOrderChannel() string`

GetOrderChannel returns the OrderChannel field if non-nil, zero value otherwise.

### GetOrderChannelOk

`func (o *EstimatedOrderModel) GetOrderChannelOk() (*string, bool)`

GetOrderChannelOk returns a tuple with the OrderChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderChannel

`func (o *EstimatedOrderModel) SetOrderChannel(v string)`

SetOrderChannel sets OrderChannel field to given value.

### HasOrderChannel

`func (o *EstimatedOrderModel) HasOrderChannel() bool`

HasOrderChannel returns a boolean if a field has been set.

### GetOrderId

`func (o *EstimatedOrderModel) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *EstimatedOrderModel) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *EstimatedOrderModel) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *EstimatedOrderModel) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderItemList

`func (o *EstimatedOrderModel) GetOrderItemList() []EstimatedOrderItemModel`

GetOrderItemList returns the OrderItemList field if non-nil, zero value otherwise.

### GetOrderItemListOk

`func (o *EstimatedOrderModel) GetOrderItemListOk() (*[]EstimatedOrderItemModel, bool)`

GetOrderItemListOk returns a tuple with the OrderItemList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemList

`func (o *EstimatedOrderModel) SetOrderItemList(v []EstimatedOrderItemModel)`

SetOrderItemList sets OrderItemList field to given value.

### HasOrderItemList

`func (o *EstimatedOrderModel) HasOrderItemList() bool`

HasOrderItemList returns a boolean if a field has been set.

### GetPreferredDeliveryDate

`func (o *EstimatedOrderModel) GetPreferredDeliveryDate() string`

GetPreferredDeliveryDate returns the PreferredDeliveryDate field if non-nil, zero value otherwise.

### GetPreferredDeliveryDateOk

`func (o *EstimatedOrderModel) GetPreferredDeliveryDateOk() (*string, bool)`

GetPreferredDeliveryDateOk returns a tuple with the PreferredDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDeliveryDate

`func (o *EstimatedOrderModel) SetPreferredDeliveryDate(v string)`

SetPreferredDeliveryDate sets PreferredDeliveryDate field to given value.

### HasPreferredDeliveryDate

`func (o *EstimatedOrderModel) HasPreferredDeliveryDate() bool`

HasPreferredDeliveryDate returns a boolean if a field has been set.

### GetPurchaseOrderNumber

`func (o *EstimatedOrderModel) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *EstimatedOrderModel) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *EstimatedOrderModel) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *EstimatedOrderModel) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetShippingAddress

`func (o *EstimatedOrderModel) GetShippingAddress() ShippingAddressModel`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *EstimatedOrderModel) GetShippingAddressOk() (*ShippingAddressModel, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *EstimatedOrderModel) SetShippingAddress(v ShippingAddressModel)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *EstimatedOrderModel) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetShippingAddressId

`func (o *EstimatedOrderModel) GetShippingAddressId() string`

GetShippingAddressId returns the ShippingAddressId field if non-nil, zero value otherwise.

### GetShippingAddressIdOk

`func (o *EstimatedOrderModel) GetShippingAddressIdOk() (*string, bool)`

GetShippingAddressIdOk returns a tuple with the ShippingAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddressId

`func (o *EstimatedOrderModel) SetShippingAddressId(v string)`

SetShippingAddressId sets ShippingAddressId field to given value.

### HasShippingAddressId

`func (o *EstimatedOrderModel) HasShippingAddressId() bool`

HasShippingAddressId returns a boolean if a field has been set.

### GetShippingCost

`func (o *EstimatedOrderModel) GetShippingCost() float64`

GetShippingCost returns the ShippingCost field if non-nil, zero value otherwise.

### GetShippingCostOk

`func (o *EstimatedOrderModel) GetShippingCostOk() (*float64, bool)`

GetShippingCostOk returns a tuple with the ShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCost

`func (o *EstimatedOrderModel) SetShippingCost(v float64)`

SetShippingCost sets ShippingCost field to given value.

### HasShippingCost

`func (o *EstimatedOrderModel) HasShippingCost() bool`

HasShippingCost returns a boolean if a field has been set.

### GetTaxAmount

`func (o *EstimatedOrderModel) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *EstimatedOrderModel) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *EstimatedOrderModel) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *EstimatedOrderModel) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTaxIncludedInShippingCost

`func (o *EstimatedOrderModel) GetTaxIncludedInShippingCost() float64`

GetTaxIncludedInShippingCost returns the TaxIncludedInShippingCost field if non-nil, zero value otherwise.

### GetTaxIncludedInShippingCostOk

`func (o *EstimatedOrderModel) GetTaxIncludedInShippingCostOk() (*float64, bool)`

GetTaxIncludedInShippingCostOk returns a tuple with the TaxIncludedInShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncludedInShippingCost

`func (o *EstimatedOrderModel) SetTaxIncludedInShippingCost(v float64)`

SetTaxIncludedInShippingCost sets TaxIncludedInShippingCost field to given value.

### HasTaxIncludedInShippingCost

`func (o *EstimatedOrderModel) HasTaxIncludedInShippingCost() bool`

HasTaxIncludedInShippingCost returns a boolean if a field has been set.

### GetTaxOnShippingCost

`func (o *EstimatedOrderModel) GetTaxOnShippingCost() bool`

GetTaxOnShippingCost returns the TaxOnShippingCost field if non-nil, zero value otherwise.

### GetTaxOnShippingCostOk

`func (o *EstimatedOrderModel) GetTaxOnShippingCostOk() (*bool, bool)`

GetTaxOnShippingCostOk returns a tuple with the TaxOnShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxOnShippingCost

`func (o *EstimatedOrderModel) SetTaxOnShippingCost(v bool)`

SetTaxOnShippingCost sets TaxOnShippingCost field to given value.

### HasTaxOnShippingCost

`func (o *EstimatedOrderModel) HasTaxOnShippingCost() bool`

HasTaxOnShippingCost returns a boolean if a field has been set.

### GetTotalAmount

`func (o *EstimatedOrderModel) GetTotalAmount() float64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *EstimatedOrderModel) GetTotalAmountOk() (*float64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *EstimatedOrderModel) SetTotalAmount(v float64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *EstimatedOrderModel) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetWithholdingTaxAmount

`func (o *EstimatedOrderModel) GetWithholdingTaxAmount() float64`

GetWithholdingTaxAmount returns the WithholdingTaxAmount field if non-nil, zero value otherwise.

### GetWithholdingTaxAmountOk

`func (o *EstimatedOrderModel) GetWithholdingTaxAmountOk() (*float64, bool)`

GetWithholdingTaxAmountOk returns a tuple with the WithholdingTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithholdingTaxAmount

`func (o *EstimatedOrderModel) SetWithholdingTaxAmount(v float64)`

SetWithholdingTaxAmount sets WithholdingTaxAmount field to given value.

### HasWithholdingTaxAmount

`func (o *EstimatedOrderModel) HasWithholdingTaxAmount() bool`

HasWithholdingTaxAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


