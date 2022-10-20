# GetOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedShippingOptions** | Pointer to **[]map[string]interface{}** | 適用済み配送オプション | [optional] 
**BalanceDue** | Pointer to **float64** | 差引請求額 | [optional] 
**ContainsTaxIncludedAmounts** | Pointer to **bool** |  | [optional] 
**Currency** | Pointer to **string** | 通貨 | [optional] 
**Email** | Pointer to **string** | メールアドレス | [optional] 
**OrderChannel** | Pointer to **string** | 発注チャンネル | [optional] 
**OrderDateTime** | Pointer to **string** | 発注日時 (yyyyMMddHHmmss) | [optional] 
**OrderId** | Pointer to **string** | 発注 ID | [optional] 
**OrderItemList** | Pointer to [**[]EstimatedOrderItemModel**](EstimatedOrderItemModel.md) | 発注商品リスト | [optional] 
**OrderStatus** | Pointer to **string** | 発注ステータス | [optional] 
**PreferredDeliveryDate** | Pointer to **string** | 希望納品日 | [optional] 
**PurchaseOrderNumber** | Pointer to **string** | 注文番号 | [optional] 
**ShippingAddress** | Pointer to [**ShippingAddressModel**](ShippingAddressModel.md) |  | [optional] 
**ShippingAddressId** | Pointer to **string** | 商品発送先 ID | [optional] 
**ShippingCost** | Pointer to **float64** | 配送料 | [optional] 
**ShippingLabelNumber** | Pointer to **string** | 宅配便送り状番号 (代表) | [optional] 
**ShippingLabelNumbers** | Pointer to **[]string** | 宅配便送り状番号 | [optional] 
**TaxAmount** | Pointer to **float64** | 税額 | [optional] 
**TaxIncludedInShippingCost** | Pointer to **float64** | 送料に含まれる税額 | [optional] 
**TaxOnShippingCost** | Pointer to **bool** | 配送料の課税有無 | [optional] 
**TotalAmount** | Pointer to **float64** | 合計額 (税込) | [optional] 
**WithholdingTaxAmount** | Pointer to **float64** | 源泉徴収税額 | [optional] 

## Methods

### NewGetOrderResponse

`func NewGetOrderResponse() *GetOrderResponse`

NewGetOrderResponse instantiates a new GetOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderResponseWithDefaults

`func NewGetOrderResponseWithDefaults() *GetOrderResponse`

NewGetOrderResponseWithDefaults instantiates a new GetOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedShippingOptions

`func (o *GetOrderResponse) GetAppliedShippingOptions() []map[string]interface{}`

GetAppliedShippingOptions returns the AppliedShippingOptions field if non-nil, zero value otherwise.

### GetAppliedShippingOptionsOk

`func (o *GetOrderResponse) GetAppliedShippingOptionsOk() (*[]map[string]interface{}, bool)`

GetAppliedShippingOptionsOk returns a tuple with the AppliedShippingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedShippingOptions

`func (o *GetOrderResponse) SetAppliedShippingOptions(v []map[string]interface{})`

SetAppliedShippingOptions sets AppliedShippingOptions field to given value.

### HasAppliedShippingOptions

`func (o *GetOrderResponse) HasAppliedShippingOptions() bool`

HasAppliedShippingOptions returns a boolean if a field has been set.

### GetBalanceDue

`func (o *GetOrderResponse) GetBalanceDue() float64`

GetBalanceDue returns the BalanceDue field if non-nil, zero value otherwise.

### GetBalanceDueOk

`func (o *GetOrderResponse) GetBalanceDueOk() (*float64, bool)`

GetBalanceDueOk returns a tuple with the BalanceDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceDue

`func (o *GetOrderResponse) SetBalanceDue(v float64)`

SetBalanceDue sets BalanceDue field to given value.

### HasBalanceDue

`func (o *GetOrderResponse) HasBalanceDue() bool`

HasBalanceDue returns a boolean if a field has been set.

### GetContainsTaxIncludedAmounts

`func (o *GetOrderResponse) GetContainsTaxIncludedAmounts() bool`

GetContainsTaxIncludedAmounts returns the ContainsTaxIncludedAmounts field if non-nil, zero value otherwise.

### GetContainsTaxIncludedAmountsOk

`func (o *GetOrderResponse) GetContainsTaxIncludedAmountsOk() (*bool, bool)`

GetContainsTaxIncludedAmountsOk returns a tuple with the ContainsTaxIncludedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsTaxIncludedAmounts

`func (o *GetOrderResponse) SetContainsTaxIncludedAmounts(v bool)`

SetContainsTaxIncludedAmounts sets ContainsTaxIncludedAmounts field to given value.

### HasContainsTaxIncludedAmounts

`func (o *GetOrderResponse) HasContainsTaxIncludedAmounts() bool`

HasContainsTaxIncludedAmounts returns a boolean if a field has been set.

### GetCurrency

`func (o *GetOrderResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetOrderResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetOrderResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetOrderResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEmail

`func (o *GetOrderResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetOrderResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetOrderResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetOrderResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetOrderChannel

`func (o *GetOrderResponse) GetOrderChannel() string`

GetOrderChannel returns the OrderChannel field if non-nil, zero value otherwise.

### GetOrderChannelOk

`func (o *GetOrderResponse) GetOrderChannelOk() (*string, bool)`

GetOrderChannelOk returns a tuple with the OrderChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderChannel

`func (o *GetOrderResponse) SetOrderChannel(v string)`

SetOrderChannel sets OrderChannel field to given value.

### HasOrderChannel

`func (o *GetOrderResponse) HasOrderChannel() bool`

HasOrderChannel returns a boolean if a field has been set.

### GetOrderDateTime

`func (o *GetOrderResponse) GetOrderDateTime() string`

GetOrderDateTime returns the OrderDateTime field if non-nil, zero value otherwise.

### GetOrderDateTimeOk

`func (o *GetOrderResponse) GetOrderDateTimeOk() (*string, bool)`

GetOrderDateTimeOk returns a tuple with the OrderDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDateTime

`func (o *GetOrderResponse) SetOrderDateTime(v string)`

SetOrderDateTime sets OrderDateTime field to given value.

### HasOrderDateTime

`func (o *GetOrderResponse) HasOrderDateTime() bool`

HasOrderDateTime returns a boolean if a field has been set.

### GetOrderId

`func (o *GetOrderResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetOrderResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetOrderResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetOrderResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderItemList

`func (o *GetOrderResponse) GetOrderItemList() []EstimatedOrderItemModel`

GetOrderItemList returns the OrderItemList field if non-nil, zero value otherwise.

### GetOrderItemListOk

`func (o *GetOrderResponse) GetOrderItemListOk() (*[]EstimatedOrderItemModel, bool)`

GetOrderItemListOk returns a tuple with the OrderItemList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemList

`func (o *GetOrderResponse) SetOrderItemList(v []EstimatedOrderItemModel)`

SetOrderItemList sets OrderItemList field to given value.

### HasOrderItemList

`func (o *GetOrderResponse) HasOrderItemList() bool`

HasOrderItemList returns a boolean if a field has been set.

### GetOrderStatus

`func (o *GetOrderResponse) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *GetOrderResponse) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *GetOrderResponse) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *GetOrderResponse) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetPreferredDeliveryDate

`func (o *GetOrderResponse) GetPreferredDeliveryDate() string`

GetPreferredDeliveryDate returns the PreferredDeliveryDate field if non-nil, zero value otherwise.

### GetPreferredDeliveryDateOk

`func (o *GetOrderResponse) GetPreferredDeliveryDateOk() (*string, bool)`

GetPreferredDeliveryDateOk returns a tuple with the PreferredDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDeliveryDate

`func (o *GetOrderResponse) SetPreferredDeliveryDate(v string)`

SetPreferredDeliveryDate sets PreferredDeliveryDate field to given value.

### HasPreferredDeliveryDate

`func (o *GetOrderResponse) HasPreferredDeliveryDate() bool`

HasPreferredDeliveryDate returns a boolean if a field has been set.

### GetPurchaseOrderNumber

`func (o *GetOrderResponse) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *GetOrderResponse) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *GetOrderResponse) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *GetOrderResponse) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetShippingAddress

`func (o *GetOrderResponse) GetShippingAddress() ShippingAddressModel`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *GetOrderResponse) GetShippingAddressOk() (*ShippingAddressModel, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *GetOrderResponse) SetShippingAddress(v ShippingAddressModel)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *GetOrderResponse) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetShippingAddressId

`func (o *GetOrderResponse) GetShippingAddressId() string`

GetShippingAddressId returns the ShippingAddressId field if non-nil, zero value otherwise.

### GetShippingAddressIdOk

`func (o *GetOrderResponse) GetShippingAddressIdOk() (*string, bool)`

GetShippingAddressIdOk returns a tuple with the ShippingAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddressId

`func (o *GetOrderResponse) SetShippingAddressId(v string)`

SetShippingAddressId sets ShippingAddressId field to given value.

### HasShippingAddressId

`func (o *GetOrderResponse) HasShippingAddressId() bool`

HasShippingAddressId returns a boolean if a field has been set.

### GetShippingCost

`func (o *GetOrderResponse) GetShippingCost() float64`

GetShippingCost returns the ShippingCost field if non-nil, zero value otherwise.

### GetShippingCostOk

`func (o *GetOrderResponse) GetShippingCostOk() (*float64, bool)`

GetShippingCostOk returns a tuple with the ShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCost

`func (o *GetOrderResponse) SetShippingCost(v float64)`

SetShippingCost sets ShippingCost field to given value.

### HasShippingCost

`func (o *GetOrderResponse) HasShippingCost() bool`

HasShippingCost returns a boolean if a field has been set.

### GetShippingLabelNumber

`func (o *GetOrderResponse) GetShippingLabelNumber() string`

GetShippingLabelNumber returns the ShippingLabelNumber field if non-nil, zero value otherwise.

### GetShippingLabelNumberOk

`func (o *GetOrderResponse) GetShippingLabelNumberOk() (*string, bool)`

GetShippingLabelNumberOk returns a tuple with the ShippingLabelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelNumber

`func (o *GetOrderResponse) SetShippingLabelNumber(v string)`

SetShippingLabelNumber sets ShippingLabelNumber field to given value.

### HasShippingLabelNumber

`func (o *GetOrderResponse) HasShippingLabelNumber() bool`

HasShippingLabelNumber returns a boolean if a field has been set.

### GetShippingLabelNumbers

`func (o *GetOrderResponse) GetShippingLabelNumbers() []string`

GetShippingLabelNumbers returns the ShippingLabelNumbers field if non-nil, zero value otherwise.

### GetShippingLabelNumbersOk

`func (o *GetOrderResponse) GetShippingLabelNumbersOk() (*[]string, bool)`

GetShippingLabelNumbersOk returns a tuple with the ShippingLabelNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelNumbers

`func (o *GetOrderResponse) SetShippingLabelNumbers(v []string)`

SetShippingLabelNumbers sets ShippingLabelNumbers field to given value.

### HasShippingLabelNumbers

`func (o *GetOrderResponse) HasShippingLabelNumbers() bool`

HasShippingLabelNumbers returns a boolean if a field has been set.

### GetTaxAmount

`func (o *GetOrderResponse) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *GetOrderResponse) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *GetOrderResponse) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *GetOrderResponse) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTaxIncludedInShippingCost

`func (o *GetOrderResponse) GetTaxIncludedInShippingCost() float64`

GetTaxIncludedInShippingCost returns the TaxIncludedInShippingCost field if non-nil, zero value otherwise.

### GetTaxIncludedInShippingCostOk

`func (o *GetOrderResponse) GetTaxIncludedInShippingCostOk() (*float64, bool)`

GetTaxIncludedInShippingCostOk returns a tuple with the TaxIncludedInShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncludedInShippingCost

`func (o *GetOrderResponse) SetTaxIncludedInShippingCost(v float64)`

SetTaxIncludedInShippingCost sets TaxIncludedInShippingCost field to given value.

### HasTaxIncludedInShippingCost

`func (o *GetOrderResponse) HasTaxIncludedInShippingCost() bool`

HasTaxIncludedInShippingCost returns a boolean if a field has been set.

### GetTaxOnShippingCost

`func (o *GetOrderResponse) GetTaxOnShippingCost() bool`

GetTaxOnShippingCost returns the TaxOnShippingCost field if non-nil, zero value otherwise.

### GetTaxOnShippingCostOk

`func (o *GetOrderResponse) GetTaxOnShippingCostOk() (*bool, bool)`

GetTaxOnShippingCostOk returns a tuple with the TaxOnShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxOnShippingCost

`func (o *GetOrderResponse) SetTaxOnShippingCost(v bool)`

SetTaxOnShippingCost sets TaxOnShippingCost field to given value.

### HasTaxOnShippingCost

`func (o *GetOrderResponse) HasTaxOnShippingCost() bool`

HasTaxOnShippingCost returns a boolean if a field has been set.

### GetTotalAmount

`func (o *GetOrderResponse) GetTotalAmount() float64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *GetOrderResponse) GetTotalAmountOk() (*float64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *GetOrderResponse) SetTotalAmount(v float64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *GetOrderResponse) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetWithholdingTaxAmount

`func (o *GetOrderResponse) GetWithholdingTaxAmount() float64`

GetWithholdingTaxAmount returns the WithholdingTaxAmount field if non-nil, zero value otherwise.

### GetWithholdingTaxAmountOk

`func (o *GetOrderResponse) GetWithholdingTaxAmountOk() (*float64, bool)`

GetWithholdingTaxAmountOk returns a tuple with the WithholdingTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithholdingTaxAmount

`func (o *GetOrderResponse) SetWithholdingTaxAmount(v float64)`

SetWithholdingTaxAmount sets WithholdingTaxAmount field to given value.

### HasWithholdingTaxAmount

`func (o *GetOrderResponse) HasWithholdingTaxAmount() bool`

HasWithholdingTaxAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


