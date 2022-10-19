/*
SORACOM API

SORACOM API v1

API version: VERSION_PLACEHOLDER
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"encoding/json"
	"time"
)

// LoraGateway struct for LoraGateway
type LoraGateway struct {
	Address *string `json:"address,omitempty"`
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	GatewayId *string `json:"gatewayId,omitempty"`
	LastModifiedTime *time.Time `json:"lastModifiedTime,omitempty"`
	NetworkSetId *string `json:"networkSetId,omitempty"`
	Online *bool `json:"online,omitempty"`
	OperatorId *string `json:"operatorId,omitempty"`
	Owned *bool `json:"owned,omitempty"`
	Status *string `json:"status,omitempty"`
	Tags *map[string]string `json:"tags,omitempty"`
	TerminationEnabled *bool `json:"terminationEnabled,omitempty"`
}

// NewLoraGateway instantiates a new LoraGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoraGateway() *LoraGateway {
	this := LoraGateway{}
	var online bool = false
	this.Online = &online
	var owned bool = false
	this.Owned = &owned
	var terminationEnabled bool = false
	this.TerminationEnabled = &terminationEnabled
	return &this
}

// NewLoraGatewayWithDefaults instantiates a new LoraGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoraGatewayWithDefaults() *LoraGateway {
	this := LoraGateway{}
	var online bool = false
	this.Online = &online
	var owned bool = false
	this.Owned = &owned
	var terminationEnabled bool = false
	this.TerminationEnabled = &terminationEnabled
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *LoraGateway) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraGateway) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *LoraGateway) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *LoraGateway) SetAddress(v string) {
	o.Address = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *LoraGateway) GetCreatedTime() time.Time {
	if o == nil || o.CreatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraGateway) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedTime == nil {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *LoraGateway) HasCreatedTime() bool {
	if o != nil && o.CreatedTime != nil {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *LoraGateway) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *LoraGateway) GetGatewayId() string {
	if o == nil || o.GatewayId == nil {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraGateway) GetGatewayIdOk() (*string, bool) {
	if o == nil || o.GatewayId == nil {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *LoraGateway) HasGatewayId() bool {
	if o != nil && o.GatewayId != nil {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *LoraGateway) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *LoraGateway) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraGateway) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *LoraGateway) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *LoraGateway) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetNetworkSetId returns the NetworkSetId field value if set, zero value otherwise.
func (o *LoraGateway) GetNetworkSetId() string {
	if o == nil || o.NetworkSetId == nil {
		var ret string
		return ret
	}
	return *o.NetworkSetId
}

// GetNetworkSetIdOk returns a tuple with the NetworkSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraGateway) GetNetworkSetIdOk() (*string, bool) {
	if o == nil || o.NetworkSetId == nil {
		return nil, false
	}
	return o.NetworkSetId, true
}

// HasNetworkSetId returns a boolean if a field has been set.
func (o *LoraGateway) HasNetworkSetId() bool {
	if o != nil && o.NetworkSetId != nil {
		return true
	}

	return false
}

// SetNetworkSetId gets a reference to the given string and assigns it to the NetworkSetId field.
func (o *LoraGateway) SetNetworkSetId(v string) {
	o.NetworkSetId = &v
}

// GetOnline returns the Online field value if set, zero value otherwise.
func (o *LoraGateway) GetOnline() bool {
	if o == nil || o.Online == nil {
		var ret bool
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraGateway) GetOnlineOk() (*bool, bool) {
	if o == nil || o.Online == nil {
		return nil, false
	}
	return o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *LoraGateway) HasOnline() bool {
	if o != nil && o.Online != nil {
		return true
	}

	return false
}

// SetOnline gets a reference to the given bool and assigns it to the Online field.
func (o *LoraGateway) SetOnline(v bool) {
	o.Online = &v
}

// GetOperatorId returns the OperatorId field value if set, zero value otherwise.
func (o *LoraGateway) GetOperatorId() string {
	if o == nil || o.OperatorId == nil {
		var ret string
		return ret
	}
	return *o.OperatorId
}

// GetOperatorIdOk returns a tuple with the OperatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraGateway) GetOperatorIdOk() (*string, bool) {
	if o == nil || o.OperatorId == nil {
		return nil, false
	}
	return o.OperatorId, true
}

// HasOperatorId returns a boolean if a field has been set.
func (o *LoraGateway) HasOperatorId() bool {
	if o != nil && o.OperatorId != nil {
		return true
	}

	return false
}

// SetOperatorId gets a reference to the given string and assigns it to the OperatorId field.
func (o *LoraGateway) SetOperatorId(v string) {
	o.OperatorId = &v
}

// GetOwned returns the Owned field value if set, zero value otherwise.
func (o *LoraGateway) GetOwned() bool {
	if o == nil || o.Owned == nil {
		var ret bool
		return ret
	}
	return *o.Owned
}

// GetOwnedOk returns a tuple with the Owned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraGateway) GetOwnedOk() (*bool, bool) {
	if o == nil || o.Owned == nil {
		return nil, false
	}
	return o.Owned, true
}

// HasOwned returns a boolean if a field has been set.
func (o *LoraGateway) HasOwned() bool {
	if o != nil && o.Owned != nil {
		return true
	}

	return false
}

// SetOwned gets a reference to the given bool and assigns it to the Owned field.
func (o *LoraGateway) SetOwned(v bool) {
	o.Owned = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LoraGateway) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraGateway) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LoraGateway) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LoraGateway) SetStatus(v string) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LoraGateway) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraGateway) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LoraGateway) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *LoraGateway) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetTerminationEnabled returns the TerminationEnabled field value if set, zero value otherwise.
func (o *LoraGateway) GetTerminationEnabled() bool {
	if o == nil || o.TerminationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TerminationEnabled
}

// GetTerminationEnabledOk returns a tuple with the TerminationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraGateway) GetTerminationEnabledOk() (*bool, bool) {
	if o == nil || o.TerminationEnabled == nil {
		return nil, false
	}
	return o.TerminationEnabled, true
}

// HasTerminationEnabled returns a boolean if a field has been set.
func (o *LoraGateway) HasTerminationEnabled() bool {
	if o != nil && o.TerminationEnabled != nil {
		return true
	}

	return false
}

// SetTerminationEnabled gets a reference to the given bool and assigns it to the TerminationEnabled field.
func (o *LoraGateway) SetTerminationEnabled(v bool) {
	o.TerminationEnabled = &v
}

func (o LoraGateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.CreatedTime != nil {
		toSerialize["createdTime"] = o.CreatedTime
	}
	if o.GatewayId != nil {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if o.LastModifiedTime != nil {
		toSerialize["lastModifiedTime"] = o.LastModifiedTime
	}
	if o.NetworkSetId != nil {
		toSerialize["networkSetId"] = o.NetworkSetId
	}
	if o.Online != nil {
		toSerialize["online"] = o.Online
	}
	if o.OperatorId != nil {
		toSerialize["operatorId"] = o.OperatorId
	}
	if o.Owned != nil {
		toSerialize["owned"] = o.Owned
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TerminationEnabled != nil {
		toSerialize["terminationEnabled"] = o.TerminationEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableLoraGateway struct {
	value *LoraGateway
	isSet bool
}

func (v NullableLoraGateway) Get() *LoraGateway {
	return v.value
}

func (v *NullableLoraGateway) Set(val *LoraGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableLoraGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableLoraGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoraGateway(val *LoraGateway) *NullableLoraGateway {
	return &NullableLoraGateway{value: val, isSet: true}
}

func (v NullableLoraGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoraGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


