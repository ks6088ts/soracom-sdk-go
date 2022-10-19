# NapterAuditLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Direction** | Pointer to [**NapterAuditLogDirection**](NapterAuditLogDirection.md) |  | [optional] 
**Imsi** | Pointer to **string** |  | [optional] 
**IsTLS** | Pointer to **bool** |  | [optional] 
**OperatorId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewNapterAuditLogEntry

`func NewNapterAuditLogEntry() *NapterAuditLogEntry`

NewNapterAuditLogEntry instantiates a new NapterAuditLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNapterAuditLogEntryWithDefaults

`func NewNapterAuditLogEntryWithDefaults() *NapterAuditLogEntry`

NewNapterAuditLogEntryWithDefaults instantiates a new NapterAuditLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *NapterAuditLogEntry) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *NapterAuditLogEntry) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *NapterAuditLogEntry) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *NapterAuditLogEntry) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NapterAuditLogEntry) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NapterAuditLogEntry) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NapterAuditLogEntry) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NapterAuditLogEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDirection

`func (o *NapterAuditLogEntry) GetDirection() NapterAuditLogDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *NapterAuditLogEntry) GetDirectionOk() (*NapterAuditLogDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *NapterAuditLogEntry) SetDirection(v NapterAuditLogDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *NapterAuditLogEntry) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetImsi

`func (o *NapterAuditLogEntry) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *NapterAuditLogEntry) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *NapterAuditLogEntry) SetImsi(v string)`

SetImsi sets Imsi field to given value.

### HasImsi

`func (o *NapterAuditLogEntry) HasImsi() bool`

HasImsi returns a boolean if a field has been set.

### GetIsTLS

`func (o *NapterAuditLogEntry) GetIsTLS() bool`

GetIsTLS returns the IsTLS field if non-nil, zero value otherwise.

### GetIsTLSOk

`func (o *NapterAuditLogEntry) GetIsTLSOk() (*bool, bool)`

GetIsTLSOk returns a tuple with the IsTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTLS

`func (o *NapterAuditLogEntry) SetIsTLS(v bool)`

SetIsTLS sets IsTLS field to given value.

### HasIsTLS

`func (o *NapterAuditLogEntry) HasIsTLS() bool`

HasIsTLS returns a boolean if a field has been set.

### GetOperatorId

`func (o *NapterAuditLogEntry) GetOperatorId() string`

GetOperatorId returns the OperatorId field if non-nil, zero value otherwise.

### GetOperatorIdOk

`func (o *NapterAuditLogEntry) GetOperatorIdOk() (*string, bool)`

GetOperatorIdOk returns a tuple with the OperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorId

`func (o *NapterAuditLogEntry) SetOperatorId(v string)`

SetOperatorId sets OperatorId field to given value.

### HasOperatorId

`func (o *NapterAuditLogEntry) HasOperatorId() bool`

HasOperatorId returns a boolean if a field has been set.

### GetType

`func (o *NapterAuditLogEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NapterAuditLogEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NapterAuditLogEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NapterAuditLogEntry) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


