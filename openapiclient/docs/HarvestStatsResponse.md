# HarvestStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HarvestStatsMap** | Pointer to [**map[string]SoracomHarvestStats**](SoracomHarvestStats.md) |  | [optional] 
**Unixtime** | Pointer to **int64** |  | [optional] 

## Methods

### NewHarvestStatsResponse

`func NewHarvestStatsResponse() *HarvestStatsResponse`

NewHarvestStatsResponse instantiates a new HarvestStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvestStatsResponseWithDefaults

`func NewHarvestStatsResponseWithDefaults() *HarvestStatsResponse`

NewHarvestStatsResponseWithDefaults instantiates a new HarvestStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHarvestStatsMap

`func (o *HarvestStatsResponse) GetHarvestStatsMap() map[string]SoracomHarvestStats`

GetHarvestStatsMap returns the HarvestStatsMap field if non-nil, zero value otherwise.

### GetHarvestStatsMapOk

`func (o *HarvestStatsResponse) GetHarvestStatsMapOk() (*map[string]SoracomHarvestStats, bool)`

GetHarvestStatsMapOk returns a tuple with the HarvestStatsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarvestStatsMap

`func (o *HarvestStatsResponse) SetHarvestStatsMap(v map[string]SoracomHarvestStats)`

SetHarvestStatsMap sets HarvestStatsMap field to given value.

### HasHarvestStatsMap

`func (o *HarvestStatsResponse) HasHarvestStatsMap() bool`

HasHarvestStatsMap returns a boolean if a field has been set.

### GetUnixtime

`func (o *HarvestStatsResponse) GetUnixtime() int64`

GetUnixtime returns the Unixtime field if non-nil, zero value otherwise.

### GetUnixtimeOk

`func (o *HarvestStatsResponse) GetUnixtimeOk() (*int64, bool)`

GetUnixtimeOk returns a tuple with the Unixtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixtime

`func (o *HarvestStatsResponse) SetUnixtime(v int64)`

SetUnixtime sets Unixtime field to given value.

### HasUnixtime

`func (o *HarvestStatsResponse) HasUnixtime() bool`

HasUnixtime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


