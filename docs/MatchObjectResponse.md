# MatchObjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchId** | Pointer to **int32** | The ID number of the match assigned by Valve | [optional] 
**Duration** | Pointer to **int32** | Duration of the game in seconds | [optional] 
**StartTime** | Pointer to **int32** | The Unix timestamp at which the game started | [optional] 
**RadiantTeamId** | Pointer to **int32** | The Radiant&#39;s team_id | [optional] 
**RadiantName** | Pointer to **string** | The Radiant&#39;s team name | [optional] 
**DireTeamId** | Pointer to **int32** | The Dire&#39;s team_id | [optional] 
**DireName** | Pointer to **string** | The Dire&#39;s team name | [optional] 
**Leagueid** | Pointer to **int32** | Identifier for the league the match took place in | [optional] 
**LeagueName** | Pointer to **string** | Name of league the match took place in | [optional] 
**SeriesId** | Pointer to **int32** | Identifier for the series of the match | [optional] 
**SeriesType** | Pointer to **int32** | Type of series the match was | [optional] 
**RadiantScore** | Pointer to **int32** | Number of kills the Radiant team had when the match ended | [optional] 
**DireScore** | Pointer to **int32** | Number of kills the Dire team had when the match ended | [optional] 
**RadiantWin** | Pointer to **NullableBool** | Boolean indicating whether Radiant won the match | [optional] 
**Radiant** | Pointer to **bool** | Whether the team/player/hero was on Radiant | [optional] 

## Methods

### NewMatchObjectResponse

`func NewMatchObjectResponse() *MatchObjectResponse`

NewMatchObjectResponse instantiates a new MatchObjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchObjectResponseWithDefaults

`func NewMatchObjectResponseWithDefaults() *MatchObjectResponse`

NewMatchObjectResponseWithDefaults instantiates a new MatchObjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchId

`func (o *MatchObjectResponse) GetMatchId() int32`

GetMatchId returns the MatchId field if non-nil, zero value otherwise.

### GetMatchIdOk

`func (o *MatchObjectResponse) GetMatchIdOk() (*int32, bool)`

GetMatchIdOk returns a tuple with the MatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchId

`func (o *MatchObjectResponse) SetMatchId(v int32)`

SetMatchId sets MatchId field to given value.

### HasMatchId

`func (o *MatchObjectResponse) HasMatchId() bool`

HasMatchId returns a boolean if a field has been set.

### GetDuration

`func (o *MatchObjectResponse) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MatchObjectResponse) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *MatchObjectResponse) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *MatchObjectResponse) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetStartTime

`func (o *MatchObjectResponse) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MatchObjectResponse) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MatchObjectResponse) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MatchObjectResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetRadiantTeamId

`func (o *MatchObjectResponse) GetRadiantTeamId() int32`

GetRadiantTeamId returns the RadiantTeamId field if non-nil, zero value otherwise.

### GetRadiantTeamIdOk

`func (o *MatchObjectResponse) GetRadiantTeamIdOk() (*int32, bool)`

GetRadiantTeamIdOk returns a tuple with the RadiantTeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiantTeamId

`func (o *MatchObjectResponse) SetRadiantTeamId(v int32)`

SetRadiantTeamId sets RadiantTeamId field to given value.

### HasRadiantTeamId

`func (o *MatchObjectResponse) HasRadiantTeamId() bool`

HasRadiantTeamId returns a boolean if a field has been set.

### GetRadiantName

`func (o *MatchObjectResponse) GetRadiantName() string`

GetRadiantName returns the RadiantName field if non-nil, zero value otherwise.

### GetRadiantNameOk

`func (o *MatchObjectResponse) GetRadiantNameOk() (*string, bool)`

GetRadiantNameOk returns a tuple with the RadiantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiantName

`func (o *MatchObjectResponse) SetRadiantName(v string)`

SetRadiantName sets RadiantName field to given value.

### HasRadiantName

`func (o *MatchObjectResponse) HasRadiantName() bool`

HasRadiantName returns a boolean if a field has been set.

### GetDireTeamId

`func (o *MatchObjectResponse) GetDireTeamId() int32`

GetDireTeamId returns the DireTeamId field if non-nil, zero value otherwise.

### GetDireTeamIdOk

`func (o *MatchObjectResponse) GetDireTeamIdOk() (*int32, bool)`

GetDireTeamIdOk returns a tuple with the DireTeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDireTeamId

`func (o *MatchObjectResponse) SetDireTeamId(v int32)`

SetDireTeamId sets DireTeamId field to given value.

### HasDireTeamId

`func (o *MatchObjectResponse) HasDireTeamId() bool`

HasDireTeamId returns a boolean if a field has been set.

### GetDireName

`func (o *MatchObjectResponse) GetDireName() string`

GetDireName returns the DireName field if non-nil, zero value otherwise.

### GetDireNameOk

`func (o *MatchObjectResponse) GetDireNameOk() (*string, bool)`

GetDireNameOk returns a tuple with the DireName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDireName

`func (o *MatchObjectResponse) SetDireName(v string)`

SetDireName sets DireName field to given value.

### HasDireName

`func (o *MatchObjectResponse) HasDireName() bool`

HasDireName returns a boolean if a field has been set.

### GetLeagueid

`func (o *MatchObjectResponse) GetLeagueid() int32`

GetLeagueid returns the Leagueid field if non-nil, zero value otherwise.

### GetLeagueidOk

`func (o *MatchObjectResponse) GetLeagueidOk() (*int32, bool)`

GetLeagueidOk returns a tuple with the Leagueid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeagueid

`func (o *MatchObjectResponse) SetLeagueid(v int32)`

SetLeagueid sets Leagueid field to given value.

### HasLeagueid

`func (o *MatchObjectResponse) HasLeagueid() bool`

HasLeagueid returns a boolean if a field has been set.

### GetLeagueName

`func (o *MatchObjectResponse) GetLeagueName() string`

GetLeagueName returns the LeagueName field if non-nil, zero value otherwise.

### GetLeagueNameOk

`func (o *MatchObjectResponse) GetLeagueNameOk() (*string, bool)`

GetLeagueNameOk returns a tuple with the LeagueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeagueName

`func (o *MatchObjectResponse) SetLeagueName(v string)`

SetLeagueName sets LeagueName field to given value.

### HasLeagueName

`func (o *MatchObjectResponse) HasLeagueName() bool`

HasLeagueName returns a boolean if a field has been set.

### GetSeriesId

`func (o *MatchObjectResponse) GetSeriesId() int32`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *MatchObjectResponse) GetSeriesIdOk() (*int32, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *MatchObjectResponse) SetSeriesId(v int32)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *MatchObjectResponse) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetSeriesType

`func (o *MatchObjectResponse) GetSeriesType() int32`

GetSeriesType returns the SeriesType field if non-nil, zero value otherwise.

### GetSeriesTypeOk

`func (o *MatchObjectResponse) GetSeriesTypeOk() (*int32, bool)`

GetSeriesTypeOk returns a tuple with the SeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesType

`func (o *MatchObjectResponse) SetSeriesType(v int32)`

SetSeriesType sets SeriesType field to given value.

### HasSeriesType

`func (o *MatchObjectResponse) HasSeriesType() bool`

HasSeriesType returns a boolean if a field has been set.

### GetRadiantScore

`func (o *MatchObjectResponse) GetRadiantScore() int32`

GetRadiantScore returns the RadiantScore field if non-nil, zero value otherwise.

### GetRadiantScoreOk

`func (o *MatchObjectResponse) GetRadiantScoreOk() (*int32, bool)`

GetRadiantScoreOk returns a tuple with the RadiantScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiantScore

`func (o *MatchObjectResponse) SetRadiantScore(v int32)`

SetRadiantScore sets RadiantScore field to given value.

### HasRadiantScore

`func (o *MatchObjectResponse) HasRadiantScore() bool`

HasRadiantScore returns a boolean if a field has been set.

### GetDireScore

`func (o *MatchObjectResponse) GetDireScore() int32`

GetDireScore returns the DireScore field if non-nil, zero value otherwise.

### GetDireScoreOk

`func (o *MatchObjectResponse) GetDireScoreOk() (*int32, bool)`

GetDireScoreOk returns a tuple with the DireScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDireScore

`func (o *MatchObjectResponse) SetDireScore(v int32)`

SetDireScore sets DireScore field to given value.

### HasDireScore

`func (o *MatchObjectResponse) HasDireScore() bool`

HasDireScore returns a boolean if a field has been set.

### GetRadiantWin

`func (o *MatchObjectResponse) GetRadiantWin() bool`

GetRadiantWin returns the RadiantWin field if non-nil, zero value otherwise.

### GetRadiantWinOk

`func (o *MatchObjectResponse) GetRadiantWinOk() (*bool, bool)`

GetRadiantWinOk returns a tuple with the RadiantWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiantWin

`func (o *MatchObjectResponse) SetRadiantWin(v bool)`

SetRadiantWin sets RadiantWin field to given value.

### HasRadiantWin

`func (o *MatchObjectResponse) HasRadiantWin() bool`

HasRadiantWin returns a boolean if a field has been set.

### SetRadiantWinNil

`func (o *MatchObjectResponse) SetRadiantWinNil(b bool)`

 SetRadiantWinNil sets the value for RadiantWin to be an explicit nil

### UnsetRadiantWin
`func (o *MatchObjectResponse) UnsetRadiantWin()`

UnsetRadiantWin ensures that no value is present for RadiantWin, not even an explicit nil
### GetRadiant

`func (o *MatchObjectResponse) GetRadiant() bool`

GetRadiant returns the Radiant field if non-nil, zero value otherwise.

### GetRadiantOk

`func (o *MatchObjectResponse) GetRadiantOk() (*bool, bool)`

GetRadiantOk returns a tuple with the Radiant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiant

`func (o *MatchObjectResponse) SetRadiant(v bool)`

SetRadiant sets Radiant field to given value.

### HasRadiant

`func (o *MatchObjectResponse) HasRadiant() bool`

HasRadiant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


