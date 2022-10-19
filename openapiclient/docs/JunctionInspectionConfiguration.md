# JunctionInspectionConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**Report** | Pointer to [**FunnelConfiguration**](FunnelConfiguration.md) |  | [optional] 

## Methods

### NewJunctionInspectionConfiguration

`func NewJunctionInspectionConfiguration() *JunctionInspectionConfiguration`

NewJunctionInspectionConfiguration instantiates a new JunctionInspectionConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunctionInspectionConfigurationWithDefaults

`func NewJunctionInspectionConfigurationWithDefaults() *JunctionInspectionConfiguration`

NewJunctionInspectionConfigurationWithDefaults instantiates a new JunctionInspectionConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *JunctionInspectionConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JunctionInspectionConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JunctionInspectionConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *JunctionInspectionConfiguration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetReport

`func (o *JunctionInspectionConfiguration) GetReport() FunnelConfiguration`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *JunctionInspectionConfiguration) GetReportOk() (*FunnelConfiguration, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *JunctionInspectionConfiguration) SetReport(v FunnelConfiguration)`

SetReport sets Report field to given value.

### HasReport

`func (o *JunctionInspectionConfiguration) HasReport() bool`

HasReport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


