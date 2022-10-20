# VerifyPasswordResetTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** |  | 
**Token** | **string** |  | 

## Methods

### NewVerifyPasswordResetTokenRequest

`func NewVerifyPasswordResetTokenRequest(password string, token string, ) *VerifyPasswordResetTokenRequest`

NewVerifyPasswordResetTokenRequest instantiates a new VerifyPasswordResetTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyPasswordResetTokenRequestWithDefaults

`func NewVerifyPasswordResetTokenRequestWithDefaults() *VerifyPasswordResetTokenRequest`

NewVerifyPasswordResetTokenRequestWithDefaults instantiates a new VerifyPasswordResetTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *VerifyPasswordResetTokenRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VerifyPasswordResetTokenRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VerifyPasswordResetTokenRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetToken

`func (o *VerifyPasswordResetTokenRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *VerifyPasswordResetTokenRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *VerifyPasswordResetTokenRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


