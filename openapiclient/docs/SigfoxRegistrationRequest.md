# SigfoxRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistrationSecret** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSigfoxRegistrationRequest

`func NewSigfoxRegistrationRequest() *SigfoxRegistrationRequest`

NewSigfoxRegistrationRequest instantiates a new SigfoxRegistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigfoxRegistrationRequestWithDefaults

`func NewSigfoxRegistrationRequestWithDefaults() *SigfoxRegistrationRequest`

NewSigfoxRegistrationRequestWithDefaults instantiates a new SigfoxRegistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistrationSecret

`func (o *SigfoxRegistrationRequest) GetRegistrationSecret() string`

GetRegistrationSecret returns the RegistrationSecret field if non-nil, zero value otherwise.

### GetRegistrationSecretOk

`func (o *SigfoxRegistrationRequest) GetRegistrationSecretOk() (*string, bool)`

GetRegistrationSecretOk returns a tuple with the RegistrationSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationSecret

`func (o *SigfoxRegistrationRequest) SetRegistrationSecret(v string)`

SetRegistrationSecret sets RegistrationSecret field to given value.

### HasRegistrationSecret

`func (o *SigfoxRegistrationRequest) HasRegistrationSecret() bool`

HasRegistrationSecret returns a boolean if a field has been set.

### GetTags

`func (o *SigfoxRegistrationRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SigfoxRegistrationRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SigfoxRegistrationRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SigfoxRegistrationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


