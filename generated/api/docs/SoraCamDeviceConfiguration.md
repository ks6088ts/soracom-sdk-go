# SoraCamDeviceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioAlarmEnabled** | Pointer to **bool** | 音声アラームが有効化されているかどうか | [optional] 
**MotionDetectionEnabled** | Pointer to **bool** | モーション検知が有効化されているかどうか | [optional] 
**SmokeAlarmEnabled** | Pointer to **bool** | 煙検知が有効化されているかどうか | [optional] 

## Methods

### NewSoraCamDeviceConfiguration

`func NewSoraCamDeviceConfiguration() *SoraCamDeviceConfiguration`

NewSoraCamDeviceConfiguration instantiates a new SoraCamDeviceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoraCamDeviceConfigurationWithDefaults

`func NewSoraCamDeviceConfigurationWithDefaults() *SoraCamDeviceConfiguration`

NewSoraCamDeviceConfigurationWithDefaults instantiates a new SoraCamDeviceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioAlarmEnabled

`func (o *SoraCamDeviceConfiguration) GetAudioAlarmEnabled() bool`

GetAudioAlarmEnabled returns the AudioAlarmEnabled field if non-nil, zero value otherwise.

### GetAudioAlarmEnabledOk

`func (o *SoraCamDeviceConfiguration) GetAudioAlarmEnabledOk() (*bool, bool)`

GetAudioAlarmEnabledOk returns a tuple with the AudioAlarmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioAlarmEnabled

`func (o *SoraCamDeviceConfiguration) SetAudioAlarmEnabled(v bool)`

SetAudioAlarmEnabled sets AudioAlarmEnabled field to given value.

### HasAudioAlarmEnabled

`func (o *SoraCamDeviceConfiguration) HasAudioAlarmEnabled() bool`

HasAudioAlarmEnabled returns a boolean if a field has been set.

### GetMotionDetectionEnabled

`func (o *SoraCamDeviceConfiguration) GetMotionDetectionEnabled() bool`

GetMotionDetectionEnabled returns the MotionDetectionEnabled field if non-nil, zero value otherwise.

### GetMotionDetectionEnabledOk

`func (o *SoraCamDeviceConfiguration) GetMotionDetectionEnabledOk() (*bool, bool)`

GetMotionDetectionEnabledOk returns a tuple with the MotionDetectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotionDetectionEnabled

`func (o *SoraCamDeviceConfiguration) SetMotionDetectionEnabled(v bool)`

SetMotionDetectionEnabled sets MotionDetectionEnabled field to given value.

### HasMotionDetectionEnabled

`func (o *SoraCamDeviceConfiguration) HasMotionDetectionEnabled() bool`

HasMotionDetectionEnabled returns a boolean if a field has been set.

### GetSmokeAlarmEnabled

`func (o *SoraCamDeviceConfiguration) GetSmokeAlarmEnabled() bool`

GetSmokeAlarmEnabled returns the SmokeAlarmEnabled field if non-nil, zero value otherwise.

### GetSmokeAlarmEnabledOk

`func (o *SoraCamDeviceConfiguration) GetSmokeAlarmEnabledOk() (*bool, bool)`

GetSmokeAlarmEnabledOk returns a tuple with the SmokeAlarmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmokeAlarmEnabled

`func (o *SoraCamDeviceConfiguration) SetSmokeAlarmEnabled(v bool)`

SetSmokeAlarmEnabled sets SmokeAlarmEnabled field to given value.

### HasSmokeAlarmEnabled

`func (o *SoraCamDeviceConfiguration) HasSmokeAlarmEnabled() bool`

HasSmokeAlarmEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


