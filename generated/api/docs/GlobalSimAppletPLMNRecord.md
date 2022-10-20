# GlobalSimAppletPLMNRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerId** | **int64** |  | 
**Mcc** | **string** |  | 
**Mnc** | Pointer to **string** |  | [optional] 

## Methods

### NewGlobalSimAppletPLMNRecord

`func NewGlobalSimAppletPLMNRecord(containerId int64, mcc string, ) *GlobalSimAppletPLMNRecord`

NewGlobalSimAppletPLMNRecord instantiates a new GlobalSimAppletPLMNRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalSimAppletPLMNRecordWithDefaults

`func NewGlobalSimAppletPLMNRecordWithDefaults() *GlobalSimAppletPLMNRecord`

NewGlobalSimAppletPLMNRecordWithDefaults instantiates a new GlobalSimAppletPLMNRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerId

`func (o *GlobalSimAppletPLMNRecord) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *GlobalSimAppletPLMNRecord) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *GlobalSimAppletPLMNRecord) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.


### GetMcc

`func (o *GlobalSimAppletPLMNRecord) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *GlobalSimAppletPLMNRecord) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *GlobalSimAppletPLMNRecord) SetMcc(v string)`

SetMcc sets Mcc field to given value.


### GetMnc

`func (o *GlobalSimAppletPLMNRecord) GetMnc() string`

GetMnc returns the Mnc field if non-nil, zero value otherwise.

### GetMncOk

`func (o *GlobalSimAppletPLMNRecord) GetMncOk() (*string, bool)`

GetMncOk returns a tuple with the Mnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnc

`func (o *GlobalSimAppletPLMNRecord) SetMnc(v string)`

SetMnc sets Mnc field to given value.

### HasMnc

`func (o *GlobalSimAppletPLMNRecord) HasMnc() bool`

HasMnc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


