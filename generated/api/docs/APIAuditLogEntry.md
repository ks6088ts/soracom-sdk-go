# APIAuditLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKind** | Pointer to **string** | API 種別 (例: &#x60;/v1/auth&#x60;)。 | [optional] 
**OperatorId** | Pointer to **string** | API を呼び出したオペレーターID。 | [optional] 
**RemoteIpAddress** | Pointer to **string** | API の呼び出し元 IP アドレス。 | [optional] 
**RequestPath** | Pointer to **string** | リクエストされた API のパス。 | [optional] 
**RequestedTimeEpochMs** | Pointer to **int64** | API が呼び出された時刻。この値はページネーションのために &#x60;last_evaluated_key&#x60; リクエストパラメーターに利用することができます。 | [optional] 
**UserName** | Pointer to **string** | API を呼び出した SAM ユーザー名。もしこの値が空の場合は呼び出しユーザーは root ユーザーです。 | [optional] 

## Methods

### NewAPIAuditLogEntry

`func NewAPIAuditLogEntry() *APIAuditLogEntry`

NewAPIAuditLogEntry instantiates a new APIAuditLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIAuditLogEntryWithDefaults

`func NewAPIAuditLogEntryWithDefaults() *APIAuditLogEntry`

NewAPIAuditLogEntryWithDefaults instantiates a new APIAuditLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKind

`func (o *APIAuditLogEntry) GetApiKind() string`

GetApiKind returns the ApiKind field if non-nil, zero value otherwise.

### GetApiKindOk

`func (o *APIAuditLogEntry) GetApiKindOk() (*string, bool)`

GetApiKindOk returns a tuple with the ApiKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKind

`func (o *APIAuditLogEntry) SetApiKind(v string)`

SetApiKind sets ApiKind field to given value.

### HasApiKind

`func (o *APIAuditLogEntry) HasApiKind() bool`

HasApiKind returns a boolean if a field has been set.

### GetOperatorId

`func (o *APIAuditLogEntry) GetOperatorId() string`

GetOperatorId returns the OperatorId field if non-nil, zero value otherwise.

### GetOperatorIdOk

`func (o *APIAuditLogEntry) GetOperatorIdOk() (*string, bool)`

GetOperatorIdOk returns a tuple with the OperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorId

`func (o *APIAuditLogEntry) SetOperatorId(v string)`

SetOperatorId sets OperatorId field to given value.

### HasOperatorId

`func (o *APIAuditLogEntry) HasOperatorId() bool`

HasOperatorId returns a boolean if a field has been set.

### GetRemoteIpAddress

`func (o *APIAuditLogEntry) GetRemoteIpAddress() string`

GetRemoteIpAddress returns the RemoteIpAddress field if non-nil, zero value otherwise.

### GetRemoteIpAddressOk

`func (o *APIAuditLogEntry) GetRemoteIpAddressOk() (*string, bool)`

GetRemoteIpAddressOk returns a tuple with the RemoteIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpAddress

`func (o *APIAuditLogEntry) SetRemoteIpAddress(v string)`

SetRemoteIpAddress sets RemoteIpAddress field to given value.

### HasRemoteIpAddress

`func (o *APIAuditLogEntry) HasRemoteIpAddress() bool`

HasRemoteIpAddress returns a boolean if a field has been set.

### GetRequestPath

`func (o *APIAuditLogEntry) GetRequestPath() string`

GetRequestPath returns the RequestPath field if non-nil, zero value otherwise.

### GetRequestPathOk

`func (o *APIAuditLogEntry) GetRequestPathOk() (*string, bool)`

GetRequestPathOk returns a tuple with the RequestPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPath

`func (o *APIAuditLogEntry) SetRequestPath(v string)`

SetRequestPath sets RequestPath field to given value.

### HasRequestPath

`func (o *APIAuditLogEntry) HasRequestPath() bool`

HasRequestPath returns a boolean if a field has been set.

### GetRequestedTimeEpochMs

`func (o *APIAuditLogEntry) GetRequestedTimeEpochMs() int64`

GetRequestedTimeEpochMs returns the RequestedTimeEpochMs field if non-nil, zero value otherwise.

### GetRequestedTimeEpochMsOk

`func (o *APIAuditLogEntry) GetRequestedTimeEpochMsOk() (*int64, bool)`

GetRequestedTimeEpochMsOk returns a tuple with the RequestedTimeEpochMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTimeEpochMs

`func (o *APIAuditLogEntry) SetRequestedTimeEpochMs(v int64)`

SetRequestedTimeEpochMs sets RequestedTimeEpochMs field to given value.

### HasRequestedTimeEpochMs

`func (o *APIAuditLogEntry) HasRequestedTimeEpochMs() bool`

HasRequestedTimeEpochMs returns a boolean if a field has been set.

### GetUserName

`func (o *APIAuditLogEntry) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *APIAuditLogEntry) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *APIAuditLogEntry) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *APIAuditLogEntry) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


