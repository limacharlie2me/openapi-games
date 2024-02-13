# MatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchId** | Pointer to **int32** | The ID number of the match assigned by Valve | [optional] 
**BarracksStatusDire** | Pointer to **int32** | Bitmask. An integer that represents a binary of which barracks are still standing. 63 would mean all barracks still stand at the end of the game. | [optional] 
**BarracksStatusRadiant** | Pointer to **int32** | Bitmask. An integer that represents a binary of which barracks are still standing. 63 would mean all barracks still stand at the end of the game. | [optional] 
**Chat** | Pointer to [**[]MatchResponseChatInner**](MatchResponseChatInner.md) | Array containing information on the chat of the game | [optional] 
**Cluster** | Pointer to **int32** | cluster | [optional] 
**Cosmetics** | Pointer to **map[string]int32** | cosmetics | [optional] 
**DireScore** | Pointer to **int32** | Number of kills the Dire team had when the match ended | [optional] 
**DraftTimings** | Pointer to [**[]MatchResponseDraftTimingsInner**](MatchResponseDraftTimingsInner.md) | draft_timings | [optional] 
**Duration** | Pointer to **int32** | Duration of the game in seconds | [optional] 
**Engine** | Pointer to **int32** | engine | [optional] 
**FirstBloodTime** | Pointer to **int32** | Time in seconds at which first blood occurred | [optional] 
**GameMode** | Pointer to **int32** | Integer corresponding to game mode played. List of constants can be found here: https://github.com/odota/dotaconstants/blob/master/json/game_mode.json | [optional] 
**HumanPlayers** | Pointer to **int32** | Number of human players in the game | [optional] 
**Leagueid** | Pointer to **int32** | leagueid | [optional] 
**LobbyType** | Pointer to **int32** | Integer corresponding to lobby type of match. List of constants can be found here: https://github.com/odota/dotaconstants/blob/master/json/lobby_type.json | [optional] 
**MatchSeqNum** | Pointer to **int32** | match_seq_num | [optional] 
**NegativeVotes** | Pointer to **int32** | Number of negative votes the replay received in the in-game client | [optional] 
**Objectives** | Pointer to **[]map[string]interface{}** | objectives | [optional] 
**PicksBans** | Pointer to [**[]MatchResponsePicksBansInner**](MatchResponsePicksBansInner.md) | Array containing information on the draft. Each item contains a boolean relating to whether the choice is a pick or a ban, the hero ID, the team the picked or banned it, and the order. | [optional] 
**PositiveVotes** | Pointer to **int32** | Number of positive votes the replay received in the in-game client | [optional] 
**RadiantGoldAdv** | Pointer to **[]float32** | Array of the Radiant gold advantage at each minute in the game. A negative number means that Radiant is behind, and thus it is their gold disadvantage.  | [optional] 
**RadiantScore** | Pointer to **int32** | Number of kills the Radiant team had when the match ended | [optional] 
**RadiantWin** | Pointer to **NullableBool** | Boolean indicating whether Radiant won the match | [optional] 
**RadiantXpAdv** | Pointer to **[]float32** | Array of the Radiant experience advantage at each minute in the game. A negative number means that Radiant is behind, and thus it is their experience disadvantage.  | [optional] 
**StartTime** | Pointer to **int32** | The Unix timestamp at which the game started | [optional] 
**Teamfights** | Pointer to **[]map[string]interface{}** | teamfights | [optional] 
**TowerStatusDire** | Pointer to **int32** | Bitmask. An integer that represents a binary of which Dire towers are still standing. | [optional] 
**TowerStatusRadiant** | Pointer to **int32** | Bitmask. An integer that represents a binary of which Radiant towers are still standing. | [optional] 
**Version** | Pointer to **int32** | Parse version, used internally by OpenDota | [optional] 
**ReplaySalt** | Pointer to **int32** | replay_salt | [optional] 
**SeriesId** | Pointer to **int32** | series_id | [optional] 
**SeriesType** | Pointer to **int32** | series_type | [optional] 
**RadiantTeam** | Pointer to **map[string]interface{}** | radiant_team | [optional] 
**DireTeam** | Pointer to **map[string]interface{}** | dire_team | [optional] 
**League** | Pointer to **map[string]interface{}** | league | [optional] 
**Skill** | Pointer to **NullableInt32** | Skill bracket assigned by Valve (Normal, High, Very High) | [optional] 
**Players** | Pointer to [**[]MatchResponsePlayersInner**](MatchResponsePlayersInner.md) | Array of information on individual players | [optional] 
**Patch** | Pointer to **int32** | Patch ID, from dotaconstants | [optional] 
**Region** | Pointer to **int32** | Integer corresponding to the region the game was played on | [optional] 
**AllWordCounts** | Pointer to **map[string]interface{}** | Word counts of the all chat messages in the player&#39;s games | [optional] 
**MyWordCounts** | Pointer to **map[string]interface{}** | Word counts of the player&#39;s all chat messages | [optional] 
**Throw** | Pointer to **int32** | Maximum gold advantage of the player&#39;s team if they lost the match | [optional] 
**Comeback** | Pointer to **int32** | Maximum gold disadvantage of the player&#39;s team if they won the match | [optional] 
**Loss** | Pointer to **int32** | Maximum gold disadvantage of the player&#39;s team if they lost the match | [optional] 
**Win** | Pointer to **int32** | Maximum gold advantage of the player&#39;s team if they won the match | [optional] 
**ReplayUrl** | Pointer to **string** | replay_url | [optional] 

## Methods

### NewMatchResponse

`func NewMatchResponse() *MatchResponse`

NewMatchResponse instantiates a new MatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchResponseWithDefaults

`func NewMatchResponseWithDefaults() *MatchResponse`

NewMatchResponseWithDefaults instantiates a new MatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchId

`func (o *MatchResponse) GetMatchId() int32`

GetMatchId returns the MatchId field if non-nil, zero value otherwise.

### GetMatchIdOk

`func (o *MatchResponse) GetMatchIdOk() (*int32, bool)`

GetMatchIdOk returns a tuple with the MatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchId

`func (o *MatchResponse) SetMatchId(v int32)`

SetMatchId sets MatchId field to given value.

### HasMatchId

`func (o *MatchResponse) HasMatchId() bool`

HasMatchId returns a boolean if a field has been set.

### GetBarracksStatusDire

`func (o *MatchResponse) GetBarracksStatusDire() int32`

GetBarracksStatusDire returns the BarracksStatusDire field if non-nil, zero value otherwise.

### GetBarracksStatusDireOk

`func (o *MatchResponse) GetBarracksStatusDireOk() (*int32, bool)`

GetBarracksStatusDireOk returns a tuple with the BarracksStatusDire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarracksStatusDire

`func (o *MatchResponse) SetBarracksStatusDire(v int32)`

SetBarracksStatusDire sets BarracksStatusDire field to given value.

### HasBarracksStatusDire

`func (o *MatchResponse) HasBarracksStatusDire() bool`

HasBarracksStatusDire returns a boolean if a field has been set.

### GetBarracksStatusRadiant

`func (o *MatchResponse) GetBarracksStatusRadiant() int32`

GetBarracksStatusRadiant returns the BarracksStatusRadiant field if non-nil, zero value otherwise.

### GetBarracksStatusRadiantOk

`func (o *MatchResponse) GetBarracksStatusRadiantOk() (*int32, bool)`

GetBarracksStatusRadiantOk returns a tuple with the BarracksStatusRadiant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarracksStatusRadiant

`func (o *MatchResponse) SetBarracksStatusRadiant(v int32)`

SetBarracksStatusRadiant sets BarracksStatusRadiant field to given value.

### HasBarracksStatusRadiant

`func (o *MatchResponse) HasBarracksStatusRadiant() bool`

HasBarracksStatusRadiant returns a boolean if a field has been set.

### GetChat

`func (o *MatchResponse) GetChat() []MatchResponseChatInner`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *MatchResponse) GetChatOk() (*[]MatchResponseChatInner, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *MatchResponse) SetChat(v []MatchResponseChatInner)`

SetChat sets Chat field to given value.

### HasChat

`func (o *MatchResponse) HasChat() bool`

HasChat returns a boolean if a field has been set.

### GetCluster

`func (o *MatchResponse) GetCluster() int32`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *MatchResponse) GetClusterOk() (*int32, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *MatchResponse) SetCluster(v int32)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *MatchResponse) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCosmetics

`func (o *MatchResponse) GetCosmetics() map[string]int32`

GetCosmetics returns the Cosmetics field if non-nil, zero value otherwise.

### GetCosmeticsOk

`func (o *MatchResponse) GetCosmeticsOk() (*map[string]int32, bool)`

GetCosmeticsOk returns a tuple with the Cosmetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosmetics

`func (o *MatchResponse) SetCosmetics(v map[string]int32)`

SetCosmetics sets Cosmetics field to given value.

### HasCosmetics

`func (o *MatchResponse) HasCosmetics() bool`

HasCosmetics returns a boolean if a field has been set.

### GetDireScore

`func (o *MatchResponse) GetDireScore() int32`

GetDireScore returns the DireScore field if non-nil, zero value otherwise.

### GetDireScoreOk

`func (o *MatchResponse) GetDireScoreOk() (*int32, bool)`

GetDireScoreOk returns a tuple with the DireScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDireScore

`func (o *MatchResponse) SetDireScore(v int32)`

SetDireScore sets DireScore field to given value.

### HasDireScore

`func (o *MatchResponse) HasDireScore() bool`

HasDireScore returns a boolean if a field has been set.

### GetDraftTimings

`func (o *MatchResponse) GetDraftTimings() []MatchResponseDraftTimingsInner`

GetDraftTimings returns the DraftTimings field if non-nil, zero value otherwise.

### GetDraftTimingsOk

`func (o *MatchResponse) GetDraftTimingsOk() (*[]MatchResponseDraftTimingsInner, bool)`

GetDraftTimingsOk returns a tuple with the DraftTimings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftTimings

`func (o *MatchResponse) SetDraftTimings(v []MatchResponseDraftTimingsInner)`

SetDraftTimings sets DraftTimings field to given value.

### HasDraftTimings

`func (o *MatchResponse) HasDraftTimings() bool`

HasDraftTimings returns a boolean if a field has been set.

### GetDuration

`func (o *MatchResponse) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MatchResponse) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *MatchResponse) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *MatchResponse) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEngine

`func (o *MatchResponse) GetEngine() int32`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *MatchResponse) GetEngineOk() (*int32, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *MatchResponse) SetEngine(v int32)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *MatchResponse) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetFirstBloodTime

`func (o *MatchResponse) GetFirstBloodTime() int32`

GetFirstBloodTime returns the FirstBloodTime field if non-nil, zero value otherwise.

### GetFirstBloodTimeOk

`func (o *MatchResponse) GetFirstBloodTimeOk() (*int32, bool)`

GetFirstBloodTimeOk returns a tuple with the FirstBloodTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstBloodTime

`func (o *MatchResponse) SetFirstBloodTime(v int32)`

SetFirstBloodTime sets FirstBloodTime field to given value.

### HasFirstBloodTime

`func (o *MatchResponse) HasFirstBloodTime() bool`

HasFirstBloodTime returns a boolean if a field has been set.

### GetGameMode

`func (o *MatchResponse) GetGameMode() int32`

GetGameMode returns the GameMode field if non-nil, zero value otherwise.

### GetGameModeOk

`func (o *MatchResponse) GetGameModeOk() (*int32, bool)`

GetGameModeOk returns a tuple with the GameMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameMode

`func (o *MatchResponse) SetGameMode(v int32)`

SetGameMode sets GameMode field to given value.

### HasGameMode

`func (o *MatchResponse) HasGameMode() bool`

HasGameMode returns a boolean if a field has been set.

### GetHumanPlayers

`func (o *MatchResponse) GetHumanPlayers() int32`

GetHumanPlayers returns the HumanPlayers field if non-nil, zero value otherwise.

### GetHumanPlayersOk

`func (o *MatchResponse) GetHumanPlayersOk() (*int32, bool)`

GetHumanPlayersOk returns a tuple with the HumanPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumanPlayers

`func (o *MatchResponse) SetHumanPlayers(v int32)`

SetHumanPlayers sets HumanPlayers field to given value.

### HasHumanPlayers

`func (o *MatchResponse) HasHumanPlayers() bool`

HasHumanPlayers returns a boolean if a field has been set.

### GetLeagueid

`func (o *MatchResponse) GetLeagueid() int32`

GetLeagueid returns the Leagueid field if non-nil, zero value otherwise.

### GetLeagueidOk

`func (o *MatchResponse) GetLeagueidOk() (*int32, bool)`

GetLeagueidOk returns a tuple with the Leagueid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeagueid

`func (o *MatchResponse) SetLeagueid(v int32)`

SetLeagueid sets Leagueid field to given value.

### HasLeagueid

`func (o *MatchResponse) HasLeagueid() bool`

HasLeagueid returns a boolean if a field has been set.

### GetLobbyType

`func (o *MatchResponse) GetLobbyType() int32`

GetLobbyType returns the LobbyType field if non-nil, zero value otherwise.

### GetLobbyTypeOk

`func (o *MatchResponse) GetLobbyTypeOk() (*int32, bool)`

GetLobbyTypeOk returns a tuple with the LobbyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLobbyType

`func (o *MatchResponse) SetLobbyType(v int32)`

SetLobbyType sets LobbyType field to given value.

### HasLobbyType

`func (o *MatchResponse) HasLobbyType() bool`

HasLobbyType returns a boolean if a field has been set.

### GetMatchSeqNum

`func (o *MatchResponse) GetMatchSeqNum() int32`

GetMatchSeqNum returns the MatchSeqNum field if non-nil, zero value otherwise.

### GetMatchSeqNumOk

`func (o *MatchResponse) GetMatchSeqNumOk() (*int32, bool)`

GetMatchSeqNumOk returns a tuple with the MatchSeqNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchSeqNum

`func (o *MatchResponse) SetMatchSeqNum(v int32)`

SetMatchSeqNum sets MatchSeqNum field to given value.

### HasMatchSeqNum

`func (o *MatchResponse) HasMatchSeqNum() bool`

HasMatchSeqNum returns a boolean if a field has been set.

### GetNegativeVotes

`func (o *MatchResponse) GetNegativeVotes() int32`

GetNegativeVotes returns the NegativeVotes field if non-nil, zero value otherwise.

### GetNegativeVotesOk

`func (o *MatchResponse) GetNegativeVotesOk() (*int32, bool)`

GetNegativeVotesOk returns a tuple with the NegativeVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeVotes

`func (o *MatchResponse) SetNegativeVotes(v int32)`

SetNegativeVotes sets NegativeVotes field to given value.

### HasNegativeVotes

`func (o *MatchResponse) HasNegativeVotes() bool`

HasNegativeVotes returns a boolean if a field has been set.

### GetObjectives

`func (o *MatchResponse) GetObjectives() []map[string]interface{}`

GetObjectives returns the Objectives field if non-nil, zero value otherwise.

### GetObjectivesOk

`func (o *MatchResponse) GetObjectivesOk() (*[]map[string]interface{}, bool)`

GetObjectivesOk returns a tuple with the Objectives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectives

`func (o *MatchResponse) SetObjectives(v []map[string]interface{})`

SetObjectives sets Objectives field to given value.

### HasObjectives

`func (o *MatchResponse) HasObjectives() bool`

HasObjectives returns a boolean if a field has been set.

### GetPicksBans

`func (o *MatchResponse) GetPicksBans() []MatchResponsePicksBansInner`

GetPicksBans returns the PicksBans field if non-nil, zero value otherwise.

### GetPicksBansOk

`func (o *MatchResponse) GetPicksBansOk() (*[]MatchResponsePicksBansInner, bool)`

GetPicksBansOk returns a tuple with the PicksBans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicksBans

`func (o *MatchResponse) SetPicksBans(v []MatchResponsePicksBansInner)`

SetPicksBans sets PicksBans field to given value.

### HasPicksBans

`func (o *MatchResponse) HasPicksBans() bool`

HasPicksBans returns a boolean if a field has been set.

### GetPositiveVotes

`func (o *MatchResponse) GetPositiveVotes() int32`

GetPositiveVotes returns the PositiveVotes field if non-nil, zero value otherwise.

### GetPositiveVotesOk

`func (o *MatchResponse) GetPositiveVotesOk() (*int32, bool)`

GetPositiveVotesOk returns a tuple with the PositiveVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositiveVotes

`func (o *MatchResponse) SetPositiveVotes(v int32)`

SetPositiveVotes sets PositiveVotes field to given value.

### HasPositiveVotes

`func (o *MatchResponse) HasPositiveVotes() bool`

HasPositiveVotes returns a boolean if a field has been set.

### GetRadiantGoldAdv

`func (o *MatchResponse) GetRadiantGoldAdv() []float32`

GetRadiantGoldAdv returns the RadiantGoldAdv field if non-nil, zero value otherwise.

### GetRadiantGoldAdvOk

`func (o *MatchResponse) GetRadiantGoldAdvOk() (*[]float32, bool)`

GetRadiantGoldAdvOk returns a tuple with the RadiantGoldAdv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiantGoldAdv

`func (o *MatchResponse) SetRadiantGoldAdv(v []float32)`

SetRadiantGoldAdv sets RadiantGoldAdv field to given value.

### HasRadiantGoldAdv

`func (o *MatchResponse) HasRadiantGoldAdv() bool`

HasRadiantGoldAdv returns a boolean if a field has been set.

### GetRadiantScore

`func (o *MatchResponse) GetRadiantScore() int32`

GetRadiantScore returns the RadiantScore field if non-nil, zero value otherwise.

### GetRadiantScoreOk

`func (o *MatchResponse) GetRadiantScoreOk() (*int32, bool)`

GetRadiantScoreOk returns a tuple with the RadiantScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiantScore

`func (o *MatchResponse) SetRadiantScore(v int32)`

SetRadiantScore sets RadiantScore field to given value.

### HasRadiantScore

`func (o *MatchResponse) HasRadiantScore() bool`

HasRadiantScore returns a boolean if a field has been set.

### GetRadiantWin

`func (o *MatchResponse) GetRadiantWin() bool`

GetRadiantWin returns the RadiantWin field if non-nil, zero value otherwise.

### GetRadiantWinOk

`func (o *MatchResponse) GetRadiantWinOk() (*bool, bool)`

GetRadiantWinOk returns a tuple with the RadiantWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiantWin

`func (o *MatchResponse) SetRadiantWin(v bool)`

SetRadiantWin sets RadiantWin field to given value.

### HasRadiantWin

`func (o *MatchResponse) HasRadiantWin() bool`

HasRadiantWin returns a boolean if a field has been set.

### SetRadiantWinNil

`func (o *MatchResponse) SetRadiantWinNil(b bool)`

 SetRadiantWinNil sets the value for RadiantWin to be an explicit nil

### UnsetRadiantWin
`func (o *MatchResponse) UnsetRadiantWin()`

UnsetRadiantWin ensures that no value is present for RadiantWin, not even an explicit nil
### GetRadiantXpAdv

`func (o *MatchResponse) GetRadiantXpAdv() []float32`

GetRadiantXpAdv returns the RadiantXpAdv field if non-nil, zero value otherwise.

### GetRadiantXpAdvOk

`func (o *MatchResponse) GetRadiantXpAdvOk() (*[]float32, bool)`

GetRadiantXpAdvOk returns a tuple with the RadiantXpAdv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiantXpAdv

`func (o *MatchResponse) SetRadiantXpAdv(v []float32)`

SetRadiantXpAdv sets RadiantXpAdv field to given value.

### HasRadiantXpAdv

`func (o *MatchResponse) HasRadiantXpAdv() bool`

HasRadiantXpAdv returns a boolean if a field has been set.

### GetStartTime

`func (o *MatchResponse) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MatchResponse) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MatchResponse) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MatchResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTeamfights

`func (o *MatchResponse) GetTeamfights() []map[string]interface{}`

GetTeamfights returns the Teamfights field if non-nil, zero value otherwise.

### GetTeamfightsOk

`func (o *MatchResponse) GetTeamfightsOk() (*[]map[string]interface{}, bool)`

GetTeamfightsOk returns a tuple with the Teamfights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamfights

`func (o *MatchResponse) SetTeamfights(v []map[string]interface{})`

SetTeamfights sets Teamfights field to given value.

### HasTeamfights

`func (o *MatchResponse) HasTeamfights() bool`

HasTeamfights returns a boolean if a field has been set.

### SetTeamfightsNil

`func (o *MatchResponse) SetTeamfightsNil(b bool)`

 SetTeamfightsNil sets the value for Teamfights to be an explicit nil

### UnsetTeamfights
`func (o *MatchResponse) UnsetTeamfights()`

UnsetTeamfights ensures that no value is present for Teamfights, not even an explicit nil
### GetTowerStatusDire

`func (o *MatchResponse) GetTowerStatusDire() int32`

GetTowerStatusDire returns the TowerStatusDire field if non-nil, zero value otherwise.

### GetTowerStatusDireOk

`func (o *MatchResponse) GetTowerStatusDireOk() (*int32, bool)`

GetTowerStatusDireOk returns a tuple with the TowerStatusDire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTowerStatusDire

`func (o *MatchResponse) SetTowerStatusDire(v int32)`

SetTowerStatusDire sets TowerStatusDire field to given value.

### HasTowerStatusDire

`func (o *MatchResponse) HasTowerStatusDire() bool`

HasTowerStatusDire returns a boolean if a field has been set.

### GetTowerStatusRadiant

`func (o *MatchResponse) GetTowerStatusRadiant() int32`

GetTowerStatusRadiant returns the TowerStatusRadiant field if non-nil, zero value otherwise.

### GetTowerStatusRadiantOk

`func (o *MatchResponse) GetTowerStatusRadiantOk() (*int32, bool)`

GetTowerStatusRadiantOk returns a tuple with the TowerStatusRadiant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTowerStatusRadiant

`func (o *MatchResponse) SetTowerStatusRadiant(v int32)`

SetTowerStatusRadiant sets TowerStatusRadiant field to given value.

### HasTowerStatusRadiant

`func (o *MatchResponse) HasTowerStatusRadiant() bool`

HasTowerStatusRadiant returns a boolean if a field has been set.

### GetVersion

`func (o *MatchResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MatchResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MatchResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MatchResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetReplaySalt

`func (o *MatchResponse) GetReplaySalt() int32`

GetReplaySalt returns the ReplaySalt field if non-nil, zero value otherwise.

### GetReplaySaltOk

`func (o *MatchResponse) GetReplaySaltOk() (*int32, bool)`

GetReplaySaltOk returns a tuple with the ReplaySalt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaySalt

`func (o *MatchResponse) SetReplaySalt(v int32)`

SetReplaySalt sets ReplaySalt field to given value.

### HasReplaySalt

`func (o *MatchResponse) HasReplaySalt() bool`

HasReplaySalt returns a boolean if a field has been set.

### GetSeriesId

`func (o *MatchResponse) GetSeriesId() int32`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *MatchResponse) GetSeriesIdOk() (*int32, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *MatchResponse) SetSeriesId(v int32)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *MatchResponse) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetSeriesType

`func (o *MatchResponse) GetSeriesType() int32`

GetSeriesType returns the SeriesType field if non-nil, zero value otherwise.

### GetSeriesTypeOk

`func (o *MatchResponse) GetSeriesTypeOk() (*int32, bool)`

GetSeriesTypeOk returns a tuple with the SeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesType

`func (o *MatchResponse) SetSeriesType(v int32)`

SetSeriesType sets SeriesType field to given value.

### HasSeriesType

`func (o *MatchResponse) HasSeriesType() bool`

HasSeriesType returns a boolean if a field has been set.

### GetRadiantTeam

`func (o *MatchResponse) GetRadiantTeam() map[string]interface{}`

GetRadiantTeam returns the RadiantTeam field if non-nil, zero value otherwise.

### GetRadiantTeamOk

`func (o *MatchResponse) GetRadiantTeamOk() (*map[string]interface{}, bool)`

GetRadiantTeamOk returns a tuple with the RadiantTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiantTeam

`func (o *MatchResponse) SetRadiantTeam(v map[string]interface{})`

SetRadiantTeam sets RadiantTeam field to given value.

### HasRadiantTeam

`func (o *MatchResponse) HasRadiantTeam() bool`

HasRadiantTeam returns a boolean if a field has been set.

### GetDireTeam

`func (o *MatchResponse) GetDireTeam() map[string]interface{}`

GetDireTeam returns the DireTeam field if non-nil, zero value otherwise.

### GetDireTeamOk

`func (o *MatchResponse) GetDireTeamOk() (*map[string]interface{}, bool)`

GetDireTeamOk returns a tuple with the DireTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDireTeam

`func (o *MatchResponse) SetDireTeam(v map[string]interface{})`

SetDireTeam sets DireTeam field to given value.

### HasDireTeam

`func (o *MatchResponse) HasDireTeam() bool`

HasDireTeam returns a boolean if a field has been set.

### GetLeague

`func (o *MatchResponse) GetLeague() map[string]interface{}`

GetLeague returns the League field if non-nil, zero value otherwise.

### GetLeagueOk

`func (o *MatchResponse) GetLeagueOk() (*map[string]interface{}, bool)`

GetLeagueOk returns a tuple with the League field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeague

`func (o *MatchResponse) SetLeague(v map[string]interface{})`

SetLeague sets League field to given value.

### HasLeague

`func (o *MatchResponse) HasLeague() bool`

HasLeague returns a boolean if a field has been set.

### GetSkill

`func (o *MatchResponse) GetSkill() int32`

GetSkill returns the Skill field if non-nil, zero value otherwise.

### GetSkillOk

`func (o *MatchResponse) GetSkillOk() (*int32, bool)`

GetSkillOk returns a tuple with the Skill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkill

`func (o *MatchResponse) SetSkill(v int32)`

SetSkill sets Skill field to given value.

### HasSkill

`func (o *MatchResponse) HasSkill() bool`

HasSkill returns a boolean if a field has been set.

### SetSkillNil

`func (o *MatchResponse) SetSkillNil(b bool)`

 SetSkillNil sets the value for Skill to be an explicit nil

### UnsetSkill
`func (o *MatchResponse) UnsetSkill()`

UnsetSkill ensures that no value is present for Skill, not even an explicit nil
### GetPlayers

`func (o *MatchResponse) GetPlayers() []MatchResponsePlayersInner`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *MatchResponse) GetPlayersOk() (*[]MatchResponsePlayersInner, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *MatchResponse) SetPlayers(v []MatchResponsePlayersInner)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *MatchResponse) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.

### GetPatch

`func (o *MatchResponse) GetPatch() int32`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *MatchResponse) GetPatchOk() (*int32, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *MatchResponse) SetPatch(v int32)`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *MatchResponse) HasPatch() bool`

HasPatch returns a boolean if a field has been set.

### GetRegion

`func (o *MatchResponse) GetRegion() int32`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MatchResponse) GetRegionOk() (*int32, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MatchResponse) SetRegion(v int32)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MatchResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetAllWordCounts

`func (o *MatchResponse) GetAllWordCounts() map[string]interface{}`

GetAllWordCounts returns the AllWordCounts field if non-nil, zero value otherwise.

### GetAllWordCountsOk

`func (o *MatchResponse) GetAllWordCountsOk() (*map[string]interface{}, bool)`

GetAllWordCountsOk returns a tuple with the AllWordCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllWordCounts

`func (o *MatchResponse) SetAllWordCounts(v map[string]interface{})`

SetAllWordCounts sets AllWordCounts field to given value.

### HasAllWordCounts

`func (o *MatchResponse) HasAllWordCounts() bool`

HasAllWordCounts returns a boolean if a field has been set.

### GetMyWordCounts

`func (o *MatchResponse) GetMyWordCounts() map[string]interface{}`

GetMyWordCounts returns the MyWordCounts field if non-nil, zero value otherwise.

### GetMyWordCountsOk

`func (o *MatchResponse) GetMyWordCountsOk() (*map[string]interface{}, bool)`

GetMyWordCountsOk returns a tuple with the MyWordCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyWordCounts

`func (o *MatchResponse) SetMyWordCounts(v map[string]interface{})`

SetMyWordCounts sets MyWordCounts field to given value.

### HasMyWordCounts

`func (o *MatchResponse) HasMyWordCounts() bool`

HasMyWordCounts returns a boolean if a field has been set.

### GetThrow

`func (o *MatchResponse) GetThrow() int32`

GetThrow returns the Throw field if non-nil, zero value otherwise.

### GetThrowOk

`func (o *MatchResponse) GetThrowOk() (*int32, bool)`

GetThrowOk returns a tuple with the Throw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrow

`func (o *MatchResponse) SetThrow(v int32)`

SetThrow sets Throw field to given value.

### HasThrow

`func (o *MatchResponse) HasThrow() bool`

HasThrow returns a boolean if a field has been set.

### GetComeback

`func (o *MatchResponse) GetComeback() int32`

GetComeback returns the Comeback field if non-nil, zero value otherwise.

### GetComebackOk

`func (o *MatchResponse) GetComebackOk() (*int32, bool)`

GetComebackOk returns a tuple with the Comeback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComeback

`func (o *MatchResponse) SetComeback(v int32)`

SetComeback sets Comeback field to given value.

### HasComeback

`func (o *MatchResponse) HasComeback() bool`

HasComeback returns a boolean if a field has been set.

### GetLoss

`func (o *MatchResponse) GetLoss() int32`

GetLoss returns the Loss field if non-nil, zero value otherwise.

### GetLossOk

`func (o *MatchResponse) GetLossOk() (*int32, bool)`

GetLossOk returns a tuple with the Loss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoss

`func (o *MatchResponse) SetLoss(v int32)`

SetLoss sets Loss field to given value.

### HasLoss

`func (o *MatchResponse) HasLoss() bool`

HasLoss returns a boolean if a field has been set.

### GetWin

`func (o *MatchResponse) GetWin() int32`

GetWin returns the Win field if non-nil, zero value otherwise.

### GetWinOk

`func (o *MatchResponse) GetWinOk() (*int32, bool)`

GetWinOk returns a tuple with the Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWin

`func (o *MatchResponse) SetWin(v int32)`

SetWin sets Win field to given value.

### HasWin

`func (o *MatchResponse) HasWin() bool`

HasWin returns a boolean if a field has been set.

### GetReplayUrl

`func (o *MatchResponse) GetReplayUrl() string`

GetReplayUrl returns the ReplayUrl field if non-nil, zero value otherwise.

### GetReplayUrlOk

`func (o *MatchResponse) GetReplayUrlOk() (*string, bool)`

GetReplayUrlOk returns a tuple with the ReplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayUrl

`func (o *MatchResponse) SetReplayUrl(v string)`

SetReplayUrl sets ReplayUrl field to given value.

### HasReplayUrl

`func (o *MatchResponse) HasReplayUrl() bool`

HasReplayUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


