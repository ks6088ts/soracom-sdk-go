# SubscriptionContainerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | Pointer to [**[]SubscriptionContainer**](SubscriptionContainer.md) |  | [optional] 
**CountryMapping** | Pointer to [**map[string]SubscriptionContainerStatusCountryMappingValue**](SubscriptionContainerStatusCountryMappingValue.md) |  | [optional] 

## Methods

### NewSubscriptionContainerStatus

`func NewSubscriptionContainerStatus() *SubscriptionContainerStatus`

NewSubscriptionContainerStatus instantiates a new SubscriptionContainerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionContainerStatusWithDefaults

`func NewSubscriptionContainerStatusWithDefaults() *SubscriptionContainerStatus`

NewSubscriptionContainerStatusWithDefaults instantiates a new SubscriptionContainerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *SubscriptionContainerStatus) GetContainers() []SubscriptionContainer`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *SubscriptionContainerStatus) GetContainersOk() (*[]SubscriptionContainer, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *SubscriptionContainerStatus) SetContainers(v []SubscriptionContainer)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *SubscriptionContainerStatus) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetCountryMapping

`func (o *SubscriptionContainerStatus) GetCountryMapping() map[string]SubscriptionContainerStatusCountryMappingValue`

GetCountryMapping returns the CountryMapping field if non-nil, zero value otherwise.

### GetCountryMappingOk

`func (o *SubscriptionContainerStatus) GetCountryMappingOk() (*map[string]SubscriptionContainerStatusCountryMappingValue, bool)`

GetCountryMappingOk returns a tuple with the CountryMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryMapping

`func (o *SubscriptionContainerStatus) SetCountryMapping(v map[string]SubscriptionContainerStatusCountryMappingValue)`

SetCountryMapping sets CountryMapping field to given value.

### HasCountryMapping

`func (o *SubscriptionContainerStatus) HasCountryMapping() bool`

HasCountryMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


