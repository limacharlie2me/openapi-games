# PlayersResponseProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** | The player account ID | [optional] 
**Personaname** | Pointer to **NullableString** | Player&#39;s Steam name | [optional] 
**Name** | Pointer to **NullableString** | name | [optional] 
**Plus** | Pointer to **bool** | Boolean indicating status of current Dota Plus subscription | [optional] 
**Cheese** | Pointer to **NullableInt32** | cheese | [optional] 
**Steamid** | Pointer to **NullableString** | steamid | [optional] 
**Avatar** | Pointer to **NullableString** | avatar | [optional] 
**Avatarmedium** | Pointer to **NullableString** | avatarmedium | [optional] 
**Avatarfull** | Pointer to **NullableString** | avatarfull | [optional] 
**Profileurl** | Pointer to **NullableString** | profileurl | [optional] 
**LastLogin** | Pointer to **NullableString** | last_login | [optional] 
**Loccountrycode** | Pointer to **NullableString** | loccountrycode | [optional] 
**IsContributor** | Pointer to **bool** | Boolean indicating if the user contributed to the development of OpenDota | [optional] [default to false]
**IsSubscriber** | Pointer to **bool** | Boolean indicating if the user subscribed to OpenDota | [optional] [default to false]

## Methods

### NewPlayersResponseProfile

`func NewPlayersResponseProfile() *PlayersResponseProfile`

NewPlayersResponseProfile instantiates a new PlayersResponseProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayersResponseProfileWithDefaults

`func NewPlayersResponseProfileWithDefaults() *PlayersResponseProfile`

NewPlayersResponseProfileWithDefaults instantiates a new PlayersResponseProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *PlayersResponseProfile) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PlayersResponseProfile) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PlayersResponseProfile) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *PlayersResponseProfile) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPersonaname

`func (o *PlayersResponseProfile) GetPersonaname() string`

GetPersonaname returns the Personaname field if non-nil, zero value otherwise.

### GetPersonanameOk

`func (o *PlayersResponseProfile) GetPersonanameOk() (*string, bool)`

GetPersonanameOk returns a tuple with the Personaname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonaname

`func (o *PlayersResponseProfile) SetPersonaname(v string)`

SetPersonaname sets Personaname field to given value.

### HasPersonaname

`func (o *PlayersResponseProfile) HasPersonaname() bool`

HasPersonaname returns a boolean if a field has been set.

### SetPersonanameNil

`func (o *PlayersResponseProfile) SetPersonanameNil(b bool)`

 SetPersonanameNil sets the value for Personaname to be an explicit nil

### UnsetPersonaname
`func (o *PlayersResponseProfile) UnsetPersonaname()`

UnsetPersonaname ensures that no value is present for Personaname, not even an explicit nil
### GetName

`func (o *PlayersResponseProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlayersResponseProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlayersResponseProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlayersResponseProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PlayersResponseProfile) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PlayersResponseProfile) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPlus

`func (o *PlayersResponseProfile) GetPlus() bool`

GetPlus returns the Plus field if non-nil, zero value otherwise.

### GetPlusOk

`func (o *PlayersResponseProfile) GetPlusOk() (*bool, bool)`

GetPlusOk returns a tuple with the Plus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlus

`func (o *PlayersResponseProfile) SetPlus(v bool)`

SetPlus sets Plus field to given value.

### HasPlus

`func (o *PlayersResponseProfile) HasPlus() bool`

HasPlus returns a boolean if a field has been set.

### GetCheese

`func (o *PlayersResponseProfile) GetCheese() int32`

GetCheese returns the Cheese field if non-nil, zero value otherwise.

### GetCheeseOk

`func (o *PlayersResponseProfile) GetCheeseOk() (*int32, bool)`

GetCheeseOk returns a tuple with the Cheese field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheese

`func (o *PlayersResponseProfile) SetCheese(v int32)`

SetCheese sets Cheese field to given value.

### HasCheese

`func (o *PlayersResponseProfile) HasCheese() bool`

HasCheese returns a boolean if a field has been set.

### SetCheeseNil

`func (o *PlayersResponseProfile) SetCheeseNil(b bool)`

 SetCheeseNil sets the value for Cheese to be an explicit nil

### UnsetCheese
`func (o *PlayersResponseProfile) UnsetCheese()`

UnsetCheese ensures that no value is present for Cheese, not even an explicit nil
### GetSteamid

`func (o *PlayersResponseProfile) GetSteamid() string`

GetSteamid returns the Steamid field if non-nil, zero value otherwise.

### GetSteamidOk

`func (o *PlayersResponseProfile) GetSteamidOk() (*string, bool)`

GetSteamidOk returns a tuple with the Steamid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamid

`func (o *PlayersResponseProfile) SetSteamid(v string)`

SetSteamid sets Steamid field to given value.

### HasSteamid

`func (o *PlayersResponseProfile) HasSteamid() bool`

HasSteamid returns a boolean if a field has been set.

### SetSteamidNil

`func (o *PlayersResponseProfile) SetSteamidNil(b bool)`

 SetSteamidNil sets the value for Steamid to be an explicit nil

### UnsetSteamid
`func (o *PlayersResponseProfile) UnsetSteamid()`

UnsetSteamid ensures that no value is present for Steamid, not even an explicit nil
### GetAvatar

`func (o *PlayersResponseProfile) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *PlayersResponseProfile) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *PlayersResponseProfile) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *PlayersResponseProfile) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *PlayersResponseProfile) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *PlayersResponseProfile) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetAvatarmedium

`func (o *PlayersResponseProfile) GetAvatarmedium() string`

GetAvatarmedium returns the Avatarmedium field if non-nil, zero value otherwise.

### GetAvatarmediumOk

`func (o *PlayersResponseProfile) GetAvatarmediumOk() (*string, bool)`

GetAvatarmediumOk returns a tuple with the Avatarmedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarmedium

`func (o *PlayersResponseProfile) SetAvatarmedium(v string)`

SetAvatarmedium sets Avatarmedium field to given value.

### HasAvatarmedium

`func (o *PlayersResponseProfile) HasAvatarmedium() bool`

HasAvatarmedium returns a boolean if a field has been set.

### SetAvatarmediumNil

`func (o *PlayersResponseProfile) SetAvatarmediumNil(b bool)`

 SetAvatarmediumNil sets the value for Avatarmedium to be an explicit nil

### UnsetAvatarmedium
`func (o *PlayersResponseProfile) UnsetAvatarmedium()`

UnsetAvatarmedium ensures that no value is present for Avatarmedium, not even an explicit nil
### GetAvatarfull

`func (o *PlayersResponseProfile) GetAvatarfull() string`

GetAvatarfull returns the Avatarfull field if non-nil, zero value otherwise.

### GetAvatarfullOk

`func (o *PlayersResponseProfile) GetAvatarfullOk() (*string, bool)`

GetAvatarfullOk returns a tuple with the Avatarfull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarfull

`func (o *PlayersResponseProfile) SetAvatarfull(v string)`

SetAvatarfull sets Avatarfull field to given value.

### HasAvatarfull

`func (o *PlayersResponseProfile) HasAvatarfull() bool`

HasAvatarfull returns a boolean if a field has been set.

### SetAvatarfullNil

`func (o *PlayersResponseProfile) SetAvatarfullNil(b bool)`

 SetAvatarfullNil sets the value for Avatarfull to be an explicit nil

### UnsetAvatarfull
`func (o *PlayersResponseProfile) UnsetAvatarfull()`

UnsetAvatarfull ensures that no value is present for Avatarfull, not even an explicit nil
### GetProfileurl

`func (o *PlayersResponseProfile) GetProfileurl() string`

GetProfileurl returns the Profileurl field if non-nil, zero value otherwise.

### GetProfileurlOk

`func (o *PlayersResponseProfile) GetProfileurlOk() (*string, bool)`

GetProfileurlOk returns a tuple with the Profileurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileurl

`func (o *PlayersResponseProfile) SetProfileurl(v string)`

SetProfileurl sets Profileurl field to given value.

### HasProfileurl

`func (o *PlayersResponseProfile) HasProfileurl() bool`

HasProfileurl returns a boolean if a field has been set.

### SetProfileurlNil

`func (o *PlayersResponseProfile) SetProfileurlNil(b bool)`

 SetProfileurlNil sets the value for Profileurl to be an explicit nil

### UnsetProfileurl
`func (o *PlayersResponseProfile) UnsetProfileurl()`

UnsetProfileurl ensures that no value is present for Profileurl, not even an explicit nil
### GetLastLogin

`func (o *PlayersResponseProfile) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *PlayersResponseProfile) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *PlayersResponseProfile) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *PlayersResponseProfile) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *PlayersResponseProfile) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *PlayersResponseProfile) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetLoccountrycode

`func (o *PlayersResponseProfile) GetLoccountrycode() string`

GetLoccountrycode returns the Loccountrycode field if non-nil, zero value otherwise.

### GetLoccountrycodeOk

`func (o *PlayersResponseProfile) GetLoccountrycodeOk() (*string, bool)`

GetLoccountrycodeOk returns a tuple with the Loccountrycode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoccountrycode

`func (o *PlayersResponseProfile) SetLoccountrycode(v string)`

SetLoccountrycode sets Loccountrycode field to given value.

### HasLoccountrycode

`func (o *PlayersResponseProfile) HasLoccountrycode() bool`

HasLoccountrycode returns a boolean if a field has been set.

### SetLoccountrycodeNil

`func (o *PlayersResponseProfile) SetLoccountrycodeNil(b bool)`

 SetLoccountrycodeNil sets the value for Loccountrycode to be an explicit nil

### UnsetLoccountrycode
`func (o *PlayersResponseProfile) UnsetLoccountrycode()`

UnsetLoccountrycode ensures that no value is present for Loccountrycode, not even an explicit nil
### GetIsContributor

`func (o *PlayersResponseProfile) GetIsContributor() bool`

GetIsContributor returns the IsContributor field if non-nil, zero value otherwise.

### GetIsContributorOk

`func (o *PlayersResponseProfile) GetIsContributorOk() (*bool, bool)`

GetIsContributorOk returns a tuple with the IsContributor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContributor

`func (o *PlayersResponseProfile) SetIsContributor(v bool)`

SetIsContributor sets IsContributor field to given value.

### HasIsContributor

`func (o *PlayersResponseProfile) HasIsContributor() bool`

HasIsContributor returns a boolean if a field has been set.

### GetIsSubscriber

`func (o *PlayersResponseProfile) GetIsSubscriber() bool`

GetIsSubscriber returns the IsSubscriber field if non-nil, zero value otherwise.

### GetIsSubscriberOk

`func (o *PlayersResponseProfile) GetIsSubscriberOk() (*bool, bool)`

GetIsSubscriberOk returns a tuple with the IsSubscriber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscriber

`func (o *PlayersResponseProfile) SetIsSubscriber(v bool)`

SetIsSubscriber sets IsSubscriber field to given value.

### HasIsSubscriber

`func (o *PlayersResponseProfile) HasIsSubscriber() bool`

HasIsSubscriber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


