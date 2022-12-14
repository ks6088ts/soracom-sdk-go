/*
SORACOM SANDBOX API

SORACOM SANDBOX API v1

API version: 20160218
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sandbox

import (
	"encoding/json"
)

// SandboxInitRequest struct for SandboxInitRequest
type SandboxInitRequest struct {
	AuthKey string `json:"authKey"`
	AuthKeyId string `json:"authKeyId"`
	// カバレッジタイプ - `g`: グローバルカバレッジ - `jp`: 日本カバレッジ 
	CoverageTypes []string `json:"coverageTypes,omitempty"`
	Email string `json:"email"`
	Password string `json:"password"`
	RegisterPaymentMethod *bool `json:"registerPaymentMethod,omitempty"`
}

// NewSandboxInitRequest instantiates a new SandboxInitRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxInitRequest(authKey string, authKeyId string, email string, password string) *SandboxInitRequest {
	this := SandboxInitRequest{}
	this.AuthKey = authKey
	this.AuthKeyId = authKeyId
	this.Email = email
	this.Password = password
	var registerPaymentMethod bool = true
	this.RegisterPaymentMethod = &registerPaymentMethod
	return &this
}

// NewSandboxInitRequestWithDefaults instantiates a new SandboxInitRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxInitRequestWithDefaults() *SandboxInitRequest {
	this := SandboxInitRequest{}
	var registerPaymentMethod bool = true
	this.RegisterPaymentMethod = &registerPaymentMethod
	return &this
}

// GetAuthKey returns the AuthKey field value
func (o *SandboxInitRequest) GetAuthKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthKey
}

// GetAuthKeyOk returns a tuple with the AuthKey field value
// and a boolean to check if the value has been set.
func (o *SandboxInitRequest) GetAuthKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthKey, true
}

// SetAuthKey sets field value
func (o *SandboxInitRequest) SetAuthKey(v string) {
	o.AuthKey = v
}

// GetAuthKeyId returns the AuthKeyId field value
func (o *SandboxInitRequest) GetAuthKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthKeyId
}

// GetAuthKeyIdOk returns a tuple with the AuthKeyId field value
// and a boolean to check if the value has been set.
func (o *SandboxInitRequest) GetAuthKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthKeyId, true
}

// SetAuthKeyId sets field value
func (o *SandboxInitRequest) SetAuthKeyId(v string) {
	o.AuthKeyId = v
}

// GetCoverageTypes returns the CoverageTypes field value if set, zero value otherwise.
func (o *SandboxInitRequest) GetCoverageTypes() []string {
	if o == nil || o.CoverageTypes == nil {
		var ret []string
		return ret
	}
	return o.CoverageTypes
}

// GetCoverageTypesOk returns a tuple with the CoverageTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxInitRequest) GetCoverageTypesOk() ([]string, bool) {
	if o == nil || o.CoverageTypes == nil {
		return nil, false
	}
	return o.CoverageTypes, true
}

// HasCoverageTypes returns a boolean if a field has been set.
func (o *SandboxInitRequest) HasCoverageTypes() bool {
	if o != nil && o.CoverageTypes != nil {
		return true
	}

	return false
}

// SetCoverageTypes gets a reference to the given []string and assigns it to the CoverageTypes field.
func (o *SandboxInitRequest) SetCoverageTypes(v []string) {
	o.CoverageTypes = v
}

// GetEmail returns the Email field value
func (o *SandboxInitRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SandboxInitRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SandboxInitRequest) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *SandboxInitRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SandboxInitRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SandboxInitRequest) SetPassword(v string) {
	o.Password = v
}

// GetRegisterPaymentMethod returns the RegisterPaymentMethod field value if set, zero value otherwise.
func (o *SandboxInitRequest) GetRegisterPaymentMethod() bool {
	if o == nil || o.RegisterPaymentMethod == nil {
		var ret bool
		return ret
	}
	return *o.RegisterPaymentMethod
}

// GetRegisterPaymentMethodOk returns a tuple with the RegisterPaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxInitRequest) GetRegisterPaymentMethodOk() (*bool, bool) {
	if o == nil || o.RegisterPaymentMethod == nil {
		return nil, false
	}
	return o.RegisterPaymentMethod, true
}

// HasRegisterPaymentMethod returns a boolean if a field has been set.
func (o *SandboxInitRequest) HasRegisterPaymentMethod() bool {
	if o != nil && o.RegisterPaymentMethod != nil {
		return true
	}

	return false
}

// SetRegisterPaymentMethod gets a reference to the given bool and assigns it to the RegisterPaymentMethod field.
func (o *SandboxInitRequest) SetRegisterPaymentMethod(v bool) {
	o.RegisterPaymentMethod = &v
}

func (o SandboxInitRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["authKey"] = o.AuthKey
	}
	if true {
		toSerialize["authKeyId"] = o.AuthKeyId
	}
	if o.CoverageTypes != nil {
		toSerialize["coverageTypes"] = o.CoverageTypes
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if o.RegisterPaymentMethod != nil {
		toSerialize["registerPaymentMethod"] = o.RegisterPaymentMethod
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxInitRequest struct {
	value *SandboxInitRequest
	isSet bool
}

func (v NullableSandboxInitRequest) Get() *SandboxInitRequest {
	return v.value
}

func (v *NullableSandboxInitRequest) Set(val *SandboxInitRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxInitRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxInitRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxInitRequest(val *SandboxInitRequest) *NullableSandboxInitRequest {
	return &NullableSandboxInitRequest{value: val, isSet: true}
}

func (v NullableSandboxInitRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxInitRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


