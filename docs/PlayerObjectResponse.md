# PlayerObjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** | The player account ID | [optional] 
**Steamid** | Pointer to **string** | Player&#39;s steam identifier | [optional] 
**Avatar** | Pointer to **string** | Steam picture URL (small picture) | [optional] 
**Avatarmedium** | Pointer to **string** | Steam picture URL (medium picture) | [optional] 
**Avatarfull** | Pointer to **string** | Steam picture URL (full picture) | [optional] 
**Profileurl** | Pointer to **string** | Steam profile URL | [optional] 
**Personaname** | Pointer to **NullableString** | Player&#39;s Steam name | [optional] 
**LastLogin** | Pointer to **time.Time** | Date and time of last login to OpenDota | [optional] 
**FullHistoryTime** | Pointer to **time.Time** | Date and time of last request to refresh player&#39;s match history | [optional] 
**Cheese** | Pointer to **int32** | Amount of dollars the player has donated to OpenDota | [optional] 
**FhUnavailable** | Pointer to **bool** | Whether the refresh of player&#39; match history failed | [optional] 
**Loccountrycode** | Pointer to **string** | Player&#39;s country identifier, e.g. US | [optional] 
**Name** | Pointer to **string** | Verified player name, e.g. &#39;Miracle-&#39; | [optional] 
**CountryCode** | Pointer to **string** | Player&#39;s country code | [optional] 
**FantasyRole** | Pointer to **int32** | Player&#39;s ingame role (core: 1 or support: 2) | [optional] 
**TeamId** | Pointer to **int32** | Player&#39;s team identifier | [optional] 
**TeamName** | Pointer to **NullableString** | Team name | [optional] 
**TeamTag** | Pointer to **string** | Player&#39;s team shorthand tag, e.g. &#39;EG&#39; | [optional] 
**IsLocked** | Pointer to **bool** | Whether the roster lock is active | [optional] 
**IsPro** | Pointer to **bool** | Whether the player is professional or not | [optional] 
**LockedUntil** | Pointer to **int32** | When the roster lock will end | [optional] 

## Methods

### NewPlayerObjectResponse

`func NewPlayerObjectResponse() *PlayerObjectResponse`

NewPlayerObjectResponse instantiates a new PlayerObjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerObjectResponseWithDefaults

`func NewPlayerObjectResponseWithDefaults() *PlayerObjectResponse`

NewPlayerObjectResponseWithDefaults instantiates a new PlayerObjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *PlayerObjectResponse) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PlayerObjectResponse) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PlayerObjectResponse) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *PlayerObjectResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetSteamid

`func (o *PlayerObjectResponse) GetSteamid() string`

GetSteamid returns the Steamid field if non-nil, zero value otherwise.

### GetSteamidOk

`func (o *PlayerObjectResponse) GetSteamidOk() (*string, bool)`

GetSteamidOk returns a tuple with the Steamid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamid

`func (o *PlayerObjectResponse) SetSteamid(v string)`

SetSteamid sets Steamid field to given value.

### HasSteamid

`func (o *PlayerObjectResponse) HasSteamid() bool`

HasSteamid returns a boolean if a field has been set.

### GetAvatar

`func (o *PlayerObjectResponse) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *PlayerObjectResponse) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *PlayerObjectResponse) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *PlayerObjectResponse) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetAvatarmedium

`func (o *PlayerObjectResponse) GetAvatarmedium() string`

GetAvatarmedium returns the Avatarmedium field if non-nil, zero value otherwise.

### GetAvatarmediumOk

`func (o *PlayerObjectResponse) GetAvatarmediumOk() (*string, bool)`

GetAvatarmediumOk returns a tuple with the Avatarmedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarmedium

`func (o *PlayerObjectResponse) SetAvatarmedium(v string)`

SetAvatarmedium sets Avatarmedium field to given value.

### HasAvatarmedium

`func (o *PlayerObjectResponse) HasAvatarmedium() bool`

HasAvatarmedium returns a boolean if a field has been set.

### GetAvatarfull

`func (o *PlayerObjectResponse) GetAvatarfull() string`

GetAvatarfull returns the Avatarfull field if non-nil, zero value otherwise.

### GetAvatarfullOk

`func (o *PlayerObjectResponse) GetAvatarfullOk() (*string, bool)`

GetAvatarfullOk returns a tuple with the Avatarfull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarfull

`func (o *PlayerObjectResponse) SetAvatarfull(v string)`

SetAvatarfull sets Avatarfull field to given value.

### HasAvatarfull

`func (o *PlayerObjectResponse) HasAvatarfull() bool`

HasAvatarfull returns a boolean if a field has been set.

### GetProfileurl

`func (o *PlayerObjectResponse) GetProfileurl() string`

GetProfileurl returns the Profileurl field if non-nil, zero value otherwise.

### GetProfileurlOk

`func (o *PlayerObjectResponse) GetProfileurlOk() (*string, bool)`

GetProfileurlOk returns a tuple with the Profileurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileurl

`func (o *PlayerObjectResponse) SetProfileurl(v string)`

SetProfileurl sets Profileurl field to given value.

### HasProfileurl

`func (o *PlayerObjectResponse) HasProfileurl() bool`

HasProfileurl returns a boolean if a field has been set.

### GetPersonaname

`func (o *PlayerObjectResponse) GetPersonaname() string`

GetPersonaname returns the Personaname field if non-nil, zero value otherwise.

### GetPersonanameOk

`func (o *PlayerObjectResponse) GetPersonanameOk() (*string, bool)`

GetPersonanameOk returns a tuple with the Personaname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonaname

`func (o *PlayerObjectResponse) SetPersonaname(v string)`

SetPersonaname sets Personaname field to given value.

### HasPersonaname

`func (o *PlayerObjectResponse) HasPersonaname() bool`

HasPersonaname returns a boolean if a field has been set.

### SetPersonanameNil

`func (o *PlayerObjectResponse) SetPersonanameNil(b bool)`

 SetPersonanameNil sets the value for Personaname to be an explicit nil

### UnsetPersonaname
`func (o *PlayerObjectResponse) UnsetPersonaname()`

UnsetPersonaname ensures that no value is present for Personaname, not even an explicit nil
### GetLastLogin

`func (o *PlayerObjectResponse) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *PlayerObjectResponse) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *PlayerObjectResponse) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *PlayerObjectResponse) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetFullHistoryTime

`func (o *PlayerObjectResponse) GetFullHistoryTime() time.Time`

GetFullHistoryTime returns the FullHistoryTime field if non-nil, zero value otherwise.

### GetFullHistoryTimeOk

`func (o *PlayerObjectResponse) GetFullHistoryTimeOk() (*time.Time, bool)`

GetFullHistoryTimeOk returns a tuple with the FullHistoryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullHistoryTime

`func (o *PlayerObjectResponse) SetFullHistoryTime(v time.Time)`

SetFullHistoryTime sets FullHistoryTime field to given value.

### HasFullHistoryTime

`func (o *PlayerObjectResponse) HasFullHistoryTime() bool`

HasFullHistoryTime returns a boolean if a field has been set.

### GetCheese

`func (o *PlayerObjectResponse) GetCheese() int32`

GetCheese returns the Cheese field if non-nil, zero value otherwise.

### GetCheeseOk

`func (o *PlayerObjectResponse) GetCheeseOk() (*int32, bool)`

GetCheeseOk returns a tuple with the Cheese field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheese

`func (o *PlayerObjectResponse) SetCheese(v int32)`

SetCheese sets Cheese field to given value.

### HasCheese

`func (o *PlayerObjectResponse) HasCheese() bool`

HasCheese returns a boolean if a field has been set.

### GetFhUnavailable

`func (o *PlayerObjectResponse) GetFhUnavailable() bool`

GetFhUnavailable returns the FhUnavailable field if non-nil, zero value otherwise.

### GetFhUnavailableOk

`func (o *PlayerObjectResponse) GetFhUnavailableOk() (*bool, bool)`

GetFhUnavailableOk returns a tuple with the FhUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFhUnavailable

`func (o *PlayerObjectResponse) SetFhUnavailable(v bool)`

SetFhUnavailable sets FhUnavailable field to given value.

### HasFhUnavailable

`func (o *PlayerObjectResponse) HasFhUnavailable() bool`

HasFhUnavailable returns a boolean if a field has been set.

### GetLoccountrycode

`func (o *PlayerObjectResponse) GetLoccountrycode() string`

GetLoccountrycode returns the Loccountrycode field if non-nil, zero value otherwise.

### GetLoccountrycodeOk

`func (o *PlayerObjectResponse) GetLoccountrycodeOk() (*string, bool)`

GetLoccountrycodeOk returns a tuple with the Loccountrycode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoccountrycode

`func (o *PlayerObjectResponse) SetLoccountrycode(v string)`

SetLoccountrycode sets Loccountrycode field to given value.

### HasLoccountrycode

`func (o *PlayerObjectResponse) HasLoccountrycode() bool`

HasLoccountrycode returns a boolean if a field has been set.

### GetName

`func (o *PlayerObjectResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlayerObjectResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlayerObjectResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlayerObjectResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCountryCode

`func (o *PlayerObjectResponse) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PlayerObjectResponse) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PlayerObjectResponse) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PlayerObjectResponse) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetFantasyRole

`func (o *PlayerObjectResponse) GetFantasyRole() int32`

GetFantasyRole returns the FantasyRole field if non-nil, zero value otherwise.

### GetFantasyRoleOk

`func (o *PlayerObjectResponse) GetFantasyRoleOk() (*int32, bool)`

GetFantasyRoleOk returns a tuple with the FantasyRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFantasyRole

`func (o *PlayerObjectResponse) SetFantasyRole(v int32)`

SetFantasyRole sets FantasyRole field to given value.

### HasFantasyRole

`func (o *PlayerObjectResponse) HasFantasyRole() bool`

HasFantasyRole returns a boolean if a field has been set.

### GetTeamId

`func (o *PlayerObjectResponse) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *PlayerObjectResponse) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *PlayerObjectResponse) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *PlayerObjectResponse) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetTeamName

`func (o *PlayerObjectResponse) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *PlayerObjectResponse) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *PlayerObjectResponse) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *PlayerObjectResponse) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### SetTeamNameNil

`func (o *PlayerObjectResponse) SetTeamNameNil(b bool)`

 SetTeamNameNil sets the value for TeamName to be an explicit nil

### UnsetTeamName
`func (o *PlayerObjectResponse) UnsetTeamName()`

UnsetTeamName ensures that no value is present for TeamName, not even an explicit nil
### GetTeamTag

`func (o *PlayerObjectResponse) GetTeamTag() string`

GetTeamTag returns the TeamTag field if non-nil, zero value otherwise.

### GetTeamTagOk

`func (o *PlayerObjectResponse) GetTeamTagOk() (*string, bool)`

GetTeamTagOk returns a tuple with the TeamTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamTag

`func (o *PlayerObjectResponse) SetTeamTag(v string)`

SetTeamTag sets TeamTag field to given value.

### HasTeamTag

`func (o *PlayerObjectResponse) HasTeamTag() bool`

HasTeamTag returns a boolean if a field has been set.

### GetIsLocked

`func (o *PlayerObjectResponse) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *PlayerObjectResponse) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *PlayerObjectResponse) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *PlayerObjectResponse) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsPro

`func (o *PlayerObjectResponse) GetIsPro() bool`

GetIsPro returns the IsPro field if non-nil, zero value otherwise.

### GetIsProOk

`func (o *PlayerObjectResponse) GetIsProOk() (*bool, bool)`

GetIsProOk returns a tuple with the IsPro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPro

`func (o *PlayerObjectResponse) SetIsPro(v bool)`

SetIsPro sets IsPro field to given value.

### HasIsPro

`func (o *PlayerObjectResponse) HasIsPro() bool`

HasIsPro returns a boolean if a field has been set.

### GetLockedUntil

`func (o *PlayerObjectResponse) GetLockedUntil() int32`

GetLockedUntil returns the LockedUntil field if non-nil, zero value otherwise.

### GetLockedUntilOk

`func (o *PlayerObjectResponse) GetLockedUntilOk() (*int32, bool)`

GetLockedUntilOk returns a tuple with the LockedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedUntil

`func (o *PlayerObjectResponse) SetLockedUntil(v int32)`

SetLockedUntil sets LockedUntil field to given value.

### HasLockedUntil

`func (o *PlayerObjectResponse) HasLockedUntil() bool`

HasLockedUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


