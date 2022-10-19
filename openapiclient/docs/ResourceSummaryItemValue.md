# ResourceSummaryItemValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | リソースの要約アイテムの集計値の名前  - &#x60;resourceSummaryType&#x60; が &#x60;simsPerStatus&#x60; の場合:   - &#x60;readySims&#x60;: ステータスが「準備完了 (ready)」の IoT SIM の数   - &#x60;activeSims&#x60;: ステータスが「使用中 (active)」の IoT SIM の数   - &#x60;inactiveSims&#x60;: ステータスが「休止中 (inactive)」の IoT SIM の数   - &#x60;standbySims&#x60;: ステータスが「利用開始待ち (standby)」の IoT SIM の数   - &#x60;suspendedSims&#x60;: ステータスが「利用中断中 (suspended)」の IoT SIM の数  | [optional] 
**Value** | Pointer to **float32** | リソースの要約アイテムの集計値の値 | [optional] 

## Methods

### NewResourceSummaryItemValue

`func NewResourceSummaryItemValue() *ResourceSummaryItemValue`

NewResourceSummaryItemValue instantiates a new ResourceSummaryItemValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSummaryItemValueWithDefaults

`func NewResourceSummaryItemValueWithDefaults() *ResourceSummaryItemValue`

NewResourceSummaryItemValueWithDefaults instantiates a new ResourceSummaryItemValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResourceSummaryItemValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceSummaryItemValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceSummaryItemValue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceSummaryItemValue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ResourceSummaryItemValue) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResourceSummaryItemValue) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResourceSummaryItemValue) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ResourceSummaryItemValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


