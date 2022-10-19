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

// LagoonUserEmailUpdatingRequest struct for LagoonUserEmailUpdatingRequest
type LagoonUserEmailUpdatingRequest struct {
	UserEmail *string `json:"userEmail,omitempty"`
}

// NewLagoonUserEmailUpdatingRequest instantiates a new LagoonUserEmailUpdatingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLagoonUserEmailUpdatingRequest() *LagoonUserEmailUpdatingRequest {
	this := LagoonUserEmailUpdatingRequest{}
	return &this
}

// NewLagoonUserEmailUpdatingRequestWithDefaults instantiates a new LagoonUserEmailUpdatingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLagoonUserEmailUpdatingRequestWithDefaults() *LagoonUserEmailUpdatingRequest {
	this := LagoonUserEmailUpdatingRequest{}
	return &this
}

// GetUserEmail returns the UserEmail field value if set, zero value otherwise.
func (o *LagoonUserEmailUpdatingRequest) GetUserEmail() string {
	if o == nil || o.UserEmail == nil {
		var ret string
		return ret
	}
	return *o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LagoonUserEmailUpdatingRequest) GetUserEmailOk() (*string, bool) {
	if o == nil || o.UserEmail == nil {
		return nil, false
	}
	return o.UserEmail, true
}

// HasUserEmail returns a boolean if a field has been set.
func (o *LagoonUserEmailUpdatingRequest) HasUserEmail() bool {
	if o != nil && o.UserEmail != nil {
		return true
	}

	return false
}

// SetUserEmail gets a reference to the given string and assigns it to the UserEmail field.
func (o *LagoonUserEmailUpdatingRequest) SetUserEmail(v string) {
	o.UserEmail = &v
}

func (o LagoonUserEmailUpdatingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserEmail != nil {
		toSerialize["userEmail"] = o.UserEmail
	}
	return json.Marshal(toSerialize)
}

type NullableLagoonUserEmailUpdatingRequest struct {
	value *LagoonUserEmailUpdatingRequest
	isSet bool
}

func (v NullableLagoonUserEmailUpdatingRequest) Get() *LagoonUserEmailUpdatingRequest {
	return v.value
}

func (v *NullableLagoonUserEmailUpdatingRequest) Set(val *LagoonUserEmailUpdatingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLagoonUserEmailUpdatingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLagoonUserEmailUpdatingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLagoonUserEmailUpdatingRequest(val *LagoonUserEmailUpdatingRequest) *NullableLagoonUserEmailUpdatingRequest {
	return &NullableLagoonUserEmailUpdatingRequest{value: val, isSet: true}
}

func (v NullableLagoonUserEmailUpdatingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLagoonUserEmailUpdatingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

