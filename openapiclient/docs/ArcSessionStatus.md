# ArcSessionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArcAllowedIPs** | Pointer to **[]string** |  | [optional] 
**ArcClientPeerIpAddress** | Pointer to **string** |  | [optional] 
**ArcClientPeerPublicKey** | Pointer to **string** |  | [optional] 
**ArcServerEndpoint** | Pointer to **string** |  | [optional] 
**ArcServerPeerPublicKey** | Pointer to **string** |  | [optional] 
**GatewayPrivateIpAddress** | Pointer to **string** |  | [optional] 
**GatewayPublicIpAddress** | Pointer to **string** |  | [optional] 
**LastUpdatedAt** | Pointer to **int64** |  | [optional] 
**VpgId** | Pointer to **string** |  | [optional] 

## Methods

### NewArcSessionStatus

`func NewArcSessionStatus() *ArcSessionStatus`

NewArcSessionStatus instantiates a new ArcSessionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcSessionStatusWithDefaults

`func NewArcSessionStatusWithDefaults() *ArcSessionStatus`

NewArcSessionStatusWithDefaults instantiates a new ArcSessionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArcAllowedIPs

`func (o *ArcSessionStatus) GetArcAllowedIPs() []string`

GetArcAllowedIPs returns the ArcAllowedIPs field if non-nil, zero value otherwise.

### GetArcAllowedIPsOk

`func (o *ArcSessionStatus) GetArcAllowedIPsOk() (*[]string, bool)`

GetArcAllowedIPsOk returns a tuple with the ArcAllowedIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcAllowedIPs

`func (o *ArcSessionStatus) SetArcAllowedIPs(v []string)`

SetArcAllowedIPs sets ArcAllowedIPs field to given value.

### HasArcAllowedIPs

`func (o *ArcSessionStatus) HasArcAllowedIPs() bool`

HasArcAllowedIPs returns a boolean if a field has been set.

### GetArcClientPeerIpAddress

`func (o *ArcSessionStatus) GetArcClientPeerIpAddress() string`

GetArcClientPeerIpAddress returns the ArcClientPeerIpAddress field if non-nil, zero value otherwise.

### GetArcClientPeerIpAddressOk

`func (o *ArcSessionStatus) GetArcClientPeerIpAddressOk() (*string, bool)`

GetArcClientPeerIpAddressOk returns a tuple with the ArcClientPeerIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcClientPeerIpAddress

`func (o *ArcSessionStatus) SetArcClientPeerIpAddress(v string)`

SetArcClientPeerIpAddress sets ArcClientPeerIpAddress field to given value.

### HasArcClientPeerIpAddress

`func (o *ArcSessionStatus) HasArcClientPeerIpAddress() bool`

HasArcClientPeerIpAddress returns a boolean if a field has been set.

### GetArcClientPeerPublicKey

`func (o *ArcSessionStatus) GetArcClientPeerPublicKey() string`

GetArcClientPeerPublicKey returns the ArcClientPeerPublicKey field if non-nil, zero value otherwise.

### GetArcClientPeerPublicKeyOk

`func (o *ArcSessionStatus) GetArcClientPeerPublicKeyOk() (*string, bool)`

GetArcClientPeerPublicKeyOk returns a tuple with the ArcClientPeerPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcClientPeerPublicKey

`func (o *ArcSessionStatus) SetArcClientPeerPublicKey(v string)`

SetArcClientPeerPublicKey sets ArcClientPeerPublicKey field to given value.

### HasArcClientPeerPublicKey

`func (o *ArcSessionStatus) HasArcClientPeerPublicKey() bool`

HasArcClientPeerPublicKey returns a boolean if a field has been set.

### GetArcServerEndpoint

`func (o *ArcSessionStatus) GetArcServerEndpoint() string`

GetArcServerEndpoint returns the ArcServerEndpoint field if non-nil, zero value otherwise.

### GetArcServerEndpointOk

`func (o *ArcSessionStatus) GetArcServerEndpointOk() (*string, bool)`

GetArcServerEndpointOk returns a tuple with the ArcServerEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcServerEndpoint

`func (o *ArcSessionStatus) SetArcServerEndpoint(v string)`

SetArcServerEndpoint sets ArcServerEndpoint field to given value.

### HasArcServerEndpoint

`func (o *ArcSessionStatus) HasArcServerEndpoint() bool`

HasArcServerEndpoint returns a boolean if a field has been set.

### GetArcServerPeerPublicKey

`func (o *ArcSessionStatus) GetArcServerPeerPublicKey() string`

GetArcServerPeerPublicKey returns the ArcServerPeerPublicKey field if non-nil, zero value otherwise.

### GetArcServerPeerPublicKeyOk

`func (o *ArcSessionStatus) GetArcServerPeerPublicKeyOk() (*string, bool)`

GetArcServerPeerPublicKeyOk returns a tuple with the ArcServerPeerPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcServerPeerPublicKey

`func (o *ArcSessionStatus) SetArcServerPeerPublicKey(v string)`

SetArcServerPeerPublicKey sets ArcServerPeerPublicKey field to given value.

### HasArcServerPeerPublicKey

`func (o *ArcSessionStatus) HasArcServerPeerPublicKey() bool`

HasArcServerPeerPublicKey returns a boolean if a field has been set.

### GetGatewayPrivateIpAddress

`func (o *ArcSessionStatus) GetGatewayPrivateIpAddress() string`

GetGatewayPrivateIpAddress returns the GatewayPrivateIpAddress field if non-nil, zero value otherwise.

### GetGatewayPrivateIpAddressOk

`func (o *ArcSessionStatus) GetGatewayPrivateIpAddressOk() (*string, bool)`

GetGatewayPrivateIpAddressOk returns a tuple with the GatewayPrivateIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPrivateIpAddress

`func (o *ArcSessionStatus) SetGatewayPrivateIpAddress(v string)`

SetGatewayPrivateIpAddress sets GatewayPrivateIpAddress field to given value.

### HasGatewayPrivateIpAddress

`func (o *ArcSessionStatus) HasGatewayPrivateIpAddress() bool`

HasGatewayPrivateIpAddress returns a boolean if a field has been set.

### GetGatewayPublicIpAddress

`func (o *ArcSessionStatus) GetGatewayPublicIpAddress() string`

GetGatewayPublicIpAddress returns the GatewayPublicIpAddress field if non-nil, zero value otherwise.

### GetGatewayPublicIpAddressOk

`func (o *ArcSessionStatus) GetGatewayPublicIpAddressOk() (*string, bool)`

GetGatewayPublicIpAddressOk returns a tuple with the GatewayPublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPublicIpAddress

`func (o *ArcSessionStatus) SetGatewayPublicIpAddress(v string)`

SetGatewayPublicIpAddress sets GatewayPublicIpAddress field to given value.

### HasGatewayPublicIpAddress

`func (o *ArcSessionStatus) HasGatewayPublicIpAddress() bool`

HasGatewayPublicIpAddress returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *ArcSessionStatus) GetLastUpdatedAt() int64`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *ArcSessionStatus) GetLastUpdatedAtOk() (*int64, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *ArcSessionStatus) SetLastUpdatedAt(v int64)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *ArcSessionStatus) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetVpgId

`func (o *ArcSessionStatus) GetVpgId() string`

GetVpgId returns the VpgId field if non-nil, zero value otherwise.

### GetVpgIdOk

`func (o *ArcSessionStatus) GetVpgIdOk() (*string, bool)`

GetVpgIdOk returns a tuple with the VpgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpgId

`func (o *ArcSessionStatus) SetVpgId(v string)`

SetVpgId sets VpgId field to given value.

### HasVpgId

`func (o *ArcSessionStatus) HasVpgId() bool`

HasVpgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


