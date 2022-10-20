# CouponResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float64** | クーポン額 | [optional] 
**Balance** | Pointer to **float64** | 残高 | [optional] 
**BillItemName** | Pointer to **string** | 対象課金アイテム | [optional] 
**CouponCode** | Pointer to **string** | クーポンコード | [optional] 
**Currency** | Pointer to **string** | 通貨 | [optional] 
**ExpiryYearMonth** | Pointer to **string** | 有効期限 | [optional] 
**OrderId** | Pointer to **string** | 発注 ID | [optional] 

## Methods

### NewCouponResponse

`func NewCouponResponse() *CouponResponse`

NewCouponResponse instantiates a new CouponResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponResponseWithDefaults

`func NewCouponResponseWithDefaults() *CouponResponse`

NewCouponResponseWithDefaults instantiates a new CouponResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CouponResponse) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CouponResponse) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CouponResponse) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CouponResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBalance

`func (o *CouponResponse) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CouponResponse) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CouponResponse) SetBalance(v float64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *CouponResponse) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetBillItemName

`func (o *CouponResponse) GetBillItemName() string`

GetBillItemName returns the BillItemName field if non-nil, zero value otherwise.

### GetBillItemNameOk

`func (o *CouponResponse) GetBillItemNameOk() (*string, bool)`

GetBillItemNameOk returns a tuple with the BillItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillItemName

`func (o *CouponResponse) SetBillItemName(v string)`

SetBillItemName sets BillItemName field to given value.

### HasBillItemName

`func (o *CouponResponse) HasBillItemName() bool`

HasBillItemName returns a boolean if a field has been set.

### GetCouponCode

`func (o *CouponResponse) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *CouponResponse) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *CouponResponse) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.

### HasCouponCode

`func (o *CouponResponse) HasCouponCode() bool`

HasCouponCode returns a boolean if a field has been set.

### GetCurrency

`func (o *CouponResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CouponResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CouponResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CouponResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExpiryYearMonth

`func (o *CouponResponse) GetExpiryYearMonth() string`

GetExpiryYearMonth returns the ExpiryYearMonth field if non-nil, zero value otherwise.

### GetExpiryYearMonthOk

`func (o *CouponResponse) GetExpiryYearMonthOk() (*string, bool)`

GetExpiryYearMonthOk returns a tuple with the ExpiryYearMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryYearMonth

`func (o *CouponResponse) SetExpiryYearMonth(v string)`

SetExpiryYearMonth sets ExpiryYearMonth field to given value.

### HasExpiryYearMonth

`func (o *CouponResponse) HasExpiryYearMonth() bool`

HasExpiryYearMonth returns a boolean if a field has been set.

### GetOrderId

`func (o *CouponResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CouponResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CouponResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CouponResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


