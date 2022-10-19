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

// GenerateUserAuthKeyResponse struct for GenerateUserAuthKeyResponse
type GenerateUserAuthKeyResponse struct {
	AuthKey *string `json:"authKey,omitempty"`
	AuthKeyId *string `json:"authKeyId,omitempty"`
}

// NewGenerateUserAuthKeyResponse instantiates a new GenerateUserAuthKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateUserAuthKeyResponse() *GenerateUserAuthKeyResponse {
	this := GenerateUserAuthKeyResponse{}
	return &this
}

// NewGenerateUserAuthKeyResponseWithDefaults instantiates a new GenerateUserAuthKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateUserAuthKeyResponseWithDefaults() *GenerateUserAuthKeyResponse {
	this := GenerateUserAuthKeyResponse{}
	return &this
}

// GetAuthKey returns the AuthKey field value if set, zero value otherwise.
func (o *GenerateUserAuthKeyResponse) GetAuthKey() string {
	if o == nil || o.AuthKey == nil {
		var ret string
		return ret
	}
	return *o.AuthKey
}

// GetAuthKeyOk returns a tuple with the AuthKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateUserAuthKeyResponse) GetAuthKeyOk() (*string, bool) {
	if o == nil || o.AuthKey == nil {
		return nil, false
	}
	return o.AuthKey, true
}

// HasAuthKey returns a boolean if a field has been set.
func (o *GenerateUserAuthKeyResponse) HasAuthKey() bool {
	if o != nil && o.AuthKey != nil {
		return true
	}

	return false
}

// SetAuthKey gets a reference to the given string and assigns it to the AuthKey field.
func (o *GenerateUserAuthKeyResponse) SetAuthKey(v string) {
	o.AuthKey = &v
}

// GetAuthKeyId returns the AuthKeyId field value if set, zero value otherwise.
func (o *GenerateUserAuthKeyResponse) GetAuthKeyId() string {
	if o == nil || o.AuthKeyId == nil {
		var ret string
		return ret
	}
	return *o.AuthKeyId
}

// GetAuthKeyIdOk returns a tuple with the AuthKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateUserAuthKeyResponse) GetAuthKeyIdOk() (*string, bool) {
	if o == nil || o.AuthKeyId == nil {
		return nil, false
	}
	return o.AuthKeyId, true
}

// HasAuthKeyId returns a boolean if a field has been set.
func (o *GenerateUserAuthKeyResponse) HasAuthKeyId() bool {
	if o != nil && o.AuthKeyId != nil {
		return true
	}

	return false
}

// SetAuthKeyId gets a reference to the given string and assigns it to the AuthKeyId field.
func (o *GenerateUserAuthKeyResponse) SetAuthKeyId(v string) {
	o.AuthKeyId = &v
}

func (o GenerateUserAuthKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthKey != nil {
		toSerialize["authKey"] = o.AuthKey
	}
	if o.AuthKeyId != nil {
		toSerialize["authKeyId"] = o.AuthKeyId
	}
	return json.Marshal(toSerialize)
}

type NullableGenerateUserAuthKeyResponse struct {
	value *GenerateUserAuthKeyResponse
	isSet bool
}

func (v NullableGenerateUserAuthKeyResponse) Get() *GenerateUserAuthKeyResponse {
	return v.value
}

func (v *NullableGenerateUserAuthKeyResponse) Set(val *GenerateUserAuthKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateUserAuthKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateUserAuthKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateUserAuthKeyResponse(val *GenerateUserAuthKeyResponse) *NullableGenerateUserAuthKeyResponse {
	return &NullableGenerateUserAuthKeyResponse{value: val, isSet: true}
}

func (v NullableGenerateUserAuthKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateUserAuthKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


