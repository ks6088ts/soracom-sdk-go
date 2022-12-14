/*
SORACOM API

SORACOM API v1

API version: VERSION_PLACEHOLDER
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SoraCamImageExportInfo struct for SoraCamImageExportInfo
type SoraCamImageExportInfo struct {
	// ソラカメ対応カメラのデバイス ID
	DeviceId *string `json:"deviceId,omitempty"`
	// URL の有効期限 (UNIX 時間 (ミリ秒))。`status` が `completed` の場合のみ含まれます。
	ExpiryTime *int64 `json:"expiryTime,omitempty"`
	// エクスポート ID。[`SoraCam:exportSoraCamDeviceRecordedImage API`](/ja-jp/tools/api/reference/#/SoraCam/exportSoraCamDeviceRecordedImage) で取得したエクスポート ID を、[`SoraCam:getSoraCamDeviceExportedImage API`](/ja-jp/tools/api/reference/#/SoraCam/getSoraCamDeviceExportedImage) で指定すると、jpg ファイルをダウンロードするための URL を取得できます。 
	ExportId *string `json:"exportId,omitempty"`
	// [`SoraCam:exportSoraCamDeviceRecordedImage API`](/ja-jp/tools/api/reference/#/SoraCam/exportSoraCamDeviceRecordedImage) を呼び出したオペレーターの ID 
	OperatorId *string `json:"operatorId,omitempty"`
	// [`SoraCam:exportSoraCamDeviceRecordedImage API`](/ja-jp/tools/api/reference/#/SoraCam/exportSoraCamDeviceRecordedImage) によるエクスポートのリクエストを SORACOM プラットフォームが受け付けた日時 (UNIX 時間 (ミリ秒)) 
	RequestedTime *int64 `json:"requestedTime,omitempty"`
	// エクスポート処理の現在の状況。  - `initializing`: 初期状態 - `processing`: 処理中 - `completed`: エクスポート完了 - `failed`: エクスポート失敗 - `expired`: URL の有効期限切れ 
	Status *string `json:"status,omitempty"`
	// エクスポートされた jpg ファイルをダウンロードするための URL。`status` が `completed` の場合のみ含まれます。
	Url *string `json:"url,omitempty"`
}

// NewSoraCamImageExportInfo instantiates a new SoraCamImageExportInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoraCamImageExportInfo() *SoraCamImageExportInfo {
	this := SoraCamImageExportInfo{}
	return &this
}

// NewSoraCamImageExportInfoWithDefaults instantiates a new SoraCamImageExportInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoraCamImageExportInfoWithDefaults() *SoraCamImageExportInfo {
	this := SoraCamImageExportInfo{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *SoraCamImageExportInfo) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoraCamImageExportInfo) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *SoraCamImageExportInfo) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *SoraCamImageExportInfo) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetExpiryTime returns the ExpiryTime field value if set, zero value otherwise.
func (o *SoraCamImageExportInfo) GetExpiryTime() int64 {
	if o == nil || o.ExpiryTime == nil {
		var ret int64
		return ret
	}
	return *o.ExpiryTime
}

// GetExpiryTimeOk returns a tuple with the ExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoraCamImageExportInfo) GetExpiryTimeOk() (*int64, bool) {
	if o == nil || o.ExpiryTime == nil {
		return nil, false
	}
	return o.ExpiryTime, true
}

// HasExpiryTime returns a boolean if a field has been set.
func (o *SoraCamImageExportInfo) HasExpiryTime() bool {
	if o != nil && o.ExpiryTime != nil {
		return true
	}

	return false
}

// SetExpiryTime gets a reference to the given int64 and assigns it to the ExpiryTime field.
func (o *SoraCamImageExportInfo) SetExpiryTime(v int64) {
	o.ExpiryTime = &v
}

// GetExportId returns the ExportId field value if set, zero value otherwise.
func (o *SoraCamImageExportInfo) GetExportId() string {
	if o == nil || o.ExportId == nil {
		var ret string
		return ret
	}
	return *o.ExportId
}

// GetExportIdOk returns a tuple with the ExportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoraCamImageExportInfo) GetExportIdOk() (*string, bool) {
	if o == nil || o.ExportId == nil {
		return nil, false
	}
	return o.ExportId, true
}

// HasExportId returns a boolean if a field has been set.
func (o *SoraCamImageExportInfo) HasExportId() bool {
	if o != nil && o.ExportId != nil {
		return true
	}

	return false
}

// SetExportId gets a reference to the given string and assigns it to the ExportId field.
func (o *SoraCamImageExportInfo) SetExportId(v string) {
	o.ExportId = &v
}

// GetOperatorId returns the OperatorId field value if set, zero value otherwise.
func (o *SoraCamImageExportInfo) GetOperatorId() string {
	if o == nil || o.OperatorId == nil {
		var ret string
		return ret
	}
	return *o.OperatorId
}

// GetOperatorIdOk returns a tuple with the OperatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoraCamImageExportInfo) GetOperatorIdOk() (*string, bool) {
	if o == nil || o.OperatorId == nil {
		return nil, false
	}
	return o.OperatorId, true
}

// HasOperatorId returns a boolean if a field has been set.
func (o *SoraCamImageExportInfo) HasOperatorId() bool {
	if o != nil && o.OperatorId != nil {
		return true
	}

	return false
}

// SetOperatorId gets a reference to the given string and assigns it to the OperatorId field.
func (o *SoraCamImageExportInfo) SetOperatorId(v string) {
	o.OperatorId = &v
}

// GetRequestedTime returns the RequestedTime field value if set, zero value otherwise.
func (o *SoraCamImageExportInfo) GetRequestedTime() int64 {
	if o == nil || o.RequestedTime == nil {
		var ret int64
		return ret
	}
	return *o.RequestedTime
}

// GetRequestedTimeOk returns a tuple with the RequestedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoraCamImageExportInfo) GetRequestedTimeOk() (*int64, bool) {
	if o == nil || o.RequestedTime == nil {
		return nil, false
	}
	return o.RequestedTime, true
}

// HasRequestedTime returns a boolean if a field has been set.
func (o *SoraCamImageExportInfo) HasRequestedTime() bool {
	if o != nil && o.RequestedTime != nil {
		return true
	}

	return false
}

// SetRequestedTime gets a reference to the given int64 and assigns it to the RequestedTime field.
func (o *SoraCamImageExportInfo) SetRequestedTime(v int64) {
	o.RequestedTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SoraCamImageExportInfo) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoraCamImageExportInfo) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SoraCamImageExportInfo) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SoraCamImageExportInfo) SetStatus(v string) {
	o.Status = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SoraCamImageExportInfo) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoraCamImageExportInfo) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SoraCamImageExportInfo) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SoraCamImageExportInfo) SetUrl(v string) {
	o.Url = &v
}

func (o SoraCamImageExportInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.ExpiryTime != nil {
		toSerialize["expiryTime"] = o.ExpiryTime
	}
	if o.ExportId != nil {
		toSerialize["exportId"] = o.ExportId
	}
	if o.OperatorId != nil {
		toSerialize["operatorId"] = o.OperatorId
	}
	if o.RequestedTime != nil {
		toSerialize["requestedTime"] = o.RequestedTime
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableSoraCamImageExportInfo struct {
	value *SoraCamImageExportInfo
	isSet bool
}

func (v NullableSoraCamImageExportInfo) Get() *SoraCamImageExportInfo {
	return v.value
}

func (v *NullableSoraCamImageExportInfo) Set(val *SoraCamImageExportInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSoraCamImageExportInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSoraCamImageExportInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoraCamImageExportInfo(val *SoraCamImageExportInfo) *NullableSoraCamImageExportInfo {
	return &NullableSoraCamImageExportInfo{value: val, isSet: true}
}

func (v NullableSoraCamImageExportInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoraCamImageExportInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


