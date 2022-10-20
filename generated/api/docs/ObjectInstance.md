# ObjectInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Observed** | Pointer to **bool** |  | [optional] [default to false]
**Resources** | Pointer to [**map[string]ResourceInstance**](ResourceInstance.md) |  | [optional] 

## Methods

### NewObjectInstance

`func NewObjectInstance() *ObjectInstance`

NewObjectInstance instantiates a new ObjectInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectInstanceWithDefaults

`func NewObjectInstanceWithDefaults() *ObjectInstance`

NewObjectInstanceWithDefaults instantiates a new ObjectInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectInstance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectInstance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectInstance) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObserved

`func (o *ObjectInstance) GetObserved() bool`

GetObserved returns the Observed field if non-nil, zero value otherwise.

### GetObservedOk

`func (o *ObjectInstance) GetObservedOk() (*bool, bool)`

GetObservedOk returns a tuple with the Observed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObserved

`func (o *ObjectInstance) SetObserved(v bool)`

SetObserved sets Observed field to given value.

### HasObserved

`func (o *ObjectInstance) HasObserved() bool`

HasObserved returns a boolean if a field has been set.

### GetResources

`func (o *ObjectInstance) GetResources() map[string]ResourceInstance`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ObjectInstance) GetResourcesOk() (*map[string]ResourceInstance, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ObjectInstance) SetResources(v map[string]ResourceInstance)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ObjectInstance) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


