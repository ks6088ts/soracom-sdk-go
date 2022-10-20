# UserDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthKeyList** | Pointer to [**[]AuthKeyResponse**](AuthKeyResponse.md) |  | [optional] 
**CreateDateTime** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**HasPassword** | Pointer to **bool** |  | [optional] 
**Permission** | Pointer to **string** |  | [optional] 
**RoleList** | Pointer to [**[]ListRolesResponse**](ListRolesResponse.md) |  | [optional] 
**UpdateDateTime** | Pointer to **int64** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewUserDetailResponse

`func NewUserDetailResponse() *UserDetailResponse`

NewUserDetailResponse instantiates a new UserDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDetailResponseWithDefaults

`func NewUserDetailResponseWithDefaults() *UserDetailResponse`

NewUserDetailResponseWithDefaults instantiates a new UserDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthKeyList

`func (o *UserDetailResponse) GetAuthKeyList() []AuthKeyResponse`

GetAuthKeyList returns the AuthKeyList field if non-nil, zero value otherwise.

### GetAuthKeyListOk

`func (o *UserDetailResponse) GetAuthKeyListOk() (*[]AuthKeyResponse, bool)`

GetAuthKeyListOk returns a tuple with the AuthKeyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKeyList

`func (o *UserDetailResponse) SetAuthKeyList(v []AuthKeyResponse)`

SetAuthKeyList sets AuthKeyList field to given value.

### HasAuthKeyList

`func (o *UserDetailResponse) HasAuthKeyList() bool`

HasAuthKeyList returns a boolean if a field has been set.

### GetCreateDateTime

`func (o *UserDetailResponse) GetCreateDateTime() int64`

GetCreateDateTime returns the CreateDateTime field if non-nil, zero value otherwise.

### GetCreateDateTimeOk

`func (o *UserDetailResponse) GetCreateDateTimeOk() (*int64, bool)`

GetCreateDateTimeOk returns a tuple with the CreateDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDateTime

`func (o *UserDetailResponse) SetCreateDateTime(v int64)`

SetCreateDateTime sets CreateDateTime field to given value.

### HasCreateDateTime

`func (o *UserDetailResponse) HasCreateDateTime() bool`

HasCreateDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *UserDetailResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserDetailResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserDetailResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserDetailResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHasPassword

`func (o *UserDetailResponse) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *UserDetailResponse) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *UserDetailResponse) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *UserDetailResponse) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### GetPermission

`func (o *UserDetailResponse) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *UserDetailResponse) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *UserDetailResponse) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *UserDetailResponse) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetRoleList

`func (o *UserDetailResponse) GetRoleList() []ListRolesResponse`

GetRoleList returns the RoleList field if non-nil, zero value otherwise.

### GetRoleListOk

`func (o *UserDetailResponse) GetRoleListOk() (*[]ListRolesResponse, bool)`

GetRoleListOk returns a tuple with the RoleList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleList

`func (o *UserDetailResponse) SetRoleList(v []ListRolesResponse)`

SetRoleList sets RoleList field to given value.

### HasRoleList

`func (o *UserDetailResponse) HasRoleList() bool`

HasRoleList returns a boolean if a field has been set.

### GetUpdateDateTime

`func (o *UserDetailResponse) GetUpdateDateTime() int64`

GetUpdateDateTime returns the UpdateDateTime field if non-nil, zero value otherwise.

### GetUpdateDateTimeOk

`func (o *UserDetailResponse) GetUpdateDateTimeOk() (*int64, bool)`

GetUpdateDateTimeOk returns a tuple with the UpdateDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDateTime

`func (o *UserDetailResponse) SetUpdateDateTime(v int64)`

SetUpdateDateTime sets UpdateDateTime field to given value.

### HasUpdateDateTime

`func (o *UserDetailResponse) HasUpdateDateTime() bool`

HasUpdateDateTime returns a boolean if a field has been set.

### GetUserName

`func (o *UserDetailResponse) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UserDetailResponse) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UserDetailResponse) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UserDetailResponse) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


