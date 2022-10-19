# CreateVirtualPrivateGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceSubnetCidrRange** | Pointer to **string** |  | [optional] [default to "10.128.0.0/9"]
**Type** | **int32** | VPG のタイプ  - &#x60;14&#x60; : Type-E  - &#x60;15&#x60; : Type-F  | 
**UseInternetGateway** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewCreateVirtualPrivateGatewayRequest

`func NewCreateVirtualPrivateGatewayRequest(type_ int32, ) *CreateVirtualPrivateGatewayRequest`

NewCreateVirtualPrivateGatewayRequest instantiates a new CreateVirtualPrivateGatewayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVirtualPrivateGatewayRequestWithDefaults

`func NewCreateVirtualPrivateGatewayRequestWithDefaults() *CreateVirtualPrivateGatewayRequest`

NewCreateVirtualPrivateGatewayRequestWithDefaults instantiates a new CreateVirtualPrivateGatewayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceSubnetCidrRange

`func (o *CreateVirtualPrivateGatewayRequest) GetDeviceSubnetCidrRange() string`

GetDeviceSubnetCidrRange returns the DeviceSubnetCidrRange field if non-nil, zero value otherwise.

### GetDeviceSubnetCidrRangeOk

`func (o *CreateVirtualPrivateGatewayRequest) GetDeviceSubnetCidrRangeOk() (*string, bool)`

GetDeviceSubnetCidrRangeOk returns a tuple with the DeviceSubnetCidrRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSubnetCidrRange

`func (o *CreateVirtualPrivateGatewayRequest) SetDeviceSubnetCidrRange(v string)`

SetDeviceSubnetCidrRange sets DeviceSubnetCidrRange field to given value.

### HasDeviceSubnetCidrRange

`func (o *CreateVirtualPrivateGatewayRequest) HasDeviceSubnetCidrRange() bool`

HasDeviceSubnetCidrRange returns a boolean if a field has been set.

### GetType

`func (o *CreateVirtualPrivateGatewayRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateVirtualPrivateGatewayRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateVirtualPrivateGatewayRequest) SetType(v int32)`

SetType sets Type field to given value.


### GetUseInternetGateway

`func (o *CreateVirtualPrivateGatewayRequest) GetUseInternetGateway() bool`

GetUseInternetGateway returns the UseInternetGateway field if non-nil, zero value otherwise.

### GetUseInternetGatewayOk

`func (o *CreateVirtualPrivateGatewayRequest) GetUseInternetGatewayOk() (*bool, bool)`

GetUseInternetGatewayOk returns a tuple with the UseInternetGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseInternetGateway

`func (o *CreateVirtualPrivateGatewayRequest) SetUseInternetGateway(v bool)`

SetUseInternetGateway sets UseInternetGateway field to given value.

### HasUseInternetGateway

`func (o *CreateVirtualPrivateGatewayRequest) HasUseInternetGateway() bool`

HasUseInternetGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


