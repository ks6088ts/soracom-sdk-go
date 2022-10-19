# DailyBill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | 金額 | [optional] 
**Currency** | Pointer to **string** | 通貨 | [optional] 
**Date** | Pointer to **string** | 年月日 | [optional] 

## Methods

### NewDailyBill

`func NewDailyBill() *DailyBill`

NewDailyBill instantiates a new DailyBill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyBillWithDefaults

`func NewDailyBillWithDefaults() *DailyBill`

NewDailyBillWithDefaults instantiates a new DailyBill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DailyBill) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DailyBill) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DailyBill) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DailyBill) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *DailyBill) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DailyBill) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DailyBill) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DailyBill) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDate

`func (o *DailyBill) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DailyBill) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DailyBill) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *DailyBill) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


