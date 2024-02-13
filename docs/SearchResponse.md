# SearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** | The player account ID | [optional] 
**Avatarfull** | Pointer to **NullableString** | avatarfull | [optional] 
**Personaname** | Pointer to **NullableString** | Player&#39;s Steam name | [optional] 
**LastMatchTime** | Pointer to **string** | last_match_time. May not be present or null. | [optional] 
**Similarity** | Pointer to **float32** | similarity | [optional] 

## Methods

### NewSearchResponse

`func NewSearchResponse() *SearchResponse`

NewSearchResponse instantiates a new SearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseWithDefaults

`func NewSearchResponseWithDefaults() *SearchResponse`

NewSearchResponseWithDefaults instantiates a new SearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SearchResponse) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SearchResponse) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SearchResponse) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *SearchResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAvatarfull

`func (o *SearchResponse) GetAvatarfull() string`

GetAvatarfull returns the Avatarfull field if non-nil, zero value otherwise.

### GetAvatarfullOk

`func (o *SearchResponse) GetAvatarfullOk() (*string, bool)`

GetAvatarfullOk returns a tuple with the Avatarfull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarfull

`func (o *SearchResponse) SetAvatarfull(v string)`

SetAvatarfull sets Avatarfull field to given value.

### HasAvatarfull

`func (o *SearchResponse) HasAvatarfull() bool`

HasAvatarfull returns a boolean if a field has been set.

### SetAvatarfullNil

`func (o *SearchResponse) SetAvatarfullNil(b bool)`

 SetAvatarfullNil sets the value for Avatarfull to be an explicit nil

### UnsetAvatarfull
`func (o *SearchResponse) UnsetAvatarfull()`

UnsetAvatarfull ensures that no value is present for Avatarfull, not even an explicit nil
### GetPersonaname

`func (o *SearchResponse) GetPersonaname() string`

GetPersonaname returns the Personaname field if non-nil, zero value otherwise.

### GetPersonanameOk

`func (o *SearchResponse) GetPersonanameOk() (*string, bool)`

GetPersonanameOk returns a tuple with the Personaname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonaname

`func (o *SearchResponse) SetPersonaname(v string)`

SetPersonaname sets Personaname field to given value.

### HasPersonaname

`func (o *SearchResponse) HasPersonaname() bool`

HasPersonaname returns a boolean if a field has been set.

### SetPersonanameNil

`func (o *SearchResponse) SetPersonanameNil(b bool)`

 SetPersonanameNil sets the value for Personaname to be an explicit nil

### UnsetPersonaname
`func (o *SearchResponse) UnsetPersonaname()`

UnsetPersonaname ensures that no value is present for Personaname, not even an explicit nil
### GetLastMatchTime

`func (o *SearchResponse) GetLastMatchTime() string`

GetLastMatchTime returns the LastMatchTime field if non-nil, zero value otherwise.

### GetLastMatchTimeOk

`func (o *SearchResponse) GetLastMatchTimeOk() (*string, bool)`

GetLastMatchTimeOk returns a tuple with the LastMatchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMatchTime

`func (o *SearchResponse) SetLastMatchTime(v string)`

SetLastMatchTime sets LastMatchTime field to given value.

### HasLastMatchTime

`func (o *SearchResponse) HasLastMatchTime() bool`

HasLastMatchTime returns a boolean if a field has been set.

### GetSimilarity

`func (o *SearchResponse) GetSimilarity() float32`

GetSimilarity returns the Similarity field if non-nil, zero value otherwise.

### GetSimilarityOk

`func (o *SearchResponse) GetSimilarityOk() (*float32, bool)`

GetSimilarityOk returns a tuple with the Similarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimilarity

`func (o *SearchResponse) SetSimilarity(v float32)`

SetSimilarity sets Similarity field to given value.

### HasSimilarity

`func (o *SearchResponse) HasSimilarity() bool`

HasSimilarity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


