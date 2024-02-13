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

// checks if the PlayerHeroesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlayerHeroesResponse{}

// PlayerHeroesResponse hero
type PlayerHeroesResponse struct {
	// The ID value of the hero played
	HeroId *int32 `json:"hero_id,omitempty"`
	// last_played
	LastPlayed *int32 `json:"last_played,omitempty"`
	// games
	Games *int32 `json:"games,omitempty"`
	// win
	Win *int32 `json:"win,omitempty"`
	// with_games
	WithGames *int32 `json:"with_games,omitempty"`
	// with_win
	WithWin *int32 `json:"with_win,omitempty"`
	// against_games
	AgainstGames *int32 `json:"against_games,omitempty"`
	// against_win
	AgainstWin *int32 `json:"against_win,omitempty"`
}

// NewPlayerHeroesResponse instantiates a new PlayerHeroesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlayerHeroesResponse() *PlayerHeroesResponse {
	this := PlayerHeroesResponse{}
	return &this
}

// NewPlayerHeroesResponseWithDefaults instantiates a new PlayerHeroesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayerHeroesResponseWithDefaults() *PlayerHeroesResponse {
	this := PlayerHeroesResponse{}
	return &this
}

// GetHeroId returns the HeroId field value if set, zero value otherwise.
func (o *PlayerHeroesResponse) GetHeroId() int32 {
	if o == nil || IsNil(o.HeroId) {
		var ret int32
		return ret
	}
	return *o.HeroId
}

// GetHeroIdOk returns a tuple with the HeroId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerHeroesResponse) GetHeroIdOk() (*int32, bool) {
	if o == nil || IsNil(o.HeroId) {
		return nil, false
	}
	return o.HeroId, true
}

// HasHeroId returns a boolean if a field has been set.
func (o *PlayerHeroesResponse) HasHeroId() bool {
	if o != nil && !IsNil(o.HeroId) {
		return true
	}

	return false
}

// SetHeroId gets a reference to the given int32 and assigns it to the HeroId field.
func (o *PlayerHeroesResponse) SetHeroId(v int32) {
	o.HeroId = &v
}

// GetLastPlayed returns the LastPlayed field value if set, zero value otherwise.
func (o *PlayerHeroesResponse) GetLastPlayed() int32 {
	if o == nil || IsNil(o.LastPlayed) {
		var ret int32
		return ret
	}
	return *o.LastPlayed
}

// GetLastPlayedOk returns a tuple with the LastPlayed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerHeroesResponse) GetLastPlayedOk() (*int32, bool) {
	if o == nil || IsNil(o.LastPlayed) {
		return nil, false
	}
	return o.LastPlayed, true
}

// HasLastPlayed returns a boolean if a field has been set.
func (o *PlayerHeroesResponse) HasLastPlayed() bool {
	if o != nil && !IsNil(o.LastPlayed) {
		return true
	}

	return false
}

// SetLastPlayed gets a reference to the given int32 and assigns it to the LastPlayed field.
func (o *PlayerHeroesResponse) SetLastPlayed(v int32) {
	o.LastPlayed = &v
}

// GetGames returns the Games field value if set, zero value otherwise.
func (o *PlayerHeroesResponse) GetGames() int32 {
	if o == nil || IsNil(o.Games) {
		var ret int32
		return ret
	}
	return *o.Games
}

// GetGamesOk returns a tuple with the Games field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerHeroesResponse) GetGamesOk() (*int32, bool) {
	if o == nil || IsNil(o.Games) {
		return nil, false
	}
	return o.Games, true
}

// HasGames returns a boolean if a field has been set.
func (o *PlayerHeroesResponse) HasGames() bool {
	if o != nil && !IsNil(o.Games) {
		return true
	}

	return false
}

// SetGames gets a reference to the given int32 and assigns it to the Games field.
func (o *PlayerHeroesResponse) SetGames(v int32) {
	o.Games = &v
}

// GetWin returns the Win field value if set, zero value otherwise.
func (o *PlayerHeroesResponse) GetWin() int32 {
	if o == nil || IsNil(o.Win) {
		var ret int32
		return ret
	}
	return *o.Win
}

// GetWinOk returns a tuple with the Win field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerHeroesResponse) GetWinOk() (*int32, bool) {
	if o == nil || IsNil(o.Win) {
		return nil, false
	}
	return o.Win, true
}

// HasWin returns a boolean if a field has been set.
func (o *PlayerHeroesResponse) HasWin() bool {
	if o != nil && !IsNil(o.Win) {
		return true
	}

	return false
}

// SetWin gets a reference to the given int32 and assigns it to the Win field.
func (o *PlayerHeroesResponse) SetWin(v int32) {
	o.Win = &v
}

// GetWithGames returns the WithGames field value if set, zero value otherwise.
func (o *PlayerHeroesResponse) GetWithGames() int32 {
	if o == nil || IsNil(o.WithGames) {
		var ret int32
		return ret
	}
	return *o.WithGames
}

// GetWithGamesOk returns a tuple with the WithGames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerHeroesResponse) GetWithGamesOk() (*int32, bool) {
	if o == nil || IsNil(o.WithGames) {
		return nil, false
	}
	return o.WithGames, true
}

// HasWithGames returns a boolean if a field has been set.
func (o *PlayerHeroesResponse) HasWithGames() bool {
	if o != nil && !IsNil(o.WithGames) {
		return true
	}

	return false
}

// SetWithGames gets a reference to the given int32 and assigns it to the WithGames field.
func (o *PlayerHeroesResponse) SetWithGames(v int32) {
	o.WithGames = &v
}

// GetWithWin returns the WithWin field value if set, zero value otherwise.
func (o *PlayerHeroesResponse) GetWithWin() int32 {
	if o == nil || IsNil(o.WithWin) {
		var ret int32
		return ret
	}
	return *o.WithWin
}

// GetWithWinOk returns a tuple with the WithWin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerHeroesResponse) GetWithWinOk() (*int32, bool) {
	if o == nil || IsNil(o.WithWin) {
		return nil, false
	}
	return o.WithWin, true
}

// HasWithWin returns a boolean if a field has been set.
func (o *PlayerHeroesResponse) HasWithWin() bool {
	if o != nil && !IsNil(o.WithWin) {
		return true
	}

	return false
}

// SetWithWin gets a reference to the given int32 and assigns it to the WithWin field.
func (o *PlayerHeroesResponse) SetWithWin(v int32) {
	o.WithWin = &v
}

// GetAgainstGames returns the AgainstGames field value if set, zero value otherwise.
func (o *PlayerHeroesResponse) GetAgainstGames() int32 {
	if o == nil || IsNil(o.AgainstGames) {
		var ret int32
		return ret
	}
	return *o.AgainstGames
}

// GetAgainstGamesOk returns a tuple with the AgainstGames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerHeroesResponse) GetAgainstGamesOk() (*int32, bool) {
	if o == nil || IsNil(o.AgainstGames) {
		return nil, false
	}
	return o.AgainstGames, true
}

// HasAgainstGames returns a boolean if a field has been set.
func (o *PlayerHeroesResponse) HasAgainstGames() bool {
	if o != nil && !IsNil(o.AgainstGames) {
		return true
	}

	return false
}

// SetAgainstGames gets a reference to the given int32 and assigns it to the AgainstGames field.
func (o *PlayerHeroesResponse) SetAgainstGames(v int32) {
	o.AgainstGames = &v
}

// GetAgainstWin returns the AgainstWin field value if set, zero value otherwise.
func (o *PlayerHeroesResponse) GetAgainstWin() int32 {
	if o == nil || IsNil(o.AgainstWin) {
		var ret int32
		return ret
	}
	return *o.AgainstWin
}

// GetAgainstWinOk returns a tuple with the AgainstWin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerHeroesResponse) GetAgainstWinOk() (*int32, bool) {
	if o == nil || IsNil(o.AgainstWin) {
		return nil, false
	}
	return o.AgainstWin, true
}

// HasAgainstWin returns a boolean if a field has been set.
func (o *PlayerHeroesResponse) HasAgainstWin() bool {
	if o != nil && !IsNil(o.AgainstWin) {
		return true
	}

	return false
}

// SetAgainstWin gets a reference to the given int32 and assigns it to the AgainstWin field.
func (o *PlayerHeroesResponse) SetAgainstWin(v int32) {
	o.AgainstWin = &v
}

func (o PlayerHeroesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlayerHeroesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HeroId) {
		toSerialize["hero_id"] = o.HeroId
	}
	if !IsNil(o.LastPlayed) {
		toSerialize["last_played"] = o.LastPlayed
	}
	if !IsNil(o.Games) {
		toSerialize["games"] = o.Games
	}
	if !IsNil(o.Win) {
		toSerialize["win"] = o.Win
	}
	if !IsNil(o.WithGames) {
		toSerialize["with_games"] = o.WithGames
	}
	if !IsNil(o.WithWin) {
		toSerialize["with_win"] = o.WithWin
	}
	if !IsNil(o.AgainstGames) {
		toSerialize["against_games"] = o.AgainstGames
	}
	if !IsNil(o.AgainstWin) {
		toSerialize["against_win"] = o.AgainstWin
	}
	return toSerialize, nil
}

type NullablePlayerHeroesResponse struct {
	value *PlayerHeroesResponse
	isSet bool
}

func (v NullablePlayerHeroesResponse) Get() *PlayerHeroesResponse {
	return v.value
}

func (v *NullablePlayerHeroesResponse) Set(val *PlayerHeroesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayerHeroesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayerHeroesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayerHeroesResponse(val *PlayerHeroesResponse) *NullablePlayerHeroesResponse {
	return &NullablePlayerHeroesResponse{value: val, isSet: true}
}

func (v NullablePlayerHeroesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayerHeroesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

