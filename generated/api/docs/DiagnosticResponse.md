# DiagnosticResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiagnosticId** | Pointer to **string** |  | [optional] 
**From** | Pointer to **int64** | 診断対象期間の始まり (UNIX 時間 (ミリ秒)) | [optional] 
**Insights** | Pointer to [**[]map[string]Insight**](map[string]Insight.md) |  | [optional] 
**ResourceId** | Pointer to **string** | resourceType に応じた識別子 | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**To** | Pointer to **int64** | 診断対象期間の終わり (UNIX 時間 (ミリ秒)) | [optional] 

## Methods

### NewDiagnosticResponse

`func NewDiagnosticResponse() *DiagnosticResponse`

NewDiagnosticResponse instantiates a new DiagnosticResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosticResponseWithDefaults

`func NewDiagnosticResponseWithDefaults() *DiagnosticResponse`

NewDiagnosticResponseWithDefaults instantiates a new DiagnosticResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiagnosticId

`func (o *DiagnosticResponse) GetDiagnosticId() string`

GetDiagnosticId returns the DiagnosticId field if non-nil, zero value otherwise.

### GetDiagnosticIdOk

`func (o *DiagnosticResponse) GetDiagnosticIdOk() (*string, bool)`

GetDiagnosticIdOk returns a tuple with the DiagnosticId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticId

`func (o *DiagnosticResponse) SetDiagnosticId(v string)`

SetDiagnosticId sets DiagnosticId field to given value.

### HasDiagnosticId

`func (o *DiagnosticResponse) HasDiagnosticId() bool`

HasDiagnosticId returns a boolean if a field has been set.

### GetFrom

`func (o *DiagnosticResponse) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *DiagnosticResponse) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *DiagnosticResponse) SetFrom(v int64)`

SetFrom sets From field to given value.

### HasFrom

`func (o *DiagnosticResponse) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetInsights

`func (o *DiagnosticResponse) GetInsights() []map[string]Insight`

GetInsights returns the Insights field if non-nil, zero value otherwise.

### GetInsightsOk

`func (o *DiagnosticResponse) GetInsightsOk() (*[]map[string]Insight, bool)`

GetInsightsOk returns a tuple with the Insights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsights

`func (o *DiagnosticResponse) SetInsights(v []map[string]Insight)`

SetInsights sets Insights field to given value.

### HasInsights

`func (o *DiagnosticResponse) HasInsights() bool`

HasInsights returns a boolean if a field has been set.

### GetResourceId

`func (o *DiagnosticResponse) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *DiagnosticResponse) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *DiagnosticResponse) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *DiagnosticResponse) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *DiagnosticResponse) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *DiagnosticResponse) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *DiagnosticResponse) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *DiagnosticResponse) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetService

`func (o *DiagnosticResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *DiagnosticResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *DiagnosticResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *DiagnosticResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *DiagnosticResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DiagnosticResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DiagnosticResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DiagnosticResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTo

`func (o *DiagnosticResponse) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *DiagnosticResponse) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *DiagnosticResponse) SetTo(v int64)`

SetTo sets To field to given value.

### HasTo

`func (o *DiagnosticResponse) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


