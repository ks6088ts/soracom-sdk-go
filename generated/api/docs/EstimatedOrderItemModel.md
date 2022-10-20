# EstimatedOrderItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**ProductModel**](ProductModel.md) |  | [optional] 
**ProductAmount** | Pointer to **float64** | 商品金額 | [optional] 
**Quantity** | Pointer to **int32** | 購入数 | [optional] 
**TaxIncludedProductAmount** | Pointer to **float32** |  | [optional] 

## Methods

### NewEstimatedOrderItemModel

`func NewEstimatedOrderItemModel() *EstimatedOrderItemModel`

NewEstimatedOrderItemModel instantiates a new EstimatedOrderItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimatedOrderItemModelWithDefaults

`func NewEstimatedOrderItemModelWithDefaults() *EstimatedOrderItemModel`

NewEstimatedOrderItemModelWithDefaults instantiates a new EstimatedOrderItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *EstimatedOrderItemModel) GetProduct() ProductModel`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *EstimatedOrderItemModel) GetProductOk() (*ProductModel, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *EstimatedOrderItemModel) SetProduct(v ProductModel)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *EstimatedOrderItemModel) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetProductAmount

`func (o *EstimatedOrderItemModel) GetProductAmount() float64`

GetProductAmount returns the ProductAmount field if non-nil, zero value otherwise.

### GetProductAmountOk

`func (o *EstimatedOrderItemModel) GetProductAmountOk() (*float64, bool)`

GetProductAmountOk returns a tuple with the ProductAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAmount

`func (o *EstimatedOrderItemModel) SetProductAmount(v float64)`

SetProductAmount sets ProductAmount field to given value.

### HasProductAmount

`func (o *EstimatedOrderItemModel) HasProductAmount() bool`

HasProductAmount returns a boolean if a field has been set.

### GetQuantity

`func (o *EstimatedOrderItemModel) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *EstimatedOrderItemModel) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *EstimatedOrderItemModel) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *EstimatedOrderItemModel) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTaxIncludedProductAmount

`func (o *EstimatedOrderItemModel) GetTaxIncludedProductAmount() float32`

GetTaxIncludedProductAmount returns the TaxIncludedProductAmount field if non-nil, zero value otherwise.

### GetTaxIncludedProductAmountOk

`func (o *EstimatedOrderItemModel) GetTaxIncludedProductAmountOk() (*float32, bool)`

GetTaxIncludedProductAmountOk returns a tuple with the TaxIncludedProductAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncludedProductAmount

`func (o *EstimatedOrderItemModel) SetTaxIncludedProductAmount(v float32)`

SetTaxIncludedProductAmount sets TaxIncludedProductAmount field to given value.

### HasTaxIncludedProductAmount

`func (o *EstimatedOrderItemModel) HasTaxIncludedProductAmount() bool`

HasTaxIncludedProductAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


