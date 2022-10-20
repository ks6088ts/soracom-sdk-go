# AuthKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthKeyId** | Pointer to **string** |  | [optional] 
**CreateDateTime** | Pointer to **int64** |  | [optional] 
**LastUsedDateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewAuthKeyResponse

`func NewAuthKeyResponse() *AuthKeyResponse`

NewAuthKeyResponse instantiates a new AuthKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthKeyResponseWithDefaults

`func NewAuthKeyResponseWithDefaults() *AuthKeyResponse`

NewAuthKeyResponseWithDefaults instantiates a new AuthKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthKeyId

`func (o *AuthKeyResponse) GetAuthKeyId() string`

GetAuthKeyId returns the AuthKeyId field if non-nil, zero value otherwise.

### GetAuthKeyIdOk

`func (o *AuthKeyResponse) GetAuthKeyIdOk() (*string, bool)`

GetAuthKeyIdOk returns a tuple with the AuthKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKeyId

`func (o *AuthKeyResponse) SetAuthKeyId(v string)`

SetAuthKeyId sets AuthKeyId field to given value.

### HasAuthKeyId

`func (o *AuthKeyResponse) HasAuthKeyId() bool`

HasAuthKeyId returns a boolean if a field has been set.

### GetCreateDateTime

`func (o *AuthKeyResponse) GetCreateDateTime() int64`

GetCreateDateTime returns the CreateDateTime field if non-nil, zero value otherwise.

### GetCreateDateTimeOk

`func (o *AuthKeyResponse) GetCreateDateTimeOk() (*int64, bool)`

GetCreateDateTimeOk returns a tuple with the CreateDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDateTime

`func (o *AuthKeyResponse) SetCreateDateTime(v int64)`

SetCreateDateTime sets CreateDateTime field to given value.

### HasCreateDateTime

`func (o *AuthKeyResponse) HasCreateDateTime() bool`

HasCreateDateTime returns a boolean if a field has been set.

### GetLastUsedDateTime

`func (o *AuthKeyResponse) GetLastUsedDateTime() int64`

GetLastUsedDateTime returns the LastUsedDateTime field if non-nil, zero value otherwise.

### GetLastUsedDateTimeOk

`func (o *AuthKeyResponse) GetLastUsedDateTimeOk() (*int64, bool)`

GetLastUsedDateTimeOk returns a tuple with the LastUsedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedDateTime

`func (o *AuthKeyResponse) SetLastUsedDateTime(v int64)`

SetLastUsedDateTime sets LastUsedDateTime field to given value.

### HasLastUsedDateTime

`func (o *AuthKeyResponse) HasLastUsedDateTime() bool`

HasLastUsedDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


