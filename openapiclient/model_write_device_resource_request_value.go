/*
SORACOM API

SORACOM API v1

API version: VERSION_PLACEHOLDER
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"encoding/json"
	"fmt"
)

// WriteDeviceResourceRequestValue - struct for WriteDeviceResourceRequestValue
type WriteDeviceResourceRequestValue struct {
	Bool *bool
	Float32 *float32
	String *string
}

// boolAsWriteDeviceResourceRequestValue is a convenience function that returns bool wrapped in WriteDeviceResourceRequestValue
func BoolAsWriteDeviceResourceRequestValue(v *bool) WriteDeviceResourceRequestValue {
	return WriteDeviceResourceRequestValue{
		Bool: v,
	}
}

// float32AsWriteDeviceResourceRequestValue is a convenience function that returns float32 wrapped in WriteDeviceResourceRequestValue
func Float32AsWriteDeviceResourceRequestValue(v *float32) WriteDeviceResourceRequestValue {
	return WriteDeviceResourceRequestValue{
		Float32: v,
	}
}

// stringAsWriteDeviceResourceRequestValue is a convenience function that returns string wrapped in WriteDeviceResourceRequestValue
func StringAsWriteDeviceResourceRequestValue(v *string) WriteDeviceResourceRequestValue {
	return WriteDeviceResourceRequestValue{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *WriteDeviceResourceRequestValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			match++
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal data into Float32
	err = newStrictDecoder(data).Decode(&dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			match++
		}
	} else {
		dst.Float32 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Bool = nil
		dst.Float32 = nil
		dst.String = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(WriteDeviceResourceRequestValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(WriteDeviceResourceRequestValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src WriteDeviceResourceRequestValue) MarshalJSON() ([]byte, error) {
	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WriteDeviceResourceRequestValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.Float32 != nil {
		return obj.Float32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableWriteDeviceResourceRequestValue struct {
	value *WriteDeviceResourceRequestValue
	isSet bool
}

func (v NullableWriteDeviceResourceRequestValue) Get() *WriteDeviceResourceRequestValue {
	return v.value
}

func (v *NullableWriteDeviceResourceRequestValue) Set(val *WriteDeviceResourceRequestValue) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteDeviceResourceRequestValue) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteDeviceResourceRequestValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteDeviceResourceRequestValue(val *WriteDeviceResourceRequestValue) *NullableWriteDeviceResourceRequestValue {
	return &NullableWriteDeviceResourceRequestValue{value: val, isSet: true}
}

func (v NullableWriteDeviceResourceRequestValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteDeviceResourceRequestValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


