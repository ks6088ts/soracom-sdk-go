# Insight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnomalyDetectedTimes** | Pointer to **[]string** | 異常と思われる状況を検出した時間。フォーマットは category に関するイベントのタイムスタンプに依存します。（例：session の場合は UNIX 時間 (ミリ秒)） | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**InsightId** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ReferenceUrls** | Pointer to [**[]map[string]ReferenceUrl**](map[string]ReferenceUrl.md) | インサイトに関連のある参考情報の URL | [optional] 
**Severity** | Pointer to **string** |  | [optional] 

## Methods

### NewInsight

`func NewInsight() *Insight`

NewInsight instantiates a new Insight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightWithDefaults

`func NewInsightWithDefaults() *Insight`

NewInsightWithDefaults instantiates a new Insight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnomalyDetectedTimes

`func (o *Insight) GetAnomalyDetectedTimes() []string`

GetAnomalyDetectedTimes returns the AnomalyDetectedTimes field if non-nil, zero value otherwise.

### GetAnomalyDetectedTimesOk

`func (o *Insight) GetAnomalyDetectedTimesOk() (*[]string, bool)`

GetAnomalyDetectedTimesOk returns a tuple with the AnomalyDetectedTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyDetectedTimes

`func (o *Insight) SetAnomalyDetectedTimes(v []string)`

SetAnomalyDetectedTimes sets AnomalyDetectedTimes field to given value.

### HasAnomalyDetectedTimes

`func (o *Insight) HasAnomalyDetectedTimes() bool`

HasAnomalyDetectedTimes returns a boolean if a field has been set.

### GetCategory

`func (o *Insight) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Insight) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Insight) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Insight) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetInsightId

`func (o *Insight) GetInsightId() string`

GetInsightId returns the InsightId field if non-nil, zero value otherwise.

### GetInsightIdOk

`func (o *Insight) GetInsightIdOk() (*string, bool)`

GetInsightIdOk returns a tuple with the InsightId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightId

`func (o *Insight) SetInsightId(v string)`

SetInsightId sets InsightId field to given value.

### HasInsightId

`func (o *Insight) HasInsightId() bool`

HasInsightId returns a boolean if a field has been set.

### GetMessage

`func (o *Insight) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Insight) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Insight) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Insight) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReferenceUrls

`func (o *Insight) GetReferenceUrls() []map[string]ReferenceUrl`

GetReferenceUrls returns the ReferenceUrls field if non-nil, zero value otherwise.

### GetReferenceUrlsOk

`func (o *Insight) GetReferenceUrlsOk() (*[]map[string]ReferenceUrl, bool)`

GetReferenceUrlsOk returns a tuple with the ReferenceUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceUrls

`func (o *Insight) SetReferenceUrls(v []map[string]ReferenceUrl)`

SetReferenceUrls sets ReferenceUrls field to given value.

### HasReferenceUrls

`func (o *Insight) HasReferenceUrls() bool`

HasReferenceUrls returns a boolean if a field has been set.

### GetSeverity

`func (o *Insight) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Insight) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Insight) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Insight) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


