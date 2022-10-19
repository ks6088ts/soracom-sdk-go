# AuthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthKey** | Pointer to **string** |  | [optional] 
**AuthKeyId** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**MfaOTPCode** | Pointer to **string** |  | [optional] 
**OperatorId** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**TokenTimeoutSeconds** | Pointer to **int64** |  | [optional] [default to 86400]
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthRequest

`func NewAuthRequest() *AuthRequest`

NewAuthRequest instantiates a new AuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthRequestWithDefaults

`func NewAuthRequestWithDefaults() *AuthRequest`

NewAuthRequestWithDefaults instantiates a new AuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthKey

`func (o *AuthRequest) GetAuthKey() string`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *AuthRequest) GetAuthKeyOk() (*string, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *AuthRequest) SetAuthKey(v string)`

SetAuthKey sets AuthKey field to given value.

### HasAuthKey

`func (o *AuthRequest) HasAuthKey() bool`

HasAuthKey returns a boolean if a field has been set.

### GetAuthKeyId

`func (o *AuthRequest) GetAuthKeyId() string`

GetAuthKeyId returns the AuthKeyId field if non-nil, zero value otherwise.

### GetAuthKeyIdOk

`func (o *AuthRequest) GetAuthKeyIdOk() (*string, bool)`

GetAuthKeyIdOk returns a tuple with the AuthKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKeyId

`func (o *AuthRequest) SetAuthKeyId(v string)`

SetAuthKeyId sets AuthKeyId field to given value.

### HasAuthKeyId

`func (o *AuthRequest) HasAuthKeyId() bool`

HasAuthKeyId returns a boolean if a field has been set.

### GetEmail

`func (o *AuthRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuthRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMfaOTPCode

`func (o *AuthRequest) GetMfaOTPCode() string`

GetMfaOTPCode returns the MfaOTPCode field if non-nil, zero value otherwise.

### GetMfaOTPCodeOk

`func (o *AuthRequest) GetMfaOTPCodeOk() (*string, bool)`

GetMfaOTPCodeOk returns a tuple with the MfaOTPCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaOTPCode

`func (o *AuthRequest) SetMfaOTPCode(v string)`

SetMfaOTPCode sets MfaOTPCode field to given value.

### HasMfaOTPCode

`func (o *AuthRequest) HasMfaOTPCode() bool`

HasMfaOTPCode returns a boolean if a field has been set.

### GetOperatorId

`func (o *AuthRequest) GetOperatorId() string`

GetOperatorId returns the OperatorId field if non-nil, zero value otherwise.

### GetOperatorIdOk

`func (o *AuthRequest) GetOperatorIdOk() (*string, bool)`

GetOperatorIdOk returns a tuple with the OperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorId

`func (o *AuthRequest) SetOperatorId(v string)`

SetOperatorId sets OperatorId field to given value.

### HasOperatorId

`func (o *AuthRequest) HasOperatorId() bool`

HasOperatorId returns a boolean if a field has been set.

### GetPassword

`func (o *AuthRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AuthRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTokenTimeoutSeconds

`func (o *AuthRequest) GetTokenTimeoutSeconds() int64`

GetTokenTimeoutSeconds returns the TokenTimeoutSeconds field if non-nil, zero value otherwise.

### GetTokenTimeoutSecondsOk

`func (o *AuthRequest) GetTokenTimeoutSecondsOk() (*int64, bool)`

GetTokenTimeoutSecondsOk returns a tuple with the TokenTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTimeoutSeconds

`func (o *AuthRequest) SetTokenTimeoutSeconds(v int64)`

SetTokenTimeoutSeconds sets TokenTimeoutSeconds field to given value.

### HasTokenTimeoutSeconds

`func (o *AuthRequest) HasTokenTimeoutSeconds() bool`

HasTokenTimeoutSeconds returns a boolean if a field has been set.

### GetUserName

`func (o *AuthRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AuthRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AuthRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *AuthRequest) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


