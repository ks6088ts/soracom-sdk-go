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

// LagoonLicensePacksUpdatingRequest struct for LagoonLicensePacksUpdatingRequest
type LagoonLicensePacksUpdatingRequest struct {
	LicensePackQuantities []LagoonLicensePacksUpdatingRequestLicensePackQuantitiesInner `json:"licensePackQuantities,omitempty"`
}

// NewLagoonLicensePacksUpdatingRequest instantiates a new LagoonLicensePacksUpdatingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLagoonLicensePacksUpdatingRequest() *LagoonLicensePacksUpdatingRequest {
	this := LagoonLicensePacksUpdatingRequest{}
	return &this
}

// NewLagoonLicensePacksUpdatingRequestWithDefaults instantiates a new LagoonLicensePacksUpdatingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLagoonLicensePacksUpdatingRequestWithDefaults() *LagoonLicensePacksUpdatingRequest {
	this := LagoonLicensePacksUpdatingRequest{}
	return &this
}

// GetLicensePackQuantities returns the LicensePackQuantities field value if set, zero value otherwise.
func (o *LagoonLicensePacksUpdatingRequest) GetLicensePackQuantities() []LagoonLicensePacksUpdatingRequestLicensePackQuantitiesInner {
	if o == nil || o.LicensePackQuantities == nil {
		var ret []LagoonLicensePacksUpdatingRequestLicensePackQuantitiesInner
		return ret
	}
	return o.LicensePackQuantities
}

// GetLicensePackQuantitiesOk returns a tuple with the LicensePackQuantities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LagoonLicensePacksUpdatingRequest) GetLicensePackQuantitiesOk() ([]LagoonLicensePacksUpdatingRequestLicensePackQuantitiesInner, bool) {
	if o == nil || o.LicensePackQuantities == nil {
		return nil, false
	}
	return o.LicensePackQuantities, true
}

// HasLicensePackQuantities returns a boolean if a field has been set.
func (o *LagoonLicensePacksUpdatingRequest) HasLicensePackQuantities() bool {
	if o != nil && o.LicensePackQuantities != nil {
		return true
	}

	return false
}

// SetLicensePackQuantities gets a reference to the given []LagoonLicensePacksUpdatingRequestLicensePackQuantitiesInner and assigns it to the LicensePackQuantities field.
func (o *LagoonLicensePacksUpdatingRequest) SetLicensePackQuantities(v []LagoonLicensePacksUpdatingRequestLicensePackQuantitiesInner) {
	o.LicensePackQuantities = v
}

func (o LagoonLicensePacksUpdatingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LicensePackQuantities != nil {
		toSerialize["licensePackQuantities"] = o.LicensePackQuantities
	}
	return json.Marshal(toSerialize)
}

type NullableLagoonLicensePacksUpdatingRequest struct {
	value *LagoonLicensePacksUpdatingRequest
	isSet bool
}

func (v NullableLagoonLicensePacksUpdatingRequest) Get() *LagoonLicensePacksUpdatingRequest {
	return v.value
}

func (v *NullableLagoonLicensePacksUpdatingRequest) Set(val *LagoonLicensePacksUpdatingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLagoonLicensePacksUpdatingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLagoonLicensePacksUpdatingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLagoonLicensePacksUpdatingRequest(val *LagoonLicensePacksUpdatingRequest) *NullableLagoonLicensePacksUpdatingRequest {
	return &NullableLagoonLicensePacksUpdatingRequest{value: val, isSet: true}
}

func (v NullableLagoonLicensePacksUpdatingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLagoonLicensePacksUpdatingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

