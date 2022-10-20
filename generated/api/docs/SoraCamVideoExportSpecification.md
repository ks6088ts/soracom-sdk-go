# SoraCamVideoExportSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **int64** | エクスポート対象の開始時刻 (UNIX 時間 (ミリ秒)) | 
**To** | **int64** | エクスポート対象の終了時刻 (UNIX 時間 (ミリ秒))  - 一回の API 呼び出しでエクスポートできる時間は最大 300 秒 (5 分) です。&#x60;from&#x60; と &#x60;to&#x60; の差が、300 秒を超えないようにしてください。  | 

## Methods

### NewSoraCamVideoExportSpecification

`func NewSoraCamVideoExportSpecification(from int64, to int64, ) *SoraCamVideoExportSpecification`

NewSoraCamVideoExportSpecification instantiates a new SoraCamVideoExportSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoraCamVideoExportSpecificationWithDefaults

`func NewSoraCamVideoExportSpecificationWithDefaults() *SoraCamVideoExportSpecification`

NewSoraCamVideoExportSpecificationWithDefaults instantiates a new SoraCamVideoExportSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *SoraCamVideoExportSpecification) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SoraCamVideoExportSpecification) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SoraCamVideoExportSpecification) SetFrom(v int64)`

SetFrom sets From field to given value.


### GetTo

`func (o *SoraCamVideoExportSpecification) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SoraCamVideoExportSpecification) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SoraCamVideoExportSpecification) SetTo(v int64)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


