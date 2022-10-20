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

// SupportTokenResponse struct for SupportTokenResponse
type SupportTokenResponse struct {
	Token string `json:"token"`
}

// NewSupportTokenResponse instantiates a new SupportTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportTokenResponse(token string) *SupportTokenResponse {
	this := SupportTokenResponse{}
	this.Token = token
	return &this
}

// NewSupportTokenResponseWithDefaults instantiates a new SupportTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportTokenResponseWithDefaults() *SupportTokenResponse {
	this := SupportTokenResponse{}
	return &this
}

// GetToken returns the Token field value
func (o *SupportTokenResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *SupportTokenResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *SupportTokenResponse) SetToken(v string) {
	o.Token = v
}

func (o SupportTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableSupportTokenResponse struct {
	value *SupportTokenResponse
	isSet bool
}

func (v NullableSupportTokenResponse) Get() *SupportTokenResponse {
	return v.value
}

func (v *NullableSupportTokenResponse) Set(val *SupportTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportTokenResponse(val *SupportTokenResponse) *NullableSupportTokenResponse {
	return &NullableSupportTokenResponse{value: val, isSet: true}
}

func (v NullableSupportTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

