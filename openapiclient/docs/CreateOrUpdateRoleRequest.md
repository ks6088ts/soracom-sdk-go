# CreateOrUpdateRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Permission** | **string** | Permission の JSON | 

## Methods

### NewCreateOrUpdateRoleRequest

`func NewCreateOrUpdateRoleRequest(permission string, ) *CreateOrUpdateRoleRequest`

NewCreateOrUpdateRoleRequest instantiates a new CreateOrUpdateRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateRoleRequestWithDefaults

`func NewCreateOrUpdateRoleRequestWithDefaults() *CreateOrUpdateRoleRequest`

NewCreateOrUpdateRoleRequestWithDefaults instantiates a new CreateOrUpdateRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateOrUpdateRoleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrUpdateRoleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrUpdateRoleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrUpdateRoleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPermission

`func (o *CreateOrUpdateRoleRequest) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *CreateOrUpdateRoleRequest) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *CreateOrUpdateRoleRequest) SetPermission(v string)`

SetPermission sets Permission field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


