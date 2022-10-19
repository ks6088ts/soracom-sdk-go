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

// DataSourceResourceMetadata struct for DataSourceResourceMetadata
type DataSourceResourceMetadata struct {
	AttributeNames []string `json:"attributeNames,omitempty"`
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty"`
	ResourceId *string `json:"resourceId,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
}

// NewDataSourceResourceMetadata instantiates a new DataSourceResourceMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSourceResourceMetadata() *DataSourceResourceMetadata {
	this := DataSourceResourceMetadata{}
	return &this
}

// NewDataSourceResourceMetadataWithDefaults instantiates a new DataSourceResourceMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSourceResourceMetadataWithDefaults() *DataSourceResourceMetadata {
	this := DataSourceResourceMetadata{}
	return &this
}

// GetAttributeNames returns the AttributeNames field value if set, zero value otherwise.
func (o *DataSourceResourceMetadata) GetAttributeNames() []string {
	if o == nil || o.AttributeNames == nil {
		var ret []string
		return ret
	}
	return o.AttributeNames
}

// GetAttributeNamesOk returns a tuple with the AttributeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceResourceMetadata) GetAttributeNamesOk() ([]string, bool) {
	if o == nil || o.AttributeNames == nil {
		return nil, false
	}
	return o.AttributeNames, true
}

// HasAttributeNames returns a boolean if a field has been set.
func (o *DataSourceResourceMetadata) HasAttributeNames() bool {
	if o != nil && o.AttributeNames != nil {
		return true
	}

	return false
}

// SetAttributeNames gets a reference to the given []string and assigns it to the AttributeNames field.
func (o *DataSourceResourceMetadata) SetAttributeNames(v []string) {
	o.AttributeNames = v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *DataSourceResourceMetadata) GetLastModifiedTime() int64 {
	if o == nil || o.LastModifiedTime == nil {
		var ret int64
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceResourceMetadata) GetLastModifiedTimeOk() (*int64, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *DataSourceResourceMetadata) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given int64 and assigns it to the LastModifiedTime field.
func (o *DataSourceResourceMetadata) SetLastModifiedTime(v int64) {
	o.LastModifiedTime = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *DataSourceResourceMetadata) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceResourceMetadata) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *DataSourceResourceMetadata) HasResourceId() bool {
	if o != nil && o.ResourceId != nil {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *DataSourceResourceMetadata) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *DataSourceResourceMetadata) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceResourceMetadata) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *DataSourceResourceMetadata) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *DataSourceResourceMetadata) SetResourceType(v string) {
	o.ResourceType = &v
}

func (o DataSourceResourceMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AttributeNames != nil {
		toSerialize["attributeNames"] = o.AttributeNames
	}
	if o.LastModifiedTime != nil {
		toSerialize["lastModifiedTime"] = o.LastModifiedTime
	}
	if o.ResourceId != nil {
		toSerialize["resourceId"] = o.ResourceId
	}
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}
	return json.Marshal(toSerialize)
}

type NullableDataSourceResourceMetadata struct {
	value *DataSourceResourceMetadata
	isSet bool
}

func (v NullableDataSourceResourceMetadata) Get() *DataSourceResourceMetadata {
	return v.value
}

func (v *NullableDataSourceResourceMetadata) Set(val *DataSourceResourceMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceResourceMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceResourceMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceResourceMetadata(val *DataSourceResourceMetadata) *NullableDataSourceResourceMetadata {
	return &NullableDataSourceResourceMetadata{value: val, isSet: true}
}

func (v NullableDataSourceResourceMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceResourceMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


