# ExportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **int64** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**To** | Pointer to **int64** |  | [optional] 

## Methods

### NewExportRequest

`func NewExportRequest() *ExportRequest`

NewExportRequest instantiates a new ExportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportRequestWithDefaults

`func NewExportRequestWithDefaults() *ExportRequest`

NewExportRequestWithDefaults instantiates a new ExportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *ExportRequest) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ExportRequest) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ExportRequest) SetFrom(v int64)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ExportRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetPeriod

`func (o *ExportRequest) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ExportRequest) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ExportRequest) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *ExportRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetTo

`func (o *ExportRequest) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ExportRequest) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ExportRequest) SetTo(v int64)`

SetTo sets To field to given value.

### HasTo

`func (o *ExportRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


