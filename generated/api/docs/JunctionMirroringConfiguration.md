# JunctionMirroringConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Peers** | Pointer to [**map[string]JunctionMirroringPeer**](JunctionMirroringPeer.md) |  | [optional] 

## Methods

### NewJunctionMirroringConfiguration

`func NewJunctionMirroringConfiguration() *JunctionMirroringConfiguration`

NewJunctionMirroringConfiguration instantiates a new JunctionMirroringConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunctionMirroringConfigurationWithDefaults

`func NewJunctionMirroringConfigurationWithDefaults() *JunctionMirroringConfiguration`

NewJunctionMirroringConfigurationWithDefaults instantiates a new JunctionMirroringConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeers

`func (o *JunctionMirroringConfiguration) GetPeers() map[string]JunctionMirroringPeer`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *JunctionMirroringConfiguration) GetPeersOk() (*map[string]JunctionMirroringPeer, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *JunctionMirroringConfiguration) SetPeers(v map[string]JunctionMirroringPeer)`

SetPeers sets Peers field to given value.

### HasPeers

`func (o *JunctionMirroringConfiguration) HasPeers() bool`

HasPeers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


