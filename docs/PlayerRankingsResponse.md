# PlayerRankingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeroId** | Pointer to **int32** | The ID value of the hero played | [optional] 
**Score** | Pointer to **float32** | Hero score | [optional] 
**PercentRank** | Pointer to **float32** | percent_rank | [optional] 
**Card** | Pointer to **int32** | numeric_rank | [optional] 

## Methods

### NewPlayerRankingsResponse

`func NewPlayerRankingsResponse() *PlayerRankingsResponse`

NewPlayerRankingsResponse instantiates a new PlayerRankingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerRankingsResponseWithDefaults

`func NewPlayerRankingsResponseWithDefaults() *PlayerRankingsResponse`

NewPlayerRankingsResponseWithDefaults instantiates a new PlayerRankingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeroId

`func (o *PlayerRankingsResponse) GetHeroId() int32`

GetHeroId returns the HeroId field if non-nil, zero value otherwise.

### GetHeroIdOk

`func (o *PlayerRankingsResponse) GetHeroIdOk() (*int32, bool)`

GetHeroIdOk returns a tuple with the HeroId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroId

`func (o *PlayerRankingsResponse) SetHeroId(v int32)`

SetHeroId sets HeroId field to given value.

### HasHeroId

`func (o *PlayerRankingsResponse) HasHeroId() bool`

HasHeroId returns a boolean if a field has been set.

### GetScore

`func (o *PlayerRankingsResponse) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *PlayerRankingsResponse) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *PlayerRankingsResponse) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *PlayerRankingsResponse) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetPercentRank

`func (o *PlayerRankingsResponse) GetPercentRank() float32`

GetPercentRank returns the PercentRank field if non-nil, zero value otherwise.

### GetPercentRankOk

`func (o *PlayerRankingsResponse) GetPercentRankOk() (*float32, bool)`

GetPercentRankOk returns a tuple with the PercentRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentRank

`func (o *PlayerRankingsResponse) SetPercentRank(v float32)`

SetPercentRank sets PercentRank field to given value.

### HasPercentRank

`func (o *PlayerRankingsResponse) HasPercentRank() bool`

HasPercentRank returns a boolean if a field has been set.

### GetCard

`func (o *PlayerRankingsResponse) GetCard() int32`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *PlayerRankingsResponse) GetCardOk() (*int32, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *PlayerRankingsResponse) SetCard(v int32)`

SetCard sets Card field to given value.

### HasCard

`func (o *PlayerRankingsResponse) HasCard() bool`

HasCard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


