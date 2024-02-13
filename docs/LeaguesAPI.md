# \LeaguesAPI

All URIs are relative to *https://api.opendota.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLeagues**](LeaguesAPI.md#GetLeagues) | **Get** /leagues | GET /leagues
[**GetLeaguesByLeagueId**](LeaguesAPI.md#GetLeaguesByLeagueId) | **Get** /leagues/{league_id} | GET /leagues/{league_id}
[**GetLeaguesByLeagueIdSelectMatches**](LeaguesAPI.md#GetLeaguesByLeagueIdSelectMatches) | **Get** /leagues/{league_id}/matches | GET /leagues/{league_id}/matches
[**GetLeaguesByLeagueIdSelectTeams**](LeaguesAPI.md#GetLeaguesByLeagueIdSelectTeams) | **Get** /leagues/{league_id}/teams | GET /leagues/{league_id}/teams



## GetLeagues

> []LeagueObjectResponse GetLeagues(ctx).Execute()

GET /leagues



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
	resp, r, err := apiClient.LeaguesAPI.GetLeagues(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaguesAPI.GetLeagues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeagues`: []LeagueObjectResponse
	fmt.Fprintf(os.Stdout, "Response from `LeaguesAPI.GetLeagues`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaguesRequest struct via the builder pattern


### Return type

[**[]LeagueObjectResponse**](LeagueObjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeaguesByLeagueId

> []LeagueObjectResponse GetLeaguesByLeagueId(ctx, leagueId).Execute()

GET /leagues/{league_id}



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
	leagueId := int32(56) // int32 | League ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaguesAPI.GetLeaguesByLeagueId(context.Background(), leagueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaguesAPI.GetLeaguesByLeagueId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaguesByLeagueId`: []LeagueObjectResponse
	fmt.Fprintf(os.Stdout, "Response from `LeaguesAPI.GetLeaguesByLeagueId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**leagueId** | **int32** | League ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaguesByLeagueIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LeagueObjectResponse**](LeagueObjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeaguesByLeagueIdSelectMatches

> MatchObjectResponse GetLeaguesByLeagueIdSelectMatches(ctx, leagueId).Execute()

GET /leagues/{league_id}/matches



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
	leagueId := int32(56) // int32 | League ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaguesAPI.GetLeaguesByLeagueIdSelectMatches(context.Background(), leagueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaguesAPI.GetLeaguesByLeagueIdSelectMatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaguesByLeagueIdSelectMatches`: MatchObjectResponse
	fmt.Fprintf(os.Stdout, "Response from `LeaguesAPI.GetLeaguesByLeagueIdSelectMatches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**leagueId** | **int32** | League ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaguesByLeagueIdSelectMatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MatchObjectResponse**](MatchObjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeaguesByLeagueIdSelectTeams

> TeamObjectResponse GetLeaguesByLeagueIdSelectTeams(ctx, leagueId).Execute()

GET /leagues/{league_id}/teams



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
	leagueId := int32(56) // int32 | League ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaguesAPI.GetLeaguesByLeagueIdSelectTeams(context.Background(), leagueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaguesAPI.GetLeaguesByLeagueIdSelectTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaguesByLeagueIdSelectTeams`: TeamObjectResponse
	fmt.Fprintf(os.Stdout, "Response from `LeaguesAPI.GetLeaguesByLeagueIdSelectTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**leagueId** | **int32** | League ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaguesByLeagueIdSelectTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TeamObjectResponse**](TeamObjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

