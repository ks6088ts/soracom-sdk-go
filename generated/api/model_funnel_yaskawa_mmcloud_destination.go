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

// FunnelYaskawaMmcloudDestination struct for FunnelYaskawaMmcloudDestination
type FunnelYaskawaMmcloudDestination struct {
	Desthost *string `json:"desthost,omitempty"`
	Destport *int32 `json:"destport,omitempty"`
	Provider *string `json:"provider,omitempty"`
	ResourceUrl *string `json:"resourceUrl,omitempty"`
	Service *string `json:"service,omitempty"`
}

// NewFunnelYaskawaMmcloudDestination instantiates a new FunnelYaskawaMmcloudDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunnelYaskawaMmcloudDestination() *FunnelYaskawaMmcloudDestination {
	this := FunnelYaskawaMmcloudDestination{}
	return &this
}

// NewFunnelYaskawaMmcloudDestinationWithDefaults instantiates a new FunnelYaskawaMmcloudDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunnelYaskawaMmcloudDestinationWithDefaults() *FunnelYaskawaMmcloudDestination {
	this := FunnelYaskawaMmcloudDestination{}
	return &this
}

// GetDesthost returns the Desthost field value if set, zero value otherwise.
func (o *FunnelYaskawaMmcloudDestination) GetDesthost() string {
	if o == nil || o.Desthost == nil {
		var ret string
		return ret
	}
	return *o.Desthost
}

// GetDesthostOk returns a tuple with the Desthost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelYaskawaMmcloudDestination) GetDesthostOk() (*string, bool) {
	if o == nil || o.Desthost == nil {
		return nil, false
	}
	return o.Desthost, true
}

// HasDesthost returns a boolean if a field has been set.
func (o *FunnelYaskawaMmcloudDestination) HasDesthost() bool {
	if o != nil && o.Desthost != nil {
		return true
	}

	return false
}

// SetDesthost gets a reference to the given string and assigns it to the Desthost field.
func (o *FunnelYaskawaMmcloudDestination) SetDesthost(v string) {
	o.Desthost = &v
}

// GetDestport returns the Destport field value if set, zero value otherwise.
func (o *FunnelYaskawaMmcloudDestination) GetDestport() int32 {
	if o == nil || o.Destport == nil {
		var ret int32
		return ret
	}
	return *o.Destport
}

// GetDestportOk returns a tuple with the Destport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelYaskawaMmcloudDestination) GetDestportOk() (*int32, bool) {
	if o == nil || o.Destport == nil {
		return nil, false
	}
	return o.Destport, true
}

// HasDestport returns a boolean if a field has been set.
func (o *FunnelYaskawaMmcloudDestination) HasDestport() bool {
	if o != nil && o.Destport != nil {
		return true
	}

	return false
}

// SetDestport gets a reference to the given int32 and assigns it to the Destport field.
func (o *FunnelYaskawaMmcloudDestination) SetDestport(v int32) {
	o.Destport = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *FunnelYaskawaMmcloudDestination) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelYaskawaMmcloudDestination) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *FunnelYaskawaMmcloudDestination) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *FunnelYaskawaMmcloudDestination) SetProvider(v string) {
	o.Provider = &v
}

// GetResourceUrl returns the ResourceUrl field value if set, zero value otherwise.
func (o *FunnelYaskawaMmcloudDestination) GetResourceUrl() string {
	if o == nil || o.ResourceUrl == nil {
		var ret string
		return ret
	}
	return *o.ResourceUrl
}

// GetResourceUrlOk returns a tuple with the ResourceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelYaskawaMmcloudDestination) GetResourceUrlOk() (*string, bool) {
	if o == nil || o.ResourceUrl == nil {
		return nil, false
	}
	return o.ResourceUrl, true
}

// HasResourceUrl returns a boolean if a field has been set.
func (o *FunnelYaskawaMmcloudDestination) HasResourceUrl() bool {
	if o != nil && o.ResourceUrl != nil {
		return true
	}

	return false
}

// SetResourceUrl gets a reference to the given string and assigns it to the ResourceUrl field.
func (o *FunnelYaskawaMmcloudDestination) SetResourceUrl(v string) {
	o.ResourceUrl = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *FunnelYaskawaMmcloudDestination) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelYaskawaMmcloudDestination) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *FunnelYaskawaMmcloudDestination) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *FunnelYaskawaMmcloudDestination) SetService(v string) {
	o.Service = &v
}

func (o FunnelYaskawaMmcloudDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Desthost != nil {
		toSerialize["desthost"] = o.Desthost
	}
	if o.Destport != nil {
		toSerialize["destport"] = o.Destport
	}
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

type NullableFunnelYaskawaMmcloudDestination struct {
	value *FunnelYaskawaMmcloudDestination
	isSet bool
}

func (v NullableFunnelYaskawaMmcloudDestination) Get() *FunnelYaskawaMmcloudDestination {
	return v.value
}

func (v *NullableFunnelYaskawaMmcloudDestination) Set(val *FunnelYaskawaMmcloudDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableFunnelYaskawaMmcloudDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableFunnelYaskawaMmcloudDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunnelYaskawaMmcloudDestination(val *FunnelYaskawaMmcloudDestination) *NullableFunnelYaskawaMmcloudDestination {
	return &NullableFunnelYaskawaMmcloudDestination{value: val, isSet: true}
}

func (v NullableFunnelYaskawaMmcloudDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunnelYaskawaMmcloudDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

