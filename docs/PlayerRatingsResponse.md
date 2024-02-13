# PlayerRatingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** | The player account ID | [optional] 
**MatchId** | Pointer to **int32** | The ID number of the match assigned by Valve | [optional] 
**SoloCompetitiveRank** | Pointer to **NullableInt32** | solo_competitive_rank | [optional] 
**CompetitiveRank** | Pointer to **int32** | competitive_rank | [optional] 
**Time** | Pointer to **int32** | time | [optional] 

## Methods

### NewPlayerRatingsResponse

`func NewPlayerRatingsResponse() *PlayerRatingsResponse`

NewPlayerRatingsResponse instantiates a new PlayerRatingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerRatingsResponseWithDefaults

`func NewPlayerRatingsResponseWithDefaults() *PlayerRatingsResponse`

NewPlayerRatingsResponseWithDefaults instantiates a new PlayerRatingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *PlayerRatingsResponse) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PlayerRatingsResponse) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PlayerRatingsResponse) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *PlayerRatingsResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetMatchId

`func (o *PlayerRatingsResponse) GetMatchId() int32`

GetMatchId returns the MatchId field if non-nil, zero value otherwise.

### GetMatchIdOk

`func (o *PlayerRatingsResponse) GetMatchIdOk() (*int32, bool)`

GetMatchIdOk returns a tuple with the MatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchId

`func (o *PlayerRatingsResponse) SetMatchId(v int32)`

SetMatchId sets MatchId field to given value.

### HasMatchId

`func (o *PlayerRatingsResponse) HasMatchId() bool`

HasMatchId returns a boolean if a field has been set.

### GetSoloCompetitiveRank

`func (o *PlayerRatingsResponse) GetSoloCompetitiveRank() int32`

GetSoloCompetitiveRank returns the SoloCompetitiveRank field if non-nil, zero value otherwise.

### GetSoloCompetitiveRankOk

`func (o *PlayerRatingsResponse) GetSoloCompetitiveRankOk() (*int32, bool)`

GetSoloCompetitiveRankOk returns a tuple with the SoloCompetitiveRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoloCompetitiveRank

`func (o *PlayerRatingsResponse) SetSoloCompetitiveRank(v int32)`

SetSoloCompetitiveRank sets SoloCompetitiveRank field to given value.

### HasSoloCompetitiveRank

`func (o *PlayerRatingsResponse) HasSoloCompetitiveRank() bool`

HasSoloCompetitiveRank returns a boolean if a field has been set.

### SetSoloCompetitiveRankNil

`func (o *PlayerRatingsResponse) SetSoloCompetitiveRankNil(b bool)`

 SetSoloCompetitiveRankNil sets the value for SoloCompetitiveRank to be an explicit nil

### UnsetSoloCompetitiveRank
`func (o *PlayerRatingsResponse) UnsetSoloCompetitiveRank()`

UnsetSoloCompetitiveRank ensures that no value is present for SoloCompetitiveRank, not even an explicit nil
### GetCompetitiveRank

`func (o *PlayerRatingsResponse) GetCompetitiveRank() int32`

GetCompetitiveRank returns the CompetitiveRank field if non-nil, zero value otherwise.

### GetCompetitiveRankOk

`func (o *PlayerRatingsResponse) GetCompetitiveRankOk() (*int32, bool)`

GetCompetitiveRankOk returns a tuple with the CompetitiveRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitiveRank

`func (o *PlayerRatingsResponse) SetCompetitiveRank(v int32)`

SetCompetitiveRank sets CompetitiveRank field to given value.

### HasCompetitiveRank

`func (o *PlayerRatingsResponse) HasCompetitiveRank() bool`

HasCompetitiveRank returns a boolean if a field has been set.

### GetTime

`func (o *PlayerRatingsResponse) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *PlayerRatingsResponse) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *PlayerRatingsResponse) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *PlayerRatingsResponse) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


