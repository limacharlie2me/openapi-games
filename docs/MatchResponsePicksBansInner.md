# MatchResponsePicksBansInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPick** | Pointer to **bool** | Boolean indicating whether the choice is a pick or a ban | [optional] 
**HeroId** | Pointer to **int32** | The ID value of the hero played | [optional] 
**Team** | Pointer to **int32** | The team that picked or banned the hero | [optional] 
**Order** | Pointer to **int32** | The order of the pick or ban | [optional] 

## Methods

### NewMatchResponsePicksBansInner

`func NewMatchResponsePicksBansInner() *MatchResponsePicksBansInner`

NewMatchResponsePicksBansInner instantiates a new MatchResponsePicksBansInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchResponsePicksBansInnerWithDefaults

`func NewMatchResponsePicksBansInnerWithDefaults() *MatchResponsePicksBansInner`

NewMatchResponsePicksBansInnerWithDefaults instantiates a new MatchResponsePicksBansInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPick

`func (o *MatchResponsePicksBansInner) GetIsPick() bool`

GetIsPick returns the IsPick field if non-nil, zero value otherwise.

### GetIsPickOk

`func (o *MatchResponsePicksBansInner) GetIsPickOk() (*bool, bool)`

GetIsPickOk returns a tuple with the IsPick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPick

`func (o *MatchResponsePicksBansInner) SetIsPick(v bool)`

SetIsPick sets IsPick field to given value.

### HasIsPick

`func (o *MatchResponsePicksBansInner) HasIsPick() bool`

HasIsPick returns a boolean if a field has been set.

### GetHeroId

`func (o *MatchResponsePicksBansInner) GetHeroId() int32`

GetHeroId returns the HeroId field if non-nil, zero value otherwise.

### GetHeroIdOk

`func (o *MatchResponsePicksBansInner) GetHeroIdOk() (*int32, bool)`

GetHeroIdOk returns a tuple with the HeroId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroId

`func (o *MatchResponsePicksBansInner) SetHeroId(v int32)`

SetHeroId sets HeroId field to given value.

### HasHeroId

`func (o *MatchResponsePicksBansInner) HasHeroId() bool`

HasHeroId returns a boolean if a field has been set.

### GetTeam

`func (o *MatchResponsePicksBansInner) GetTeam() int32`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *MatchResponsePicksBansInner) GetTeamOk() (*int32, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *MatchResponsePicksBansInner) SetTeam(v int32)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *MatchResponsePicksBansInner) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetOrder

`func (o *MatchResponsePicksBansInner) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MatchResponsePicksBansInner) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MatchResponsePicksBansInner) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *MatchResponsePicksBansInner) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


