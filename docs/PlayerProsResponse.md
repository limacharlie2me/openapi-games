# PlayerProsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** | The player account ID | [optional] 
**Name** | Pointer to **NullableString** | name | [optional] 
**CountryCode** | Pointer to **string** | country_code | [optional] 
**FantasyRole** | Pointer to **int32** | fantasy_role | [optional] 
**TeamId** | Pointer to **int32** | team_id | [optional] 
**TeamName** | Pointer to **NullableString** | Team name | [optional] 
**TeamTag** | Pointer to **NullableString** | team_tag | [optional] 
**IsLocked** | Pointer to **bool** | is_locked | [optional] 
**IsPro** | Pointer to **bool** | is_pro | [optional] 
**LockedUntil** | Pointer to **NullableInt32** | locked_until | [optional] 
**Steamid** | Pointer to **NullableString** | steamid | [optional] 
**Avatar** | Pointer to **NullableString** | avatar | [optional] 
**Avatarmedium** | Pointer to **NullableString** | avatarmedium | [optional] 
**Avatarfull** | Pointer to **NullableString** | avatarfull | [optional] 
**Profileurl** | Pointer to **NullableString** | profileurl | [optional] 
**LastLogin** | Pointer to **NullableTime** | last_login | [optional] 
**FullHistoryTime** | Pointer to **NullableTime** | full_history_time | [optional] 
**Cheese** | Pointer to **NullableInt32** | cheese | [optional] 
**FhUnavailable** | Pointer to **NullableBool** | fh_unavailable | [optional] 
**Loccountrycode** | Pointer to **NullableString** | loccountrycode | [optional] 
**LastPlayed** | Pointer to **NullableInt32** | last_played | [optional] 
**Win** | Pointer to **int32** | win | [optional] 
**Games** | Pointer to **int32** | games | [optional] 
**WithWin** | Pointer to **int32** | with_win | [optional] 
**WithGames** | Pointer to **int32** | with_games | [optional] 
**AgainstWin** | Pointer to **int32** | against_win | [optional] 
**AgainstGames** | Pointer to **int32** | against_games | [optional] 
**WithGpmSum** | Pointer to **NullableInt32** | with_gpm_sum | [optional] 
**WithXpmSum** | Pointer to **NullableInt32** | with_xpm_sum | [optional] 

## Methods

### NewPlayerProsResponse

`func NewPlayerProsResponse() *PlayerProsResponse`

NewPlayerProsResponse instantiates a new PlayerProsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerProsResponseWithDefaults

`func NewPlayerProsResponseWithDefaults() *PlayerProsResponse`

NewPlayerProsResponseWithDefaults instantiates a new PlayerProsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *PlayerProsResponse) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PlayerProsResponse) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PlayerProsResponse) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *PlayerProsResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *PlayerProsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlayerProsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlayerProsResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlayerProsResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PlayerProsResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PlayerProsResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCountryCode

`func (o *PlayerProsResponse) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PlayerProsResponse) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PlayerProsResponse) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PlayerProsResponse) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetFantasyRole

`func (o *PlayerProsResponse) GetFantasyRole() int32`

GetFantasyRole returns the FantasyRole field if non-nil, zero value otherwise.

### GetFantasyRoleOk

`func (o *PlayerProsResponse) GetFantasyRoleOk() (*int32, bool)`

GetFantasyRoleOk returns a tuple with the FantasyRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFantasyRole

`func (o *PlayerProsResponse) SetFantasyRole(v int32)`

SetFantasyRole sets FantasyRole field to given value.

### HasFantasyRole

`func (o *PlayerProsResponse) HasFantasyRole() bool`

HasFantasyRole returns a boolean if a field has been set.

### GetTeamId

`func (o *PlayerProsResponse) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *PlayerProsResponse) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *PlayerProsResponse) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *PlayerProsResponse) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetTeamName

`func (o *PlayerProsResponse) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *PlayerProsResponse) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *PlayerProsResponse) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *PlayerProsResponse) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### SetTeamNameNil

`func (o *PlayerProsResponse) SetTeamNameNil(b bool)`

 SetTeamNameNil sets the value for TeamName to be an explicit nil

### UnsetTeamName
`func (o *PlayerProsResponse) UnsetTeamName()`

UnsetTeamName ensures that no value is present for TeamName, not even an explicit nil
### GetTeamTag

`func (o *PlayerProsResponse) GetTeamTag() string`

GetTeamTag returns the TeamTag field if non-nil, zero value otherwise.

### GetTeamTagOk

`func (o *PlayerProsResponse) GetTeamTagOk() (*string, bool)`

GetTeamTagOk returns a tuple with the TeamTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamTag

`func (o *PlayerProsResponse) SetTeamTag(v string)`

SetTeamTag sets TeamTag field to given value.

### HasTeamTag

`func (o *PlayerProsResponse) HasTeamTag() bool`

HasTeamTag returns a boolean if a field has been set.

### SetTeamTagNil

`func (o *PlayerProsResponse) SetTeamTagNil(b bool)`

 SetTeamTagNil sets the value for TeamTag to be an explicit nil

### UnsetTeamTag
`func (o *PlayerProsResponse) UnsetTeamTag()`

UnsetTeamTag ensures that no value is present for TeamTag, not even an explicit nil
### GetIsLocked

`func (o *PlayerProsResponse) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *PlayerProsResponse) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *PlayerProsResponse) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *PlayerProsResponse) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsPro

`func (o *PlayerProsResponse) GetIsPro() bool`

GetIsPro returns the IsPro field if non-nil, zero value otherwise.

### GetIsProOk

`func (o *PlayerProsResponse) GetIsProOk() (*bool, bool)`

GetIsProOk returns a tuple with the IsPro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPro

`func (o *PlayerProsResponse) SetIsPro(v bool)`

SetIsPro sets IsPro field to given value.

### HasIsPro

`func (o *PlayerProsResponse) HasIsPro() bool`

HasIsPro returns a boolean if a field has been set.

### GetLockedUntil

`func (o *PlayerProsResponse) GetLockedUntil() int32`

GetLockedUntil returns the LockedUntil field if non-nil, zero value otherwise.

### GetLockedUntilOk

`func (o *PlayerProsResponse) GetLockedUntilOk() (*int32, bool)`

GetLockedUntilOk returns a tuple with the LockedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedUntil

`func (o *PlayerProsResponse) SetLockedUntil(v int32)`

SetLockedUntil sets LockedUntil field to given value.

### HasLockedUntil

`func (o *PlayerProsResponse) HasLockedUntil() bool`

HasLockedUntil returns a boolean if a field has been set.

### SetLockedUntilNil

`func (o *PlayerProsResponse) SetLockedUntilNil(b bool)`

 SetLockedUntilNil sets the value for LockedUntil to be an explicit nil

### UnsetLockedUntil
`func (o *PlayerProsResponse) UnsetLockedUntil()`

UnsetLockedUntil ensures that no value is present for LockedUntil, not even an explicit nil
### GetSteamid

`func (o *PlayerProsResponse) GetSteamid() string`

GetSteamid returns the Steamid field if non-nil, zero value otherwise.

### GetSteamidOk

`func (o *PlayerProsResponse) GetSteamidOk() (*string, bool)`

GetSteamidOk returns a tuple with the Steamid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamid

`func (o *PlayerProsResponse) SetSteamid(v string)`

SetSteamid sets Steamid field to given value.

### HasSteamid

`func (o *PlayerProsResponse) HasSteamid() bool`

HasSteamid returns a boolean if a field has been set.

### SetSteamidNil

`func (o *PlayerProsResponse) SetSteamidNil(b bool)`

 SetSteamidNil sets the value for Steamid to be an explicit nil

### UnsetSteamid
`func (o *PlayerProsResponse) UnsetSteamid()`

UnsetSteamid ensures that no value is present for Steamid, not even an explicit nil
### GetAvatar

`func (o *PlayerProsResponse) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *PlayerProsResponse) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *PlayerProsResponse) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *PlayerProsResponse) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *PlayerProsResponse) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *PlayerProsResponse) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetAvatarmedium

`func (o *PlayerProsResponse) GetAvatarmedium() string`

GetAvatarmedium returns the Avatarmedium field if non-nil, zero value otherwise.

### GetAvatarmediumOk

`func (o *PlayerProsResponse) GetAvatarmediumOk() (*string, bool)`

GetAvatarmediumOk returns a tuple with the Avatarmedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarmedium

`func (o *PlayerProsResponse) SetAvatarmedium(v string)`

SetAvatarmedium sets Avatarmedium field to given value.

### HasAvatarmedium

`func (o *PlayerProsResponse) HasAvatarmedium() bool`

HasAvatarmedium returns a boolean if a field has been set.

### SetAvatarmediumNil

`func (o *PlayerProsResponse) SetAvatarmediumNil(b bool)`

 SetAvatarmediumNil sets the value for Avatarmedium to be an explicit nil

### UnsetAvatarmedium
`func (o *PlayerProsResponse) UnsetAvatarmedium()`

UnsetAvatarmedium ensures that no value is present for Avatarmedium, not even an explicit nil
### GetAvatarfull

`func (o *PlayerProsResponse) GetAvatarfull() string`

GetAvatarfull returns the Avatarfull field if non-nil, zero value otherwise.

### GetAvatarfullOk

`func (o *PlayerProsResponse) GetAvatarfullOk() (*string, bool)`

GetAvatarfullOk returns a tuple with the Avatarfull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarfull

`func (o *PlayerProsResponse) SetAvatarfull(v string)`

SetAvatarfull sets Avatarfull field to given value.

### HasAvatarfull

`func (o *PlayerProsResponse) HasAvatarfull() bool`

HasAvatarfull returns a boolean if a field has been set.

### SetAvatarfullNil

`func (o *PlayerProsResponse) SetAvatarfullNil(b bool)`

 SetAvatarfullNil sets the value for Avatarfull to be an explicit nil

### UnsetAvatarfull
`func (o *PlayerProsResponse) UnsetAvatarfull()`

UnsetAvatarfull ensures that no value is present for Avatarfull, not even an explicit nil
### GetProfileurl

`func (o *PlayerProsResponse) GetProfileurl() string`

GetProfileurl returns the Profileurl field if non-nil, zero value otherwise.

### GetProfileurlOk

`func (o *PlayerProsResponse) GetProfileurlOk() (*string, bool)`

GetProfileurlOk returns a tuple with the Profileurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileurl

`func (o *PlayerProsResponse) SetProfileurl(v string)`

SetProfileurl sets Profileurl field to given value.

### HasProfileurl

`func (o *PlayerProsResponse) HasProfileurl() bool`

HasProfileurl returns a boolean if a field has been set.

### SetProfileurlNil

`func (o *PlayerProsResponse) SetProfileurlNil(b bool)`

 SetProfileurlNil sets the value for Profileurl to be an explicit nil

### UnsetProfileurl
`func (o *PlayerProsResponse) UnsetProfileurl()`

UnsetProfileurl ensures that no value is present for Profileurl, not even an explicit nil
### GetLastLogin

`func (o *PlayerProsResponse) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *PlayerProsResponse) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *PlayerProsResponse) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *PlayerProsResponse) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *PlayerProsResponse) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *PlayerProsResponse) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetFullHistoryTime

`func (o *PlayerProsResponse) GetFullHistoryTime() time.Time`

GetFullHistoryTime returns the FullHistoryTime field if non-nil, zero value otherwise.

### GetFullHistoryTimeOk

`func (o *PlayerProsResponse) GetFullHistoryTimeOk() (*time.Time, bool)`

GetFullHistoryTimeOk returns a tuple with the FullHistoryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullHistoryTime

`func (o *PlayerProsResponse) SetFullHistoryTime(v time.Time)`

SetFullHistoryTime sets FullHistoryTime field to given value.

### HasFullHistoryTime

`func (o *PlayerProsResponse) HasFullHistoryTime() bool`

HasFullHistoryTime returns a boolean if a field has been set.

### SetFullHistoryTimeNil

`func (o *PlayerProsResponse) SetFullHistoryTimeNil(b bool)`

 SetFullHistoryTimeNil sets the value for FullHistoryTime to be an explicit nil

### UnsetFullHistoryTime
`func (o *PlayerProsResponse) UnsetFullHistoryTime()`

UnsetFullHistoryTime ensures that no value is present for FullHistoryTime, not even an explicit nil
### GetCheese

`func (o *PlayerProsResponse) GetCheese() int32`

GetCheese returns the Cheese field if non-nil, zero value otherwise.

### GetCheeseOk

`func (o *PlayerProsResponse) GetCheeseOk() (*int32, bool)`

GetCheeseOk returns a tuple with the Cheese field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheese

`func (o *PlayerProsResponse) SetCheese(v int32)`

SetCheese sets Cheese field to given value.

### HasCheese

`func (o *PlayerProsResponse) HasCheese() bool`

HasCheese returns a boolean if a field has been set.

### SetCheeseNil

`func (o *PlayerProsResponse) SetCheeseNil(b bool)`

 SetCheeseNil sets the value for Cheese to be an explicit nil

### UnsetCheese
`func (o *PlayerProsResponse) UnsetCheese()`

UnsetCheese ensures that no value is present for Cheese, not even an explicit nil
### GetFhUnavailable

`func (o *PlayerProsResponse) GetFhUnavailable() bool`

GetFhUnavailable returns the FhUnavailable field if non-nil, zero value otherwise.

### GetFhUnavailableOk

`func (o *PlayerProsResponse) GetFhUnavailableOk() (*bool, bool)`

GetFhUnavailableOk returns a tuple with the FhUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFhUnavailable

`func (o *PlayerProsResponse) SetFhUnavailable(v bool)`

SetFhUnavailable sets FhUnavailable field to given value.

### HasFhUnavailable

`func (o *PlayerProsResponse) HasFhUnavailable() bool`

HasFhUnavailable returns a boolean if a field has been set.

### SetFhUnavailableNil

`func (o *PlayerProsResponse) SetFhUnavailableNil(b bool)`

 SetFhUnavailableNil sets the value for FhUnavailable to be an explicit nil

### UnsetFhUnavailable
`func (o *PlayerProsResponse) UnsetFhUnavailable()`

UnsetFhUnavailable ensures that no value is present for FhUnavailable, not even an explicit nil
### GetLoccountrycode

`func (o *PlayerProsResponse) GetLoccountrycode() string`

GetLoccountrycode returns the Loccountrycode field if non-nil, zero value otherwise.

### GetLoccountrycodeOk

`func (o *PlayerProsResponse) GetLoccountrycodeOk() (*string, bool)`

GetLoccountrycodeOk returns a tuple with the Loccountrycode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoccountrycode

`func (o *PlayerProsResponse) SetLoccountrycode(v string)`

SetLoccountrycode sets Loccountrycode field to given value.

### HasLoccountrycode

`func (o *PlayerProsResponse) HasLoccountrycode() bool`

HasLoccountrycode returns a boolean if a field has been set.

### SetLoccountrycodeNil

`func (o *PlayerProsResponse) SetLoccountrycodeNil(b bool)`

 SetLoccountrycodeNil sets the value for Loccountrycode to be an explicit nil

### UnsetLoccountrycode
`func (o *PlayerProsResponse) UnsetLoccountrycode()`

UnsetLoccountrycode ensures that no value is present for Loccountrycode, not even an explicit nil
### GetLastPlayed

`func (o *PlayerProsResponse) GetLastPlayed() int32`

GetLastPlayed returns the LastPlayed field if non-nil, zero value otherwise.

### GetLastPlayedOk

`func (o *PlayerProsResponse) GetLastPlayedOk() (*int32, bool)`

GetLastPlayedOk returns a tuple with the LastPlayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPlayed

`func (o *PlayerProsResponse) SetLastPlayed(v int32)`

SetLastPlayed sets LastPlayed field to given value.

### HasLastPlayed

`func (o *PlayerProsResponse) HasLastPlayed() bool`

HasLastPlayed returns a boolean if a field has been set.

### SetLastPlayedNil

`func (o *PlayerProsResponse) SetLastPlayedNil(b bool)`

 SetLastPlayedNil sets the value for LastPlayed to be an explicit nil

### UnsetLastPlayed
`func (o *PlayerProsResponse) UnsetLastPlayed()`

UnsetLastPlayed ensures that no value is present for LastPlayed, not even an explicit nil
### GetWin

`func (o *PlayerProsResponse) GetWin() int32`

GetWin returns the Win field if non-nil, zero value otherwise.

### GetWinOk

`func (o *PlayerProsResponse) GetWinOk() (*int32, bool)`

GetWinOk returns a tuple with the Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWin

`func (o *PlayerProsResponse) SetWin(v int32)`

SetWin sets Win field to given value.

### HasWin

`func (o *PlayerProsResponse) HasWin() bool`

HasWin returns a boolean if a field has been set.

### GetGames

`func (o *PlayerProsResponse) GetGames() int32`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *PlayerProsResponse) GetGamesOk() (*int32, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *PlayerProsResponse) SetGames(v int32)`

SetGames sets Games field to given value.

### HasGames

`func (o *PlayerProsResponse) HasGames() bool`

HasGames returns a boolean if a field has been set.

### GetWithWin

`func (o *PlayerProsResponse) GetWithWin() int32`

GetWithWin returns the WithWin field if non-nil, zero value otherwise.

### GetWithWinOk

`func (o *PlayerProsResponse) GetWithWinOk() (*int32, bool)`

GetWithWinOk returns a tuple with the WithWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithWin

`func (o *PlayerProsResponse) SetWithWin(v int32)`

SetWithWin sets WithWin field to given value.

### HasWithWin

`func (o *PlayerProsResponse) HasWithWin() bool`

HasWithWin returns a boolean if a field has been set.

### GetWithGames

`func (o *PlayerProsResponse) GetWithGames() int32`

GetWithGames returns the WithGames field if non-nil, zero value otherwise.

### GetWithGamesOk

`func (o *PlayerProsResponse) GetWithGamesOk() (*int32, bool)`

GetWithGamesOk returns a tuple with the WithGames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithGames

`func (o *PlayerProsResponse) SetWithGames(v int32)`

SetWithGames sets WithGames field to given value.

### HasWithGames

`func (o *PlayerProsResponse) HasWithGames() bool`

HasWithGames returns a boolean if a field has been set.

### GetAgainstWin

`func (o *PlayerProsResponse) GetAgainstWin() int32`

GetAgainstWin returns the AgainstWin field if non-nil, zero value otherwise.

### GetAgainstWinOk

`func (o *PlayerProsResponse) GetAgainstWinOk() (*int32, bool)`

GetAgainstWinOk returns a tuple with the AgainstWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgainstWin

`func (o *PlayerProsResponse) SetAgainstWin(v int32)`

SetAgainstWin sets AgainstWin field to given value.

### HasAgainstWin

`func (o *PlayerProsResponse) HasAgainstWin() bool`

HasAgainstWin returns a boolean if a field has been set.

### GetAgainstGames

`func (o *PlayerProsResponse) GetAgainstGames() int32`

GetAgainstGames returns the AgainstGames field if non-nil, zero value otherwise.

### GetAgainstGamesOk

`func (o *PlayerProsResponse) GetAgainstGamesOk() (*int32, bool)`

GetAgainstGamesOk returns a tuple with the AgainstGames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgainstGames

`func (o *PlayerProsResponse) SetAgainstGames(v int32)`

SetAgainstGames sets AgainstGames field to given value.

### HasAgainstGames

`func (o *PlayerProsResponse) HasAgainstGames() bool`

HasAgainstGames returns a boolean if a field has been set.

### GetWithGpmSum

`func (o *PlayerProsResponse) GetWithGpmSum() int32`

GetWithGpmSum returns the WithGpmSum field if non-nil, zero value otherwise.

### GetWithGpmSumOk

`func (o *PlayerProsResponse) GetWithGpmSumOk() (*int32, bool)`

GetWithGpmSumOk returns a tuple with the WithGpmSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithGpmSum

`func (o *PlayerProsResponse) SetWithGpmSum(v int32)`

SetWithGpmSum sets WithGpmSum field to given value.

### HasWithGpmSum

`func (o *PlayerProsResponse) HasWithGpmSum() bool`

HasWithGpmSum returns a boolean if a field has been set.

### SetWithGpmSumNil

`func (o *PlayerProsResponse) SetWithGpmSumNil(b bool)`

 SetWithGpmSumNil sets the value for WithGpmSum to be an explicit nil

### UnsetWithGpmSum
`func (o *PlayerProsResponse) UnsetWithGpmSum()`

UnsetWithGpmSum ensures that no value is present for WithGpmSum, not even an explicit nil
### GetWithXpmSum

`func (o *PlayerProsResponse) GetWithXpmSum() int32`

GetWithXpmSum returns the WithXpmSum field if non-nil, zero value otherwise.

### GetWithXpmSumOk

`func (o *PlayerProsResponse) GetWithXpmSumOk() (*int32, bool)`

GetWithXpmSumOk returns a tuple with the WithXpmSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithXpmSum

`func (o *PlayerProsResponse) SetWithXpmSum(v int32)`

SetWithXpmSum sets WithXpmSum field to given value.

### HasWithXpmSum

`func (o *PlayerProsResponse) HasWithXpmSum() bool`

HasWithXpmSum returns a boolean if a field has been set.

### SetWithXpmSumNil

`func (o *PlayerProsResponse) SetWithXpmSumNil(b bool)`

 SetWithXpmSumNil sets the value for WithXpmSum to be an explicit nil

### UnsetWithXpmSum
`func (o *PlayerProsResponse) UnsetWithXpmSum()`

UnsetWithXpmSum ensures that no value is present for WithXpmSum, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


