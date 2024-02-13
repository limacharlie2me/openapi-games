# TeamHeroesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeroId** | Pointer to **int32** | The ID value of the hero played | [optional] 
**Name** | Pointer to **string** | Hero name | [optional] 
**GamesPlayed** | Pointer to **int32** | Number of games played | [optional] 
**Wins** | Pointer to **int32** | Number of wins | [optional] 

## Methods

### NewTeamHeroesResponse

`func NewTeamHeroesResponse() *TeamHeroesResponse`

NewTeamHeroesResponse instantiates a new TeamHeroesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamHeroesResponseWithDefaults

`func NewTeamHeroesResponseWithDefaults() *TeamHeroesResponse`

NewTeamHeroesResponseWithDefaults instantiates a new TeamHeroesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeroId

`func (o *TeamHeroesResponse) GetHeroId() int32`

GetHeroId returns the HeroId field if non-nil, zero value otherwise.

### GetHeroIdOk

`func (o *TeamHeroesResponse) GetHeroIdOk() (*int32, bool)`

GetHeroIdOk returns a tuple with the HeroId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroId

`func (o *TeamHeroesResponse) SetHeroId(v int32)`

SetHeroId sets HeroId field to given value.

### HasHeroId

`func (o *TeamHeroesResponse) HasHeroId() bool`

HasHeroId returns a boolean if a field has been set.

### GetName

`func (o *TeamHeroesResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamHeroesResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamHeroesResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamHeroesResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGamesPlayed

`func (o *TeamHeroesResponse) GetGamesPlayed() int32`

GetGamesPlayed returns the GamesPlayed field if non-nil, zero value otherwise.

### GetGamesPlayedOk

`func (o *TeamHeroesResponse) GetGamesPlayedOk() (*int32, bool)`

GetGamesPlayedOk returns a tuple with the GamesPlayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGamesPlayed

`func (o *TeamHeroesResponse) SetGamesPlayed(v int32)`

SetGamesPlayed sets GamesPlayed field to given value.

### HasGamesPlayed

`func (o *TeamHeroesResponse) HasGamesPlayed() bool`

HasGamesPlayed returns a boolean if a field has been set.

### GetWins

`func (o *TeamHeroesResponse) GetWins() int32`

GetWins returns the Wins field if non-nil, zero value otherwise.

### GetWinsOk

`func (o *TeamHeroesResponse) GetWinsOk() (*int32, bool)`

GetWinsOk returns a tuple with the Wins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWins

`func (o *TeamHeroesResponse) SetWins(v int32)`

SetWins sets Wins field to given value.

### HasWins

`func (o *TeamHeroesResponse) HasWins() bool`

HasWins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


