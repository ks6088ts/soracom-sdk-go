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

// ExpiryTime struct for ExpiryTime
type ExpiryTime struct {
	// 期限切れ時のアクション。以下のいずれかを指定します。各設定について詳しくは、[IoT SIM の有効期限とアクションを設定する](/ja-jp/docs/air/set-expiry/) を参照してください。なお、`terminate` を指定する場合は、あらかじめ解約プロテクションを解除してください。  省略した場合は、null 値が設定されます。 - `doNothing` : 保留 - `deleteSession` : セッション切断 - `deactivate` : 休止 - `suspend` : 利用中断 - `terminate` : 解約 - null 値 : (なし) (`doNothing` と同じ動作です) 
	ExpiryAction NullableString `json:"expiryAction,omitempty"`
	// 有効期限として設定された日付のタイムスタンプ (UNIX 時間 (ミリ秒))
	ExpiryTime int64 `json:"expiryTime"`
}

// NewExpiryTime instantiates a new ExpiryTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpiryTime(expiryTime int64) *ExpiryTime {
	this := ExpiryTime{}
	this.ExpiryTime = expiryTime
	return &this
}

// NewExpiryTimeWithDefaults instantiates a new ExpiryTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpiryTimeWithDefaults() *ExpiryTime {
	this := ExpiryTime{}
	return &this
}

// GetExpiryAction returns the ExpiryAction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpiryTime) GetExpiryAction() string {
	if o == nil || o.ExpiryAction.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExpiryAction.Get()
}

// GetExpiryActionOk returns a tuple with the ExpiryAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpiryTime) GetExpiryActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiryAction.Get(), o.ExpiryAction.IsSet()
}

// HasExpiryAction returns a boolean if a field has been set.
func (o *ExpiryTime) HasExpiryAction() bool {
	if o != nil && o.ExpiryAction.IsSet() {
		return true
	}

	return false
}

// SetExpiryAction gets a reference to the given NullableString and assigns it to the ExpiryAction field.
func (o *ExpiryTime) SetExpiryAction(v string) {
	o.ExpiryAction.Set(&v)
}
// SetExpiryActionNil sets the value for ExpiryAction to be an explicit nil
func (o *ExpiryTime) SetExpiryActionNil() {
	o.ExpiryAction.Set(nil)
}

// UnsetExpiryAction ensures that no value is present for ExpiryAction, not even an explicit nil
func (o *ExpiryTime) UnsetExpiryAction() {
	o.ExpiryAction.Unset()
}

// GetExpiryTime returns the ExpiryTime field value
func (o *ExpiryTime) GetExpiryTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ExpiryTime
}

// GetExpiryTimeOk returns a tuple with the ExpiryTime field value
// and a boolean to check if the value has been set.
func (o *ExpiryTime) GetExpiryTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiryTime, true
}

// SetExpiryTime sets field value
func (o *ExpiryTime) SetExpiryTime(v int64) {
	o.ExpiryTime = v
}

func (o ExpiryTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiryAction.IsSet() {
		toSerialize["expiryAction"] = o.ExpiryAction.Get()
	}
	if true {
		toSerialize["expiryTime"] = o.ExpiryTime
	}
	return json.Marshal(toSerialize)
}

type NullableExpiryTime struct {
	value *ExpiryTime
	isSet bool
}

func (v NullableExpiryTime) Get() *ExpiryTime {
	return v.value
}

func (v *NullableExpiryTime) Set(val *ExpiryTime) {
	v.value = val
	v.isSet = true
}

func (v NullableExpiryTime) IsSet() bool {
	return v.isSet
}

func (v *NullableExpiryTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpiryTime(val *ExpiryTime) *NullableExpiryTime {
	return &NullableExpiryTime{value: val, isSet: true}
}

func (v NullableExpiryTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpiryTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


