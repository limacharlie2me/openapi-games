# ScenarioLaneRolesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeroId** | Pointer to **int32** | The ID value of the hero played | [optional] 
**LaneRole** | Pointer to **int32** | The hero&#39;s lane role | [optional] 
**Time** | Pointer to **int32** | Maximum game length in seconds | [optional] 
**Games** | Pointer to **string** | The number of games where the hero played in this lane role | [optional] 
**Wins** | Pointer to **string** | The number of games won where the hero played in this lane role | [optional] 

## Methods

### NewScenarioLaneRolesResponse

`func NewScenarioLaneRolesResponse() *ScenarioLaneRolesResponse`

NewScenarioLaneRolesResponse instantiates a new ScenarioLaneRolesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScenarioLaneRolesResponseWithDefaults

`func NewScenarioLaneRolesResponseWithDefaults() *ScenarioLaneRolesResponse`

NewScenarioLaneRolesResponseWithDefaults instantiates a new ScenarioLaneRolesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeroId

`func (o *ScenarioLaneRolesResponse) GetHeroId() int32`

GetHeroId returns the HeroId field if non-nil, zero value otherwise.

### GetHeroIdOk

`func (o *ScenarioLaneRolesResponse) GetHeroIdOk() (*int32, bool)`

GetHeroIdOk returns a tuple with the HeroId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroId

`func (o *ScenarioLaneRolesResponse) SetHeroId(v int32)`

SetHeroId sets HeroId field to given value.

### HasHeroId

`func (o *ScenarioLaneRolesResponse) HasHeroId() bool`

HasHeroId returns a boolean if a field has been set.

### GetLaneRole

`func (o *ScenarioLaneRolesResponse) GetLaneRole() int32`

GetLaneRole returns the LaneRole field if non-nil, zero value otherwise.

### GetLaneRoleOk

`func (o *ScenarioLaneRolesResponse) GetLaneRoleOk() (*int32, bool)`

GetLaneRoleOk returns a tuple with the LaneRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaneRole

`func (o *ScenarioLaneRolesResponse) SetLaneRole(v int32)`

SetLaneRole sets LaneRole field to given value.

### HasLaneRole

`func (o *ScenarioLaneRolesResponse) HasLaneRole() bool`

HasLaneRole returns a boolean if a field has been set.

### GetTime

`func (o *ScenarioLaneRolesResponse) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ScenarioLaneRolesResponse) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ScenarioLaneRolesResponse) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *ScenarioLaneRolesResponse) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetGames

`func (o *ScenarioLaneRolesResponse) GetGames() string`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *ScenarioLaneRolesResponse) GetGamesOk() (*string, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *ScenarioLaneRolesResponse) SetGames(v string)`

SetGames sets Games field to given value.

### HasGames

`func (o *ScenarioLaneRolesResponse) HasGames() bool`

HasGames returns a boolean if a field has been set.

### GetWins

`func (o *ScenarioLaneRolesResponse) GetWins() string`

GetWins returns the Wins field if non-nil, zero value otherwise.

### GetWinsOk

`func (o *ScenarioLaneRolesResponse) GetWinsOk() (*string, bool)`

GetWinsOk returns a tuple with the Wins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWins

`func (o *ScenarioLaneRolesResponse) SetWins(v string)`

SetWins sets Wins field to given value.

### HasWins

`func (o *ScenarioLaneRolesResponse) HasWins() bool`

HasWins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


