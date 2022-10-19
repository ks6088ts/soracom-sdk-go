# CreateEstimatedVolumeDiscountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractTermMonth** | **int32** | 契約月数 | [default to 12]
**Quantity** | **int32** | 数量 | 
**StartDate** | Pointer to **string** | 適用開始日 | [optional] 
**VolumeDiscountPaymentType** | **string** | お支払い方法 | 
**VolumeDiscountType** | **string** | 料金計算方法 | 

## Methods

### NewCreateEstimatedVolumeDiscountRequest

`func NewCreateEstimatedVolumeDiscountRequest(contractTermMonth int32, quantity int32, volumeDiscountPaymentType string, volumeDiscountType string, ) *CreateEstimatedVolumeDiscountRequest`

NewCreateEstimatedVolumeDiscountRequest instantiates a new CreateEstimatedVolumeDiscountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEstimatedVolumeDiscountRequestWithDefaults

`func NewCreateEstimatedVolumeDiscountRequestWithDefaults() *CreateEstimatedVolumeDiscountRequest`

NewCreateEstimatedVolumeDiscountRequestWithDefaults instantiates a new CreateEstimatedVolumeDiscountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractTermMonth

`func (o *CreateEstimatedVolumeDiscountRequest) GetContractTermMonth() int32`

GetContractTermMonth returns the ContractTermMonth field if non-nil, zero value otherwise.

### GetContractTermMonthOk

`func (o *CreateEstimatedVolumeDiscountRequest) GetContractTermMonthOk() (*int32, bool)`

GetContractTermMonthOk returns a tuple with the ContractTermMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTermMonth

`func (o *CreateEstimatedVolumeDiscountRequest) SetContractTermMonth(v int32)`

SetContractTermMonth sets ContractTermMonth field to given value.


### GetQuantity

`func (o *CreateEstimatedVolumeDiscountRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateEstimatedVolumeDiscountRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateEstimatedVolumeDiscountRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetStartDate

`func (o *CreateEstimatedVolumeDiscountRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateEstimatedVolumeDiscountRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateEstimatedVolumeDiscountRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateEstimatedVolumeDiscountRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetVolumeDiscountPaymentType

`func (o *CreateEstimatedVolumeDiscountRequest) GetVolumeDiscountPaymentType() string`

GetVolumeDiscountPaymentType returns the VolumeDiscountPaymentType field if non-nil, zero value otherwise.

### GetVolumeDiscountPaymentTypeOk

`func (o *CreateEstimatedVolumeDiscountRequest) GetVolumeDiscountPaymentTypeOk() (*string, bool)`

GetVolumeDiscountPaymentTypeOk returns a tuple with the VolumeDiscountPaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDiscountPaymentType

`func (o *CreateEstimatedVolumeDiscountRequest) SetVolumeDiscountPaymentType(v string)`

SetVolumeDiscountPaymentType sets VolumeDiscountPaymentType field to given value.


### GetVolumeDiscountType

`func (o *CreateEstimatedVolumeDiscountRequest) GetVolumeDiscountType() string`

GetVolumeDiscountType returns the VolumeDiscountType field if non-nil, zero value otherwise.

### GetVolumeDiscountTypeOk

`func (o *CreateEstimatedVolumeDiscountRequest) GetVolumeDiscountTypeOk() (*string, bool)`

GetVolumeDiscountTypeOk returns a tuple with the VolumeDiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDiscountType

`func (o *CreateEstimatedVolumeDiscountRequest) SetVolumeDiscountType(v string)`

SetVolumeDiscountType sets VolumeDiscountType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


