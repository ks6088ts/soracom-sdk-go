# RegisterOperatorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Password** | **string** | パスワードは以下の条件を満たしている必要があります：長さ 8 文字以上 100 文字以内、アルファベット小文字 (a-z) を 1 文字以上使用、アルファベット大文字 (A-Z) を 1 文字以上使用、数字を 1 文字以上使用。記号なども使用できます。 | 

## Methods

### NewRegisterOperatorsRequest

`func NewRegisterOperatorsRequest(email string, password string, ) *RegisterOperatorsRequest`

NewRegisterOperatorsRequest instantiates a new RegisterOperatorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterOperatorsRequestWithDefaults

`func NewRegisterOperatorsRequestWithDefaults() *RegisterOperatorsRequest`

NewRegisterOperatorsRequestWithDefaults instantiates a new RegisterOperatorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *RegisterOperatorsRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RegisterOperatorsRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RegisterOperatorsRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *RegisterOperatorsRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegisterOperatorsRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegisterOperatorsRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


