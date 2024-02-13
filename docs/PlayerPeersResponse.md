# PlayerPeersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** | The player account ID | [optional] 
**LastPlayed** | Pointer to **int32** | last_played | [optional] 
**Win** | Pointer to **int32** | win | [optional] 
**Games** | Pointer to **int32** | games | [optional] 
**WithWin** | Pointer to **int32** | with_win | [optional] 
**WithGames** | Pointer to **int32** | with_games | [optional] 
**AgainstWin** | Pointer to **int32** | against_win | [optional] 
**AgainstGames** | Pointer to **int32** | against_games | [optional] 
**WithGpmSum** | Pointer to **int32** | with_gpm_sum | [optional] 
**WithXpmSum** | Pointer to **int32** | with_xpm_sum | [optional] 
**Personaname** | Pointer to **NullableString** | Player&#39;s Steam name | [optional] 
**Name** | Pointer to **NullableString** | name | [optional] 
**IsContributor** | Pointer to **bool** | is_contributor | [optional] 
**IsSubscriber** | Pointer to **bool** | is_subscriber | [optional] 
**LastLogin** | Pointer to **NullableString** | last_login | [optional] 
**Avatar** | Pointer to **NullableString** | avatar | [optional] 
**Avatarfull** | Pointer to **NullableString** | avatarfull | [optional] 

## Methods

### NewPlayerPeersResponse

`func NewPlayerPeersResponse() *PlayerPeersResponse`

NewPlayerPeersResponse instantiates a new PlayerPeersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerPeersResponseWithDefaults

`func NewPlayerPeersResponseWithDefaults() *PlayerPeersResponse`

NewPlayerPeersResponseWithDefaults instantiates a new PlayerPeersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *PlayerPeersResponse) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PlayerPeersResponse) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PlayerPeersResponse) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *PlayerPeersResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLastPlayed

`func (o *PlayerPeersResponse) GetLastPlayed() int32`

GetLastPlayed returns the LastPlayed field if non-nil, zero value otherwise.

### GetLastPlayedOk

`func (o *PlayerPeersResponse) GetLastPlayedOk() (*int32, bool)`

GetLastPlayedOk returns a tuple with the LastPlayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPlayed

`func (o *PlayerPeersResponse) SetLastPlayed(v int32)`

SetLastPlayed sets LastPlayed field to given value.

### HasLastPlayed

`func (o *PlayerPeersResponse) HasLastPlayed() bool`

HasLastPlayed returns a boolean if a field has been set.

### GetWin

`func (o *PlayerPeersResponse) GetWin() int32`

GetWin returns the Win field if non-nil, zero value otherwise.

### GetWinOk

`func (o *PlayerPeersResponse) GetWinOk() (*int32, bool)`

GetWinOk returns a tuple with the Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWin

`func (o *PlayerPeersResponse) SetWin(v int32)`

SetWin sets Win field to given value.

### HasWin

`func (o *PlayerPeersResponse) HasWin() bool`

HasWin returns a boolean if a field has been set.

### GetGames

`func (o *PlayerPeersResponse) GetGames() int32`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *PlayerPeersResponse) GetGamesOk() (*int32, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *PlayerPeersResponse) SetGames(v int32)`

SetGames sets Games field to given value.

### HasGames

`func (o *PlayerPeersResponse) HasGames() bool`

HasGames returns a boolean if a field has been set.

### GetWithWin

`func (o *PlayerPeersResponse) GetWithWin() int32`

GetWithWin returns the WithWin field if non-nil, zero value otherwise.

### GetWithWinOk

`func (o *PlayerPeersResponse) GetWithWinOk() (*int32, bool)`

GetWithWinOk returns a tuple with the WithWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithWin

`func (o *PlayerPeersResponse) SetWithWin(v int32)`

SetWithWin sets WithWin field to given value.

### HasWithWin

`func (o *PlayerPeersResponse) HasWithWin() bool`

HasWithWin returns a boolean if a field has been set.

### GetWithGames

`func (o *PlayerPeersResponse) GetWithGames() int32`

GetWithGames returns the WithGames field if non-nil, zero value otherwise.

### GetWithGamesOk

`func (o *PlayerPeersResponse) GetWithGamesOk() (*int32, bool)`

GetWithGamesOk returns a tuple with the WithGames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithGames

`func (o *PlayerPeersResponse) SetWithGames(v int32)`

SetWithGames sets WithGames field to given value.

### HasWithGames

`func (o *PlayerPeersResponse) HasWithGames() bool`

HasWithGames returns a boolean if a field has been set.

### GetAgainstWin

`func (o *PlayerPeersResponse) GetAgainstWin() int32`

GetAgainstWin returns the AgainstWin field if non-nil, zero value otherwise.

### GetAgainstWinOk

`func (o *PlayerPeersResponse) GetAgainstWinOk() (*int32, bool)`

GetAgainstWinOk returns a tuple with the AgainstWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgainstWin

`func (o *PlayerPeersResponse) SetAgainstWin(v int32)`

SetAgainstWin sets AgainstWin field to given value.

### HasAgainstWin

`func (o *PlayerPeersResponse) HasAgainstWin() bool`

HasAgainstWin returns a boolean if a field has been set.

### GetAgainstGames

`func (o *PlayerPeersResponse) GetAgainstGames() int32`

GetAgainstGames returns the AgainstGames field if non-nil, zero value otherwise.

### GetAgainstGamesOk

`func (o *PlayerPeersResponse) GetAgainstGamesOk() (*int32, bool)`

GetAgainstGamesOk returns a tuple with the AgainstGames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgainstGames

`func (o *PlayerPeersResponse) SetAgainstGames(v int32)`

SetAgainstGames sets AgainstGames field to given value.

### HasAgainstGames

`func (o *PlayerPeersResponse) HasAgainstGames() bool`

HasAgainstGames returns a boolean if a field has been set.

### GetWithGpmSum

`func (o *PlayerPeersResponse) GetWithGpmSum() int32`

GetWithGpmSum returns the WithGpmSum field if non-nil, zero value otherwise.

### GetWithGpmSumOk

`func (o *PlayerPeersResponse) GetWithGpmSumOk() (*int32, bool)`

GetWithGpmSumOk returns a tuple with the WithGpmSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithGpmSum

`func (o *PlayerPeersResponse) SetWithGpmSum(v int32)`

SetWithGpmSum sets WithGpmSum field to given value.

### HasWithGpmSum

`func (o *PlayerPeersResponse) HasWithGpmSum() bool`

HasWithGpmSum returns a boolean if a field has been set.

### GetWithXpmSum

`func (o *PlayerPeersResponse) GetWithXpmSum() int32`

GetWithXpmSum returns the WithXpmSum field if non-nil, zero value otherwise.

### GetWithXpmSumOk

`func (o *PlayerPeersResponse) GetWithXpmSumOk() (*int32, bool)`

GetWithXpmSumOk returns a tuple with the WithXpmSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithXpmSum

`func (o *PlayerPeersResponse) SetWithXpmSum(v int32)`

SetWithXpmSum sets WithXpmSum field to given value.

### HasWithXpmSum

`func (o *PlayerPeersResponse) HasWithXpmSum() bool`

HasWithXpmSum returns a boolean if a field has been set.

### GetPersonaname

`func (o *PlayerPeersResponse) GetPersonaname() string`

GetPersonaname returns the Personaname field if non-nil, zero value otherwise.

### GetPersonanameOk

`func (o *PlayerPeersResponse) GetPersonanameOk() (*string, bool)`

GetPersonanameOk returns a tuple with the Personaname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonaname

`func (o *PlayerPeersResponse) SetPersonaname(v string)`

SetPersonaname sets Personaname field to given value.

### HasPersonaname

`func (o *PlayerPeersResponse) HasPersonaname() bool`

HasPersonaname returns a boolean if a field has been set.

### SetPersonanameNil

`func (o *PlayerPeersResponse) SetPersonanameNil(b bool)`

 SetPersonanameNil sets the value for Personaname to be an explicit nil

### UnsetPersonaname
`func (o *PlayerPeersResponse) UnsetPersonaname()`

UnsetPersonaname ensures that no value is present for Personaname, not even an explicit nil
### GetName

`func (o *PlayerPeersResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlayerPeersResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlayerPeersResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlayerPeersResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PlayerPeersResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PlayerPeersResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsContributor

`func (o *PlayerPeersResponse) GetIsContributor() bool`

GetIsContributor returns the IsContributor field if non-nil, zero value otherwise.

### GetIsContributorOk

`func (o *PlayerPeersResponse) GetIsContributorOk() (*bool, bool)`

GetIsContributorOk returns a tuple with the IsContributor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContributor

`func (o *PlayerPeersResponse) SetIsContributor(v bool)`

SetIsContributor sets IsContributor field to given value.

### HasIsContributor

`func (o *PlayerPeersResponse) HasIsContributor() bool`

HasIsContributor returns a boolean if a field has been set.

### GetIsSubscriber

`func (o *PlayerPeersResponse) GetIsSubscriber() bool`

GetIsSubscriber returns the IsSubscriber field if non-nil, zero value otherwise.

### GetIsSubscriberOk

`func (o *PlayerPeersResponse) GetIsSubscriberOk() (*bool, bool)`

GetIsSubscriberOk returns a tuple with the IsSubscriber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscriber

`func (o *PlayerPeersResponse) SetIsSubscriber(v bool)`

SetIsSubscriber sets IsSubscriber field to given value.

### HasIsSubscriber

`func (o *PlayerPeersResponse) HasIsSubscriber() bool`

HasIsSubscriber returns a boolean if a field has been set.

### GetLastLogin

`func (o *PlayerPeersResponse) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *PlayerPeersResponse) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *PlayerPeersResponse) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *PlayerPeersResponse) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *PlayerPeersResponse) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *PlayerPeersResponse) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetAvatar

`func (o *PlayerPeersResponse) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *PlayerPeersResponse) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *PlayerPeersResponse) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *PlayerPeersResponse) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *PlayerPeersResponse) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *PlayerPeersResponse) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetAvatarfull

`func (o *PlayerPeersResponse) GetAvatarfull() string`

GetAvatarfull returns the Avatarfull field if non-nil, zero value otherwise.

### GetAvatarfullOk

`func (o *PlayerPeersResponse) GetAvatarfullOk() (*string, bool)`

GetAvatarfullOk returns a tuple with the Avatarfull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarfull

`func (o *PlayerPeersResponse) SetAvatarfull(v string)`

SetAvatarfull sets Avatarfull field to given value.

### HasAvatarfull

`func (o *PlayerPeersResponse) HasAvatarfull() bool`

HasAvatarfull returns a boolean if a field has been set.

### SetAvatarfullNil

`func (o *PlayerPeersResponse) SetAvatarfullNil(b bool)`

 SetAvatarfullNil sets the value for Avatarfull to be an explicit nil

### UnsetAvatarfull
`func (o *PlayerPeersResponse) UnsetAvatarfull()`

UnsetAvatarfull ensures that no value is present for Avatarfull, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


