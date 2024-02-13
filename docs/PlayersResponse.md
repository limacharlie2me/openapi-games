# PlayersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SoloCompetitiveRank** | Pointer to **NullableInt32** | solo_competitive_rank | [optional] 
**CompetitiveRank** | Pointer to **NullableInt32** | competitive_rank | [optional] 
**RankTier** | Pointer to **NullableFloat32** | rank_tier | [optional] 
**LeaderboardRank** | Pointer to **NullableFloat32** | leaderboard_rank | [optional] 
**Profile** | Pointer to [**PlayersResponseProfile**](PlayersResponseProfile.md) |  | [optional] 

## Methods

### NewPlayersResponse

`func NewPlayersResponse() *PlayersResponse`

NewPlayersResponse instantiates a new PlayersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayersResponseWithDefaults

`func NewPlayersResponseWithDefaults() *PlayersResponse`

NewPlayersResponseWithDefaults instantiates a new PlayersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSoloCompetitiveRank

`func (o *PlayersResponse) GetSoloCompetitiveRank() int32`

GetSoloCompetitiveRank returns the SoloCompetitiveRank field if non-nil, zero value otherwise.

### GetSoloCompetitiveRankOk

`func (o *PlayersResponse) GetSoloCompetitiveRankOk() (*int32, bool)`

GetSoloCompetitiveRankOk returns a tuple with the SoloCompetitiveRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoloCompetitiveRank

`func (o *PlayersResponse) SetSoloCompetitiveRank(v int32)`

SetSoloCompetitiveRank sets SoloCompetitiveRank field to given value.

### HasSoloCompetitiveRank

`func (o *PlayersResponse) HasSoloCompetitiveRank() bool`

HasSoloCompetitiveRank returns a boolean if a field has been set.

### SetSoloCompetitiveRankNil

`func (o *PlayersResponse) SetSoloCompetitiveRankNil(b bool)`

 SetSoloCompetitiveRankNil sets the value for SoloCompetitiveRank to be an explicit nil

### UnsetSoloCompetitiveRank
`func (o *PlayersResponse) UnsetSoloCompetitiveRank()`

UnsetSoloCompetitiveRank ensures that no value is present for SoloCompetitiveRank, not even an explicit nil
### GetCompetitiveRank

`func (o *PlayersResponse) GetCompetitiveRank() int32`

GetCompetitiveRank returns the CompetitiveRank field if non-nil, zero value otherwise.

### GetCompetitiveRankOk

`func (o *PlayersResponse) GetCompetitiveRankOk() (*int32, bool)`

GetCompetitiveRankOk returns a tuple with the CompetitiveRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitiveRank

`func (o *PlayersResponse) SetCompetitiveRank(v int32)`

SetCompetitiveRank sets CompetitiveRank field to given value.

### HasCompetitiveRank

`func (o *PlayersResponse) HasCompetitiveRank() bool`

HasCompetitiveRank returns a boolean if a field has been set.

### SetCompetitiveRankNil

`func (o *PlayersResponse) SetCompetitiveRankNil(b bool)`

 SetCompetitiveRankNil sets the value for CompetitiveRank to be an explicit nil

### UnsetCompetitiveRank
`func (o *PlayersResponse) UnsetCompetitiveRank()`

UnsetCompetitiveRank ensures that no value is present for CompetitiveRank, not even an explicit nil
### GetRankTier

`func (o *PlayersResponse) GetRankTier() float32`

GetRankTier returns the RankTier field if non-nil, zero value otherwise.

### GetRankTierOk

`func (o *PlayersResponse) GetRankTierOk() (*float32, bool)`

GetRankTierOk returns a tuple with the RankTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankTier

`func (o *PlayersResponse) SetRankTier(v float32)`

SetRankTier sets RankTier field to given value.

### HasRankTier

`func (o *PlayersResponse) HasRankTier() bool`

HasRankTier returns a boolean if a field has been set.

### SetRankTierNil

`func (o *PlayersResponse) SetRankTierNil(b bool)`

 SetRankTierNil sets the value for RankTier to be an explicit nil

### UnsetRankTier
`func (o *PlayersResponse) UnsetRankTier()`

UnsetRankTier ensures that no value is present for RankTier, not even an explicit nil
### GetLeaderboardRank

`func (o *PlayersResponse) GetLeaderboardRank() float32`

GetLeaderboardRank returns the LeaderboardRank field if non-nil, zero value otherwise.

### GetLeaderboardRankOk

`func (o *PlayersResponse) GetLeaderboardRankOk() (*float32, bool)`

GetLeaderboardRankOk returns a tuple with the LeaderboardRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderboardRank

`func (o *PlayersResponse) SetLeaderboardRank(v float32)`

SetLeaderboardRank sets LeaderboardRank field to given value.

### HasLeaderboardRank

`func (o *PlayersResponse) HasLeaderboardRank() bool`

HasLeaderboardRank returns a boolean if a field has been set.

### SetLeaderboardRankNil

`func (o *PlayersResponse) SetLeaderboardRankNil(b bool)`

 SetLeaderboardRankNil sets the value for LeaderboardRank to be an explicit nil

### UnsetLeaderboardRank
`func (o *PlayersResponse) UnsetLeaderboardRank()`

UnsetLeaderboardRank ensures that no value is present for LeaderboardRank, not even an explicit nil
### GetProfile

`func (o *PlayersResponse) GetProfile() PlayersResponseProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PlayersResponse) GetProfileOk() (*PlayersResponseProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PlayersResponse) SetProfile(v PlayersResponseProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PlayersResponse) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


