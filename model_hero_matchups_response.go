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

// checks if the HeroMatchupsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeroMatchupsResponse{}

// HeroMatchupsResponse struct for HeroMatchupsResponse
type HeroMatchupsResponse struct {
	// The ID value of the hero played
	HeroId *int32 `json:"hero_id,omitempty"`
	// Number of games played
	GamesPlayed *int32 `json:"games_played,omitempty"`
	// Number of games won
	Wins *int32 `json:"wins,omitempty"`
}

// NewHeroMatchupsResponse instantiates a new HeroMatchupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeroMatchupsResponse() *HeroMatchupsResponse {
	this := HeroMatchupsResponse{}
	return &this
}

// NewHeroMatchupsResponseWithDefaults instantiates a new HeroMatchupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeroMatchupsResponseWithDefaults() *HeroMatchupsResponse {
	this := HeroMatchupsResponse{}
	return &this
}

// GetHeroId returns the HeroId field value if set, zero value otherwise.
func (o *HeroMatchupsResponse) GetHeroId() int32 {
	if o == nil || IsNil(o.HeroId) {
		var ret int32
		return ret
	}
	return *o.HeroId
}

// GetHeroIdOk returns a tuple with the HeroId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeroMatchupsResponse) GetHeroIdOk() (*int32, bool) {
	if o == nil || IsNil(o.HeroId) {
		return nil, false
	}
	return o.HeroId, true
}

// HasHeroId returns a boolean if a field has been set.
func (o *HeroMatchupsResponse) HasHeroId() bool {
	if o != nil && !IsNil(o.HeroId) {
		return true
	}

	return false
}

// SetHeroId gets a reference to the given int32 and assigns it to the HeroId field.
func (o *HeroMatchupsResponse) SetHeroId(v int32) {
	o.HeroId = &v
}

// GetGamesPlayed returns the GamesPlayed field value if set, zero value otherwise.
func (o *HeroMatchupsResponse) GetGamesPlayed() int32 {
	if o == nil || IsNil(o.GamesPlayed) {
		var ret int32
		return ret
	}
	return *o.GamesPlayed
}

// GetGamesPlayedOk returns a tuple with the GamesPlayed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeroMatchupsResponse) GetGamesPlayedOk() (*int32, bool) {
	if o == nil || IsNil(o.GamesPlayed) {
		return nil, false
	}
	return o.GamesPlayed, true
}

// HasGamesPlayed returns a boolean if a field has been set.
func (o *HeroMatchupsResponse) HasGamesPlayed() bool {
	if o != nil && !IsNil(o.GamesPlayed) {
		return true
	}

	return false
}

// SetGamesPlayed gets a reference to the given int32 and assigns it to the GamesPlayed field.
func (o *HeroMatchupsResponse) SetGamesPlayed(v int32) {
	o.GamesPlayed = &v
}

// GetWins returns the Wins field value if set, zero value otherwise.
func (o *HeroMatchupsResponse) GetWins() int32 {
	if o == nil || IsNil(o.Wins) {
		var ret int32
		return ret
	}
	return *o.Wins
}

// GetWinsOk returns a tuple with the Wins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeroMatchupsResponse) GetWinsOk() (*int32, bool) {
	if o == nil || IsNil(o.Wins) {
		return nil, false
	}
	return o.Wins, true
}

// HasWins returns a boolean if a field has been set.
func (o *HeroMatchupsResponse) HasWins() bool {
	if o != nil && !IsNil(o.Wins) {
		return true
	}

	return false
}

// SetWins gets a reference to the given int32 and assigns it to the Wins field.
func (o *HeroMatchupsResponse) SetWins(v int32) {
	o.Wins = &v
}

func (o HeroMatchupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HeroMatchupsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HeroId) {
		toSerialize["hero_id"] = o.HeroId
	}
	if !IsNil(o.GamesPlayed) {
		toSerialize["games_played"] = o.GamesPlayed
	}
	if !IsNil(o.Wins) {
		toSerialize["wins"] = o.Wins
	}
	return toSerialize, nil
}

type NullableHeroMatchupsResponse struct {
	value *HeroMatchupsResponse
	isSet bool
}

func (v NullableHeroMatchupsResponse) Get() *HeroMatchupsResponse {
	return v.value
}

func (v *NullableHeroMatchupsResponse) Set(val *HeroMatchupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHeroMatchupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHeroMatchupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeroMatchupsResponse(val *HeroMatchupsResponse) *NullableHeroMatchupsResponse {
	return &NullableHeroMatchupsResponse{value: val, isSet: true}
}

func (v NullableHeroMatchupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeroMatchupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

