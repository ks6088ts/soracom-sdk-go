# FunnelConfigurationDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **string** |  | [optional] 
**ResourceUrl** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**RandomizePartitionKey** | Pointer to **bool** |  | [optional] 
**EndpointId** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **string** |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 
**CollectionId** | Pointer to **string** |  | [optional] 
**PassAll** | Pointer to **bool** |  | [optional] 
**Values** | Pointer to **string** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**DataFormat** | Pointer to **string** |  | [optional] 
**Desthost** | Pointer to **string** |  | [optional] 
**Destport** | Pointer to **int32** |  | [optional] 
**AdditionalData** | Pointer to **string** |  | [optional] 

## Methods

### NewFunnelConfigurationDestination

`func NewFunnelConfigurationDestination() *FunnelConfigurationDestination`

NewFunnelConfigurationDestination instantiates a new FunnelConfigurationDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunnelConfigurationDestinationWithDefaults

`func NewFunnelConfigurationDestinationWithDefaults() *FunnelConfigurationDestination`

NewFunnelConfigurationDestinationWithDefaults instantiates a new FunnelConfigurationDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *FunnelConfigurationDestination) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *FunnelConfigurationDestination) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *FunnelConfigurationDestination) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *FunnelConfigurationDestination) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetResourceUrl

`func (o *FunnelConfigurationDestination) GetResourceUrl() string`

GetResourceUrl returns the ResourceUrl field if non-nil, zero value otherwise.

### GetResourceUrlOk

`func (o *FunnelConfigurationDestination) GetResourceUrlOk() (*string, bool)`

GetResourceUrlOk returns a tuple with the ResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUrl

`func (o *FunnelConfigurationDestination) SetResourceUrl(v string)`

SetResourceUrl sets ResourceUrl field to given value.

### HasResourceUrl

`func (o *FunnelConfigurationDestination) HasResourceUrl() bool`

HasResourceUrl returns a boolean if a field has been set.

### GetService

`func (o *FunnelConfigurationDestination) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *FunnelConfigurationDestination) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *FunnelConfigurationDestination) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *FunnelConfigurationDestination) HasService() bool`

HasService returns a boolean if a field has been set.

### GetRandomizePartitionKey

`func (o *FunnelConfigurationDestination) GetRandomizePartitionKey() bool`

GetRandomizePartitionKey returns the RandomizePartitionKey field if non-nil, zero value otherwise.

### GetRandomizePartitionKeyOk

`func (o *FunnelConfigurationDestination) GetRandomizePartitionKeyOk() (*bool, bool)`

GetRandomizePartitionKeyOk returns a tuple with the RandomizePartitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizePartitionKey

`func (o *FunnelConfigurationDestination) SetRandomizePartitionKey(v bool)`

SetRandomizePartitionKey sets RandomizePartitionKey field to given value.

### HasRandomizePartitionKey

`func (o *FunnelConfigurationDestination) HasRandomizePartitionKey() bool`

HasRandomizePartitionKey returns a boolean if a field has been set.

### GetEndpointId

`func (o *FunnelConfigurationDestination) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *FunnelConfigurationDestination) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *FunnelConfigurationDestination) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *FunnelConfigurationDestination) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetTenantId

`func (o *FunnelConfigurationDestination) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *FunnelConfigurationDestination) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *FunnelConfigurationDestination) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *FunnelConfigurationDestination) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetConfig

`func (o *FunnelConfigurationDestination) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *FunnelConfigurationDestination) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *FunnelConfigurationDestination) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *FunnelConfigurationDestination) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetApplicationId

`func (o *FunnelConfigurationDestination) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *FunnelConfigurationDestination) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *FunnelConfigurationDestination) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *FunnelConfigurationDestination) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetCollectionId

`func (o *FunnelConfigurationDestination) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *FunnelConfigurationDestination) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *FunnelConfigurationDestination) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *FunnelConfigurationDestination) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetPassAll

`func (o *FunnelConfigurationDestination) GetPassAll() bool`

GetPassAll returns the PassAll field if non-nil, zero value otherwise.

### GetPassAllOk

`func (o *FunnelConfigurationDestination) GetPassAllOk() (*bool, bool)`

GetPassAllOk returns a tuple with the PassAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassAll

`func (o *FunnelConfigurationDestination) SetPassAll(v bool)`

SetPassAll sets PassAll field to given value.

### HasPassAll

`func (o *FunnelConfigurationDestination) HasPassAll() bool`

HasPassAll returns a boolean if a field has been set.

### GetValues

`func (o *FunnelConfigurationDestination) GetValues() string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *FunnelConfigurationDestination) GetValuesOk() (*string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *FunnelConfigurationDestination) SetValues(v string)`

SetValues sets Values field to given value.

### HasValues

`func (o *FunnelConfigurationDestination) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetChannelId

`func (o *FunnelConfigurationDestination) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *FunnelConfigurationDestination) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *FunnelConfigurationDestination) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *FunnelConfigurationDestination) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetDataFormat

`func (o *FunnelConfigurationDestination) GetDataFormat() string`

GetDataFormat returns the DataFormat field if non-nil, zero value otherwise.

### GetDataFormatOk

`func (o *FunnelConfigurationDestination) GetDataFormatOk() (*string, bool)`

GetDataFormatOk returns a tuple with the DataFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFormat

`func (o *FunnelConfigurationDestination) SetDataFormat(v string)`

SetDataFormat sets DataFormat field to given value.

### HasDataFormat

`func (o *FunnelConfigurationDestination) HasDataFormat() bool`

HasDataFormat returns a boolean if a field has been set.

### GetDesthost

`func (o *FunnelConfigurationDestination) GetDesthost() string`

GetDesthost returns the Desthost field if non-nil, zero value otherwise.

### GetDesthostOk

`func (o *FunnelConfigurationDestination) GetDesthostOk() (*string, bool)`

GetDesthostOk returns a tuple with the Desthost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesthost

`func (o *FunnelConfigurationDestination) SetDesthost(v string)`

SetDesthost sets Desthost field to given value.

### HasDesthost

`func (o *FunnelConfigurationDestination) HasDesthost() bool`

HasDesthost returns a boolean if a field has been set.

### GetDestport

`func (o *FunnelConfigurationDestination) GetDestport() int32`

GetDestport returns the Destport field if non-nil, zero value otherwise.

### GetDestportOk

`func (o *FunnelConfigurationDestination) GetDestportOk() (*int32, bool)`

GetDestportOk returns a tuple with the Destport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestport

`func (o *FunnelConfigurationDestination) SetDestport(v int32)`

SetDestport sets Destport field to given value.

### HasDestport

`func (o *FunnelConfigurationDestination) HasDestport() bool`

HasDestport returns a boolean if a field has been set.

### GetAdditionalData

`func (o *FunnelConfigurationDestination) GetAdditionalData() string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *FunnelConfigurationDestination) GetAdditionalDataOk() (*string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *FunnelConfigurationDestination) SetAdditionalData(v string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *FunnelConfigurationDestination) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


