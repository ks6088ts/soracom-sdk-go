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

// AddSubscriptionRequest struct for AddSubscriptionRequest
type AddSubscriptionRequest struct {
	// セカンダリサブスクリプションのダウンロードが完了すると、すぐにそのセカンダリサブスクリプションを有効化 (active) するかどうかを指定します。 - `true`: すぐに有効化 (active) する - `false`: 無効 (inactive) を維持する。次の [ネットワーク登録](/ja-jp/resources/glossary/#ネットワーク登録) 時に、セカンダリサブスクリプションが有効化 (active) されます。 
	Enable *bool `json:"enable,omitempty"`
	// 追加する [セカンダリサブスクリプション](/ja-jp/resources/glossary/subscription/) - `planP1`、`planX1`、`planX2`、`planX3`: 追加サブスクリプション - `planArc01`: バーチャル SIM/Subscriber 
	Subscription string `json:"subscription"`
	// - `virtual`: `subscription` で `planArc01` を指定した場合 - `cellular`: `subscription` で `planArc01` 以外を指定した場合 
	Type *string `json:"type,omitempty"`
}

// NewAddSubscriptionRequest instantiates a new AddSubscriptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSubscriptionRequest(subscription string) *AddSubscriptionRequest {
	this := AddSubscriptionRequest{}
	this.Subscription = subscription
	return &this
}

// NewAddSubscriptionRequestWithDefaults instantiates a new AddSubscriptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSubscriptionRequestWithDefaults() *AddSubscriptionRequest {
	this := AddSubscriptionRequest{}
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *AddSubscriptionRequest) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSubscriptionRequest) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *AddSubscriptionRequest) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *AddSubscriptionRequest) SetEnable(v bool) {
	o.Enable = &v
}

// GetSubscription returns the Subscription field value
func (o *AddSubscriptionRequest) GetSubscription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *AddSubscriptionRequest) GetSubscriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *AddSubscriptionRequest) SetSubscription(v string) {
	o.Subscription = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AddSubscriptionRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSubscriptionRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AddSubscriptionRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AddSubscriptionRequest) SetType(v string) {
	o.Type = &v
}

func (o AddSubscriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enable != nil {
		toSerialize["enable"] = o.Enable
	}
	if true {
		toSerialize["subscription"] = o.Subscription
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableAddSubscriptionRequest struct {
	value *AddSubscriptionRequest
	isSet bool
}

func (v NullableAddSubscriptionRequest) Get() *AddSubscriptionRequest {
	return v.value
}

func (v *NullableAddSubscriptionRequest) Set(val *AddSubscriptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSubscriptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSubscriptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSubscriptionRequest(val *AddSubscriptionRequest) *NullableAddSubscriptionRequest {
	return &NullableAddSubscriptionRequest{value: val, isSet: true}
}

func (v NullableAddSubscriptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSubscriptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

