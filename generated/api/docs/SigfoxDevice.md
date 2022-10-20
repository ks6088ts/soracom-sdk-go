# SigfoxDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**LastModifiedTime** | Pointer to **time.Time** |  | [optional] 
**LastSeen** | Pointer to [**LastSeen**](LastSeen.md) |  | [optional] 
**OperatorId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**TerminationEnabled** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSigfoxDevice

`func NewSigfoxDevice() *SigfoxDevice`

NewSigfoxDevice instantiates a new SigfoxDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigfoxDeviceWithDefaults

`func NewSigfoxDeviceWithDefaults() *SigfoxDevice`

NewSigfoxDeviceWithDefaults instantiates a new SigfoxDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *SigfoxDevice) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SigfoxDevice) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SigfoxDevice) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *SigfoxDevice) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetGroupId

`func (o *SigfoxDevice) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SigfoxDevice) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SigfoxDevice) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *SigfoxDevice) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *SigfoxDevice) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *SigfoxDevice) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *SigfoxDevice) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *SigfoxDevice) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetLastSeen

`func (o *SigfoxDevice) GetLastSeen() LastSeen`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *SigfoxDevice) GetLastSeenOk() (*LastSeen, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *SigfoxDevice) SetLastSeen(v LastSeen)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *SigfoxDevice) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetOperatorId

`func (o *SigfoxDevice) GetOperatorId() string`

GetOperatorId returns the OperatorId field if non-nil, zero value otherwise.

### GetOperatorIdOk

`func (o *SigfoxDevice) GetOperatorIdOk() (*string, bool)`

GetOperatorIdOk returns a tuple with the OperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorId

`func (o *SigfoxDevice) SetOperatorId(v string)`

SetOperatorId sets OperatorId field to given value.

### HasOperatorId

`func (o *SigfoxDevice) HasOperatorId() bool`

HasOperatorId returns a boolean if a field has been set.

### GetStatus

`func (o *SigfoxDevice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SigfoxDevice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SigfoxDevice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SigfoxDevice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *SigfoxDevice) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SigfoxDevice) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SigfoxDevice) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SigfoxDevice) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTerminationEnabled

`func (o *SigfoxDevice) GetTerminationEnabled() bool`

GetTerminationEnabled returns the TerminationEnabled field if non-nil, zero value otherwise.

### GetTerminationEnabledOk

`func (o *SigfoxDevice) GetTerminationEnabledOk() (*bool, bool)`

GetTerminationEnabledOk returns a tuple with the TerminationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationEnabled

`func (o *SigfoxDevice) SetTerminationEnabled(v bool)`

SetTerminationEnabled sets TerminationEnabled field to given value.

### HasTerminationEnabled

`func (o *SigfoxDevice) HasTerminationEnabled() bool`

HasTerminationEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


