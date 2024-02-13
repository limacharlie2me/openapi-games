# PlayerMatchesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchId** | Pointer to **int32** | The ID number of the match assigned by Valve | [optional] 
**PlayerSlot** | Pointer to **NullableInt32** | Which slot the player is in. 0-127 are Radiant, 128-255 are Dire | [optional] 
**RadiantWin** | Pointer to **NullableBool** | Boolean indicating whether Radiant won the match | [optional] 
**Duration** | Pointer to **int32** | Duration of the game in seconds | [optional] 
**GameMode** | Pointer to **int32** | Integer corresponding to game mode played. List of constants can be found here: https://github.com/odota/dotaconstants/blob/master/json/game_mode.json | [optional] 
**LobbyType** | Pointer to **int32** | Integer corresponding to lobby type of match. List of constants can be found here: https://github.com/odota/dotaconstants/blob/master/json/lobby_type.json | [optional] 
**HeroId** | Pointer to **int32** | The ID value of the hero played | [optional] 
**StartTime** | Pointer to **int32** | The Unix timestamp at which the game started | [optional] 
**Version** | Pointer to **NullableInt32** | version | [optional] 
**Kills** | Pointer to **int32** | Total kills the player had at the end of the game | [optional] 
**Deaths** | Pointer to **int32** | Total deaths the player had at the end of the game | [optional] 
**Assists** | Pointer to **int32** | Total assists the player had at the end of the game | [optional] 
**Skill** | Pointer to **NullableInt32** | Skill bracket assigned by Valve (Normal, High, Very High) | [optional] 
**AverageRank** | Pointer to **NullableInt32** | Average rank of players with public match data | [optional] 
**LeaverStatus** | Pointer to **int32** | Integer describing whether or not the player left the game. 0: didn&#39;t leave. 1: left safely. 2+: Abandoned | [optional] 
**PartySize** | Pointer to **NullableInt32** | Size of the player&#39;s party | [optional] 

## Methods

### NewPlayerMatchesResponse

`func NewPlayerMatchesResponse() *PlayerMatchesResponse`

NewPlayerMatchesResponse instantiates a new PlayerMatchesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerMatchesResponseWithDefaults

`func NewPlayerMatchesResponseWithDefaults() *PlayerMatchesResponse`

NewPlayerMatchesResponseWithDefaults instantiates a new PlayerMatchesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchId

`func (o *PlayerMatchesResponse) GetMatchId() int32`

GetMatchId returns the MatchId field if non-nil, zero value otherwise.

### GetMatchIdOk

`func (o *PlayerMatchesResponse) GetMatchIdOk() (*int32, bool)`

GetMatchIdOk returns a tuple with the MatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchId

`func (o *PlayerMatchesResponse) SetMatchId(v int32)`

SetMatchId sets MatchId field to given value.

### HasMatchId

`func (o *PlayerMatchesResponse) HasMatchId() bool`

HasMatchId returns a boolean if a field has been set.

### GetPlayerSlot

`func (o *PlayerMatchesResponse) GetPlayerSlot() int32`

GetPlayerSlot returns the PlayerSlot field if non-nil, zero value otherwise.

### GetPlayerSlotOk

`func (o *PlayerMatchesResponse) GetPlayerSlotOk() (*int32, bool)`

GetPlayerSlotOk returns a tuple with the PlayerSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerSlot

`func (o *PlayerMatchesResponse) SetPlayerSlot(v int32)`

SetPlayerSlot sets PlayerSlot field to given value.

### HasPlayerSlot

`func (o *PlayerMatchesResponse) HasPlayerSlot() bool`

HasPlayerSlot returns a boolean if a field has been set.

### SetPlayerSlotNil

`func (o *PlayerMatchesResponse) SetPlayerSlotNil(b bool)`

 SetPlayerSlotNil sets the value for PlayerSlot to be an explicit nil

### UnsetPlayerSlot
`func (o *PlayerMatchesResponse) UnsetPlayerSlot()`

UnsetPlayerSlot ensures that no value is present for PlayerSlot, not even an explicit nil
### GetRadiantWin

`func (o *PlayerMatchesResponse) GetRadiantWin() bool`

GetRadiantWin returns the RadiantWin field if non-nil, zero value otherwise.

### GetRadiantWinOk

`func (o *PlayerMatchesResponse) GetRadiantWinOk() (*bool, bool)`

GetRadiantWinOk returns a tuple with the RadiantWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiantWin

`func (o *PlayerMatchesResponse) SetRadiantWin(v bool)`

SetRadiantWin sets RadiantWin field to given value.

### HasRadiantWin

`func (o *PlayerMatchesResponse) HasRadiantWin() bool`

HasRadiantWin returns a boolean if a field has been set.

### SetRadiantWinNil

`func (o *PlayerMatchesResponse) SetRadiantWinNil(b bool)`

 SetRadiantWinNil sets the value for RadiantWin to be an explicit nil

### UnsetRadiantWin
`func (o *PlayerMatchesResponse) UnsetRadiantWin()`

UnsetRadiantWin ensures that no value is present for RadiantWin, not even an explicit nil
### GetDuration

`func (o *PlayerMatchesResponse) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PlayerMatchesResponse) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PlayerMatchesResponse) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *PlayerMatchesResponse) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetGameMode

`func (o *PlayerMatchesResponse) GetGameMode() int32`

GetGameMode returns the GameMode field if non-nil, zero value otherwise.

### GetGameModeOk

`func (o *PlayerMatchesResponse) GetGameModeOk() (*int32, bool)`

GetGameModeOk returns a tuple with the GameMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameMode

`func (o *PlayerMatchesResponse) SetGameMode(v int32)`

SetGameMode sets GameMode field to given value.

### HasGameMode

`func (o *PlayerMatchesResponse) HasGameMode() bool`

HasGameMode returns a boolean if a field has been set.

### GetLobbyType

`func (o *PlayerMatchesResponse) GetLobbyType() int32`

GetLobbyType returns the LobbyType field if non-nil, zero value otherwise.

### GetLobbyTypeOk

`func (o *PlayerMatchesResponse) GetLobbyTypeOk() (*int32, bool)`

GetLobbyTypeOk returns a tuple with the LobbyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLobbyType

`func (o *PlayerMatchesResponse) SetLobbyType(v int32)`

SetLobbyType sets LobbyType field to given value.

### HasLobbyType

`func (o *PlayerMatchesResponse) HasLobbyType() bool`

HasLobbyType returns a boolean if a field has been set.

### GetHeroId

`func (o *PlayerMatchesResponse) GetHeroId() int32`

GetHeroId returns the HeroId field if non-nil, zero value otherwise.

### GetHeroIdOk

`func (o *PlayerMatchesResponse) GetHeroIdOk() (*int32, bool)`

GetHeroIdOk returns a tuple with the HeroId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroId

`func (o *PlayerMatchesResponse) SetHeroId(v int32)`

SetHeroId sets HeroId field to given value.

### HasHeroId

`func (o *PlayerMatchesResponse) HasHeroId() bool`

HasHeroId returns a boolean if a field has been set.

### GetStartTime

`func (o *PlayerMatchesResponse) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *PlayerMatchesResponse) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *PlayerMatchesResponse) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *PlayerMatchesResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetVersion

`func (o *PlayerMatchesResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PlayerMatchesResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PlayerMatchesResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PlayerMatchesResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *PlayerMatchesResponse) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *PlayerMatchesResponse) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetKills

`func (o *PlayerMatchesResponse) GetKills() int32`

GetKills returns the Kills field if non-nil, zero value otherwise.

### GetKillsOk

`func (o *PlayerMatchesResponse) GetKillsOk() (*int32, bool)`

GetKillsOk returns a tuple with the Kills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKills

`func (o *PlayerMatchesResponse) SetKills(v int32)`

SetKills sets Kills field to given value.

### HasKills

`func (o *PlayerMatchesResponse) HasKills() bool`

HasKills returns a boolean if a field has been set.

### GetDeaths

`func (o *PlayerMatchesResponse) GetDeaths() int32`

GetDeaths returns the Deaths field if non-nil, zero value otherwise.

### GetDeathsOk

`func (o *PlayerMatchesResponse) GetDeathsOk() (*int32, bool)`

GetDeathsOk returns a tuple with the Deaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeaths

`func (o *PlayerMatchesResponse) SetDeaths(v int32)`

SetDeaths sets Deaths field to given value.

### HasDeaths

`func (o *PlayerMatchesResponse) HasDeaths() bool`

HasDeaths returns a boolean if a field has been set.

### GetAssists

`func (o *PlayerMatchesResponse) GetAssists() int32`

GetAssists returns the Assists field if non-nil, zero value otherwise.

### GetAssistsOk

`func (o *PlayerMatchesResponse) GetAssistsOk() (*int32, bool)`

GetAssistsOk returns a tuple with the Assists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssists

`func (o *PlayerMatchesResponse) SetAssists(v int32)`

SetAssists sets Assists field to given value.

### HasAssists

`func (o *PlayerMatchesResponse) HasAssists() bool`

HasAssists returns a boolean if a field has been set.

### GetSkill

`func (o *PlayerMatchesResponse) GetSkill() int32`

GetSkill returns the Skill field if non-nil, zero value otherwise.

### GetSkillOk

`func (o *PlayerMatchesResponse) GetSkillOk() (*int32, bool)`

GetSkillOk returns a tuple with the Skill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkill

`func (o *PlayerMatchesResponse) SetSkill(v int32)`

SetSkill sets Skill field to given value.

### HasSkill

`func (o *PlayerMatchesResponse) HasSkill() bool`

HasSkill returns a boolean if a field has been set.

### SetSkillNil

`func (o *PlayerMatchesResponse) SetSkillNil(b bool)`

 SetSkillNil sets the value for Skill to be an explicit nil

### UnsetSkill
`func (o *PlayerMatchesResponse) UnsetSkill()`

UnsetSkill ensures that no value is present for Skill, not even an explicit nil
### GetAverageRank

`func (o *PlayerMatchesResponse) GetAverageRank() int32`

GetAverageRank returns the AverageRank field if non-nil, zero value otherwise.

### GetAverageRankOk

`func (o *PlayerMatchesResponse) GetAverageRankOk() (*int32, bool)`

GetAverageRankOk returns a tuple with the AverageRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRank

`func (o *PlayerMatchesResponse) SetAverageRank(v int32)`

SetAverageRank sets AverageRank field to given value.

### HasAverageRank

`func (o *PlayerMatchesResponse) HasAverageRank() bool`

HasAverageRank returns a boolean if a field has been set.

### SetAverageRankNil

`func (o *PlayerMatchesResponse) SetAverageRankNil(b bool)`

 SetAverageRankNil sets the value for AverageRank to be an explicit nil

### UnsetAverageRank
`func (o *PlayerMatchesResponse) UnsetAverageRank()`

UnsetAverageRank ensures that no value is present for AverageRank, not even an explicit nil
### GetLeaverStatus

`func (o *PlayerMatchesResponse) GetLeaverStatus() int32`

GetLeaverStatus returns the LeaverStatus field if non-nil, zero value otherwise.

### GetLeaverStatusOk

`func (o *PlayerMatchesResponse) GetLeaverStatusOk() (*int32, bool)`

GetLeaverStatusOk returns a tuple with the LeaverStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaverStatus

`func (o *PlayerMatchesResponse) SetLeaverStatus(v int32)`

SetLeaverStatus sets LeaverStatus field to given value.

### HasLeaverStatus

`func (o *PlayerMatchesResponse) HasLeaverStatus() bool`

HasLeaverStatus returns a boolean if a field has been set.

### GetPartySize

`func (o *PlayerMatchesResponse) GetPartySize() int32`

GetPartySize returns the PartySize field if non-nil, zero value otherwise.

### GetPartySizeOk

`func (o *PlayerMatchesResponse) GetPartySizeOk() (*int32, bool)`

GetPartySizeOk returns a tuple with the PartySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartySize

`func (o *PlayerMatchesResponse) SetPartySize(v int32)`

SetPartySize sets PartySize field to given value.

### HasPartySize

`func (o *PlayerMatchesResponse) HasPartySize() bool`

HasPartySize returns a boolean if a field has been set.

### SetPartySizeNil

`func (o *PlayerMatchesResponse) SetPartySizeNil(b bool)`

 SetPartySizeNil sets the value for PartySize to be an explicit nil

### UnsetPartySize
`func (o *PlayerMatchesResponse) UnsetPartySize()`

UnsetPartySize ensures that no value is present for PartySize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


