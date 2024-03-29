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

// checks if the PlayerRecentMatchesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlayerRecentMatchesResponse{}

// PlayerRecentMatchesResponse match
type PlayerRecentMatchesResponse struct {
	// The ID number of the match assigned by Valve
	MatchId *int32 `json:"match_id,omitempty"`
	// Which slot the player is in. 0-127 are Radiant, 128-255 are Dire
	PlayerSlot NullableInt32 `json:"player_slot,omitempty"`
	// Boolean indicating whether Radiant won the match
	RadiantWin NullableBool `json:"radiant_win,omitempty"`
	// Duration of the game in seconds
	Duration *int32 `json:"duration,omitempty"`
	// Integer corresponding to game mode played. List of constants can be found here: https://github.com/odota/dotaconstants/blob/master/json/game_mode.json
	GameMode *int32 `json:"game_mode,omitempty"`
	// Integer corresponding to lobby type of match. List of constants can be found here: https://github.com/odota/dotaconstants/blob/master/json/lobby_type.json
	LobbyType *int32 `json:"lobby_type,omitempty"`
	// The ID value of the hero played
	HeroId *int32 `json:"hero_id,omitempty"`
	// The Unix timestamp at which the game started
	StartTime *int32 `json:"start_time,omitempty"`
	// version
	Version NullableInt32 `json:"version,omitempty"`
	// Total kills the player had at the end of the match
	Kills *int32 `json:"kills,omitempty"`
	// Total deaths the player had at the end of the match
	Deaths *int32 `json:"deaths,omitempty"`
	// Total assists the player had at the end of the match
	Assists *int32 `json:"assists,omitempty"`
	// Skill bracket assigned by Valve (Normal, High, Very High). If the skill is unknown, will return null.
	Skill NullableInt32 `json:"skill,omitempty"`
	// Average rank of players with public match data
	AverageRank NullableInt32 `json:"average_rank,omitempty"`
	// Experience Per Minute obtained by the player
	XpPerMin *int32 `json:"xp_per_min,omitempty"`
	// Average gold per minute of the player
	GoldPerMin *int32 `json:"gold_per_min,omitempty"`
	// Total hero damage to enemy heroes
	HeroDamage *int32 `json:"hero_damage,omitempty"`
	// Total healing of ally heroes
	HeroHealing *int32 `json:"hero_healing,omitempty"`
	// Total last hits the player had at the end of the match
	LastHits *int32 `json:"last_hits,omitempty"`
	// Integer corresponding to which lane the player laned in for the match
	Lane NullableInt32 `json:"lane,omitempty"`
	// lane_role
	LaneRole NullableInt32 `json:"lane_role,omitempty"`
	// Boolean describing whether or not the player roamed
	IsRoaming NullableBool `json:"is_roaming,omitempty"`
	// cluster
	Cluster *int32 `json:"cluster,omitempty"`
	// Integer describing whether or not the player left the game. 0: didn't leave. 1: left safely. 2+: Abandoned
	LeaverStatus *int32 `json:"leaver_status,omitempty"`
	// Size of the players party. If not in a party, will return 1.
	PartySize NullableInt32 `json:"party_size,omitempty"`
}

// NewPlayerRecentMatchesResponse instantiates a new PlayerRecentMatchesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlayerRecentMatchesResponse() *PlayerRecentMatchesResponse {
	this := PlayerRecentMatchesResponse{}
	return &this
}

// NewPlayerRecentMatchesResponseWithDefaults instantiates a new PlayerRecentMatchesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayerRecentMatchesResponseWithDefaults() *PlayerRecentMatchesResponse {
	this := PlayerRecentMatchesResponse{}
	return &this
}

// GetMatchId returns the MatchId field value if set, zero value otherwise.
func (o *PlayerRecentMatchesResponse) GetMatchId() int32 {
	if o == nil || IsNil(o.MatchId) {
		var ret int32
		return ret
	}
	return *o.MatchId
}

// GetMatchIdOk returns a tuple with the MatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerRecentMatchesResponse) GetMatchIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MatchId) {
		return nil, false
	}
	return o.MatchId, true
}

// HasMatchId returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasMatchId() bool {
	if o != nil && !IsNil(o.MatchId) {
		return true
	}

	return false
}

// SetMatchId gets a reference to the given int32 and assigns it to the MatchId field.
func (o *PlayerRecentMatchesResponse) SetMatchId(v int32) {
	o.MatchId = &v
}

// GetPlayerSlot returns the PlayerSlot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlayerRecentMatchesResponse) GetPlayerSlot() int32 {
	if o == nil || IsNil(o.PlayerSlot.Get()) {
		var ret int32
		return ret
	}
	return *o.PlayerSlot.Get()
}

// GetPlayerSlotOk returns a tuple with the PlayerSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayerRecentMatchesResponse) GetPlayerSlotOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlayerSlot.Get(), o.PlayerSlot.IsSet()
}

// HasPlayerSlot returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasPlayerSlot() bool {
	if o != nil && o.PlayerSlot.IsSet() {
		return true
	}

	return false
}

// SetPlayerSlot gets a reference to the given NullableInt32 and assigns it to the PlayerSlot field.
func (o *PlayerRecentMatchesResponse) SetPlayerSlot(v int32) {
	o.PlayerSlot.Set(&v)
}
// SetPlayerSlotNil sets the value for PlayerSlot to be an explicit nil
func (o *PlayerRecentMatchesResponse) SetPlayerSlotNil() {
	o.PlayerSlot.Set(nil)
}

// UnsetPlayerSlot ensures that no value is present for PlayerSlot, not even an explicit nil
func (o *PlayerRecentMatchesResponse) UnsetPlayerSlot() {
	o.PlayerSlot.Unset()
}

// GetRadiantWin returns the RadiantWin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlayerRecentMatchesResponse) GetRadiantWin() bool {
	if o == nil || IsNil(o.RadiantWin.Get()) {
		var ret bool
		return ret
	}
	return *o.RadiantWin.Get()
}

// GetRadiantWinOk returns a tuple with the RadiantWin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayerRecentMatchesResponse) GetRadiantWinOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RadiantWin.Get(), o.RadiantWin.IsSet()
}

// HasRadiantWin returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasRadiantWin() bool {
	if o != nil && o.RadiantWin.IsSet() {
		return true
	}

	return false
}

// SetRadiantWin gets a reference to the given NullableBool and assigns it to the RadiantWin field.
func (o *PlayerRecentMatchesResponse) SetRadiantWin(v bool) {
	o.RadiantWin.Set(&v)
}
// SetRadiantWinNil sets the value for RadiantWin to be an explicit nil
func (o *PlayerRecentMatchesResponse) SetRadiantWinNil() {
	o.RadiantWin.Set(nil)
}

// UnsetRadiantWin ensures that no value is present for RadiantWin, not even an explicit nil
func (o *PlayerRecentMatchesResponse) UnsetRadiantWin() {
	o.RadiantWin.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *PlayerRecentMatchesResponse) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerRecentMatchesResponse) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *PlayerRecentMatchesResponse) SetDuration(v int32) {
	o.Duration = &v
}

// GetGameMode returns the GameMode field value if set, zero value otherwise.
func (o *PlayerRecentMatchesResponse) GetGameMode() int32 {
	if o == nil || IsNil(o.GameMode) {
		var ret int32
		return ret
	}
	return *o.GameMode
}

// GetGameModeOk returns a tuple with the GameMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerRecentMatchesResponse) GetGameModeOk() (*int32, bool) {
	if o == nil || IsNil(o.GameMode) {
		return nil, false
	}
	return o.GameMode, true
}

// HasGameMode returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasGameMode() bool {
	if o != nil && !IsNil(o.GameMode) {
		return true
	}

	return false
}

// SetGameMode gets a reference to the given int32 and assigns it to the GameMode field.
func (o *PlayerRecentMatchesResponse) SetGameMode(v int32) {
	o.GameMode = &v
}

// GetLobbyType returns the LobbyType field value if set, zero value otherwise.
func (o *PlayerRecentMatchesResponse) GetLobbyType() int32 {
	if o == nil || IsNil(o.LobbyType) {
		var ret int32
		return ret
	}
	return *o.LobbyType
}

// GetLobbyTypeOk returns a tuple with the LobbyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerRecentMatchesResponse) GetLobbyTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.LobbyType) {
		return nil, false
	}
	return o.LobbyType, true
}

// HasLobbyType returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasLobbyType() bool {
	if o != nil && !IsNil(o.LobbyType) {
		return true
	}

	return false
}

// SetLobbyType gets a reference to the given int32 and assigns it to the LobbyType field.
func (o *PlayerRecentMatchesResponse) SetLobbyType(v int32) {
	o.LobbyType = &v
}

// GetHeroId returns the HeroId field value if set, zero value otherwise.
func (o *PlayerRecentMatchesResponse) GetHeroId() int32 {
	if o == nil || IsNil(o.HeroId) {
		var ret int32
		return ret
	}
	return *o.HeroId
}

// GetHeroIdOk returns a tuple with the HeroId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerRecentMatchesResponse) GetHeroIdOk() (*int32, bool) {
	if o == nil || IsNil(o.HeroId) {
		return nil, false
	}
	return o.HeroId, true
}

// HasHeroId returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasHeroId() bool {
	if o != nil && !IsNil(o.HeroId) {
		return true
	}

	return false
}

// SetHeroId gets a reference to the given int32 and assigns it to the HeroId field.
func (o *PlayerRecentMatchesResponse) SetHeroId(v int32) {
	o.HeroId = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *PlayerRecentMatchesResponse) GetStartTime() int32 {
	if o == nil || IsNil(o.StartTime) {
		var ret int32
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerRecentMatchesResponse) GetStartTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int32 and assigns it to the StartTime field.
func (o *PlayerRecentMatchesResponse) SetStartTime(v int32) {
	o.StartTime = &v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlayerRecentMatchesResponse) GetVersion() int32 {
	if o == nil || IsNil(o.Version.Get()) {
		var ret int32
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayerRecentMatchesResponse) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableInt32 and assigns it to the Version field.
func (o *PlayerRecentMatchesResponse) SetVersion(v int32) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *PlayerRecentMatchesResponse) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *PlayerRecentMatchesResponse) UnsetVersion() {
	o.Version.Unset()
}

// GetKills returns the Kills field value if set, zero value otherwise.
func (o *PlayerRecentMatchesResponse) GetKills() int32 {
	if o == nil || IsNil(o.Kills) {
		var ret int32
		return ret
	}
	return *o.Kills
}

// GetKillsOk returns a tuple with the Kills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerRecentMatchesResponse) GetKillsOk() (*int32, bool) {
	if o == nil || IsNil(o.Kills) {
		return nil, false
	}
	return o.Kills, true
}

// HasKills returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasKills() bool {
	if o != nil && !IsNil(o.Kills) {
		return true
	}

	return false
}

// SetKills gets a reference to the given int32 and assigns it to the Kills field.
func (o *PlayerRecentMatchesResponse) SetKills(v int32) {
	o.Kills = &v
}

// GetDeaths returns the Deaths field value if set, zero value otherwise.
func (o *PlayerRecentMatchesResponse) GetDeaths() int32 {
	if o == nil || IsNil(o.Deaths) {
		var ret int32
		return ret
	}
	return *o.Deaths
}

// GetDeathsOk returns a tuple with the Deaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerRecentMatchesResponse) GetDeathsOk() (*int32, bool) {
	if o == nil || IsNil(o.Deaths) {
		return nil, false
	}
	return o.Deaths, true
}

// HasDeaths returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasDeaths() bool {
	if o != nil && !IsNil(o.Deaths) {
		return true
	}

	return false
}

// SetDeaths gets a reference to the given int32 and assigns it to the Deaths field.
func (o *PlayerRecentMatchesResponse) SetDeaths(v int32) {
	o.Deaths = &v
}

// GetAssists returns the Assists field value if set, zero value otherwise.
func (o *PlayerRecentMatchesResponse) GetAssists() int32 {
	if o == nil || IsNil(o.Assists) {
		var ret int32
		return ret
	}
	return *o.Assists
}

// GetAssistsOk returns a tuple with the Assists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerRecentMatchesResponse) GetAssistsOk() (*int32, bool) {
	if o == nil || IsNil(o.Assists) {
		return nil, false
	}
	return o.Assists, true
}

// HasAssists returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasAssists() bool {
	if o != nil && !IsNil(o.Assists) {
		return true
	}

	return false
}

// SetAssists gets a reference to the given int32 and assigns it to the Assists field.
func (o *PlayerRecentMatchesResponse) SetAssists(v int32) {
	o.Assists = &v
}

// GetSkill returns the Skill field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlayerRecentMatchesResponse) GetSkill() int32 {
	if o == nil || IsNil(o.Skill.Get()) {
		var ret int32
		return ret
	}
	return *o.Skill.Get()
}

// GetSkillOk returns a tuple with the Skill field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayerRecentMatchesResponse) GetSkillOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Skill.Get(), o.Skill.IsSet()
}

// HasSkill returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasSkill() bool {
	if o != nil && o.Skill.IsSet() {
		return true
	}

	return false
}

// SetSkill gets a reference to the given NullableInt32 and assigns it to the Skill field.
func (o *PlayerRecentMatchesResponse) SetSkill(v int32) {
	o.Skill.Set(&v)
}
// SetSkillNil sets the value for Skill to be an explicit nil
func (o *PlayerRecentMatchesResponse) SetSkillNil() {
	o.Skill.Set(nil)
}

// UnsetSkill ensures that no value is present for Skill, not even an explicit nil
func (o *PlayerRecentMatchesResponse) UnsetSkill() {
	o.Skill.Unset()
}

// GetAverageRank returns the AverageRank field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlayerRecentMatchesResponse) GetAverageRank() int32 {
	if o == nil || IsNil(o.AverageRank.Get()) {
		var ret int32
		return ret
	}
	return *o.AverageRank.Get()
}

// GetAverageRankOk returns a tuple with the AverageRank field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayerRecentMatchesResponse) GetAverageRankOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AverageRank.Get(), o.AverageRank.IsSet()
}

// HasAverageRank returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasAverageRank() bool {
	if o != nil && o.AverageRank.IsSet() {
		return true
	}

	return false
}

// SetAverageRank gets a reference to the given NullableInt32 and assigns it to the AverageRank field.
func (o *PlayerRecentMatchesResponse) SetAverageRank(v int32) {
	o.AverageRank.Set(&v)
}
// SetAverageRankNil sets the value for AverageRank to be an explicit nil
func (o *PlayerRecentMatchesResponse) SetAverageRankNil() {
	o.AverageRank.Set(nil)
}

// UnsetAverageRank ensures that no value is present for AverageRank, not even an explicit nil
func (o *PlayerRecentMatchesResponse) UnsetAverageRank() {
	o.AverageRank.Unset()
}

// GetXpPerMin returns the XpPerMin field value if set, zero value otherwise.
func (o *PlayerRecentMatchesResponse) GetXpPerMin() int32 {
	if o == nil || IsNil(o.XpPerMin) {
		var ret int32
		return ret
	}
	return *o.XpPerMin
}

// GetXpPerMinOk returns a tuple with the XpPerMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerRecentMatchesResponse) GetXpPerMinOk() (*int32, bool) {
	if o == nil || IsNil(o.XpPerMin) {
		return nil, false
	}
	return o.XpPerMin, true
}

// HasXpPerMin returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasXpPerMin() bool {
	if o != nil && !IsNil(o.XpPerMin) {
		return true
	}

	return false
}

// SetXpPerMin gets a reference to the given int32 and assigns it to the XpPerMin field.
func (o *PlayerRecentMatchesResponse) SetXpPerMin(v int32) {
	o.XpPerMin = &v
}

// GetGoldPerMin returns the GoldPerMin field value if set, zero value otherwise.
func (o *PlayerRecentMatchesResponse) GetGoldPerMin() int32 {
	if o == nil || IsNil(o.GoldPerMin) {
		var ret int32
		return ret
	}
	return *o.GoldPerMin
}

// GetGoldPerMinOk returns a tuple with the GoldPerMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerRecentMatchesResponse) GetGoldPerMinOk() (*int32, bool) {
	if o == nil || IsNil(o.GoldPerMin) {
		return nil, false
	}
	return o.GoldPerMin, true
}

// HasGoldPerMin returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasGoldPerMin() bool {
	if o != nil && !IsNil(o.GoldPerMin) {
		return true
	}

	return false
}

// SetGoldPerMin gets a reference to the given int32 and assigns it to the GoldPerMin field.
func (o *PlayerRecentMatchesResponse) SetGoldPerMin(v int32) {
	o.GoldPerMin = &v
}

// GetHeroDamage returns the HeroDamage field value if set, zero value otherwise.
func (o *PlayerRecentMatchesResponse) GetHeroDamage() int32 {
	if o == nil || IsNil(o.HeroDamage) {
		var ret int32
		return ret
	}
	return *o.HeroDamage
}

// GetHeroDamageOk returns a tuple with the HeroDamage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerRecentMatchesResponse) GetHeroDamageOk() (*int32, bool) {
	if o == nil || IsNil(o.HeroDamage) {
		return nil, false
	}
	return o.HeroDamage, true
}

// HasHeroDamage returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasHeroDamage() bool {
	if o != nil && !IsNil(o.HeroDamage) {
		return true
	}

	return false
}

// SetHeroDamage gets a reference to the given int32 and assigns it to the HeroDamage field.
func (o *PlayerRecentMatchesResponse) SetHeroDamage(v int32) {
	o.HeroDamage = &v
}

// GetHeroHealing returns the HeroHealing field value if set, zero value otherwise.
func (o *PlayerRecentMatchesResponse) GetHeroHealing() int32 {
	if o == nil || IsNil(o.HeroHealing) {
		var ret int32
		return ret
	}
	return *o.HeroHealing
}

// GetHeroHealingOk returns a tuple with the HeroHealing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerRecentMatchesResponse) GetHeroHealingOk() (*int32, bool) {
	if o == nil || IsNil(o.HeroHealing) {
		return nil, false
	}
	return o.HeroHealing, true
}

// HasHeroHealing returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasHeroHealing() bool {
	if o != nil && !IsNil(o.HeroHealing) {
		return true
	}

	return false
}

// SetHeroHealing gets a reference to the given int32 and assigns it to the HeroHealing field.
func (o *PlayerRecentMatchesResponse) SetHeroHealing(v int32) {
	o.HeroHealing = &v
}

// GetLastHits returns the LastHits field value if set, zero value otherwise.
func (o *PlayerRecentMatchesResponse) GetLastHits() int32 {
	if o == nil || IsNil(o.LastHits) {
		var ret int32
		return ret
	}
	return *o.LastHits
}

// GetLastHitsOk returns a tuple with the LastHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerRecentMatchesResponse) GetLastHitsOk() (*int32, bool) {
	if o == nil || IsNil(o.LastHits) {
		return nil, false
	}
	return o.LastHits, true
}

// HasLastHits returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasLastHits() bool {
	if o != nil && !IsNil(o.LastHits) {
		return true
	}

	return false
}

// SetLastHits gets a reference to the given int32 and assigns it to the LastHits field.
func (o *PlayerRecentMatchesResponse) SetLastHits(v int32) {
	o.LastHits = &v
}

// GetLane returns the Lane field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlayerRecentMatchesResponse) GetLane() int32 {
	if o == nil || IsNil(o.Lane.Get()) {
		var ret int32
		return ret
	}
	return *o.Lane.Get()
}

// GetLaneOk returns a tuple with the Lane field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayerRecentMatchesResponse) GetLaneOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lane.Get(), o.Lane.IsSet()
}

// HasLane returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasLane() bool {
	if o != nil && o.Lane.IsSet() {
		return true
	}

	return false
}

// SetLane gets a reference to the given NullableInt32 and assigns it to the Lane field.
func (o *PlayerRecentMatchesResponse) SetLane(v int32) {
	o.Lane.Set(&v)
}
// SetLaneNil sets the value for Lane to be an explicit nil
func (o *PlayerRecentMatchesResponse) SetLaneNil() {
	o.Lane.Set(nil)
}

// UnsetLane ensures that no value is present for Lane, not even an explicit nil
func (o *PlayerRecentMatchesResponse) UnsetLane() {
	o.Lane.Unset()
}

// GetLaneRole returns the LaneRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlayerRecentMatchesResponse) GetLaneRole() int32 {
	if o == nil || IsNil(o.LaneRole.Get()) {
		var ret int32
		return ret
	}
	return *o.LaneRole.Get()
}

// GetLaneRoleOk returns a tuple with the LaneRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayerRecentMatchesResponse) GetLaneRoleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LaneRole.Get(), o.LaneRole.IsSet()
}

// HasLaneRole returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasLaneRole() bool {
	if o != nil && o.LaneRole.IsSet() {
		return true
	}

	return false
}

// SetLaneRole gets a reference to the given NullableInt32 and assigns it to the LaneRole field.
func (o *PlayerRecentMatchesResponse) SetLaneRole(v int32) {
	o.LaneRole.Set(&v)
}
// SetLaneRoleNil sets the value for LaneRole to be an explicit nil
func (o *PlayerRecentMatchesResponse) SetLaneRoleNil() {
	o.LaneRole.Set(nil)
}

// UnsetLaneRole ensures that no value is present for LaneRole, not even an explicit nil
func (o *PlayerRecentMatchesResponse) UnsetLaneRole() {
	o.LaneRole.Unset()
}

// GetIsRoaming returns the IsRoaming field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlayerRecentMatchesResponse) GetIsRoaming() bool {
	if o == nil || IsNil(o.IsRoaming.Get()) {
		var ret bool
		return ret
	}
	return *o.IsRoaming.Get()
}

// GetIsRoamingOk returns a tuple with the IsRoaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayerRecentMatchesResponse) GetIsRoamingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRoaming.Get(), o.IsRoaming.IsSet()
}

// HasIsRoaming returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasIsRoaming() bool {
	if o != nil && o.IsRoaming.IsSet() {
		return true
	}

	return false
}

// SetIsRoaming gets a reference to the given NullableBool and assigns it to the IsRoaming field.
func (o *PlayerRecentMatchesResponse) SetIsRoaming(v bool) {
	o.IsRoaming.Set(&v)
}
// SetIsRoamingNil sets the value for IsRoaming to be an explicit nil
func (o *PlayerRecentMatchesResponse) SetIsRoamingNil() {
	o.IsRoaming.Set(nil)
}

// UnsetIsRoaming ensures that no value is present for IsRoaming, not even an explicit nil
func (o *PlayerRecentMatchesResponse) UnsetIsRoaming() {
	o.IsRoaming.Unset()
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *PlayerRecentMatchesResponse) GetCluster() int32 {
	if o == nil || IsNil(o.Cluster) {
		var ret int32
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerRecentMatchesResponse) GetClusterOk() (*int32, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given int32 and assigns it to the Cluster field.
func (o *PlayerRecentMatchesResponse) SetCluster(v int32) {
	o.Cluster = &v
}

// GetLeaverStatus returns the LeaverStatus field value if set, zero value otherwise.
func (o *PlayerRecentMatchesResponse) GetLeaverStatus() int32 {
	if o == nil || IsNil(o.LeaverStatus) {
		var ret int32
		return ret
	}
	return *o.LeaverStatus
}

// GetLeaverStatusOk returns a tuple with the LeaverStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerRecentMatchesResponse) GetLeaverStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.LeaverStatus) {
		return nil, false
	}
	return o.LeaverStatus, true
}

// HasLeaverStatus returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasLeaverStatus() bool {
	if o != nil && !IsNil(o.LeaverStatus) {
		return true
	}

	return false
}

// SetLeaverStatus gets a reference to the given int32 and assigns it to the LeaverStatus field.
func (o *PlayerRecentMatchesResponse) SetLeaverStatus(v int32) {
	o.LeaverStatus = &v
}

// GetPartySize returns the PartySize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlayerRecentMatchesResponse) GetPartySize() int32 {
	if o == nil || IsNil(o.PartySize.Get()) {
		var ret int32
		return ret
	}
	return *o.PartySize.Get()
}

// GetPartySizeOk returns a tuple with the PartySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayerRecentMatchesResponse) GetPartySizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartySize.Get(), o.PartySize.IsSet()
}

// HasPartySize returns a boolean if a field has been set.
func (o *PlayerRecentMatchesResponse) HasPartySize() bool {
	if o != nil && o.PartySize.IsSet() {
		return true
	}

	return false
}

// SetPartySize gets a reference to the given NullableInt32 and assigns it to the PartySize field.
func (o *PlayerRecentMatchesResponse) SetPartySize(v int32) {
	o.PartySize.Set(&v)
}
// SetPartySizeNil sets the value for PartySize to be an explicit nil
func (o *PlayerRecentMatchesResponse) SetPartySizeNil() {
	o.PartySize.Set(nil)
}

// UnsetPartySize ensures that no value is present for PartySize, not even an explicit nil
func (o *PlayerRecentMatchesResponse) UnsetPartySize() {
	o.PartySize.Unset()
}

func (o PlayerRecentMatchesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlayerRecentMatchesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MatchId) {
		toSerialize["match_id"] = o.MatchId
	}
	if o.PlayerSlot.IsSet() {
		toSerialize["player_slot"] = o.PlayerSlot.Get()
	}
	if o.RadiantWin.IsSet() {
		toSerialize["radiant_win"] = o.RadiantWin.Get()
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.GameMode) {
		toSerialize["game_mode"] = o.GameMode
	}
	if !IsNil(o.LobbyType) {
		toSerialize["lobby_type"] = o.LobbyType
	}
	if !IsNil(o.HeroId) {
		toSerialize["hero_id"] = o.HeroId
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if !IsNil(o.Kills) {
		toSerialize["kills"] = o.Kills
	}
	if !IsNil(o.Deaths) {
		toSerialize["deaths"] = o.Deaths
	}
	if !IsNil(o.Assists) {
		toSerialize["assists"] = o.Assists
	}
	if o.Skill.IsSet() {
		toSerialize["skill"] = o.Skill.Get()
	}
	if o.AverageRank.IsSet() {
		toSerialize["average_rank"] = o.AverageRank.Get()
	}
	if !IsNil(o.XpPerMin) {
		toSerialize["xp_per_min"] = o.XpPerMin
	}
	if !IsNil(o.GoldPerMin) {
		toSerialize["gold_per_min"] = o.GoldPerMin
	}
	if !IsNil(o.HeroDamage) {
		toSerialize["hero_damage"] = o.HeroDamage
	}
	if !IsNil(o.HeroHealing) {
		toSerialize["hero_healing"] = o.HeroHealing
	}
	if !IsNil(o.LastHits) {
		toSerialize["last_hits"] = o.LastHits
	}
	if o.Lane.IsSet() {
		toSerialize["lane"] = o.Lane.Get()
	}
	if o.LaneRole.IsSet() {
		toSerialize["lane_role"] = o.LaneRole.Get()
	}
	if o.IsRoaming.IsSet() {
		toSerialize["is_roaming"] = o.IsRoaming.Get()
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.LeaverStatus) {
		toSerialize["leaver_status"] = o.LeaverStatus
	}
	if o.PartySize.IsSet() {
		toSerialize["party_size"] = o.PartySize.Get()
	}
	return toSerialize, nil
}

type NullablePlayerRecentMatchesResponse struct {
	value *PlayerRecentMatchesResponse
	isSet bool
}

func (v NullablePlayerRecentMatchesResponse) Get() *PlayerRecentMatchesResponse {
	return v.value
}

func (v *NullablePlayerRecentMatchesResponse) Set(val *PlayerRecentMatchesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayerRecentMatchesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayerRecentMatchesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayerRecentMatchesResponse(val *PlayerRecentMatchesResponse) *NullablePlayerRecentMatchesResponse {
	return &NullablePlayerRecentMatchesResponse{value: val, isSet: true}
}

func (v NullablePlayerRecentMatchesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayerRecentMatchesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


