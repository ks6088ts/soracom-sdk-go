# LoraGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**GatewayId** | Pointer to **string** |  | [optional] 
**LastModifiedTime** | Pointer to **time.Time** |  | [optional] 
**NetworkSetId** | Pointer to **string** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] [default to false]
**OperatorId** | Pointer to **string** |  | [optional] 
**Owned** | Pointer to **bool** |  | [optional] [default to false]
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**TerminationEnabled** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewLoraGateway

`func NewLoraGateway() *LoraGateway`

NewLoraGateway instantiates a new LoraGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoraGatewayWithDefaults

`func NewLoraGatewayWithDefaults() *LoraGateway`

NewLoraGatewayWithDefaults instantiates a new LoraGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *LoraGateway) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LoraGateway) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LoraGateway) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *LoraGateway) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCreatedTime

`func (o *LoraGateway) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *LoraGateway) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *LoraGateway) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *LoraGateway) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetGatewayId

`func (o *LoraGateway) GetGatewayId() string`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *LoraGateway) GetGatewayIdOk() (*string, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *LoraGateway) SetGatewayId(v string)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *LoraGateway) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *LoraGateway) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *LoraGateway) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *LoraGateway) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *LoraGateway) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetNetworkSetId

`func (o *LoraGateway) GetNetworkSetId() string`

GetNetworkSetId returns the NetworkSetId field if non-nil, zero value otherwise.

### GetNetworkSetIdOk

`func (o *LoraGateway) GetNetworkSetIdOk() (*string, bool)`

GetNetworkSetIdOk returns a tuple with the NetworkSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSetId

`func (o *LoraGateway) SetNetworkSetId(v string)`

SetNetworkSetId sets NetworkSetId field to given value.

### HasNetworkSetId

`func (o *LoraGateway) HasNetworkSetId() bool`

HasNetworkSetId returns a boolean if a field has been set.

### GetOnline

`func (o *LoraGateway) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *LoraGateway) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *LoraGateway) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *LoraGateway) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetOperatorId

`func (o *LoraGateway) GetOperatorId() string`

GetOperatorId returns the OperatorId field if non-nil, zero value otherwise.

### GetOperatorIdOk

`func (o *LoraGateway) GetOperatorIdOk() (*string, bool)`

GetOperatorIdOk returns a tuple with the OperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorId

`func (o *LoraGateway) SetOperatorId(v string)`

SetOperatorId sets OperatorId field to given value.

### HasOperatorId

`func (o *LoraGateway) HasOperatorId() bool`

HasOperatorId returns a boolean if a field has been set.

### GetOwned

`func (o *LoraGateway) GetOwned() bool`

GetOwned returns the Owned field if non-nil, zero value otherwise.

### GetOwnedOk

`func (o *LoraGateway) GetOwnedOk() (*bool, bool)`

GetOwnedOk returns a tuple with the Owned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwned

`func (o *LoraGateway) SetOwned(v bool)`

SetOwned sets Owned field to given value.

### HasOwned

`func (o *LoraGateway) HasOwned() bool`

HasOwned returns a boolean if a field has been set.

### GetStatus

`func (o *LoraGateway) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LoraGateway) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LoraGateway) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LoraGateway) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *LoraGateway) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LoraGateway) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LoraGateway) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LoraGateway) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTerminationEnabled

`func (o *LoraGateway) GetTerminationEnabled() bool`

GetTerminationEnabled returns the TerminationEnabled field if non-nil, zero value otherwise.

### GetTerminationEnabledOk

`func (o *LoraGateway) GetTerminationEnabledOk() (*bool, bool)`

GetTerminationEnabledOk returns a tuple with the TerminationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationEnabled

`func (o *LoraGateway) SetTerminationEnabled(v bool)`

SetTerminationEnabled sets TerminationEnabled field to given value.

### HasTerminationEnabled

`func (o *LoraGateway) HasTerminationEnabled() bool`

HasTerminationEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


