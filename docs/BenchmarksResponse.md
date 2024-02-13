# BenchmarksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeroId** | Pointer to **int32** | The ID value of the hero played | [optional] 
**Result** | Pointer to [**BenchmarksResponseResult**](BenchmarksResponseResult.md) |  | [optional] 

## Methods

### NewBenchmarksResponse

`func NewBenchmarksResponse() *BenchmarksResponse`

NewBenchmarksResponse instantiates a new BenchmarksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarksResponseWithDefaults

`func NewBenchmarksResponseWithDefaults() *BenchmarksResponse`

NewBenchmarksResponseWithDefaults instantiates a new BenchmarksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeroId

`func (o *BenchmarksResponse) GetHeroId() int32`

GetHeroId returns the HeroId field if non-nil, zero value otherwise.

### GetHeroIdOk

`func (o *BenchmarksResponse) GetHeroIdOk() (*int32, bool)`

GetHeroIdOk returns a tuple with the HeroId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroId

`func (o *BenchmarksResponse) SetHeroId(v int32)`

SetHeroId sets HeroId field to given value.

### HasHeroId

`func (o *BenchmarksResponse) HasHeroId() bool`

HasHeroId returns a boolean if a field has been set.

### GetResult

`func (o *BenchmarksResponse) GetResult() BenchmarksResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BenchmarksResponse) GetResultOk() (*BenchmarksResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BenchmarksResponse) SetResult(v BenchmarksResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *BenchmarksResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


