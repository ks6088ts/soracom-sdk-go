# SandboxCreateSubscriberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apn** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**ExpiryTime** | Pointer to **int64** |  | [optional] 
**Imsi** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**LastModifiedAt** | Pointer to **int64** |  | [optional] 
**Msisdn** | Pointer to **string** |  | [optional] 
**OperatorId** | Pointer to **string** |  | [optional] 
**RegistrationSecret** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**SpeedClass** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Subscription** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSandboxCreateSubscriberResponse

`func NewSandboxCreateSubscriberResponse() *SandboxCreateSubscriberResponse`

NewSandboxCreateSubscriberResponse instantiates a new SandboxCreateSubscriberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxCreateSubscriberResponseWithDefaults

`func NewSandboxCreateSubscriberResponseWithDefaults() *SandboxCreateSubscriberResponse`

NewSandboxCreateSubscriberResponseWithDefaults instantiates a new SandboxCreateSubscriberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApn

`func (o *SandboxCreateSubscriberResponse) GetApn() string`

GetApn returns the Apn field if non-nil, zero value otherwise.

### GetApnOk

`func (o *SandboxCreateSubscriberResponse) GetApnOk() (*string, bool)`

GetApnOk returns a tuple with the Apn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApn

`func (o *SandboxCreateSubscriberResponse) SetApn(v string)`

SetApn sets Apn field to given value.

### HasApn

`func (o *SandboxCreateSubscriberResponse) HasApn() bool`

HasApn returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SandboxCreateSubscriberResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SandboxCreateSubscriberResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SandboxCreateSubscriberResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SandboxCreateSubscriberResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiryTime

`func (o *SandboxCreateSubscriberResponse) GetExpiryTime() int64`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *SandboxCreateSubscriberResponse) GetExpiryTimeOk() (*int64, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *SandboxCreateSubscriberResponse) SetExpiryTime(v int64)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *SandboxCreateSubscriberResponse) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetImsi

`func (o *SandboxCreateSubscriberResponse) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *SandboxCreateSubscriberResponse) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *SandboxCreateSubscriberResponse) SetImsi(v string)`

SetImsi sets Imsi field to given value.

### HasImsi

`func (o *SandboxCreateSubscriberResponse) HasImsi() bool`

HasImsi returns a boolean if a field has been set.

### GetIpAddress

`func (o *SandboxCreateSubscriberResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *SandboxCreateSubscriberResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *SandboxCreateSubscriberResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *SandboxCreateSubscriberResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLastModifiedAt

`func (o *SandboxCreateSubscriberResponse) GetLastModifiedAt() int64`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *SandboxCreateSubscriberResponse) GetLastModifiedAtOk() (*int64, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *SandboxCreateSubscriberResponse) SetLastModifiedAt(v int64)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *SandboxCreateSubscriberResponse) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.

### GetMsisdn

`func (o *SandboxCreateSubscriberResponse) GetMsisdn() string`

GetMsisdn returns the Msisdn field if non-nil, zero value otherwise.

### GetMsisdnOk

`func (o *SandboxCreateSubscriberResponse) GetMsisdnOk() (*string, bool)`

GetMsisdnOk returns a tuple with the Msisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdn

`func (o *SandboxCreateSubscriberResponse) SetMsisdn(v string)`

SetMsisdn sets Msisdn field to given value.

### HasMsisdn

`func (o *SandboxCreateSubscriberResponse) HasMsisdn() bool`

HasMsisdn returns a boolean if a field has been set.

### GetOperatorId

`func (o *SandboxCreateSubscriberResponse) GetOperatorId() string`

GetOperatorId returns the OperatorId field if non-nil, zero value otherwise.

### GetOperatorIdOk

`func (o *SandboxCreateSubscriberResponse) GetOperatorIdOk() (*string, bool)`

GetOperatorIdOk returns a tuple with the OperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorId

`func (o *SandboxCreateSubscriberResponse) SetOperatorId(v string)`

SetOperatorId sets OperatorId field to given value.

### HasOperatorId

`func (o *SandboxCreateSubscriberResponse) HasOperatorId() bool`

HasOperatorId returns a boolean if a field has been set.

### GetRegistrationSecret

`func (o *SandboxCreateSubscriberResponse) GetRegistrationSecret() string`

GetRegistrationSecret returns the RegistrationSecret field if non-nil, zero value otherwise.

### GetRegistrationSecretOk

`func (o *SandboxCreateSubscriberResponse) GetRegistrationSecretOk() (*string, bool)`

GetRegistrationSecretOk returns a tuple with the RegistrationSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationSecret

`func (o *SandboxCreateSubscriberResponse) SetRegistrationSecret(v string)`

SetRegistrationSecret sets RegistrationSecret field to given value.

### HasRegistrationSecret

`func (o *SandboxCreateSubscriberResponse) HasRegistrationSecret() bool`

HasRegistrationSecret returns a boolean if a field has been set.

### GetSerialNumber

`func (o *SandboxCreateSubscriberResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *SandboxCreateSubscriberResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *SandboxCreateSubscriberResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *SandboxCreateSubscriberResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSpeedClass

`func (o *SandboxCreateSubscriberResponse) GetSpeedClass() string`

GetSpeedClass returns the SpeedClass field if non-nil, zero value otherwise.

### GetSpeedClassOk

`func (o *SandboxCreateSubscriberResponse) GetSpeedClassOk() (*string, bool)`

GetSpeedClassOk returns a tuple with the SpeedClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedClass

`func (o *SandboxCreateSubscriberResponse) SetSpeedClass(v string)`

SetSpeedClass sets SpeedClass field to given value.

### HasSpeedClass

`func (o *SandboxCreateSubscriberResponse) HasSpeedClass() bool`

HasSpeedClass returns a boolean if a field has been set.

### GetStatus

`func (o *SandboxCreateSubscriberResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SandboxCreateSubscriberResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SandboxCreateSubscriberResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SandboxCreateSubscriberResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscription

`func (o *SandboxCreateSubscriberResponse) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *SandboxCreateSubscriberResponse) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *SandboxCreateSubscriberResponse) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *SandboxCreateSubscriberResponse) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetTags

`func (o *SandboxCreateSubscriberResponse) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SandboxCreateSubscriberResponse) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SandboxCreateSubscriberResponse) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SandboxCreateSubscriberResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


