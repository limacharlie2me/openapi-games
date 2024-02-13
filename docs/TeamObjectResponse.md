# TeamObjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamId** | Pointer to **int32** | Team&#39;s identifier | [optional] 
**Rating** | Pointer to **float32** | The Elo rating of the team | [optional] 
**Wins** | Pointer to **int32** | The number of games won by this team | [optional] 
**Losses** | Pointer to **int32** | The number of losses by this team | [optional] 
**LastMatchTime** | Pointer to **int32** | The Unix timestamp of the last match played by this team | [optional] 
**Name** | Pointer to **NullableString** | Team name | [optional] 
**Tag** | Pointer to **string** | The team tag/abbreviation | [optional] 

## Methods

### NewTeamObjectResponse

`func NewTeamObjectResponse() *TeamObjectResponse`

NewTeamObjectResponse instantiates a new TeamObjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamObjectResponseWithDefaults

`func NewTeamObjectResponseWithDefaults() *TeamObjectResponse`

NewTeamObjectResponseWithDefaults instantiates a new TeamObjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamId

`func (o *TeamObjectResponse) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *TeamObjectResponse) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *TeamObjectResponse) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *TeamObjectResponse) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetRating

`func (o *TeamObjectResponse) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *TeamObjectResponse) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *TeamObjectResponse) SetRating(v float32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *TeamObjectResponse) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetWins

`func (o *TeamObjectResponse) GetWins() int32`

GetWins returns the Wins field if non-nil, zero value otherwise.

### GetWinsOk

`func (o *TeamObjectResponse) GetWinsOk() (*int32, bool)`

GetWinsOk returns a tuple with the Wins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWins

`func (o *TeamObjectResponse) SetWins(v int32)`

SetWins sets Wins field to given value.

### HasWins

`func (o *TeamObjectResponse) HasWins() bool`

HasWins returns a boolean if a field has been set.

### GetLosses

`func (o *TeamObjectResponse) GetLosses() int32`

GetLosses returns the Losses field if non-nil, zero value otherwise.

### GetLossesOk

`func (o *TeamObjectResponse) GetLossesOk() (*int32, bool)`

GetLossesOk returns a tuple with the Losses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLosses

`func (o *TeamObjectResponse) SetLosses(v int32)`

SetLosses sets Losses field to given value.

### HasLosses

`func (o *TeamObjectResponse) HasLosses() bool`

HasLosses returns a boolean if a field has been set.

### GetLastMatchTime

`func (o *TeamObjectResponse) GetLastMatchTime() int32`

GetLastMatchTime returns the LastMatchTime field if non-nil, zero value otherwise.

### GetLastMatchTimeOk

`func (o *TeamObjectResponse) GetLastMatchTimeOk() (*int32, bool)`

GetLastMatchTimeOk returns a tuple with the LastMatchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMatchTime

`func (o *TeamObjectResponse) SetLastMatchTime(v int32)`

SetLastMatchTime sets LastMatchTime field to given value.

### HasLastMatchTime

`func (o *TeamObjectResponse) HasLastMatchTime() bool`

HasLastMatchTime returns a boolean if a field has been set.

### GetName

`func (o *TeamObjectResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamObjectResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamObjectResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamObjectResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TeamObjectResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TeamObjectResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTag

`func (o *TeamObjectResponse) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *TeamObjectResponse) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *TeamObjectResponse) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *TeamObjectResponse) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


