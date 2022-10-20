# LoraData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** |  | [optional] 
**FPort** | Pointer to **int32** |  | [optional] [default to 2]

## Methods

### NewLoraData

`func NewLoraData() *LoraData`

NewLoraData instantiates a new LoraData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoraDataWithDefaults

`func NewLoraDataWithDefaults() *LoraData`

NewLoraDataWithDefaults instantiates a new LoraData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *LoraData) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LoraData) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LoraData) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *LoraData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetFPort

`func (o *LoraData) GetFPort() int32`

GetFPort returns the FPort field if non-nil, zero value otherwise.

### GetFPortOk

`func (o *LoraData) GetFPortOk() (*int32, bool)`

GetFPortOk returns a tuple with the FPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFPort

`func (o *LoraData) SetFPort(v int32)`

SetFPort sets FPort field to given value.

### HasFPort

`func (o *LoraData) HasFPort() bool`

HasFPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


