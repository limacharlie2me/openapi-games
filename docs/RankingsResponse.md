# RankingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeroId** | Pointer to **int32** | The ID value of the hero played | [optional] 
**Rankings** | Pointer to [**[]RankingsResponseRankingsInner**](RankingsResponseRankingsInner.md) | rankings | [optional] 

## Methods

### NewRankingsResponse

`func NewRankingsResponse() *RankingsResponse`

NewRankingsResponse instantiates a new RankingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRankingsResponseWithDefaults

`func NewRankingsResponseWithDefaults() *RankingsResponse`

NewRankingsResponseWithDefaults instantiates a new RankingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeroId

`func (o *RankingsResponse) GetHeroId() int32`

GetHeroId returns the HeroId field if non-nil, zero value otherwise.

### GetHeroIdOk

`func (o *RankingsResponse) GetHeroIdOk() (*int32, bool)`

GetHeroIdOk returns a tuple with the HeroId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroId

`func (o *RankingsResponse) SetHeroId(v int32)`

SetHeroId sets HeroId field to given value.

### HasHeroId

`func (o *RankingsResponse) HasHeroId() bool`

HasHeroId returns a boolean if a field has been set.

### GetRankings

`func (o *RankingsResponse) GetRankings() []RankingsResponseRankingsInner`

GetRankings returns the Rankings field if non-nil, zero value otherwise.

### GetRankingsOk

`func (o *RankingsResponse) GetRankingsOk() (*[]RankingsResponseRankingsInner, bool)`

GetRankingsOk returns a tuple with the Rankings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankings

`func (o *RankingsResponse) SetRankings(v []RankingsResponseRankingsInner)`

SetRankings sets Rankings field to given value.

### HasRankings

`func (o *RankingsResponse) HasRankings() bool`

HasRankings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


