# LagoonUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | This value used on login. | [optional] 
**Id** | Pointer to **float32** |  | [optional] 
**LastSeenAt** | Pointer to **string** | The last login datetime. | [optional] 
**LastSeenAtAge** | Pointer to **string** | The last login datetime as age. | [optional] 
**Role** | Pointer to **string** | A role that represents the permission. | [optional] 

## Methods

### NewLagoonUser

`func NewLagoonUser() *LagoonUser`

NewLagoonUser instantiates a new LagoonUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLagoonUserWithDefaults

`func NewLagoonUserWithDefaults() *LagoonUser`

NewLagoonUserWithDefaults instantiates a new LagoonUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *LagoonUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LagoonUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LagoonUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LagoonUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *LagoonUser) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LagoonUser) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LagoonUser) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *LagoonUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastSeenAt

`func (o *LagoonUser) GetLastSeenAt() string`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *LagoonUser) GetLastSeenAtOk() (*string, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *LagoonUser) SetLastSeenAt(v string)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *LagoonUser) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### GetLastSeenAtAge

`func (o *LagoonUser) GetLastSeenAtAge() string`

GetLastSeenAtAge returns the LastSeenAtAge field if non-nil, zero value otherwise.

### GetLastSeenAtAgeOk

`func (o *LagoonUser) GetLastSeenAtAgeOk() (*string, bool)`

GetLastSeenAtAgeOk returns a tuple with the LastSeenAtAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAtAge

`func (o *LagoonUser) SetLastSeenAtAge(v string)`

SetLastSeenAtAge sets LastSeenAtAge field to given value.

### HasLastSeenAtAge

`func (o *LagoonUser) HasLastSeenAtAge() bool`

HasLastSeenAtAge returns a boolean if a field has been set.

### GetRole

`func (o *LagoonUser) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *LagoonUser) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *LagoonUser) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *LagoonUser) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


