# Subscriber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apn** | Pointer to **string** | この SIM で使用可能な APN（アクセスポイント名） | [optional] 
**Bundles** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **int64** | この SIM が作成された日時のタイムスタンプ | [optional] 
**ExpiredAt** | Pointer to **int64** | この SIM の有効期限が切れた日時のタイムスタンプ | [optional] 
**ExpiredTime** | Pointer to **int64** | この SIM の有効期限が切れた日時のタイムスタンプ | [optional] 
**ExpiryAction** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** | この SIM が所属しているグループの ID | [optional] 
**Iccid** | Pointer to **string** | この SIM の ICCID | [optional] 
**ImeiLock** | Pointer to [**ImeiLock**](ImeiLock.md) |  | [optional] 
**Imsi** | Pointer to **string** | この SIM の IMSI | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**LastModifiedAt** | Pointer to **int64** | この SIM の情報が変更された日時のタイムスタンプ | [optional] 
**LastPortMappingCreatedTime** | Pointer to **int64** | この SIM で Napter オンデマンドリモートアクセスサービスが利用された日時のタイムスタンプ（ミリ秒単位の Unix 時間）。 この SIM で一度も Napter を利用したことがなければ null となります。 | [optional] 
**LocationRegistrationStatus** | Pointer to [**LocationRegistrationStatus**](LocationRegistrationStatus.md) |  | [optional] 
**ModuleType** | Pointer to **string** | 物理的な SIM の形状（フォームファクタ）。取りうる値は、標準 (2FF) サイズの \&quot;mini\&quot;、 3FF サイズの \&quot;micro\&quot;、4FF サイズの \&quot;nano\&quot;、切り取りかた次第で 2FF/3FF/4FF いずれのサイズにもできる \&quot;trio\&quot;、MFF2 とも呼ばれる組込み型 (eSIM) の \&quot;embedded\&quot; です。 | [optional] 
**Msisdn** | Pointer to **string** | この SIM の MSISDN | [optional] 
**OperatorId** | Pointer to **string** | この SIM を所有しているオペレーターのオペレーター ID | [optional] 
**Plan** | Pointer to **int32** | SMS 機能を有しているかどうか。0 なら SMS 未対応、1 なら SMS 対応。 | [optional] 
**PreviousSession** | Pointer to [**PreviousSessionStatus**](PreviousSessionStatus.md) |  | [optional] 
**RegisteredTime** | Pointer to **int64** | この SIM が手動でオペレーターに登録された日時を表すタイムスタンプ（ミリ秒単位の Unix 時刻） ユーザーコンソールから SIM を購入し、自動的に登録された場合はこのフィールドは存在しません。 | [optional] 
**SerialNumber** | Pointer to **string** | この SIM のシリアル番号 | [optional] 
**SessionStatus** | Pointer to [**SessionStatus**](SessionStatus.md) |  | [optional] 
**SimId** | Pointer to **string** | この SIM の SIM ID | [optional] 
**SpeedClass** | Pointer to **string** | この SIM の速度クラス | [optional] 
**Status** | Pointer to **string** | この SIM の契約状態。取りうる値は \&quot;ready\&quot;、\&quot;active\&quot;、\&quot;inactive\&quot;、\&quot;standby\&quot;、 \&quot;suspended\&quot; もしくは \&quot;terminated\&quot; | [optional] 
**Subscription** | Pointer to **string** | この SIM のサブスクリプションの名前 | [optional] 
**Tags** | Pointer to **map[string]string** | An object which always contains at least one property \&quot;name\&quot; with a string value. If you give a subscriber/SIM a name, the name will be returned as the value of the \&quot;name\&quot; property. If the subscriber/SIM does not have a name, an empty string \&quot;\&quot; is returned. In addition, if you create any custom tags for the subscriber/SIM, each custom tag will appear as additional properties in the object. | [optional] 
**TerminationEnabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** | この SIM の速度クラス | [optional] 

## Methods

### NewSubscriber

`func NewSubscriber() *Subscriber`

NewSubscriber instantiates a new Subscriber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriberWithDefaults

`func NewSubscriberWithDefaults() *Subscriber`

NewSubscriberWithDefaults instantiates a new Subscriber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApn

`func (o *Subscriber) GetApn() string`

GetApn returns the Apn field if non-nil, zero value otherwise.

### GetApnOk

`func (o *Subscriber) GetApnOk() (*string, bool)`

GetApnOk returns a tuple with the Apn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApn

`func (o *Subscriber) SetApn(v string)`

SetApn sets Apn field to given value.

### HasApn

`func (o *Subscriber) HasApn() bool`

HasApn returns a boolean if a field has been set.

### GetBundles

`func (o *Subscriber) GetBundles() []string`

GetBundles returns the Bundles field if non-nil, zero value otherwise.

### GetBundlesOk

`func (o *Subscriber) GetBundlesOk() (*[]string, bool)`

GetBundlesOk returns a tuple with the Bundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundles

`func (o *Subscriber) SetBundles(v []string)`

SetBundles sets Bundles field to given value.

### HasBundles

`func (o *Subscriber) HasBundles() bool`

HasBundles returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Subscriber) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Subscriber) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Subscriber) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Subscriber) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiredAt

`func (o *Subscriber) GetExpiredAt() int64`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *Subscriber) GetExpiredAtOk() (*int64, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *Subscriber) SetExpiredAt(v int64)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *Subscriber) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetExpiredTime

`func (o *Subscriber) GetExpiredTime() int64`

GetExpiredTime returns the ExpiredTime field if non-nil, zero value otherwise.

### GetExpiredTimeOk

`func (o *Subscriber) GetExpiredTimeOk() (*int64, bool)`

GetExpiredTimeOk returns a tuple with the ExpiredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredTime

`func (o *Subscriber) SetExpiredTime(v int64)`

SetExpiredTime sets ExpiredTime field to given value.

### HasExpiredTime

`func (o *Subscriber) HasExpiredTime() bool`

HasExpiredTime returns a boolean if a field has been set.

### GetExpiryAction

`func (o *Subscriber) GetExpiryAction() string`

GetExpiryAction returns the ExpiryAction field if non-nil, zero value otherwise.

### GetExpiryActionOk

`func (o *Subscriber) GetExpiryActionOk() (*string, bool)`

GetExpiryActionOk returns a tuple with the ExpiryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryAction

`func (o *Subscriber) SetExpiryAction(v string)`

SetExpiryAction sets ExpiryAction field to given value.

### HasExpiryAction

`func (o *Subscriber) HasExpiryAction() bool`

HasExpiryAction returns a boolean if a field has been set.

### GetGroupId

`func (o *Subscriber) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Subscriber) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Subscriber) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Subscriber) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetIccid

`func (o *Subscriber) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *Subscriber) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *Subscriber) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *Subscriber) HasIccid() bool`

HasIccid returns a boolean if a field has been set.

### GetImeiLock

`func (o *Subscriber) GetImeiLock() ImeiLock`

GetImeiLock returns the ImeiLock field if non-nil, zero value otherwise.

### GetImeiLockOk

`func (o *Subscriber) GetImeiLockOk() (*ImeiLock, bool)`

GetImeiLockOk returns a tuple with the ImeiLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImeiLock

`func (o *Subscriber) SetImeiLock(v ImeiLock)`

SetImeiLock sets ImeiLock field to given value.

### HasImeiLock

`func (o *Subscriber) HasImeiLock() bool`

HasImeiLock returns a boolean if a field has been set.

### GetImsi

`func (o *Subscriber) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *Subscriber) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *Subscriber) SetImsi(v string)`

SetImsi sets Imsi field to given value.

### HasImsi

`func (o *Subscriber) HasImsi() bool`

HasImsi returns a boolean if a field has been set.

### GetIpAddress

`func (o *Subscriber) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Subscriber) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Subscriber) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *Subscriber) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLastModifiedAt

`func (o *Subscriber) GetLastModifiedAt() int64`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *Subscriber) GetLastModifiedAtOk() (*int64, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *Subscriber) SetLastModifiedAt(v int64)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *Subscriber) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.

### GetLastPortMappingCreatedTime

`func (o *Subscriber) GetLastPortMappingCreatedTime() int64`

GetLastPortMappingCreatedTime returns the LastPortMappingCreatedTime field if non-nil, zero value otherwise.

### GetLastPortMappingCreatedTimeOk

`func (o *Subscriber) GetLastPortMappingCreatedTimeOk() (*int64, bool)`

GetLastPortMappingCreatedTimeOk returns a tuple with the LastPortMappingCreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPortMappingCreatedTime

`func (o *Subscriber) SetLastPortMappingCreatedTime(v int64)`

SetLastPortMappingCreatedTime sets LastPortMappingCreatedTime field to given value.

### HasLastPortMappingCreatedTime

`func (o *Subscriber) HasLastPortMappingCreatedTime() bool`

HasLastPortMappingCreatedTime returns a boolean if a field has been set.

### GetLocationRegistrationStatus

`func (o *Subscriber) GetLocationRegistrationStatus() LocationRegistrationStatus`

GetLocationRegistrationStatus returns the LocationRegistrationStatus field if non-nil, zero value otherwise.

### GetLocationRegistrationStatusOk

`func (o *Subscriber) GetLocationRegistrationStatusOk() (*LocationRegistrationStatus, bool)`

GetLocationRegistrationStatusOk returns a tuple with the LocationRegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationRegistrationStatus

`func (o *Subscriber) SetLocationRegistrationStatus(v LocationRegistrationStatus)`

SetLocationRegistrationStatus sets LocationRegistrationStatus field to given value.

### HasLocationRegistrationStatus

`func (o *Subscriber) HasLocationRegistrationStatus() bool`

HasLocationRegistrationStatus returns a boolean if a field has been set.

### GetModuleType

`func (o *Subscriber) GetModuleType() string`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *Subscriber) GetModuleTypeOk() (*string, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *Subscriber) SetModuleType(v string)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *Subscriber) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### GetMsisdn

`func (o *Subscriber) GetMsisdn() string`

GetMsisdn returns the Msisdn field if non-nil, zero value otherwise.

### GetMsisdnOk

`func (o *Subscriber) GetMsisdnOk() (*string, bool)`

GetMsisdnOk returns a tuple with the Msisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdn

`func (o *Subscriber) SetMsisdn(v string)`

SetMsisdn sets Msisdn field to given value.

### HasMsisdn

`func (o *Subscriber) HasMsisdn() bool`

HasMsisdn returns a boolean if a field has been set.

### GetOperatorId

`func (o *Subscriber) GetOperatorId() string`

GetOperatorId returns the OperatorId field if non-nil, zero value otherwise.

### GetOperatorIdOk

`func (o *Subscriber) GetOperatorIdOk() (*string, bool)`

GetOperatorIdOk returns a tuple with the OperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorId

`func (o *Subscriber) SetOperatorId(v string)`

SetOperatorId sets OperatorId field to given value.

### HasOperatorId

`func (o *Subscriber) HasOperatorId() bool`

HasOperatorId returns a boolean if a field has been set.

### GetPlan

`func (o *Subscriber) GetPlan() int32`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Subscriber) GetPlanOk() (*int32, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Subscriber) SetPlan(v int32)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *Subscriber) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPreviousSession

`func (o *Subscriber) GetPreviousSession() PreviousSessionStatus`

GetPreviousSession returns the PreviousSession field if non-nil, zero value otherwise.

### GetPreviousSessionOk

`func (o *Subscriber) GetPreviousSessionOk() (*PreviousSessionStatus, bool)`

GetPreviousSessionOk returns a tuple with the PreviousSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousSession

`func (o *Subscriber) SetPreviousSession(v PreviousSessionStatus)`

SetPreviousSession sets PreviousSession field to given value.

### HasPreviousSession

`func (o *Subscriber) HasPreviousSession() bool`

HasPreviousSession returns a boolean if a field has been set.

### GetRegisteredTime

`func (o *Subscriber) GetRegisteredTime() int64`

GetRegisteredTime returns the RegisteredTime field if non-nil, zero value otherwise.

### GetRegisteredTimeOk

`func (o *Subscriber) GetRegisteredTimeOk() (*int64, bool)`

GetRegisteredTimeOk returns a tuple with the RegisteredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredTime

`func (o *Subscriber) SetRegisteredTime(v int64)`

SetRegisteredTime sets RegisteredTime field to given value.

### HasRegisteredTime

`func (o *Subscriber) HasRegisteredTime() bool`

HasRegisteredTime returns a boolean if a field has been set.

### GetSerialNumber

`func (o *Subscriber) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Subscriber) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Subscriber) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Subscriber) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSessionStatus

`func (o *Subscriber) GetSessionStatus() SessionStatus`

GetSessionStatus returns the SessionStatus field if non-nil, zero value otherwise.

### GetSessionStatusOk

`func (o *Subscriber) GetSessionStatusOk() (*SessionStatus, bool)`

GetSessionStatusOk returns a tuple with the SessionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStatus

`func (o *Subscriber) SetSessionStatus(v SessionStatus)`

SetSessionStatus sets SessionStatus field to given value.

### HasSessionStatus

`func (o *Subscriber) HasSessionStatus() bool`

HasSessionStatus returns a boolean if a field has been set.

### GetSimId

`func (o *Subscriber) GetSimId() string`

GetSimId returns the SimId field if non-nil, zero value otherwise.

### GetSimIdOk

`func (o *Subscriber) GetSimIdOk() (*string, bool)`

GetSimIdOk returns a tuple with the SimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimId

`func (o *Subscriber) SetSimId(v string)`

SetSimId sets SimId field to given value.

### HasSimId

`func (o *Subscriber) HasSimId() bool`

HasSimId returns a boolean if a field has been set.

### GetSpeedClass

`func (o *Subscriber) GetSpeedClass() string`

GetSpeedClass returns the SpeedClass field if non-nil, zero value otherwise.

### GetSpeedClassOk

`func (o *Subscriber) GetSpeedClassOk() (*string, bool)`

GetSpeedClassOk returns a tuple with the SpeedClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedClass

`func (o *Subscriber) SetSpeedClass(v string)`

SetSpeedClass sets SpeedClass field to given value.

### HasSpeedClass

`func (o *Subscriber) HasSpeedClass() bool`

HasSpeedClass returns a boolean if a field has been set.

### GetStatus

`func (o *Subscriber) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Subscriber) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Subscriber) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Subscriber) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscription

`func (o *Subscriber) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *Subscriber) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *Subscriber) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *Subscriber) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetTags

`func (o *Subscriber) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Subscriber) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Subscriber) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Subscriber) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTerminationEnabled

`func (o *Subscriber) GetTerminationEnabled() bool`

GetTerminationEnabled returns the TerminationEnabled field if non-nil, zero value otherwise.

### GetTerminationEnabledOk

`func (o *Subscriber) GetTerminationEnabledOk() (*bool, bool)`

GetTerminationEnabledOk returns a tuple with the TerminationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationEnabled

`func (o *Subscriber) SetTerminationEnabled(v bool)`

SetTerminationEnabled sets TerminationEnabled field to given value.

### HasTerminationEnabled

`func (o *Subscriber) HasTerminationEnabled() bool`

HasTerminationEnabled returns a boolean if a field has been set.

### GetType

`func (o *Subscriber) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Subscriber) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Subscriber) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Subscriber) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


