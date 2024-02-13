# HeroDurationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationBin** | Pointer to **string** | Lower bound of number of seconds the match lasted | [optional] 
**GamesPlayed** | Pointer to **int32** | Number of games played | [optional] 
**Wins** | Pointer to **int32** | Number of wins | [optional] 

## Methods

### NewHeroDurationsResponse

`func NewHeroDurationsResponse() *HeroDurationsResponse`

NewHeroDurationsResponse instantiates a new HeroDurationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeroDurationsResponseWithDefaults

`func NewHeroDurationsResponseWithDefaults() *HeroDurationsResponse`

NewHeroDurationsResponseWithDefaults instantiates a new HeroDurationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationBin

`func (o *HeroDurationsResponse) GetDurationBin() string`

GetDurationBin returns the DurationBin field if non-nil, zero value otherwise.

### GetDurationBinOk

`func (o *HeroDurationsResponse) GetDurationBinOk() (*string, bool)`

GetDurationBinOk returns a tuple with the DurationBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationBin

`func (o *HeroDurationsResponse) SetDurationBin(v string)`

SetDurationBin sets DurationBin field to given value.

### HasDurationBin

`func (o *HeroDurationsResponse) HasDurationBin() bool`

HasDurationBin returns a boolean if a field has been set.

### GetGamesPlayed

`func (o *HeroDurationsResponse) GetGamesPlayed() int32`

GetGamesPlayed returns the GamesPlayed field if non-nil, zero value otherwise.

### GetGamesPlayedOk

`func (o *HeroDurationsResponse) GetGamesPlayedOk() (*int32, bool)`

GetGamesPlayedOk returns a tuple with the GamesPlayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGamesPlayed

`func (o *HeroDurationsResponse) SetGamesPlayed(v int32)`

SetGamesPlayed sets GamesPlayed field to given value.

### HasGamesPlayed

`func (o *HeroDurationsResponse) HasGamesPlayed() bool`

HasGamesPlayed returns a boolean if a field has been set.

### GetWins

`func (o *HeroDurationsResponse) GetWins() int32`

GetWins returns the Wins field if non-nil, zero value otherwise.

### GetWinsOk

`func (o *HeroDurationsResponse) GetWinsOk() (*int32, bool)`

GetWinsOk returns a tuple with the Wins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWins

`func (o *HeroDurationsResponse) SetWins(v int32)`

SetWins sets Wins field to given value.

### HasWins

`func (o *HeroDurationsResponse) HasWins() bool`

HasWins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


