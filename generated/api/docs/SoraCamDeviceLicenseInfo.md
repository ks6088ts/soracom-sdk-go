# SoraCamDeviceLicenseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | ソラカメ対応カメラに適用されているライセンスの名前  - &#x60;Cloud Continuous Recording&#x60;: クラウド常時録画ライセンス  | [optional] 
**Value** | Pointer to **string** | ソラカメ対応カメラに適用されているライセンスのプロパティ  &#x60;name&#x60; が &#x60;Cloud Continuous Recording&#x60; の場合:  - &#x60;storage_life&#x60;: 保存期間  | [optional] 

## Methods

### NewSoraCamDeviceLicenseInfo

`func NewSoraCamDeviceLicenseInfo() *SoraCamDeviceLicenseInfo`

NewSoraCamDeviceLicenseInfo instantiates a new SoraCamDeviceLicenseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoraCamDeviceLicenseInfoWithDefaults

`func NewSoraCamDeviceLicenseInfoWithDefaults() *SoraCamDeviceLicenseInfo`

NewSoraCamDeviceLicenseInfoWithDefaults instantiates a new SoraCamDeviceLicenseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SoraCamDeviceLicenseInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SoraCamDeviceLicenseInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SoraCamDeviceLicenseInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SoraCamDeviceLicenseInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *SoraCamDeviceLicenseInfo) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SoraCamDeviceLicenseInfo) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SoraCamDeviceLicenseInfo) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SoraCamDeviceLicenseInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


