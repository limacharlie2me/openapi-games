# MatchResponseDraftTimingsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to **int32** | order | [optional] 
**Pick** | Pointer to **bool** | pick | [optional] 
**ActiveTeam** | Pointer to **int32** | active_team | [optional] 
**HeroId** | Pointer to **int32** | The ID value of the hero played | [optional] 
**PlayerSlot** | Pointer to **NullableInt32** | Which slot the player is in. 0-127 are Radiant, 128-255 are Dire | [optional] 
**ExtraTime** | Pointer to **int32** | extra_time | [optional] 
**TotalTimeTaken** | Pointer to **int32** | total_time_taken | [optional] 

## Methods

### NewMatchResponseDraftTimingsInner

`func NewMatchResponseDraftTimingsInner() *MatchResponseDraftTimingsInner`

NewMatchResponseDraftTimingsInner instantiates a new MatchResponseDraftTimingsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchResponseDraftTimingsInnerWithDefaults

`func NewMatchResponseDraftTimingsInnerWithDefaults() *MatchResponseDraftTimingsInner`

NewMatchResponseDraftTimingsInnerWithDefaults instantiates a new MatchResponseDraftTimingsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *MatchResponseDraftTimingsInner) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MatchResponseDraftTimingsInner) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MatchResponseDraftTimingsInner) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *MatchResponseDraftTimingsInner) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPick

`func (o *MatchResponseDraftTimingsInner) GetPick() bool`

GetPick returns the Pick field if non-nil, zero value otherwise.

### GetPickOk

`func (o *MatchResponseDraftTimingsInner) GetPickOk() (*bool, bool)`

GetPickOk returns a tuple with the Pick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPick

`func (o *MatchResponseDraftTimingsInner) SetPick(v bool)`

SetPick sets Pick field to given value.

### HasPick

`func (o *MatchResponseDraftTimingsInner) HasPick() bool`

HasPick returns a boolean if a field has been set.

### GetActiveTeam

`func (o *MatchResponseDraftTimingsInner) GetActiveTeam() int32`

GetActiveTeam returns the ActiveTeam field if non-nil, zero value otherwise.

### GetActiveTeamOk

`func (o *MatchResponseDraftTimingsInner) GetActiveTeamOk() (*int32, bool)`

GetActiveTeamOk returns a tuple with the ActiveTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTeam

`func (o *MatchResponseDraftTimingsInner) SetActiveTeam(v int32)`

SetActiveTeam sets ActiveTeam field to given value.

### HasActiveTeam

`func (o *MatchResponseDraftTimingsInner) HasActiveTeam() bool`

HasActiveTeam returns a boolean if a field has been set.

### GetHeroId

`func (o *MatchResponseDraftTimingsInner) GetHeroId() int32`

GetHeroId returns the HeroId field if non-nil, zero value otherwise.

### GetHeroIdOk

`func (o *MatchResponseDraftTimingsInner) GetHeroIdOk() (*int32, bool)`

GetHeroIdOk returns a tuple with the HeroId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroId

`func (o *MatchResponseDraftTimingsInner) SetHeroId(v int32)`

SetHeroId sets HeroId field to given value.

### HasHeroId

`func (o *MatchResponseDraftTimingsInner) HasHeroId() bool`

HasHeroId returns a boolean if a field has been set.

### GetPlayerSlot

`func (o *MatchResponseDraftTimingsInner) GetPlayerSlot() int32`

GetPlayerSlot returns the PlayerSlot field if non-nil, zero value otherwise.

### GetPlayerSlotOk

`func (o *MatchResponseDraftTimingsInner) GetPlayerSlotOk() (*int32, bool)`

GetPlayerSlotOk returns a tuple with the PlayerSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerSlot

`func (o *MatchResponseDraftTimingsInner) SetPlayerSlot(v int32)`

SetPlayerSlot sets PlayerSlot field to given value.

### HasPlayerSlot

`func (o *MatchResponseDraftTimingsInner) HasPlayerSlot() bool`

HasPlayerSlot returns a boolean if a field has been set.

### SetPlayerSlotNil

`func (o *MatchResponseDraftTimingsInner) SetPlayerSlotNil(b bool)`

 SetPlayerSlotNil sets the value for PlayerSlot to be an explicit nil

### UnsetPlayerSlot
`func (o *MatchResponseDraftTimingsInner) UnsetPlayerSlot()`

UnsetPlayerSlot ensures that no value is present for PlayerSlot, not even an explicit nil
### GetExtraTime

`func (o *MatchResponseDraftTimingsInner) GetExtraTime() int32`

GetExtraTime returns the ExtraTime field if non-nil, zero value otherwise.

### GetExtraTimeOk

`func (o *MatchResponseDraftTimingsInner) GetExtraTimeOk() (*int32, bool)`

GetExtraTimeOk returns a tuple with the ExtraTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraTime

`func (o *MatchResponseDraftTimingsInner) SetExtraTime(v int32)`

SetExtraTime sets ExtraTime field to given value.

### HasExtraTime

`func (o *MatchResponseDraftTimingsInner) HasExtraTime() bool`

HasExtraTime returns a boolean if a field has been set.

### GetTotalTimeTaken

`func (o *MatchResponseDraftTimingsInner) GetTotalTimeTaken() int32`

GetTotalTimeTaken returns the TotalTimeTaken field if non-nil, zero value otherwise.

### GetTotalTimeTakenOk

`func (o *MatchResponseDraftTimingsInner) GetTotalTimeTakenOk() (*int32, bool)`

GetTotalTimeTakenOk returns a tuple with the TotalTimeTaken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimeTaken

`func (o *MatchResponseDraftTimingsInner) SetTotalTimeTaken(v int32)`

SetTotalTimeTaken sets TotalTimeTaken field to given value.

### HasTotalTimeTaken

`func (o *MatchResponseDraftTimingsInner) HasTotalTimeTaken() bool`

HasTotalTimeTaken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


