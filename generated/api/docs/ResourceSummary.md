# ResourceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceSummaryItems** | Pointer to [**[]ResourceSummaryItem**](ResourceSummaryItem.md) | リソースの要約アイテムの一覧 | [optional] 
**ResourceSummaryType** | Pointer to **string** | リソースの要約の種別  - &#x60;simsPerStatus&#x60;: ステータスごとの IoT SIM の数  | [optional] 

## Methods

### NewResourceSummary

`func NewResourceSummary() *ResourceSummary`

NewResourceSummary instantiates a new ResourceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSummaryWithDefaults

`func NewResourceSummaryWithDefaults() *ResourceSummary`

NewResourceSummaryWithDefaults instantiates a new ResourceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceSummaryItems

`func (o *ResourceSummary) GetResourceSummaryItems() []ResourceSummaryItem`

GetResourceSummaryItems returns the ResourceSummaryItems field if non-nil, zero value otherwise.

### GetResourceSummaryItemsOk

`func (o *ResourceSummary) GetResourceSummaryItemsOk() (*[]ResourceSummaryItem, bool)`

GetResourceSummaryItemsOk returns a tuple with the ResourceSummaryItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSummaryItems

`func (o *ResourceSummary) SetResourceSummaryItems(v []ResourceSummaryItem)`

SetResourceSummaryItems sets ResourceSummaryItems field to given value.

### HasResourceSummaryItems

`func (o *ResourceSummary) HasResourceSummaryItems() bool`

HasResourceSummaryItems returns a boolean if a field has been set.

### GetResourceSummaryType

`func (o *ResourceSummary) GetResourceSummaryType() string`

GetResourceSummaryType returns the ResourceSummaryType field if non-nil, zero value otherwise.

### GetResourceSummaryTypeOk

`func (o *ResourceSummary) GetResourceSummaryTypeOk() (*string, bool)`

GetResourceSummaryTypeOk returns a tuple with the ResourceSummaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSummaryType

`func (o *ResourceSummary) SetResourceSummaryType(v string)`

SetResourceSummaryType sets ResourceSummaryType field to given value.

### HasResourceSummaryType

`func (o *ResourceSummary) HasResourceSummaryType() bool`

HasResourceSummaryType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


