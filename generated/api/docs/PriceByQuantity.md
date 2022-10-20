# PriceByQuantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinQuantity** | Pointer to **int32** | 最低数量 | [optional] 
**Price** | Pointer to **float64** | 販売価格 | [optional] 
**TaxIncludedPrice** | Pointer to **float64** | 税込販売価格 | [optional] 

## Methods

### NewPriceByQuantity

`func NewPriceByQuantity() *PriceByQuantity`

NewPriceByQuantity instantiates a new PriceByQuantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceByQuantityWithDefaults

`func NewPriceByQuantityWithDefaults() *PriceByQuantity`

NewPriceByQuantityWithDefaults instantiates a new PriceByQuantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinQuantity

`func (o *PriceByQuantity) GetMinQuantity() int32`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *PriceByQuantity) GetMinQuantityOk() (*int32, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *PriceByQuantity) SetMinQuantity(v int32)`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *PriceByQuantity) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.

### GetPrice

`func (o *PriceByQuantity) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PriceByQuantity) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PriceByQuantity) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PriceByQuantity) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTaxIncludedPrice

`func (o *PriceByQuantity) GetTaxIncludedPrice() float64`

GetTaxIncludedPrice returns the TaxIncludedPrice field if non-nil, zero value otherwise.

### GetTaxIncludedPriceOk

`func (o *PriceByQuantity) GetTaxIncludedPriceOk() (*float64, bool)`

GetTaxIncludedPriceOk returns a tuple with the TaxIncludedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncludedPrice

`func (o *PriceByQuantity) SetTaxIncludedPrice(v float64)`

SetTaxIncludedPrice sets TaxIncludedPrice field to given value.

### HasTaxIncludedPrice

`func (o *PriceByQuantity) HasTaxIncludedPrice() bool`

HasTaxIncludedPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


