# SandboxInitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthKey** | **string** |  | 
**AuthKeyId** | **string** |  | 
**CoverageTypes** | Pointer to **[]string** | カバレッジタイプ - &#x60;g&#x60;: グローバルカバレッジ - &#x60;jp&#x60;: 日本カバレッジ  | [optional] 
**Email** | **string** |  | 
**Password** | **string** |  | 
**RegisterPaymentMethod** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewSandboxInitRequest

`func NewSandboxInitRequest(authKey string, authKeyId string, email string, password string, ) *SandboxInitRequest`

NewSandboxInitRequest instantiates a new SandboxInitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxInitRequestWithDefaults

`func NewSandboxInitRequestWithDefaults() *SandboxInitRequest`

NewSandboxInitRequestWithDefaults instantiates a new SandboxInitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthKey

`func (o *SandboxInitRequest) GetAuthKey() string`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *SandboxInitRequest) GetAuthKeyOk() (*string, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *SandboxInitRequest) SetAuthKey(v string)`

SetAuthKey sets AuthKey field to given value.


### GetAuthKeyId

`func (o *SandboxInitRequest) GetAuthKeyId() string`

GetAuthKeyId returns the AuthKeyId field if non-nil, zero value otherwise.

### GetAuthKeyIdOk

`func (o *SandboxInitRequest) GetAuthKeyIdOk() (*string, bool)`

GetAuthKeyIdOk returns a tuple with the AuthKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKeyId

`func (o *SandboxInitRequest) SetAuthKeyId(v string)`

SetAuthKeyId sets AuthKeyId field to given value.


### GetCoverageTypes

`func (o *SandboxInitRequest) GetCoverageTypes() []string`

GetCoverageTypes returns the CoverageTypes field if non-nil, zero value otherwise.

### GetCoverageTypesOk

`func (o *SandboxInitRequest) GetCoverageTypesOk() (*[]string, bool)`

GetCoverageTypesOk returns a tuple with the CoverageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageTypes

`func (o *SandboxInitRequest) SetCoverageTypes(v []string)`

SetCoverageTypes sets CoverageTypes field to given value.

### HasCoverageTypes

`func (o *SandboxInitRequest) HasCoverageTypes() bool`

HasCoverageTypes returns a boolean if a field has been set.

### GetEmail

`func (o *SandboxInitRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SandboxInitRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SandboxInitRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *SandboxInitRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SandboxInitRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SandboxInitRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRegisterPaymentMethod

`func (o *SandboxInitRequest) GetRegisterPaymentMethod() bool`

GetRegisterPaymentMethod returns the RegisterPaymentMethod field if non-nil, zero value otherwise.

### GetRegisterPaymentMethodOk

`func (o *SandboxInitRequest) GetRegisterPaymentMethodOk() (*bool, bool)`

GetRegisterPaymentMethodOk returns a tuple with the RegisterPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterPaymentMethod

`func (o *SandboxInitRequest) SetRegisterPaymentMethod(v bool)`

SetRegisterPaymentMethod sets RegisterPaymentMethod field to given value.

### HasRegisterPaymentMethod

`func (o *SandboxInitRequest) HasRegisterPaymentMethod() bool`

HasRegisterPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


