# ExecuteSoraletRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** |  | 
**Direction** | **string** |  | 
**EncodingType** | Pointer to **string** |  | [optional] 
**Payload** | **string** |  | 
**Source** | [**map[string]SoraletDataSource**](SoraletDataSource.md) |  | 
**Userdata** | Pointer to **string** |  | [optional] 
**Version** | **string** |  | 

## Methods

### NewExecuteSoraletRequest

`func NewExecuteSoraletRequest(contentType string, direction string, payload string, source map[string]SoraletDataSource, version string, ) *ExecuteSoraletRequest`

NewExecuteSoraletRequest instantiates a new ExecuteSoraletRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteSoraletRequestWithDefaults

`func NewExecuteSoraletRequestWithDefaults() *ExecuteSoraletRequest`

NewExecuteSoraletRequestWithDefaults instantiates a new ExecuteSoraletRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *ExecuteSoraletRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ExecuteSoraletRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ExecuteSoraletRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetDirection

`func (o *ExecuteSoraletRequest) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ExecuteSoraletRequest) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ExecuteSoraletRequest) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetEncodingType

`func (o *ExecuteSoraletRequest) GetEncodingType() string`

GetEncodingType returns the EncodingType field if non-nil, zero value otherwise.

### GetEncodingTypeOk

`func (o *ExecuteSoraletRequest) GetEncodingTypeOk() (*string, bool)`

GetEncodingTypeOk returns a tuple with the EncodingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingType

`func (o *ExecuteSoraletRequest) SetEncodingType(v string)`

SetEncodingType sets EncodingType field to given value.

### HasEncodingType

`func (o *ExecuteSoraletRequest) HasEncodingType() bool`

HasEncodingType returns a boolean if a field has been set.

### GetPayload

`func (o *ExecuteSoraletRequest) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ExecuteSoraletRequest) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ExecuteSoraletRequest) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetSource

`func (o *ExecuteSoraletRequest) GetSource() map[string]SoraletDataSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ExecuteSoraletRequest) GetSourceOk() (*map[string]SoraletDataSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ExecuteSoraletRequest) SetSource(v map[string]SoraletDataSource)`

SetSource sets Source field to given value.


### GetUserdata

`func (o *ExecuteSoraletRequest) GetUserdata() string`

GetUserdata returns the Userdata field if non-nil, zero value otherwise.

### GetUserdataOk

`func (o *ExecuteSoraletRequest) GetUserdataOk() (*string, bool)`

GetUserdataOk returns a tuple with the Userdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdata

`func (o *ExecuteSoraletRequest) SetUserdata(v string)`

SetUserdata sets Userdata field to given value.

### HasUserdata

`func (o *ExecuteSoraletRequest) HasUserdata() bool`

HasUserdata returns a boolean if a field has been set.

### GetVersion

`func (o *ExecuteSoraletRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExecuteSoraletRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExecuteSoraletRequest) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


