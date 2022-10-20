# ListOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderList** | Pointer to [**[]GetOrderResponse**](GetOrderResponse.md) | 発注リスト | [optional] 

## Methods

### NewListOrderResponse

`func NewListOrderResponse() *ListOrderResponse`

NewListOrderResponse instantiates a new ListOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrderResponseWithDefaults

`func NewListOrderResponseWithDefaults() *ListOrderResponse`

NewListOrderResponseWithDefaults instantiates a new ListOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderList

`func (o *ListOrderResponse) GetOrderList() []GetOrderResponse`

GetOrderList returns the OrderList field if non-nil, zero value otherwise.

### GetOrderListOk

`func (o *ListOrderResponse) GetOrderListOk() (*[]GetOrderResponse, bool)`

GetOrderListOk returns a tuple with the OrderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderList

`func (o *ListOrderResponse) SetOrderList(v []GetOrderResponse)`

SetOrderList sets OrderList field to given value.

### HasOrderList

`func (o *ListOrderResponse) HasOrderList() bool`

HasOrderList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


