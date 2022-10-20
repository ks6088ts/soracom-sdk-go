# ProductModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignCode** | Pointer to **string** | キャンペーンコード | [optional] 
**ConsignorId** | Pointer to **string** | 委託者 ID | [optional] 
**ConsignorName** | Pointer to **string** | 委託者名 | [optional] 
**Count** | Pointer to **float64** | 入数 | [optional] 
**Currency** | Pointer to **string** | 通貨 | [optional] 
**Description** | Pointer to **string** | 商品詳細 | [optional] 
**MaxQuantity** | Pointer to **int32** | 注文あたりの最大購入数量 | [optional] 
**Price** | Pointer to **float64** | 販売価格 | [optional] 
**PriceByQuantityList** | Pointer to [**[]PriceByQuantity**](PriceByQuantity.md) | 数量別価格リスト | [optional] 
**ProductCode** | Pointer to **string** | 商品コード | [optional] 
**ProductInfoURL** | Pointer to **string** | 商品説明ページ URL | [optional] 
**ProductName** | Pointer to **string** | 商品名 | [optional] 
**ProductType** | Pointer to **string** | 商品種別 | [optional] 
**Properties** | Pointer to **map[string]string** | 商品プロパティ | [optional] 
**RegularPrice** | Pointer to **float64** | 通常価格 | [optional] 
**TaxIncludedPrice** | Pointer to **float64** | 税込販売価格 | [optional] 
**TaxIncludedRegularPrice** | Pointer to **float64** | 税込通常価格 | [optional] 

## Methods

### NewProductModel

`func NewProductModel() *ProductModel`

NewProductModel instantiates a new ProductModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductModelWithDefaults

`func NewProductModelWithDefaults() *ProductModel`

NewProductModelWithDefaults instantiates a new ProductModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignCode

`func (o *ProductModel) GetCampaignCode() string`

GetCampaignCode returns the CampaignCode field if non-nil, zero value otherwise.

### GetCampaignCodeOk

`func (o *ProductModel) GetCampaignCodeOk() (*string, bool)`

GetCampaignCodeOk returns a tuple with the CampaignCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignCode

`func (o *ProductModel) SetCampaignCode(v string)`

SetCampaignCode sets CampaignCode field to given value.

### HasCampaignCode

`func (o *ProductModel) HasCampaignCode() bool`

HasCampaignCode returns a boolean if a field has been set.

### GetConsignorId

`func (o *ProductModel) GetConsignorId() string`

GetConsignorId returns the ConsignorId field if non-nil, zero value otherwise.

### GetConsignorIdOk

`func (o *ProductModel) GetConsignorIdOk() (*string, bool)`

GetConsignorIdOk returns a tuple with the ConsignorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsignorId

`func (o *ProductModel) SetConsignorId(v string)`

SetConsignorId sets ConsignorId field to given value.

### HasConsignorId

`func (o *ProductModel) HasConsignorId() bool`

HasConsignorId returns a boolean if a field has been set.

### GetConsignorName

`func (o *ProductModel) GetConsignorName() string`

GetConsignorName returns the ConsignorName field if non-nil, zero value otherwise.

### GetConsignorNameOk

`func (o *ProductModel) GetConsignorNameOk() (*string, bool)`

GetConsignorNameOk returns a tuple with the ConsignorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsignorName

`func (o *ProductModel) SetConsignorName(v string)`

SetConsignorName sets ConsignorName field to given value.

### HasConsignorName

`func (o *ProductModel) HasConsignorName() bool`

HasConsignorName returns a boolean if a field has been set.

### GetCount

`func (o *ProductModel) GetCount() float64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ProductModel) GetCountOk() (*float64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ProductModel) SetCount(v float64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ProductModel) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCurrency

`func (o *ProductModel) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ProductModel) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ProductModel) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ProductModel) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *ProductModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaxQuantity

`func (o *ProductModel) GetMaxQuantity() int32`

GetMaxQuantity returns the MaxQuantity field if non-nil, zero value otherwise.

### GetMaxQuantityOk

`func (o *ProductModel) GetMaxQuantityOk() (*int32, bool)`

GetMaxQuantityOk returns a tuple with the MaxQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQuantity

`func (o *ProductModel) SetMaxQuantity(v int32)`

SetMaxQuantity sets MaxQuantity field to given value.

### HasMaxQuantity

`func (o *ProductModel) HasMaxQuantity() bool`

HasMaxQuantity returns a boolean if a field has been set.

### GetPrice

`func (o *ProductModel) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductModel) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductModel) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductModel) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceByQuantityList

`func (o *ProductModel) GetPriceByQuantityList() []PriceByQuantity`

GetPriceByQuantityList returns the PriceByQuantityList field if non-nil, zero value otherwise.

### GetPriceByQuantityListOk

`func (o *ProductModel) GetPriceByQuantityListOk() (*[]PriceByQuantity, bool)`

GetPriceByQuantityListOk returns a tuple with the PriceByQuantityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceByQuantityList

`func (o *ProductModel) SetPriceByQuantityList(v []PriceByQuantity)`

SetPriceByQuantityList sets PriceByQuantityList field to given value.

### HasPriceByQuantityList

`func (o *ProductModel) HasPriceByQuantityList() bool`

HasPriceByQuantityList returns a boolean if a field has been set.

### GetProductCode

`func (o *ProductModel) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *ProductModel) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *ProductModel) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *ProductModel) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### GetProductInfoURL

`func (o *ProductModel) GetProductInfoURL() string`

GetProductInfoURL returns the ProductInfoURL field if non-nil, zero value otherwise.

### GetProductInfoURLOk

`func (o *ProductModel) GetProductInfoURLOk() (*string, bool)`

GetProductInfoURLOk returns a tuple with the ProductInfoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductInfoURL

`func (o *ProductModel) SetProductInfoURL(v string)`

SetProductInfoURL sets ProductInfoURL field to given value.

### HasProductInfoURL

`func (o *ProductModel) HasProductInfoURL() bool`

HasProductInfoURL returns a boolean if a field has been set.

### GetProductName

`func (o *ProductModel) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *ProductModel) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *ProductModel) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *ProductModel) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductType

`func (o *ProductModel) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ProductModel) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ProductModel) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *ProductModel) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetProperties

`func (o *ProductModel) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ProductModel) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ProductModel) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ProductModel) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRegularPrice

`func (o *ProductModel) GetRegularPrice() float64`

GetRegularPrice returns the RegularPrice field if non-nil, zero value otherwise.

### GetRegularPriceOk

`func (o *ProductModel) GetRegularPriceOk() (*float64, bool)`

GetRegularPriceOk returns a tuple with the RegularPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularPrice

`func (o *ProductModel) SetRegularPrice(v float64)`

SetRegularPrice sets RegularPrice field to given value.

### HasRegularPrice

`func (o *ProductModel) HasRegularPrice() bool`

HasRegularPrice returns a boolean if a field has been set.

### GetTaxIncludedPrice

`func (o *ProductModel) GetTaxIncludedPrice() float64`

GetTaxIncludedPrice returns the TaxIncludedPrice field if non-nil, zero value otherwise.

### GetTaxIncludedPriceOk

`func (o *ProductModel) GetTaxIncludedPriceOk() (*float64, bool)`

GetTaxIncludedPriceOk returns a tuple with the TaxIncludedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncludedPrice

`func (o *ProductModel) SetTaxIncludedPrice(v float64)`

SetTaxIncludedPrice sets TaxIncludedPrice field to given value.

### HasTaxIncludedPrice

`func (o *ProductModel) HasTaxIncludedPrice() bool`

HasTaxIncludedPrice returns a boolean if a field has been set.

### GetTaxIncludedRegularPrice

`func (o *ProductModel) GetTaxIncludedRegularPrice() float64`

GetTaxIncludedRegularPrice returns the TaxIncludedRegularPrice field if non-nil, zero value otherwise.

### GetTaxIncludedRegularPriceOk

`func (o *ProductModel) GetTaxIncludedRegularPriceOk() (*float64, bool)`

GetTaxIncludedRegularPriceOk returns a tuple with the TaxIncludedRegularPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncludedRegularPrice

`func (o *ProductModel) SetTaxIncludedRegularPrice(v float64)`

SetTaxIncludedRegularPrice sets TaxIncludedRegularPrice field to given value.

### HasTaxIncludedRegularPrice

`func (o *ProductModel) HasTaxIncludedRegularPrice() bool`

HasTaxIncludedRegularPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


