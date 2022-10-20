# CreatePortMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | [**PortMappingDestination**](PortMappingDestination.md) |  | 
**Duration** | Pointer to **float32** | Specify the duration, in seconds, that remote access should be enabled. | [optional] 
**Source** | Pointer to [**PortMappingSource**](PortMappingSource.md) |  | [optional] 
**TlsRequired** | Pointer to **bool** | Specify whether access uses TLS. | [optional] 

## Methods

### NewCreatePortMappingRequest

`func NewCreatePortMappingRequest(destination PortMappingDestination, ) *CreatePortMappingRequest`

NewCreatePortMappingRequest instantiates a new CreatePortMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePortMappingRequestWithDefaults

`func NewCreatePortMappingRequestWithDefaults() *CreatePortMappingRequest`

NewCreatePortMappingRequestWithDefaults instantiates a new CreatePortMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *CreatePortMappingRequest) GetDestination() PortMappingDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *CreatePortMappingRequest) GetDestinationOk() (*PortMappingDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *CreatePortMappingRequest) SetDestination(v PortMappingDestination)`

SetDestination sets Destination field to given value.


### GetDuration

`func (o *CreatePortMappingRequest) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CreatePortMappingRequest) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CreatePortMappingRequest) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CreatePortMappingRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetSource

`func (o *CreatePortMappingRequest) GetSource() PortMappingSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreatePortMappingRequest) GetSourceOk() (*PortMappingSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreatePortMappingRequest) SetSource(v PortMappingSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *CreatePortMappingRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTlsRequired

`func (o *CreatePortMappingRequest) GetTlsRequired() bool`

GetTlsRequired returns the TlsRequired field if non-nil, zero value otherwise.

### GetTlsRequiredOk

`func (o *CreatePortMappingRequest) GetTlsRequiredOk() (*bool, bool)`

GetTlsRequiredOk returns a tuple with the TlsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsRequired

`func (o *CreatePortMappingRequest) SetTlsRequired(v bool)`

SetTlsRequired sets TlsRequired field to given value.

### HasTlsRequired

`func (o *CreatePortMappingRequest) HasTlsRequired() bool`

HasTlsRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


