# GenerateTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenTimeoutSeconds** | Pointer to **int32** | 新しい API トークンの有効期限の長さ（秒単位）。 指定しなければデフォルトは 86400 [秒]（24 時間）。 最大値は 172800 [秒]（48 時間）。  | [optional] [default to 86400]

## Methods

### NewGenerateTokenRequest

`func NewGenerateTokenRequest() *GenerateTokenRequest`

NewGenerateTokenRequest instantiates a new GenerateTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateTokenRequestWithDefaults

`func NewGenerateTokenRequestWithDefaults() *GenerateTokenRequest`

NewGenerateTokenRequestWithDefaults instantiates a new GenerateTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenTimeoutSeconds

`func (o *GenerateTokenRequest) GetTokenTimeoutSeconds() int32`

GetTokenTimeoutSeconds returns the TokenTimeoutSeconds field if non-nil, zero value otherwise.

### GetTokenTimeoutSecondsOk

`func (o *GenerateTokenRequest) GetTokenTimeoutSecondsOk() (*int32, bool)`

GetTokenTimeoutSecondsOk returns a tuple with the TokenTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTimeoutSeconds

`func (o *GenerateTokenRequest) SetTokenTimeoutSeconds(v int32)`

SetTokenTimeoutSeconds sets TokenTimeoutSeconds field to given value.

### HasTokenTimeoutSeconds

`func (o *GenerateTokenRequest) HasTokenTimeoutSeconds() bool`

HasTokenTimeoutSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


