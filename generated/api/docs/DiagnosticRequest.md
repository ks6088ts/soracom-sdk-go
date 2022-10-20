# DiagnosticRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **int64** | 診断対象期間の始まり (UNIX 時間 (ミリ秒)) | [optional] 
**ResourceId** | **string** | resourceType に応じた識別子  - resourceType が &#x60;sim&#x60; の場合は、SIM ID を指定します。  | 
**ResourceType** | **string** |  | 
**Service** | **string** |  | 
**To** | Pointer to **int64** | 診断対象期間の終わり (UNIX 時間 (ミリ秒)) | [optional] 

## Methods

### NewDiagnosticRequest

`func NewDiagnosticRequest(resourceId string, resourceType string, service string, ) *DiagnosticRequest`

NewDiagnosticRequest instantiates a new DiagnosticRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosticRequestWithDefaults

`func NewDiagnosticRequestWithDefaults() *DiagnosticRequest`

NewDiagnosticRequestWithDefaults instantiates a new DiagnosticRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *DiagnosticRequest) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *DiagnosticRequest) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *DiagnosticRequest) SetFrom(v int64)`

SetFrom sets From field to given value.

### HasFrom

`func (o *DiagnosticRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetResourceId

`func (o *DiagnosticRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *DiagnosticRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *DiagnosticRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *DiagnosticRequest) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *DiagnosticRequest) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *DiagnosticRequest) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetService

`func (o *DiagnosticRequest) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *DiagnosticRequest) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *DiagnosticRequest) SetService(v string)`

SetService sets Service field to given value.


### GetTo

`func (o *DiagnosticRequest) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *DiagnosticRequest) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *DiagnosticRequest) SetTo(v int64)`

SetTo sets To field to given value.

### HasTo

`func (o *DiagnosticRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


