# SoraCamExportUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** | ソラカメ対応カメラのデバイス ID | [optional] 
**Image** | Pointer to [**SoraCamImageExportUsage**](SoraCamImageExportUsage.md) |  | [optional] 
**MeteredYearMonth** | Pointer to **string** | 静止画の残りのエクスポート可能枚数や、録画映像の残りのエクスポート可能時間などを取得した年月 (フォーマット: &#39;YYYYMM&#39;) | [optional] 
**Video** | Pointer to [**SoraCamVideoExportUsage**](SoraCamVideoExportUsage.md) |  | [optional] 

## Methods

### NewSoraCamExportUsage

`func NewSoraCamExportUsage() *SoraCamExportUsage`

NewSoraCamExportUsage instantiates a new SoraCamExportUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoraCamExportUsageWithDefaults

`func NewSoraCamExportUsageWithDefaults() *SoraCamExportUsage`

NewSoraCamExportUsageWithDefaults instantiates a new SoraCamExportUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *SoraCamExportUsage) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SoraCamExportUsage) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SoraCamExportUsage) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *SoraCamExportUsage) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetImage

`func (o *SoraCamExportUsage) GetImage() SoraCamImageExportUsage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *SoraCamExportUsage) GetImageOk() (*SoraCamImageExportUsage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *SoraCamExportUsage) SetImage(v SoraCamImageExportUsage)`

SetImage sets Image field to given value.

### HasImage

`func (o *SoraCamExportUsage) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetMeteredYearMonth

`func (o *SoraCamExportUsage) GetMeteredYearMonth() string`

GetMeteredYearMonth returns the MeteredYearMonth field if non-nil, zero value otherwise.

### GetMeteredYearMonthOk

`func (o *SoraCamExportUsage) GetMeteredYearMonthOk() (*string, bool)`

GetMeteredYearMonthOk returns a tuple with the MeteredYearMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeteredYearMonth

`func (o *SoraCamExportUsage) SetMeteredYearMonth(v string)`

SetMeteredYearMonth sets MeteredYearMonth field to given value.

### HasMeteredYearMonth

`func (o *SoraCamExportUsage) HasMeteredYearMonth() bool`

HasMeteredYearMonth returns a boolean if a field has been set.

### GetVideo

`func (o *SoraCamExportUsage) GetVideo() SoraCamVideoExportUsage`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *SoraCamExportUsage) GetVideoOk() (*SoraCamVideoExportUsage, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *SoraCamExportUsage) SetVideo(v SoraCamVideoExportUsage)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *SoraCamExportUsage) HasVideo() bool`

HasVideo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


