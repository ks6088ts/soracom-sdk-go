# SoraCamVideoExportInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** | ソラカメ対応カメラのデバイス ID | [optional] 
**ExpiryTime** | Pointer to **int64** | URL の有効期限 (UNIX 時間 (ミリ秒)) | [optional] 
**ExportId** | Pointer to **string** | エクスポート ID。[&#x60;SoraCam:exportSoraCamDeviceRecordedVideo API&#x60;](/ja-jp/tools/api/reference/#/SoraCam/exportSoraCamDeviceRecordedVideo) で取得したエクスポート ID を、[&#x60;SoraCam:getSoraCamDeviceExportedVideo API&#x60;](/ja-jp/tools/api/reference/#/SoraCam/getSoraCamDeviceExportedVideo) で指定すると、mp4 ファイルを zip 形式で圧縮したファイルをダウンロードするための URL を取得できます。  | [optional] 
**OperatorId** | Pointer to **string** | [&#x60;SoraCam:exportSoraCamDeviceRecordedVideo API&#x60;](/ja-jp/tools/api/reference/#/SoraCam/exportSoraCamDeviceRecordedVideo) を呼び出したオペレーターの ID  | [optional] 
**RequestedTime** | Pointer to **int64** | [&#x60;SoraCam:exportSoraCamDeviceRecordedVideo API&#x60;](/ja-jp/tools/api/reference/#/SoraCam/exportSoraCamDeviceRecordedVideo) によるエクスポートのリクエストを SORACOM プラットフォームが受け付けた日時 (UNIX 時間 (ミリ秒))  | [optional] 
**Status** | Pointer to **string** | エクスポート処理の現在の状況。  - &#x60;initializing&#x60;: 初期状態 - &#x60;processing&#x60;: 処理中 - &#x60;completed&#x60;: エクスポート完了 - &#x60;failed&#x60;: エクスポート失敗 - &#x60;expired&#x60;: URL の有効期限切れ  | [optional] 
**Url** | Pointer to **string** | mp4 ファイルを zip 形式で圧縮したファイルをダウンロードするための URL。&#x60;status&#x60; が &#x60;completed&#x60; の場合のみ含まれます。  - ダウンロードした zip 形式のファイルを展開すると、mp4 ファイルを取得できます。  | [optional] 

## Methods

### NewSoraCamVideoExportInfo

`func NewSoraCamVideoExportInfo() *SoraCamVideoExportInfo`

NewSoraCamVideoExportInfo instantiates a new SoraCamVideoExportInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoraCamVideoExportInfoWithDefaults

`func NewSoraCamVideoExportInfoWithDefaults() *SoraCamVideoExportInfo`

NewSoraCamVideoExportInfoWithDefaults instantiates a new SoraCamVideoExportInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *SoraCamVideoExportInfo) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SoraCamVideoExportInfo) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SoraCamVideoExportInfo) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *SoraCamVideoExportInfo) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetExpiryTime

`func (o *SoraCamVideoExportInfo) GetExpiryTime() int64`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *SoraCamVideoExportInfo) GetExpiryTimeOk() (*int64, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *SoraCamVideoExportInfo) SetExpiryTime(v int64)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *SoraCamVideoExportInfo) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetExportId

`func (o *SoraCamVideoExportInfo) GetExportId() string`

GetExportId returns the ExportId field if non-nil, zero value otherwise.

### GetExportIdOk

`func (o *SoraCamVideoExportInfo) GetExportIdOk() (*string, bool)`

GetExportIdOk returns a tuple with the ExportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportId

`func (o *SoraCamVideoExportInfo) SetExportId(v string)`

SetExportId sets ExportId field to given value.

### HasExportId

`func (o *SoraCamVideoExportInfo) HasExportId() bool`

HasExportId returns a boolean if a field has been set.

### GetOperatorId

`func (o *SoraCamVideoExportInfo) GetOperatorId() string`

GetOperatorId returns the OperatorId field if non-nil, zero value otherwise.

### GetOperatorIdOk

`func (o *SoraCamVideoExportInfo) GetOperatorIdOk() (*string, bool)`

GetOperatorIdOk returns a tuple with the OperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorId

`func (o *SoraCamVideoExportInfo) SetOperatorId(v string)`

SetOperatorId sets OperatorId field to given value.

### HasOperatorId

`func (o *SoraCamVideoExportInfo) HasOperatorId() bool`

HasOperatorId returns a boolean if a field has been set.

### GetRequestedTime

`func (o *SoraCamVideoExportInfo) GetRequestedTime() int64`

GetRequestedTime returns the RequestedTime field if non-nil, zero value otherwise.

### GetRequestedTimeOk

`func (o *SoraCamVideoExportInfo) GetRequestedTimeOk() (*int64, bool)`

GetRequestedTimeOk returns a tuple with the RequestedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTime

`func (o *SoraCamVideoExportInfo) SetRequestedTime(v int64)`

SetRequestedTime sets RequestedTime field to given value.

### HasRequestedTime

`func (o *SoraCamVideoExportInfo) HasRequestedTime() bool`

HasRequestedTime returns a boolean if a field has been set.

### GetStatus

`func (o *SoraCamVideoExportInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SoraCamVideoExportInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SoraCamVideoExportInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SoraCamVideoExportInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *SoraCamVideoExportInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SoraCamVideoExportInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SoraCamVideoExportInfo) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SoraCamVideoExportInfo) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


