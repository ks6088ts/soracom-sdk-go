# EventHandlerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionConfigList** | [**[]ActionConfig**](ActionConfig.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**HandlerId** | **string** |  | 
**Name** | **string** |  | 
**RuleConfig** | [**RuleConfig**](RuleConfig.md) |  | 
**Status** | **string** |  | 
**TargetGroupId** | Pointer to **string** |  | [optional] 
**TargetImsi** | Pointer to **string** |  | [optional] 
**TargetOperatorId** | Pointer to **string** |  | [optional] 
**TargetSimId** | Pointer to **string** |  | [optional] 

## Methods

### NewEventHandlerModel

`func NewEventHandlerModel(actionConfigList []ActionConfig, handlerId string, name string, ruleConfig RuleConfig, status string, ) *EventHandlerModel`

NewEventHandlerModel instantiates a new EventHandlerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventHandlerModelWithDefaults

`func NewEventHandlerModelWithDefaults() *EventHandlerModel`

NewEventHandlerModelWithDefaults instantiates a new EventHandlerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionConfigList

`func (o *EventHandlerModel) GetActionConfigList() []ActionConfig`

GetActionConfigList returns the ActionConfigList field if non-nil, zero value otherwise.

### GetActionConfigListOk

`func (o *EventHandlerModel) GetActionConfigListOk() (*[]ActionConfig, bool)`

GetActionConfigListOk returns a tuple with the ActionConfigList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionConfigList

`func (o *EventHandlerModel) SetActionConfigList(v []ActionConfig)`

SetActionConfigList sets ActionConfigList field to given value.


### GetDescription

`func (o *EventHandlerModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventHandlerModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventHandlerModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventHandlerModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHandlerId

`func (o *EventHandlerModel) GetHandlerId() string`

GetHandlerId returns the HandlerId field if non-nil, zero value otherwise.

### GetHandlerIdOk

`func (o *EventHandlerModel) GetHandlerIdOk() (*string, bool)`

GetHandlerIdOk returns a tuple with the HandlerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerId

`func (o *EventHandlerModel) SetHandlerId(v string)`

SetHandlerId sets HandlerId field to given value.


### GetName

`func (o *EventHandlerModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventHandlerModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventHandlerModel) SetName(v string)`

SetName sets Name field to given value.


### GetRuleConfig

`func (o *EventHandlerModel) GetRuleConfig() RuleConfig`

GetRuleConfig returns the RuleConfig field if non-nil, zero value otherwise.

### GetRuleConfigOk

`func (o *EventHandlerModel) GetRuleConfigOk() (*RuleConfig, bool)`

GetRuleConfigOk returns a tuple with the RuleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleConfig

`func (o *EventHandlerModel) SetRuleConfig(v RuleConfig)`

SetRuleConfig sets RuleConfig field to given value.


### GetStatus

`func (o *EventHandlerModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventHandlerModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventHandlerModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTargetGroupId

`func (o *EventHandlerModel) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *EventHandlerModel) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *EventHandlerModel) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *EventHandlerModel) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### GetTargetImsi

`func (o *EventHandlerModel) GetTargetImsi() string`

GetTargetImsi returns the TargetImsi field if non-nil, zero value otherwise.

### GetTargetImsiOk

`func (o *EventHandlerModel) GetTargetImsiOk() (*string, bool)`

GetTargetImsiOk returns a tuple with the TargetImsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetImsi

`func (o *EventHandlerModel) SetTargetImsi(v string)`

SetTargetImsi sets TargetImsi field to given value.

### HasTargetImsi

`func (o *EventHandlerModel) HasTargetImsi() bool`

HasTargetImsi returns a boolean if a field has been set.

### GetTargetOperatorId

`func (o *EventHandlerModel) GetTargetOperatorId() string`

GetTargetOperatorId returns the TargetOperatorId field if non-nil, zero value otherwise.

### GetTargetOperatorIdOk

`func (o *EventHandlerModel) GetTargetOperatorIdOk() (*string, bool)`

GetTargetOperatorIdOk returns a tuple with the TargetOperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetOperatorId

`func (o *EventHandlerModel) SetTargetOperatorId(v string)`

SetTargetOperatorId sets TargetOperatorId field to given value.

### HasTargetOperatorId

`func (o *EventHandlerModel) HasTargetOperatorId() bool`

HasTargetOperatorId returns a boolean if a field has been set.

### GetTargetSimId

`func (o *EventHandlerModel) GetTargetSimId() string`

GetTargetSimId returns the TargetSimId field if non-nil, zero value otherwise.

### GetTargetSimIdOk

`func (o *EventHandlerModel) GetTargetSimIdOk() (*string, bool)`

GetTargetSimIdOk returns a tuple with the TargetSimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSimId

`func (o *EventHandlerModel) SetTargetSimId(v string)`

SetTargetSimId sets TargetSimId field to given value.

### HasTargetSimId

`func (o *EventHandlerModel) HasTargetSimId() bool`

HasTargetSimId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


