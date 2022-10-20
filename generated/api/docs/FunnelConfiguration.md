# FunnelConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddSimId** | Pointer to **bool** |  | [optional] [default to false]
**ContentType** | Pointer to [**FunnelContentType**](FunnelContentType.md) |  | [optional] 
**CredentialsId** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to [**FunnelConfigurationDestination**](FunnelConfigurationDestination.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewFunnelConfiguration

`func NewFunnelConfiguration() *FunnelConfiguration`

NewFunnelConfiguration instantiates a new FunnelConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunnelConfigurationWithDefaults

`func NewFunnelConfigurationWithDefaults() *FunnelConfiguration`

NewFunnelConfigurationWithDefaults instantiates a new FunnelConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddSimId

`func (o *FunnelConfiguration) GetAddSimId() bool`

GetAddSimId returns the AddSimId field if non-nil, zero value otherwise.

### GetAddSimIdOk

`func (o *FunnelConfiguration) GetAddSimIdOk() (*bool, bool)`

GetAddSimIdOk returns a tuple with the AddSimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddSimId

`func (o *FunnelConfiguration) SetAddSimId(v bool)`

SetAddSimId sets AddSimId field to given value.

### HasAddSimId

`func (o *FunnelConfiguration) HasAddSimId() bool`

HasAddSimId returns a boolean if a field has been set.

### GetContentType

`func (o *FunnelConfiguration) GetContentType() FunnelContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *FunnelConfiguration) GetContentTypeOk() (*FunnelContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *FunnelConfiguration) SetContentType(v FunnelContentType)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *FunnelConfiguration) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetCredentialsId

`func (o *FunnelConfiguration) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *FunnelConfiguration) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *FunnelConfiguration) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.

### HasCredentialsId

`func (o *FunnelConfiguration) HasCredentialsId() bool`

HasCredentialsId returns a boolean if a field has been set.

### GetDestination

`func (o *FunnelConfiguration) GetDestination() FunnelConfigurationDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *FunnelConfiguration) GetDestinationOk() (*FunnelConfigurationDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *FunnelConfiguration) SetDestination(v FunnelConfigurationDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *FunnelConfiguration) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetEnabled

`func (o *FunnelConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FunnelConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FunnelConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FunnelConfiguration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


