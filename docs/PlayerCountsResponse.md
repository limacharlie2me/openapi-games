# PlayerCountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LeaverStatus** | Pointer to **map[string]interface{}** | Integer describing whether or not the player left the game. 0: didn&#39;t leave. 1: left safely. 2+: Abandoned | [optional] 
**GameMode** | Pointer to **map[string]interface{}** | Integer corresponding to game mode played. List of constants can be found here: https://github.com/odota/dotaconstants/blob/master/json/game_mode.json | [optional] 
**LobbyType** | Pointer to **map[string]interface{}** | Integer corresponding to lobby type of match. List of constants can be found here: https://github.com/odota/dotaconstants/blob/master/json/lobby_type.json | [optional] 
**LaneRole** | Pointer to **map[string]interface{}** | lane_role | [optional] 
**Region** | Pointer to **map[string]interface{}** | Integer corresponding to the region the game was played on | [optional] 
**Patch** | Pointer to **map[string]interface{}** | Patch ID, from dotaconstants | [optional] 

## Methods

### NewPlayerCountsResponse

`func NewPlayerCountsResponse() *PlayerCountsResponse`

NewPlayerCountsResponse instantiates a new PlayerCountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerCountsResponseWithDefaults

`func NewPlayerCountsResponseWithDefaults() *PlayerCountsResponse`

NewPlayerCountsResponseWithDefaults instantiates a new PlayerCountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeaverStatus

`func (o *PlayerCountsResponse) GetLeaverStatus() map[string]interface{}`

GetLeaverStatus returns the LeaverStatus field if non-nil, zero value otherwise.

### GetLeaverStatusOk

`func (o *PlayerCountsResponse) GetLeaverStatusOk() (*map[string]interface{}, bool)`

GetLeaverStatusOk returns a tuple with the LeaverStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaverStatus

`func (o *PlayerCountsResponse) SetLeaverStatus(v map[string]interface{})`

SetLeaverStatus sets LeaverStatus field to given value.

### HasLeaverStatus

`func (o *PlayerCountsResponse) HasLeaverStatus() bool`

HasLeaverStatus returns a boolean if a field has been set.

### GetGameMode

`func (o *PlayerCountsResponse) GetGameMode() map[string]interface{}`

GetGameMode returns the GameMode field if non-nil, zero value otherwise.

### GetGameModeOk

`func (o *PlayerCountsResponse) GetGameModeOk() (*map[string]interface{}, bool)`

GetGameModeOk returns a tuple with the GameMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameMode

`func (o *PlayerCountsResponse) SetGameMode(v map[string]interface{})`

SetGameMode sets GameMode field to given value.

### HasGameMode

`func (o *PlayerCountsResponse) HasGameMode() bool`

HasGameMode returns a boolean if a field has been set.

### GetLobbyType

`func (o *PlayerCountsResponse) GetLobbyType() map[string]interface{}`

GetLobbyType returns the LobbyType field if non-nil, zero value otherwise.

### GetLobbyTypeOk

`func (o *PlayerCountsResponse) GetLobbyTypeOk() (*map[string]interface{}, bool)`

GetLobbyTypeOk returns a tuple with the LobbyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLobbyType

`func (o *PlayerCountsResponse) SetLobbyType(v map[string]interface{})`

SetLobbyType sets LobbyType field to given value.

### HasLobbyType

`func (o *PlayerCountsResponse) HasLobbyType() bool`

HasLobbyType returns a boolean if a field has been set.

### GetLaneRole

`func (o *PlayerCountsResponse) GetLaneRole() map[string]interface{}`

GetLaneRole returns the LaneRole field if non-nil, zero value otherwise.

### GetLaneRoleOk

`func (o *PlayerCountsResponse) GetLaneRoleOk() (*map[string]interface{}, bool)`

GetLaneRoleOk returns a tuple with the LaneRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaneRole

`func (o *PlayerCountsResponse) SetLaneRole(v map[string]interface{})`

SetLaneRole sets LaneRole field to given value.

### HasLaneRole

`func (o *PlayerCountsResponse) HasLaneRole() bool`

HasLaneRole returns a boolean if a field has been set.

### GetRegion

`func (o *PlayerCountsResponse) GetRegion() map[string]interface{}`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PlayerCountsResponse) GetRegionOk() (*map[string]interface{}, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PlayerCountsResponse) SetRegion(v map[string]interface{})`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PlayerCountsResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetPatch

`func (o *PlayerCountsResponse) GetPatch() map[string]interface{}`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *PlayerCountsResponse) GetPatchOk() (*map[string]interface{}, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *PlayerCountsResponse) SetPatch(v map[string]interface{})`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *PlayerCountsResponse) HasPatch() bool`

HasPatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


