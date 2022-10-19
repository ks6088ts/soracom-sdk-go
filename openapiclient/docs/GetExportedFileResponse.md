# GetExportedFileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | ファイル出力ステータス | [optional] 
**Url** | Pointer to **string** | ファイルダウンロード URL | [optional] 

## Methods

### NewGetExportedFileResponse

`func NewGetExportedFileResponse() *GetExportedFileResponse`

NewGetExportedFileResponse instantiates a new GetExportedFileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExportedFileResponseWithDefaults

`func NewGetExportedFileResponseWithDefaults() *GetExportedFileResponse`

NewGetExportedFileResponseWithDefaults instantiates a new GetExportedFileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetExportedFileResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetExportedFileResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetExportedFileResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetExportedFileResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *GetExportedFileResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetExportedFileResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetExportedFileResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetExportedFileResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


