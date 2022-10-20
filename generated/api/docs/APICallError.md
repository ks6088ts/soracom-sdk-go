# APICallError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to [**APICallErrorMessage**](APICallErrorMessage.md) |  | [optional] 
**HttpStatus** | Pointer to **int32** |  | [optional] 

## Methods

### NewAPICallError

`func NewAPICallError() *APICallError`

NewAPICallError instantiates a new APICallError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPICallErrorWithDefaults

`func NewAPICallErrorWithDefaults() *APICallError`

NewAPICallErrorWithDefaults instantiates a new APICallError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *APICallError) GetErrorMessage() APICallErrorMessage`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *APICallError) GetErrorMessageOk() (*APICallErrorMessage, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *APICallError) SetErrorMessage(v APICallErrorMessage)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *APICallError) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetHttpStatus

`func (o *APICallError) GetHttpStatus() int32`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *APICallError) GetHttpStatusOk() (*int32, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *APICallError) SetHttpStatus(v int32)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *APICallError) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


