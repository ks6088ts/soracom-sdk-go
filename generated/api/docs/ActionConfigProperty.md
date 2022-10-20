# ActionConfigProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | InvokeAWSLambdaAction の時のみ有効 | [optional] 
**Body** | Pointer to **string** | ExecuteWebRequestAction の時のみ有効(任意) | [optional] 
**ContentType** | Pointer to **string** | ExecuteWebRequestAction の時のみ有効 | [optional] 
**Endpoint** | Pointer to **string** | InvokeAWSLambdaAction の時のみ有効 | [optional] 
**ExecutionDateTimeConst** | **string** | executionDateTimeConst で指定したタイミングから指定分後に実行 | 
**ExecutionOffsetMinutes** | Pointer to **string** | Run in the minutes after the timing of executionDateTimeConst | [optional] 
**FunctionName** | Pointer to **string** | InvokeAWSLambdaAction の時のみ有効 | [optional] 
**Headers** | Pointer to **map[string]interface{}** | ExecuteWebRequestAction の時のみ有効(任意) | [optional] 
**HttpMethod** | Pointer to **string** | ExecuteWebRequestAction の時のみ有効 | [optional] 
**Message** | Pointer to **string** | SendMailAction, SendMailToOperatorAction の時のみ有効 | [optional] 
**Parameter1** | Pointer to **string** | InvokeAWSLambdaAction の時のみ有効 | [optional] 
**Parameter2** | Pointer to **string** | InvokeAWSLambdaAction の時のみ有効 | [optional] 
**Parameter3** | Pointer to **string** | InvokeAWSLambdaAction の時のみ有効 | [optional] 
**Parameter4** | Pointer to **string** | InvokeAWSLambdaAction の時のみ有効 | [optional] 
**Parameter5** | Pointer to **string** | InvokeAWSLambdaAction の時のみ有効 | [optional] 
**SecretAccessKey** | Pointer to **string** | InvokeAWSLambdaAction の時のみ有効 | [optional] 
**SpeedClass** | Pointer to **string** | ChangeSpeedClassAction の時のみ有効 | [optional] 
**Title** | Pointer to **string** | SendMailAction, SendMailToOperatorAction の時のみ有効 | [optional] 
**To** | Pointer to **string** | SendMailAction の時のみ有効 | [optional] 
**Url** | Pointer to **string** | 接続先 URL とパラメーター ExecuteWebRequestAction の時のみ有効 | [optional] 

## Methods

### NewActionConfigProperty

`func NewActionConfigProperty(executionDateTimeConst string, ) *ActionConfigProperty`

NewActionConfigProperty instantiates a new ActionConfigProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionConfigPropertyWithDefaults

`func NewActionConfigPropertyWithDefaults() *ActionConfigProperty`

NewActionConfigPropertyWithDefaults instantiates a new ActionConfigProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *ActionConfigProperty) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *ActionConfigProperty) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *ActionConfigProperty) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *ActionConfigProperty) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetBody

`func (o *ActionConfigProperty) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ActionConfigProperty) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ActionConfigProperty) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ActionConfigProperty) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetContentType

`func (o *ActionConfigProperty) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ActionConfigProperty) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ActionConfigProperty) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ActionConfigProperty) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetEndpoint

`func (o *ActionConfigProperty) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ActionConfigProperty) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ActionConfigProperty) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ActionConfigProperty) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetExecutionDateTimeConst

`func (o *ActionConfigProperty) GetExecutionDateTimeConst() string`

GetExecutionDateTimeConst returns the ExecutionDateTimeConst field if non-nil, zero value otherwise.

### GetExecutionDateTimeConstOk

`func (o *ActionConfigProperty) GetExecutionDateTimeConstOk() (*string, bool)`

GetExecutionDateTimeConstOk returns a tuple with the ExecutionDateTimeConst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDateTimeConst

`func (o *ActionConfigProperty) SetExecutionDateTimeConst(v string)`

SetExecutionDateTimeConst sets ExecutionDateTimeConst field to given value.


### GetExecutionOffsetMinutes

`func (o *ActionConfigProperty) GetExecutionOffsetMinutes() string`

GetExecutionOffsetMinutes returns the ExecutionOffsetMinutes field if non-nil, zero value otherwise.

### GetExecutionOffsetMinutesOk

`func (o *ActionConfigProperty) GetExecutionOffsetMinutesOk() (*string, bool)`

GetExecutionOffsetMinutesOk returns a tuple with the ExecutionOffsetMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionOffsetMinutes

`func (o *ActionConfigProperty) SetExecutionOffsetMinutes(v string)`

SetExecutionOffsetMinutes sets ExecutionOffsetMinutes field to given value.

### HasExecutionOffsetMinutes

`func (o *ActionConfigProperty) HasExecutionOffsetMinutes() bool`

HasExecutionOffsetMinutes returns a boolean if a field has been set.

### GetFunctionName

`func (o *ActionConfigProperty) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *ActionConfigProperty) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *ActionConfigProperty) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.

### HasFunctionName

`func (o *ActionConfigProperty) HasFunctionName() bool`

HasFunctionName returns a boolean if a field has been set.

### GetHeaders

`func (o *ActionConfigProperty) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ActionConfigProperty) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ActionConfigProperty) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ActionConfigProperty) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetHttpMethod

`func (o *ActionConfigProperty) GetHttpMethod() string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *ActionConfigProperty) GetHttpMethodOk() (*string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *ActionConfigProperty) SetHttpMethod(v string)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *ActionConfigProperty) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetMessage

`func (o *ActionConfigProperty) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ActionConfigProperty) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ActionConfigProperty) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ActionConfigProperty) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetParameter1

`func (o *ActionConfigProperty) GetParameter1() string`

GetParameter1 returns the Parameter1 field if non-nil, zero value otherwise.

### GetParameter1Ok

`func (o *ActionConfigProperty) GetParameter1Ok() (*string, bool)`

GetParameter1Ok returns a tuple with the Parameter1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter1

`func (o *ActionConfigProperty) SetParameter1(v string)`

SetParameter1 sets Parameter1 field to given value.

### HasParameter1

`func (o *ActionConfigProperty) HasParameter1() bool`

HasParameter1 returns a boolean if a field has been set.

### GetParameter2

`func (o *ActionConfigProperty) GetParameter2() string`

GetParameter2 returns the Parameter2 field if non-nil, zero value otherwise.

### GetParameter2Ok

`func (o *ActionConfigProperty) GetParameter2Ok() (*string, bool)`

GetParameter2Ok returns a tuple with the Parameter2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter2

`func (o *ActionConfigProperty) SetParameter2(v string)`

SetParameter2 sets Parameter2 field to given value.

### HasParameter2

`func (o *ActionConfigProperty) HasParameter2() bool`

HasParameter2 returns a boolean if a field has been set.

### GetParameter3

`func (o *ActionConfigProperty) GetParameter3() string`

GetParameter3 returns the Parameter3 field if non-nil, zero value otherwise.

### GetParameter3Ok

`func (o *ActionConfigProperty) GetParameter3Ok() (*string, bool)`

GetParameter3Ok returns a tuple with the Parameter3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter3

`func (o *ActionConfigProperty) SetParameter3(v string)`

SetParameter3 sets Parameter3 field to given value.

### HasParameter3

`func (o *ActionConfigProperty) HasParameter3() bool`

HasParameter3 returns a boolean if a field has been set.

### GetParameter4

`func (o *ActionConfigProperty) GetParameter4() string`

GetParameter4 returns the Parameter4 field if non-nil, zero value otherwise.

### GetParameter4Ok

`func (o *ActionConfigProperty) GetParameter4Ok() (*string, bool)`

GetParameter4Ok returns a tuple with the Parameter4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter4

`func (o *ActionConfigProperty) SetParameter4(v string)`

SetParameter4 sets Parameter4 field to given value.

### HasParameter4

`func (o *ActionConfigProperty) HasParameter4() bool`

HasParameter4 returns a boolean if a field has been set.

### GetParameter5

`func (o *ActionConfigProperty) GetParameter5() string`

GetParameter5 returns the Parameter5 field if non-nil, zero value otherwise.

### GetParameter5Ok

`func (o *ActionConfigProperty) GetParameter5Ok() (*string, bool)`

GetParameter5Ok returns a tuple with the Parameter5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter5

`func (o *ActionConfigProperty) SetParameter5(v string)`

SetParameter5 sets Parameter5 field to given value.

### HasParameter5

`func (o *ActionConfigProperty) HasParameter5() bool`

HasParameter5 returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *ActionConfigProperty) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *ActionConfigProperty) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *ActionConfigProperty) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *ActionConfigProperty) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### GetSpeedClass

`func (o *ActionConfigProperty) GetSpeedClass() string`

GetSpeedClass returns the SpeedClass field if non-nil, zero value otherwise.

### GetSpeedClassOk

`func (o *ActionConfigProperty) GetSpeedClassOk() (*string, bool)`

GetSpeedClassOk returns a tuple with the SpeedClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedClass

`func (o *ActionConfigProperty) SetSpeedClass(v string)`

SetSpeedClass sets SpeedClass field to given value.

### HasSpeedClass

`func (o *ActionConfigProperty) HasSpeedClass() bool`

HasSpeedClass returns a boolean if a field has been set.

### GetTitle

`func (o *ActionConfigProperty) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ActionConfigProperty) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ActionConfigProperty) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ActionConfigProperty) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTo

`func (o *ActionConfigProperty) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ActionConfigProperty) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ActionConfigProperty) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ActionConfigProperty) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetUrl

`func (o *ActionConfigProperty) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ActionConfigProperty) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ActionConfigProperty) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ActionConfigProperty) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


