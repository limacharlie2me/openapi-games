# TeamMatchObjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchId** | Pointer to **int32** | The ID number of the match assigned by Valve | [optional] 
**Radiant** | Pointer to **bool** | Whether the team/player/hero was on Radiant | [optional] 
**RadiantWin** | Pointer to **NullableBool** | Boolean indicating whether Radiant won the match | [optional] 
**RadiantScore** | Pointer to **int32** | Number of kills the Radiant team had when the match ended | [optional] 
**DireScore** | Pointer to **int32** | Number of kills the Dire team had when the match ended | [optional] 
**Duration** | Pointer to **int32** | Duration of the game in seconds | [optional] 
**StartTime** | Pointer to **int32** | The Unix timestamp at which the game started | [optional] 
**Leagueid** | Pointer to **int32** | Identifier for the league the match took place in | [optional] 
**LeagueName** | Pointer to **string** | Name of league the match took place in | [optional] 
**Cluster** | Pointer to **int32** | cluster | [optional] 
**OpposingTeamId** | Pointer to **int32** | Opposing team identifier | [optional] 
**OpposingTeamName** | Pointer to **NullableString** | Opposing team name, e.g. &#39;Evil Geniuses&#39; | [optional] 
**OpposingTeamLogo** | Pointer to **string** | Opposing team logo url | [optional] 

## Methods

### NewTeamMatchObjectResponse

`func NewTeamMatchObjectResponse() *TeamMatchObjectResponse`

NewTeamMatchObjectResponse instantiates a new TeamMatchObjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamMatchObjectResponseWithDefaults

`func NewTeamMatchObjectResponseWithDefaults() *TeamMatchObjectResponse`

NewTeamMatchObjectResponseWithDefaults instantiates a new TeamMatchObjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchId

`func (o *TeamMatchObjectResponse) GetMatchId() int32`

GetMatchId returns the MatchId field if non-nil, zero value otherwise.

### GetMatchIdOk

`func (o *TeamMatchObjectResponse) GetMatchIdOk() (*int32, bool)`

GetMatchIdOk returns a tuple with the MatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchId

`func (o *TeamMatchObjectResponse) SetMatchId(v int32)`

SetMatchId sets MatchId field to given value.

### HasMatchId

`func (o *TeamMatchObjectResponse) HasMatchId() bool`

HasMatchId returns a boolean if a field has been set.

### GetRadiant

`func (o *TeamMatchObjectResponse) GetRadiant() bool`

GetRadiant returns the Radiant field if non-nil, zero value otherwise.

### GetRadiantOk

`func (o *TeamMatchObjectResponse) GetRadiantOk() (*bool, bool)`

GetRadiantOk returns a tuple with the Radiant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiant

`func (o *TeamMatchObjectResponse) SetRadiant(v bool)`

SetRadiant sets Radiant field to given value.

### HasRadiant

`func (o *TeamMatchObjectResponse) HasRadiant() bool`

HasRadiant returns a boolean if a field has been set.

### GetRadiantWin

`func (o *TeamMatchObjectResponse) GetRadiantWin() bool`

GetRadiantWin returns the RadiantWin field if non-nil, zero value otherwise.

### GetRadiantWinOk

`func (o *TeamMatchObjectResponse) GetRadiantWinOk() (*bool, bool)`

GetRadiantWinOk returns a tuple with the RadiantWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiantWin

`func (o *TeamMatchObjectResponse) SetRadiantWin(v bool)`

SetRadiantWin sets RadiantWin field to given value.

### HasRadiantWin

`func (o *TeamMatchObjectResponse) HasRadiantWin() bool`

HasRadiantWin returns a boolean if a field has been set.

### SetRadiantWinNil

`func (o *TeamMatchObjectResponse) SetRadiantWinNil(b bool)`

 SetRadiantWinNil sets the value for RadiantWin to be an explicit nil

### UnsetRadiantWin
`func (o *TeamMatchObjectResponse) UnsetRadiantWin()`

UnsetRadiantWin ensures that no value is present for RadiantWin, not even an explicit nil
### GetRadiantScore

`func (o *TeamMatchObjectResponse) GetRadiantScore() int32`

GetRadiantScore returns the RadiantScore field if non-nil, zero value otherwise.

### GetRadiantScoreOk

`func (o *TeamMatchObjectResponse) GetRadiantScoreOk() (*int32, bool)`

GetRadiantScoreOk returns a tuple with the RadiantScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiantScore

`func (o *TeamMatchObjectResponse) SetRadiantScore(v int32)`

SetRadiantScore sets RadiantScore field to given value.

### HasRadiantScore

`func (o *TeamMatchObjectResponse) HasRadiantScore() bool`

HasRadiantScore returns a boolean if a field has been set.

### GetDireScore

`func (o *TeamMatchObjectResponse) GetDireScore() int32`

GetDireScore returns the DireScore field if non-nil, zero value otherwise.

### GetDireScoreOk

`func (o *TeamMatchObjectResponse) GetDireScoreOk() (*int32, bool)`

GetDireScoreOk returns a tuple with the DireScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDireScore

`func (o *TeamMatchObjectResponse) SetDireScore(v int32)`

SetDireScore sets DireScore field to given value.

### HasDireScore

`func (o *TeamMatchObjectResponse) HasDireScore() bool`

HasDireScore returns a boolean if a field has been set.

### GetDuration

`func (o *TeamMatchObjectResponse) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TeamMatchObjectResponse) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TeamMatchObjectResponse) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TeamMatchObjectResponse) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetStartTime

`func (o *TeamMatchObjectResponse) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TeamMatchObjectResponse) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TeamMatchObjectResponse) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TeamMatchObjectResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetLeagueid

`func (o *TeamMatchObjectResponse) GetLeagueid() int32`

GetLeagueid returns the Leagueid field if non-nil, zero value otherwise.

### GetLeagueidOk

`func (o *TeamMatchObjectResponse) GetLeagueidOk() (*int32, bool)`

GetLeagueidOk returns a tuple with the Leagueid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeagueid

`func (o *TeamMatchObjectResponse) SetLeagueid(v int32)`

SetLeagueid sets Leagueid field to given value.

### HasLeagueid

`func (o *TeamMatchObjectResponse) HasLeagueid() bool`

HasLeagueid returns a boolean if a field has been set.

### GetLeagueName

`func (o *TeamMatchObjectResponse) GetLeagueName() string`

GetLeagueName returns the LeagueName field if non-nil, zero value otherwise.

### GetLeagueNameOk

`func (o *TeamMatchObjectResponse) GetLeagueNameOk() (*string, bool)`

GetLeagueNameOk returns a tuple with the LeagueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeagueName

`func (o *TeamMatchObjectResponse) SetLeagueName(v string)`

SetLeagueName sets LeagueName field to given value.

### HasLeagueName

`func (o *TeamMatchObjectResponse) HasLeagueName() bool`

HasLeagueName returns a boolean if a field has been set.

### GetCluster

`func (o *TeamMatchObjectResponse) GetCluster() int32`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *TeamMatchObjectResponse) GetClusterOk() (*int32, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *TeamMatchObjectResponse) SetCluster(v int32)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *TeamMatchObjectResponse) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetOpposingTeamId

`func (o *TeamMatchObjectResponse) GetOpposingTeamId() int32`

GetOpposingTeamId returns the OpposingTeamId field if non-nil, zero value otherwise.

### GetOpposingTeamIdOk

`func (o *TeamMatchObjectResponse) GetOpposingTeamIdOk() (*int32, bool)`

GetOpposingTeamIdOk returns a tuple with the OpposingTeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpposingTeamId

`func (o *TeamMatchObjectResponse) SetOpposingTeamId(v int32)`

SetOpposingTeamId sets OpposingTeamId field to given value.

### HasOpposingTeamId

`func (o *TeamMatchObjectResponse) HasOpposingTeamId() bool`

HasOpposingTeamId returns a boolean if a field has been set.

### GetOpposingTeamName

`func (o *TeamMatchObjectResponse) GetOpposingTeamName() string`

GetOpposingTeamName returns the OpposingTeamName field if non-nil, zero value otherwise.

### GetOpposingTeamNameOk

`func (o *TeamMatchObjectResponse) GetOpposingTeamNameOk() (*string, bool)`

GetOpposingTeamNameOk returns a tuple with the OpposingTeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpposingTeamName

`func (o *TeamMatchObjectResponse) SetOpposingTeamName(v string)`

SetOpposingTeamName sets OpposingTeamName field to given value.

### HasOpposingTeamName

`func (o *TeamMatchObjectResponse) HasOpposingTeamName() bool`

HasOpposingTeamName returns a boolean if a field has been set.

### SetOpposingTeamNameNil

`func (o *TeamMatchObjectResponse) SetOpposingTeamNameNil(b bool)`

 SetOpposingTeamNameNil sets the value for OpposingTeamName to be an explicit nil

### UnsetOpposingTeamName
`func (o *TeamMatchObjectResponse) UnsetOpposingTeamName()`

UnsetOpposingTeamName ensures that no value is present for OpposingTeamName, not even an explicit nil
### GetOpposingTeamLogo

`func (o *TeamMatchObjectResponse) GetOpposingTeamLogo() string`

GetOpposingTeamLogo returns the OpposingTeamLogo field if non-nil, zero value otherwise.

### GetOpposingTeamLogoOk

`func (o *TeamMatchObjectResponse) GetOpposingTeamLogoOk() (*string, bool)`

GetOpposingTeamLogoOk returns a tuple with the OpposingTeamLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpposingTeamLogo

`func (o *TeamMatchObjectResponse) SetOpposingTeamLogo(v string)`

SetOpposingTeamLogo sets OpposingTeamLogo field to given value.

### HasOpposingTeamLogo

`func (o *TeamMatchObjectResponse) HasOpposingTeamLogo() bool`

HasOpposingTeamLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


