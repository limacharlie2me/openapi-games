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

// checks if the RankingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RankingsResponse{}

// RankingsResponse struct for RankingsResponse
type RankingsResponse struct {
	// The ID value of the hero played
	HeroId *int32 `json:"hero_id,omitempty"`
	// rankings
	Rankings []RankingsResponseRankingsInner `json:"rankings,omitempty"`
}

// NewRankingsResponse instantiates a new RankingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRankingsResponse() *RankingsResponse {
	this := RankingsResponse{}
	return &this
}

// NewRankingsResponseWithDefaults instantiates a new RankingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRankingsResponseWithDefaults() *RankingsResponse {
	this := RankingsResponse{}
	return &this
}

// GetHeroId returns the HeroId field value if set, zero value otherwise.
func (o *RankingsResponse) GetHeroId() int32 {
	if o == nil || IsNil(o.HeroId) {
		var ret int32
		return ret
	}
	return *o.HeroId
}

// GetHeroIdOk returns a tuple with the HeroId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RankingsResponse) GetHeroIdOk() (*int32, bool) {
	if o == nil || IsNil(o.HeroId) {
		return nil, false
	}
	return o.HeroId, true
}

// HasHeroId returns a boolean if a field has been set.
func (o *RankingsResponse) HasHeroId() bool {
	if o != nil && !IsNil(o.HeroId) {
		return true
	}

	return false
}

// SetHeroId gets a reference to the given int32 and assigns it to the HeroId field.
func (o *RankingsResponse) SetHeroId(v int32) {
	o.HeroId = &v
}

// GetRankings returns the Rankings field value if set, zero value otherwise.
func (o *RankingsResponse) GetRankings() []RankingsResponseRankingsInner {
	if o == nil || IsNil(o.Rankings) {
		var ret []RankingsResponseRankingsInner
		return ret
	}
	return o.Rankings
}

// GetRankingsOk returns a tuple with the Rankings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RankingsResponse) GetRankingsOk() ([]RankingsResponseRankingsInner, bool) {
	if o == nil || IsNil(o.Rankings) {
		return nil, false
	}
	return o.Rankings, true
}

// HasRankings returns a boolean if a field has been set.
func (o *RankingsResponse) HasRankings() bool {
	if o != nil && !IsNil(o.Rankings) {
		return true
	}

	return false
}

// SetRankings gets a reference to the given []RankingsResponseRankingsInner and assigns it to the Rankings field.
func (o *RankingsResponse) SetRankings(v []RankingsResponseRankingsInner) {
	o.Rankings = v
}

func (o RankingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RankingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HeroId) {
		toSerialize["hero_id"] = o.HeroId
	}
	if !IsNil(o.Rankings) {
		toSerialize["rankings"] = o.Rankings
	}
	return toSerialize, nil
}

type NullableRankingsResponse struct {
	value *RankingsResponse
	isSet bool
}

func (v NullableRankingsResponse) Get() *RankingsResponse {
	return v.value
}

func (v *NullableRankingsResponse) Set(val *RankingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRankingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRankingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRankingsResponse(val *RankingsResponse) *NullableRankingsResponse {
	return &NullableRankingsResponse{value: val, isSet: true}
}

func (v NullableRankingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRankingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


