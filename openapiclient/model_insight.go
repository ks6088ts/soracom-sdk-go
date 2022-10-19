/*
SORACOM API

SORACOM API v1

API version: VERSION_PLACEHOLDER
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"encoding/json"
)

// Insight struct for Insight
type Insight struct {
	// 異常と思われる状況を検出した時間。フォーマットは category に関するイベントのタイムスタンプに依存します。（例：session の場合は UNIX 時間 (ミリ秒)）
	AnomalyDetectedTimes []string `json:"anomalyDetectedTimes,omitempty"`
	Category *string `json:"category,omitempty"`
	InsightId *string `json:"insightId,omitempty"`
	Message *string `json:"message,omitempty"`
	// インサイトに関連のある参考情報の URL
	ReferenceUrls []map[string]ReferenceUrl `json:"referenceUrls,omitempty"`
	Severity *string `json:"severity,omitempty"`
}

// NewInsight instantiates a new Insight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsight() *Insight {
	this := Insight{}
	return &this
}

// NewInsightWithDefaults instantiates a new Insight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsightWithDefaults() *Insight {
	this := Insight{}
	return &this
}

// GetAnomalyDetectedTimes returns the AnomalyDetectedTimes field value if set, zero value otherwise.
func (o *Insight) GetAnomalyDetectedTimes() []string {
	if o == nil || o.AnomalyDetectedTimes == nil {
		var ret []string
		return ret
	}
	return o.AnomalyDetectedTimes
}

// GetAnomalyDetectedTimesOk returns a tuple with the AnomalyDetectedTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insight) GetAnomalyDetectedTimesOk() ([]string, bool) {
	if o == nil || o.AnomalyDetectedTimes == nil {
		return nil, false
	}
	return o.AnomalyDetectedTimes, true
}

// HasAnomalyDetectedTimes returns a boolean if a field has been set.
func (o *Insight) HasAnomalyDetectedTimes() bool {
	if o != nil && o.AnomalyDetectedTimes != nil {
		return true
	}

	return false
}

// SetAnomalyDetectedTimes gets a reference to the given []string and assigns it to the AnomalyDetectedTimes field.
func (o *Insight) SetAnomalyDetectedTimes(v []string) {
	o.AnomalyDetectedTimes = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Insight) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insight) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Insight) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *Insight) SetCategory(v string) {
	o.Category = &v
}

// GetInsightId returns the InsightId field value if set, zero value otherwise.
func (o *Insight) GetInsightId() string {
	if o == nil || o.InsightId == nil {
		var ret string
		return ret
	}
	return *o.InsightId
}

// GetInsightIdOk returns a tuple with the InsightId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insight) GetInsightIdOk() (*string, bool) {
	if o == nil || o.InsightId == nil {
		return nil, false
	}
	return o.InsightId, true
}

// HasInsightId returns a boolean if a field has been set.
func (o *Insight) HasInsightId() bool {
	if o != nil && o.InsightId != nil {
		return true
	}

	return false
}

// SetInsightId gets a reference to the given string and assigns it to the InsightId field.
func (o *Insight) SetInsightId(v string) {
	o.InsightId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Insight) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insight) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Insight) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Insight) SetMessage(v string) {
	o.Message = &v
}

// GetReferenceUrls returns the ReferenceUrls field value if set, zero value otherwise.
func (o *Insight) GetReferenceUrls() []map[string]ReferenceUrl {
	if o == nil || o.ReferenceUrls == nil {
		var ret []map[string]ReferenceUrl
		return ret
	}
	return o.ReferenceUrls
}

// GetReferenceUrlsOk returns a tuple with the ReferenceUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insight) GetReferenceUrlsOk() ([]map[string]ReferenceUrl, bool) {
	if o == nil || o.ReferenceUrls == nil {
		return nil, false
	}
	return o.ReferenceUrls, true
}

// HasReferenceUrls returns a boolean if a field has been set.
func (o *Insight) HasReferenceUrls() bool {
	if o != nil && o.ReferenceUrls != nil {
		return true
	}

	return false
}

// SetReferenceUrls gets a reference to the given []map[string]ReferenceUrl and assigns it to the ReferenceUrls field.
func (o *Insight) SetReferenceUrls(v []map[string]ReferenceUrl) {
	o.ReferenceUrls = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *Insight) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insight) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *Insight) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *Insight) SetSeverity(v string) {
	o.Severity = &v
}

func (o Insight) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnomalyDetectedTimes != nil {
		toSerialize["anomalyDetectedTimes"] = o.AnomalyDetectedTimes
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.InsightId != nil {
		toSerialize["insightId"] = o.InsightId
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ReferenceUrls != nil {
		toSerialize["referenceUrls"] = o.ReferenceUrls
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	return json.Marshal(toSerialize)
}

type NullableInsight struct {
	value *Insight
	isSet bool
}

func (v NullableInsight) Get() *Insight {
	return v.value
}

func (v *NullableInsight) Set(val *Insight) {
	v.value = val
	v.isSet = true
}

func (v NullableInsight) IsSet() bool {
	return v.isSet
}

func (v *NullableInsight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsight(val *Insight) *NullableInsight {
	return &NullableInsight{value: val, isSet: true}
}

func (v NullableInsight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

