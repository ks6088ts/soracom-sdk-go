# Gadget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastModifiedTime** | Pointer to **time.Time** |  | [optional] 
**LastSeen** | Pointer to **map[string]interface{}** |  | [optional] 
**OperatorId** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**TerminationEnabled** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewGadget

`func NewGadget() *Gadget`

NewGadget instantiates a new Gadget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGadgetWithDefaults

`func NewGadgetWithDefaults() *Gadget`

NewGadgetWithDefaults instantiates a new Gadget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *Gadget) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Gadget) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Gadget) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Gadget) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetGroupId

`func (o *Gadget) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Gadget) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Gadget) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Gadget) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *Gadget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Gadget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Gadget) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Gadget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *Gadget) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *Gadget) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *Gadget) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *Gadget) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetLastSeen

`func (o *Gadget) GetLastSeen() map[string]interface{}`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *Gadget) GetLastSeenOk() (*map[string]interface{}, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *Gadget) SetLastSeen(v map[string]interface{})`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *Gadget) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetOperatorId

`func (o *Gadget) GetOperatorId() string`

GetOperatorId returns the OperatorId field if non-nil, zero value otherwise.

### GetOperatorIdOk

`func (o *Gadget) GetOperatorIdOk() (*string, bool)`

GetOperatorIdOk returns a tuple with the OperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorId

`func (o *Gadget) SetOperatorId(v string)`

SetOperatorId sets OperatorId field to given value.

### HasOperatorId

`func (o *Gadget) HasOperatorId() bool`

HasOperatorId returns a boolean if a field has been set.

### GetProductId

`func (o *Gadget) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Gadget) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Gadget) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *Gadget) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *Gadget) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Gadget) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Gadget) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Gadget) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetStatus

`func (o *Gadget) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Gadget) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Gadget) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Gadget) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *Gadget) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Gadget) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Gadget) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Gadget) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTerminationEnabled

`func (o *Gadget) GetTerminationEnabled() bool`

GetTerminationEnabled returns the TerminationEnabled field if non-nil, zero value otherwise.

### GetTerminationEnabledOk

`func (o *Gadget) GetTerminationEnabledOk() (*bool, bool)`

GetTerminationEnabledOk returns a tuple with the TerminationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationEnabled

`func (o *Gadget) SetTerminationEnabled(v bool)`

SetTerminationEnabled sets TerminationEnabled field to given value.

### HasTerminationEnabled

`func (o *Gadget) HasTerminationEnabled() bool`

HasTerminationEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


