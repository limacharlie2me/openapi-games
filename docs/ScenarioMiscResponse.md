# ScenarioMiscResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scenario** | Pointer to **string** | The scenario&#39;s name or description | [optional] 
**IsRadiant** | Pointer to **bool** | Boolean indicating whether Radiant executed this scenario | [optional] 
**Region** | Pointer to **int32** | Region the game was played in | [optional] 
**Games** | Pointer to **string** | The number of games where this scenario occurred | [optional] 
**Wins** | Pointer to **string** | The number of games won where this scenario occured | [optional] 

## Methods

### NewScenarioMiscResponse

`func NewScenarioMiscResponse() *ScenarioMiscResponse`

NewScenarioMiscResponse instantiates a new ScenarioMiscResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScenarioMiscResponseWithDefaults

`func NewScenarioMiscResponseWithDefaults() *ScenarioMiscResponse`

NewScenarioMiscResponseWithDefaults instantiates a new ScenarioMiscResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScenario

`func (o *ScenarioMiscResponse) GetScenario() string`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *ScenarioMiscResponse) GetScenarioOk() (*string, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *ScenarioMiscResponse) SetScenario(v string)`

SetScenario sets Scenario field to given value.

### HasScenario

`func (o *ScenarioMiscResponse) HasScenario() bool`

HasScenario returns a boolean if a field has been set.

### GetIsRadiant

`func (o *ScenarioMiscResponse) GetIsRadiant() bool`

GetIsRadiant returns the IsRadiant field if non-nil, zero value otherwise.

### GetIsRadiantOk

`func (o *ScenarioMiscResponse) GetIsRadiantOk() (*bool, bool)`

GetIsRadiantOk returns a tuple with the IsRadiant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRadiant

`func (o *ScenarioMiscResponse) SetIsRadiant(v bool)`

SetIsRadiant sets IsRadiant field to given value.

### HasIsRadiant

`func (o *ScenarioMiscResponse) HasIsRadiant() bool`

HasIsRadiant returns a boolean if a field has been set.

### GetRegion

`func (o *ScenarioMiscResponse) GetRegion() int32`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ScenarioMiscResponse) GetRegionOk() (*int32, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ScenarioMiscResponse) SetRegion(v int32)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ScenarioMiscResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetGames

`func (o *ScenarioMiscResponse) GetGames() string`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *ScenarioMiscResponse) GetGamesOk() (*string, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *ScenarioMiscResponse) SetGames(v string)`

SetGames sets Games field to given value.

### HasGames

`func (o *ScenarioMiscResponse) HasGames() bool`

HasGames returns a boolean if a field has been set.

### GetWins

`func (o *ScenarioMiscResponse) GetWins() string`

GetWins returns the Wins field if non-nil, zero value otherwise.

### GetWinsOk

`func (o *ScenarioMiscResponse) GetWinsOk() (*string, bool)`

GetWinsOk returns a tuple with the Wins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWins

`func (o *ScenarioMiscResponse) SetWins(v string)`

SetWins sets Wins field to given value.

### HasWins

`func (o *ScenarioMiscResponse) HasWins() bool`

HasWins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


