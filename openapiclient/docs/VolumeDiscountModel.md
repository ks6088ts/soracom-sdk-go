# VolumeDiscountModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractTermMonth** | Pointer to **int32** | 契約月数 | [optional] [default to 12]
**Currency** | Pointer to **string** | 通貨 | [optional] 
**EndDate** | Pointer to **string** | 適用終了日 | [optional] 
**InitialCost** | Pointer to **float64** | 初期費用 | [optional] 
**Quantity** | Pointer to **int32** | 数量 | [optional] 
**StartDate** | Pointer to **string** | 適用開始日 | [optional] 
**TaxIncludedInitialCost** | Pointer to **float64** | 税込初期費用 | [optional] 
**TaxIncludedUnitPrice** | Pointer to **float64** | 税込単価 | [optional] 
**UnitPrice** | Pointer to **float64** | 単価 | [optional] 
**VolumeDiscountPaymentType** | Pointer to **string** | お支払い方法 | [optional] 
**VolumeDiscountType** | Pointer to **string** | 料金計算方法 | [optional] 

## Methods

### NewVolumeDiscountModel

`func NewVolumeDiscountModel() *VolumeDiscountModel`

NewVolumeDiscountModel instantiates a new VolumeDiscountModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeDiscountModelWithDefaults

`func NewVolumeDiscountModelWithDefaults() *VolumeDiscountModel`

NewVolumeDiscountModelWithDefaults instantiates a new VolumeDiscountModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractTermMonth

`func (o *VolumeDiscountModel) GetContractTermMonth() int32`

GetContractTermMonth returns the ContractTermMonth field if non-nil, zero value otherwise.

### GetContractTermMonthOk

`func (o *VolumeDiscountModel) GetContractTermMonthOk() (*int32, bool)`

GetContractTermMonthOk returns a tuple with the ContractTermMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTermMonth

`func (o *VolumeDiscountModel) SetContractTermMonth(v int32)`

SetContractTermMonth sets ContractTermMonth field to given value.

### HasContractTermMonth

`func (o *VolumeDiscountModel) HasContractTermMonth() bool`

HasContractTermMonth returns a boolean if a field has been set.

### GetCurrency

`func (o *VolumeDiscountModel) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *VolumeDiscountModel) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *VolumeDiscountModel) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *VolumeDiscountModel) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEndDate

`func (o *VolumeDiscountModel) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *VolumeDiscountModel) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *VolumeDiscountModel) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *VolumeDiscountModel) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetInitialCost

`func (o *VolumeDiscountModel) GetInitialCost() float64`

GetInitialCost returns the InitialCost field if non-nil, zero value otherwise.

### GetInitialCostOk

`func (o *VolumeDiscountModel) GetInitialCostOk() (*float64, bool)`

GetInitialCostOk returns a tuple with the InitialCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialCost

`func (o *VolumeDiscountModel) SetInitialCost(v float64)`

SetInitialCost sets InitialCost field to given value.

### HasInitialCost

`func (o *VolumeDiscountModel) HasInitialCost() bool`

HasInitialCost returns a boolean if a field has been set.

### GetQuantity

`func (o *VolumeDiscountModel) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *VolumeDiscountModel) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *VolumeDiscountModel) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *VolumeDiscountModel) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStartDate

`func (o *VolumeDiscountModel) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *VolumeDiscountModel) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *VolumeDiscountModel) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *VolumeDiscountModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTaxIncludedInitialCost

`func (o *VolumeDiscountModel) GetTaxIncludedInitialCost() float64`

GetTaxIncludedInitialCost returns the TaxIncludedInitialCost field if non-nil, zero value otherwise.

### GetTaxIncludedInitialCostOk

`func (o *VolumeDiscountModel) GetTaxIncludedInitialCostOk() (*float64, bool)`

GetTaxIncludedInitialCostOk returns a tuple with the TaxIncludedInitialCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncludedInitialCost

`func (o *VolumeDiscountModel) SetTaxIncludedInitialCost(v float64)`

SetTaxIncludedInitialCost sets TaxIncludedInitialCost field to given value.

### HasTaxIncludedInitialCost

`func (o *VolumeDiscountModel) HasTaxIncludedInitialCost() bool`

HasTaxIncludedInitialCost returns a boolean if a field has been set.

### GetTaxIncludedUnitPrice

`func (o *VolumeDiscountModel) GetTaxIncludedUnitPrice() float64`

GetTaxIncludedUnitPrice returns the TaxIncludedUnitPrice field if non-nil, zero value otherwise.

### GetTaxIncludedUnitPriceOk

`func (o *VolumeDiscountModel) GetTaxIncludedUnitPriceOk() (*float64, bool)`

GetTaxIncludedUnitPriceOk returns a tuple with the TaxIncludedUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncludedUnitPrice

`func (o *VolumeDiscountModel) SetTaxIncludedUnitPrice(v float64)`

SetTaxIncludedUnitPrice sets TaxIncludedUnitPrice field to given value.

### HasTaxIncludedUnitPrice

`func (o *VolumeDiscountModel) HasTaxIncludedUnitPrice() bool`

HasTaxIncludedUnitPrice returns a boolean if a field has been set.

### GetUnitPrice

`func (o *VolumeDiscountModel) GetUnitPrice() float64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *VolumeDiscountModel) GetUnitPriceOk() (*float64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *VolumeDiscountModel) SetUnitPrice(v float64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *VolumeDiscountModel) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetVolumeDiscountPaymentType

`func (o *VolumeDiscountModel) GetVolumeDiscountPaymentType() string`

GetVolumeDiscountPaymentType returns the VolumeDiscountPaymentType field if non-nil, zero value otherwise.

### GetVolumeDiscountPaymentTypeOk

`func (o *VolumeDiscountModel) GetVolumeDiscountPaymentTypeOk() (*string, bool)`

GetVolumeDiscountPaymentTypeOk returns a tuple with the VolumeDiscountPaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDiscountPaymentType

`func (o *VolumeDiscountModel) SetVolumeDiscountPaymentType(v string)`

SetVolumeDiscountPaymentType sets VolumeDiscountPaymentType field to given value.

### HasVolumeDiscountPaymentType

`func (o *VolumeDiscountModel) HasVolumeDiscountPaymentType() bool`

HasVolumeDiscountPaymentType returns a boolean if a field has been set.

### GetVolumeDiscountType

`func (o *VolumeDiscountModel) GetVolumeDiscountType() string`

GetVolumeDiscountType returns the VolumeDiscountType field if non-nil, zero value otherwise.

### GetVolumeDiscountTypeOk

`func (o *VolumeDiscountModel) GetVolumeDiscountTypeOk() (*string, bool)`

GetVolumeDiscountTypeOk returns a tuple with the VolumeDiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDiscountType

`func (o *VolumeDiscountModel) SetVolumeDiscountType(v string)`

SetVolumeDiscountType sets VolumeDiscountType field to given value.

### HasVolumeDiscountType

`func (o *VolumeDiscountModel) HasVolumeDiscountType() bool`

HasVolumeDiscountType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


