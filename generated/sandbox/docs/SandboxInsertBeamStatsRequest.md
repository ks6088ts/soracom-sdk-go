# SandboxInsertBeamStatsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeamStatsMap** | Pointer to [**SandboxInsertBeamStatsRequestBeamStatsMap**](SandboxInsertBeamStatsRequestBeamStatsMap.md) |  | [optional] 
**Unixtime** | Pointer to **int64** | UNIX 時刻（ミリ秒単位） | [optional] 

## Methods

### NewSandboxInsertBeamStatsRequest

`func NewSandboxInsertBeamStatsRequest() *SandboxInsertBeamStatsRequest`

NewSandboxInsertBeamStatsRequest instantiates a new SandboxInsertBeamStatsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxInsertBeamStatsRequestWithDefaults

`func NewSandboxInsertBeamStatsRequestWithDefaults() *SandboxInsertBeamStatsRequest`

NewSandboxInsertBeamStatsRequestWithDefaults instantiates a new SandboxInsertBeamStatsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeamStatsMap

`func (o *SandboxInsertBeamStatsRequest) GetBeamStatsMap() SandboxInsertBeamStatsRequestBeamStatsMap`

GetBeamStatsMap returns the BeamStatsMap field if non-nil, zero value otherwise.

### GetBeamStatsMapOk

`func (o *SandboxInsertBeamStatsRequest) GetBeamStatsMapOk() (*SandboxInsertBeamStatsRequestBeamStatsMap, bool)`

GetBeamStatsMapOk returns a tuple with the BeamStatsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeamStatsMap

`func (o *SandboxInsertBeamStatsRequest) SetBeamStatsMap(v SandboxInsertBeamStatsRequestBeamStatsMap)`

SetBeamStatsMap sets BeamStatsMap field to given value.

### HasBeamStatsMap

`func (o *SandboxInsertBeamStatsRequest) HasBeamStatsMap() bool`

HasBeamStatsMap returns a boolean if a field has been set.

### GetUnixtime

`func (o *SandboxInsertBeamStatsRequest) GetUnixtime() int64`

GetUnixtime returns the Unixtime field if non-nil, zero value otherwise.

### GetUnixtimeOk

`func (o *SandboxInsertBeamStatsRequest) GetUnixtimeOk() (*int64, bool)`

GetUnixtimeOk returns a tuple with the Unixtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixtime

`func (o *SandboxInsertBeamStatsRequest) SetUnixtime(v int64)`

SetUnixtime sets Unixtime field to given value.

### HasUnixtime

`func (o *SandboxInsertBeamStatsRequest) HasUnixtime() bool`

HasUnixtime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


