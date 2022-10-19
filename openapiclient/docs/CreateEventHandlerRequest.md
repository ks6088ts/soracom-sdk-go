# CreateEventHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionConfigList** | [**[]ActionConfig**](ActionConfig.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RuleConfig** | [**RuleConfig**](RuleConfig.md) |  | 
**Status** | **string** |  | 
**TargetGroupId** | Pointer to **string** |  | [optional] 
**TargetImsi** | Pointer to **string** |  | [optional] 
**TargetOperatorId** | Pointer to **string** |  | [optional] 
**TargetSimId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateEventHandlerRequest

`func NewCreateEventHandlerRequest(actionConfigList []ActionConfig, ruleConfig RuleConfig, status string, ) *CreateEventHandlerRequest`

NewCreateEventHandlerRequest instantiates a new CreateEventHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEventHandlerRequestWithDefaults

`func NewCreateEventHandlerRequestWithDefaults() *CreateEventHandlerRequest`

NewCreateEventHandlerRequestWithDefaults instantiates a new CreateEventHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionConfigList

`func (o *CreateEventHandlerRequest) GetActionConfigList() []ActionConfig`

GetActionConfigList returns the ActionConfigList field if non-nil, zero value otherwise.

### GetActionConfigListOk

`func (o *CreateEventHandlerRequest) GetActionConfigListOk() (*[]ActionConfig, bool)`

GetActionConfigListOk returns a tuple with the ActionConfigList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionConfigList

`func (o *CreateEventHandlerRequest) SetActionConfigList(v []ActionConfig)`

SetActionConfigList sets ActionConfigList field to given value.


### GetDescription

`func (o *CreateEventHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateEventHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateEventHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateEventHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CreateEventHandlerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEventHandlerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEventHandlerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateEventHandlerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuleConfig

`func (o *CreateEventHandlerRequest) GetRuleConfig() RuleConfig`

GetRuleConfig returns the RuleConfig field if non-nil, zero value otherwise.

### GetRuleConfigOk

`func (o *CreateEventHandlerRequest) GetRuleConfigOk() (*RuleConfig, bool)`

GetRuleConfigOk returns a tuple with the RuleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleConfig

`func (o *CreateEventHandlerRequest) SetRuleConfig(v RuleConfig)`

SetRuleConfig sets RuleConfig field to given value.


### GetStatus

`func (o *CreateEventHandlerRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateEventHandlerRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateEventHandlerRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTargetGroupId

`func (o *CreateEventHandlerRequest) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *CreateEventHandlerRequest) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *CreateEventHandlerRequest) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *CreateEventHandlerRequest) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### GetTargetImsi

`func (o *CreateEventHandlerRequest) GetTargetImsi() string`

GetTargetImsi returns the TargetImsi field if non-nil, zero value otherwise.

### GetTargetImsiOk

`func (o *CreateEventHandlerRequest) GetTargetImsiOk() (*string, bool)`

GetTargetImsiOk returns a tuple with the TargetImsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetImsi

`func (o *CreateEventHandlerRequest) SetTargetImsi(v string)`

SetTargetImsi sets TargetImsi field to given value.

### HasTargetImsi

`func (o *CreateEventHandlerRequest) HasTargetImsi() bool`

HasTargetImsi returns a boolean if a field has been set.

### GetTargetOperatorId

`func (o *CreateEventHandlerRequest) GetTargetOperatorId() string`

GetTargetOperatorId returns the TargetOperatorId field if non-nil, zero value otherwise.

### GetTargetOperatorIdOk

`func (o *CreateEventHandlerRequest) GetTargetOperatorIdOk() (*string, bool)`

GetTargetOperatorIdOk returns a tuple with the TargetOperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetOperatorId

`func (o *CreateEventHandlerRequest) SetTargetOperatorId(v string)`

SetTargetOperatorId sets TargetOperatorId field to given value.

### HasTargetOperatorId

`func (o *CreateEventHandlerRequest) HasTargetOperatorId() bool`

HasTargetOperatorId returns a boolean if a field has been set.

### GetTargetSimId

`func (o *CreateEventHandlerRequest) GetTargetSimId() string`

GetTargetSimId returns the TargetSimId field if non-nil, zero value otherwise.

### GetTargetSimIdOk

`func (o *CreateEventHandlerRequest) GetTargetSimIdOk() (*string, bool)`

GetTargetSimIdOk returns a tuple with the TargetSimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSimId

`func (o *CreateEventHandlerRequest) SetTargetSimId(v string)`

SetTargetSimId sets TargetSimId field to given value.

### HasTargetSimId

`func (o *CreateEventHandlerRequest) HasTargetSimId() bool`

HasTargetSimId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


