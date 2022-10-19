# DownlinkPingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfPingRequests** | Pointer to **int32** | ping の試行回数 | [optional] [default to 1]
**TimeoutSeconds** | Pointer to **int32** | 各 ping 実行のタイムアウト秒数 | [optional] [default to 1]

## Methods

### NewDownlinkPingRequest

`func NewDownlinkPingRequest() *DownlinkPingRequest`

NewDownlinkPingRequest instantiates a new DownlinkPingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownlinkPingRequestWithDefaults

`func NewDownlinkPingRequestWithDefaults() *DownlinkPingRequest`

NewDownlinkPingRequestWithDefaults instantiates a new DownlinkPingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfPingRequests

`func (o *DownlinkPingRequest) GetNumberOfPingRequests() int32`

GetNumberOfPingRequests returns the NumberOfPingRequests field if non-nil, zero value otherwise.

### GetNumberOfPingRequestsOk

`func (o *DownlinkPingRequest) GetNumberOfPingRequestsOk() (*int32, bool)`

GetNumberOfPingRequestsOk returns a tuple with the NumberOfPingRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPingRequests

`func (o *DownlinkPingRequest) SetNumberOfPingRequests(v int32)`

SetNumberOfPingRequests sets NumberOfPingRequests field to given value.

### HasNumberOfPingRequests

`func (o *DownlinkPingRequest) HasNumberOfPingRequests() bool`

HasNumberOfPingRequests returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *DownlinkPingRequest) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *DownlinkPingRequest) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *DownlinkPingRequest) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *DownlinkPingRequest) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


