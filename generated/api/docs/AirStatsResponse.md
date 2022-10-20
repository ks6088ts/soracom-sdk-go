# AirStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTrafficStatsMap** | Pointer to [**map[string]DataTrafficStats**](DataTrafficStats.md) |  | [optional] 
**Date** | Pointer to **string** | 日時 (協定世界時 (UTC)) | [optional] 
**Unixtime** | Pointer to **int64** |  | [optional] 

## Methods

### NewAirStatsResponse

`func NewAirStatsResponse() *AirStatsResponse`

NewAirStatsResponse instantiates a new AirStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirStatsResponseWithDefaults

`func NewAirStatsResponseWithDefaults() *AirStatsResponse`

NewAirStatsResponseWithDefaults instantiates a new AirStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTrafficStatsMap

`func (o *AirStatsResponse) GetDataTrafficStatsMap() map[string]DataTrafficStats`

GetDataTrafficStatsMap returns the DataTrafficStatsMap field if non-nil, zero value otherwise.

### GetDataTrafficStatsMapOk

`func (o *AirStatsResponse) GetDataTrafficStatsMapOk() (*map[string]DataTrafficStats, bool)`

GetDataTrafficStatsMapOk returns a tuple with the DataTrafficStatsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTrafficStatsMap

`func (o *AirStatsResponse) SetDataTrafficStatsMap(v map[string]DataTrafficStats)`

SetDataTrafficStatsMap sets DataTrafficStatsMap field to given value.

### HasDataTrafficStatsMap

`func (o *AirStatsResponse) HasDataTrafficStatsMap() bool`

HasDataTrafficStatsMap returns a boolean if a field has been set.

### GetDate

`func (o *AirStatsResponse) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AirStatsResponse) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AirStatsResponse) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *AirStatsResponse) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetUnixtime

`func (o *AirStatsResponse) GetUnixtime() int64`

GetUnixtime returns the Unixtime field if non-nil, zero value otherwise.

### GetUnixtimeOk

`func (o *AirStatsResponse) GetUnixtimeOk() (*int64, bool)`

GetUnixtimeOk returns a tuple with the Unixtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixtime

`func (o *AirStatsResponse) SetUnixtime(v int64)`

SetUnixtime sets Unixtime field to given value.

### HasUnixtime

`func (o *AirStatsResponse) HasUnixtime() bool`

HasUnixtime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


