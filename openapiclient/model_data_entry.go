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

// DataEntry struct for DataEntry
type DataEntry struct {
	Content *string `json:"content,omitempty"`
	ContentType *string `json:"contentType,omitempty"`
	Time *int64 `json:"time,omitempty"`
}

// NewDataEntry instantiates a new DataEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataEntry() *DataEntry {
	this := DataEntry{}
	return &this
}

// NewDataEntryWithDefaults instantiates a new DataEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataEntryWithDefaults() *DataEntry {
	this := DataEntry{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *DataEntry) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataEntry) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *DataEntry) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *DataEntry) SetContent(v string) {
	o.Content = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *DataEntry) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataEntry) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *DataEntry) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *DataEntry) SetContentType(v string) {
	o.ContentType = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *DataEntry) GetTime() int64 {
	if o == nil || o.Time == nil {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataEntry) GetTimeOk() (*int64, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *DataEntry) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *DataEntry) SetTime(v int64) {
	o.Time = &v
}

func (o DataEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.ContentType != nil {
		toSerialize["contentType"] = o.ContentType
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableDataEntry struct {
	value *DataEntry
	isSet bool
}

func (v NullableDataEntry) Get() *DataEntry {
	return v.value
}

func (v *NullableDataEntry) Set(val *DataEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableDataEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableDataEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataEntry(val *DataEntry) *NullableDataEntry {
	return &NullableDataEntry{value: val, isSet: true}
}

func (v NullableDataEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


