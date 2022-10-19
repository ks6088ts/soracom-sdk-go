# FunkStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunkStatsMap** | Pointer to [**map[string]SoracomFunkStats**](SoracomFunkStats.md) |  | [optional] 
**Unixtime** | Pointer to **int64** |  | [optional] 

## Methods

### NewFunkStatsResponse

`func NewFunkStatsResponse() *FunkStatsResponse`

NewFunkStatsResponse instantiates a new FunkStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunkStatsResponseWithDefaults

`func NewFunkStatsResponseWithDefaults() *FunkStatsResponse`

NewFunkStatsResponseWithDefaults instantiates a new FunkStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunkStatsMap

`func (o *FunkStatsResponse) GetFunkStatsMap() map[string]SoracomFunkStats`

GetFunkStatsMap returns the FunkStatsMap field if non-nil, zero value otherwise.

### GetFunkStatsMapOk

`func (o *FunkStatsResponse) GetFunkStatsMapOk() (*map[string]SoracomFunkStats, bool)`

GetFunkStatsMapOk returns a tuple with the FunkStatsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunkStatsMap

`func (o *FunkStatsResponse) SetFunkStatsMap(v map[string]SoracomFunkStats)`

SetFunkStatsMap sets FunkStatsMap field to given value.

### HasFunkStatsMap

`func (o *FunkStatsResponse) HasFunkStatsMap() bool`

HasFunkStatsMap returns a boolean if a field has been set.

### GetUnixtime

`func (o *FunkStatsResponse) GetUnixtime() int64`

GetUnixtime returns the Unixtime field if non-nil, zero value otherwise.

### GetUnixtimeOk

`func (o *FunkStatsResponse) GetUnixtimeOk() (*int64, bool)`

GetUnixtimeOk returns a tuple with the Unixtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixtime

`func (o *FunkStatsResponse) SetUnixtime(v int64)`

SetUnixtime sets Unixtime field to given value.

### HasUnixtime

`func (o *FunkStatsResponse) HasUnixtime() bool`

HasUnixtime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


