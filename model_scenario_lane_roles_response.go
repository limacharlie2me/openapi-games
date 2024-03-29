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

// checks if the ScenarioLaneRolesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScenarioLaneRolesResponse{}

// ScenarioLaneRolesResponse struct for ScenarioLaneRolesResponse
type ScenarioLaneRolesResponse struct {
	// The ID value of the hero played
	HeroId *int32 `json:"hero_id,omitempty"`
	// The hero's lane role
	LaneRole *int32 `json:"lane_role,omitempty"`
	// Maximum game length in seconds
	Time *int32 `json:"time,omitempty"`
	// The number of games where the hero played in this lane role
	Games *string `json:"games,omitempty"`
	// The number of games won where the hero played in this lane role
	Wins *string `json:"wins,omitempty"`
}

// NewScenarioLaneRolesResponse instantiates a new ScenarioLaneRolesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScenarioLaneRolesResponse() *ScenarioLaneRolesResponse {
	this := ScenarioLaneRolesResponse{}
	return &this
}

// NewScenarioLaneRolesResponseWithDefaults instantiates a new ScenarioLaneRolesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScenarioLaneRolesResponseWithDefaults() *ScenarioLaneRolesResponse {
	this := ScenarioLaneRolesResponse{}
	return &this
}

// GetHeroId returns the HeroId field value if set, zero value otherwise.
func (o *ScenarioLaneRolesResponse) GetHeroId() int32 {
	if o == nil || IsNil(o.HeroId) {
		var ret int32
		return ret
	}
	return *o.HeroId
}

// GetHeroIdOk returns a tuple with the HeroId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioLaneRolesResponse) GetHeroIdOk() (*int32, bool) {
	if o == nil || IsNil(o.HeroId) {
		return nil, false
	}
	return o.HeroId, true
}

// HasHeroId returns a boolean if a field has been set.
func (o *ScenarioLaneRolesResponse) HasHeroId() bool {
	if o != nil && !IsNil(o.HeroId) {
		return true
	}

	return false
}

// SetHeroId gets a reference to the given int32 and assigns it to the HeroId field.
func (o *ScenarioLaneRolesResponse) SetHeroId(v int32) {
	o.HeroId = &v
}

// GetLaneRole returns the LaneRole field value if set, zero value otherwise.
func (o *ScenarioLaneRolesResponse) GetLaneRole() int32 {
	if o == nil || IsNil(o.LaneRole) {
		var ret int32
		return ret
	}
	return *o.LaneRole
}

// GetLaneRoleOk returns a tuple with the LaneRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioLaneRolesResponse) GetLaneRoleOk() (*int32, bool) {
	if o == nil || IsNil(o.LaneRole) {
		return nil, false
	}
	return o.LaneRole, true
}

// HasLaneRole returns a boolean if a field has been set.
func (o *ScenarioLaneRolesResponse) HasLaneRole() bool {
	if o != nil && !IsNil(o.LaneRole) {
		return true
	}

	return false
}

// SetLaneRole gets a reference to the given int32 and assigns it to the LaneRole field.
func (o *ScenarioLaneRolesResponse) SetLaneRole(v int32) {
	o.LaneRole = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ScenarioLaneRolesResponse) GetTime() int32 {
	if o == nil || IsNil(o.Time) {
		var ret int32
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioLaneRolesResponse) GetTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ScenarioLaneRolesResponse) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int32 and assigns it to the Time field.
func (o *ScenarioLaneRolesResponse) SetTime(v int32) {
	o.Time = &v
}

// GetGames returns the Games field value if set, zero value otherwise.
func (o *ScenarioLaneRolesResponse) GetGames() string {
	if o == nil || IsNil(o.Games) {
		var ret string
		return ret
	}
	return *o.Games
}

// GetGamesOk returns a tuple with the Games field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioLaneRolesResponse) GetGamesOk() (*string, bool) {
	if o == nil || IsNil(o.Games) {
		return nil, false
	}
	return o.Games, true
}

// HasGames returns a boolean if a field has been set.
func (o *ScenarioLaneRolesResponse) HasGames() bool {
	if o != nil && !IsNil(o.Games) {
		return true
	}

	return false
}

// SetGames gets a reference to the given string and assigns it to the Games field.
func (o *ScenarioLaneRolesResponse) SetGames(v string) {
	o.Games = &v
}

// GetWins returns the Wins field value if set, zero value otherwise.
func (o *ScenarioLaneRolesResponse) GetWins() string {
	if o == nil || IsNil(o.Wins) {
		var ret string
		return ret
	}
	return *o.Wins
}

// GetWinsOk returns a tuple with the Wins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioLaneRolesResponse) GetWinsOk() (*string, bool) {
	if o == nil || IsNil(o.Wins) {
		return nil, false
	}
	return o.Wins, true
}

// HasWins returns a boolean if a field has been set.
func (o *ScenarioLaneRolesResponse) HasWins() bool {
	if o != nil && !IsNil(o.Wins) {
		return true
	}

	return false
}

// SetWins gets a reference to the given string and assigns it to the Wins field.
func (o *ScenarioLaneRolesResponse) SetWins(v string) {
	o.Wins = &v
}

func (o ScenarioLaneRolesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScenarioLaneRolesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HeroId) {
		toSerialize["hero_id"] = o.HeroId
	}
	if !IsNil(o.LaneRole) {
		toSerialize["lane_role"] = o.LaneRole
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.Games) {
		toSerialize["games"] = o.Games
	}
	if !IsNil(o.Wins) {
		toSerialize["wins"] = o.Wins
	}
	return toSerialize, nil
}

type NullableScenarioLaneRolesResponse struct {
	value *ScenarioLaneRolesResponse
	isSet bool
}

func (v NullableScenarioLaneRolesResponse) Get() *ScenarioLaneRolesResponse {
	return v.value
}

func (v *NullableScenarioLaneRolesResponse) Set(val *ScenarioLaneRolesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScenarioLaneRolesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScenarioLaneRolesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScenarioLaneRolesResponse(val *ScenarioLaneRolesResponse) *NullableScenarioLaneRolesResponse {
	return &NullableScenarioLaneRolesResponse{value: val, isSet: true}
}

func (v NullableScenarioLaneRolesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScenarioLaneRolesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


