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

// FunnelBrainsTechImpulseDestination struct for FunnelBrainsTechImpulseDestination
type FunnelBrainsTechImpulseDestination struct {
	Provider *string `json:"provider,omitempty"`
	ResourceUrl *string `json:"resourceUrl,omitempty"`
	Service *string `json:"service,omitempty"`
}

// NewFunnelBrainsTechImpulseDestination instantiates a new FunnelBrainsTechImpulseDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunnelBrainsTechImpulseDestination() *FunnelBrainsTechImpulseDestination {
	this := FunnelBrainsTechImpulseDestination{}
	return &this
}

// NewFunnelBrainsTechImpulseDestinationWithDefaults instantiates a new FunnelBrainsTechImpulseDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunnelBrainsTechImpulseDestinationWithDefaults() *FunnelBrainsTechImpulseDestination {
	this := FunnelBrainsTechImpulseDestination{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *FunnelBrainsTechImpulseDestination) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelBrainsTechImpulseDestination) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *FunnelBrainsTechImpulseDestination) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *FunnelBrainsTechImpulseDestination) SetProvider(v string) {
	o.Provider = &v
}

// GetResourceUrl returns the ResourceUrl field value if set, zero value otherwise.
func (o *FunnelBrainsTechImpulseDestination) GetResourceUrl() string {
	if o == nil || o.ResourceUrl == nil {
		var ret string
		return ret
	}
	return *o.ResourceUrl
}

// GetResourceUrlOk returns a tuple with the ResourceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelBrainsTechImpulseDestination) GetResourceUrlOk() (*string, bool) {
	if o == nil || o.ResourceUrl == nil {
		return nil, false
	}
	return o.ResourceUrl, true
}

// HasResourceUrl returns a boolean if a field has been set.
func (o *FunnelBrainsTechImpulseDestination) HasResourceUrl() bool {
	if o != nil && o.ResourceUrl != nil {
		return true
	}

	return false
}

// SetResourceUrl gets a reference to the given string and assigns it to the ResourceUrl field.
func (o *FunnelBrainsTechImpulseDestination) SetResourceUrl(v string) {
	o.ResourceUrl = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *FunnelBrainsTechImpulseDestination) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelBrainsTechImpulseDestination) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *FunnelBrainsTechImpulseDestination) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *FunnelBrainsTechImpulseDestination) SetService(v string) {
	o.Service = &v
}

func (o FunnelBrainsTechImpulseDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.ResourceUrl != nil {
		toSerialize["resourceUrl"] = o.ResourceUrl
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	return json.Marshal(toSerialize)
}

type NullableFunnelBrainsTechImpulseDestination struct {
	value *FunnelBrainsTechImpulseDestination
	isSet bool
}

func (v NullableFunnelBrainsTechImpulseDestination) Get() *FunnelBrainsTechImpulseDestination {
	return v.value
}

func (v *NullableFunnelBrainsTechImpulseDestination) Set(val *FunnelBrainsTechImpulseDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableFunnelBrainsTechImpulseDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableFunnelBrainsTechImpulseDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunnelBrainsTechImpulseDestination(val *FunnelBrainsTechImpulseDestination) *NullableFunnelBrainsTechImpulseDestination {
	return &NullableFunnelBrainsTechImpulseDestination{value: val, isSet: true}
}

func (v NullableFunnelBrainsTechImpulseDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunnelBrainsTechImpulseDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


