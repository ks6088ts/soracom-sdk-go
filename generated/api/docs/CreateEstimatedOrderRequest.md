# CreateEstimatedOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderItemList** | Pointer to [**[]OrderItemModel**](OrderItemModel.md) | 発注商品リスト | [optional] 
**ShippingAddressId** | Pointer to **string** | 商品発送先 ID | [optional] 

## Methods

### NewCreateEstimatedOrderRequest

`func NewCreateEstimatedOrderRequest() *CreateEstimatedOrderRequest`

NewCreateEstimatedOrderRequest instantiates a new CreateEstimatedOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEstimatedOrderRequestWithDefaults

`func NewCreateEstimatedOrderRequestWithDefaults() *CreateEstimatedOrderRequest`

NewCreateEstimatedOrderRequestWithDefaults instantiates a new CreateEstimatedOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderItemList

`func (o *CreateEstimatedOrderRequest) GetOrderItemList() []OrderItemModel`

GetOrderItemList returns the OrderItemList field if non-nil, zero value otherwise.

### GetOrderItemListOk

`func (o *CreateEstimatedOrderRequest) GetOrderItemListOk() (*[]OrderItemModel, bool)`

GetOrderItemListOk returns a tuple with the OrderItemList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemList

`func (o *CreateEstimatedOrderRequest) SetOrderItemList(v []OrderItemModel)`

SetOrderItemList sets OrderItemList field to given value.

### HasOrderItemList

`func (o *CreateEstimatedOrderRequest) HasOrderItemList() bool`

HasOrderItemList returns a boolean if a field has been set.

### GetShippingAddressId

`func (o *CreateEstimatedOrderRequest) GetShippingAddressId() string`

GetShippingAddressId returns the ShippingAddressId field if non-nil, zero value otherwise.

### GetShippingAddressIdOk

`func (o *CreateEstimatedOrderRequest) GetShippingAddressIdOk() (*string, bool)`

GetShippingAddressIdOk returns a tuple with the ShippingAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddressId

`func (o *CreateEstimatedOrderRequest) SetShippingAddressId(v string)`

SetShippingAddressId sets ShippingAddressId field to given value.

### HasShippingAddressId

`func (o *CreateEstimatedOrderRequest) HasShippingAddressId() bool`

HasShippingAddressId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


