# PortMappingDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Imsi** | **string** | The target IMSI of the subscriber. | 
**Port** | **float32** | The port on your device used for access. | 

## Methods

### NewPortMappingDestination

`func NewPortMappingDestination(imsi string, port float32, ) *PortMappingDestination`

NewPortMappingDestination instantiates a new PortMappingDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortMappingDestinationWithDefaults

`func NewPortMappingDestinationWithDefaults() *PortMappingDestination`

NewPortMappingDestinationWithDefaults instantiates a new PortMappingDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImsi

`func (o *PortMappingDestination) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *PortMappingDestination) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *PortMappingDestination) SetImsi(v string)`

SetImsi sets Imsi field to given value.


### GetPort

`func (o *PortMappingDestination) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PortMappingDestination) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PortMappingDestination) SetPort(v float32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


