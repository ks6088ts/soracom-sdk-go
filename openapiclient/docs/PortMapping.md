# PortMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to [**PortMappingDestination**](PortMappingDestination.md) |  | [optional] 
**Duration** | Pointer to **float32** | The duration that remote access is enabled, in seconds. | [optional] 
**Endpoint** | Pointer to **string** | SORACOM Napter endpoint (IP address and port number) for remote access. | [optional] 
**Hostname** | Pointer to **string** | SORACOM Napter hostname for remote access. | [optional] 
**IpAddress** | Pointer to **string** | SORACOM Napter IP Address for remote access. | [optional] 
**Port** | Pointer to **float32** | SORACOM Napter port number for remote access. | [optional] 
**Source** | Pointer to [**PortMappingSource**](PortMappingSource.md) |  | [optional] 
**TlsRequired** | Pointer to **bool** | Indicates TLS is required. | [optional] 

## Methods

### NewPortMapping

`func NewPortMapping() *PortMapping`

NewPortMapping instantiates a new PortMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortMappingWithDefaults

`func NewPortMappingWithDefaults() *PortMapping`

NewPortMappingWithDefaults instantiates a new PortMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *PortMapping) GetDestination() PortMappingDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PortMapping) GetDestinationOk() (*PortMappingDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PortMapping) SetDestination(v PortMappingDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *PortMapping) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDuration

`func (o *PortMapping) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PortMapping) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PortMapping) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *PortMapping) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEndpoint

`func (o *PortMapping) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *PortMapping) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *PortMapping) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *PortMapping) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetHostname

`func (o *PortMapping) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *PortMapping) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *PortMapping) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *PortMapping) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIpAddress

`func (o *PortMapping) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *PortMapping) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *PortMapping) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *PortMapping) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetPort

`func (o *PortMapping) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PortMapping) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PortMapping) SetPort(v float32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PortMapping) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSource

`func (o *PortMapping) GetSource() PortMappingSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PortMapping) GetSourceOk() (*PortMappingSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PortMapping) SetSource(v PortMappingSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *PortMapping) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTlsRequired

`func (o *PortMapping) GetTlsRequired() bool`

GetTlsRequired returns the TlsRequired field if non-nil, zero value otherwise.

### GetTlsRequiredOk

`func (o *PortMapping) GetTlsRequiredOk() (*bool, bool)`

GetTlsRequiredOk returns a tuple with the TlsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsRequired

`func (o *PortMapping) SetTlsRequired(v bool)`

SetTlsRequired sets TlsRequired field to given value.

### HasTlsRequired

`func (o *PortMapping) HasTlsRequired() bool`

HasTlsRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


