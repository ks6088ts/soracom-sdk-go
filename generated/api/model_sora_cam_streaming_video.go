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

// SoraCamStreamingVideo struct for SoraCamStreamingVideo
type SoraCamStreamingVideo struct {
	// ストリーミング映像 (リアルタイム映像 / 録画映像) のプレイリストエントリーの配列  - リアルタイム映像の場合は、返却される動画ファイルは常にひとつです。また、`from` および `to` は省略されます。 - 録画映像の場合は、録画の状態によって動画ファイルが分割されていることがあります。 
	PlayList []SoraCamPlayListEntry `json:"playList,omitempty"`
}

// NewSoraCamStreamingVideo instantiates a new SoraCamStreamingVideo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoraCamStreamingVideo() *SoraCamStreamingVideo {
	this := SoraCamStreamingVideo{}
	return &this
}

// NewSoraCamStreamingVideoWithDefaults instantiates a new SoraCamStreamingVideo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoraCamStreamingVideoWithDefaults() *SoraCamStreamingVideo {
	this := SoraCamStreamingVideo{}
	return &this
}

// GetPlayList returns the PlayList field value if set, zero value otherwise.
func (o *SoraCamStreamingVideo) GetPlayList() []SoraCamPlayListEntry {
	if o == nil || o.PlayList == nil {
		var ret []SoraCamPlayListEntry
		return ret
	}
	return o.PlayList
}

// GetPlayListOk returns a tuple with the PlayList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoraCamStreamingVideo) GetPlayListOk() ([]SoraCamPlayListEntry, bool) {
	if o == nil || o.PlayList == nil {
		return nil, false
	}
	return o.PlayList, true
}

// HasPlayList returns a boolean if a field has been set.
func (o *SoraCamStreamingVideo) HasPlayList() bool {
	if o != nil && o.PlayList != nil {
		return true
	}

	return false
}

// SetPlayList gets a reference to the given []SoraCamPlayListEntry and assigns it to the PlayList field.
func (o *SoraCamStreamingVideo) SetPlayList(v []SoraCamPlayListEntry) {
	o.PlayList = v
}

func (o SoraCamStreamingVideo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PlayList != nil {
		toSerialize["playList"] = o.PlayList
	}
	return json.Marshal(toSerialize)
}

type NullableSoraCamStreamingVideo struct {
	value *SoraCamStreamingVideo
	isSet bool
}

func (v NullableSoraCamStreamingVideo) Get() *SoraCamStreamingVideo {
	return v.value
}

func (v *NullableSoraCamStreamingVideo) Set(val *SoraCamStreamingVideo) {
	v.value = val
	v.isSet = true
}

func (v NullableSoraCamStreamingVideo) IsSet() bool {
	return v.isSet
}

func (v *NullableSoraCamStreamingVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoraCamStreamingVideo(val *SoraCamStreamingVideo) *NullableSoraCamStreamingVideo {
	return &NullableSoraCamStreamingVideo{value: val, isSet: true}
}

func (v NullableSoraCamStreamingVideo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoraCamStreamingVideo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

