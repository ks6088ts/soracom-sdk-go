# RoleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateDateTime** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Permission** | Pointer to **string** | Permission の JSON | [optional] 
**RoleId** | Pointer to **string** |  | [optional] 
**UpdateDateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewRoleResponse

`func NewRoleResponse() *RoleResponse`

NewRoleResponse instantiates a new RoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleResponseWithDefaults

`func NewRoleResponseWithDefaults() *RoleResponse`

NewRoleResponseWithDefaults instantiates a new RoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateDateTime

`func (o *RoleResponse) GetCreateDateTime() int64`

GetCreateDateTime returns the CreateDateTime field if non-nil, zero value otherwise.

### GetCreateDateTimeOk

`func (o *RoleResponse) GetCreateDateTimeOk() (*int64, bool)`

GetCreateDateTimeOk returns a tuple with the CreateDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDateTime

`func (o *RoleResponse) SetCreateDateTime(v int64)`

SetCreateDateTime sets CreateDateTime field to given value.

### HasCreateDateTime

`func (o *RoleResponse) HasCreateDateTime() bool`

HasCreateDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *RoleResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPermission

`func (o *RoleResponse) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *RoleResponse) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *RoleResponse) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *RoleResponse) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetRoleId

`func (o *RoleResponse) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *RoleResponse) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *RoleResponse) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *RoleResponse) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetUpdateDateTime

`func (o *RoleResponse) GetUpdateDateTime() int64`

GetUpdateDateTime returns the UpdateDateTime field if non-nil, zero value otherwise.

### GetUpdateDateTimeOk

`func (o *RoleResponse) GetUpdateDateTimeOk() (*int64, bool)`

GetUpdateDateTimeOk returns a tuple with the UpdateDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDateTime

`func (o *RoleResponse) SetUpdateDateTime(v int64)`

SetUpdateDateTime sets UpdateDateTime field to given value.

### HasUpdateDateTime

`func (o *RoleResponse) HasUpdateDateTime() bool`

HasUpdateDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


