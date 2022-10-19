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

// RegisterLoraDeviceRequest struct for RegisterLoraDeviceRequest
type RegisterLoraDeviceRequest struct {
	GroupId *string `json:"groupId,omitempty"`
	RegistrationSecret *string `json:"registrationSecret,omitempty"`
	Tags *map[string]string `json:"tags,omitempty"`
}

// NewRegisterLoraDeviceRequest instantiates a new RegisterLoraDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterLoraDeviceRequest() *RegisterLoraDeviceRequest {
	this := RegisterLoraDeviceRequest{}
	return &this
}

// NewRegisterLoraDeviceRequestWithDefaults instantiates a new RegisterLoraDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterLoraDeviceRequestWithDefaults() *RegisterLoraDeviceRequest {
	this := RegisterLoraDeviceRequest{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *RegisterLoraDeviceRequest) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterLoraDeviceRequest) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *RegisterLoraDeviceRequest) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *RegisterLoraDeviceRequest) SetGroupId(v string) {
	o.GroupId = &v
}

// GetRegistrationSecret returns the RegistrationSecret field value if set, zero value otherwise.
func (o *RegisterLoraDeviceRequest) GetRegistrationSecret() string {
	if o == nil || o.RegistrationSecret == nil {
		var ret string
		return ret
	}
	return *o.RegistrationSecret
}

// GetRegistrationSecretOk returns a tuple with the RegistrationSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterLoraDeviceRequest) GetRegistrationSecretOk() (*string, bool) {
	if o == nil || o.RegistrationSecret == nil {
		return nil, false
	}
	return o.RegistrationSecret, true
}

// HasRegistrationSecret returns a boolean if a field has been set.
func (o *RegisterLoraDeviceRequest) HasRegistrationSecret() bool {
	if o != nil && o.RegistrationSecret != nil {
		return true
	}

	return false
}

// SetRegistrationSecret gets a reference to the given string and assigns it to the RegistrationSecret field.
func (o *RegisterLoraDeviceRequest) SetRegistrationSecret(v string) {
	o.RegistrationSecret = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RegisterLoraDeviceRequest) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterLoraDeviceRequest) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RegisterLoraDeviceRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *RegisterLoraDeviceRequest) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o RegisterLoraDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.RegistrationSecret != nil {
		toSerialize["registrationSecret"] = o.RegistrationSecret
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableRegisterLoraDeviceRequest struct {
	value *RegisterLoraDeviceRequest
	isSet bool
}

func (v NullableRegisterLoraDeviceRequest) Get() *RegisterLoraDeviceRequest {
	return v.value
}

func (v *NullableRegisterLoraDeviceRequest) Set(val *RegisterLoraDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterLoraDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterLoraDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterLoraDeviceRequest(val *RegisterLoraDeviceRequest) *NullableRegisterLoraDeviceRequest {
	return &NullableRegisterLoraDeviceRequest{value: val, isSet: true}
}

func (v NullableRegisterLoraDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterLoraDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

