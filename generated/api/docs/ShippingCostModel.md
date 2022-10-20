# ShippingCostModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedShippingOptions** | Pointer to **[]map[string]interface{}** | 適用済み配送オプション | [optional] 
**Currency** | Pointer to **string** | 通貨 | [optional] 
**ShippingArea** | Pointer to **string** | (日本カバレッジのみ) 発送先エリア | [optional] 
**ShippingAreaName** | Pointer to **string** | 発送先エリア名称 | [optional] 
**ShippingCost** | Pointer to **float64** | 送料 | [optional] 
**Size** | Pointer to **int32** | 発送サイズ | [optional] 

## Methods

### NewShippingCostModel

`func NewShippingCostModel() *ShippingCostModel`

NewShippingCostModel instantiates a new ShippingCostModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingCostModelWithDefaults

`func NewShippingCostModelWithDefaults() *ShippingCostModel`

NewShippingCostModelWithDefaults instantiates a new ShippingCostModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedShippingOptions

`func (o *ShippingCostModel) GetAppliedShippingOptions() []map[string]interface{}`

GetAppliedShippingOptions returns the AppliedShippingOptions field if non-nil, zero value otherwise.

### GetAppliedShippingOptionsOk

`func (o *ShippingCostModel) GetAppliedShippingOptionsOk() (*[]map[string]interface{}, bool)`

GetAppliedShippingOptionsOk returns a tuple with the AppliedShippingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedShippingOptions

`func (o *ShippingCostModel) SetAppliedShippingOptions(v []map[string]interface{})`

SetAppliedShippingOptions sets AppliedShippingOptions field to given value.

### HasAppliedShippingOptions

`func (o *ShippingCostModel) HasAppliedShippingOptions() bool`

HasAppliedShippingOptions returns a boolean if a field has been set.

### GetCurrency

`func (o *ShippingCostModel) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ShippingCostModel) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ShippingCostModel) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ShippingCostModel) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetShippingArea

`func (o *ShippingCostModel) GetShippingArea() string`

GetShippingArea returns the ShippingArea field if non-nil, zero value otherwise.

### GetShippingAreaOk

`func (o *ShippingCostModel) GetShippingAreaOk() (*string, bool)`

GetShippingAreaOk returns a tuple with the ShippingArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingArea

`func (o *ShippingCostModel) SetShippingArea(v string)`

SetShippingArea sets ShippingArea field to given value.

### HasShippingArea

`func (o *ShippingCostModel) HasShippingArea() bool`

HasShippingArea returns a boolean if a field has been set.

### GetShippingAreaName

`func (o *ShippingCostModel) GetShippingAreaName() string`

GetShippingAreaName returns the ShippingAreaName field if non-nil, zero value otherwise.

### GetShippingAreaNameOk

`func (o *ShippingCostModel) GetShippingAreaNameOk() (*string, bool)`

GetShippingAreaNameOk returns a tuple with the ShippingAreaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAreaName

`func (o *ShippingCostModel) SetShippingAreaName(v string)`

SetShippingAreaName sets ShippingAreaName field to given value.

### HasShippingAreaName

`func (o *ShippingCostModel) HasShippingAreaName() bool`

HasShippingAreaName returns a boolean if a field has been set.

### GetShippingCost

`func (o *ShippingCostModel) GetShippingCost() float64`

GetShippingCost returns the ShippingCost field if non-nil, zero value otherwise.

### GetShippingCostOk

`func (o *ShippingCostModel) GetShippingCostOk() (*float64, bool)`

GetShippingCostOk returns a tuple with the ShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCost

`func (o *ShippingCostModel) SetShippingCost(v float64)`

SetShippingCost sets ShippingCost field to given value.

### HasShippingCost

`func (o *ShippingCostModel) HasShippingCost() bool`

HasShippingCost returns a boolean if a field has been set.

### GetSize

`func (o *ShippingCostModel) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ShippingCostModel) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ShippingCostModel) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ShippingCostModel) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


