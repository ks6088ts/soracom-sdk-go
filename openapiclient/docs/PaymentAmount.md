# PaymentAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | 通貨 | [optional] 
**PaymentAmount** | Pointer to **float32** |  | [optional] 
**SubTotalAmount** | Pointer to **float32** |  | [optional] 
**TaxAmount** | Pointer to **float64** | 税額 | [optional] 
**TaxFreeSubTotalAmount** | Pointer to **float64** | 非課税小計 | [optional] 
**TaxRate** | Pointer to **map[string]interface{}** |  | [optional] 
**TaxableSubTotalAmount** | Pointer to **float64** | 課税小計 | [optional] 
**TotalAmount** | Pointer to **float64** | 合計額 (税込) | [optional] 
**WithholdingTaxAmount** | Pointer to **float64** | 源泉徴収税額 | [optional] 

## Methods

### NewPaymentAmount

`func NewPaymentAmount() *PaymentAmount`

NewPaymentAmount instantiates a new PaymentAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAmountWithDefaults

`func NewPaymentAmountWithDefaults() *PaymentAmount`

NewPaymentAmountWithDefaults instantiates a new PaymentAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *PaymentAmount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentAmount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentAmount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentAmount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *PaymentAmount) GetPaymentAmount() float32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *PaymentAmount) GetPaymentAmountOk() (*float32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *PaymentAmount) SetPaymentAmount(v float32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *PaymentAmount) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetSubTotalAmount

`func (o *PaymentAmount) GetSubTotalAmount() float32`

GetSubTotalAmount returns the SubTotalAmount field if non-nil, zero value otherwise.

### GetSubTotalAmountOk

`func (o *PaymentAmount) GetSubTotalAmountOk() (*float32, bool)`

GetSubTotalAmountOk returns a tuple with the SubTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotalAmount

`func (o *PaymentAmount) SetSubTotalAmount(v float32)`

SetSubTotalAmount sets SubTotalAmount field to given value.

### HasSubTotalAmount

`func (o *PaymentAmount) HasSubTotalAmount() bool`

HasSubTotalAmount returns a boolean if a field has been set.

### GetTaxAmount

`func (o *PaymentAmount) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *PaymentAmount) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *PaymentAmount) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *PaymentAmount) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTaxFreeSubTotalAmount

`func (o *PaymentAmount) GetTaxFreeSubTotalAmount() float64`

GetTaxFreeSubTotalAmount returns the TaxFreeSubTotalAmount field if non-nil, zero value otherwise.

### GetTaxFreeSubTotalAmountOk

`func (o *PaymentAmount) GetTaxFreeSubTotalAmountOk() (*float64, bool)`

GetTaxFreeSubTotalAmountOk returns a tuple with the TaxFreeSubTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxFreeSubTotalAmount

`func (o *PaymentAmount) SetTaxFreeSubTotalAmount(v float64)`

SetTaxFreeSubTotalAmount sets TaxFreeSubTotalAmount field to given value.

### HasTaxFreeSubTotalAmount

`func (o *PaymentAmount) HasTaxFreeSubTotalAmount() bool`

HasTaxFreeSubTotalAmount returns a boolean if a field has been set.

### GetTaxRate

`func (o *PaymentAmount) GetTaxRate() map[string]interface{}`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *PaymentAmount) GetTaxRateOk() (*map[string]interface{}, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *PaymentAmount) SetTaxRate(v map[string]interface{})`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *PaymentAmount) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetTaxableSubTotalAmount

`func (o *PaymentAmount) GetTaxableSubTotalAmount() float64`

GetTaxableSubTotalAmount returns the TaxableSubTotalAmount field if non-nil, zero value otherwise.

### GetTaxableSubTotalAmountOk

`func (o *PaymentAmount) GetTaxableSubTotalAmountOk() (*float64, bool)`

GetTaxableSubTotalAmountOk returns a tuple with the TaxableSubTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableSubTotalAmount

`func (o *PaymentAmount) SetTaxableSubTotalAmount(v float64)`

SetTaxableSubTotalAmount sets TaxableSubTotalAmount field to given value.

### HasTaxableSubTotalAmount

`func (o *PaymentAmount) HasTaxableSubTotalAmount() bool`

HasTaxableSubTotalAmount returns a boolean if a field has been set.

### GetTotalAmount

`func (o *PaymentAmount) GetTotalAmount() float64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *PaymentAmount) GetTotalAmountOk() (*float64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *PaymentAmount) SetTotalAmount(v float64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *PaymentAmount) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetWithholdingTaxAmount

`func (o *PaymentAmount) GetWithholdingTaxAmount() float64`

GetWithholdingTaxAmount returns the WithholdingTaxAmount field if non-nil, zero value otherwise.

### GetWithholdingTaxAmountOk

`func (o *PaymentAmount) GetWithholdingTaxAmountOk() (*float64, bool)`

GetWithholdingTaxAmountOk returns a tuple with the WithholdingTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithholdingTaxAmount

`func (o *PaymentAmount) SetWithholdingTaxAmount(v float64)`

SetWithholdingTaxAmount sets WithholdingTaxAmount field to given value.

### HasWithholdingTaxAmount

`func (o *PaymentAmount) HasWithholdingTaxAmount() bool`

HasWithholdingTaxAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


