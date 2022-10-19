# RuleConfigProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InactiveTimeoutDateConst** | **string** |  | 
**InactiveTimeoutOffsetMinutes** | Pointer to **string** |  | [optional] 
**LimitTotalAmount** | Pointer to **string** |  | [optional] 
**LimitTotalTrafficMegaByte** | Pointer to **string** |  | [optional] 
**RunOnceAmongTarget** | Pointer to **string** |  | [optional] 
**TargetOtaStatus** | Pointer to **string** | SimSubscriptionStatusRule の時のみ有効 | [optional] 
**TargetSpeedClass** | Pointer to **string** | SubscriberSpeedClassAttributeRule の時のみ有効 | [optional] 
**TargetStatus** | Pointer to **string** | SubscriberStatusAttributeRule の時のみ有効 | [optional] 

## Methods

### NewRuleConfigProperty

`func NewRuleConfigProperty(inactiveTimeoutDateConst string, ) *RuleConfigProperty`

NewRuleConfigProperty instantiates a new RuleConfigProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleConfigPropertyWithDefaults

`func NewRuleConfigPropertyWithDefaults() *RuleConfigProperty`

NewRuleConfigPropertyWithDefaults instantiates a new RuleConfigProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInactiveTimeoutDateConst

`func (o *RuleConfigProperty) GetInactiveTimeoutDateConst() string`

GetInactiveTimeoutDateConst returns the InactiveTimeoutDateConst field if non-nil, zero value otherwise.

### GetInactiveTimeoutDateConstOk

`func (o *RuleConfigProperty) GetInactiveTimeoutDateConstOk() (*string, bool)`

GetInactiveTimeoutDateConstOk returns a tuple with the InactiveTimeoutDateConst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveTimeoutDateConst

`func (o *RuleConfigProperty) SetInactiveTimeoutDateConst(v string)`

SetInactiveTimeoutDateConst sets InactiveTimeoutDateConst field to given value.


### GetInactiveTimeoutOffsetMinutes

`func (o *RuleConfigProperty) GetInactiveTimeoutOffsetMinutes() string`

GetInactiveTimeoutOffsetMinutes returns the InactiveTimeoutOffsetMinutes field if non-nil, zero value otherwise.

### GetInactiveTimeoutOffsetMinutesOk

`func (o *RuleConfigProperty) GetInactiveTimeoutOffsetMinutesOk() (*string, bool)`

GetInactiveTimeoutOffsetMinutesOk returns a tuple with the InactiveTimeoutOffsetMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveTimeoutOffsetMinutes

`func (o *RuleConfigProperty) SetInactiveTimeoutOffsetMinutes(v string)`

SetInactiveTimeoutOffsetMinutes sets InactiveTimeoutOffsetMinutes field to given value.

### HasInactiveTimeoutOffsetMinutes

`func (o *RuleConfigProperty) HasInactiveTimeoutOffsetMinutes() bool`

HasInactiveTimeoutOffsetMinutes returns a boolean if a field has been set.

### GetLimitTotalAmount

`func (o *RuleConfigProperty) GetLimitTotalAmount() string`

GetLimitTotalAmount returns the LimitTotalAmount field if non-nil, zero value otherwise.

### GetLimitTotalAmountOk

`func (o *RuleConfigProperty) GetLimitTotalAmountOk() (*string, bool)`

GetLimitTotalAmountOk returns a tuple with the LimitTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitTotalAmount

`func (o *RuleConfigProperty) SetLimitTotalAmount(v string)`

SetLimitTotalAmount sets LimitTotalAmount field to given value.

### HasLimitTotalAmount

`func (o *RuleConfigProperty) HasLimitTotalAmount() bool`

HasLimitTotalAmount returns a boolean if a field has been set.

### GetLimitTotalTrafficMegaByte

`func (o *RuleConfigProperty) GetLimitTotalTrafficMegaByte() string`

GetLimitTotalTrafficMegaByte returns the LimitTotalTrafficMegaByte field if non-nil, zero value otherwise.

### GetLimitTotalTrafficMegaByteOk

`func (o *RuleConfigProperty) GetLimitTotalTrafficMegaByteOk() (*string, bool)`

GetLimitTotalTrafficMegaByteOk returns a tuple with the LimitTotalTrafficMegaByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitTotalTrafficMegaByte

`func (o *RuleConfigProperty) SetLimitTotalTrafficMegaByte(v string)`

SetLimitTotalTrafficMegaByte sets LimitTotalTrafficMegaByte field to given value.

### HasLimitTotalTrafficMegaByte

`func (o *RuleConfigProperty) HasLimitTotalTrafficMegaByte() bool`

HasLimitTotalTrafficMegaByte returns a boolean if a field has been set.

### GetRunOnceAmongTarget

`func (o *RuleConfigProperty) GetRunOnceAmongTarget() string`

GetRunOnceAmongTarget returns the RunOnceAmongTarget field if non-nil, zero value otherwise.

### GetRunOnceAmongTargetOk

`func (o *RuleConfigProperty) GetRunOnceAmongTargetOk() (*string, bool)`

GetRunOnceAmongTargetOk returns a tuple with the RunOnceAmongTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOnceAmongTarget

`func (o *RuleConfigProperty) SetRunOnceAmongTarget(v string)`

SetRunOnceAmongTarget sets RunOnceAmongTarget field to given value.

### HasRunOnceAmongTarget

`func (o *RuleConfigProperty) HasRunOnceAmongTarget() bool`

HasRunOnceAmongTarget returns a boolean if a field has been set.

### GetTargetOtaStatus

`func (o *RuleConfigProperty) GetTargetOtaStatus() string`

GetTargetOtaStatus returns the TargetOtaStatus field if non-nil, zero value otherwise.

### GetTargetOtaStatusOk

`func (o *RuleConfigProperty) GetTargetOtaStatusOk() (*string, bool)`

GetTargetOtaStatusOk returns a tuple with the TargetOtaStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetOtaStatus

`func (o *RuleConfigProperty) SetTargetOtaStatus(v string)`

SetTargetOtaStatus sets TargetOtaStatus field to given value.

### HasTargetOtaStatus

`func (o *RuleConfigProperty) HasTargetOtaStatus() bool`

HasTargetOtaStatus returns a boolean if a field has been set.

### GetTargetSpeedClass

`func (o *RuleConfigProperty) GetTargetSpeedClass() string`

GetTargetSpeedClass returns the TargetSpeedClass field if non-nil, zero value otherwise.

### GetTargetSpeedClassOk

`func (o *RuleConfigProperty) GetTargetSpeedClassOk() (*string, bool)`

GetTargetSpeedClassOk returns a tuple with the TargetSpeedClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSpeedClass

`func (o *RuleConfigProperty) SetTargetSpeedClass(v string)`

SetTargetSpeedClass sets TargetSpeedClass field to given value.

### HasTargetSpeedClass

`func (o *RuleConfigProperty) HasTargetSpeedClass() bool`

HasTargetSpeedClass returns a boolean if a field has been set.

### GetTargetStatus

`func (o *RuleConfigProperty) GetTargetStatus() string`

GetTargetStatus returns the TargetStatus field if non-nil, zero value otherwise.

### GetTargetStatusOk

`func (o *RuleConfigProperty) GetTargetStatusOk() (*string, bool)`

GetTargetStatusOk returns a tuple with the TargetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStatus

`func (o *RuleConfigProperty) SetTargetStatus(v string)`

SetTargetStatus sets TargetStatus field to given value.

### HasTargetStatus

`func (o *RuleConfigProperty) HasTargetStatus() bool`

HasTargetStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


