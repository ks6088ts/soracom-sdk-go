# SoraCamStreamingVideo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayList** | Pointer to [**[]SoraCamPlayListEntry**](SoraCamPlayListEntry.md) | ストリーミング映像 (リアルタイム映像 / 録画映像) のプレイリストエントリーの配列  - リアルタイム映像の場合は、返却される動画ファイルは常にひとつです。また、&#x60;from&#x60; および &#x60;to&#x60; は省略されます。 - 録画映像の場合は、録画の状態によって動画ファイルが分割されていることがあります。  | [optional] 

## Methods

### NewSoraCamStreamingVideo

`func NewSoraCamStreamingVideo() *SoraCamStreamingVideo`

NewSoraCamStreamingVideo instantiates a new SoraCamStreamingVideo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoraCamStreamingVideoWithDefaults

`func NewSoraCamStreamingVideoWithDefaults() *SoraCamStreamingVideo`

NewSoraCamStreamingVideoWithDefaults instantiates a new SoraCamStreamingVideo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayList

`func (o *SoraCamStreamingVideo) GetPlayList() []SoraCamPlayListEntry`

GetPlayList returns the PlayList field if non-nil, zero value otherwise.

### GetPlayListOk

`func (o *SoraCamStreamingVideo) GetPlayListOk() (*[]SoraCamPlayListEntry, bool)`

GetPlayListOk returns a tuple with the PlayList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayList

`func (o *SoraCamStreamingVideo) SetPlayList(v []SoraCamPlayListEntry)`

SetPlayList sets PlayList field to given value.

### HasPlayList

`func (o *SoraCamStreamingVideo) HasPlayList() bool`

HasPlayList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


