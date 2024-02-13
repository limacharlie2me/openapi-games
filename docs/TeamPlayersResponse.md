# TeamPlayersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** | The player account ID | [optional] 
**Name** | Pointer to **NullableString** | name | [optional] 
**GamesPlayed** | Pointer to **int32** | Number of games played | [optional] 
**Wins** | Pointer to **int32** | Number of wins | [optional] 
**IsCurrentTeamMember** | Pointer to **bool** | If this player is on the current roster | [optional] 

## Methods

### NewTeamPlayersResponse

`func NewTeamPlayersResponse() *TeamPlayersResponse`

NewTeamPlayersResponse instantiates a new TeamPlayersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamPlayersResponseWithDefaults

`func NewTeamPlayersResponseWithDefaults() *TeamPlayersResponse`

NewTeamPlayersResponseWithDefaults instantiates a new TeamPlayersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *TeamPlayersResponse) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TeamPlayersResponse) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TeamPlayersResponse) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TeamPlayersResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *TeamPlayersResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamPlayersResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamPlayersResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamPlayersResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TeamPlayersResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TeamPlayersResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetGamesPlayed

`func (o *TeamPlayersResponse) GetGamesPlayed() int32`

GetGamesPlayed returns the GamesPlayed field if non-nil, zero value otherwise.

### GetGamesPlayedOk

`func (o *TeamPlayersResponse) GetGamesPlayedOk() (*int32, bool)`

GetGamesPlayedOk returns a tuple with the GamesPlayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGamesPlayed

`func (o *TeamPlayersResponse) SetGamesPlayed(v int32)`

SetGamesPlayed sets GamesPlayed field to given value.

### HasGamesPlayed

`func (o *TeamPlayersResponse) HasGamesPlayed() bool`

HasGamesPlayed returns a boolean if a field has been set.

### GetWins

`func (o *TeamPlayersResponse) GetWins() int32`

GetWins returns the Wins field if non-nil, zero value otherwise.

### GetWinsOk

`func (o *TeamPlayersResponse) GetWinsOk() (*int32, bool)`

GetWinsOk returns a tuple with the Wins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWins

`func (o *TeamPlayersResponse) SetWins(v int32)`

SetWins sets Wins field to given value.

### HasWins

`func (o *TeamPlayersResponse) HasWins() bool`

HasWins returns a boolean if a field has been set.

### GetIsCurrentTeamMember

`func (o *TeamPlayersResponse) GetIsCurrentTeamMember() bool`

GetIsCurrentTeamMember returns the IsCurrentTeamMember field if non-nil, zero value otherwise.

### GetIsCurrentTeamMemberOk

`func (o *TeamPlayersResponse) GetIsCurrentTeamMemberOk() (*bool, bool)`

GetIsCurrentTeamMemberOk returns a tuple with the IsCurrentTeamMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCurrentTeamMember

`func (o *TeamPlayersResponse) SetIsCurrentTeamMember(v bool)`

SetIsCurrentTeamMember sets IsCurrentTeamMember field to given value.

### HasIsCurrentTeamMember

`func (o *TeamPlayersResponse) HasIsCurrentTeamMember() bool`

HasIsCurrentTeamMember returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


