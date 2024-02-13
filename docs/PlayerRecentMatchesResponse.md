# PlayerRecentMatchesResponse

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
**Kills** | Pointer to **int32** | Total kills the player had at the end of the match | [optional] 
**Deaths** | Pointer to **int32** | Total deaths the player had at the end of the match | [optional] 
**Assists** | Pointer to **int32** | Total assists the player had at the end of the match | [optional] 
**Skill** | Pointer to **NullableInt32** | Skill bracket assigned by Valve (Normal, High, Very High). If the skill is unknown, will return null. | [optional] 
**AverageRank** | Pointer to **NullableInt32** | Average rank of players with public match data | [optional] 
**XpPerMin** | Pointer to **int32** | Experience Per Minute obtained by the player | [optional] 
**GoldPerMin** | Pointer to **int32** | Average gold per minute of the player | [optional] 
**HeroDamage** | Pointer to **int32** | Total hero damage to enemy heroes | [optional] 
**HeroHealing** | Pointer to **int32** | Total healing of ally heroes | [optional] 
**LastHits** | Pointer to **int32** | Total last hits the player had at the end of the match | [optional] 
**Lane** | Pointer to **NullableInt32** | Integer corresponding to which lane the player laned in for the match | [optional] 
**LaneRole** | Pointer to **NullableInt32** | lane_role | [optional] 
**IsRoaming** | Pointer to **NullableBool** | Boolean describing whether or not the player roamed | [optional] 
**Cluster** | Pointer to **int32** | cluster | [optional] 
**LeaverStatus** | Pointer to **int32** | Integer describing whether or not the player left the game. 0: didn&#39;t leave. 1: left safely. 2+: Abandoned | [optional] 
**PartySize** | Pointer to **NullableInt32** | Size of the players party. If not in a party, will return 1. | [optional] 

## Methods

### NewPlayerRecentMatchesResponse

`func NewPlayerRecentMatchesResponse() *PlayerRecentMatchesResponse`

NewPlayerRecentMatchesResponse instantiates a new PlayerRecentMatchesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerRecentMatchesResponseWithDefaults

`func NewPlayerRecentMatchesResponseWithDefaults() *PlayerRecentMatchesResponse`

NewPlayerRecentMatchesResponseWithDefaults instantiates a new PlayerRecentMatchesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchId

`func (o *PlayerRecentMatchesResponse) GetMatchId() int32`

GetMatchId returns the MatchId field if non-nil, zero value otherwise.

### GetMatchIdOk

`func (o *PlayerRecentMatchesResponse) GetMatchIdOk() (*int32, bool)`

GetMatchIdOk returns a tuple with the MatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchId

`func (o *PlayerRecentMatchesResponse) SetMatchId(v int32)`

SetMatchId sets MatchId field to given value.

### HasMatchId

`func (o *PlayerRecentMatchesResponse) HasMatchId() bool`

HasMatchId returns a boolean if a field has been set.

### GetPlayerSlot

`func (o *PlayerRecentMatchesResponse) GetPlayerSlot() int32`

GetPlayerSlot returns the PlayerSlot field if non-nil, zero value otherwise.

### GetPlayerSlotOk

`func (o *PlayerRecentMatchesResponse) GetPlayerSlotOk() (*int32, bool)`

GetPlayerSlotOk returns a tuple with the PlayerSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerSlot

`func (o *PlayerRecentMatchesResponse) SetPlayerSlot(v int32)`

SetPlayerSlot sets PlayerSlot field to given value.

### HasPlayerSlot

`func (o *PlayerRecentMatchesResponse) HasPlayerSlot() bool`

HasPlayerSlot returns a boolean if a field has been set.

### SetPlayerSlotNil

`func (o *PlayerRecentMatchesResponse) SetPlayerSlotNil(b bool)`

 SetPlayerSlotNil sets the value for PlayerSlot to be an explicit nil

### UnsetPlayerSlot
`func (o *PlayerRecentMatchesResponse) UnsetPlayerSlot()`

UnsetPlayerSlot ensures that no value is present for PlayerSlot, not even an explicit nil
### GetRadiantWin

`func (o *PlayerRecentMatchesResponse) GetRadiantWin() bool`

GetRadiantWin returns the RadiantWin field if non-nil, zero value otherwise.

### GetRadiantWinOk

`func (o *PlayerRecentMatchesResponse) GetRadiantWinOk() (*bool, bool)`

GetRadiantWinOk returns a tuple with the RadiantWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiantWin

`func (o *PlayerRecentMatchesResponse) SetRadiantWin(v bool)`

SetRadiantWin sets RadiantWin field to given value.

### HasRadiantWin

`func (o *PlayerRecentMatchesResponse) HasRadiantWin() bool`

HasRadiantWin returns a boolean if a field has been set.

### SetRadiantWinNil

`func (o *PlayerRecentMatchesResponse) SetRadiantWinNil(b bool)`

 SetRadiantWinNil sets the value for RadiantWin to be an explicit nil

### UnsetRadiantWin
`func (o *PlayerRecentMatchesResponse) UnsetRadiantWin()`

UnsetRadiantWin ensures that no value is present for RadiantWin, not even an explicit nil
### GetDuration

`func (o *PlayerRecentMatchesResponse) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PlayerRecentMatchesResponse) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PlayerRecentMatchesResponse) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *PlayerRecentMatchesResponse) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetGameMode

`func (o *PlayerRecentMatchesResponse) GetGameMode() int32`

GetGameMode returns the GameMode field if non-nil, zero value otherwise.

### GetGameModeOk

`func (o *PlayerRecentMatchesResponse) GetGameModeOk() (*int32, bool)`

GetGameModeOk returns a tuple with the GameMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameMode

`func (o *PlayerRecentMatchesResponse) SetGameMode(v int32)`

SetGameMode sets GameMode field to given value.

### HasGameMode

`func (o *PlayerRecentMatchesResponse) HasGameMode() bool`

HasGameMode returns a boolean if a field has been set.

### GetLobbyType

`func (o *PlayerRecentMatchesResponse) GetLobbyType() int32`

GetLobbyType returns the LobbyType field if non-nil, zero value otherwise.

### GetLobbyTypeOk

`func (o *PlayerRecentMatchesResponse) GetLobbyTypeOk() (*int32, bool)`

GetLobbyTypeOk returns a tuple with the LobbyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLobbyType

`func (o *PlayerRecentMatchesResponse) SetLobbyType(v int32)`

SetLobbyType sets LobbyType field to given value.

### HasLobbyType

`func (o *PlayerRecentMatchesResponse) HasLobbyType() bool`

HasLobbyType returns a boolean if a field has been set.

### GetHeroId

`func (o *PlayerRecentMatchesResponse) GetHeroId() int32`

GetHeroId returns the HeroId field if non-nil, zero value otherwise.

### GetHeroIdOk

`func (o *PlayerRecentMatchesResponse) GetHeroIdOk() (*int32, bool)`

GetHeroIdOk returns a tuple with the HeroId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroId

`func (o *PlayerRecentMatchesResponse) SetHeroId(v int32)`

SetHeroId sets HeroId field to given value.

### HasHeroId

`func (o *PlayerRecentMatchesResponse) HasHeroId() bool`

HasHeroId returns a boolean if a field has been set.

### GetStartTime

`func (o *PlayerRecentMatchesResponse) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *PlayerRecentMatchesResponse) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *PlayerRecentMatchesResponse) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *PlayerRecentMatchesResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetVersion

`func (o *PlayerRecentMatchesResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PlayerRecentMatchesResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PlayerRecentMatchesResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PlayerRecentMatchesResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *PlayerRecentMatchesResponse) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *PlayerRecentMatchesResponse) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetKills

`func (o *PlayerRecentMatchesResponse) GetKills() int32`

GetKills returns the Kills field if non-nil, zero value otherwise.

### GetKillsOk

`func (o *PlayerRecentMatchesResponse) GetKillsOk() (*int32, bool)`

GetKillsOk returns a tuple with the Kills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKills

`func (o *PlayerRecentMatchesResponse) SetKills(v int32)`

SetKills sets Kills field to given value.

### HasKills

`func (o *PlayerRecentMatchesResponse) HasKills() bool`

HasKills returns a boolean if a field has been set.

### GetDeaths

`func (o *PlayerRecentMatchesResponse) GetDeaths() int32`

GetDeaths returns the Deaths field if non-nil, zero value otherwise.

### GetDeathsOk

`func (o *PlayerRecentMatchesResponse) GetDeathsOk() (*int32, bool)`

GetDeathsOk returns a tuple with the Deaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeaths

`func (o *PlayerRecentMatchesResponse) SetDeaths(v int32)`

SetDeaths sets Deaths field to given value.

### HasDeaths

`func (o *PlayerRecentMatchesResponse) HasDeaths() bool`

HasDeaths returns a boolean if a field has been set.

### GetAssists

`func (o *PlayerRecentMatchesResponse) GetAssists() int32`

GetAssists returns the Assists field if non-nil, zero value otherwise.

### GetAssistsOk

`func (o *PlayerRecentMatchesResponse) GetAssistsOk() (*int32, bool)`

GetAssistsOk returns a tuple with the Assists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssists

`func (o *PlayerRecentMatchesResponse) SetAssists(v int32)`

SetAssists sets Assists field to given value.

### HasAssists

`func (o *PlayerRecentMatchesResponse) HasAssists() bool`

HasAssists returns a boolean if a field has been set.

### GetSkill

`func (o *PlayerRecentMatchesResponse) GetSkill() int32`

GetSkill returns the Skill field if non-nil, zero value otherwise.

### GetSkillOk

`func (o *PlayerRecentMatchesResponse) GetSkillOk() (*int32, bool)`

GetSkillOk returns a tuple with the Skill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkill

`func (o *PlayerRecentMatchesResponse) SetSkill(v int32)`

SetSkill sets Skill field to given value.

### HasSkill

`func (o *PlayerRecentMatchesResponse) HasSkill() bool`

HasSkill returns a boolean if a field has been set.

### SetSkillNil

`func (o *PlayerRecentMatchesResponse) SetSkillNil(b bool)`

 SetSkillNil sets the value for Skill to be an explicit nil

### UnsetSkill
`func (o *PlayerRecentMatchesResponse) UnsetSkill()`

UnsetSkill ensures that no value is present for Skill, not even an explicit nil
### GetAverageRank

`func (o *PlayerRecentMatchesResponse) GetAverageRank() int32`

GetAverageRank returns the AverageRank field if non-nil, zero value otherwise.

### GetAverageRankOk

`func (o *PlayerRecentMatchesResponse) GetAverageRankOk() (*int32, bool)`

GetAverageRankOk returns a tuple with the AverageRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRank

`func (o *PlayerRecentMatchesResponse) SetAverageRank(v int32)`

SetAverageRank sets AverageRank field to given value.

### HasAverageRank

`func (o *PlayerRecentMatchesResponse) HasAverageRank() bool`

HasAverageRank returns a boolean if a field has been set.

### SetAverageRankNil

`func (o *PlayerRecentMatchesResponse) SetAverageRankNil(b bool)`

 SetAverageRankNil sets the value for AverageRank to be an explicit nil

### UnsetAverageRank
`func (o *PlayerRecentMatchesResponse) UnsetAverageRank()`

UnsetAverageRank ensures that no value is present for AverageRank, not even an explicit nil
### GetXpPerMin

`func (o *PlayerRecentMatchesResponse) GetXpPerMin() int32`

GetXpPerMin returns the XpPerMin field if non-nil, zero value otherwise.

### GetXpPerMinOk

`func (o *PlayerRecentMatchesResponse) GetXpPerMinOk() (*int32, bool)`

GetXpPerMinOk returns a tuple with the XpPerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXpPerMin

`func (o *PlayerRecentMatchesResponse) SetXpPerMin(v int32)`

SetXpPerMin sets XpPerMin field to given value.

### HasXpPerMin

`func (o *PlayerRecentMatchesResponse) HasXpPerMin() bool`

HasXpPerMin returns a boolean if a field has been set.

### GetGoldPerMin

`func (o *PlayerRecentMatchesResponse) GetGoldPerMin() int32`

GetGoldPerMin returns the GoldPerMin field if non-nil, zero value otherwise.

### GetGoldPerMinOk

`func (o *PlayerRecentMatchesResponse) GetGoldPerMinOk() (*int32, bool)`

GetGoldPerMinOk returns a tuple with the GoldPerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoldPerMin

`func (o *PlayerRecentMatchesResponse) SetGoldPerMin(v int32)`

SetGoldPerMin sets GoldPerMin field to given value.

### HasGoldPerMin

`func (o *PlayerRecentMatchesResponse) HasGoldPerMin() bool`

HasGoldPerMin returns a boolean if a field has been set.

### GetHeroDamage

`func (o *PlayerRecentMatchesResponse) GetHeroDamage() int32`

GetHeroDamage returns the HeroDamage field if non-nil, zero value otherwise.

### GetHeroDamageOk

`func (o *PlayerRecentMatchesResponse) GetHeroDamageOk() (*int32, bool)`

GetHeroDamageOk returns a tuple with the HeroDamage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroDamage

`func (o *PlayerRecentMatchesResponse) SetHeroDamage(v int32)`

SetHeroDamage sets HeroDamage field to given value.

### HasHeroDamage

`func (o *PlayerRecentMatchesResponse) HasHeroDamage() bool`

HasHeroDamage returns a boolean if a field has been set.

### GetHeroHealing

`func (o *PlayerRecentMatchesResponse) GetHeroHealing() int32`

GetHeroHealing returns the HeroHealing field if non-nil, zero value otherwise.

### GetHeroHealingOk

`func (o *PlayerRecentMatchesResponse) GetHeroHealingOk() (*int32, bool)`

GetHeroHealingOk returns a tuple with the HeroHealing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroHealing

`func (o *PlayerRecentMatchesResponse) SetHeroHealing(v int32)`

SetHeroHealing sets HeroHealing field to given value.

### HasHeroHealing

`func (o *PlayerRecentMatchesResponse) HasHeroHealing() bool`

HasHeroHealing returns a boolean if a field has been set.

### GetLastHits

`func (o *PlayerRecentMatchesResponse) GetLastHits() int32`

GetLastHits returns the LastHits field if non-nil, zero value otherwise.

### GetLastHitsOk

`func (o *PlayerRecentMatchesResponse) GetLastHitsOk() (*int32, bool)`

GetLastHitsOk returns a tuple with the LastHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHits

`func (o *PlayerRecentMatchesResponse) SetLastHits(v int32)`

SetLastHits sets LastHits field to given value.

### HasLastHits

`func (o *PlayerRecentMatchesResponse) HasLastHits() bool`

HasLastHits returns a boolean if a field has been set.

### GetLane

`func (o *PlayerRecentMatchesResponse) GetLane() int32`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *PlayerRecentMatchesResponse) GetLaneOk() (*int32, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *PlayerRecentMatchesResponse) SetLane(v int32)`

SetLane sets Lane field to given value.

### HasLane

`func (o *PlayerRecentMatchesResponse) HasLane() bool`

HasLane returns a boolean if a field has been set.

### SetLaneNil

`func (o *PlayerRecentMatchesResponse) SetLaneNil(b bool)`

 SetLaneNil sets the value for Lane to be an explicit nil

### UnsetLane
`func (o *PlayerRecentMatchesResponse) UnsetLane()`

UnsetLane ensures that no value is present for Lane, not even an explicit nil
### GetLaneRole

`func (o *PlayerRecentMatchesResponse) GetLaneRole() int32`

GetLaneRole returns the LaneRole field if non-nil, zero value otherwise.

### GetLaneRoleOk

`func (o *PlayerRecentMatchesResponse) GetLaneRoleOk() (*int32, bool)`

GetLaneRoleOk returns a tuple with the LaneRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaneRole

`func (o *PlayerRecentMatchesResponse) SetLaneRole(v int32)`

SetLaneRole sets LaneRole field to given value.

### HasLaneRole

`func (o *PlayerRecentMatchesResponse) HasLaneRole() bool`

HasLaneRole returns a boolean if a field has been set.

### SetLaneRoleNil

`func (o *PlayerRecentMatchesResponse) SetLaneRoleNil(b bool)`

 SetLaneRoleNil sets the value for LaneRole to be an explicit nil

### UnsetLaneRole
`func (o *PlayerRecentMatchesResponse) UnsetLaneRole()`

UnsetLaneRole ensures that no value is present for LaneRole, not even an explicit nil
### GetIsRoaming

`func (o *PlayerRecentMatchesResponse) GetIsRoaming() bool`

GetIsRoaming returns the IsRoaming field if non-nil, zero value otherwise.

### GetIsRoamingOk

`func (o *PlayerRecentMatchesResponse) GetIsRoamingOk() (*bool, bool)`

GetIsRoamingOk returns a tuple with the IsRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoaming

`func (o *PlayerRecentMatchesResponse) SetIsRoaming(v bool)`

SetIsRoaming sets IsRoaming field to given value.

### HasIsRoaming

`func (o *PlayerRecentMatchesResponse) HasIsRoaming() bool`

HasIsRoaming returns a boolean if a field has been set.

### SetIsRoamingNil

`func (o *PlayerRecentMatchesResponse) SetIsRoamingNil(b bool)`

 SetIsRoamingNil sets the value for IsRoaming to be an explicit nil

### UnsetIsRoaming
`func (o *PlayerRecentMatchesResponse) UnsetIsRoaming()`

UnsetIsRoaming ensures that no value is present for IsRoaming, not even an explicit nil
### GetCluster

`func (o *PlayerRecentMatchesResponse) GetCluster() int32`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PlayerRecentMatchesResponse) GetClusterOk() (*int32, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PlayerRecentMatchesResponse) SetCluster(v int32)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PlayerRecentMatchesResponse) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetLeaverStatus

`func (o *PlayerRecentMatchesResponse) GetLeaverStatus() int32`

GetLeaverStatus returns the LeaverStatus field if non-nil, zero value otherwise.

### GetLeaverStatusOk

`func (o *PlayerRecentMatchesResponse) GetLeaverStatusOk() (*int32, bool)`

GetLeaverStatusOk returns a tuple with the LeaverStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaverStatus

`func (o *PlayerRecentMatchesResponse) SetLeaverStatus(v int32)`

SetLeaverStatus sets LeaverStatus field to given value.

### HasLeaverStatus

`func (o *PlayerRecentMatchesResponse) HasLeaverStatus() bool`

HasLeaverStatus returns a boolean if a field has been set.

### GetPartySize

`func (o *PlayerRecentMatchesResponse) GetPartySize() int32`

GetPartySize returns the PartySize field if non-nil, zero value otherwise.

### GetPartySizeOk

`func (o *PlayerRecentMatchesResponse) GetPartySizeOk() (*int32, bool)`

GetPartySizeOk returns a tuple with the PartySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartySize

`func (o *PlayerRecentMatchesResponse) SetPartySize(v int32)`

SetPartySize sets PartySize field to given value.

### HasPartySize

`func (o *PlayerRecentMatchesResponse) HasPartySize() bool`

HasPartySize returns a boolean if a field has been set.

### SetPartySizeNil

`func (o *PlayerRecentMatchesResponse) SetPartySizeNil(b bool)`

 SetPartySizeNil sets the value for PartySize to be an explicit nil

### UnsetPartySize
`func (o *PlayerRecentMatchesResponse) UnsetPartySize()`

UnsetPartySize ensures that no value is present for PartySize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


