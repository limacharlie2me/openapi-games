# \HeroesAPI

All URIs are relative to *https://api.opendota.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHeroes**](HeroesAPI.md#GetHeroes) | **Get** /heroes | GET /heroes
[**GetHeroesByHeroIdSelectDurations**](HeroesAPI.md#GetHeroesByHeroIdSelectDurations) | **Get** /heroes/{hero_id}/durations | GET /heroes/{hero_id}/durations
[**GetHeroesByHeroIdSelectItemPopularity**](HeroesAPI.md#GetHeroesByHeroIdSelectItemPopularity) | **Get** /heroes/{hero_id}/itemPopularity | GET /heroes/{hero_id}/itemPopularity
[**GetHeroesByHeroIdSelectMatches**](HeroesAPI.md#GetHeroesByHeroIdSelectMatches) | **Get** /heroes/{hero_id}/matches | GET /heroes/{hero_id}/matches
[**GetHeroesByHeroIdSelectMatchups**](HeroesAPI.md#GetHeroesByHeroIdSelectMatchups) | **Get** /heroes/{hero_id}/matchups | GET /heroes/{hero_id}/matchups
[**GetHeroesByHeroIdSelectPlayers**](HeroesAPI.md#GetHeroesByHeroIdSelectPlayers) | **Get** /heroes/{hero_id}/players | GET /heroes/{hero_id}/players



## GetHeroes

> []HeroObjectResponse GetHeroes(ctx).Execute()

GET /heroes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HeroesAPI.GetHeroes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HeroesAPI.GetHeroes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHeroes`: []HeroObjectResponse
	fmt.Fprintf(os.Stdout, "Response from `HeroesAPI.GetHeroes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHeroesRequest struct via the builder pattern


### Return type

[**[]HeroObjectResponse**](HeroObjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHeroesByHeroIdSelectDurations

> []HeroDurationsResponse GetHeroesByHeroIdSelectDurations(ctx, heroId).Execute()

GET /heroes/{hero_id}/durations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	heroId := int32(56) // int32 | Hero ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HeroesAPI.GetHeroesByHeroIdSelectDurations(context.Background(), heroId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HeroesAPI.GetHeroesByHeroIdSelectDurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHeroesByHeroIdSelectDurations`: []HeroDurationsResponse
	fmt.Fprintf(os.Stdout, "Response from `HeroesAPI.GetHeroesByHeroIdSelectDurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**heroId** | **int32** | Hero ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHeroesByHeroIdSelectDurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HeroDurationsResponse**](HeroDurationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHeroesByHeroIdSelectItemPopularity

> HeroItemPopularityResponse GetHeroesByHeroIdSelectItemPopularity(ctx, heroId).Execute()

GET /heroes/{hero_id}/itemPopularity



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	heroId := int32(56) // int32 | Hero ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HeroesAPI.GetHeroesByHeroIdSelectItemPopularity(context.Background(), heroId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HeroesAPI.GetHeroesByHeroIdSelectItemPopularity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHeroesByHeroIdSelectItemPopularity`: HeroItemPopularityResponse
	fmt.Fprintf(os.Stdout, "Response from `HeroesAPI.GetHeroesByHeroIdSelectItemPopularity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**heroId** | **int32** | Hero ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHeroesByHeroIdSelectItemPopularityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HeroItemPopularityResponse**](HeroItemPopularityResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHeroesByHeroIdSelectMatches

> []MatchObjectResponse GetHeroesByHeroIdSelectMatches(ctx, heroId).Execute()

GET /heroes/{hero_id}/matches



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	heroId := int32(56) // int32 | Hero ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HeroesAPI.GetHeroesByHeroIdSelectMatches(context.Background(), heroId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HeroesAPI.GetHeroesByHeroIdSelectMatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHeroesByHeroIdSelectMatches`: []MatchObjectResponse
	fmt.Fprintf(os.Stdout, "Response from `HeroesAPI.GetHeroesByHeroIdSelectMatches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**heroId** | **int32** | Hero ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHeroesByHeroIdSelectMatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MatchObjectResponse**](MatchObjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHeroesByHeroIdSelectMatchups

> []HeroMatchupsResponse GetHeroesByHeroIdSelectMatchups(ctx, heroId).Execute()

GET /heroes/{hero_id}/matchups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	heroId := int32(56) // int32 | Hero ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HeroesAPI.GetHeroesByHeroIdSelectMatchups(context.Background(), heroId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HeroesAPI.GetHeroesByHeroIdSelectMatchups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHeroesByHeroIdSelectMatchups`: []HeroMatchupsResponse
	fmt.Fprintf(os.Stdout, "Response from `HeroesAPI.GetHeroesByHeroIdSelectMatchups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**heroId** | **int32** | Hero ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHeroesByHeroIdSelectMatchupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HeroMatchupsResponse**](HeroMatchupsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHeroesByHeroIdSelectPlayers

> [][]PlayerObjectResponse GetHeroesByHeroIdSelectPlayers(ctx, heroId).Execute()

GET /heroes/{hero_id}/players



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	heroId := int32(56) // int32 | Hero ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HeroesAPI.GetHeroesByHeroIdSelectPlayers(context.Background(), heroId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HeroesAPI.GetHeroesByHeroIdSelectPlayers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHeroesByHeroIdSelectPlayers`: [][]PlayerObjectResponse
	fmt.Fprintf(os.Stdout, "Response from `HeroesAPI.GetHeroesByHeroIdSelectPlayers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**heroId** | **int32** | Hero ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHeroesByHeroIdSelectPlayersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[][]PlayerObjectResponse**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

