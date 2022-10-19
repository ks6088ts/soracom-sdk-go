# SoraCamImageExportUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemainingFrames** | Pointer to **int32** | 今月中にエクスポート可能な静止画の残りの枚数 | [optional] 
**UsedFrames** | Pointer to **int32** | 今月消費した動画の視聴可能時間をすべて静止画でエクスポートしたと仮定した場合の静止画の枚数 | [optional] 

## Methods

### NewSoraCamImageExportUsage

`func NewSoraCamImageExportUsage() *SoraCamImageExportUsage`

NewSoraCamImageExportUsage instantiates a new SoraCamImageExportUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoraCamImageExportUsageWithDefaults

`func NewSoraCamImageExportUsageWithDefaults() *SoraCamImageExportUsage`

NewSoraCamImageExportUsageWithDefaults instantiates a new SoraCamImageExportUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemainingFrames

`func (o *SoraCamImageExportUsage) GetRemainingFrames() int32`

GetRemainingFrames returns the RemainingFrames field if non-nil, zero value otherwise.

### GetRemainingFramesOk

`func (o *SoraCamImageExportUsage) GetRemainingFramesOk() (*int32, bool)`

GetRemainingFramesOk returns a tuple with the RemainingFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingFrames

`func (o *SoraCamImageExportUsage) SetRemainingFrames(v int32)`

SetRemainingFrames sets RemainingFrames field to given value.

### HasRemainingFrames

`func (o *SoraCamImageExportUsage) HasRemainingFrames() bool`

HasRemainingFrames returns a boolean if a field has been set.

### GetUsedFrames

`func (o *SoraCamImageExportUsage) GetUsedFrames() int32`

GetUsedFrames returns the UsedFrames field if non-nil, zero value otherwise.

### GetUsedFramesOk

`func (o *SoraCamImageExportUsage) GetUsedFramesOk() (*int32, bool)`

GetUsedFramesOk returns a tuple with the UsedFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedFrames

`func (o *SoraCamImageExportUsage) SetUsedFrames(v int32)`

SetUsedFrames sets UsedFrames field to given value.

### HasUsedFrames

`func (o *SoraCamImageExportUsage) HasUsedFrames() bool`

HasUsedFrames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


