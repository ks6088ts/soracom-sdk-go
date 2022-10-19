# Placement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfrastructureProvider** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** | &#x60;infrastructureProvider&#x60; が &#x60;aws&#x60; の場合は、ランデブーポイントです。  - &#x60;ap-northeast-1&#x60;: 東京（日本） - &#x60;eu-central-1&#x60;: フランクフルト (ドイツ) - &#x60;us-west-2&#x60;: オレゴン (米国) - &#x60;ap-southeast-2&#x60;: シドニー (オーストラリア)  | [optional] [default to "eu-central-1"]

## Methods

### NewPlacement

`func NewPlacement() *Placement`

NewPlacement instantiates a new Placement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlacementWithDefaults

`func NewPlacementWithDefaults() *Placement`

NewPlacementWithDefaults instantiates a new Placement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfrastructureProvider

`func (o *Placement) GetInfrastructureProvider() string`

GetInfrastructureProvider returns the InfrastructureProvider field if non-nil, zero value otherwise.

### GetInfrastructureProviderOk

`func (o *Placement) GetInfrastructureProviderOk() (*string, bool)`

GetInfrastructureProviderOk returns a tuple with the InfrastructureProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureProvider

`func (o *Placement) SetInfrastructureProvider(v string)`

SetInfrastructureProvider sets InfrastructureProvider field to given value.

### HasInfrastructureProvider

`func (o *Placement) HasInfrastructureProvider() bool`

HasInfrastructureProvider returns a boolean if a field has been set.

### GetRegion

`func (o *Placement) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Placement) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Placement) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Placement) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


