# SessionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cell** | Pointer to [**Cell**](Cell.md) |  | [optional] 
**DnsServers** | Pointer to **[]string** | Array of IP address of the DNS servers. | [optional] 
**GatewayPrivateIpAddress** | Pointer to **string** |  | [optional] 
**GatewayPublicIpAddress** | Pointer to **string** |  | [optional] 
**Imei** | Pointer to **string** | The IMEI of the device using the SIM. | [optional] 
**Imsi** | Pointer to **string** | IMSI | [optional] 
**LastUpdatedAt** | Pointer to **int64** |  | [optional] 
**Location** | Pointer to **map[string]interface{}** |  | [optional] 
**Online** | Pointer to **bool** | セッション状態 - &#x60;true&#x60;: オンライン - &#x60;false&#x60;: オフライン  | [optional] 
**Placement** | Pointer to [**Placement**](Placement.md) |  | [optional] 
**Subscription** | Pointer to **string** | サブスクリプション | [optional] 
**UeIpAddress** | Pointer to **string** | The IP address of the device. | [optional] 
**VpgId** | Pointer to **string** | The Virtual Private Gateway IP address configured. | [optional] 

## Methods

### NewSessionStatus

`func NewSessionStatus() *SessionStatus`

NewSessionStatus instantiates a new SessionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionStatusWithDefaults

`func NewSessionStatusWithDefaults() *SessionStatus`

NewSessionStatusWithDefaults instantiates a new SessionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCell

`func (o *SessionStatus) GetCell() Cell`

GetCell returns the Cell field if non-nil, zero value otherwise.

### GetCellOk

`func (o *SessionStatus) GetCellOk() (*Cell, bool)`

GetCellOk returns a tuple with the Cell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCell

`func (o *SessionStatus) SetCell(v Cell)`

SetCell sets Cell field to given value.

### HasCell

`func (o *SessionStatus) HasCell() bool`

HasCell returns a boolean if a field has been set.

### GetDnsServers

`func (o *SessionStatus) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *SessionStatus) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *SessionStatus) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *SessionStatus) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetGatewayPrivateIpAddress

`func (o *SessionStatus) GetGatewayPrivateIpAddress() string`

GetGatewayPrivateIpAddress returns the GatewayPrivateIpAddress field if non-nil, zero value otherwise.

### GetGatewayPrivateIpAddressOk

`func (o *SessionStatus) GetGatewayPrivateIpAddressOk() (*string, bool)`

GetGatewayPrivateIpAddressOk returns a tuple with the GatewayPrivateIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPrivateIpAddress

`func (o *SessionStatus) SetGatewayPrivateIpAddress(v string)`

SetGatewayPrivateIpAddress sets GatewayPrivateIpAddress field to given value.

### HasGatewayPrivateIpAddress

`func (o *SessionStatus) HasGatewayPrivateIpAddress() bool`

HasGatewayPrivateIpAddress returns a boolean if a field has been set.

### GetGatewayPublicIpAddress

`func (o *SessionStatus) GetGatewayPublicIpAddress() string`

GetGatewayPublicIpAddress returns the GatewayPublicIpAddress field if non-nil, zero value otherwise.

### GetGatewayPublicIpAddressOk

`func (o *SessionStatus) GetGatewayPublicIpAddressOk() (*string, bool)`

GetGatewayPublicIpAddressOk returns a tuple with the GatewayPublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPublicIpAddress

`func (o *SessionStatus) SetGatewayPublicIpAddress(v string)`

SetGatewayPublicIpAddress sets GatewayPublicIpAddress field to given value.

### HasGatewayPublicIpAddress

`func (o *SessionStatus) HasGatewayPublicIpAddress() bool`

HasGatewayPublicIpAddress returns a boolean if a field has been set.

### GetImei

`func (o *SessionStatus) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *SessionStatus) GetImeiOk() (*string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImei

`func (o *SessionStatus) SetImei(v string)`

SetImei sets Imei field to given value.

### HasImei

`func (o *SessionStatus) HasImei() bool`

HasImei returns a boolean if a field has been set.

### GetImsi

`func (o *SessionStatus) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *SessionStatus) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *SessionStatus) SetImsi(v string)`

SetImsi sets Imsi field to given value.

### HasImsi

`func (o *SessionStatus) HasImsi() bool`

HasImsi returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *SessionStatus) GetLastUpdatedAt() int64`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *SessionStatus) GetLastUpdatedAtOk() (*int64, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *SessionStatus) SetLastUpdatedAt(v int64)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *SessionStatus) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetLocation

`func (o *SessionStatus) GetLocation() map[string]interface{}`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SessionStatus) GetLocationOk() (*map[string]interface{}, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SessionStatus) SetLocation(v map[string]interface{})`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SessionStatus) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetOnline

`func (o *SessionStatus) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *SessionStatus) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *SessionStatus) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *SessionStatus) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetPlacement

`func (o *SessionStatus) GetPlacement() Placement`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *SessionStatus) GetPlacementOk() (*Placement, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *SessionStatus) SetPlacement(v Placement)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *SessionStatus) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetSubscription

`func (o *SessionStatus) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *SessionStatus) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *SessionStatus) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *SessionStatus) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetUeIpAddress

`func (o *SessionStatus) GetUeIpAddress() string`

GetUeIpAddress returns the UeIpAddress field if non-nil, zero value otherwise.

### GetUeIpAddressOk

`func (o *SessionStatus) GetUeIpAddressOk() (*string, bool)`

GetUeIpAddressOk returns a tuple with the UeIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpAddress

`func (o *SessionStatus) SetUeIpAddress(v string)`

SetUeIpAddress sets UeIpAddress field to given value.

### HasUeIpAddress

`func (o *SessionStatus) HasUeIpAddress() bool`

HasUeIpAddress returns a boolean if a field has been set.

### GetVpgId

`func (o *SessionStatus) GetVpgId() string`

GetVpgId returns the VpgId field if non-nil, zero value otherwise.

### GetVpgIdOk

`func (o *SessionStatus) GetVpgIdOk() (*string, bool)`

GetVpgIdOk returns a tuple with the VpgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpgId

`func (o *SessionStatus) SetVpgId(v string)`

SetVpgId sets VpgId field to given value.

### HasVpgId

`func (o *SessionStatus) HasVpgId() bool`

HasVpgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


