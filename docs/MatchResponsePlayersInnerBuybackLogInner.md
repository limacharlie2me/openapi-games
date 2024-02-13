# MatchResponsePlayersInnerBuybackLogInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **int32** | Time in seconds the buyback occurred | [optional] 
**Slot** | Pointer to **int32** | slot | [optional] 
**PlayerSlot** | Pointer to **NullableInt32** | Which slot the player is in. 0-127 are Radiant, 128-255 are Dire | [optional] 

## Methods

### NewMatchResponsePlayersInnerBuybackLogInner

`func NewMatchResponsePlayersInnerBuybackLogInner() *MatchResponsePlayersInnerBuybackLogInner`

NewMatchResponsePlayersInnerBuybackLogInner instantiates a new MatchResponsePlayersInnerBuybackLogInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchResponsePlayersInnerBuybackLogInnerWithDefaults

`func NewMatchResponsePlayersInnerBuybackLogInnerWithDefaults() *MatchResponsePlayersInnerBuybackLogInner`

NewMatchResponsePlayersInnerBuybackLogInnerWithDefaults instantiates a new MatchResponsePlayersInnerBuybackLogInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *MatchResponsePlayersInnerBuybackLogInner) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MatchResponsePlayersInnerBuybackLogInner) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *MatchResponsePlayersInnerBuybackLogInner) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *MatchResponsePlayersInnerBuybackLogInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetSlot

`func (o *MatchResponsePlayersInnerBuybackLogInner) GetSlot() int32`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *MatchResponsePlayersInnerBuybackLogInner) GetSlotOk() (*int32, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *MatchResponsePlayersInnerBuybackLogInner) SetSlot(v int32)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *MatchResponsePlayersInnerBuybackLogInner) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetPlayerSlot

`func (o *MatchResponsePlayersInnerBuybackLogInner) GetPlayerSlot() int32`

GetPlayerSlot returns the PlayerSlot field if non-nil, zero value otherwise.

### GetPlayerSlotOk

`func (o *MatchResponsePlayersInnerBuybackLogInner) GetPlayerSlotOk() (*int32, bool)`

GetPlayerSlotOk returns a tuple with the PlayerSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerSlot

`func (o *MatchResponsePlayersInnerBuybackLogInner) SetPlayerSlot(v int32)`

SetPlayerSlot sets PlayerSlot field to given value.

### HasPlayerSlot

`func (o *MatchResponsePlayersInnerBuybackLogInner) HasPlayerSlot() bool`

HasPlayerSlot returns a boolean if a field has been set.

### SetPlayerSlotNil

`func (o *MatchResponsePlayersInnerBuybackLogInner) SetPlayerSlotNil(b bool)`

 SetPlayerSlotNil sets the value for PlayerSlot to be an explicit nil

### UnsetPlayerSlot
`func (o *MatchResponsePlayersInnerBuybackLogInner) UnsetPlayerSlot()`

UnsetPlayerSlot ensures that no value is present for PlayerSlot, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


