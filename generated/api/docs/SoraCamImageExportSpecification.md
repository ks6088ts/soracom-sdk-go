# SoraCamImageExportSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageFilters** | Pointer to **[]string** | エクスポートした静止画に適用するフィルタ | [optional] 
**Time** | **int64** | エクスポート対象の時刻 (UNIX 時間 (ミリ秒))。指定した時刻に撮影された録画映像から静止画がエクスポートされます。 | 

## Methods

### NewSoraCamImageExportSpecification

`func NewSoraCamImageExportSpecification(time int64, ) *SoraCamImageExportSpecification`

NewSoraCamImageExportSpecification instantiates a new SoraCamImageExportSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoraCamImageExportSpecificationWithDefaults

`func NewSoraCamImageExportSpecificationWithDefaults() *SoraCamImageExportSpecification`

NewSoraCamImageExportSpecificationWithDefaults instantiates a new SoraCamImageExportSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageFilters

`func (o *SoraCamImageExportSpecification) GetImageFilters() []string`

GetImageFilters returns the ImageFilters field if non-nil, zero value otherwise.

### GetImageFiltersOk

`func (o *SoraCamImageExportSpecification) GetImageFiltersOk() (*[]string, bool)`

GetImageFiltersOk returns a tuple with the ImageFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFilters

`func (o *SoraCamImageExportSpecification) SetImageFilters(v []string)`

SetImageFilters sets ImageFilters field to given value.

### HasImageFilters

`func (o *SoraCamImageExportSpecification) HasImageFilters() bool`

HasImageFilters returns a boolean if a field has been set.

### GetTime

`func (o *SoraCamImageExportSpecification) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *SoraCamImageExportSpecification) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *SoraCamImageExportSpecification) SetTime(v int64)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


