# OpenGateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivacySeparatorEnabled** | Pointer to **bool** |  | [optional] [default to false]
**VxlanId** | Pointer to **int32** |  | [optional] [default to 10]

## Methods

### NewOpenGateRequest

`func NewOpenGateRequest() *OpenGateRequest`

NewOpenGateRequest instantiates a new OpenGateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenGateRequestWithDefaults

`func NewOpenGateRequestWithDefaults() *OpenGateRequest`

NewOpenGateRequestWithDefaults instantiates a new OpenGateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivacySeparatorEnabled

`func (o *OpenGateRequest) GetPrivacySeparatorEnabled() bool`

GetPrivacySeparatorEnabled returns the PrivacySeparatorEnabled field if non-nil, zero value otherwise.

### GetPrivacySeparatorEnabledOk

`func (o *OpenGateRequest) GetPrivacySeparatorEnabledOk() (*bool, bool)`

GetPrivacySeparatorEnabledOk returns a tuple with the PrivacySeparatorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacySeparatorEnabled

`func (o *OpenGateRequest) SetPrivacySeparatorEnabled(v bool)`

SetPrivacySeparatorEnabled sets PrivacySeparatorEnabled field to given value.

### HasPrivacySeparatorEnabled

`func (o *OpenGateRequest) HasPrivacySeparatorEnabled() bool`

HasPrivacySeparatorEnabled returns a boolean if a field has been set.

### GetVxlanId

`func (o *OpenGateRequest) GetVxlanId() int32`

GetVxlanId returns the VxlanId field if non-nil, zero value otherwise.

### GetVxlanIdOk

`func (o *OpenGateRequest) GetVxlanIdOk() (*int32, bool)`

GetVxlanIdOk returns a tuple with the VxlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlanId

`func (o *OpenGateRequest) SetVxlanId(v int32)`

SetVxlanId sets VxlanId field to given value.

### HasVxlanId

`func (o *OpenGateRequest) HasVxlanId() bool`

HasVxlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


