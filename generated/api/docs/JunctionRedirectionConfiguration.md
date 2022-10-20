# JunctionRedirectionConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**Gateway** | Pointer to **string** |  | [optional] 

## Methods

### NewJunctionRedirectionConfiguration

`func NewJunctionRedirectionConfiguration() *JunctionRedirectionConfiguration`

NewJunctionRedirectionConfiguration instantiates a new JunctionRedirectionConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunctionRedirectionConfigurationWithDefaults

`func NewJunctionRedirectionConfigurationWithDefaults() *JunctionRedirectionConfiguration`

NewJunctionRedirectionConfigurationWithDefaults instantiates a new JunctionRedirectionConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *JunctionRedirectionConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JunctionRedirectionConfiguration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JunctionRedirectionConfiguration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JunctionRedirectionConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *JunctionRedirectionConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JunctionRedirectionConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JunctionRedirectionConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *JunctionRedirectionConfiguration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGateway

`func (o *JunctionRedirectionConfiguration) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *JunctionRedirectionConfiguration) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *JunctionRedirectionConfiguration) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *JunctionRedirectionConfiguration) HasGateway() bool`

HasGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


