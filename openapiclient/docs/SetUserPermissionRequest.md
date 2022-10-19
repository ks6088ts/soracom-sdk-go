# SetUserPermissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Permission** | **string** | Permission の JSON | 

## Methods

### NewSetUserPermissionRequest

`func NewSetUserPermissionRequest(permission string, ) *SetUserPermissionRequest`

NewSetUserPermissionRequest instantiates a new SetUserPermissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetUserPermissionRequestWithDefaults

`func NewSetUserPermissionRequestWithDefaults() *SetUserPermissionRequest`

NewSetUserPermissionRequestWithDefaults instantiates a new SetUserPermissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SetUserPermissionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SetUserPermissionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SetUserPermissionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SetUserPermissionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPermission

`func (o *SetUserPermissionRequest) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *SetUserPermissionRequest) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *SetUserPermissionRequest) SetPermission(v string)`

SetPermission sets Permission field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


