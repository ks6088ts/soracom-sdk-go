# ResourceInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Mandatory** | Pointer to **bool** |  | [optional] [default to false]
**Multiple** | Pointer to **bool** |  | [optional] [default to false]
**Name** | Pointer to **string** |  | [optional] 
**Observed** | Pointer to **bool** |  | [optional] [default to false]
**Operations** | Pointer to **string** |  | [optional] 
**RangeEnumeration** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Units** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 
**Values** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResourceInstance

`func NewResourceInstance() *ResourceInstance`

NewResourceInstance instantiates a new ResourceInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceInstanceWithDefaults

`func NewResourceInstanceWithDefaults() *ResourceInstance`

NewResourceInstanceWithDefaults instantiates a new ResourceInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ResourceInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ResourceInstance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceInstance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceInstance) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMandatory

`func (o *ResourceInstance) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *ResourceInstance) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *ResourceInstance) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *ResourceInstance) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.

### GetMultiple

`func (o *ResourceInstance) GetMultiple() bool`

GetMultiple returns the Multiple field if non-nil, zero value otherwise.

### GetMultipleOk

`func (o *ResourceInstance) GetMultipleOk() (*bool, bool)`

GetMultipleOk returns a tuple with the Multiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiple

`func (o *ResourceInstance) SetMultiple(v bool)`

SetMultiple sets Multiple field to given value.

### HasMultiple

`func (o *ResourceInstance) HasMultiple() bool`

HasMultiple returns a boolean if a field has been set.

### GetName

`func (o *ResourceInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObserved

`func (o *ResourceInstance) GetObserved() bool`

GetObserved returns the Observed field if non-nil, zero value otherwise.

### GetObservedOk

`func (o *ResourceInstance) GetObservedOk() (*bool, bool)`

GetObservedOk returns a tuple with the Observed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObserved

`func (o *ResourceInstance) SetObserved(v bool)`

SetObserved sets Observed field to given value.

### HasObserved

`func (o *ResourceInstance) HasObserved() bool`

HasObserved returns a boolean if a field has been set.

### GetOperations

`func (o *ResourceInstance) GetOperations() string`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *ResourceInstance) GetOperationsOk() (*string, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *ResourceInstance) SetOperations(v string)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *ResourceInstance) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetRangeEnumeration

`func (o *ResourceInstance) GetRangeEnumeration() string`

GetRangeEnumeration returns the RangeEnumeration field if non-nil, zero value otherwise.

### GetRangeEnumerationOk

`func (o *ResourceInstance) GetRangeEnumerationOk() (*string, bool)`

GetRangeEnumerationOk returns a tuple with the RangeEnumeration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeEnumeration

`func (o *ResourceInstance) SetRangeEnumeration(v string)`

SetRangeEnumeration sets RangeEnumeration field to given value.

### HasRangeEnumeration

`func (o *ResourceInstance) HasRangeEnumeration() bool`

HasRangeEnumeration returns a boolean if a field has been set.

### GetType

`func (o *ResourceInstance) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceInstance) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceInstance) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResourceInstance) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnits

`func (o *ResourceInstance) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ResourceInstance) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ResourceInstance) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *ResourceInstance) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetValue

`func (o *ResourceInstance) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResourceInstance) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResourceInstance) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ResourceInstance) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValues

`func (o *ResourceInstance) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ResourceInstance) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ResourceInstance) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *ResourceInstance) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


