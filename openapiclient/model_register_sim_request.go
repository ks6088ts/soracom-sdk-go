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

// RegisterSimRequest struct for RegisterSimRequest
type RegisterSimRequest struct {
	GroupId *string `json:"groupId,omitempty"`
	// PUK or PASSCODE on SIM card
	RegistrationSecret string `json:"registrationSecret"`
	// An object which always contains at least one property \"name\" with a string value. If you give a subscriber/SIM a name, the name will be returned as the value of the \"name\" property. If the subscriber/SIM does not have a name, an empty string \"\" is returned. In addition, if you create any custom tags for the subscriber/SIM, each custom tag will appear as additional properties in the object.
	Tags *map[string]string `json:"tags,omitempty"`
}

// NewRegisterSimRequest instantiates a new RegisterSimRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterSimRequest(registrationSecret string) *RegisterSimRequest {
	this := RegisterSimRequest{}
	this.RegistrationSecret = registrationSecret
	return &this
}

// NewRegisterSimRequestWithDefaults instantiates a new RegisterSimRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterSimRequestWithDefaults() *RegisterSimRequest {
	this := RegisterSimRequest{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *RegisterSimRequest) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterSimRequest) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *RegisterSimRequest) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *RegisterSimRequest) SetGroupId(v string) {
	o.GroupId = &v
}

// GetRegistrationSecret returns the RegistrationSecret field value
func (o *RegisterSimRequest) GetRegistrationSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegistrationSecret
}

// GetRegistrationSecretOk returns a tuple with the RegistrationSecret field value
// and a boolean to check if the value has been set.
func (o *RegisterSimRequest) GetRegistrationSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegistrationSecret, true
}

// SetRegistrationSecret sets field value
func (o *RegisterSimRequest) SetRegistrationSecret(v string) {
	o.RegistrationSecret = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RegisterSimRequest) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterSimRequest) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RegisterSimRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *RegisterSimRequest) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o RegisterSimRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if true {
		toSerialize["registrationSecret"] = o.RegistrationSecret
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableRegisterSimRequest struct {
	value *RegisterSimRequest
	isSet bool
}

func (v NullableRegisterSimRequest) Get() *RegisterSimRequest {
	return v.value
}

func (v *NullableRegisterSimRequest) Set(val *RegisterSimRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterSimRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterSimRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterSimRequest(val *RegisterSimRequest) *NullableRegisterSimRequest {
	return &NullableRegisterSimRequest{value: val, isSet: true}
}

func (v NullableRegisterSimRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterSimRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

