# GetVolumeDiscountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractDateTime** | Pointer to **string** | 契約日時 | [optional] 
**ContractId** | Pointer to **string** | 契約 ID | [optional] 
**ContractTermMonth** | Pointer to **int32** | 契約月数 | [optional] [default to 12]
**Currency** | Pointer to **string** | Currency | [optional] 
**EndDate** | Pointer to **string** | 適用終了日 | [optional] 
**InitialCost** | Pointer to **float64** | 初期費用 | [optional] 
**Quantity** | Pointer to **int32** | 数量 | [optional] 
**StartDate** | Pointer to **string** | 適用開始日 | [optional] 
**Status** | Pointer to **string** | ステータス | [optional] 
**UnitPrice** | Pointer to **float64** | 単価 | [optional] 
**VolumeDiscountPaymentType** | Pointer to **string** | お支払い方法 | [optional] 
**VolumeDiscountType** | Pointer to **string** | 料金計算方法 | [optional] 

## Methods

### NewGetVolumeDiscountResponse

`func NewGetVolumeDiscountResponse() *GetVolumeDiscountResponse`

NewGetVolumeDiscountResponse instantiates a new GetVolumeDiscountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVolumeDiscountResponseWithDefaults

`func NewGetVolumeDiscountResponseWithDefaults() *GetVolumeDiscountResponse`

NewGetVolumeDiscountResponseWithDefaults instantiates a new GetVolumeDiscountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractDateTime

`func (o *GetVolumeDiscountResponse) GetContractDateTime() string`

GetContractDateTime returns the ContractDateTime field if non-nil, zero value otherwise.

### GetContractDateTimeOk

`func (o *GetVolumeDiscountResponse) GetContractDateTimeOk() (*string, bool)`

GetContractDateTimeOk returns a tuple with the ContractDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractDateTime

`func (o *GetVolumeDiscountResponse) SetContractDateTime(v string)`

SetContractDateTime sets ContractDateTime field to given value.

### HasContractDateTime

`func (o *GetVolumeDiscountResponse) HasContractDateTime() bool`

HasContractDateTime returns a boolean if a field has been set.

### GetContractId

`func (o *GetVolumeDiscountResponse) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *GetVolumeDiscountResponse) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *GetVolumeDiscountResponse) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *GetVolumeDiscountResponse) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetContractTermMonth

`func (o *GetVolumeDiscountResponse) GetContractTermMonth() int32`

GetContractTermMonth returns the ContractTermMonth field if non-nil, zero value otherwise.

### GetContractTermMonthOk

`func (o *GetVolumeDiscountResponse) GetContractTermMonthOk() (*int32, bool)`

GetContractTermMonthOk returns a tuple with the ContractTermMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTermMonth

`func (o *GetVolumeDiscountResponse) SetContractTermMonth(v int32)`

SetContractTermMonth sets ContractTermMonth field to given value.

### HasContractTermMonth

`func (o *GetVolumeDiscountResponse) HasContractTermMonth() bool`

HasContractTermMonth returns a boolean if a field has been set.

### GetCurrency

`func (o *GetVolumeDiscountResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetVolumeDiscountResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetVolumeDiscountResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetVolumeDiscountResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEndDate

`func (o *GetVolumeDiscountResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetVolumeDiscountResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetVolumeDiscountResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetVolumeDiscountResponse) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetInitialCost

`func (o *GetVolumeDiscountResponse) GetInitialCost() float64`

GetInitialCost returns the InitialCost field if non-nil, zero value otherwise.

### GetInitialCostOk

`func (o *GetVolumeDiscountResponse) GetInitialCostOk() (*float64, bool)`

GetInitialCostOk returns a tuple with the InitialCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialCost

`func (o *GetVolumeDiscountResponse) SetInitialCost(v float64)`

SetInitialCost sets InitialCost field to given value.

### HasInitialCost

`func (o *GetVolumeDiscountResponse) HasInitialCost() bool`

HasInitialCost returns a boolean if a field has been set.

### GetQuantity

`func (o *GetVolumeDiscountResponse) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GetVolumeDiscountResponse) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GetVolumeDiscountResponse) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GetVolumeDiscountResponse) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStartDate

`func (o *GetVolumeDiscountResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetVolumeDiscountResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetVolumeDiscountResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetVolumeDiscountResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *GetVolumeDiscountResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetVolumeDiscountResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetVolumeDiscountResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetVolumeDiscountResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUnitPrice

`func (o *GetVolumeDiscountResponse) GetUnitPrice() float64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *GetVolumeDiscountResponse) GetUnitPriceOk() (*float64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *GetVolumeDiscountResponse) SetUnitPrice(v float64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *GetVolumeDiscountResponse) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetVolumeDiscountPaymentType

`func (o *GetVolumeDiscountResponse) GetVolumeDiscountPaymentType() string`

GetVolumeDiscountPaymentType returns the VolumeDiscountPaymentType field if non-nil, zero value otherwise.

### GetVolumeDiscountPaymentTypeOk

`func (o *GetVolumeDiscountResponse) GetVolumeDiscountPaymentTypeOk() (*string, bool)`

GetVolumeDiscountPaymentTypeOk returns a tuple with the VolumeDiscountPaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDiscountPaymentType

`func (o *GetVolumeDiscountResponse) SetVolumeDiscountPaymentType(v string)`

SetVolumeDiscountPaymentType sets VolumeDiscountPaymentType field to given value.

### HasVolumeDiscountPaymentType

`func (o *GetVolumeDiscountResponse) HasVolumeDiscountPaymentType() bool`

HasVolumeDiscountPaymentType returns a boolean if a field has been set.

### GetVolumeDiscountType

`func (o *GetVolumeDiscountResponse) GetVolumeDiscountType() string`

GetVolumeDiscountType returns the VolumeDiscountType field if non-nil, zero value otherwise.

### GetVolumeDiscountTypeOk

`func (o *GetVolumeDiscountResponse) GetVolumeDiscountTypeOk() (*string, bool)`

GetVolumeDiscountTypeOk returns a tuple with the VolumeDiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDiscountType

`func (o *GetVolumeDiscountResponse) SetVolumeDiscountType(v string)`

SetVolumeDiscountType sets VolumeDiscountType field to given value.

### HasVolumeDiscountType

`func (o *GetVolumeDiscountResponse) HasVolumeDiscountType() bool`

HasVolumeDiscountType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


