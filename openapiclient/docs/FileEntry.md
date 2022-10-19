# FileEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentLength** | Pointer to **int64** | コンテントの長さ | [optional] 
**ContentType** | Pointer to **string** | コンテントタイプ | [optional] 
**CreatedTime** | Pointer to **int64** | ファイルの作成時刻 | [optional] 
**Directory** | Pointer to **string** | 親ディレクトリ名 | [optional] 
**Etag** | Pointer to **string** | ファイルの ETag | [optional] 
**FilePath** | Pointer to **string** | ファイルの絶対パス | [optional] 
**Filename** | Pointer to **string** | ファイル名 | [optional] 
**IsDirectory** | Pointer to **bool** | ディレクトリか否か | [optional] 
**LastModifiedTime** | Pointer to **int64** | ファイルの更新時刻 | [optional] 

## Methods

### NewFileEntry

`func NewFileEntry() *FileEntry`

NewFileEntry instantiates a new FileEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileEntryWithDefaults

`func NewFileEntryWithDefaults() *FileEntry`

NewFileEntryWithDefaults instantiates a new FileEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentLength

`func (o *FileEntry) GetContentLength() int64`

GetContentLength returns the ContentLength field if non-nil, zero value otherwise.

### GetContentLengthOk

`func (o *FileEntry) GetContentLengthOk() (*int64, bool)`

GetContentLengthOk returns a tuple with the ContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLength

`func (o *FileEntry) SetContentLength(v int64)`

SetContentLength sets ContentLength field to given value.

### HasContentLength

`func (o *FileEntry) HasContentLength() bool`

HasContentLength returns a boolean if a field has been set.

### GetContentType

`func (o *FileEntry) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *FileEntry) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *FileEntry) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *FileEntry) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetCreatedTime

`func (o *FileEntry) GetCreatedTime() int64`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *FileEntry) GetCreatedTimeOk() (*int64, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *FileEntry) SetCreatedTime(v int64)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *FileEntry) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDirectory

`func (o *FileEntry) GetDirectory() string`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *FileEntry) GetDirectoryOk() (*string, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *FileEntry) SetDirectory(v string)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *FileEntry) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### GetEtag

`func (o *FileEntry) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *FileEntry) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *FileEntry) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *FileEntry) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetFilePath

`func (o *FileEntry) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *FileEntry) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *FileEntry) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *FileEntry) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetFilename

`func (o *FileEntry) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *FileEntry) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *FileEntry) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *FileEntry) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetIsDirectory

`func (o *FileEntry) GetIsDirectory() bool`

GetIsDirectory returns the IsDirectory field if non-nil, zero value otherwise.

### GetIsDirectoryOk

`func (o *FileEntry) GetIsDirectoryOk() (*bool, bool)`

GetIsDirectoryOk returns a tuple with the IsDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectory

`func (o *FileEntry) SetIsDirectory(v bool)`

SetIsDirectory sets IsDirectory field to given value.

### HasIsDirectory

`func (o *FileEntry) HasIsDirectory() bool`

HasIsDirectory returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *FileEntry) GetLastModifiedTime() int64`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *FileEntry) GetLastModifiedTimeOk() (*int64, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *FileEntry) SetLastModifiedTime(v int64)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *FileEntry) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


