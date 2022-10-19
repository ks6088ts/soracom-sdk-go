# SoraCamImageExportInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** | ソラカメ対応カメラのデバイス ID | [optional] 
**ExpiryTime** | Pointer to **int64** | URL の有効期限 (UNIX 時間 (ミリ秒))。&#x60;status&#x60; が &#x60;completed&#x60; の場合のみ含まれます。 | [optional] 
**ExportId** | Pointer to **string** | エクスポート ID。[&#x60;SoraCam:exportSoraCamDeviceRecordedImage API&#x60;](/ja-jp/tools/api/reference/#/SoraCam/exportSoraCamDeviceRecordedImage) で取得したエクスポート ID を、[&#x60;SoraCam:getSoraCamDeviceExportedImage API&#x60;](/ja-jp/tools/api/reference/#/SoraCam/getSoraCamDeviceExportedImage) で指定すると、jpg ファイルをダウンロードするための URL を取得できます。  | [optional] 
**OperatorId** | Pointer to **string** | [&#x60;SoraCam:exportSoraCamDeviceRecordedImage API&#x60;](/ja-jp/tools/api/reference/#/SoraCam/exportSoraCamDeviceRecordedImage) を呼び出したオペレーターの ID  | [optional] 
**RequestedTime** | Pointer to **int64** | [&#x60;SoraCam:exportSoraCamDeviceRecordedImage API&#x60;](/ja-jp/tools/api/reference/#/SoraCam/exportSoraCamDeviceRecordedImage) によるエクスポートのリクエストを SORACOM プラットフォームが受け付けた日時 (UNIX 時間 (ミリ秒))  | [optional] 
**Status** | Pointer to **string** | エクスポート処理の現在の状況。  - &#x60;initializing&#x60;: 初期状態 - &#x60;processing&#x60;: 処理中 - &#x60;completed&#x60;: エクスポート完了 - &#x60;failed&#x60;: エクスポート失敗 - &#x60;expired&#x60;: URL の有効期限切れ  | [optional] 
**Url** | Pointer to **string** | エクスポートされた jpg ファイルをダウンロードするための URL。&#x60;status&#x60; が &#x60;completed&#x60; の場合のみ含まれます。 | [optional] 

## Methods

### NewSoraCamImageExportInfo

`func NewSoraCamImageExportInfo() *SoraCamImageExportInfo`

NewSoraCamImageExportInfo instantiates a new SoraCamImageExportInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoraCamImageExportInfoWithDefaults

`func NewSoraCamImageExportInfoWithDefaults() *SoraCamImageExportInfo`

NewSoraCamImageExportInfoWithDefaults instantiates a new SoraCamImageExportInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *SoraCamImageExportInfo) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SoraCamImageExportInfo) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SoraCamImageExportInfo) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *SoraCamImageExportInfo) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetExpiryTime

`func (o *SoraCamImageExportInfo) GetExpiryTime() int64`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *SoraCamImageExportInfo) GetExpiryTimeOk() (*int64, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *SoraCamImageExportInfo) SetExpiryTime(v int64)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *SoraCamImageExportInfo) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetExportId

`func (o *SoraCamImageExportInfo) GetExportId() string`

GetExportId returns the ExportId field if non-nil, zero value otherwise.

### GetExportIdOk

`func (o *SoraCamImageExportInfo) GetExportIdOk() (*string, bool)`

GetExportIdOk returns a tuple with the ExportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportId

`func (o *SoraCamImageExportInfo) SetExportId(v string)`

SetExportId sets ExportId field to given value.

### HasExportId

`func (o *SoraCamImageExportInfo) HasExportId() bool`

HasExportId returns a boolean if a field has been set.

### GetOperatorId

`func (o *SoraCamImageExportInfo) GetOperatorId() string`

GetOperatorId returns the OperatorId field if non-nil, zero value otherwise.

### GetOperatorIdOk

`func (o *SoraCamImageExportInfo) GetOperatorIdOk() (*string, bool)`

GetOperatorIdOk returns a tuple with the OperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorId

`func (o *SoraCamImageExportInfo) SetOperatorId(v string)`

SetOperatorId sets OperatorId field to given value.

### HasOperatorId

`func (o *SoraCamImageExportInfo) HasOperatorId() bool`

HasOperatorId returns a boolean if a field has been set.

### GetRequestedTime

`func (o *SoraCamImageExportInfo) GetRequestedTime() int64`

GetRequestedTime returns the RequestedTime field if non-nil, zero value otherwise.

### GetRequestedTimeOk

`func (o *SoraCamImageExportInfo) GetRequestedTimeOk() (*int64, bool)`

GetRequestedTimeOk returns a tuple with the RequestedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTime

`func (o *SoraCamImageExportInfo) SetRequestedTime(v int64)`

SetRequestedTime sets RequestedTime field to given value.

### HasRequestedTime

`func (o *SoraCamImageExportInfo) HasRequestedTime() bool`

HasRequestedTime returns a boolean if a field has been set.

### GetStatus

`func (o *SoraCamImageExportInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SoraCamImageExportInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SoraCamImageExportInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SoraCamImageExportInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *SoraCamImageExportInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SoraCamImageExportInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SoraCamImageExportInfo) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SoraCamImageExportInfo) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


