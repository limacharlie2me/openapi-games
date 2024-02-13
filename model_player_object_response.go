/*
OpenDota API

# Introduction The OpenDota API provides Dota 2 related data including advanced match data extracted from match replays.  You can find data that can be used to convert hero and ability IDs and other information provided by the API from the [dotaconstants](https://github.com/odota/dotaconstants) repository.  Without a key, you can make 2,000 free calls per day at a rate limit of 60 requests/minute. We also offer a Premium Tier with unlimited API calls and higher rate limits. Check out the [API page](https://www.opendota.com/api-keys) to learn more.     

API version: 25.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the PlayerObjectResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlayerObjectResponse{}

// PlayerObjectResponse struct for PlayerObjectResponse
type PlayerObjectResponse struct {
	// The player account ID
	AccountId *int32 `json:"account_id,omitempty"`
	// Player's steam identifier
	Steamid *string `json:"steamid,omitempty"`
	// Steam picture URL (small picture)
	Avatar *string `json:"avatar,omitempty"`
	// Steam picture URL (medium picture)
	Avatarmedium *string `json:"avatarmedium,omitempty"`
	// Steam picture URL (full picture)
	Avatarfull *string `json:"avatarfull,omitempty"`
	// Steam profile URL
	Profileurl *string `json:"profileurl,omitempty"`
	// Player's Steam name
	Personaname NullableString `json:"personaname,omitempty"`
	// Date and time of last login to OpenDota
	LastLogin *time.Time `json:"last_login,omitempty"`
	// Date and time of last request to refresh player's match history
	FullHistoryTime *time.Time `json:"full_history_time,omitempty"`
	// Amount of dollars the player has donated to OpenDota
	Cheese *int32 `json:"cheese,omitempty"`
	// Whether the refresh of player' match history failed
	FhUnavailable *bool `json:"fh_unavailable,omitempty"`
	// Player's country identifier, e.g. US
	Loccountrycode *string `json:"loccountrycode,omitempty"`
	// Verified player name, e.g. 'Miracle-'
	Name *string `json:"name,omitempty"`
	// Player's country code
	CountryCode *string `json:"country_code,omitempty"`
	// Player's ingame role (core: 1 or support: 2)
	FantasyRole *int32 `json:"fantasy_role,omitempty"`
	// Player's team identifier
	TeamId *int32 `json:"team_id,omitempty"`
	// Team name
	TeamName NullableString `json:"team_name,omitempty"`
	// Player's team shorthand tag, e.g. 'EG'
	TeamTag *string `json:"team_tag,omitempty"`
	// Whether the roster lock is active
	IsLocked *bool `json:"is_locked,omitempty"`
	// Whether the player is professional or not
	IsPro *bool `json:"is_pro,omitempty"`
	// When the roster lock will end
	LockedUntil *int32 `json:"locked_until,omitempty"`
}

// NewPlayerObjectResponse instantiates a new PlayerObjectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlayerObjectResponse() *PlayerObjectResponse {
	this := PlayerObjectResponse{}
	return &this
}

// NewPlayerObjectResponseWithDefaults instantiates a new PlayerObjectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayerObjectResponseWithDefaults() *PlayerObjectResponse {
	this := PlayerObjectResponse{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *PlayerObjectResponse) GetAccountId() int32 {
	if o == nil || IsNil(o.AccountId) {
		var ret int32
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerObjectResponse) GetAccountIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.
func (o *PlayerObjectResponse) SetAccountId(v int32) {
	o.AccountId = &v
}

// GetSteamid returns the Steamid field value if set, zero value otherwise.
func (o *PlayerObjectResponse) GetSteamid() string {
	if o == nil || IsNil(o.Steamid) {
		var ret string
		return ret
	}
	return *o.Steamid
}

// GetSteamidOk returns a tuple with the Steamid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerObjectResponse) GetSteamidOk() (*string, bool) {
	if o == nil || IsNil(o.Steamid) {
		return nil, false
	}
	return o.Steamid, true
}

// HasSteamid returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasSteamid() bool {
	if o != nil && !IsNil(o.Steamid) {
		return true
	}

	return false
}

// SetSteamid gets a reference to the given string and assigns it to the Steamid field.
func (o *PlayerObjectResponse) SetSteamid(v string) {
	o.Steamid = &v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *PlayerObjectResponse) GetAvatar() string {
	if o == nil || IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerObjectResponse) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *PlayerObjectResponse) SetAvatar(v string) {
	o.Avatar = &v
}

// GetAvatarmedium returns the Avatarmedium field value if set, zero value otherwise.
func (o *PlayerObjectResponse) GetAvatarmedium() string {
	if o == nil || IsNil(o.Avatarmedium) {
		var ret string
		return ret
	}
	return *o.Avatarmedium
}

// GetAvatarmediumOk returns a tuple with the Avatarmedium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerObjectResponse) GetAvatarmediumOk() (*string, bool) {
	if o == nil || IsNil(o.Avatarmedium) {
		return nil, false
	}
	return o.Avatarmedium, true
}

// HasAvatarmedium returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasAvatarmedium() bool {
	if o != nil && !IsNil(o.Avatarmedium) {
		return true
	}

	return false
}

// SetAvatarmedium gets a reference to the given string and assigns it to the Avatarmedium field.
func (o *PlayerObjectResponse) SetAvatarmedium(v string) {
	o.Avatarmedium = &v
}

// GetAvatarfull returns the Avatarfull field value if set, zero value otherwise.
func (o *PlayerObjectResponse) GetAvatarfull() string {
	if o == nil || IsNil(o.Avatarfull) {
		var ret string
		return ret
	}
	return *o.Avatarfull
}

// GetAvatarfullOk returns a tuple with the Avatarfull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerObjectResponse) GetAvatarfullOk() (*string, bool) {
	if o == nil || IsNil(o.Avatarfull) {
		return nil, false
	}
	return o.Avatarfull, true
}

// HasAvatarfull returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasAvatarfull() bool {
	if o != nil && !IsNil(o.Avatarfull) {
		return true
	}

	return false
}

// SetAvatarfull gets a reference to the given string and assigns it to the Avatarfull field.
func (o *PlayerObjectResponse) SetAvatarfull(v string) {
	o.Avatarfull = &v
}

// GetProfileurl returns the Profileurl field value if set, zero value otherwise.
func (o *PlayerObjectResponse) GetProfileurl() string {
	if o == nil || IsNil(o.Profileurl) {
		var ret string
		return ret
	}
	return *o.Profileurl
}

// GetProfileurlOk returns a tuple with the Profileurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerObjectResponse) GetProfileurlOk() (*string, bool) {
	if o == nil || IsNil(o.Profileurl) {
		return nil, false
	}
	return o.Profileurl, true
}

// HasProfileurl returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasProfileurl() bool {
	if o != nil && !IsNil(o.Profileurl) {
		return true
	}

	return false
}

// SetProfileurl gets a reference to the given string and assigns it to the Profileurl field.
func (o *PlayerObjectResponse) SetProfileurl(v string) {
	o.Profileurl = &v
}

// GetPersonaname returns the Personaname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlayerObjectResponse) GetPersonaname() string {
	if o == nil || IsNil(o.Personaname.Get()) {
		var ret string
		return ret
	}
	return *o.Personaname.Get()
}

// GetPersonanameOk returns a tuple with the Personaname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayerObjectResponse) GetPersonanameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Personaname.Get(), o.Personaname.IsSet()
}

// HasPersonaname returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasPersonaname() bool {
	if o != nil && o.Personaname.IsSet() {
		return true
	}

	return false
}

// SetPersonaname gets a reference to the given NullableString and assigns it to the Personaname field.
func (o *PlayerObjectResponse) SetPersonaname(v string) {
	o.Personaname.Set(&v)
}
// SetPersonanameNil sets the value for Personaname to be an explicit nil
func (o *PlayerObjectResponse) SetPersonanameNil() {
	o.Personaname.Set(nil)
}

// UnsetPersonaname ensures that no value is present for Personaname, not even an explicit nil
func (o *PlayerObjectResponse) UnsetPersonaname() {
	o.Personaname.Unset()
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *PlayerObjectResponse) GetLastLogin() time.Time {
	if o == nil || IsNil(o.LastLogin) {
		var ret time.Time
		return ret
	}
	return *o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerObjectResponse) GetLastLoginOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastLogin) {
		return nil, false
	}
	return o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasLastLogin() bool {
	if o != nil && !IsNil(o.LastLogin) {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given time.Time and assigns it to the LastLogin field.
func (o *PlayerObjectResponse) SetLastLogin(v time.Time) {
	o.LastLogin = &v
}

// GetFullHistoryTime returns the FullHistoryTime field value if set, zero value otherwise.
func (o *PlayerObjectResponse) GetFullHistoryTime() time.Time {
	if o == nil || IsNil(o.FullHistoryTime) {
		var ret time.Time
		return ret
	}
	return *o.FullHistoryTime
}

// GetFullHistoryTimeOk returns a tuple with the FullHistoryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerObjectResponse) GetFullHistoryTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FullHistoryTime) {
		return nil, false
	}
	return o.FullHistoryTime, true
}

// HasFullHistoryTime returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasFullHistoryTime() bool {
	if o != nil && !IsNil(o.FullHistoryTime) {
		return true
	}

	return false
}

// SetFullHistoryTime gets a reference to the given time.Time and assigns it to the FullHistoryTime field.
func (o *PlayerObjectResponse) SetFullHistoryTime(v time.Time) {
	o.FullHistoryTime = &v
}

// GetCheese returns the Cheese field value if set, zero value otherwise.
func (o *PlayerObjectResponse) GetCheese() int32 {
	if o == nil || IsNil(o.Cheese) {
		var ret int32
		return ret
	}
	return *o.Cheese
}

// GetCheeseOk returns a tuple with the Cheese field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerObjectResponse) GetCheeseOk() (*int32, bool) {
	if o == nil || IsNil(o.Cheese) {
		return nil, false
	}
	return o.Cheese, true
}

// HasCheese returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasCheese() bool {
	if o != nil && !IsNil(o.Cheese) {
		return true
	}

	return false
}

// SetCheese gets a reference to the given int32 and assigns it to the Cheese field.
func (o *PlayerObjectResponse) SetCheese(v int32) {
	o.Cheese = &v
}

// GetFhUnavailable returns the FhUnavailable field value if set, zero value otherwise.
func (o *PlayerObjectResponse) GetFhUnavailable() bool {
	if o == nil || IsNil(o.FhUnavailable) {
		var ret bool
		return ret
	}
	return *o.FhUnavailable
}

// GetFhUnavailableOk returns a tuple with the FhUnavailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerObjectResponse) GetFhUnavailableOk() (*bool, bool) {
	if o == nil || IsNil(o.FhUnavailable) {
		return nil, false
	}
	return o.FhUnavailable, true
}

// HasFhUnavailable returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasFhUnavailable() bool {
	if o != nil && !IsNil(o.FhUnavailable) {
		return true
	}

	return false
}

// SetFhUnavailable gets a reference to the given bool and assigns it to the FhUnavailable field.
func (o *PlayerObjectResponse) SetFhUnavailable(v bool) {
	o.FhUnavailable = &v
}

// GetLoccountrycode returns the Loccountrycode field value if set, zero value otherwise.
func (o *PlayerObjectResponse) GetLoccountrycode() string {
	if o == nil || IsNil(o.Loccountrycode) {
		var ret string
		return ret
	}
	return *o.Loccountrycode
}

// GetLoccountrycodeOk returns a tuple with the Loccountrycode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerObjectResponse) GetLoccountrycodeOk() (*string, bool) {
	if o == nil || IsNil(o.Loccountrycode) {
		return nil, false
	}
	return o.Loccountrycode, true
}

// HasLoccountrycode returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasLoccountrycode() bool {
	if o != nil && !IsNil(o.Loccountrycode) {
		return true
	}

	return false
}

// SetLoccountrycode gets a reference to the given string and assigns it to the Loccountrycode field.
func (o *PlayerObjectResponse) SetLoccountrycode(v string) {
	o.Loccountrycode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PlayerObjectResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerObjectResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PlayerObjectResponse) SetName(v string) {
	o.Name = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *PlayerObjectResponse) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerObjectResponse) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *PlayerObjectResponse) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetFantasyRole returns the FantasyRole field value if set, zero value otherwise.
func (o *PlayerObjectResponse) GetFantasyRole() int32 {
	if o == nil || IsNil(o.FantasyRole) {
		var ret int32
		return ret
	}
	return *o.FantasyRole
}

// GetFantasyRoleOk returns a tuple with the FantasyRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerObjectResponse) GetFantasyRoleOk() (*int32, bool) {
	if o == nil || IsNil(o.FantasyRole) {
		return nil, false
	}
	return o.FantasyRole, true
}

// HasFantasyRole returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasFantasyRole() bool {
	if o != nil && !IsNil(o.FantasyRole) {
		return true
	}

	return false
}

// SetFantasyRole gets a reference to the given int32 and assigns it to the FantasyRole field.
func (o *PlayerObjectResponse) SetFantasyRole(v int32) {
	o.FantasyRole = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *PlayerObjectResponse) GetTeamId() int32 {
	if o == nil || IsNil(o.TeamId) {
		var ret int32
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerObjectResponse) GetTeamIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TeamId) {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasTeamId() bool {
	if o != nil && !IsNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given int32 and assigns it to the TeamId field.
func (o *PlayerObjectResponse) SetTeamId(v int32) {
	o.TeamId = &v
}

// GetTeamName returns the TeamName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlayerObjectResponse) GetTeamName() string {
	if o == nil || IsNil(o.TeamName.Get()) {
		var ret string
		return ret
	}
	return *o.TeamName.Get()
}

// GetTeamNameOk returns a tuple with the TeamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayerObjectResponse) GetTeamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TeamName.Get(), o.TeamName.IsSet()
}

// HasTeamName returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasTeamName() bool {
	if o != nil && o.TeamName.IsSet() {
		return true
	}

	return false
}

// SetTeamName gets a reference to the given NullableString and assigns it to the TeamName field.
func (o *PlayerObjectResponse) SetTeamName(v string) {
	o.TeamName.Set(&v)
}
// SetTeamNameNil sets the value for TeamName to be an explicit nil
func (o *PlayerObjectResponse) SetTeamNameNil() {
	o.TeamName.Set(nil)
}

// UnsetTeamName ensures that no value is present for TeamName, not even an explicit nil
func (o *PlayerObjectResponse) UnsetTeamName() {
	o.TeamName.Unset()
}

// GetTeamTag returns the TeamTag field value if set, zero value otherwise.
func (o *PlayerObjectResponse) GetTeamTag() string {
	if o == nil || IsNil(o.TeamTag) {
		var ret string
		return ret
	}
	return *o.TeamTag
}

// GetTeamTagOk returns a tuple with the TeamTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerObjectResponse) GetTeamTagOk() (*string, bool) {
	if o == nil || IsNil(o.TeamTag) {
		return nil, false
	}
	return o.TeamTag, true
}

// HasTeamTag returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasTeamTag() bool {
	if o != nil && !IsNil(o.TeamTag) {
		return true
	}

	return false
}

// SetTeamTag gets a reference to the given string and assigns it to the TeamTag field.
func (o *PlayerObjectResponse) SetTeamTag(v string) {
	o.TeamTag = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *PlayerObjectResponse) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerObjectResponse) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *PlayerObjectResponse) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetIsPro returns the IsPro field value if set, zero value otherwise.
func (o *PlayerObjectResponse) GetIsPro() bool {
	if o == nil || IsNil(o.IsPro) {
		var ret bool
		return ret
	}
	return *o.IsPro
}

// GetIsProOk returns a tuple with the IsPro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerObjectResponse) GetIsProOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPro) {
		return nil, false
	}
	return o.IsPro, true
}

// HasIsPro returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasIsPro() bool {
	if o != nil && !IsNil(o.IsPro) {
		return true
	}

	return false
}

// SetIsPro gets a reference to the given bool and assigns it to the IsPro field.
func (o *PlayerObjectResponse) SetIsPro(v bool) {
	o.IsPro = &v
}

// GetLockedUntil returns the LockedUntil field value if set, zero value otherwise.
func (o *PlayerObjectResponse) GetLockedUntil() int32 {
	if o == nil || IsNil(o.LockedUntil) {
		var ret int32
		return ret
	}
	return *o.LockedUntil
}

// GetLockedUntilOk returns a tuple with the LockedUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerObjectResponse) GetLockedUntilOk() (*int32, bool) {
	if o == nil || IsNil(o.LockedUntil) {
		return nil, false
	}
	return o.LockedUntil, true
}

// HasLockedUntil returns a boolean if a field has been set.
func (o *PlayerObjectResponse) HasLockedUntil() bool {
	if o != nil && !IsNil(o.LockedUntil) {
		return true
	}

	return false
}

// SetLockedUntil gets a reference to the given int32 and assigns it to the LockedUntil field.
func (o *PlayerObjectResponse) SetLockedUntil(v int32) {
	o.LockedUntil = &v
}

func (o PlayerObjectResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlayerObjectResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.Steamid) {
		toSerialize["steamid"] = o.Steamid
	}
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}
	if !IsNil(o.Avatarmedium) {
		toSerialize["avatarmedium"] = o.Avatarmedium
	}
	if !IsNil(o.Avatarfull) {
		toSerialize["avatarfull"] = o.Avatarfull
	}
	if !IsNil(o.Profileurl) {
		toSerialize["profileurl"] = o.Profileurl
	}
	if o.Personaname.IsSet() {
		toSerialize["personaname"] = o.Personaname.Get()
	}
	if !IsNil(o.LastLogin) {
		toSerialize["last_login"] = o.LastLogin
	}
	if !IsNil(o.FullHistoryTime) {
		toSerialize["full_history_time"] = o.FullHistoryTime
	}
	if !IsNil(o.Cheese) {
		toSerialize["cheese"] = o.Cheese
	}
	if !IsNil(o.FhUnavailable) {
		toSerialize["fh_unavailable"] = o.FhUnavailable
	}
	if !IsNil(o.Loccountrycode) {
		toSerialize["loccountrycode"] = o.Loccountrycode
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}
	if !IsNil(o.FantasyRole) {
		toSerialize["fantasy_role"] = o.FantasyRole
	}
	if !IsNil(o.TeamId) {
		toSerialize["team_id"] = o.TeamId
	}
	if o.TeamName.IsSet() {
		toSerialize["team_name"] = o.TeamName.Get()
	}
	if !IsNil(o.TeamTag) {
		toSerialize["team_tag"] = o.TeamTag
	}
	if !IsNil(o.IsLocked) {
		toSerialize["is_locked"] = o.IsLocked
	}
	if !IsNil(o.IsPro) {
		toSerialize["is_pro"] = o.IsPro
	}
	if !IsNil(o.LockedUntil) {
		toSerialize["locked_until"] = o.LockedUntil
	}
	return toSerialize, nil
}

type NullablePlayerObjectResponse struct {
	value *PlayerObjectResponse
	isSet bool
}

func (v NullablePlayerObjectResponse) Get() *PlayerObjectResponse {
	return v.value
}

func (v *NullablePlayerObjectResponse) Set(val *PlayerObjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayerObjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayerObjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayerObjectResponse(val *PlayerObjectResponse) *NullablePlayerObjectResponse {
	return &NullablePlayerObjectResponse{value: val, isSet: true}
}

func (v NullablePlayerObjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayerObjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

