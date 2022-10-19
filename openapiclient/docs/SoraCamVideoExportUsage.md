# SoraCamVideoExportUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemainingSeconds** | Pointer to **int32** | 今月中にエクスポート可能な録画映像の残りの秒数 | [optional] 
**UsedSeconds** | Pointer to **int32** | 今月消費した動画の視聴可能時間をすべて録画映像でエクスポートしたと仮定した場合の、エクスポート可能な秒数 | [optional] 

## Methods

### NewSoraCamVideoExportUsage

`func NewSoraCamVideoExportUsage() *SoraCamVideoExportUsage`

NewSoraCamVideoExportUsage instantiates a new SoraCamVideoExportUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoraCamVideoExportUsageWithDefaults

`func NewSoraCamVideoExportUsageWithDefaults() *SoraCamVideoExportUsage`

NewSoraCamVideoExportUsageWithDefaults instantiates a new SoraCamVideoExportUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemainingSeconds

`func (o *SoraCamVideoExportUsage) GetRemainingSeconds() int32`

GetRemainingSeconds returns the RemainingSeconds field if non-nil, zero value otherwise.

### GetRemainingSecondsOk

`func (o *SoraCamVideoExportUsage) GetRemainingSecondsOk() (*int32, bool)`

GetRemainingSecondsOk returns a tuple with the RemainingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingSeconds

`func (o *SoraCamVideoExportUsage) SetRemainingSeconds(v int32)`

SetRemainingSeconds sets RemainingSeconds field to given value.

### HasRemainingSeconds

`func (o *SoraCamVideoExportUsage) HasRemainingSeconds() bool`

HasRemainingSeconds returns a boolean if a field has been set.

### GetUsedSeconds

`func (o *SoraCamVideoExportUsage) GetUsedSeconds() int32`

GetUsedSeconds returns the UsedSeconds field if non-nil, zero value otherwise.

### GetUsedSecondsOk

`func (o *SoraCamVideoExportUsage) GetUsedSecondsOk() (*int32, bool)`

GetUsedSecondsOk returns a tuple with the UsedSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSeconds

`func (o *SoraCamVideoExportUsage) SetUsedSeconds(v int32)`

SetUsedSeconds sets UsedSeconds field to given value.

### HasUsedSeconds

`func (o *SoraCamVideoExportUsage) HasUsedSeconds() bool`

HasUsedSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


