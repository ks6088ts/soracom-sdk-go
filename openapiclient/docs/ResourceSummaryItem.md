# ResourceSummaryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimensions** | Pointer to [**[]ResourceSummaryItemDimension**](ResourceSummaryItemDimension.md) | リソースの要約アイテムのディメンションの一覧 | [optional] 
**UpdatedTime** | Pointer to **int64** | リソースの要約アイテムの最終更新日時 (unixtime ミリ秒単位) | [optional] 
**Values** | Pointer to [**[]ResourceSummaryItemValue**](ResourceSummaryItemValue.md) | リソースの要約アイテムの集計値の一覧 | [optional] 

## Methods

### NewResourceSummaryItem

`func NewResourceSummaryItem() *ResourceSummaryItem`

NewResourceSummaryItem instantiates a new ResourceSummaryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSummaryItemWithDefaults

`func NewResourceSummaryItemWithDefaults() *ResourceSummaryItem`

NewResourceSummaryItemWithDefaults instantiates a new ResourceSummaryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensions

`func (o *ResourceSummaryItem) GetDimensions() []ResourceSummaryItemDimension`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *ResourceSummaryItem) GetDimensionsOk() (*[]ResourceSummaryItemDimension, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *ResourceSummaryItem) SetDimensions(v []ResourceSummaryItemDimension)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *ResourceSummaryItem) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *ResourceSummaryItem) GetUpdatedTime() int64`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *ResourceSummaryItem) GetUpdatedTimeOk() (*int64, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *ResourceSummaryItem) SetUpdatedTime(v int64)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *ResourceSummaryItem) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetValues

`func (o *ResourceSummaryItem) GetValues() []ResourceSummaryItemValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ResourceSummaryItem) GetValuesOk() (*[]ResourceSummaryItemValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ResourceSummaryItem) SetValues(v []ResourceSummaryItemValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *ResourceSummaryItem) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


