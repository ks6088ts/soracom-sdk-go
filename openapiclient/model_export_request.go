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

// ExportRequest struct for ExportRequest
type ExportRequest struct {
	From *int64 `json:"from,omitempty"`
	Period *string `json:"period,omitempty"`
	To *int64 `json:"to,omitempty"`
}

// NewExportRequest instantiates a new ExportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportRequest() *ExportRequest {
	this := ExportRequest{}
	return &this
}

// NewExportRequestWithDefaults instantiates a new ExportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportRequestWithDefaults() *ExportRequest {
	this := ExportRequest{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *ExportRequest) GetFrom() int64 {
	if o == nil || o.From == nil {
		var ret int64
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportRequest) GetFromOk() (*int64, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *ExportRequest) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given int64 and assigns it to the From field.
func (o *ExportRequest) SetFrom(v int64) {
	o.From = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *ExportRequest) GetPeriod() string {
	if o == nil || o.Period == nil {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportRequest) GetPeriodOk() (*string, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *ExportRequest) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *ExportRequest) SetPeriod(v string) {
	o.Period = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *ExportRequest) GetTo() int64 {
	if o == nil || o.To == nil {
		var ret int64
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportRequest) GetToOk() (*int64, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *ExportRequest) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given int64 and assigns it to the To field.
func (o *ExportRequest) SetTo(v int64) {
	o.To = &v
}

func (o ExportRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableExportRequest struct {
	value *ExportRequest
	isSet bool
}

func (v NullableExportRequest) Get() *ExportRequest {
	return v.value
}

func (v *NullableExportRequest) Set(val *ExportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportRequest(val *ExportRequest) *NullableExportRequest {
	return &NullableExportRequest{value: val, isSet: true}
}

func (v NullableExportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

