# SessionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apn** | Pointer to **string** | The Access Point Name configured. | [optional] 
**Dns0** | Pointer to **string** | First IP address of DNS servers. | [optional] 
**Dns1** | Pointer to **string** | Second IP address of DNS servers. | [optional] 
**Event** | Pointer to **string** | The type of behavior for the event. Possible values are \&quot;Created\&quot; indicates the device created a new session, \&quot;Modified\&quot; indicates that an existing connection was modified, \&quot;Deleted\&quot; indicates a network connection was closed. | [optional] 
**GatewayPrivateIpAddress** | Pointer to **string** |  | [optional] 
**GatewayPublicIpAddress** | Pointer to **string** |  | [optional] 
**Imei** | Pointer to **string** | The IMEI of the device using the SIM. | [optional] 
**Imsi** | Pointer to **string** | The IMSI of the SIM. | [optional] 
**OperatorId** | Pointer to **string** | The operator ID of the session event. | [optional] 
**Time** | Pointer to **int64** | The timestamp of the session event. | [optional] 
**UeIpAddress** | Pointer to **string** | The IP address of the device. | [optional] 
**VpgId** | Pointer to **string** | The Virtual Private Gateway IP address configured. | [optional] 

## Methods

### NewSessionEvent

`func NewSessionEvent() *SessionEvent`

NewSessionEvent instantiates a new SessionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionEventWithDefaults

`func NewSessionEventWithDefaults() *SessionEvent`

NewSessionEventWithDefaults instantiates a new SessionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApn

`func (o *SessionEvent) GetApn() string`

GetApn returns the Apn field if non-nil, zero value otherwise.

### GetApnOk

`func (o *SessionEvent) GetApnOk() (*string, bool)`

GetApnOk returns a tuple with the Apn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApn

`func (o *SessionEvent) SetApn(v string)`

SetApn sets Apn field to given value.

### HasApn

`func (o *SessionEvent) HasApn() bool`

HasApn returns a boolean if a field has been set.

### GetDns0

`func (o *SessionEvent) GetDns0() string`

GetDns0 returns the Dns0 field if non-nil, zero value otherwise.

### GetDns0Ok

`func (o *SessionEvent) GetDns0Ok() (*string, bool)`

GetDns0Ok returns a tuple with the Dns0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns0

`func (o *SessionEvent) SetDns0(v string)`

SetDns0 sets Dns0 field to given value.

### HasDns0

`func (o *SessionEvent) HasDns0() bool`

HasDns0 returns a boolean if a field has been set.

### GetDns1

`func (o *SessionEvent) GetDns1() string`

GetDns1 returns the Dns1 field if non-nil, zero value otherwise.

### GetDns1Ok

`func (o *SessionEvent) GetDns1Ok() (*string, bool)`

GetDns1Ok returns a tuple with the Dns1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns1

`func (o *SessionEvent) SetDns1(v string)`

SetDns1 sets Dns1 field to given value.

### HasDns1

`func (o *SessionEvent) HasDns1() bool`

HasDns1 returns a boolean if a field has been set.

### GetEvent

`func (o *SessionEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *SessionEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *SessionEvent) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *SessionEvent) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetGatewayPrivateIpAddress

`func (o *SessionEvent) GetGatewayPrivateIpAddress() string`

GetGatewayPrivateIpAddress returns the GatewayPrivateIpAddress field if non-nil, zero value otherwise.

### GetGatewayPrivateIpAddressOk

`func (o *SessionEvent) GetGatewayPrivateIpAddressOk() (*string, bool)`

GetGatewayPrivateIpAddressOk returns a tuple with the GatewayPrivateIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPrivateIpAddress

`func (o *SessionEvent) SetGatewayPrivateIpAddress(v string)`

SetGatewayPrivateIpAddress sets GatewayPrivateIpAddress field to given value.

### HasGatewayPrivateIpAddress

`func (o *SessionEvent) HasGatewayPrivateIpAddress() bool`

HasGatewayPrivateIpAddress returns a boolean if a field has been set.

### GetGatewayPublicIpAddress

`func (o *SessionEvent) GetGatewayPublicIpAddress() string`

GetGatewayPublicIpAddress returns the GatewayPublicIpAddress field if non-nil, zero value otherwise.

### GetGatewayPublicIpAddressOk

`func (o *SessionEvent) GetGatewayPublicIpAddressOk() (*string, bool)`

GetGatewayPublicIpAddressOk returns a tuple with the GatewayPublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPublicIpAddress

`func (o *SessionEvent) SetGatewayPublicIpAddress(v string)`

SetGatewayPublicIpAddress sets GatewayPublicIpAddress field to given value.

### HasGatewayPublicIpAddress

`func (o *SessionEvent) HasGatewayPublicIpAddress() bool`

HasGatewayPublicIpAddress returns a boolean if a field has been set.

### GetImei

`func (o *SessionEvent) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *SessionEvent) GetImeiOk() (*string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImei

`func (o *SessionEvent) SetImei(v string)`

SetImei sets Imei field to given value.

### HasImei

`func (o *SessionEvent) HasImei() bool`

HasImei returns a boolean if a field has been set.

### GetImsi

`func (o *SessionEvent) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *SessionEvent) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *SessionEvent) SetImsi(v string)`

SetImsi sets Imsi field to given value.

### HasImsi

`func (o *SessionEvent) HasImsi() bool`

HasImsi returns a boolean if a field has been set.

### GetOperatorId

`func (o *SessionEvent) GetOperatorId() string`

GetOperatorId returns the OperatorId field if non-nil, zero value otherwise.

### GetOperatorIdOk

`func (o *SessionEvent) GetOperatorIdOk() (*string, bool)`

GetOperatorIdOk returns a tuple with the OperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorId

`func (o *SessionEvent) SetOperatorId(v string)`

SetOperatorId sets OperatorId field to given value.

### HasOperatorId

`func (o *SessionEvent) HasOperatorId() bool`

HasOperatorId returns a boolean if a field has been set.

### GetTime

`func (o *SessionEvent) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *SessionEvent) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *SessionEvent) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *SessionEvent) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetUeIpAddress

`func (o *SessionEvent) GetUeIpAddress() string`

GetUeIpAddress returns the UeIpAddress field if non-nil, zero value otherwise.

### GetUeIpAddressOk

`func (o *SessionEvent) GetUeIpAddressOk() (*string, bool)`

GetUeIpAddressOk returns a tuple with the UeIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpAddress

`func (o *SessionEvent) SetUeIpAddress(v string)`

SetUeIpAddress sets UeIpAddress field to given value.

### HasUeIpAddress

`func (o *SessionEvent) HasUeIpAddress() bool`

HasUeIpAddress returns a boolean if a field has been set.

### GetVpgId

`func (o *SessionEvent) GetVpgId() string`

GetVpgId returns the VpgId field if non-nil, zero value otherwise.

### GetVpgIdOk

`func (o *SessionEvent) GetVpgIdOk() (*string, bool)`

GetVpgIdOk returns a tuple with the VpgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpgId

`func (o *SessionEvent) SetVpgId(v string)`

SetVpgId sets VpgId field to given value.

### HasVpgId

`func (o *SessionEvent) HasVpgId() bool`

HasVpgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


