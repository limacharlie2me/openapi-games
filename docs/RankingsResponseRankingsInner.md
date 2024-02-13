# RankingsResponseRankingsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** | The player account ID | [optional] 
**Score** | Pointer to **float32** | Score | [optional] 
**Steamid** | Pointer to **NullableString** | steamid | [optional] 
**Avatar** | Pointer to **NullableString** | avatar | [optional] 
**Avatarmedium** | Pointer to **NullableString** | avatarmedium | [optional] 
**Avatarfull** | Pointer to **NullableString** | avatarfull | [optional] 
**Profileurl** | Pointer to **NullableString** | profileurl | [optional] 
**Personaname** | Pointer to **NullableString** | Player&#39;s Steam name | [optional] 
**LastLogin** | Pointer to **NullableTime** | last_login | [optional] 
**FullHistoryTime** | Pointer to **time.Time** | full_history_time | [optional] 
**Cheese** | Pointer to **NullableInt32** | cheese | [optional] 
**FhUnavailable** | Pointer to **NullableBool** | fh_unavailable | [optional] 
**Loccountrycode** | Pointer to **NullableString** | loccountrycode | [optional] 
**RankTier** | Pointer to **NullableInt32** | rank_tier | [optional] 

## Methods

### NewRankingsResponseRankingsInner

`func NewRankingsResponseRankingsInner() *RankingsResponseRankingsInner`

NewRankingsResponseRankingsInner instantiates a new RankingsResponseRankingsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRankingsResponseRankingsInnerWithDefaults

`func NewRankingsResponseRankingsInnerWithDefaults() *RankingsResponseRankingsInner`

NewRankingsResponseRankingsInnerWithDefaults instantiates a new RankingsResponseRankingsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *RankingsResponseRankingsInner) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RankingsResponseRankingsInner) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RankingsResponseRankingsInner) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *RankingsResponseRankingsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetScore

`func (o *RankingsResponseRankingsInner) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *RankingsResponseRankingsInner) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *RankingsResponseRankingsInner) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *RankingsResponseRankingsInner) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSteamid

`func (o *RankingsResponseRankingsInner) GetSteamid() string`

GetSteamid returns the Steamid field if non-nil, zero value otherwise.

### GetSteamidOk

`func (o *RankingsResponseRankingsInner) GetSteamidOk() (*string, bool)`

GetSteamidOk returns a tuple with the Steamid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamid

`func (o *RankingsResponseRankingsInner) SetSteamid(v string)`

SetSteamid sets Steamid field to given value.

### HasSteamid

`func (o *RankingsResponseRankingsInner) HasSteamid() bool`

HasSteamid returns a boolean if a field has been set.

### SetSteamidNil

`func (o *RankingsResponseRankingsInner) SetSteamidNil(b bool)`

 SetSteamidNil sets the value for Steamid to be an explicit nil

### UnsetSteamid
`func (o *RankingsResponseRankingsInner) UnsetSteamid()`

UnsetSteamid ensures that no value is present for Steamid, not even an explicit nil
### GetAvatar

`func (o *RankingsResponseRankingsInner) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *RankingsResponseRankingsInner) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *RankingsResponseRankingsInner) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *RankingsResponseRankingsInner) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *RankingsResponseRankingsInner) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *RankingsResponseRankingsInner) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetAvatarmedium

`func (o *RankingsResponseRankingsInner) GetAvatarmedium() string`

GetAvatarmedium returns the Avatarmedium field if non-nil, zero value otherwise.

### GetAvatarmediumOk

`func (o *RankingsResponseRankingsInner) GetAvatarmediumOk() (*string, bool)`

GetAvatarmediumOk returns a tuple with the Avatarmedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarmedium

`func (o *RankingsResponseRankingsInner) SetAvatarmedium(v string)`

SetAvatarmedium sets Avatarmedium field to given value.

### HasAvatarmedium

`func (o *RankingsResponseRankingsInner) HasAvatarmedium() bool`

HasAvatarmedium returns a boolean if a field has been set.

### SetAvatarmediumNil

`func (o *RankingsResponseRankingsInner) SetAvatarmediumNil(b bool)`

 SetAvatarmediumNil sets the value for Avatarmedium to be an explicit nil

### UnsetAvatarmedium
`func (o *RankingsResponseRankingsInner) UnsetAvatarmedium()`

UnsetAvatarmedium ensures that no value is present for Avatarmedium, not even an explicit nil
### GetAvatarfull

`func (o *RankingsResponseRankingsInner) GetAvatarfull() string`

GetAvatarfull returns the Avatarfull field if non-nil, zero value otherwise.

### GetAvatarfullOk

`func (o *RankingsResponseRankingsInner) GetAvatarfullOk() (*string, bool)`

GetAvatarfullOk returns a tuple with the Avatarfull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarfull

`func (o *RankingsResponseRankingsInner) SetAvatarfull(v string)`

SetAvatarfull sets Avatarfull field to given value.

### HasAvatarfull

`func (o *RankingsResponseRankingsInner) HasAvatarfull() bool`

HasAvatarfull returns a boolean if a field has been set.

### SetAvatarfullNil

`func (o *RankingsResponseRankingsInner) SetAvatarfullNil(b bool)`

 SetAvatarfullNil sets the value for Avatarfull to be an explicit nil

### UnsetAvatarfull
`func (o *RankingsResponseRankingsInner) UnsetAvatarfull()`

UnsetAvatarfull ensures that no value is present for Avatarfull, not even an explicit nil
### GetProfileurl

`func (o *RankingsResponseRankingsInner) GetProfileurl() string`

GetProfileurl returns the Profileurl field if non-nil, zero value otherwise.

### GetProfileurlOk

`func (o *RankingsResponseRankingsInner) GetProfileurlOk() (*string, bool)`

GetProfileurlOk returns a tuple with the Profileurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileurl

`func (o *RankingsResponseRankingsInner) SetProfileurl(v string)`

SetProfileurl sets Profileurl field to given value.

### HasProfileurl

`func (o *RankingsResponseRankingsInner) HasProfileurl() bool`

HasProfileurl returns a boolean if a field has been set.

### SetProfileurlNil

`func (o *RankingsResponseRankingsInner) SetProfileurlNil(b bool)`

 SetProfileurlNil sets the value for Profileurl to be an explicit nil

### UnsetProfileurl
`func (o *RankingsResponseRankingsInner) UnsetProfileurl()`

UnsetProfileurl ensures that no value is present for Profileurl, not even an explicit nil
### GetPersonaname

`func (o *RankingsResponseRankingsInner) GetPersonaname() string`

GetPersonaname returns the Personaname field if non-nil, zero value otherwise.

### GetPersonanameOk

`func (o *RankingsResponseRankingsInner) GetPersonanameOk() (*string, bool)`

GetPersonanameOk returns a tuple with the Personaname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonaname

`func (o *RankingsResponseRankingsInner) SetPersonaname(v string)`

SetPersonaname sets Personaname field to given value.

### HasPersonaname

`func (o *RankingsResponseRankingsInner) HasPersonaname() bool`

HasPersonaname returns a boolean if a field has been set.

### SetPersonanameNil

`func (o *RankingsResponseRankingsInner) SetPersonanameNil(b bool)`

 SetPersonanameNil sets the value for Personaname to be an explicit nil

### UnsetPersonaname
`func (o *RankingsResponseRankingsInner) UnsetPersonaname()`

UnsetPersonaname ensures that no value is present for Personaname, not even an explicit nil
### GetLastLogin

`func (o *RankingsResponseRankingsInner) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *RankingsResponseRankingsInner) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *RankingsResponseRankingsInner) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *RankingsResponseRankingsInner) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *RankingsResponseRankingsInner) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *RankingsResponseRankingsInner) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetFullHistoryTime

`func (o *RankingsResponseRankingsInner) GetFullHistoryTime() time.Time`

GetFullHistoryTime returns the FullHistoryTime field if non-nil, zero value otherwise.

### GetFullHistoryTimeOk

`func (o *RankingsResponseRankingsInner) GetFullHistoryTimeOk() (*time.Time, bool)`

GetFullHistoryTimeOk returns a tuple with the FullHistoryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullHistoryTime

`func (o *RankingsResponseRankingsInner) SetFullHistoryTime(v time.Time)`

SetFullHistoryTime sets FullHistoryTime field to given value.

### HasFullHistoryTime

`func (o *RankingsResponseRankingsInner) HasFullHistoryTime() bool`

HasFullHistoryTime returns a boolean if a field has been set.

### GetCheese

`func (o *RankingsResponseRankingsInner) GetCheese() int32`

GetCheese returns the Cheese field if non-nil, zero value otherwise.

### GetCheeseOk

`func (o *RankingsResponseRankingsInner) GetCheeseOk() (*int32, bool)`

GetCheeseOk returns a tuple with the Cheese field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheese

`func (o *RankingsResponseRankingsInner) SetCheese(v int32)`

SetCheese sets Cheese field to given value.

### HasCheese

`func (o *RankingsResponseRankingsInner) HasCheese() bool`

HasCheese returns a boolean if a field has been set.

### SetCheeseNil

`func (o *RankingsResponseRankingsInner) SetCheeseNil(b bool)`

 SetCheeseNil sets the value for Cheese to be an explicit nil

### UnsetCheese
`func (o *RankingsResponseRankingsInner) UnsetCheese()`

UnsetCheese ensures that no value is present for Cheese, not even an explicit nil
### GetFhUnavailable

`func (o *RankingsResponseRankingsInner) GetFhUnavailable() bool`

GetFhUnavailable returns the FhUnavailable field if non-nil, zero value otherwise.

### GetFhUnavailableOk

`func (o *RankingsResponseRankingsInner) GetFhUnavailableOk() (*bool, bool)`

GetFhUnavailableOk returns a tuple with the FhUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFhUnavailable

`func (o *RankingsResponseRankingsInner) SetFhUnavailable(v bool)`

SetFhUnavailable sets FhUnavailable field to given value.

### HasFhUnavailable

`func (o *RankingsResponseRankingsInner) HasFhUnavailable() bool`

HasFhUnavailable returns a boolean if a field has been set.

### SetFhUnavailableNil

`func (o *RankingsResponseRankingsInner) SetFhUnavailableNil(b bool)`

 SetFhUnavailableNil sets the value for FhUnavailable to be an explicit nil

### UnsetFhUnavailable
`func (o *RankingsResponseRankingsInner) UnsetFhUnavailable()`

UnsetFhUnavailable ensures that no value is present for FhUnavailable, not even an explicit nil
### GetLoccountrycode

`func (o *RankingsResponseRankingsInner) GetLoccountrycode() string`

GetLoccountrycode returns the Loccountrycode field if non-nil, zero value otherwise.

### GetLoccountrycodeOk

`func (o *RankingsResponseRankingsInner) GetLoccountrycodeOk() (*string, bool)`

GetLoccountrycodeOk returns a tuple with the Loccountrycode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoccountrycode

`func (o *RankingsResponseRankingsInner) SetLoccountrycode(v string)`

SetLoccountrycode sets Loccountrycode field to given value.

### HasLoccountrycode

`func (o *RankingsResponseRankingsInner) HasLoccountrycode() bool`

HasLoccountrycode returns a boolean if a field has been set.

### SetLoccountrycodeNil

`func (o *RankingsResponseRankingsInner) SetLoccountrycodeNil(b bool)`

 SetLoccountrycodeNil sets the value for Loccountrycode to be an explicit nil

### UnsetLoccountrycode
`func (o *RankingsResponseRankingsInner) UnsetLoccountrycode()`

UnsetLoccountrycode ensures that no value is present for Loccountrycode, not even an explicit nil
### GetRankTier

`func (o *RankingsResponseRankingsInner) GetRankTier() int32`

GetRankTier returns the RankTier field if non-nil, zero value otherwise.

### GetRankTierOk

`func (o *RankingsResponseRankingsInner) GetRankTierOk() (*int32, bool)`

GetRankTierOk returns a tuple with the RankTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankTier

`func (o *RankingsResponseRankingsInner) SetRankTier(v int32)`

SetRankTier sets RankTier field to given value.

### HasRankTier

`func (o *RankingsResponseRankingsInner) HasRankTier() bool`

HasRankTier returns a boolean if a field has been set.

### SetRankTierNil

`func (o *RankingsResponseRankingsInner) SetRankTierNil(b bool)`

 SetRankTierNil sets the value for RankTier to be an explicit nil

### UnsetRankTier
`func (o *RankingsResponseRankingsInner) UnsetRankTier()`

UnsetRankTier ensures that no value is present for RankTier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


