# VirtualPrivateGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedOperators** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedTime** | Pointer to **int64** |  | [optional] 
**DeviceSubnetCidrRange** | Pointer to **string** | デバイスサブネット IP アドレスレンジ | [optional] 
**FixedIpAddressesEnabled** | Pointer to **bool** | 固定グローバル IP アドレスオプション  - &#x60;true&#x60;: 有効 - &#x60;false&#x60;: 無効  | [optional] 
**GateVxlanId** | Pointer to **float32** |  | [optional] 
**ImplicitTerminationProtected** | Pointer to **bool** |  | [optional] 
**JunctionEnabled** | Pointer to **bool** |  | [optional] 
**LastModifiedTime** | Pointer to **int64** |  | [optional] 
**OperatorId** | Pointer to **string** |  | [optional] 
**Placement** | Pointer to [**Placement**](Placement.md) |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** | An object which always contains at least one property \&quot;name\&quot; with a string value. If you give a subscriber/SIM a name, the name will be returned as the value of the \&quot;name\&quot; property. If the subscriber/SIM does not have a name, an empty string \&quot;\&quot; is returned. In addition, if you create any custom tags for the subscriber/SIM, each custom tag will appear as additional properties in the object. | [optional] 
**Type** | Pointer to **int32** | VPG のタイプ  - &#x60;14&#x60; : Type-E - &#x60;15&#x60; : Type-F  | [optional] 
**UeSubnetCidrRange** | Pointer to **string** | デバイスサブネット IP アドレスレンジ | [optional] 
**UseInternetGateway** | Pointer to **bool** | インターネットゲートウェイ  - &#x60;true&#x60;: 有効 - &#x60;false&#x60;: 無効  | [optional] 
**VirtualInterfaces** | Pointer to **map[string]string** |  | [optional] 
**VpcPeeringConnections** | Pointer to **map[string]string** |  | [optional] 
**VpgId** | Pointer to **string** | ID | [optional] 

## Methods

### NewVirtualPrivateGateway

`func NewVirtualPrivateGateway() *VirtualPrivateGateway`

NewVirtualPrivateGateway instantiates a new VirtualPrivateGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualPrivateGatewayWithDefaults

`func NewVirtualPrivateGatewayWithDefaults() *VirtualPrivateGateway`

NewVirtualPrivateGatewayWithDefaults instantiates a new VirtualPrivateGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedOperators

`func (o *VirtualPrivateGateway) GetAllowedOperators() map[string]interface{}`

GetAllowedOperators returns the AllowedOperators field if non-nil, zero value otherwise.

### GetAllowedOperatorsOk

`func (o *VirtualPrivateGateway) GetAllowedOperatorsOk() (*map[string]interface{}, bool)`

GetAllowedOperatorsOk returns a tuple with the AllowedOperators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOperators

`func (o *VirtualPrivateGateway) SetAllowedOperators(v map[string]interface{})`

SetAllowedOperators sets AllowedOperators field to given value.

### HasAllowedOperators

`func (o *VirtualPrivateGateway) HasAllowedOperators() bool`

HasAllowedOperators returns a boolean if a field has been set.

### GetCreatedTime

`func (o *VirtualPrivateGateway) GetCreatedTime() int64`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *VirtualPrivateGateway) GetCreatedTimeOk() (*int64, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *VirtualPrivateGateway) SetCreatedTime(v int64)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *VirtualPrivateGateway) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDeviceSubnetCidrRange

`func (o *VirtualPrivateGateway) GetDeviceSubnetCidrRange() string`

GetDeviceSubnetCidrRange returns the DeviceSubnetCidrRange field if non-nil, zero value otherwise.

### GetDeviceSubnetCidrRangeOk

`func (o *VirtualPrivateGateway) GetDeviceSubnetCidrRangeOk() (*string, bool)`

GetDeviceSubnetCidrRangeOk returns a tuple with the DeviceSubnetCidrRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSubnetCidrRange

`func (o *VirtualPrivateGateway) SetDeviceSubnetCidrRange(v string)`

SetDeviceSubnetCidrRange sets DeviceSubnetCidrRange field to given value.

### HasDeviceSubnetCidrRange

`func (o *VirtualPrivateGateway) HasDeviceSubnetCidrRange() bool`

HasDeviceSubnetCidrRange returns a boolean if a field has been set.

### GetFixedIpAddressesEnabled

`func (o *VirtualPrivateGateway) GetFixedIpAddressesEnabled() bool`

GetFixedIpAddressesEnabled returns the FixedIpAddressesEnabled field if non-nil, zero value otherwise.

### GetFixedIpAddressesEnabledOk

`func (o *VirtualPrivateGateway) GetFixedIpAddressesEnabledOk() (*bool, bool)`

GetFixedIpAddressesEnabledOk returns a tuple with the FixedIpAddressesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAddressesEnabled

`func (o *VirtualPrivateGateway) SetFixedIpAddressesEnabled(v bool)`

SetFixedIpAddressesEnabled sets FixedIpAddressesEnabled field to given value.

### HasFixedIpAddressesEnabled

`func (o *VirtualPrivateGateway) HasFixedIpAddressesEnabled() bool`

HasFixedIpAddressesEnabled returns a boolean if a field has been set.

### GetGateVxlanId

`func (o *VirtualPrivateGateway) GetGateVxlanId() float32`

GetGateVxlanId returns the GateVxlanId field if non-nil, zero value otherwise.

### GetGateVxlanIdOk

`func (o *VirtualPrivateGateway) GetGateVxlanIdOk() (*float32, bool)`

GetGateVxlanIdOk returns a tuple with the GateVxlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateVxlanId

`func (o *VirtualPrivateGateway) SetGateVxlanId(v float32)`

SetGateVxlanId sets GateVxlanId field to given value.

### HasGateVxlanId

`func (o *VirtualPrivateGateway) HasGateVxlanId() bool`

HasGateVxlanId returns a boolean if a field has been set.

### GetImplicitTerminationProtected

`func (o *VirtualPrivateGateway) GetImplicitTerminationProtected() bool`

GetImplicitTerminationProtected returns the ImplicitTerminationProtected field if non-nil, zero value otherwise.

### GetImplicitTerminationProtectedOk

`func (o *VirtualPrivateGateway) GetImplicitTerminationProtectedOk() (*bool, bool)`

GetImplicitTerminationProtectedOk returns a tuple with the ImplicitTerminationProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicitTerminationProtected

`func (o *VirtualPrivateGateway) SetImplicitTerminationProtected(v bool)`

SetImplicitTerminationProtected sets ImplicitTerminationProtected field to given value.

### HasImplicitTerminationProtected

`func (o *VirtualPrivateGateway) HasImplicitTerminationProtected() bool`

HasImplicitTerminationProtected returns a boolean if a field has been set.

### GetJunctionEnabled

`func (o *VirtualPrivateGateway) GetJunctionEnabled() bool`

GetJunctionEnabled returns the JunctionEnabled field if non-nil, zero value otherwise.

### GetJunctionEnabledOk

`func (o *VirtualPrivateGateway) GetJunctionEnabledOk() (*bool, bool)`

GetJunctionEnabledOk returns a tuple with the JunctionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJunctionEnabled

`func (o *VirtualPrivateGateway) SetJunctionEnabled(v bool)`

SetJunctionEnabled sets JunctionEnabled field to given value.

### HasJunctionEnabled

`func (o *VirtualPrivateGateway) HasJunctionEnabled() bool`

HasJunctionEnabled returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *VirtualPrivateGateway) GetLastModifiedTime() int64`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *VirtualPrivateGateway) GetLastModifiedTimeOk() (*int64, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *VirtualPrivateGateway) SetLastModifiedTime(v int64)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *VirtualPrivateGateway) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetOperatorId

`func (o *VirtualPrivateGateway) GetOperatorId() string`

GetOperatorId returns the OperatorId field if non-nil, zero value otherwise.

### GetOperatorIdOk

`func (o *VirtualPrivateGateway) GetOperatorIdOk() (*string, bool)`

GetOperatorIdOk returns a tuple with the OperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorId

`func (o *VirtualPrivateGateway) SetOperatorId(v string)`

SetOperatorId sets OperatorId field to given value.

### HasOperatorId

`func (o *VirtualPrivateGateway) HasOperatorId() bool`

HasOperatorId returns a boolean if a field has been set.

### GetPlacement

`func (o *VirtualPrivateGateway) GetPlacement() Placement`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VirtualPrivateGateway) GetPlacementOk() (*Placement, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VirtualPrivateGateway) SetPlacement(v Placement)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VirtualPrivateGateway) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetSize

`func (o *VirtualPrivateGateway) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VirtualPrivateGateway) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VirtualPrivateGateway) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *VirtualPrivateGateway) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualPrivateGateway) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualPrivateGateway) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualPrivateGateway) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualPrivateGateway) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *VirtualPrivateGateway) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualPrivateGateway) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualPrivateGateway) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualPrivateGateway) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *VirtualPrivateGateway) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualPrivateGateway) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualPrivateGateway) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *VirtualPrivateGateway) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUeSubnetCidrRange

`func (o *VirtualPrivateGateway) GetUeSubnetCidrRange() string`

GetUeSubnetCidrRange returns the UeSubnetCidrRange field if non-nil, zero value otherwise.

### GetUeSubnetCidrRangeOk

`func (o *VirtualPrivateGateway) GetUeSubnetCidrRangeOk() (*string, bool)`

GetUeSubnetCidrRangeOk returns a tuple with the UeSubnetCidrRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeSubnetCidrRange

`func (o *VirtualPrivateGateway) SetUeSubnetCidrRange(v string)`

SetUeSubnetCidrRange sets UeSubnetCidrRange field to given value.

### HasUeSubnetCidrRange

`func (o *VirtualPrivateGateway) HasUeSubnetCidrRange() bool`

HasUeSubnetCidrRange returns a boolean if a field has been set.

### GetUseInternetGateway

`func (o *VirtualPrivateGateway) GetUseInternetGateway() bool`

GetUseInternetGateway returns the UseInternetGateway field if non-nil, zero value otherwise.

### GetUseInternetGatewayOk

`func (o *VirtualPrivateGateway) GetUseInternetGatewayOk() (*bool, bool)`

GetUseInternetGatewayOk returns a tuple with the UseInternetGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseInternetGateway

`func (o *VirtualPrivateGateway) SetUseInternetGateway(v bool)`

SetUseInternetGateway sets UseInternetGateway field to given value.

### HasUseInternetGateway

`func (o *VirtualPrivateGateway) HasUseInternetGateway() bool`

HasUseInternetGateway returns a boolean if a field has been set.

### GetVirtualInterfaces

`func (o *VirtualPrivateGateway) GetVirtualInterfaces() map[string]string`

GetVirtualInterfaces returns the VirtualInterfaces field if non-nil, zero value otherwise.

### GetVirtualInterfacesOk

`func (o *VirtualPrivateGateway) GetVirtualInterfacesOk() (*map[string]string, bool)`

GetVirtualInterfacesOk returns a tuple with the VirtualInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualInterfaces

`func (o *VirtualPrivateGateway) SetVirtualInterfaces(v map[string]string)`

SetVirtualInterfaces sets VirtualInterfaces field to given value.

### HasVirtualInterfaces

`func (o *VirtualPrivateGateway) HasVirtualInterfaces() bool`

HasVirtualInterfaces returns a boolean if a field has been set.

### GetVpcPeeringConnections

`func (o *VirtualPrivateGateway) GetVpcPeeringConnections() map[string]string`

GetVpcPeeringConnections returns the VpcPeeringConnections field if non-nil, zero value otherwise.

### GetVpcPeeringConnectionsOk

`func (o *VirtualPrivateGateway) GetVpcPeeringConnectionsOk() (*map[string]string, bool)`

GetVpcPeeringConnectionsOk returns a tuple with the VpcPeeringConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcPeeringConnections

`func (o *VirtualPrivateGateway) SetVpcPeeringConnections(v map[string]string)`

SetVpcPeeringConnections sets VpcPeeringConnections field to given value.

### HasVpcPeeringConnections

`func (o *VirtualPrivateGateway) HasVpcPeeringConnections() bool`

HasVpcPeeringConnections returns a boolean if a field has been set.

### GetVpgId

`func (o *VirtualPrivateGateway) GetVpgId() string`

GetVpgId returns the VpgId field if non-nil, zero value otherwise.

### GetVpgIdOk

`func (o *VirtualPrivateGateway) GetVpgIdOk() (*string, bool)`

GetVpgIdOk returns a tuple with the VpgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpgId

`func (o *VirtualPrivateGateway) SetVpgId(v string)`

SetVpgId sets VpgId field to given value.

### HasVpgId

`func (o *VirtualPrivateGateway) HasVpgId() bool`

HasVpgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


