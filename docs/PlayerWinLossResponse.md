# PlayerWinLossResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Win** | Pointer to **int32** | Number of wins | [optional] 
**Lose** | Pointer to **int32** | Number of loses | [optional] 

## Methods

### NewPlayerWinLossResponse

`func NewPlayerWinLossResponse() *PlayerWinLossResponse`

NewPlayerWinLossResponse instantiates a new PlayerWinLossResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerWinLossResponseWithDefaults

`func NewPlayerWinLossResponseWithDefaults() *PlayerWinLossResponse`

NewPlayerWinLossResponseWithDefaults instantiates a new PlayerWinLossResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWin

`func (o *PlayerWinLossResponse) GetWin() int32`

GetWin returns the Win field if non-nil, zero value otherwise.

### GetWinOk

`func (o *PlayerWinLossResponse) GetWinOk() (*int32, bool)`

GetWinOk returns a tuple with the Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWin

`func (o *PlayerWinLossResponse) SetWin(v int32)`

SetWin sets Win field to given value.

### HasWin

`func (o *PlayerWinLossResponse) HasWin() bool`

HasWin returns a boolean if a field has been set.

### GetLose

`func (o *PlayerWinLossResponse) GetLose() int32`

GetLose returns the Lose field if non-nil, zero value otherwise.

### GetLoseOk

`func (o *PlayerWinLossResponse) GetLoseOk() (*int32, bool)`

GetLoseOk returns a tuple with the Lose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLose

`func (o *PlayerWinLossResponse) SetLose(v int32)`

SetLose sets Lose field to given value.

### HasLose

`func (o *PlayerWinLossResponse) HasLose() bool`

HasLose returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


