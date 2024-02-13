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

// checks if the BenchmarksResponseResultTowerDamageInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BenchmarksResponseResultTowerDamageInner{}

// BenchmarksResponseResultTowerDamageInner struct for BenchmarksResponseResultTowerDamageInner
type BenchmarksResponseResultTowerDamageInner struct {
	// percentile
	Percentile *float32 `json:"percentile,omitempty"`
	// value
	Value *int32 `json:"value,omitempty"`
}

// NewBenchmarksResponseResultTowerDamageInner instantiates a new BenchmarksResponseResultTowerDamageInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBenchmarksResponseResultTowerDamageInner() *BenchmarksResponseResultTowerDamageInner {
	this := BenchmarksResponseResultTowerDamageInner{}
	return &this
}

// NewBenchmarksResponseResultTowerDamageInnerWithDefaults instantiates a new BenchmarksResponseResultTowerDamageInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBenchmarksResponseResultTowerDamageInnerWithDefaults() *BenchmarksResponseResultTowerDamageInner {
	this := BenchmarksResponseResultTowerDamageInner{}
	return &this
}

// GetPercentile returns the Percentile field value if set, zero value otherwise.
func (o *BenchmarksResponseResultTowerDamageInner) GetPercentile() float32 {
	if o == nil || IsNil(o.Percentile) {
		var ret float32
		return ret
	}
	return *o.Percentile
}

// GetPercentileOk returns a tuple with the Percentile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BenchmarksResponseResultTowerDamageInner) GetPercentileOk() (*float32, bool) {
	if o == nil || IsNil(o.Percentile) {
		return nil, false
	}
	return o.Percentile, true
}

// HasPercentile returns a boolean if a field has been set.
func (o *BenchmarksResponseResultTowerDamageInner) HasPercentile() bool {
	if o != nil && !IsNil(o.Percentile) {
		return true
	}

	return false
}

// SetPercentile gets a reference to the given float32 and assigns it to the Percentile field.
func (o *BenchmarksResponseResultTowerDamageInner) SetPercentile(v float32) {
	o.Percentile = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BenchmarksResponseResultTowerDamageInner) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BenchmarksResponseResultTowerDamageInner) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BenchmarksResponseResultTowerDamageInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *BenchmarksResponseResultTowerDamageInner) SetValue(v int32) {
	o.Value = &v
}

func (o BenchmarksResponseResultTowerDamageInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BenchmarksResponseResultTowerDamageInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Percentile) {
		toSerialize["percentile"] = o.Percentile
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableBenchmarksResponseResultTowerDamageInner struct {
	value *BenchmarksResponseResultTowerDamageInner
	isSet bool
}

func (v NullableBenchmarksResponseResultTowerDamageInner) Get() *BenchmarksResponseResultTowerDamageInner {
	return v.value
}

func (v *NullableBenchmarksResponseResultTowerDamageInner) Set(val *BenchmarksResponseResultTowerDamageInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBenchmarksResponseResultTowerDamageInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBenchmarksResponseResultTowerDamageInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBenchmarksResponseResultTowerDamageInner(val *BenchmarksResponseResultTowerDamageInner) *NullableBenchmarksResponseResultTowerDamageInner {
	return &NullableBenchmarksResponseResultTowerDamageInner{value: val, isSet: true}
}

func (v NullableBenchmarksResponseResultTowerDamageInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBenchmarksResponseResultTowerDamageInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


