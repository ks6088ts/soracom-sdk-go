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

// LogEntry struct for LogEntry
type LogEntry struct {
	Body map[string]interface{} `json:"body,omitempty"`
	ResourceId *string `json:"resourceId,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Service *string `json:"service,omitempty"`
	Time *int64 `json:"time,omitempty"`
}

// NewLogEntry instantiates a new LogEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogEntry() *LogEntry {
	this := LogEntry{}
	return &this
}

// NewLogEntryWithDefaults instantiates a new LogEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogEntryWithDefaults() *LogEntry {
	this := LogEntry{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *LogEntry) GetBody() map[string]interface{} {
	if o == nil || o.Body == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEntry) GetBodyOk() (map[string]interface{}, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *LogEntry) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given map[string]interface{} and assigns it to the Body field.
func (o *LogEntry) SetBody(v map[string]interface{}) {
	o.Body = v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *LogEntry) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEntry) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *LogEntry) HasResourceId() bool {
	if o != nil && o.ResourceId != nil {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *LogEntry) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *LogEntry) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEntry) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *LogEntry) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *LogEntry) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *LogEntry) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEntry) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *LogEntry) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *LogEntry) SetService(v string) {
	o.Service = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *LogEntry) GetTime() int64 {
	if o == nil || o.Time == nil {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEntry) GetTimeOk() (*int64, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *LogEntry) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *LogEntry) SetTime(v int64) {
	o.Time = &v
}

func (o LogEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.ResourceId != nil {
		toSerialize["resourceId"] = o.ResourceId
	}
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableLogEntry struct {
	value *LogEntry
	isSet bool
}

func (v NullableLogEntry) Get() *LogEntry {
	return v.value
}

func (v *NullableLogEntry) Set(val *LogEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableLogEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableLogEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogEntry(val *LogEntry) *NullableLogEntry {
	return &NullableLogEntry{value: val, isSet: true}
}

func (v NullableLogEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

