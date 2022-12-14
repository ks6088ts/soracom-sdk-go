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

// RegisterOperatorsRequest struct for RegisterOperatorsRequest
type RegisterOperatorsRequest struct {
	Email string `json:"email"`
	// パスワードは以下の条件を満たしている必要があります：長さ 8 文字以上 100 文字以内、アルファベット小文字 (a-z) を 1 文字以上使用、アルファベット大文字 (A-Z) を 1 文字以上使用、数字を 1 文字以上使用。記号なども使用できます。
	Password string `json:"password"`
}

// NewRegisterOperatorsRequest instantiates a new RegisterOperatorsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterOperatorsRequest(email string, password string) *RegisterOperatorsRequest {
	this := RegisterOperatorsRequest{}
	this.Email = email
	this.Password = password
	return &this
}

// NewRegisterOperatorsRequestWithDefaults instantiates a new RegisterOperatorsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterOperatorsRequestWithDefaults() *RegisterOperatorsRequest {
	this := RegisterOperatorsRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *RegisterOperatorsRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *RegisterOperatorsRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *RegisterOperatorsRequest) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *RegisterOperatorsRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *RegisterOperatorsRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *RegisterOperatorsRequest) SetPassword(v string) {
	o.Password = v
}

func (o RegisterOperatorsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableRegisterOperatorsRequest struct {
	value *RegisterOperatorsRequest
	isSet bool
}

func (v NullableRegisterOperatorsRequest) Get() *RegisterOperatorsRequest {
	return v.value
}

func (v *NullableRegisterOperatorsRequest) Set(val *RegisterOperatorsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterOperatorsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterOperatorsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterOperatorsRequest(val *RegisterOperatorsRequest) *NullableRegisterOperatorsRequest {
	return &NullableRegisterOperatorsRequest{value: val, isSet: true}
}

func (v NullableRegisterOperatorsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterOperatorsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


