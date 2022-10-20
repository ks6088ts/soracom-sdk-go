# SoraCamPlayListEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **int64** | 録画映像の開始時刻 (UNIX 時間 (ミリ秒))。リアルタイム映像の場合は省略されます。 | [optional] 
**To** | Pointer to **int64** | 録画映像の終了時刻 (UNIX 時間 (ミリ秒))。リアルタイム映像の場合は省略されます。 | [optional] 
**Url** | Pointer to **string** | ストリーミング映像 (リアルタイム映像 / 録画映像) の mpd ファイルを取得するための URL  - 以下の場合は URL が返されません。   - リアルタイム映像を取得しようとした場合に、ソラカメ対応カメラの電源が入っていない。   - リアルタイム映像を取得しようとした場合に、ソラカメ対応カメラでクラウド録画が開始されていない。   - 録画映像を取得しようとした場合に、指定した &#x60;from&#x60; から &#x60;to&#x60; の間に録画映像が存在しない。  | [optional] 

## Methods

### NewSoraCamPlayListEntry

`func NewSoraCamPlayListEntry() *SoraCamPlayListEntry`

NewSoraCamPlayListEntry instantiates a new SoraCamPlayListEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoraCamPlayListEntryWithDefaults

`func NewSoraCamPlayListEntryWithDefaults() *SoraCamPlayListEntry`

NewSoraCamPlayListEntryWithDefaults instantiates a new SoraCamPlayListEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *SoraCamPlayListEntry) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SoraCamPlayListEntry) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SoraCamPlayListEntry) SetFrom(v int64)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SoraCamPlayListEntry) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *SoraCamPlayListEntry) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SoraCamPlayListEntry) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SoraCamPlayListEntry) SetTo(v int64)`

SetTo sets To field to given value.

### HasTo

`func (o *SoraCamPlayListEntry) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetUrl

`func (o *SoraCamPlayListEntry) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SoraCamPlayListEntry) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SoraCamPlayListEntry) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SoraCamPlayListEntry) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


