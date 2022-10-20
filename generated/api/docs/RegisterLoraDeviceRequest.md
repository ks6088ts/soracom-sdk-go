# RegisterLoraDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** |  | [optional] 
**RegistrationSecret** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewRegisterLoraDeviceRequest

`func NewRegisterLoraDeviceRequest() *RegisterLoraDeviceRequest`

NewRegisterLoraDeviceRequest instantiates a new RegisterLoraDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterLoraDeviceRequestWithDefaults

`func NewRegisterLoraDeviceRequestWithDefaults() *RegisterLoraDeviceRequest`

NewRegisterLoraDeviceRequestWithDefaults instantiates a new RegisterLoraDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *RegisterLoraDeviceRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *RegisterLoraDeviceRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *RegisterLoraDeviceRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *RegisterLoraDeviceRequest) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetRegistrationSecret

`func (o *RegisterLoraDeviceRequest) GetRegistrationSecret() string`

GetRegistrationSecret returns the RegistrationSecret field if non-nil, zero value otherwise.

### GetRegistrationSecretOk

`func (o *RegisterLoraDeviceRequest) GetRegistrationSecretOk() (*string, bool)`

GetRegistrationSecretOk returns a tuple with the RegistrationSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationSecret

`func (o *RegisterLoraDeviceRequest) SetRegistrationSecret(v string)`

SetRegistrationSecret sets RegistrationSecret field to given value.

### HasRegistrationSecret

`func (o *RegisterLoraDeviceRequest) HasRegistrationSecret() bool`

HasRegistrationSecret returns a boolean if a field has been set.

### GetTags

`func (o *RegisterLoraDeviceRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RegisterLoraDeviceRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RegisterLoraDeviceRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RegisterLoraDeviceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


