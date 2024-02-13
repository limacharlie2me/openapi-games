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

// checks if the HeroItemPopularityResponseMidGameItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeroItemPopularityResponseMidGameItems{}

// HeroItemPopularityResponseMidGameItems Items bought between 10 and 25 min of the game, with cost at least 2000
type HeroItemPopularityResponseMidGameItems struct {
	// Number of item bought
	Item *int32 `json:"item,omitempty"`
}

// NewHeroItemPopularityResponseMidGameItems instantiates a new HeroItemPopularityResponseMidGameItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeroItemPopularityResponseMidGameItems() *HeroItemPopularityResponseMidGameItems {
	this := HeroItemPopularityResponseMidGameItems{}
	return &this
}

// NewHeroItemPopularityResponseMidGameItemsWithDefaults instantiates a new HeroItemPopularityResponseMidGameItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeroItemPopularityResponseMidGameItemsWithDefaults() *HeroItemPopularityResponseMidGameItems {
	this := HeroItemPopularityResponseMidGameItems{}
	return &this
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *HeroItemPopularityResponseMidGameItems) GetItem() int32 {
	if o == nil || IsNil(o.Item) {
		var ret int32
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeroItemPopularityResponseMidGameItems) GetItemOk() (*int32, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *HeroItemPopularityResponseMidGameItems) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given int32 and assigns it to the Item field.
func (o *HeroItemPopularityResponseMidGameItems) SetItem(v int32) {
	o.Item = &v
}

func (o HeroItemPopularityResponseMidGameItems) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HeroItemPopularityResponseMidGameItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	return toSerialize, nil
}

type NullableHeroItemPopularityResponseMidGameItems struct {
	value *HeroItemPopularityResponseMidGameItems
	isSet bool
}

func (v NullableHeroItemPopularityResponseMidGameItems) Get() *HeroItemPopularityResponseMidGameItems {
	return v.value
}

func (v *NullableHeroItemPopularityResponseMidGameItems) Set(val *HeroItemPopularityResponseMidGameItems) {
	v.value = val
	v.isSet = true
}

func (v NullableHeroItemPopularityResponseMidGameItems) IsSet() bool {
	return v.isSet
}

func (v *NullableHeroItemPopularityResponseMidGameItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeroItemPopularityResponseMidGameItems(val *HeroItemPopularityResponseMidGameItems) *NullableHeroItemPopularityResponseMidGameItems {
	return &NullableHeroItemPopularityResponseMidGameItems{value: val, isSet: true}
}

func (v NullableHeroItemPopularityResponseMidGameItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeroItemPopularityResponseMidGameItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


