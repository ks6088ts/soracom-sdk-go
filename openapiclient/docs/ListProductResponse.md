# ListProductResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductList** | Pointer to [**[]ProductModel**](ProductModel.md) |  | [optional] 
**ShippingCostList** | Pointer to [**[]ShippingCostModel**](ShippingCostModel.md) |  | [optional] 

## Methods

### NewListProductResponse

`func NewListProductResponse() *ListProductResponse`

NewListProductResponse instantiates a new ListProductResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProductResponseWithDefaults

`func NewListProductResponseWithDefaults() *ListProductResponse`

NewListProductResponseWithDefaults instantiates a new ListProductResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductList

`func (o *ListProductResponse) GetProductList() []ProductModel`

GetProductList returns the ProductList field if non-nil, zero value otherwise.

### GetProductListOk

`func (o *ListProductResponse) GetProductListOk() (*[]ProductModel, bool)`

GetProductListOk returns a tuple with the ProductList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductList

`func (o *ListProductResponse) SetProductList(v []ProductModel)`

SetProductList sets ProductList field to given value.

### HasProductList

`func (o *ListProductResponse) HasProductList() bool`

HasProductList returns a boolean if a field has been set.

### GetShippingCostList

`func (o *ListProductResponse) GetShippingCostList() []ShippingCostModel`

GetShippingCostList returns the ShippingCostList field if non-nil, zero value otherwise.

### GetShippingCostListOk

`func (o *ListProductResponse) GetShippingCostListOk() (*[]ShippingCostModel, bool)`

GetShippingCostListOk returns a tuple with the ShippingCostList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCostList

`func (o *ListProductResponse) SetShippingCostList(v []ShippingCostModel)`

SetShippingCostList sets ShippingCostList field to given value.

### HasShippingCostList

`func (o *ListProductResponse) HasShippingCostList() bool`

HasShippingCostList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


