# LagoonUserCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** | A role that represents the permission. | [optional] 
**UserEmail** | Pointer to **string** |  | [optional] 
**UserPassword** | Pointer to **string** |  | [optional] 

## Methods

### NewLagoonUserCreationRequest

`func NewLagoonUserCreationRequest() *LagoonUserCreationRequest`

NewLagoonUserCreationRequest instantiates a new LagoonUserCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLagoonUserCreationRequestWithDefaults

`func NewLagoonUserCreationRequestWithDefaults() *LagoonUserCreationRequest`

NewLagoonUserCreationRequestWithDefaults instantiates a new LagoonUserCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *LagoonUserCreationRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *LagoonUserCreationRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *LagoonUserCreationRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *LagoonUserCreationRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUserEmail

`func (o *LagoonUserCreationRequest) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *LagoonUserCreationRequest) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *LagoonUserCreationRequest) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *LagoonUserCreationRequest) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetUserPassword

`func (o *LagoonUserCreationRequest) GetUserPassword() string`

GetUserPassword returns the UserPassword field if non-nil, zero value otherwise.

### GetUserPasswordOk

`func (o *LagoonUserCreationRequest) GetUserPasswordOk() (*string, bool)`

GetUserPasswordOk returns a tuple with the UserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPassword

`func (o *LagoonUserCreationRequest) SetUserPassword(v string)`

SetUserPassword sets UserPassword field to given value.

### HasUserPassword

`func (o *LagoonUserCreationRequest) HasUserPassword() bool`

HasUserPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


