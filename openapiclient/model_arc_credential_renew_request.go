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

// ArcCredentialRenewRequest struct for ArcCredentialRenewRequest
type ArcCredentialRenewRequest struct {
	// SIM に付与する Arc クライアントの公開鍵。もしこの値が空の場合はサーバーがキーペアを生成します。
	ArcClientPeerPublicKey *string `json:"arcClientPeerPublicKey,omitempty"`
}

// NewArcCredentialRenewRequest instantiates a new ArcCredentialRenewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcCredentialRenewRequest() *ArcCredentialRenewRequest {
	this := ArcCredentialRenewRequest{}
	return &this
}

// NewArcCredentialRenewRequestWithDefaults instantiates a new ArcCredentialRenewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcCredentialRenewRequestWithDefaults() *ArcCredentialRenewRequest {
	this := ArcCredentialRenewRequest{}
	return &this
}

// GetArcClientPeerPublicKey returns the ArcClientPeerPublicKey field value if set, zero value otherwise.
func (o *ArcCredentialRenewRequest) GetArcClientPeerPublicKey() string {
	if o == nil || o.ArcClientPeerPublicKey == nil {
		var ret string
		return ret
	}
	return *o.ArcClientPeerPublicKey
}

// GetArcClientPeerPublicKeyOk returns a tuple with the ArcClientPeerPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcCredentialRenewRequest) GetArcClientPeerPublicKeyOk() (*string, bool) {
	if o == nil || o.ArcClientPeerPublicKey == nil {
		return nil, false
	}
	return o.ArcClientPeerPublicKey, true
}

// HasArcClientPeerPublicKey returns a boolean if a field has been set.
func (o *ArcCredentialRenewRequest) HasArcClientPeerPublicKey() bool {
	if o != nil && o.ArcClientPeerPublicKey != nil {
		return true
	}

	return false
}

// SetArcClientPeerPublicKey gets a reference to the given string and assigns it to the ArcClientPeerPublicKey field.
func (o *ArcCredentialRenewRequest) SetArcClientPeerPublicKey(v string) {
	o.ArcClientPeerPublicKey = &v
}

func (o ArcCredentialRenewRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArcClientPeerPublicKey != nil {
		toSerialize["arcClientPeerPublicKey"] = o.ArcClientPeerPublicKey
	}
	return json.Marshal(toSerialize)
}

type NullableArcCredentialRenewRequest struct {
	value *ArcCredentialRenewRequest
	isSet bool
}

func (v NullableArcCredentialRenewRequest) Get() *ArcCredentialRenewRequest {
	return v.value
}

func (v *NullableArcCredentialRenewRequest) Set(val *ArcCredentialRenewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableArcCredentialRenewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableArcCredentialRenewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArcCredentialRenewRequest(val *ArcCredentialRenewRequest) *NullableArcCredentialRenewRequest {
	return &NullableArcCredentialRenewRequest{value: val, isSet: true}
}

func (v NullableArcCredentialRenewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArcCredentialRenewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

