# ScenarioItemTimingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeroId** | Pointer to **int32** | The ID value of the hero played | [optional] 
**Item** | Pointer to **string** | Purchased item | [optional] 
**Time** | Pointer to **int32** | Ingame time in seconds before the item was purchased | [optional] 
**Games** | Pointer to **string** | The number of games where the hero bought this item before this time | [optional] 
**Wins** | Pointer to **string** | The number of games won where the hero bought this item before this time | [optional] 

## Methods

### NewScenarioItemTimingsResponse

`func NewScenarioItemTimingsResponse() *ScenarioItemTimingsResponse`

NewScenarioItemTimingsResponse instantiates a new ScenarioItemTimingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScenarioItemTimingsResponseWithDefaults

`func NewScenarioItemTimingsResponseWithDefaults() *ScenarioItemTimingsResponse`

NewScenarioItemTimingsResponseWithDefaults instantiates a new ScenarioItemTimingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeroId

`func (o *ScenarioItemTimingsResponse) GetHeroId() int32`

GetHeroId returns the HeroId field if non-nil, zero value otherwise.

### GetHeroIdOk

`func (o *ScenarioItemTimingsResponse) GetHeroIdOk() (*int32, bool)`

GetHeroIdOk returns a tuple with the HeroId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroId

`func (o *ScenarioItemTimingsResponse) SetHeroId(v int32)`

SetHeroId sets HeroId field to given value.

### HasHeroId

`func (o *ScenarioItemTimingsResponse) HasHeroId() bool`

HasHeroId returns a boolean if a field has been set.

### GetItem

`func (o *ScenarioItemTimingsResponse) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ScenarioItemTimingsResponse) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ScenarioItemTimingsResponse) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *ScenarioItemTimingsResponse) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetTime

`func (o *ScenarioItemTimingsResponse) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ScenarioItemTimingsResponse) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ScenarioItemTimingsResponse) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *ScenarioItemTimingsResponse) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetGames

`func (o *ScenarioItemTimingsResponse) GetGames() string`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *ScenarioItemTimingsResponse) GetGamesOk() (*string, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *ScenarioItemTimingsResponse) SetGames(v string)`

SetGames sets Games field to given value.

### HasGames

`func (o *ScenarioItemTimingsResponse) HasGames() bool`

HasGames returns a boolean if a field has been set.

### GetWins

`func (o *ScenarioItemTimingsResponse) GetWins() string`

GetWins returns the Wins field if non-nil, zero value otherwise.

### GetWinsOk

`func (o *ScenarioItemTimingsResponse) GetWinsOk() (*string, bool)`

GetWinsOk returns a tuple with the Wins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWins

`func (o *ScenarioItemTimingsResponse) SetWins(v string)`

SetWins sets Wins field to given value.

### HasWins

`func (o *ScenarioItemTimingsResponse) HasWins() bool`

HasWins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


