# MonthlyBill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | 利用料金 | [optional] 
**Currency** | Pointer to **string** | 通貨 | [optional] 
**PaymentStatus** | Pointer to **string** | お支払い状況。詳しくは、[ご利用料金のお支払状況について](https://soracom.zendesk.com/hc/ja/articles/360029337031) を参照してください。  - &#x60;lessThanMinimumCharge&#x60;: 翌月以降請求 (最低課金額に満たないため) - &#x60;paying&#x60;: 支払い処理中 - &#x60;paid&#x60;: 支払い済み - &#x60;fail&#x60;: 処理失敗 - &#x60;refunding&#x60;: リファンド処理中 - &#x60;refunded&#x60;: リファンド済み - &#x60;refund_fail&#x60;: リファンド処理失敗 - &#x60;delegating_to_third_party&#x60;: 債権譲渡中 - &#x60;delegated_to_third_party&#x60;: 債権譲渡済み - &#x60;charging&#x60;: 請求実施中 - &#x60;canceling&#x60;: キャンセル中 - &#x60;cancelled&#x60;: キャンセル済み - &#x60;cancel_failed&#x60;: キャンセル処理失敗  | [optional] 
**PaymentTransactionId** | Pointer to **string** | 課金処理 ID。この ID を指定して [Payment:getPaymentTransaction API](#/Payment/getPaymentTransaction) を呼び出すと課金詳細が取得できます。 | [optional] 
**State** | Pointer to **string** | 利用料金の集計状況  - &#x60;temporary&#x60;: 集計中 - &#x60;closed&#x60;: 確定済み  | [optional] 
**YearMonth** | Pointer to **string** | 年月 | [optional] 

## Methods

### NewMonthlyBill

`func NewMonthlyBill() *MonthlyBill`

NewMonthlyBill instantiates a new MonthlyBill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonthlyBillWithDefaults

`func NewMonthlyBillWithDefaults() *MonthlyBill`

NewMonthlyBillWithDefaults instantiates a new MonthlyBill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *MonthlyBill) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *MonthlyBill) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *MonthlyBill) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *MonthlyBill) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *MonthlyBill) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MonthlyBill) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MonthlyBill) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MonthlyBill) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *MonthlyBill) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *MonthlyBill) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *MonthlyBill) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *MonthlyBill) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetPaymentTransactionId

`func (o *MonthlyBill) GetPaymentTransactionId() string`

GetPaymentTransactionId returns the PaymentTransactionId field if non-nil, zero value otherwise.

### GetPaymentTransactionIdOk

`func (o *MonthlyBill) GetPaymentTransactionIdOk() (*string, bool)`

GetPaymentTransactionIdOk returns a tuple with the PaymentTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTransactionId

`func (o *MonthlyBill) SetPaymentTransactionId(v string)`

SetPaymentTransactionId sets PaymentTransactionId field to given value.

### HasPaymentTransactionId

`func (o *MonthlyBill) HasPaymentTransactionId() bool`

HasPaymentTransactionId returns a boolean if a field has been set.

### GetState

`func (o *MonthlyBill) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MonthlyBill) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MonthlyBill) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MonthlyBill) HasState() bool`

HasState returns a boolean if a field has been set.

### GetYearMonth

`func (o *MonthlyBill) GetYearMonth() string`

GetYearMonth returns the YearMonth field if non-nil, zero value otherwise.

### GetYearMonthOk

`func (o *MonthlyBill) GetYearMonthOk() (*string, bool)`

GetYearMonthOk returns a tuple with the YearMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearMonth

`func (o *MonthlyBill) SetYearMonth(v string)`

SetYearMonth sets YearMonth field to given value.

### HasYearMonth

`func (o *MonthlyBill) HasYearMonth() bool`

HasYearMonth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


