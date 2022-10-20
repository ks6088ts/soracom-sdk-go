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

// ContractUpdatingRequest struct for ContractUpdatingRequest
type ContractUpdatingRequest struct {
	ContractDetail map[string]interface{} `json:"contractDetail,omitempty"`
	ContractName *string `json:"contractName,omitempty"`
}

// NewContractUpdatingRequest instantiates a new ContractUpdatingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractUpdatingRequest() *ContractUpdatingRequest {
	this := ContractUpdatingRequest{}
	return &this
}

// NewContractUpdatingRequestWithDefaults instantiates a new ContractUpdatingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractUpdatingRequestWithDefaults() *ContractUpdatingRequest {
	this := ContractUpdatingRequest{}
	return &this
}

// GetContractDetail returns the ContractDetail field value if set, zero value otherwise.
func (o *ContractUpdatingRequest) GetContractDetail() map[string]interface{} {
	if o == nil || o.ContractDetail == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ContractDetail
}

// GetContractDetailOk returns a tuple with the ContractDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUpdatingRequest) GetContractDetailOk() (map[string]interface{}, bool) {
	if o == nil || o.ContractDetail == nil {
		return nil, false
	}
	return o.ContractDetail, true
}

// HasContractDetail returns a boolean if a field has been set.
func (o *ContractUpdatingRequest) HasContractDetail() bool {
	if o != nil && o.ContractDetail != nil {
		return true
	}

	return false
}

// SetContractDetail gets a reference to the given map[string]interface{} and assigns it to the ContractDetail field.
func (o *ContractUpdatingRequest) SetContractDetail(v map[string]interface{}) {
	o.ContractDetail = v
}

// GetContractName returns the ContractName field value if set, zero value otherwise.
func (o *ContractUpdatingRequest) GetContractName() string {
	if o == nil || o.ContractName == nil {
		var ret string
		return ret
	}
	return *o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUpdatingRequest) GetContractNameOk() (*string, bool) {
	if o == nil || o.ContractName == nil {
		return nil, false
	}
	return o.ContractName, true
}

// HasContractName returns a boolean if a field has been set.
func (o *ContractUpdatingRequest) HasContractName() bool {
	if o != nil && o.ContractName != nil {
		return true
	}

	return false
}

// SetContractName gets a reference to the given string and assigns it to the ContractName field.
func (o *ContractUpdatingRequest) SetContractName(v string) {
	o.ContractName = &v
}

func (o ContractUpdatingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContractDetail != nil {
		toSerialize["contractDetail"] = o.ContractDetail
	}
	if o.ContractName != nil {
		toSerialize["contractName"] = o.ContractName
	}
	return json.Marshal(toSerialize)
}

type NullableContractUpdatingRequest struct {
	value *ContractUpdatingRequest
	isSet bool
}

func (v NullableContractUpdatingRequest) Get() *ContractUpdatingRequest {
	return v.value
}

func (v *NullableContractUpdatingRequest) Set(val *ContractUpdatingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContractUpdatingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContractUpdatingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractUpdatingRequest(val *ContractUpdatingRequest) *NullableContractUpdatingRequest {
	return &NullableContractUpdatingRequest{value: val, isSet: true}
}

func (v NullableContractUpdatingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractUpdatingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

