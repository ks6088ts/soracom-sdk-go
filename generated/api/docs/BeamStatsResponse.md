# BeamStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeamStatsMap** | Pointer to [**map[string]SoracomBeamStats**](SoracomBeamStats.md) |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**Unixtime** | Pointer to **int64** |  | [optional] 

## Methods

### NewBeamStatsResponse

`func NewBeamStatsResponse() *BeamStatsResponse`

NewBeamStatsResponse instantiates a new BeamStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeamStatsResponseWithDefaults

`func NewBeamStatsResponseWithDefaults() *BeamStatsResponse`

NewBeamStatsResponseWithDefaults instantiates a new BeamStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeamStatsMap

`func (o *BeamStatsResponse) GetBeamStatsMap() map[string]SoracomBeamStats`

GetBeamStatsMap returns the BeamStatsMap field if non-nil, zero value otherwise.

### GetBeamStatsMapOk

`func (o *BeamStatsResponse) GetBeamStatsMapOk() (*map[string]SoracomBeamStats, bool)`

GetBeamStatsMapOk returns a tuple with the BeamStatsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeamStatsMap

`func (o *BeamStatsResponse) SetBeamStatsMap(v map[string]SoracomBeamStats)`

SetBeamStatsMap sets BeamStatsMap field to given value.

### HasBeamStatsMap

`func (o *BeamStatsResponse) HasBeamStatsMap() bool`

HasBeamStatsMap returns a boolean if a field has been set.

### GetDate

`func (o *BeamStatsResponse) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BeamStatsResponse) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BeamStatsResponse) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *BeamStatsResponse) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetUnixtime

`func (o *BeamStatsResponse) GetUnixtime() int64`

GetUnixtime returns the Unixtime field if non-nil, zero value otherwise.

### GetUnixtimeOk

`func (o *BeamStatsResponse) GetUnixtimeOk() (*int64, bool)`

GetUnixtimeOk returns a tuple with the Unixtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixtime

`func (o *BeamStatsResponse) SetUnixtime(v int64)`

SetUnixtime sets Unixtime field to given value.

### HasUnixtime

`func (o *BeamStatsResponse) HasUnixtime() bool`

HasUnixtime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


