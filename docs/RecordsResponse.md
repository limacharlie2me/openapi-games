# RecordsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchId** | Pointer to **int32** | The ID number of the match assigned by Valve | [optional] 
**StartTime** | Pointer to **int32** | The Unix timestamp at which the game started | [optional] 
**HeroId** | Pointer to **int32** | The ID value of the hero played | [optional] 
**Score** | Pointer to **int32** | Record score | [optional] 

## Methods

### NewRecordsResponse

`func NewRecordsResponse() *RecordsResponse`

NewRecordsResponse instantiates a new RecordsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordsResponseWithDefaults

`func NewRecordsResponseWithDefaults() *RecordsResponse`

NewRecordsResponseWithDefaults instantiates a new RecordsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchId

`func (o *RecordsResponse) GetMatchId() int32`

GetMatchId returns the MatchId field if non-nil, zero value otherwise.

### GetMatchIdOk

`func (o *RecordsResponse) GetMatchIdOk() (*int32, bool)`

GetMatchIdOk returns a tuple with the MatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchId

`func (o *RecordsResponse) SetMatchId(v int32)`

SetMatchId sets MatchId field to given value.

### HasMatchId

`func (o *RecordsResponse) HasMatchId() bool`

HasMatchId returns a boolean if a field has been set.

### GetStartTime

`func (o *RecordsResponse) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RecordsResponse) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RecordsResponse) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RecordsResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetHeroId

`func (o *RecordsResponse) GetHeroId() int32`

GetHeroId returns the HeroId field if non-nil, zero value otherwise.

### GetHeroIdOk

`func (o *RecordsResponse) GetHeroIdOk() (*int32, bool)`

GetHeroIdOk returns a tuple with the HeroId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroId

`func (o *RecordsResponse) SetHeroId(v int32)`

SetHeroId sets HeroId field to given value.

### HasHeroId

`func (o *RecordsResponse) HasHeroId() bool`

HasHeroId returns a boolean if a field has been set.

### GetScore

`func (o *RecordsResponse) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *RecordsResponse) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *RecordsResponse) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *RecordsResponse) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


