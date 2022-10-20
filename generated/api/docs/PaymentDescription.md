# PaymentDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**ItemList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPaymentDescription

`func NewPaymentDescription() *PaymentDescription`

NewPaymentDescription instantiates a new PaymentDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentDescriptionWithDefaults

`func NewPaymentDescriptionWithDefaults() *PaymentDescription`

NewPaymentDescriptionWithDefaults instantiates a new PaymentDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PaymentDescription) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentDescription) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentDescription) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentDescription) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetItemList

`func (o *PaymentDescription) GetItemList() []string`

GetItemList returns the ItemList field if non-nil, zero value otherwise.

### GetItemListOk

`func (o *PaymentDescription) GetItemListOk() (*[]string, bool)`

GetItemListOk returns a tuple with the ItemList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemList

`func (o *PaymentDescription) SetItemList(v []string)`

SetItemList sets ItemList field to given value.

### HasItemList

`func (o *PaymentDescription) HasItemList() bool`

HasItemList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


