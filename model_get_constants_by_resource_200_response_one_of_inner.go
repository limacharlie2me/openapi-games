/*
OpenDota API

# Introduction The OpenDota API provides Dota 2 related data including advanced match data extracted from match replays.  You can find data that can be used to convert hero and ability IDs and other information provided by the API from the [dotaconstants](https://github.com/odota/dotaconstants) repository.  Without a key, you can make 2,000 free calls per day at a rate limit of 60 requests/minute. We also offer a Premium Tier with unlimited API calls and higher rate limits. Check out the [API page](https://www.opendota.com/api-keys) to learn more.     

API version: 25.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// GetConstantsByResource200ResponseOneOfInner - struct for GetConstantsByResource200ResponseOneOfInner
type GetConstantsByResource200ResponseOneOfInner struct {
	Int32 *int32
	MapmapOfStringAny *map[string]interface{}
}

// int32AsGetConstantsByResource200ResponseOneOfInner is a convenience function that returns int32 wrapped in GetConstantsByResource200ResponseOneOfInner
func Int32AsGetConstantsByResource200ResponseOneOfInner(v *int32) GetConstantsByResource200ResponseOneOfInner {
	return GetConstantsByResource200ResponseOneOfInner{
		Int32: v,
	}
}

// map[string]interface{}AsGetConstantsByResource200ResponseOneOfInner is a convenience function that returns map[string]interface{} wrapped in GetConstantsByResource200ResponseOneOfInner
func MapmapOfStringAnyAsGetConstantsByResource200ResponseOneOfInner(v *map[string]interface{}) GetConstantsByResource200ResponseOneOfInner {
	return GetConstantsByResource200ResponseOneOfInner{
		MapmapOfStringAny: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetConstantsByResource200ResponseOneOfInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			match++
		}
	} else {
		dst.Int32 = nil
	}

	// try to unmarshal data into MapmapOfStringAny
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringAny)
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			match++
		}
	} else {
		dst.MapmapOfStringAny = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Int32 = nil
		dst.MapmapOfStringAny = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetConstantsByResource200ResponseOneOfInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetConstantsByResource200ResponseOneOfInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetConstantsByResource200ResponseOneOfInner) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetConstantsByResource200ResponseOneOfInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Int32 != nil {
		return obj.Int32
	}

	if obj.MapmapOfStringAny != nil {
		return obj.MapmapOfStringAny
	}

	// all schemas are nil
	return nil
}

type NullableGetConstantsByResource200ResponseOneOfInner struct {
	value *GetConstantsByResource200ResponseOneOfInner
	isSet bool
}

func (v NullableGetConstantsByResource200ResponseOneOfInner) Get() *GetConstantsByResource200ResponseOneOfInner {
	return v.value
}

func (v *NullableGetConstantsByResource200ResponseOneOfInner) Set(val *GetConstantsByResource200ResponseOneOfInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConstantsByResource200ResponseOneOfInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConstantsByResource200ResponseOneOfInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConstantsByResource200ResponseOneOfInner(val *GetConstantsByResource200ResponseOneOfInner) *NullableGetConstantsByResource200ResponseOneOfInner {
	return &NullableGetConstantsByResource200ResponseOneOfInner{value: val, isSet: true}
}

func (v NullableGetConstantsByResource200ResponseOneOfInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConstantsByResource200ResponseOneOfInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


