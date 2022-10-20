# EstimatedVolumeDiscountModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | 通貨 | [optional] 
**OrderId** | Pointer to **string** | 発注 ID | [optional] 
**TaxAmount** | Pointer to **float64** | 消費税 | [optional] 
**TotalAmount** | Pointer to **float64** | 合計金額 | [optional] 
**VolumeDiscount** | Pointer to [**VolumeDiscountModel**](VolumeDiscountModel.md) |  | [optional] 

## Methods

### NewEstimatedVolumeDiscountModel

`func NewEstimatedVolumeDiscountModel() *EstimatedVolumeDiscountModel`

NewEstimatedVolumeDiscountModel instantiates a new EstimatedVolumeDiscountModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimatedVolumeDiscountModelWithDefaults

`func NewEstimatedVolumeDiscountModelWithDefaults() *EstimatedVolumeDiscountModel`

NewEstimatedVolumeDiscountModelWithDefaults instantiates a new EstimatedVolumeDiscountModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *EstimatedVolumeDiscountModel) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *EstimatedVolumeDiscountModel) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *EstimatedVolumeDiscountModel) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *EstimatedVolumeDiscountModel) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetOrderId

`func (o *EstimatedVolumeDiscountModel) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *EstimatedVolumeDiscountModel) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *EstimatedVolumeDiscountModel) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *EstimatedVolumeDiscountModel) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetTaxAmount

`func (o *EstimatedVolumeDiscountModel) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *EstimatedVolumeDiscountModel) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *EstimatedVolumeDiscountModel) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *EstimatedVolumeDiscountModel) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTotalAmount

`func (o *EstimatedVolumeDiscountModel) GetTotalAmount() float64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *EstimatedVolumeDiscountModel) GetTotalAmountOk() (*float64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *EstimatedVolumeDiscountModel) SetTotalAmount(v float64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *EstimatedVolumeDiscountModel) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetVolumeDiscount

`func (o *EstimatedVolumeDiscountModel) GetVolumeDiscount() VolumeDiscountModel`

GetVolumeDiscount returns the VolumeDiscount field if non-nil, zero value otherwise.

### GetVolumeDiscountOk

`func (o *EstimatedVolumeDiscountModel) GetVolumeDiscountOk() (*VolumeDiscountModel, bool)`

GetVolumeDiscountOk returns a tuple with the VolumeDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDiscount

`func (o *EstimatedVolumeDiscountModel) SetVolumeDiscount(v VolumeDiscountModel)`

SetVolumeDiscount sets VolumeDiscount field to given value.

### HasVolumeDiscount

`func (o *EstimatedVolumeDiscountModel) HasVolumeDiscount() bool`

HasVolumeDiscount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


