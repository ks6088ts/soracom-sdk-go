# GetLatestBill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | 金額 | [optional] 
**Currency** | Pointer to **string** | 通貨 | [optional] 
**LastEvaluatedTime** | Pointer to **string** | 利用額計算を実施した時間 | [optional] 

## Methods

### NewGetLatestBill

`func NewGetLatestBill() *GetLatestBill`

NewGetLatestBill instantiates a new GetLatestBill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLatestBillWithDefaults

`func NewGetLatestBillWithDefaults() *GetLatestBill`

NewGetLatestBillWithDefaults instantiates a new GetLatestBill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetLatestBill) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetLatestBill) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetLatestBill) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetLatestBill) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *GetLatestBill) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetLatestBill) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetLatestBill) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetLatestBill) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetLastEvaluatedTime

`func (o *GetLatestBill) GetLastEvaluatedTime() string`

GetLastEvaluatedTime returns the LastEvaluatedTime field if non-nil, zero value otherwise.

### GetLastEvaluatedTimeOk

`func (o *GetLatestBill) GetLastEvaluatedTimeOk() (*string, bool)`

GetLastEvaluatedTimeOk returns a tuple with the LastEvaluatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEvaluatedTime

`func (o *GetLatestBill) SetLastEvaluatedTime(v string)`

SetLastEvaluatedTime sets LastEvaluatedTime field to given value.

### HasLastEvaluatedTime

`func (o *GetLatestBill) HasLastEvaluatedTime() bool`

HasLastEvaluatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


