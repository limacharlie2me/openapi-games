/*
OpenDota API

# Introduction The OpenDota API provides Dota 2 related data including advanced match data extracted from match replays.  You can find data that can be used to convert hero and ability IDs and other information provided by the API from the [dotaconstants](https://github.com/odota/dotaconstants) repository.  Without a key, you can make 2,000 free calls per day at a rate limit of 60 requests/minute. We also offer a Premium Tier with unlimited API calls and higher rate limits. Check out the [API page](https://www.opendota.com/api-keys) to learn more.     

API version: 25.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BenchmarksResponseResultGoldPerMinInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BenchmarksResponseResultGoldPerMinInner{}

// BenchmarksResponseResultGoldPerMinInner struct for BenchmarksResponseResultGoldPerMinInner
type BenchmarksResponseResultGoldPerMinInner struct {
	// percentile
	Percentile *float32 `json:"percentile,omitempty"`
	// value
	Value *float32 `json:"value,omitempty"`
}

// NewBenchmarksResponseResultGoldPerMinInner instantiates a new BenchmarksResponseResultGoldPerMinInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBenchmarksResponseResultGoldPerMinInner() *BenchmarksResponseResultGoldPerMinInner {
	this := BenchmarksResponseResultGoldPerMinInner{}
	return &this
}

// NewBenchmarksResponseResultGoldPerMinInnerWithDefaults instantiates a new BenchmarksResponseResultGoldPerMinInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBenchmarksResponseResultGoldPerMinInnerWithDefaults() *BenchmarksResponseResultGoldPerMinInner {
	this := BenchmarksResponseResultGoldPerMinInner{}
	return &this
}

// GetPercentile returns the Percentile field value if set, zero value otherwise.
func (o *BenchmarksResponseResultGoldPerMinInner) GetPercentile() float32 {
	if o == nil || IsNil(o.Percentile) {
		var ret float32
		return ret
	}
	return *o.Percentile
}

// GetPercentileOk returns a tuple with the Percentile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BenchmarksResponseResultGoldPerMinInner) GetPercentileOk() (*float32, bool) {
	if o == nil || IsNil(o.Percentile) {
		return nil, false
	}
	return o.Percentile, true
}

// HasPercentile returns a boolean if a field has been set.
func (o *BenchmarksResponseResultGoldPerMinInner) HasPercentile() bool {
	if o != nil && !IsNil(o.Percentile) {
		return true
	}

	return false
}

// SetPercentile gets a reference to the given float32 and assigns it to the Percentile field.
func (o *BenchmarksResponseResultGoldPerMinInner) SetPercentile(v float32) {
	o.Percentile = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BenchmarksResponseResultGoldPerMinInner) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BenchmarksResponseResultGoldPerMinInner) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BenchmarksResponseResultGoldPerMinInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *BenchmarksResponseResultGoldPerMinInner) SetValue(v float32) {
	o.Value = &v
}

func (o BenchmarksResponseResultGoldPerMinInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BenchmarksResponseResultGoldPerMinInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Percentile) {
		toSerialize["percentile"] = o.Percentile
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableBenchmarksResponseResultGoldPerMinInner struct {
	value *BenchmarksResponseResultGoldPerMinInner
	isSet bool
}

func (v NullableBenchmarksResponseResultGoldPerMinInner) Get() *BenchmarksResponseResultGoldPerMinInner {
	return v.value
}

func (v *NullableBenchmarksResponseResultGoldPerMinInner) Set(val *BenchmarksResponseResultGoldPerMinInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBenchmarksResponseResultGoldPerMinInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBenchmarksResponseResultGoldPerMinInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBenchmarksResponseResultGoldPerMinInner(val *BenchmarksResponseResultGoldPerMinInner) *NullableBenchmarksResponseResultGoldPerMinInner {
	return &NullableBenchmarksResponseResultGoldPerMinInner{value: val, isSet: true}
}

func (v NullableBenchmarksResponseResultGoldPerMinInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBenchmarksResponseResultGoldPerMinInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

