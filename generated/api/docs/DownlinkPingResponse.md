# DownlinkPingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rtt** | Pointer to **string** | ping の RTT 的な結果 | [optional] 
**Stat** | Pointer to **string** | ping の統計的な結果 | [optional] 
**Success** | Pointer to **bool** | ping の成功可否 | [optional] 

## Methods

### NewDownlinkPingResponse

`func NewDownlinkPingResponse() *DownlinkPingResponse`

NewDownlinkPingResponse instantiates a new DownlinkPingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownlinkPingResponseWithDefaults

`func NewDownlinkPingResponseWithDefaults() *DownlinkPingResponse`

NewDownlinkPingResponseWithDefaults instantiates a new DownlinkPingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRtt

`func (o *DownlinkPingResponse) GetRtt() string`

GetRtt returns the Rtt field if non-nil, zero value otherwise.

### GetRttOk

`func (o *DownlinkPingResponse) GetRttOk() (*string, bool)`

GetRttOk returns a tuple with the Rtt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtt

`func (o *DownlinkPingResponse) SetRtt(v string)`

SetRtt sets Rtt field to given value.

### HasRtt

`func (o *DownlinkPingResponse) HasRtt() bool`

HasRtt returns a boolean if a field has been set.

### GetStat

`func (o *DownlinkPingResponse) GetStat() string`

GetStat returns the Stat field if non-nil, zero value otherwise.

### GetStatOk

`func (o *DownlinkPingResponse) GetStatOk() (*string, bool)`

GetStatOk returns a tuple with the Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStat

`func (o *DownlinkPingResponse) SetStat(v string)`

SetStat sets Stat field to given value.

### HasStat

`func (o *DownlinkPingResponse) HasStat() bool`

HasStat returns a boolean if a field has been set.

### GetSuccess

`func (o *DownlinkPingResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DownlinkPingResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DownlinkPingResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *DownlinkPingResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


