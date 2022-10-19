# PreviousSessionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cell** | Pointer to [**Cell**](Cell.md) |  | [optional] 
**CreatedTime** | Pointer to **int64** | The time when the previous session was created. | [optional] 
**DeletedTime** | Pointer to **int64** | The time when the previous session was deleted. | [optional] 
**DnsServers** | Pointer to **[]string** | Array of IP address of the DNS servers. | [optional] 
**GatewayPrivateIpAddress** | Pointer to **string** |  | [optional] 
**GatewayPublicIpAddress** | Pointer to **string** |  | [optional] 
**Imei** | Pointer to **string** | The IMEI of the device using the SIM. | [optional] 
**UeIpAddress** | Pointer to **string** | The IP address of the device. | [optional] 
**VpgId** | Pointer to **string** | The Virtual Private Gateway IP address configured. | [optional] 

## Methods

### NewPreviousSessionStatus

`func NewPreviousSessionStatus() *PreviousSessionStatus`

NewPreviousSessionStatus instantiates a new PreviousSessionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviousSessionStatusWithDefaults

`func NewPreviousSessionStatusWithDefaults() *PreviousSessionStatus`

NewPreviousSessionStatusWithDefaults instantiates a new PreviousSessionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCell

`func (o *PreviousSessionStatus) GetCell() Cell`

GetCell returns the Cell field if non-nil, zero value otherwise.

### GetCellOk

`func (o *PreviousSessionStatus) GetCellOk() (*Cell, bool)`

GetCellOk returns a tuple with the Cell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCell

`func (o *PreviousSessionStatus) SetCell(v Cell)`

SetCell sets Cell field to given value.

### HasCell

`func (o *PreviousSessionStatus) HasCell() bool`

HasCell returns a boolean if a field has been set.

### GetCreatedTime

`func (o *PreviousSessionStatus) GetCreatedTime() int64`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PreviousSessionStatus) GetCreatedTimeOk() (*int64, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PreviousSessionStatus) SetCreatedTime(v int64)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *PreviousSessionStatus) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDeletedTime

`func (o *PreviousSessionStatus) GetDeletedTime() int64`

GetDeletedTime returns the DeletedTime field if non-nil, zero value otherwise.

### GetDeletedTimeOk

`func (o *PreviousSessionStatus) GetDeletedTimeOk() (*int64, bool)`

GetDeletedTimeOk returns a tuple with the DeletedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedTime

`func (o *PreviousSessionStatus) SetDeletedTime(v int64)`

SetDeletedTime sets DeletedTime field to given value.

### HasDeletedTime

`func (o *PreviousSessionStatus) HasDeletedTime() bool`

HasDeletedTime returns a boolean if a field has been set.

### GetDnsServers

`func (o *PreviousSessionStatus) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *PreviousSessionStatus) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *PreviousSessionStatus) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *PreviousSessionStatus) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetGatewayPrivateIpAddress

`func (o *PreviousSessionStatus) GetGatewayPrivateIpAddress() string`

GetGatewayPrivateIpAddress returns the GatewayPrivateIpAddress field if non-nil, zero value otherwise.

### GetGatewayPrivateIpAddressOk

`func (o *PreviousSessionStatus) GetGatewayPrivateIpAddressOk() (*string, bool)`

GetGatewayPrivateIpAddressOk returns a tuple with the GatewayPrivateIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPrivateIpAddress

`func (o *PreviousSessionStatus) SetGatewayPrivateIpAddress(v string)`

SetGatewayPrivateIpAddress sets GatewayPrivateIpAddress field to given value.

### HasGatewayPrivateIpAddress

`func (o *PreviousSessionStatus) HasGatewayPrivateIpAddress() bool`

HasGatewayPrivateIpAddress returns a boolean if a field has been set.

### GetGatewayPublicIpAddress

`func (o *PreviousSessionStatus) GetGatewayPublicIpAddress() string`

GetGatewayPublicIpAddress returns the GatewayPublicIpAddress field if non-nil, zero value otherwise.

### GetGatewayPublicIpAddressOk

`func (o *PreviousSessionStatus) GetGatewayPublicIpAddressOk() (*string, bool)`

GetGatewayPublicIpAddressOk returns a tuple with the GatewayPublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPublicIpAddress

`func (o *PreviousSessionStatus) SetGatewayPublicIpAddress(v string)`

SetGatewayPublicIpAddress sets GatewayPublicIpAddress field to given value.

### HasGatewayPublicIpAddress

`func (o *PreviousSessionStatus) HasGatewayPublicIpAddress() bool`

HasGatewayPublicIpAddress returns a boolean if a field has been set.

### GetImei

`func (o *PreviousSessionStatus) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *PreviousSessionStatus) GetImeiOk() (*string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImei

`func (o *PreviousSessionStatus) SetImei(v string)`

SetImei sets Imei field to given value.

### HasImei

`func (o *PreviousSessionStatus) HasImei() bool`

HasImei returns a boolean if a field has been set.

### GetUeIpAddress

`func (o *PreviousSessionStatus) GetUeIpAddress() string`

GetUeIpAddress returns the UeIpAddress field if non-nil, zero value otherwise.

### GetUeIpAddressOk

`func (o *PreviousSessionStatus) GetUeIpAddressOk() (*string, bool)`

GetUeIpAddressOk returns a tuple with the UeIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpAddress

`func (o *PreviousSessionStatus) SetUeIpAddress(v string)`

SetUeIpAddress sets UeIpAddress field to given value.

### HasUeIpAddress

`func (o *PreviousSessionStatus) HasUeIpAddress() bool`

HasUeIpAddress returns a boolean if a field has been set.

### GetVpgId

`func (o *PreviousSessionStatus) GetVpgId() string`

GetVpgId returns the VpgId field if non-nil, zero value otherwise.

### GetVpgIdOk

`func (o *PreviousSessionStatus) GetVpgIdOk() (*string, bool)`

GetVpgIdOk returns a tuple with the VpgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpgId

`func (o *PreviousSessionStatus) SetVpgId(v string)`

SetVpgId sets VpgId field to given value.

### HasVpgId

`func (o *PreviousSessionStatus) HasVpgId() bool`

HasVpgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


