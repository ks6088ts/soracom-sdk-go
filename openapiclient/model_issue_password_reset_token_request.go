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

// IssuePasswordResetTokenRequest struct for IssuePasswordResetTokenRequest
type IssuePasswordResetTokenRequest struct {
	Email string `json:"email"`
}

// NewIssuePasswordResetTokenRequest instantiates a new IssuePasswordResetTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuePasswordResetTokenRequest(email string) *IssuePasswordResetTokenRequest {
	this := IssuePasswordResetTokenRequest{}
	this.Email = email
	return &this
}

// NewIssuePasswordResetTokenRequestWithDefaults instantiates a new IssuePasswordResetTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuePasswordResetTokenRequestWithDefaults() *IssuePasswordResetTokenRequest {
	this := IssuePasswordResetTokenRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *IssuePasswordResetTokenRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *IssuePasswordResetTokenRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *IssuePasswordResetTokenRequest) SetEmail(v string) {
	o.Email = v
}

func (o IssuePasswordResetTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableIssuePasswordResetTokenRequest struct {
	value *IssuePasswordResetTokenRequest
	isSet bool
}

func (v NullableIssuePasswordResetTokenRequest) Get() *IssuePasswordResetTokenRequest {
	return v.value
}

func (v *NullableIssuePasswordResetTokenRequest) Set(val *IssuePasswordResetTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuePasswordResetTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuePasswordResetTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuePasswordResetTokenRequest(val *IssuePasswordResetTokenRequest) *NullableIssuePasswordResetTokenRequest {
	return &NullableIssuePasswordResetTokenRequest{value: val, isSet: true}
}

func (v NullableIssuePasswordResetTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuePasswordResetTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


