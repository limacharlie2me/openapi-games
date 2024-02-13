# MatchResponsePlayersInnerConnectionLogInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **int32** | Game time in seconds the event ocurred | [optional] 
**Event** | Pointer to **string** | Event that occurred | [optional] 
**PlayerSlot** | Pointer to **NullableInt32** | Which slot the player is in. 0-127 are Radiant, 128-255 are Dire | [optional] 

## Methods

### NewMatchResponsePlayersInnerConnectionLogInner

`func NewMatchResponsePlayersInnerConnectionLogInner() *MatchResponsePlayersInnerConnectionLogInner`

NewMatchResponsePlayersInnerConnectionLogInner instantiates a new MatchResponsePlayersInnerConnectionLogInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchResponsePlayersInnerConnectionLogInnerWithDefaults

`func NewMatchResponsePlayersInnerConnectionLogInnerWithDefaults() *MatchResponsePlayersInnerConnectionLogInner`

NewMatchResponsePlayersInnerConnectionLogInnerWithDefaults instantiates a new MatchResponsePlayersInnerConnectionLogInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *MatchResponsePlayersInnerConnectionLogInner) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MatchResponsePlayersInnerConnectionLogInner) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *MatchResponsePlayersInnerConnectionLogInner) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *MatchResponsePlayersInnerConnectionLogInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetEvent

`func (o *MatchResponsePlayersInnerConnectionLogInner) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *MatchResponsePlayersInnerConnectionLogInner) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *MatchResponsePlayersInnerConnectionLogInner) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *MatchResponsePlayersInnerConnectionLogInner) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetPlayerSlot

`func (o *MatchResponsePlayersInnerConnectionLogInner) GetPlayerSlot() int32`

GetPlayerSlot returns the PlayerSlot field if non-nil, zero value otherwise.

### GetPlayerSlotOk

`func (o *MatchResponsePlayersInnerConnectionLogInner) GetPlayerSlotOk() (*int32, bool)`

GetPlayerSlotOk returns a tuple with the PlayerSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerSlot

`func (o *MatchResponsePlayersInnerConnectionLogInner) SetPlayerSlot(v int32)`

SetPlayerSlot sets PlayerSlot field to given value.

### HasPlayerSlot

`func (o *MatchResponsePlayersInnerConnectionLogInner) HasPlayerSlot() bool`

HasPlayerSlot returns a boolean if a field has been set.

### SetPlayerSlotNil

`func (o *MatchResponsePlayersInnerConnectionLogInner) SetPlayerSlotNil(b bool)`

 SetPlayerSlotNil sets the value for PlayerSlot to be an explicit nil

### UnsetPlayerSlot
`func (o *MatchResponsePlayersInnerConnectionLogInner) UnsetPlayerSlot()`

UnsetPlayerSlot ensures that no value is present for PlayerSlot, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


