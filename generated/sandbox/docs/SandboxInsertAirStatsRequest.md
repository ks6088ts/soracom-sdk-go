# SandboxInsertAirStatsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTrafficStatsMap** | Pointer to [**SandboxInsertAirStatsRequestDataTrafficStatsMap**](SandboxInsertAirStatsRequestDataTrafficStatsMap.md) |  | [optional] 
**Unixtime** | Pointer to **int64** | UNIX 時刻（ミリ秒単位） | [optional] 

## Methods

### NewSandboxInsertAirStatsRequest

`func NewSandboxInsertAirStatsRequest() *SandboxInsertAirStatsRequest`

NewSandboxInsertAirStatsRequest instantiates a new SandboxInsertAirStatsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxInsertAirStatsRequestWithDefaults

`func NewSandboxInsertAirStatsRequestWithDefaults() *SandboxInsertAirStatsRequest`

NewSandboxInsertAirStatsRequestWithDefaults instantiates a new SandboxInsertAirStatsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTrafficStatsMap

`func (o *SandboxInsertAirStatsRequest) GetDataTrafficStatsMap() SandboxInsertAirStatsRequestDataTrafficStatsMap`

GetDataTrafficStatsMap returns the DataTrafficStatsMap field if non-nil, zero value otherwise.

### GetDataTrafficStatsMapOk

`func (o *SandboxInsertAirStatsRequest) GetDataTrafficStatsMapOk() (*SandboxInsertAirStatsRequestDataTrafficStatsMap, bool)`

GetDataTrafficStatsMapOk returns a tuple with the DataTrafficStatsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTrafficStatsMap

`func (o *SandboxInsertAirStatsRequest) SetDataTrafficStatsMap(v SandboxInsertAirStatsRequestDataTrafficStatsMap)`

SetDataTrafficStatsMap sets DataTrafficStatsMap field to given value.

### HasDataTrafficStatsMap

`func (o *SandboxInsertAirStatsRequest) HasDataTrafficStatsMap() bool`

HasDataTrafficStatsMap returns a boolean if a field has been set.

### GetUnixtime

`func (o *SandboxInsertAirStatsRequest) GetUnixtime() int64`

GetUnixtime returns the Unixtime field if non-nil, zero value otherwise.

### GetUnixtimeOk

`func (o *SandboxInsertAirStatsRequest) GetUnixtimeOk() (*int64, bool)`

GetUnixtimeOk returns a tuple with the Unixtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixtime

`func (o *SandboxInsertAirStatsRequest) SetUnixtime(v int64)`

SetUnixtime sets Unixtime field to given value.

### HasUnixtime

`func (o *SandboxInsertAirStatsRequest) HasUnixtime() bool`

HasUnixtime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


