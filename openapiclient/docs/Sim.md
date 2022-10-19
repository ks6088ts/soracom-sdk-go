# Sim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveProfileId** | Pointer to **string** |  | [optional] 
**ArcSessionStatus** | Pointer to [**ArcSessionStatus**](ArcSessionStatus.md) |  | [optional] 
**Capabilities** | Pointer to [**Capabilities**](Capabilities.md) |  | [optional] 
**CreatedTime** | Pointer to **int64** |  | [optional] 
**ExpiryAction** | Pointer to **string** |  | [optional] 
**ExpiryTime** | Pointer to **int64** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**ImeiLock** | Pointer to [**ImeiLock**](ImeiLock.md) |  | [optional] 
**LastModifiedTime** | Pointer to **int64** |  | [optional] 
**LastPortMappingCreatedTime** | Pointer to **int64** |  | [optional] 
**LocalInfo** | Pointer to **map[string]string** |  | [optional] 
**ModuleType** | Pointer to **string** |  | [optional] 
**OperatorId** | Pointer to **string** |  | [optional] 
**OtaSerialNumber** | Pointer to **string** |  | [optional] 
**PcapEndTime** | Pointer to **int64** |  | [optional] 
**PcapStartTime** | Pointer to **int64** |  | [optional] 
**PreviousSession** | Pointer to [**PreviousSessionStatus**](PreviousSessionStatus.md) |  | [optional] 
**Profiles** | Pointer to [**map[string]SimProfile**](SimProfile.md) |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**SessionStatus** | Pointer to [**SessionStatus**](SessionStatus.md) |  | [optional] 
**SimId** | Pointer to **string** |  | [optional] 
**SpeedClass** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | IoT SIM のステータス - &#x60;ready&#x60;: 準備完了 - &#x60;active&#x60;: 使用中 - &#x60;inactive&#x60;: 休止中 - &#x60;standby&#x60;: 利用開始待ち - &#x60;suspended&#x60;: 利用中断中 - &#x60;terminated&#x60;: 解約済み  | [optional] 
**Tags** | Pointer to **map[string]string** | An object which always contains at least one property \&quot;name\&quot; with a string value. If you give a subscriber/SIM a name, the name will be returned as the value of the \&quot;name\&quot; property. If the subscriber/SIM does not have a name, an empty string \&quot;\&quot; is returned. In addition, if you create any custom tags for the subscriber/SIM, each custom tag will appear as additional properties in the object. | [optional] 
**TerminationEnabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewSim

`func NewSim() *Sim`

NewSim instantiates a new Sim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimWithDefaults

`func NewSimWithDefaults() *Sim`

NewSimWithDefaults instantiates a new Sim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveProfileId

`func (o *Sim) GetActiveProfileId() string`

GetActiveProfileId returns the ActiveProfileId field if non-nil, zero value otherwise.

### GetActiveProfileIdOk

`func (o *Sim) GetActiveProfileIdOk() (*string, bool)`

GetActiveProfileIdOk returns a tuple with the ActiveProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProfileId

`func (o *Sim) SetActiveProfileId(v string)`

SetActiveProfileId sets ActiveProfileId field to given value.

### HasActiveProfileId

`func (o *Sim) HasActiveProfileId() bool`

HasActiveProfileId returns a boolean if a field has been set.

### GetArcSessionStatus

`func (o *Sim) GetArcSessionStatus() ArcSessionStatus`

GetArcSessionStatus returns the ArcSessionStatus field if non-nil, zero value otherwise.

### GetArcSessionStatusOk

`func (o *Sim) GetArcSessionStatusOk() (*ArcSessionStatus, bool)`

GetArcSessionStatusOk returns a tuple with the ArcSessionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcSessionStatus

`func (o *Sim) SetArcSessionStatus(v ArcSessionStatus)`

SetArcSessionStatus sets ArcSessionStatus field to given value.

### HasArcSessionStatus

`func (o *Sim) HasArcSessionStatus() bool`

HasArcSessionStatus returns a boolean if a field has been set.

### GetCapabilities

`func (o *Sim) GetCapabilities() Capabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Sim) GetCapabilitiesOk() (*Capabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Sim) SetCapabilities(v Capabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *Sim) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Sim) GetCreatedTime() int64`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Sim) GetCreatedTimeOk() (*int64, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Sim) SetCreatedTime(v int64)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Sim) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetExpiryAction

`func (o *Sim) GetExpiryAction() string`

GetExpiryAction returns the ExpiryAction field if non-nil, zero value otherwise.

### GetExpiryActionOk

`func (o *Sim) GetExpiryActionOk() (*string, bool)`

GetExpiryActionOk returns a tuple with the ExpiryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryAction

`func (o *Sim) SetExpiryAction(v string)`

SetExpiryAction sets ExpiryAction field to given value.

### HasExpiryAction

`func (o *Sim) HasExpiryAction() bool`

HasExpiryAction returns a boolean if a field has been set.

### GetExpiryTime

`func (o *Sim) GetExpiryTime() int64`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *Sim) GetExpiryTimeOk() (*int64, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *Sim) SetExpiryTime(v int64)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *Sim) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetGroupId

`func (o *Sim) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Sim) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Sim) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Sim) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetImeiLock

`func (o *Sim) GetImeiLock() ImeiLock`

GetImeiLock returns the ImeiLock field if non-nil, zero value otherwise.

### GetImeiLockOk

`func (o *Sim) GetImeiLockOk() (*ImeiLock, bool)`

GetImeiLockOk returns a tuple with the ImeiLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImeiLock

`func (o *Sim) SetImeiLock(v ImeiLock)`

SetImeiLock sets ImeiLock field to given value.

### HasImeiLock

`func (o *Sim) HasImeiLock() bool`

HasImeiLock returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *Sim) GetLastModifiedTime() int64`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *Sim) GetLastModifiedTimeOk() (*int64, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *Sim) SetLastModifiedTime(v int64)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *Sim) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetLastPortMappingCreatedTime

`func (o *Sim) GetLastPortMappingCreatedTime() int64`

GetLastPortMappingCreatedTime returns the LastPortMappingCreatedTime field if non-nil, zero value otherwise.

### GetLastPortMappingCreatedTimeOk

`func (o *Sim) GetLastPortMappingCreatedTimeOk() (*int64, bool)`

GetLastPortMappingCreatedTimeOk returns a tuple with the LastPortMappingCreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPortMappingCreatedTime

`func (o *Sim) SetLastPortMappingCreatedTime(v int64)`

SetLastPortMappingCreatedTime sets LastPortMappingCreatedTime field to given value.

### HasLastPortMappingCreatedTime

`func (o *Sim) HasLastPortMappingCreatedTime() bool`

HasLastPortMappingCreatedTime returns a boolean if a field has been set.

### GetLocalInfo

`func (o *Sim) GetLocalInfo() map[string]string`

GetLocalInfo returns the LocalInfo field if non-nil, zero value otherwise.

### GetLocalInfoOk

`func (o *Sim) GetLocalInfoOk() (*map[string]string, bool)`

GetLocalInfoOk returns a tuple with the LocalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalInfo

`func (o *Sim) SetLocalInfo(v map[string]string)`

SetLocalInfo sets LocalInfo field to given value.

### HasLocalInfo

`func (o *Sim) HasLocalInfo() bool`

HasLocalInfo returns a boolean if a field has been set.

### GetModuleType

`func (o *Sim) GetModuleType() string`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *Sim) GetModuleTypeOk() (*string, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *Sim) SetModuleType(v string)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *Sim) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### GetOperatorId

`func (o *Sim) GetOperatorId() string`

GetOperatorId returns the OperatorId field if non-nil, zero value otherwise.

### GetOperatorIdOk

`func (o *Sim) GetOperatorIdOk() (*string, bool)`

GetOperatorIdOk returns a tuple with the OperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorId

`func (o *Sim) SetOperatorId(v string)`

SetOperatorId sets OperatorId field to given value.

### HasOperatorId

`func (o *Sim) HasOperatorId() bool`

HasOperatorId returns a boolean if a field has been set.

### GetOtaSerialNumber

`func (o *Sim) GetOtaSerialNumber() string`

GetOtaSerialNumber returns the OtaSerialNumber field if non-nil, zero value otherwise.

### GetOtaSerialNumberOk

`func (o *Sim) GetOtaSerialNumberOk() (*string, bool)`

GetOtaSerialNumberOk returns a tuple with the OtaSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtaSerialNumber

`func (o *Sim) SetOtaSerialNumber(v string)`

SetOtaSerialNumber sets OtaSerialNumber field to given value.

### HasOtaSerialNumber

`func (o *Sim) HasOtaSerialNumber() bool`

HasOtaSerialNumber returns a boolean if a field has been set.

### GetPcapEndTime

`func (o *Sim) GetPcapEndTime() int64`

GetPcapEndTime returns the PcapEndTime field if non-nil, zero value otherwise.

### GetPcapEndTimeOk

`func (o *Sim) GetPcapEndTimeOk() (*int64, bool)`

GetPcapEndTimeOk returns a tuple with the PcapEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcapEndTime

`func (o *Sim) SetPcapEndTime(v int64)`

SetPcapEndTime sets PcapEndTime field to given value.

### HasPcapEndTime

`func (o *Sim) HasPcapEndTime() bool`

HasPcapEndTime returns a boolean if a field has been set.

### GetPcapStartTime

`func (o *Sim) GetPcapStartTime() int64`

GetPcapStartTime returns the PcapStartTime field if non-nil, zero value otherwise.

### GetPcapStartTimeOk

`func (o *Sim) GetPcapStartTimeOk() (*int64, bool)`

GetPcapStartTimeOk returns a tuple with the PcapStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcapStartTime

`func (o *Sim) SetPcapStartTime(v int64)`

SetPcapStartTime sets PcapStartTime field to given value.

### HasPcapStartTime

`func (o *Sim) HasPcapStartTime() bool`

HasPcapStartTime returns a boolean if a field has been set.

### GetPreviousSession

`func (o *Sim) GetPreviousSession() PreviousSessionStatus`

GetPreviousSession returns the PreviousSession field if non-nil, zero value otherwise.

### GetPreviousSessionOk

`func (o *Sim) GetPreviousSessionOk() (*PreviousSessionStatus, bool)`

GetPreviousSessionOk returns a tuple with the PreviousSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousSession

`func (o *Sim) SetPreviousSession(v PreviousSessionStatus)`

SetPreviousSession sets PreviousSession field to given value.

### HasPreviousSession

`func (o *Sim) HasPreviousSession() bool`

HasPreviousSession returns a boolean if a field has been set.

### GetProfiles

`func (o *Sim) GetProfiles() map[string]SimProfile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *Sim) GetProfilesOk() (*map[string]SimProfile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *Sim) SetProfiles(v map[string]SimProfile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *Sim) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetSerialNumber

`func (o *Sim) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Sim) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Sim) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Sim) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSessionStatus

`func (o *Sim) GetSessionStatus() SessionStatus`

GetSessionStatus returns the SessionStatus field if non-nil, zero value otherwise.

### GetSessionStatusOk

`func (o *Sim) GetSessionStatusOk() (*SessionStatus, bool)`

GetSessionStatusOk returns a tuple with the SessionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStatus

`func (o *Sim) SetSessionStatus(v SessionStatus)`

SetSessionStatus sets SessionStatus field to given value.

### HasSessionStatus

`func (o *Sim) HasSessionStatus() bool`

HasSessionStatus returns a boolean if a field has been set.

### GetSimId

`func (o *Sim) GetSimId() string`

GetSimId returns the SimId field if non-nil, zero value otherwise.

### GetSimIdOk

`func (o *Sim) GetSimIdOk() (*string, bool)`

GetSimIdOk returns a tuple with the SimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimId

`func (o *Sim) SetSimId(v string)`

SetSimId sets SimId field to given value.

### HasSimId

`func (o *Sim) HasSimId() bool`

HasSimId returns a boolean if a field has been set.

### GetSpeedClass

`func (o *Sim) GetSpeedClass() string`

GetSpeedClass returns the SpeedClass field if non-nil, zero value otherwise.

### GetSpeedClassOk

`func (o *Sim) GetSpeedClassOk() (*string, bool)`

GetSpeedClassOk returns a tuple with the SpeedClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedClass

`func (o *Sim) SetSpeedClass(v string)`

SetSpeedClass sets SpeedClass field to given value.

### HasSpeedClass

`func (o *Sim) HasSpeedClass() bool`

HasSpeedClass returns a boolean if a field has been set.

### GetStatus

`func (o *Sim) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Sim) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Sim) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Sim) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *Sim) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Sim) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Sim) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Sim) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTerminationEnabled

`func (o *Sim) GetTerminationEnabled() bool`

GetTerminationEnabled returns the TerminationEnabled field if non-nil, zero value otherwise.

### GetTerminationEnabledOk

`func (o *Sim) GetTerminationEnabledOk() (*bool, bool)`

GetTerminationEnabledOk returns a tuple with the TerminationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationEnabled

`func (o *Sim) SetTerminationEnabled(v bool)`

SetTerminationEnabled sets TerminationEnabled field to given value.

### HasTerminationEnabled

`func (o *Sim) HasTerminationEnabled() bool`

HasTerminationEnabled returns a boolean if a field has been set.

### GetType

`func (o *Sim) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Sim) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Sim) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Sim) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


