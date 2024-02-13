# PublicMatchesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchId** | Pointer to **int32** | The ID number of the match assigned by Valve | [optional] 
**MatchSeqNum** | Pointer to **int32** | match_seq_num | [optional] 
**RadiantWin** | Pointer to **NullableBool** | Boolean indicating whether Radiant won the match | [optional] 
**StartTime** | Pointer to **int32** | The Unix timestamp at which the game started | [optional] 
**Duration** | Pointer to **int32** | Duration of the game in seconds | [optional] 
**LobbyType** | Pointer to **int32** |  | [optional] 
**GameMode** | Pointer to **int32** |  | [optional] 
**AvgRankTier** | Pointer to **int32** |  | [optional] 
**NumRankTier** | Pointer to **int32** |  | [optional] 
**Cluster** | Pointer to **int32** |  | [optional] 
**RadiantTeam** | Pointer to **[]int32** | radiant_team | [optional] 
**DireTeam** | Pointer to **[]int32** | dire_team | [optional] 

## Methods

### NewPublicMatchesResponse

`func NewPublicMatchesResponse() *PublicMatchesResponse`

NewPublicMatchesResponse instantiates a new PublicMatchesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicMatchesResponseWithDefaults

`func NewPublicMatchesResponseWithDefaults() *PublicMatchesResponse`

NewPublicMatchesResponseWithDefaults instantiates a new PublicMatchesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchId

`func (o *PublicMatchesResponse) GetMatchId() int32`

GetMatchId returns the MatchId field if non-nil, zero value otherwise.

### GetMatchIdOk

`func (o *PublicMatchesResponse) GetMatchIdOk() (*int32, bool)`

GetMatchIdOk returns a tuple with the MatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchId

`func (o *PublicMatchesResponse) SetMatchId(v int32)`

SetMatchId sets MatchId field to given value.

### HasMatchId

`func (o *PublicMatchesResponse) HasMatchId() bool`

HasMatchId returns a boolean if a field has been set.

### GetMatchSeqNum

`func (o *PublicMatchesResponse) GetMatchSeqNum() int32`

GetMatchSeqNum returns the MatchSeqNum field if non-nil, zero value otherwise.

### GetMatchSeqNumOk

`func (o *PublicMatchesResponse) GetMatchSeqNumOk() (*int32, bool)`

GetMatchSeqNumOk returns a tuple with the MatchSeqNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchSeqNum

`func (o *PublicMatchesResponse) SetMatchSeqNum(v int32)`

SetMatchSeqNum sets MatchSeqNum field to given value.

### HasMatchSeqNum

`func (o *PublicMatchesResponse) HasMatchSeqNum() bool`

HasMatchSeqNum returns a boolean if a field has been set.

### GetRadiantWin

`func (o *PublicMatchesResponse) GetRadiantWin() bool`

GetRadiantWin returns the RadiantWin field if non-nil, zero value otherwise.

### GetRadiantWinOk

`func (o *PublicMatchesResponse) GetRadiantWinOk() (*bool, bool)`

GetRadiantWinOk returns a tuple with the RadiantWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiantWin

`func (o *PublicMatchesResponse) SetRadiantWin(v bool)`

SetRadiantWin sets RadiantWin field to given value.

### HasRadiantWin

`func (o *PublicMatchesResponse) HasRadiantWin() bool`

HasRadiantWin returns a boolean if a field has been set.

### SetRadiantWinNil

`func (o *PublicMatchesResponse) SetRadiantWinNil(b bool)`

 SetRadiantWinNil sets the value for RadiantWin to be an explicit nil

### UnsetRadiantWin
`func (o *PublicMatchesResponse) UnsetRadiantWin()`

UnsetRadiantWin ensures that no value is present for RadiantWin, not even an explicit nil
### GetStartTime

`func (o *PublicMatchesResponse) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *PublicMatchesResponse) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *PublicMatchesResponse) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *PublicMatchesResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetDuration

`func (o *PublicMatchesResponse) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PublicMatchesResponse) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PublicMatchesResponse) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *PublicMatchesResponse) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetLobbyType

`func (o *PublicMatchesResponse) GetLobbyType() int32`

GetLobbyType returns the LobbyType field if non-nil, zero value otherwise.

### GetLobbyTypeOk

`func (o *PublicMatchesResponse) GetLobbyTypeOk() (*int32, bool)`

GetLobbyTypeOk returns a tuple with the LobbyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLobbyType

`func (o *PublicMatchesResponse) SetLobbyType(v int32)`

SetLobbyType sets LobbyType field to given value.

### HasLobbyType

`func (o *PublicMatchesResponse) HasLobbyType() bool`

HasLobbyType returns a boolean if a field has been set.

### GetGameMode

`func (o *PublicMatchesResponse) GetGameMode() int32`

GetGameMode returns the GameMode field if non-nil, zero value otherwise.

### GetGameModeOk

`func (o *PublicMatchesResponse) GetGameModeOk() (*int32, bool)`

GetGameModeOk returns a tuple with the GameMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameMode

`func (o *PublicMatchesResponse) SetGameMode(v int32)`

SetGameMode sets GameMode field to given value.

### HasGameMode

`func (o *PublicMatchesResponse) HasGameMode() bool`

HasGameMode returns a boolean if a field has been set.

### GetAvgRankTier

`func (o *PublicMatchesResponse) GetAvgRankTier() int32`

GetAvgRankTier returns the AvgRankTier field if non-nil, zero value otherwise.

### GetAvgRankTierOk

`func (o *PublicMatchesResponse) GetAvgRankTierOk() (*int32, bool)`

GetAvgRankTierOk returns a tuple with the AvgRankTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgRankTier

`func (o *PublicMatchesResponse) SetAvgRankTier(v int32)`

SetAvgRankTier sets AvgRankTier field to given value.

### HasAvgRankTier

`func (o *PublicMatchesResponse) HasAvgRankTier() bool`

HasAvgRankTier returns a boolean if a field has been set.

### GetNumRankTier

`func (o *PublicMatchesResponse) GetNumRankTier() int32`

GetNumRankTier returns the NumRankTier field if non-nil, zero value otherwise.

### GetNumRankTierOk

`func (o *PublicMatchesResponse) GetNumRankTierOk() (*int32, bool)`

GetNumRankTierOk returns a tuple with the NumRankTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRankTier

`func (o *PublicMatchesResponse) SetNumRankTier(v int32)`

SetNumRankTier sets NumRankTier field to given value.

### HasNumRankTier

`func (o *PublicMatchesResponse) HasNumRankTier() bool`

HasNumRankTier returns a boolean if a field has been set.

### GetCluster

`func (o *PublicMatchesResponse) GetCluster() int32`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PublicMatchesResponse) GetClusterOk() (*int32, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PublicMatchesResponse) SetCluster(v int32)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PublicMatchesResponse) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetRadiantTeam

`func (o *PublicMatchesResponse) GetRadiantTeam() []int32`

GetRadiantTeam returns the RadiantTeam field if non-nil, zero value otherwise.

### GetRadiantTeamOk

`func (o *PublicMatchesResponse) GetRadiantTeamOk() (*[]int32, bool)`

GetRadiantTeamOk returns a tuple with the RadiantTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiantTeam

`func (o *PublicMatchesResponse) SetRadiantTeam(v []int32)`

SetRadiantTeam sets RadiantTeam field to given value.

### HasRadiantTeam

`func (o *PublicMatchesResponse) HasRadiantTeam() bool`

HasRadiantTeam returns a boolean if a field has been set.

### GetDireTeam

`func (o *PublicMatchesResponse) GetDireTeam() []int32`

GetDireTeam returns the DireTeam field if non-nil, zero value otherwise.

### GetDireTeamOk

`func (o *PublicMatchesResponse) GetDireTeamOk() (*[]int32, bool)`

GetDireTeamOk returns a tuple with the DireTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDireTeam

`func (o *PublicMatchesResponse) SetDireTeam(v []int32)`

SetDireTeam sets DireTeam field to given value.

### HasDireTeam

`func (o *PublicMatchesResponse) HasDireTeam() bool`

HasDireTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


