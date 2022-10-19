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

// CreateAndUpdateCredentialsModel struct for CreateAndUpdateCredentialsModel
type CreateAndUpdateCredentialsModel struct {
	Credentials map[string]interface{} `json:"credentials,omitempty"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewCreateAndUpdateCredentialsModel instantiates a new CreateAndUpdateCredentialsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAndUpdateCredentialsModel() *CreateAndUpdateCredentialsModel {
	this := CreateAndUpdateCredentialsModel{}
	return &this
}

// NewCreateAndUpdateCredentialsModelWithDefaults instantiates a new CreateAndUpdateCredentialsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAndUpdateCredentialsModelWithDefaults() *CreateAndUpdateCredentialsModel {
	this := CreateAndUpdateCredentialsModel{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *CreateAndUpdateCredentialsModel) GetCredentials() map[string]interface{} {
	if o == nil || o.Credentials == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAndUpdateCredentialsModel) GetCredentialsOk() (map[string]interface{}, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *CreateAndUpdateCredentialsModel) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given map[string]interface{} and assigns it to the Credentials field.
func (o *CreateAndUpdateCredentialsModel) SetCredentials(v map[string]interface{}) {
	o.Credentials = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateAndUpdateCredentialsModel) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAndUpdateCredentialsModel) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateAndUpdateCredentialsModel) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateAndUpdateCredentialsModel) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateAndUpdateCredentialsModel) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAndUpdateCredentialsModel) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateAndUpdateCredentialsModel) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateAndUpdateCredentialsModel) SetType(v string) {
	o.Type = &v
}

func (o CreateAndUpdateCredentialsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAndUpdateCredentialsModel struct {
	value *CreateAndUpdateCredentialsModel
	isSet bool
}

func (v NullableCreateAndUpdateCredentialsModel) Get() *CreateAndUpdateCredentialsModel {
	return v.value
}

func (v *NullableCreateAndUpdateCredentialsModel) Set(val *CreateAndUpdateCredentialsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAndUpdateCredentialsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAndUpdateCredentialsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAndUpdateCredentialsModel(val *CreateAndUpdateCredentialsModel) *NullableCreateAndUpdateCredentialsModel {
	return &NullableCreateAndUpdateCredentialsModel{value: val, isSet: true}
}

func (v NullableCreateAndUpdateCredentialsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAndUpdateCredentialsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

