# RegisterSimRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** |  | [optional] 
**RegistrationSecret** | **string** | PUK or PASSCODE on SIM card | 
**Tags** | Pointer to **map[string]string** | An object which always contains at least one property \&quot;name\&quot; with a string value. If you give a subscriber/SIM a name, the name will be returned as the value of the \&quot;name\&quot; property. If the subscriber/SIM does not have a name, an empty string \&quot;\&quot; is returned. In addition, if you create any custom tags for the subscriber/SIM, each custom tag will appear as additional properties in the object. | [optional] 

## Methods

### NewRegisterSimRequest

`func NewRegisterSimRequest(registrationSecret string, ) *RegisterSimRequest`

NewRegisterSimRequest instantiates a new RegisterSimRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterSimRequestWithDefaults

`func NewRegisterSimRequestWithDefaults() *RegisterSimRequest`

NewRegisterSimRequestWithDefaults instantiates a new RegisterSimRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *RegisterSimRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *RegisterSimRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *RegisterSimRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *RegisterSimRequest) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetRegistrationSecret

`func (o *RegisterSimRequest) GetRegistrationSecret() string`

GetRegistrationSecret returns the RegistrationSecret field if non-nil, zero value otherwise.

### GetRegistrationSecretOk

`func (o *RegisterSimRequest) GetRegistrationSecretOk() (*string, bool)`

GetRegistrationSecretOk returns a tuple with the RegistrationSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationSecret

`func (o *RegisterSimRequest) SetRegistrationSecret(v string)`

SetRegistrationSecret sets RegistrationSecret field to given value.


### GetTags

`func (o *RegisterSimRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RegisterSimRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RegisterSimRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RegisterSimRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


