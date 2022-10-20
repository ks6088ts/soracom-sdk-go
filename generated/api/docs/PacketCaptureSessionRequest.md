# PacketCaptureSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** |  | 
**Prefix** | Pointer to **string** |  | [optional] 

## Methods

### NewPacketCaptureSessionRequest

`func NewPacketCaptureSessionRequest(duration int32, ) *PacketCaptureSessionRequest`

NewPacketCaptureSessionRequest instantiates a new PacketCaptureSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPacketCaptureSessionRequestWithDefaults

`func NewPacketCaptureSessionRequestWithDefaults() *PacketCaptureSessionRequest`

NewPacketCaptureSessionRequestWithDefaults instantiates a new PacketCaptureSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *PacketCaptureSessionRequest) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PacketCaptureSessionRequest) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PacketCaptureSessionRequest) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetPrefix

`func (o *PacketCaptureSessionRequest) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PacketCaptureSessionRequest) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PacketCaptureSessionRequest) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PacketCaptureSessionRequest) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


