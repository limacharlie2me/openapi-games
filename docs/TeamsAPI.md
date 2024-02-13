# \TeamsAPI

All URIs are relative to *https://api.opendota.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTeams**](TeamsAPI.md#GetTeams) | **Get** /teams | GET /teams
[**GetTeamsByTeamId**](TeamsAPI.md#GetTeamsByTeamId) | **Get** /teams/{team_id} | GET /teams/{team_id}
[**GetTeamsByTeamIdSelectHeroes**](TeamsAPI.md#GetTeamsByTeamIdSelectHeroes) | **Get** /teams/{team_id}/heroes | GET /teams/{team_id}/heroes
[**GetTeamsByTeamIdSelectMatches**](TeamsAPI.md#GetTeamsByTeamIdSelectMatches) | **Get** /teams/{team_id}/matches | GET /teams/{team_id}/matches
[**GetTeamsByTeamIdSelectPlayers**](TeamsAPI.md#GetTeamsByTeamIdSelectPlayers) | **Get** /teams/{team_id}/players | GET /teams/{team_id}/players



## GetTeams

> []TeamObjectResponse GetTeams(ctx).Page(page).Execute()

GET /teams



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
	page := int32(56) // int32 | Page number, zero indexed. Each page returns up to 1000 entries. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.GetTeams(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GetTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeams`: []TeamObjectResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GetTeams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number, zero indexed. Each page returns up to 1000 entries. | 

### Return type

[**[]TeamObjectResponse**](TeamObjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamsByTeamId

> TeamObjectResponse GetTeamsByTeamId(ctx, teamId).Execute()

GET /teams/{team_id}



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
	teamId := int32(56) // int32 | Team ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.GetTeamsByTeamId(context.Background(), teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GetTeamsByTeamId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamsByTeamId`: TeamObjectResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GetTeamsByTeamId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | Team ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamsByTeamIdRequest struct via the builder pattern


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


## GetTeamsByTeamIdSelectHeroes

> TeamHeroesResponse GetTeamsByTeamIdSelectHeroes(ctx, teamId).Execute()

GET /teams/{team_id}/heroes



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
	teamId := int32(56) // int32 | Team ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.GetTeamsByTeamIdSelectHeroes(context.Background(), teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GetTeamsByTeamIdSelectHeroes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamsByTeamIdSelectHeroes`: TeamHeroesResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GetTeamsByTeamIdSelectHeroes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | Team ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamsByTeamIdSelectHeroesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TeamHeroesResponse**](TeamHeroesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamsByTeamIdSelectMatches

> TeamMatchObjectResponse GetTeamsByTeamIdSelectMatches(ctx, teamId).Execute()

GET /teams/{team_id}/matches



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
	teamId := int32(56) // int32 | Team ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.GetTeamsByTeamIdSelectMatches(context.Background(), teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GetTeamsByTeamIdSelectMatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamsByTeamIdSelectMatches`: TeamMatchObjectResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GetTeamsByTeamIdSelectMatches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | Team ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamsByTeamIdSelectMatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TeamMatchObjectResponse**](TeamMatchObjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamsByTeamIdSelectPlayers

> TeamPlayersResponse GetTeamsByTeamIdSelectPlayers(ctx, teamId).Execute()

GET /teams/{team_id}/players



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
	teamId := int32(56) // int32 | Team ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.GetTeamsByTeamIdSelectPlayers(context.Background(), teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GetTeamsByTeamIdSelectPlayers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamsByTeamIdSelectPlayers`: TeamPlayersResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GetTeamsByTeamIdSelectPlayers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | Team ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamsByTeamIdSelectPlayersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TeamPlayersResponse**](TeamPlayersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

