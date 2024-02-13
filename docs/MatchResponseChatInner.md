# MatchResponseChatInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **int32** | Time in seconds at which the message was said | [optional] 
**Unit** | Pointer to **string** | Name of the player who sent the message | [optional] 
**Key** | Pointer to **string** | The message the player sent | [optional] 
**Slot** | Pointer to **int32** | slot | [optional] 
**PlayerSlot** | Pointer to **NullableInt32** | Which slot the player is in. 0-127 are Radiant, 128-255 are Dire | [optional] 

## Methods

### NewMatchResponseChatInner

`func NewMatchResponseChatInner() *MatchResponseChatInner`

NewMatchResponseChatInner instantiates a new MatchResponseChatInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchResponseChatInnerWithDefaults

`func NewMatchResponseChatInnerWithDefaults() *MatchResponseChatInner`

NewMatchResponseChatInnerWithDefaults instantiates a new MatchResponseChatInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *MatchResponseChatInner) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MatchResponseChatInner) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *MatchResponseChatInner) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *MatchResponseChatInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetUnit

`func (o *MatchResponseChatInner) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MatchResponseChatInner) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MatchResponseChatInner) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MatchResponseChatInner) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetKey

`func (o *MatchResponseChatInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MatchResponseChatInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MatchResponseChatInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MatchResponseChatInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSlot

`func (o *MatchResponseChatInner) GetSlot() int32`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *MatchResponseChatInner) GetSlotOk() (*int32, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *MatchResponseChatInner) SetSlot(v int32)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *MatchResponseChatInner) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetPlayerSlot

`func (o *MatchResponseChatInner) GetPlayerSlot() int32`

GetPlayerSlot returns the PlayerSlot field if non-nil, zero value otherwise.

### GetPlayerSlotOk

`func (o *MatchResponseChatInner) GetPlayerSlotOk() (*int32, bool)`

GetPlayerSlotOk returns a tuple with the PlayerSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerSlot

`func (o *MatchResponseChatInner) SetPlayerSlot(v int32)`

SetPlayerSlot sets PlayerSlot field to given value.

### HasPlayerSlot

`func (o *MatchResponseChatInner) HasPlayerSlot() bool`

HasPlayerSlot returns a boolean if a field has been set.

### SetPlayerSlotNil

`func (o *MatchResponseChatInner) SetPlayerSlotNil(b bool)`

 SetPlayerSlotNil sets the value for PlayerSlot to be an explicit nil

### UnsetPlayerSlot
`func (o *MatchResponseChatInner) UnsetPlayerSlot()`

UnsetPlayerSlot ensures that no value is present for PlayerSlot, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


