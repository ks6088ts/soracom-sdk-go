# FileExportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportedFileId** | Pointer to **string** | ファイルエクスポート ID。この ID を [&#x60;Files:getExportedFile API&#x60;](#/Files/getExportedFile) を呼び出すときに指定すると、ファイルをダウンロードするための URL を取得できます。&#x60;export_mode&#x60; に &#x60;async&#x60; を指定した場合のみ含まれます。 | [optional] 
**ExportedFilePath** | Pointer to **string** | 出力済みファイルのパス。&#x60;export_mode&#x60; に &#x60;async&#x60; を指定した場合のみ含まれます。 | [optional] 
**Url** | Pointer to **string** | ファイルをダウンロードするための URL。&#x60;export_mode&#x60; に &#x60;sync&#x60; を指定した場合のみ含まれます。 | [optional] 

## Methods

### NewFileExportResponse

`func NewFileExportResponse() *FileExportResponse`

NewFileExportResponse instantiates a new FileExportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileExportResponseWithDefaults

`func NewFileExportResponseWithDefaults() *FileExportResponse`

NewFileExportResponseWithDefaults instantiates a new FileExportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportedFileId

`func (o *FileExportResponse) GetExportedFileId() string`

GetExportedFileId returns the ExportedFileId field if non-nil, zero value otherwise.

### GetExportedFileIdOk

`func (o *FileExportResponse) GetExportedFileIdOk() (*string, bool)`

GetExportedFileIdOk returns a tuple with the ExportedFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedFileId

`func (o *FileExportResponse) SetExportedFileId(v string)`

SetExportedFileId sets ExportedFileId field to given value.

### HasExportedFileId

`func (o *FileExportResponse) HasExportedFileId() bool`

HasExportedFileId returns a boolean if a field has been set.

### GetExportedFilePath

`func (o *FileExportResponse) GetExportedFilePath() string`

GetExportedFilePath returns the ExportedFilePath field if non-nil, zero value otherwise.

### GetExportedFilePathOk

`func (o *FileExportResponse) GetExportedFilePathOk() (*string, bool)`

GetExportedFilePathOk returns a tuple with the ExportedFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedFilePath

`func (o *FileExportResponse) SetExportedFilePath(v string)`

SetExportedFilePath sets ExportedFilePath field to given value.

### HasExportedFilePath

`func (o *FileExportResponse) HasExportedFilePath() bool`

HasExportedFilePath returns a boolean if a field has been set.

### GetUrl

`func (o *FileExportResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FileExportResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FileExportResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FileExportResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


