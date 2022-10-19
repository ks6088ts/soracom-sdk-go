# PortMappingSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpRanges** | Pointer to **[]string** | The range of IP addresses in CIDR notation which will be allowed to access the port mapping. | [optional] 

## Methods

### NewPortMappingSource

`func NewPortMappingSource() *PortMappingSource`

NewPortMappingSource instantiates a new PortMappingSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortMappingSourceWithDefaults

`func NewPortMappingSourceWithDefaults() *PortMappingSource`

NewPortMappingSourceWithDefaults instantiates a new PortMappingSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpRanges

`func (o *PortMappingSource) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *PortMappingSource) GetIpRangesOk() (*[]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *PortMappingSource) SetIpRanges(v []string)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *PortMappingSource) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


