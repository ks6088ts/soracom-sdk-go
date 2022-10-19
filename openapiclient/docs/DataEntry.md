# DataEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewDataEntry

`func NewDataEntry() *DataEntry`

NewDataEntry instantiates a new DataEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataEntryWithDefaults

`func NewDataEntryWithDefaults() *DataEntry`

NewDataEntryWithDefaults instantiates a new DataEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *DataEntry) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DataEntry) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DataEntry) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *DataEntry) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentType

`func (o *DataEntry) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *DataEntry) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *DataEntry) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *DataEntry) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetTime

`func (o *DataEntry) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DataEntry) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DataEntry) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *DataEntry) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


