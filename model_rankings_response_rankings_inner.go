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

// checks if the RankingsResponseRankingsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RankingsResponseRankingsInner{}

// RankingsResponseRankingsInner struct for RankingsResponseRankingsInner
type RankingsResponseRankingsInner struct {
	// The player account ID
	AccountId *int32 `json:"account_id,omitempty"`
	// Score
	Score *float32 `json:"score,omitempty"`
	// steamid
	Steamid NullableString `json:"steamid,omitempty"`
	// avatar
	Avatar NullableString `json:"avatar,omitempty"`
	// avatarmedium
	Avatarmedium NullableString `json:"avatarmedium,omitempty"`
	// avatarfull
	Avatarfull NullableString `json:"avatarfull,omitempty"`
	// profileurl
	Profileurl NullableString `json:"profileurl,omitempty"`
	// Player's Steam name
	Personaname NullableString `json:"personaname,omitempty"`
	// last_login
	LastLogin NullableTime `json:"last_login,omitempty"`
	// full_history_time
	FullHistoryTime *time.Time `json:"full_history_time,omitempty"`
	// cheese
	Cheese NullableInt32 `json:"cheese,omitempty"`
	// fh_unavailable
	FhUnavailable NullableBool `json:"fh_unavailable,omitempty"`
	// loccountrycode
	Loccountrycode NullableString `json:"loccountrycode,omitempty"`
	// rank_tier
	RankTier NullableInt32 `json:"rank_tier,omitempty"`
}

// NewRankingsResponseRankingsInner instantiates a new RankingsResponseRankingsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRankingsResponseRankingsInner() *RankingsResponseRankingsInner {
	this := RankingsResponseRankingsInner{}
	return &this
}

// NewRankingsResponseRankingsInnerWithDefaults instantiates a new RankingsResponseRankingsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRankingsResponseRankingsInnerWithDefaults() *RankingsResponseRankingsInner {
	this := RankingsResponseRankingsInner{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *RankingsResponseRankingsInner) GetAccountId() int32 {
	if o == nil || IsNil(o.AccountId) {
		var ret int32
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RankingsResponseRankingsInner) GetAccountIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *RankingsResponseRankingsInner) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.
func (o *RankingsResponseRankingsInner) SetAccountId(v int32) {
	o.AccountId = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *RankingsResponseRankingsInner) GetScore() float32 {
	if o == nil || IsNil(o.Score) {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RankingsResponseRankingsInner) GetScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *RankingsResponseRankingsInner) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *RankingsResponseRankingsInner) SetScore(v float32) {
	o.Score = &v
}

// GetSteamid returns the Steamid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RankingsResponseRankingsInner) GetSteamid() string {
	if o == nil || IsNil(o.Steamid.Get()) {
		var ret string
		return ret
	}
	return *o.Steamid.Get()
}

// GetSteamidOk returns a tuple with the Steamid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RankingsResponseRankingsInner) GetSteamidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Steamid.Get(), o.Steamid.IsSet()
}

// HasSteamid returns a boolean if a field has been set.
func (o *RankingsResponseRankingsInner) HasSteamid() bool {
	if o != nil && o.Steamid.IsSet() {
		return true
	}

	return false
}

// SetSteamid gets a reference to the given NullableString and assigns it to the Steamid field.
func (o *RankingsResponseRankingsInner) SetSteamid(v string) {
	o.Steamid.Set(&v)
}
// SetSteamidNil sets the value for Steamid to be an explicit nil
func (o *RankingsResponseRankingsInner) SetSteamidNil() {
	o.Steamid.Set(nil)
}

// UnsetSteamid ensures that no value is present for Steamid, not even an explicit nil
func (o *RankingsResponseRankingsInner) UnsetSteamid() {
	o.Steamid.Unset()
}

// GetAvatar returns the Avatar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RankingsResponseRankingsInner) GetAvatar() string {
	if o == nil || IsNil(o.Avatar.Get()) {
		var ret string
		return ret
	}
	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RankingsResponseRankingsInner) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// HasAvatar returns a boolean if a field has been set.
func (o *RankingsResponseRankingsInner) HasAvatar() bool {
	if o != nil && o.Avatar.IsSet() {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given NullableString and assigns it to the Avatar field.
func (o *RankingsResponseRankingsInner) SetAvatar(v string) {
	o.Avatar.Set(&v)
}
// SetAvatarNil sets the value for Avatar to be an explicit nil
func (o *RankingsResponseRankingsInner) SetAvatarNil() {
	o.Avatar.Set(nil)
}

// UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
func (o *RankingsResponseRankingsInner) UnsetAvatar() {
	o.Avatar.Unset()
}

// GetAvatarmedium returns the Avatarmedium field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RankingsResponseRankingsInner) GetAvatarmedium() string {
	if o == nil || IsNil(o.Avatarmedium.Get()) {
		var ret string
		return ret
	}
	return *o.Avatarmedium.Get()
}

// GetAvatarmediumOk returns a tuple with the Avatarmedium field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RankingsResponseRankingsInner) GetAvatarmediumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Avatarmedium.Get(), o.Avatarmedium.IsSet()
}

// HasAvatarmedium returns a boolean if a field has been set.
func (o *RankingsResponseRankingsInner) HasAvatarmedium() bool {
	if o != nil && o.Avatarmedium.IsSet() {
		return true
	}

	return false
}

// SetAvatarmedium gets a reference to the given NullableString and assigns it to the Avatarmedium field.
func (o *RankingsResponseRankingsInner) SetAvatarmedium(v string) {
	o.Avatarmedium.Set(&v)
}
// SetAvatarmediumNil sets the value for Avatarmedium to be an explicit nil
func (o *RankingsResponseRankingsInner) SetAvatarmediumNil() {
	o.Avatarmedium.Set(nil)
}

// UnsetAvatarmedium ensures that no value is present for Avatarmedium, not even an explicit nil
func (o *RankingsResponseRankingsInner) UnsetAvatarmedium() {
	o.Avatarmedium.Unset()
}

// GetAvatarfull returns the Avatarfull field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RankingsResponseRankingsInner) GetAvatarfull() string {
	if o == nil || IsNil(o.Avatarfull.Get()) {
		var ret string
		return ret
	}
	return *o.Avatarfull.Get()
}

// GetAvatarfullOk returns a tuple with the Avatarfull field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RankingsResponseRankingsInner) GetAvatarfullOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Avatarfull.Get(), o.Avatarfull.IsSet()
}

// HasAvatarfull returns a boolean if a field has been set.
func (o *RankingsResponseRankingsInner) HasAvatarfull() bool {
	if o != nil && o.Avatarfull.IsSet() {
		return true
	}

	return false
}

// SetAvatarfull gets a reference to the given NullableString and assigns it to the Avatarfull field.
func (o *RankingsResponseRankingsInner) SetAvatarfull(v string) {
	o.Avatarfull.Set(&v)
}
// SetAvatarfullNil sets the value for Avatarfull to be an explicit nil
func (o *RankingsResponseRankingsInner) SetAvatarfullNil() {
	o.Avatarfull.Set(nil)
}

// UnsetAvatarfull ensures that no value is present for Avatarfull, not even an explicit nil
func (o *RankingsResponseRankingsInner) UnsetAvatarfull() {
	o.Avatarfull.Unset()
}

// GetProfileurl returns the Profileurl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RankingsResponseRankingsInner) GetProfileurl() string {
	if o == nil || IsNil(o.Profileurl.Get()) {
		var ret string
		return ret
	}
	return *o.Profileurl.Get()
}

// GetProfileurlOk returns a tuple with the Profileurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RankingsResponseRankingsInner) GetProfileurlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Profileurl.Get(), o.Profileurl.IsSet()
}

// HasProfileurl returns a boolean if a field has been set.
func (o *RankingsResponseRankingsInner) HasProfileurl() bool {
	if o != nil && o.Profileurl.IsSet() {
		return true
	}

	return false
}

// SetProfileurl gets a reference to the given NullableString and assigns it to the Profileurl field.
func (o *RankingsResponseRankingsInner) SetProfileurl(v string) {
	o.Profileurl.Set(&v)
}
// SetProfileurlNil sets the value for Profileurl to be an explicit nil
func (o *RankingsResponseRankingsInner) SetProfileurlNil() {
	o.Profileurl.Set(nil)
}

// UnsetProfileurl ensures that no value is present for Profileurl, not even an explicit nil
func (o *RankingsResponseRankingsInner) UnsetProfileurl() {
	o.Profileurl.Unset()
}

// GetPersonaname returns the Personaname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RankingsResponseRankingsInner) GetPersonaname() string {
	if o == nil || IsNil(o.Personaname.Get()) {
		var ret string
		return ret
	}
	return *o.Personaname.Get()
}

// GetPersonanameOk returns a tuple with the Personaname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RankingsResponseRankingsInner) GetPersonanameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Personaname.Get(), o.Personaname.IsSet()
}

// HasPersonaname returns a boolean if a field has been set.
func (o *RankingsResponseRankingsInner) HasPersonaname() bool {
	if o != nil && o.Personaname.IsSet() {
		return true
	}

	return false
}

// SetPersonaname gets a reference to the given NullableString and assigns it to the Personaname field.
func (o *RankingsResponseRankingsInner) SetPersonaname(v string) {
	o.Personaname.Set(&v)
}
// SetPersonanameNil sets the value for Personaname to be an explicit nil
func (o *RankingsResponseRankingsInner) SetPersonanameNil() {
	o.Personaname.Set(nil)
}

// UnsetPersonaname ensures that no value is present for Personaname, not even an explicit nil
func (o *RankingsResponseRankingsInner) UnsetPersonaname() {
	o.Personaname.Unset()
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RankingsResponseRankingsInner) GetLastLogin() time.Time {
	if o == nil || IsNil(o.LastLogin.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastLogin.Get()
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RankingsResponseRankingsInner) GetLastLoginOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastLogin.Get(), o.LastLogin.IsSet()
}

// HasLastLogin returns a boolean if a field has been set.
func (o *RankingsResponseRankingsInner) HasLastLogin() bool {
	if o != nil && o.LastLogin.IsSet() {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given NullableTime and assigns it to the LastLogin field.
func (o *RankingsResponseRankingsInner) SetLastLogin(v time.Time) {
	o.LastLogin.Set(&v)
}
// SetLastLoginNil sets the value for LastLogin to be an explicit nil
func (o *RankingsResponseRankingsInner) SetLastLoginNil() {
	o.LastLogin.Set(nil)
}

// UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
func (o *RankingsResponseRankingsInner) UnsetLastLogin() {
	o.LastLogin.Unset()
}

// GetFullHistoryTime returns the FullHistoryTime field value if set, zero value otherwise.
func (o *RankingsResponseRankingsInner) GetFullHistoryTime() time.Time {
	if o == nil || IsNil(o.FullHistoryTime) {
		var ret time.Time
		return ret
	}
	return *o.FullHistoryTime
}

// GetFullHistoryTimeOk returns a tuple with the FullHistoryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RankingsResponseRankingsInner) GetFullHistoryTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FullHistoryTime) {
		return nil, false
	}
	return o.FullHistoryTime, true
}

// HasFullHistoryTime returns a boolean if a field has been set.
func (o *RankingsResponseRankingsInner) HasFullHistoryTime() bool {
	if o != nil && !IsNil(o.FullHistoryTime) {
		return true
	}

	return false
}

// SetFullHistoryTime gets a reference to the given time.Time and assigns it to the FullHistoryTime field.
func (o *RankingsResponseRankingsInner) SetFullHistoryTime(v time.Time) {
	o.FullHistoryTime = &v
}

// GetCheese returns the Cheese field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RankingsResponseRankingsInner) GetCheese() int32 {
	if o == nil || IsNil(o.Cheese.Get()) {
		var ret int32
		return ret
	}
	return *o.Cheese.Get()
}

// GetCheeseOk returns a tuple with the Cheese field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RankingsResponseRankingsInner) GetCheeseOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cheese.Get(), o.Cheese.IsSet()
}

// HasCheese returns a boolean if a field has been set.
func (o *RankingsResponseRankingsInner) HasCheese() bool {
	if o != nil && o.Cheese.IsSet() {
		return true
	}

	return false
}

// SetCheese gets a reference to the given NullableInt32 and assigns it to the Cheese field.
func (o *RankingsResponseRankingsInner) SetCheese(v int32) {
	o.Cheese.Set(&v)
}
// SetCheeseNil sets the value for Cheese to be an explicit nil
func (o *RankingsResponseRankingsInner) SetCheeseNil() {
	o.Cheese.Set(nil)
}

// UnsetCheese ensures that no value is present for Cheese, not even an explicit nil
func (o *RankingsResponseRankingsInner) UnsetCheese() {
	o.Cheese.Unset()
}

// GetFhUnavailable returns the FhUnavailable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RankingsResponseRankingsInner) GetFhUnavailable() bool {
	if o == nil || IsNil(o.FhUnavailable.Get()) {
		var ret bool
		return ret
	}
	return *o.FhUnavailable.Get()
}

// GetFhUnavailableOk returns a tuple with the FhUnavailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RankingsResponseRankingsInner) GetFhUnavailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FhUnavailable.Get(), o.FhUnavailable.IsSet()
}

// HasFhUnavailable returns a boolean if a field has been set.
func (o *RankingsResponseRankingsInner) HasFhUnavailable() bool {
	if o != nil && o.FhUnavailable.IsSet() {
		return true
	}

	return false
}

// SetFhUnavailable gets a reference to the given NullableBool and assigns it to the FhUnavailable field.
func (o *RankingsResponseRankingsInner) SetFhUnavailable(v bool) {
	o.FhUnavailable.Set(&v)
}
// SetFhUnavailableNil sets the value for FhUnavailable to be an explicit nil
func (o *RankingsResponseRankingsInner) SetFhUnavailableNil() {
	o.FhUnavailable.Set(nil)
}

// UnsetFhUnavailable ensures that no value is present for FhUnavailable, not even an explicit nil
func (o *RankingsResponseRankingsInner) UnsetFhUnavailable() {
	o.FhUnavailable.Unset()
}

// GetLoccountrycode returns the Loccountrycode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RankingsResponseRankingsInner) GetLoccountrycode() string {
	if o == nil || IsNil(o.Loccountrycode.Get()) {
		var ret string
		return ret
	}
	return *o.Loccountrycode.Get()
}

// GetLoccountrycodeOk returns a tuple with the Loccountrycode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RankingsResponseRankingsInner) GetLoccountrycodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Loccountrycode.Get(), o.Loccountrycode.IsSet()
}

// HasLoccountrycode returns a boolean if a field has been set.
func (o *RankingsResponseRankingsInner) HasLoccountrycode() bool {
	if o != nil && o.Loccountrycode.IsSet() {
		return true
	}

	return false
}

// SetLoccountrycode gets a reference to the given NullableString and assigns it to the Loccountrycode field.
func (o *RankingsResponseRankingsInner) SetLoccountrycode(v string) {
	o.Loccountrycode.Set(&v)
}
// SetLoccountrycodeNil sets the value for Loccountrycode to be an explicit nil
func (o *RankingsResponseRankingsInner) SetLoccountrycodeNil() {
	o.Loccountrycode.Set(nil)
}

// UnsetLoccountrycode ensures that no value is present for Loccountrycode, not even an explicit nil
func (o *RankingsResponseRankingsInner) UnsetLoccountrycode() {
	o.Loccountrycode.Unset()
}

// GetRankTier returns the RankTier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RankingsResponseRankingsInner) GetRankTier() int32 {
	if o == nil || IsNil(o.RankTier.Get()) {
		var ret int32
		return ret
	}
	return *o.RankTier.Get()
}

// GetRankTierOk returns a tuple with the RankTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RankingsResponseRankingsInner) GetRankTierOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RankTier.Get(), o.RankTier.IsSet()
}

// HasRankTier returns a boolean if a field has been set.
func (o *RankingsResponseRankingsInner) HasRankTier() bool {
	if o != nil && o.RankTier.IsSet() {
		return true
	}

	return false
}

// SetRankTier gets a reference to the given NullableInt32 and assigns it to the RankTier field.
func (o *RankingsResponseRankingsInner) SetRankTier(v int32) {
	o.RankTier.Set(&v)
}
// SetRankTierNil sets the value for RankTier to be an explicit nil
func (o *RankingsResponseRankingsInner) SetRankTierNil() {
	o.RankTier.Set(nil)
}

// UnsetRankTier ensures that no value is present for RankTier, not even an explicit nil
func (o *RankingsResponseRankingsInner) UnsetRankTier() {
	o.RankTier.Unset()
}

func (o RankingsResponseRankingsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RankingsResponseRankingsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if o.Steamid.IsSet() {
		toSerialize["steamid"] = o.Steamid.Get()
	}
	if o.Avatar.IsSet() {
		toSerialize["avatar"] = o.Avatar.Get()
	}
	if o.Avatarmedium.IsSet() {
		toSerialize["avatarmedium"] = o.Avatarmedium.Get()
	}
	if o.Avatarfull.IsSet() {
		toSerialize["avatarfull"] = o.Avatarfull.Get()
	}
	if o.Profileurl.IsSet() {
		toSerialize["profileurl"] = o.Profileurl.Get()
	}
	if o.Personaname.IsSet() {
		toSerialize["personaname"] = o.Personaname.Get()
	}
	if o.LastLogin.IsSet() {
		toSerialize["last_login"] = o.LastLogin.Get()
	}
	if !IsNil(o.FullHistoryTime) {
		toSerialize["full_history_time"] = o.FullHistoryTime
	}
	if o.Cheese.IsSet() {
		toSerialize["cheese"] = o.Cheese.Get()
	}
	if o.FhUnavailable.IsSet() {
		toSerialize["fh_unavailable"] = o.FhUnavailable.Get()
	}
	if o.Loccountrycode.IsSet() {
		toSerialize["loccountrycode"] = o.Loccountrycode.Get()
	}
	if o.RankTier.IsSet() {
		toSerialize["rank_tier"] = o.RankTier.Get()
	}
	return toSerialize, nil
}

type NullableRankingsResponseRankingsInner struct {
	value *RankingsResponseRankingsInner
	isSet bool
}

func (v NullableRankingsResponseRankingsInner) Get() *RankingsResponseRankingsInner {
	return v.value
}

func (v *NullableRankingsResponseRankingsInner) Set(val *RankingsResponseRankingsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRankingsResponseRankingsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRankingsResponseRankingsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRankingsResponseRankingsInner(val *RankingsResponseRankingsInner) *NullableRankingsResponseRankingsInner {
	return &NullableRankingsResponseRankingsInner{value: val, isSet: true}
}

func (v NullableRankingsResponseRankingsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRankingsResponseRankingsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


