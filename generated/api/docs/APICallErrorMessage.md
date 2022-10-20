# APICallErrorMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | エラーコード | 
**Message** | **string** | エラーメッセージ。リクエスト時に X-Soracom-Lang ヘッダーに言語(en,ja)を設定するとその言語のメッセージがセットされます。 | 

## Methods

### NewAPICallErrorMessage

`func NewAPICallErrorMessage(code string, message string, ) *APICallErrorMessage`

NewAPICallErrorMessage instantiates a new APICallErrorMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPICallErrorMessageWithDefaults

`func NewAPICallErrorMessageWithDefaults() *APICallErrorMessage`

NewAPICallErrorMessageWithDefaults instantiates a new APICallErrorMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *APICallErrorMessage) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *APICallErrorMessage) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *APICallErrorMessage) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *APICallErrorMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *APICallErrorMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *APICallErrorMessage) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


