# RegisterGatePeerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InnerIpAddress** | Pointer to **string** |  | [optional] 
**OuterIpAddress** | **string** |  | 

## Methods

### NewRegisterGatePeerRequest

`func NewRegisterGatePeerRequest(outerIpAddress string, ) *RegisterGatePeerRequest`

NewRegisterGatePeerRequest instantiates a new RegisterGatePeerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterGatePeerRequestWithDefaults

`func NewRegisterGatePeerRequestWithDefaults() *RegisterGatePeerRequest`

NewRegisterGatePeerRequestWithDefaults instantiates a new RegisterGatePeerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInnerIpAddress

`func (o *RegisterGatePeerRequest) GetInnerIpAddress() string`

GetInnerIpAddress returns the InnerIpAddress field if non-nil, zero value otherwise.

### GetInnerIpAddressOk

`func (o *RegisterGatePeerRequest) GetInnerIpAddressOk() (*string, bool)`

GetInnerIpAddressOk returns a tuple with the InnerIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerIpAddress

`func (o *RegisterGatePeerRequest) SetInnerIpAddress(v string)`

SetInnerIpAddress sets InnerIpAddress field to given value.

### HasInnerIpAddress

`func (o *RegisterGatePeerRequest) HasInnerIpAddress() bool`

HasInnerIpAddress returns a boolean if a field has been set.

### GetOuterIpAddress

`func (o *RegisterGatePeerRequest) GetOuterIpAddress() string`

GetOuterIpAddress returns the OuterIpAddress field if non-nil, zero value otherwise.

### GetOuterIpAddressOk

`func (o *RegisterGatePeerRequest) GetOuterIpAddressOk() (*string, bool)`

GetOuterIpAddressOk returns a tuple with the OuterIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterIpAddress

`func (o *RegisterGatePeerRequest) SetOuterIpAddress(v string)`

SetOuterIpAddress sets OuterIpAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


