# PaymentStatementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float64** | 金額 | [optional] 
**Currency** | Pointer to **string** | 通貨 | [optional] 
**PaymentDateTime** | Pointer to **int64** | 課金日時 | [optional] 
**PaymentInstrument** | Pointer to **string** | 支払情報 | [optional] 
**PaymentMethod** | Pointer to **string** | 支払方法 | [optional] 
**PaymentStatementId** | Pointer to **string** | 課金明細 ID | [optional] 
**PaymentStatementInfo** | Pointer to **string** | 支払内容 | [optional] 

## Methods

### NewPaymentStatementResponse

`func NewPaymentStatementResponse() *PaymentStatementResponse`

NewPaymentStatementResponse instantiates a new PaymentStatementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentStatementResponseWithDefaults

`func NewPaymentStatementResponseWithDefaults() *PaymentStatementResponse`

NewPaymentStatementResponseWithDefaults instantiates a new PaymentStatementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PaymentStatementResponse) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentStatementResponse) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentStatementResponse) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentStatementResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentStatementResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentStatementResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentStatementResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentStatementResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPaymentDateTime

`func (o *PaymentStatementResponse) GetPaymentDateTime() int64`

GetPaymentDateTime returns the PaymentDateTime field if non-nil, zero value otherwise.

### GetPaymentDateTimeOk

`func (o *PaymentStatementResponse) GetPaymentDateTimeOk() (*int64, bool)`

GetPaymentDateTimeOk returns a tuple with the PaymentDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDateTime

`func (o *PaymentStatementResponse) SetPaymentDateTime(v int64)`

SetPaymentDateTime sets PaymentDateTime field to given value.

### HasPaymentDateTime

`func (o *PaymentStatementResponse) HasPaymentDateTime() bool`

HasPaymentDateTime returns a boolean if a field has been set.

### GetPaymentInstrument

`func (o *PaymentStatementResponse) GetPaymentInstrument() string`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *PaymentStatementResponse) GetPaymentInstrumentOk() (*string, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *PaymentStatementResponse) SetPaymentInstrument(v string)`

SetPaymentInstrument sets PaymentInstrument field to given value.

### HasPaymentInstrument

`func (o *PaymentStatementResponse) HasPaymentInstrument() bool`

HasPaymentInstrument returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *PaymentStatementResponse) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentStatementResponse) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentStatementResponse) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *PaymentStatementResponse) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentStatementId

`func (o *PaymentStatementResponse) GetPaymentStatementId() string`

GetPaymentStatementId returns the PaymentStatementId field if non-nil, zero value otherwise.

### GetPaymentStatementIdOk

`func (o *PaymentStatementResponse) GetPaymentStatementIdOk() (*string, bool)`

GetPaymentStatementIdOk returns a tuple with the PaymentStatementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatementId

`func (o *PaymentStatementResponse) SetPaymentStatementId(v string)`

SetPaymentStatementId sets PaymentStatementId field to given value.

### HasPaymentStatementId

`func (o *PaymentStatementResponse) HasPaymentStatementId() bool`

HasPaymentStatementId returns a boolean if a field has been set.

### GetPaymentStatementInfo

`func (o *PaymentStatementResponse) GetPaymentStatementInfo() string`

GetPaymentStatementInfo returns the PaymentStatementInfo field if non-nil, zero value otherwise.

### GetPaymentStatementInfoOk

`func (o *PaymentStatementResponse) GetPaymentStatementInfoOk() (*string, bool)`

GetPaymentStatementInfoOk returns a tuple with the PaymentStatementInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatementInfo

`func (o *PaymentStatementResponse) SetPaymentStatementInfo(v string)`

SetPaymentStatementInfo sets PaymentStatementInfo field to given value.

### HasPaymentStatementInfo

`func (o *PaymentStatementResponse) HasPaymentStatementInfo() bool`

HasPaymentStatementInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


