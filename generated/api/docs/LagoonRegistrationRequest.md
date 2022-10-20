# LagoonRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | Pointer to **string** |  | [optional] 
**UserPassword** | Pointer to **string** | This password is used by the initial user&#39;s login. | [optional] 

## Methods

### NewLagoonRegistrationRequest

`func NewLagoonRegistrationRequest() *LagoonRegistrationRequest`

NewLagoonRegistrationRequest instantiates a new LagoonRegistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLagoonRegistrationRequestWithDefaults

`func NewLagoonRegistrationRequestWithDefaults() *LagoonRegistrationRequest`

NewLagoonRegistrationRequestWithDefaults instantiates a new LagoonRegistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlan

`func (o *LagoonRegistrationRequest) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *LagoonRegistrationRequest) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *LagoonRegistrationRequest) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *LagoonRegistrationRequest) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetUserPassword

`func (o *LagoonRegistrationRequest) GetUserPassword() string`

GetUserPassword returns the UserPassword field if non-nil, zero value otherwise.

### GetUserPasswordOk

`func (o *LagoonRegistrationRequest) GetUserPasswordOk() (*string, bool)`

GetUserPasswordOk returns a tuple with the UserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPassword

`func (o *LagoonRegistrationRequest) SetUserPassword(v string)`

SetUserPassword sets UserPassword field to given value.

### HasUserPassword

`func (o *LagoonRegistrationRequest) HasUserPassword() bool`

HasUserPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


