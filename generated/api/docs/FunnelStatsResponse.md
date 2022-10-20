# FunnelStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunnelStatsMap** | Pointer to [**map[string]SoracomFunnelStats**](SoracomFunnelStats.md) |  | [optional] 
**Unixtime** | Pointer to **int64** |  | [optional] 

## Methods

### NewFunnelStatsResponse

`func NewFunnelStatsResponse() *FunnelStatsResponse`

NewFunnelStatsResponse instantiates a new FunnelStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunnelStatsResponseWithDefaults

`func NewFunnelStatsResponseWithDefaults() *FunnelStatsResponse`

NewFunnelStatsResponseWithDefaults instantiates a new FunnelStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunnelStatsMap

`func (o *FunnelStatsResponse) GetFunnelStatsMap() map[string]SoracomFunnelStats`

GetFunnelStatsMap returns the FunnelStatsMap field if non-nil, zero value otherwise.

### GetFunnelStatsMapOk

`func (o *FunnelStatsResponse) GetFunnelStatsMapOk() (*map[string]SoracomFunnelStats, bool)`

GetFunnelStatsMapOk returns a tuple with the FunnelStatsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelStatsMap

`func (o *FunnelStatsResponse) SetFunnelStatsMap(v map[string]SoracomFunnelStats)`

SetFunnelStatsMap sets FunnelStatsMap field to given value.

### HasFunnelStatsMap

`func (o *FunnelStatsResponse) HasFunnelStatsMap() bool`

HasFunnelStatsMap returns a boolean if a field has been set.

### GetUnixtime

`func (o *FunnelStatsResponse) GetUnixtime() int64`

GetUnixtime returns the Unixtime field if non-nil, zero value otherwise.

### GetUnixtimeOk

`func (o *FunnelStatsResponse) GetUnixtimeOk() (*int64, bool)`

GetUnixtimeOk returns a tuple with the Unixtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixtime

`func (o *FunnelStatsResponse) SetUnixtime(v int64)`

SetUnixtime sets Unixtime field to given value.

### HasUnixtime

`func (o *FunnelStatsResponse) HasUnixtime() bool`

HasUnixtime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


