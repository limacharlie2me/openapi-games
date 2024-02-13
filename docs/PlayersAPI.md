# \PlayersAPI

All URIs are relative to *https://api.opendota.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPlayersByAccountId**](PlayersAPI.md#GetPlayersByAccountId) | **Get** /players/{account_id} | GET /players/{account_id}
[**GetPlayersByAccountIdHistogramsByField**](PlayersAPI.md#GetPlayersByAccountIdHistogramsByField) | **Get** /players/{account_id}/histograms/{field} | GET /players/{account_id}/histograms
[**GetPlayersByAccountIdSelectCounts**](PlayersAPI.md#GetPlayersByAccountIdSelectCounts) | **Get** /players/{account_id}/counts | GET /players/{account_id}/counts
[**GetPlayersByAccountIdSelectHeroes**](PlayersAPI.md#GetPlayersByAccountIdSelectHeroes) | **Get** /players/{account_id}/heroes | GET /players/{account_id}/heroes
[**GetPlayersByAccountIdSelectMatches**](PlayersAPI.md#GetPlayersByAccountIdSelectMatches) | **Get** /players/{account_id}/matches | GET /players/{account_id}/matches
[**GetPlayersByAccountIdSelectPeers**](PlayersAPI.md#GetPlayersByAccountIdSelectPeers) | **Get** /players/{account_id}/peers | GET /players/{account_id}/peers
[**GetPlayersByAccountIdSelectPros**](PlayersAPI.md#GetPlayersByAccountIdSelectPros) | **Get** /players/{account_id}/pros | GET /players/{account_id}/pros
[**GetPlayersByAccountIdSelectRankings**](PlayersAPI.md#GetPlayersByAccountIdSelectRankings) | **Get** /players/{account_id}/rankings | GET /players/{account_id}/rankings
[**GetPlayersByAccountIdSelectRatings**](PlayersAPI.md#GetPlayersByAccountIdSelectRatings) | **Get** /players/{account_id}/ratings | GET /players/{account_id}/ratings
[**GetPlayersByAccountIdSelectRecentMatches**](PlayersAPI.md#GetPlayersByAccountIdSelectRecentMatches) | **Get** /players/{account_id}/recentMatches | GET /players/{account_id}/recentMatches
[**GetPlayersByAccountIdSelectTotals**](PlayersAPI.md#GetPlayersByAccountIdSelectTotals) | **Get** /players/{account_id}/totals | GET /players/{account_id}/totals
[**GetPlayersByAccountIdSelectWardmap**](PlayersAPI.md#GetPlayersByAccountIdSelectWardmap) | **Get** /players/{account_id}/wardmap | GET /players/{account_id}/wardmap
[**GetPlayersByAccountIdSelectWl**](PlayersAPI.md#GetPlayersByAccountIdSelectWl) | **Get** /players/{account_id}/wl | GET /players/{account_id}/wl
[**GetPlayersByAccountIdSelectWordcloud**](PlayersAPI.md#GetPlayersByAccountIdSelectWordcloud) | **Get** /players/{account_id}/wordcloud | GET /players/{account_id}/wordcloud
[**PostRefresh**](PlayersAPI.md#PostRefresh) | **Post** /players/{account_id}/refresh | POST /players/{account_id}/refresh



## GetPlayersByAccountId

> PlayersResponse GetPlayersByAccountId(ctx, accountId).Execute()

GET /players/{account_id}



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
	accountId := int32(56) // int32 | Steam32 account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayersByAccountId(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayersByAccountId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayersByAccountId`: PlayersResponse
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayersByAccountId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Steam32 account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayersByAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlayersResponse**](PlayersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayersByAccountIdHistogramsByField

> []map[string]interface{} GetPlayersByAccountIdHistogramsByField(ctx, accountId, field).Limit(limit).Offset(offset).Win(win).Patch(patch).GameMode(gameMode).LobbyType(lobbyType).Region(region).Date(date).LaneRole(laneRole).HeroId(heroId).IsRadiant(isRadiant).IncludedAccountId(includedAccountId).ExcludedAccountId(excludedAccountId).WithHeroId(withHeroId).AgainstHeroId(againstHeroId).Significant(significant).Having(having).Sort(sort).Execute()

GET /players/{account_id}/histograms



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
	accountId := int32(56) // int32 | Steam32 account ID
	field := "field_example" // string | Field to aggregate on
	limit := int32(56) // int32 | Number of matches to limit to (optional)
	offset := int32(56) // int32 | Number of matches to offset start by (optional)
	win := int32(56) // int32 | Whether the player won (optional)
	patch := int32(56) // int32 | Patch ID, from dotaconstants (optional)
	gameMode := int32(56) // int32 | Game Mode ID (optional)
	lobbyType := int32(56) // int32 | Lobby type ID (optional)
	region := int32(56) // int32 | Region ID (optional)
	date := int32(56) // int32 | Days previous (optional)
	laneRole := int32(56) // int32 | Lane Role ID (optional)
	heroId := int32(56) // int32 | Hero ID (optional)
	isRadiant := int32(56) // int32 | Whether the player was radiant (optional)
	includedAccountId := int32(56) // int32 | Account IDs in the match (array) (optional)
	excludedAccountId := int32(56) // int32 | Account IDs not in the match (array) (optional)
	withHeroId := int32(56) // int32 | Hero IDs on the player's team (array) (optional)
	againstHeroId := int32(56) // int32 | Hero IDs against the player's team (array) (optional)
	significant := int32(56) // int32 | Whether the match was significant for aggregation purposes. Defaults to 1 (true), set this to 0 to return data for non-standard modes/matches. (optional)
	having := int32(56) // int32 | The minimum number of games played, for filtering hero stats (optional)
	sort := "sort_example" // string | The field to return matches sorted by in descending order (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayersByAccountIdHistogramsByField(context.Background(), accountId, field).Limit(limit).Offset(offset).Win(win).Patch(patch).GameMode(gameMode).LobbyType(lobbyType).Region(region).Date(date).LaneRole(laneRole).HeroId(heroId).IsRadiant(isRadiant).IncludedAccountId(includedAccountId).ExcludedAccountId(excludedAccountId).WithHeroId(withHeroId).AgainstHeroId(againstHeroId).Significant(significant).Having(having).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayersByAccountIdHistogramsByField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayersByAccountIdHistogramsByField`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayersByAccountIdHistogramsByField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Steam32 account ID | 
**field** | **string** | Field to aggregate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayersByAccountIdHistogramsByFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Number of matches to limit to | 
 **offset** | **int32** | Number of matches to offset start by | 
 **win** | **int32** | Whether the player won | 
 **patch** | **int32** | Patch ID, from dotaconstants | 
 **gameMode** | **int32** | Game Mode ID | 
 **lobbyType** | **int32** | Lobby type ID | 
 **region** | **int32** | Region ID | 
 **date** | **int32** | Days previous | 
 **laneRole** | **int32** | Lane Role ID | 
 **heroId** | **int32** | Hero ID | 
 **isRadiant** | **int32** | Whether the player was radiant | 
 **includedAccountId** | **int32** | Account IDs in the match (array) | 
 **excludedAccountId** | **int32** | Account IDs not in the match (array) | 
 **withHeroId** | **int32** | Hero IDs on the player&#39;s team (array) | 
 **againstHeroId** | **int32** | Hero IDs against the player&#39;s team (array) | 
 **significant** | **int32** | Whether the match was significant for aggregation purposes. Defaults to 1 (true), set this to 0 to return data for non-standard modes/matches. | 
 **having** | **int32** | The minimum number of games played, for filtering hero stats | 
 **sort** | **string** | The field to return matches sorted by in descending order | 

### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayersByAccountIdSelectCounts

> PlayerCountsResponse GetPlayersByAccountIdSelectCounts(ctx, accountId).Limit(limit).Offset(offset).Win(win).Patch(patch).GameMode(gameMode).LobbyType(lobbyType).Region(region).Date(date).LaneRole(laneRole).HeroId(heroId).IsRadiant(isRadiant).IncludedAccountId(includedAccountId).ExcludedAccountId(excludedAccountId).WithHeroId(withHeroId).AgainstHeroId(againstHeroId).Significant(significant).Having(having).Sort(sort).Execute()

GET /players/{account_id}/counts



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
	accountId := int32(56) // int32 | Steam32 account ID
	limit := int32(56) // int32 | Number of matches to limit to (optional)
	offset := int32(56) // int32 | Number of matches to offset start by (optional)
	win := int32(56) // int32 | Whether the player won (optional)
	patch := int32(56) // int32 | Patch ID, from dotaconstants (optional)
	gameMode := int32(56) // int32 | Game Mode ID (optional)
	lobbyType := int32(56) // int32 | Lobby type ID (optional)
	region := int32(56) // int32 | Region ID (optional)
	date := int32(56) // int32 | Days previous (optional)
	laneRole := int32(56) // int32 | Lane Role ID (optional)
	heroId := int32(56) // int32 | Hero ID (optional)
	isRadiant := int32(56) // int32 | Whether the player was radiant (optional)
	includedAccountId := int32(56) // int32 | Account IDs in the match (array) (optional)
	excludedAccountId := int32(56) // int32 | Account IDs not in the match (array) (optional)
	withHeroId := int32(56) // int32 | Hero IDs on the player's team (array) (optional)
	againstHeroId := int32(56) // int32 | Hero IDs against the player's team (array) (optional)
	significant := int32(56) // int32 | Whether the match was significant for aggregation purposes. Defaults to 1 (true), set this to 0 to return data for non-standard modes/matches. (optional)
	having := int32(56) // int32 | The minimum number of games played, for filtering hero stats (optional)
	sort := "sort_example" // string | The field to return matches sorted by in descending order (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectCounts(context.Background(), accountId).Limit(limit).Offset(offset).Win(win).Patch(patch).GameMode(gameMode).LobbyType(lobbyType).Region(region).Date(date).LaneRole(laneRole).HeroId(heroId).IsRadiant(isRadiant).IncludedAccountId(includedAccountId).ExcludedAccountId(excludedAccountId).WithHeroId(withHeroId).AgainstHeroId(againstHeroId).Significant(significant).Having(having).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayersByAccountIdSelectCounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayersByAccountIdSelectCounts`: PlayerCountsResponse
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayersByAccountIdSelectCounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Steam32 account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayersByAccountIdSelectCountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of matches to limit to | 
 **offset** | **int32** | Number of matches to offset start by | 
 **win** | **int32** | Whether the player won | 
 **patch** | **int32** | Patch ID, from dotaconstants | 
 **gameMode** | **int32** | Game Mode ID | 
 **lobbyType** | **int32** | Lobby type ID | 
 **region** | **int32** | Region ID | 
 **date** | **int32** | Days previous | 
 **laneRole** | **int32** | Lane Role ID | 
 **heroId** | **int32** | Hero ID | 
 **isRadiant** | **int32** | Whether the player was radiant | 
 **includedAccountId** | **int32** | Account IDs in the match (array) | 
 **excludedAccountId** | **int32** | Account IDs not in the match (array) | 
 **withHeroId** | **int32** | Hero IDs on the player&#39;s team (array) | 
 **againstHeroId** | **int32** | Hero IDs against the player&#39;s team (array) | 
 **significant** | **int32** | Whether the match was significant for aggregation purposes. Defaults to 1 (true), set this to 0 to return data for non-standard modes/matches. | 
 **having** | **int32** | The minimum number of games played, for filtering hero stats | 
 **sort** | **string** | The field to return matches sorted by in descending order | 

### Return type

[**PlayerCountsResponse**](PlayerCountsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayersByAccountIdSelectHeroes

> []PlayerHeroesResponse GetPlayersByAccountIdSelectHeroes(ctx, accountId).Limit(limit).Offset(offset).Win(win).Patch(patch).GameMode(gameMode).LobbyType(lobbyType).Region(region).Date(date).LaneRole(laneRole).HeroId(heroId).IsRadiant(isRadiant).IncludedAccountId(includedAccountId).ExcludedAccountId(excludedAccountId).WithHeroId(withHeroId).AgainstHeroId(againstHeroId).Significant(significant).Having(having).Sort(sort).Execute()

GET /players/{account_id}/heroes



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
	accountId := int32(56) // int32 | Steam32 account ID
	limit := int32(56) // int32 | Number of matches to limit to (optional)
	offset := int32(56) // int32 | Number of matches to offset start by (optional)
	win := int32(56) // int32 | Whether the player won (optional)
	patch := int32(56) // int32 | Patch ID, from dotaconstants (optional)
	gameMode := int32(56) // int32 | Game Mode ID (optional)
	lobbyType := int32(56) // int32 | Lobby type ID (optional)
	region := int32(56) // int32 | Region ID (optional)
	date := int32(56) // int32 | Days previous (optional)
	laneRole := int32(56) // int32 | Lane Role ID (optional)
	heroId := int32(56) // int32 | Hero ID (optional)
	isRadiant := int32(56) // int32 | Whether the player was radiant (optional)
	includedAccountId := int32(56) // int32 | Account IDs in the match (array) (optional)
	excludedAccountId := int32(56) // int32 | Account IDs not in the match (array) (optional)
	withHeroId := int32(56) // int32 | Hero IDs on the player's team (array) (optional)
	againstHeroId := int32(56) // int32 | Hero IDs against the player's team (array) (optional)
	significant := int32(56) // int32 | Whether the match was significant for aggregation purposes. Defaults to 1 (true), set this to 0 to return data for non-standard modes/matches. (optional)
	having := int32(56) // int32 | The minimum number of games played, for filtering hero stats (optional)
	sort := "sort_example" // string | The field to return matches sorted by in descending order (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectHeroes(context.Background(), accountId).Limit(limit).Offset(offset).Win(win).Patch(patch).GameMode(gameMode).LobbyType(lobbyType).Region(region).Date(date).LaneRole(laneRole).HeroId(heroId).IsRadiant(isRadiant).IncludedAccountId(includedAccountId).ExcludedAccountId(excludedAccountId).WithHeroId(withHeroId).AgainstHeroId(againstHeroId).Significant(significant).Having(having).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayersByAccountIdSelectHeroes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayersByAccountIdSelectHeroes`: []PlayerHeroesResponse
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayersByAccountIdSelectHeroes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Steam32 account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayersByAccountIdSelectHeroesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of matches to limit to | 
 **offset** | **int32** | Number of matches to offset start by | 
 **win** | **int32** | Whether the player won | 
 **patch** | **int32** | Patch ID, from dotaconstants | 
 **gameMode** | **int32** | Game Mode ID | 
 **lobbyType** | **int32** | Lobby type ID | 
 **region** | **int32** | Region ID | 
 **date** | **int32** | Days previous | 
 **laneRole** | **int32** | Lane Role ID | 
 **heroId** | **int32** | Hero ID | 
 **isRadiant** | **int32** | Whether the player was radiant | 
 **includedAccountId** | **int32** | Account IDs in the match (array) | 
 **excludedAccountId** | **int32** | Account IDs not in the match (array) | 
 **withHeroId** | **int32** | Hero IDs on the player&#39;s team (array) | 
 **againstHeroId** | **int32** | Hero IDs against the player&#39;s team (array) | 
 **significant** | **int32** | Whether the match was significant for aggregation purposes. Defaults to 1 (true), set this to 0 to return data for non-standard modes/matches. | 
 **having** | **int32** | The minimum number of games played, for filtering hero stats | 
 **sort** | **string** | The field to return matches sorted by in descending order | 

### Return type

[**[]PlayerHeroesResponse**](PlayerHeroesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayersByAccountIdSelectMatches

> []PlayerMatchesResponse GetPlayersByAccountIdSelectMatches(ctx, accountId).Limit(limit).Offset(offset).Win(win).Patch(patch).GameMode(gameMode).LobbyType(lobbyType).Region(region).Date(date).LaneRole(laneRole).HeroId(heroId).IsRadiant(isRadiant).IncludedAccountId(includedAccountId).ExcludedAccountId(excludedAccountId).WithHeroId(withHeroId).AgainstHeroId(againstHeroId).Significant(significant).Having(having).Sort(sort).Project(project).Execute()

GET /players/{account_id}/matches



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
	accountId := int32(56) // int32 | Steam32 account ID
	limit := int32(56) // int32 | Number of matches to limit to (optional)
	offset := int32(56) // int32 | Number of matches to offset start by (optional)
	win := int32(56) // int32 | Whether the player won (optional)
	patch := int32(56) // int32 | Patch ID, from dotaconstants (optional)
	gameMode := int32(56) // int32 | Game Mode ID (optional)
	lobbyType := int32(56) // int32 | Lobby type ID (optional)
	region := int32(56) // int32 | Region ID (optional)
	date := int32(56) // int32 | Days previous (optional)
	laneRole := int32(56) // int32 | Lane Role ID (optional)
	heroId := int32(56) // int32 | Hero ID (optional)
	isRadiant := int32(56) // int32 | Whether the player was radiant (optional)
	includedAccountId := int32(56) // int32 | Account IDs in the match (array) (optional)
	excludedAccountId := int32(56) // int32 | Account IDs not in the match (array) (optional)
	withHeroId := int32(56) // int32 | Hero IDs on the player's team (array) (optional)
	againstHeroId := int32(56) // int32 | Hero IDs against the player's team (array) (optional)
	significant := int32(56) // int32 | Whether the match was significant for aggregation purposes. Defaults to 1 (true), set this to 0 to return data for non-standard modes/matches. (optional)
	having := int32(56) // int32 | The minimum number of games played, for filtering hero stats (optional)
	sort := "sort_example" // string | The field to return matches sorted by in descending order (optional)
	project := "project_example" // string | Fields to project (array) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectMatches(context.Background(), accountId).Limit(limit).Offset(offset).Win(win).Patch(patch).GameMode(gameMode).LobbyType(lobbyType).Region(region).Date(date).LaneRole(laneRole).HeroId(heroId).IsRadiant(isRadiant).IncludedAccountId(includedAccountId).ExcludedAccountId(excludedAccountId).WithHeroId(withHeroId).AgainstHeroId(againstHeroId).Significant(significant).Having(having).Sort(sort).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayersByAccountIdSelectMatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayersByAccountIdSelectMatches`: []PlayerMatchesResponse
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayersByAccountIdSelectMatches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Steam32 account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayersByAccountIdSelectMatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of matches to limit to | 
 **offset** | **int32** | Number of matches to offset start by | 
 **win** | **int32** | Whether the player won | 
 **patch** | **int32** | Patch ID, from dotaconstants | 
 **gameMode** | **int32** | Game Mode ID | 
 **lobbyType** | **int32** | Lobby type ID | 
 **region** | **int32** | Region ID | 
 **date** | **int32** | Days previous | 
 **laneRole** | **int32** | Lane Role ID | 
 **heroId** | **int32** | Hero ID | 
 **isRadiant** | **int32** | Whether the player was radiant | 
 **includedAccountId** | **int32** | Account IDs in the match (array) | 
 **excludedAccountId** | **int32** | Account IDs not in the match (array) | 
 **withHeroId** | **int32** | Hero IDs on the player&#39;s team (array) | 
 **againstHeroId** | **int32** | Hero IDs against the player&#39;s team (array) | 
 **significant** | **int32** | Whether the match was significant for aggregation purposes. Defaults to 1 (true), set this to 0 to return data for non-standard modes/matches. | 
 **having** | **int32** | The minimum number of games played, for filtering hero stats | 
 **sort** | **string** | The field to return matches sorted by in descending order | 
 **project** | **string** | Fields to project (array) | 

### Return type

[**[]PlayerMatchesResponse**](PlayerMatchesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayersByAccountIdSelectPeers

> []PlayerPeersResponse GetPlayersByAccountIdSelectPeers(ctx, accountId).Limit(limit).Offset(offset).Win(win).Patch(patch).GameMode(gameMode).LobbyType(lobbyType).Region(region).Date(date).LaneRole(laneRole).HeroId(heroId).IsRadiant(isRadiant).IncludedAccountId(includedAccountId).ExcludedAccountId(excludedAccountId).WithHeroId(withHeroId).AgainstHeroId(againstHeroId).Significant(significant).Having(having).Sort(sort).Execute()

GET /players/{account_id}/peers



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
	accountId := int32(56) // int32 | Steam32 account ID
	limit := int32(56) // int32 | Number of matches to limit to (optional)
	offset := int32(56) // int32 | Number of matches to offset start by (optional)
	win := int32(56) // int32 | Whether the player won (optional)
	patch := int32(56) // int32 | Patch ID, from dotaconstants (optional)
	gameMode := int32(56) // int32 | Game Mode ID (optional)
	lobbyType := int32(56) // int32 | Lobby type ID (optional)
	region := int32(56) // int32 | Region ID (optional)
	date := int32(56) // int32 | Days previous (optional)
	laneRole := int32(56) // int32 | Lane Role ID (optional)
	heroId := int32(56) // int32 | Hero ID (optional)
	isRadiant := int32(56) // int32 | Whether the player was radiant (optional)
	includedAccountId := int32(56) // int32 | Account IDs in the match (array) (optional)
	excludedAccountId := int32(56) // int32 | Account IDs not in the match (array) (optional)
	withHeroId := int32(56) // int32 | Hero IDs on the player's team (array) (optional)
	againstHeroId := int32(56) // int32 | Hero IDs against the player's team (array) (optional)
	significant := int32(56) // int32 | Whether the match was significant for aggregation purposes. Defaults to 1 (true), set this to 0 to return data for non-standard modes/matches. (optional)
	having := int32(56) // int32 | The minimum number of games played, for filtering hero stats (optional)
	sort := "sort_example" // string | The field to return matches sorted by in descending order (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectPeers(context.Background(), accountId).Limit(limit).Offset(offset).Win(win).Patch(patch).GameMode(gameMode).LobbyType(lobbyType).Region(region).Date(date).LaneRole(laneRole).HeroId(heroId).IsRadiant(isRadiant).IncludedAccountId(includedAccountId).ExcludedAccountId(excludedAccountId).WithHeroId(withHeroId).AgainstHeroId(againstHeroId).Significant(significant).Having(having).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayersByAccountIdSelectPeers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayersByAccountIdSelectPeers`: []PlayerPeersResponse
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayersByAccountIdSelectPeers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Steam32 account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayersByAccountIdSelectPeersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of matches to limit to | 
 **offset** | **int32** | Number of matches to offset start by | 
 **win** | **int32** | Whether the player won | 
 **patch** | **int32** | Patch ID, from dotaconstants | 
 **gameMode** | **int32** | Game Mode ID | 
 **lobbyType** | **int32** | Lobby type ID | 
 **region** | **int32** | Region ID | 
 **date** | **int32** | Days previous | 
 **laneRole** | **int32** | Lane Role ID | 
 **heroId** | **int32** | Hero ID | 
 **isRadiant** | **int32** | Whether the player was radiant | 
 **includedAccountId** | **int32** | Account IDs in the match (array) | 
 **excludedAccountId** | **int32** | Account IDs not in the match (array) | 
 **withHeroId** | **int32** | Hero IDs on the player&#39;s team (array) | 
 **againstHeroId** | **int32** | Hero IDs against the player&#39;s team (array) | 
 **significant** | **int32** | Whether the match was significant for aggregation purposes. Defaults to 1 (true), set this to 0 to return data for non-standard modes/matches. | 
 **having** | **int32** | The minimum number of games played, for filtering hero stats | 
 **sort** | **string** | The field to return matches sorted by in descending order | 

### Return type

[**[]PlayerPeersResponse**](PlayerPeersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayersByAccountIdSelectPros

> []PlayerProsResponse GetPlayersByAccountIdSelectPros(ctx, accountId).Limit(limit).Offset(offset).Win(win).Patch(patch).GameMode(gameMode).LobbyType(lobbyType).Region(region).Date(date).LaneRole(laneRole).HeroId(heroId).IsRadiant(isRadiant).IncludedAccountId(includedAccountId).ExcludedAccountId(excludedAccountId).WithHeroId(withHeroId).AgainstHeroId(againstHeroId).Significant(significant).Having(having).Sort(sort).Execute()

GET /players/{account_id}/pros



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
	accountId := int32(56) // int32 | Steam32 account ID
	limit := int32(56) // int32 | Number of matches to limit to (optional)
	offset := int32(56) // int32 | Number of matches to offset start by (optional)
	win := int32(56) // int32 | Whether the player won (optional)
	patch := int32(56) // int32 | Patch ID, from dotaconstants (optional)
	gameMode := int32(56) // int32 | Game Mode ID (optional)
	lobbyType := int32(56) // int32 | Lobby type ID (optional)
	region := int32(56) // int32 | Region ID (optional)
	date := int32(56) // int32 | Days previous (optional)
	laneRole := int32(56) // int32 | Lane Role ID (optional)
	heroId := int32(56) // int32 | Hero ID (optional)
	isRadiant := int32(56) // int32 | Whether the player was radiant (optional)
	includedAccountId := int32(56) // int32 | Account IDs in the match (array) (optional)
	excludedAccountId := int32(56) // int32 | Account IDs not in the match (array) (optional)
	withHeroId := int32(56) // int32 | Hero IDs on the player's team (array) (optional)
	againstHeroId := int32(56) // int32 | Hero IDs against the player's team (array) (optional)
	significant := int32(56) // int32 | Whether the match was significant for aggregation purposes. Defaults to 1 (true), set this to 0 to return data for non-standard modes/matches. (optional)
	having := int32(56) // int32 | The minimum number of games played, for filtering hero stats (optional)
	sort := "sort_example" // string | The field to return matches sorted by in descending order (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectPros(context.Background(), accountId).Limit(limit).Offset(offset).Win(win).Patch(patch).GameMode(gameMode).LobbyType(lobbyType).Region(region).Date(date).LaneRole(laneRole).HeroId(heroId).IsRadiant(isRadiant).IncludedAccountId(includedAccountId).ExcludedAccountId(excludedAccountId).WithHeroId(withHeroId).AgainstHeroId(againstHeroId).Significant(significant).Having(having).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayersByAccountIdSelectPros``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayersByAccountIdSelectPros`: []PlayerProsResponse
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayersByAccountIdSelectPros`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Steam32 account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayersByAccountIdSelectProsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of matches to limit to | 
 **offset** | **int32** | Number of matches to offset start by | 
 **win** | **int32** | Whether the player won | 
 **patch** | **int32** | Patch ID, from dotaconstants | 
 **gameMode** | **int32** | Game Mode ID | 
 **lobbyType** | **int32** | Lobby type ID | 
 **region** | **int32** | Region ID | 
 **date** | **int32** | Days previous | 
 **laneRole** | **int32** | Lane Role ID | 
 **heroId** | **int32** | Hero ID | 
 **isRadiant** | **int32** | Whether the player was radiant | 
 **includedAccountId** | **int32** | Account IDs in the match (array) | 
 **excludedAccountId** | **int32** | Account IDs not in the match (array) | 
 **withHeroId** | **int32** | Hero IDs on the player&#39;s team (array) | 
 **againstHeroId** | **int32** | Hero IDs against the player&#39;s team (array) | 
 **significant** | **int32** | Whether the match was significant for aggregation purposes. Defaults to 1 (true), set this to 0 to return data for non-standard modes/matches. | 
 **having** | **int32** | The minimum number of games played, for filtering hero stats | 
 **sort** | **string** | The field to return matches sorted by in descending order | 

### Return type

[**[]PlayerProsResponse**](PlayerProsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayersByAccountIdSelectRankings

> []PlayerRankingsResponse GetPlayersByAccountIdSelectRankings(ctx, accountId).Execute()

GET /players/{account_id}/rankings



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
	accountId := int32(56) // int32 | Steam32 account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectRankings(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayersByAccountIdSelectRankings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayersByAccountIdSelectRankings`: []PlayerRankingsResponse
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayersByAccountIdSelectRankings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Steam32 account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayersByAccountIdSelectRankingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PlayerRankingsResponse**](PlayerRankingsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayersByAccountIdSelectRatings

> []PlayerRatingsResponse GetPlayersByAccountIdSelectRatings(ctx, accountId).Execute()

GET /players/{account_id}/ratings



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
	accountId := int32(56) // int32 | Steam32 account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectRatings(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayersByAccountIdSelectRatings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayersByAccountIdSelectRatings`: []PlayerRatingsResponse
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayersByAccountIdSelectRatings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Steam32 account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayersByAccountIdSelectRatingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PlayerRatingsResponse**](PlayerRatingsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayersByAccountIdSelectRecentMatches

> [][]PlayerRecentMatchesResponse GetPlayersByAccountIdSelectRecentMatches(ctx, accountId).Execute()

GET /players/{account_id}/recentMatches



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
	accountId := int32(56) // int32 | Steam32 account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectRecentMatches(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayersByAccountIdSelectRecentMatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayersByAccountIdSelectRecentMatches`: [][]PlayerRecentMatchesResponse
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayersByAccountIdSelectRecentMatches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Steam32 account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayersByAccountIdSelectRecentMatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[][]PlayerRecentMatchesResponse**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayersByAccountIdSelectTotals

> []PlayerTotalsResponse GetPlayersByAccountIdSelectTotals(ctx, accountId).Limit(limit).Offset(offset).Win(win).Patch(patch).GameMode(gameMode).LobbyType(lobbyType).Region(region).Date(date).LaneRole(laneRole).HeroId(heroId).IsRadiant(isRadiant).IncludedAccountId(includedAccountId).ExcludedAccountId(excludedAccountId).WithHeroId(withHeroId).AgainstHeroId(againstHeroId).Significant(significant).Having(having).Sort(sort).Execute()

GET /players/{account_id}/totals



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
	accountId := int32(56) // int32 | Steam32 account ID
	limit := int32(56) // int32 | Number of matches to limit to (optional)
	offset := int32(56) // int32 | Number of matches to offset start by (optional)
	win := int32(56) // int32 | Whether the player won (optional)
	patch := int32(56) // int32 | Patch ID, from dotaconstants (optional)
	gameMode := int32(56) // int32 | Game Mode ID (optional)
	lobbyType := int32(56) // int32 | Lobby type ID (optional)
	region := int32(56) // int32 | Region ID (optional)
	date := int32(56) // int32 | Days previous (optional)
	laneRole := int32(56) // int32 | Lane Role ID (optional)
	heroId := int32(56) // int32 | Hero ID (optional)
	isRadiant := int32(56) // int32 | Whether the player was radiant (optional)
	includedAccountId := int32(56) // int32 | Account IDs in the match (array) (optional)
	excludedAccountId := int32(56) // int32 | Account IDs not in the match (array) (optional)
	withHeroId := int32(56) // int32 | Hero IDs on the player's team (array) (optional)
	againstHeroId := int32(56) // int32 | Hero IDs against the player's team (array) (optional)
	significant := int32(56) // int32 | Whether the match was significant for aggregation purposes. Defaults to 1 (true), set this to 0 to return data for non-standard modes/matches. (optional)
	having := int32(56) // int32 | The minimum number of games played, for filtering hero stats (optional)
	sort := "sort_example" // string | The field to return matches sorted by in descending order (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectTotals(context.Background(), accountId).Limit(limit).Offset(offset).Win(win).Patch(patch).GameMode(gameMode).LobbyType(lobbyType).Region(region).Date(date).LaneRole(laneRole).HeroId(heroId).IsRadiant(isRadiant).IncludedAccountId(includedAccountId).ExcludedAccountId(excludedAccountId).WithHeroId(withHeroId).AgainstHeroId(againstHeroId).Significant(significant).Having(having).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayersByAccountIdSelectTotals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayersByAccountIdSelectTotals`: []PlayerTotalsResponse
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayersByAccountIdSelectTotals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Steam32 account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayersByAccountIdSelectTotalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of matches to limit to | 
 **offset** | **int32** | Number of matches to offset start by | 
 **win** | **int32** | Whether the player won | 
 **patch** | **int32** | Patch ID, from dotaconstants | 
 **gameMode** | **int32** | Game Mode ID | 
 **lobbyType** | **int32** | Lobby type ID | 
 **region** | **int32** | Region ID | 
 **date** | **int32** | Days previous | 
 **laneRole** | **int32** | Lane Role ID | 
 **heroId** | **int32** | Hero ID | 
 **isRadiant** | **int32** | Whether the player was radiant | 
 **includedAccountId** | **int32** | Account IDs in the match (array) | 
 **excludedAccountId** | **int32** | Account IDs not in the match (array) | 
 **withHeroId** | **int32** | Hero IDs on the player&#39;s team (array) | 
 **againstHeroId** | **int32** | Hero IDs against the player&#39;s team (array) | 
 **significant** | **int32** | Whether the match was significant for aggregation purposes. Defaults to 1 (true), set this to 0 to return data for non-standard modes/matches. | 
 **having** | **int32** | The minimum number of games played, for filtering hero stats | 
 **sort** | **string** | The field to return matches sorted by in descending order | 

### Return type

[**[]PlayerTotalsResponse**](PlayerTotalsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayersByAccountIdSelectWardmap

> PlayerWardMapResponse GetPlayersByAccountIdSelectWardmap(ctx, accountId).Limit(limit).Offset(offset).Win(win).Patch(patch).GameMode(gameMode).LobbyType(lobbyType).Region(region).Date(date).LaneRole(laneRole).HeroId(heroId).IsRadiant(isRadiant).IncludedAccountId(includedAccountId).ExcludedAccountId(excludedAccountId).WithHeroId(withHeroId).AgainstHeroId(againstHeroId).Significant(significant).Having(having).Sort(sort).Execute()

GET /players/{account_id}/wardmap



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
	accountId := int32(56) // int32 | Steam32 account ID
	limit := int32(56) // int32 | Number of matches to limit to (optional)
	offset := int32(56) // int32 | Number of matches to offset start by (optional)
	win := int32(56) // int32 | Whether the player won (optional)
	patch := int32(56) // int32 | Patch ID, from dotaconstants (optional)
	gameMode := int32(56) // int32 | Game Mode ID (optional)
	lobbyType := int32(56) // int32 | Lobby type ID (optional)
	region := int32(56) // int32 | Region ID (optional)
	date := int32(56) // int32 | Days previous (optional)
	laneRole := int32(56) // int32 | Lane Role ID (optional)
	heroId := int32(56) // int32 | Hero ID (optional)
	isRadiant := int32(56) // int32 | Whether the player was radiant (optional)
	includedAccountId := int32(56) // int32 | Account IDs in the match (array) (optional)
	excludedAccountId := int32(56) // int32 | Account IDs not in the match (array) (optional)
	withHeroId := int32(56) // int32 | Hero IDs on the player's team (array) (optional)
	againstHeroId := int32(56) // int32 | Hero IDs against the player's team (array) (optional)
	significant := int32(56) // int32 | Whether the match was significant for aggregation purposes. Defaults to 1 (true), set this to 0 to return data for non-standard modes/matches. (optional)
	having := int32(56) // int32 | The minimum number of games played, for filtering hero stats (optional)
	sort := "sort_example" // string | The field to return matches sorted by in descending order (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectWardmap(context.Background(), accountId).Limit(limit).Offset(offset).Win(win).Patch(patch).GameMode(gameMode).LobbyType(lobbyType).Region(region).Date(date).LaneRole(laneRole).HeroId(heroId).IsRadiant(isRadiant).IncludedAccountId(includedAccountId).ExcludedAccountId(excludedAccountId).WithHeroId(withHeroId).AgainstHeroId(againstHeroId).Significant(significant).Having(having).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayersByAccountIdSelectWardmap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayersByAccountIdSelectWardmap`: PlayerWardMapResponse
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayersByAccountIdSelectWardmap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Steam32 account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayersByAccountIdSelectWardmapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of matches to limit to | 
 **offset** | **int32** | Number of matches to offset start by | 
 **win** | **int32** | Whether the player won | 
 **patch** | **int32** | Patch ID, from dotaconstants | 
 **gameMode** | **int32** | Game Mode ID | 
 **lobbyType** | **int32** | Lobby type ID | 
 **region** | **int32** | Region ID | 
 **date** | **int32** | Days previous | 
 **laneRole** | **int32** | Lane Role ID | 
 **heroId** | **int32** | Hero ID | 
 **isRadiant** | **int32** | Whether the player was radiant | 
 **includedAccountId** | **int32** | Account IDs in the match (array) | 
 **excludedAccountId** | **int32** | Account IDs not in the match (array) | 
 **withHeroId** | **int32** | Hero IDs on the player&#39;s team (array) | 
 **againstHeroId** | **int32** | Hero IDs against the player&#39;s team (array) | 
 **significant** | **int32** | Whether the match was significant for aggregation purposes. Defaults to 1 (true), set this to 0 to return data for non-standard modes/matches. | 
 **having** | **int32** | The minimum number of games played, for filtering hero stats | 
 **sort** | **string** | The field to return matches sorted by in descending order | 

### Return type

[**PlayerWardMapResponse**](PlayerWardMapResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayersByAccountIdSelectWl

> PlayerWinLossResponse GetPlayersByAccountIdSelectWl(ctx, accountId).Limit(limit).Offset(offset).Win(win).Patch(patch).GameMode(gameMode).LobbyType(lobbyType).Region(region).Date(date).LaneRole(laneRole).HeroId(heroId).IsRadiant(isRadiant).IncludedAccountId(includedAccountId).ExcludedAccountId(excludedAccountId).WithHeroId(withHeroId).AgainstHeroId(againstHeroId).Significant(significant).Having(having).Sort(sort).Execute()

GET /players/{account_id}/wl



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
	accountId := int32(56) // int32 | Steam32 account ID
	limit := int32(56) // int32 | Number of matches to limit to (optional)
	offset := int32(56) // int32 | Number of matches to offset start by (optional)
	win := int32(56) // int32 | Whether the player won (optional)
	patch := int32(56) // int32 | Patch ID, from dotaconstants (optional)
	gameMode := int32(56) // int32 | Game Mode ID (optional)
	lobbyType := int32(56) // int32 | Lobby type ID (optional)
	region := int32(56) // int32 | Region ID (optional)
	date := int32(56) // int32 | Days previous (optional)
	laneRole := int32(56) // int32 | Lane Role ID (optional)
	heroId := int32(56) // int32 | Hero ID (optional)
	isRadiant := int32(56) // int32 | Whether the player was radiant (optional)
	includedAccountId := int32(56) // int32 | Account IDs in the match (array) (optional)
	excludedAccountId := int32(56) // int32 | Account IDs not in the match (array) (optional)
	withHeroId := int32(56) // int32 | Hero IDs on the player's team (array) (optional)
	againstHeroId := int32(56) // int32 | Hero IDs against the player's team (array) (optional)
	significant := int32(56) // int32 | Whether the match was significant for aggregation purposes. Defaults to 1 (true), set this to 0 to return data for non-standard modes/matches. (optional)
	having := int32(56) // int32 | The minimum number of games played, for filtering hero stats (optional)
	sort := "sort_example" // string | The field to return matches sorted by in descending order (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectWl(context.Background(), accountId).Limit(limit).Offset(offset).Win(win).Patch(patch).GameMode(gameMode).LobbyType(lobbyType).Region(region).Date(date).LaneRole(laneRole).HeroId(heroId).IsRadiant(isRadiant).IncludedAccountId(includedAccountId).ExcludedAccountId(excludedAccountId).WithHeroId(withHeroId).AgainstHeroId(againstHeroId).Significant(significant).Having(having).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayersByAccountIdSelectWl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayersByAccountIdSelectWl`: PlayerWinLossResponse
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayersByAccountIdSelectWl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Steam32 account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayersByAccountIdSelectWlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of matches to limit to | 
 **offset** | **int32** | Number of matches to offset start by | 
 **win** | **int32** | Whether the player won | 
 **patch** | **int32** | Patch ID, from dotaconstants | 
 **gameMode** | **int32** | Game Mode ID | 
 **lobbyType** | **int32** | Lobby type ID | 
 **region** | **int32** | Region ID | 
 **date** | **int32** | Days previous | 
 **laneRole** | **int32** | Lane Role ID | 
 **heroId** | **int32** | Hero ID | 
 **isRadiant** | **int32** | Whether the player was radiant | 
 **includedAccountId** | **int32** | Account IDs in the match (array) | 
 **excludedAccountId** | **int32** | Account IDs not in the match (array) | 
 **withHeroId** | **int32** | Hero IDs on the player&#39;s team (array) | 
 **againstHeroId** | **int32** | Hero IDs against the player&#39;s team (array) | 
 **significant** | **int32** | Whether the match was significant for aggregation purposes. Defaults to 1 (true), set this to 0 to return data for non-standard modes/matches. | 
 **having** | **int32** | The minimum number of games played, for filtering hero stats | 
 **sort** | **string** | The field to return matches sorted by in descending order | 

### Return type

[**PlayerWinLossResponse**](PlayerWinLossResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayersByAccountIdSelectWordcloud

> PlayerWordCloudResponse GetPlayersByAccountIdSelectWordcloud(ctx, accountId).Limit(limit).Offset(offset).Win(win).Patch(patch).GameMode(gameMode).LobbyType(lobbyType).Region(region).Date(date).LaneRole(laneRole).HeroId(heroId).IsRadiant(isRadiant).IncludedAccountId(includedAccountId).ExcludedAccountId(excludedAccountId).WithHeroId(withHeroId).AgainstHeroId(againstHeroId).Significant(significant).Having(having).Sort(sort).Execute()

GET /players/{account_id}/wordcloud



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
	accountId := int32(56) // int32 | Steam32 account ID
	limit := int32(56) // int32 | Number of matches to limit to (optional)
	offset := int32(56) // int32 | Number of matches to offset start by (optional)
	win := int32(56) // int32 | Whether the player won (optional)
	patch := int32(56) // int32 | Patch ID, from dotaconstants (optional)
	gameMode := int32(56) // int32 | Game Mode ID (optional)
	lobbyType := int32(56) // int32 | Lobby type ID (optional)
	region := int32(56) // int32 | Region ID (optional)
	date := int32(56) // int32 | Days previous (optional)
	laneRole := int32(56) // int32 | Lane Role ID (optional)
	heroId := int32(56) // int32 | Hero ID (optional)
	isRadiant := int32(56) // int32 | Whether the player was radiant (optional)
	includedAccountId := int32(56) // int32 | Account IDs in the match (array) (optional)
	excludedAccountId := int32(56) // int32 | Account IDs not in the match (array) (optional)
	withHeroId := int32(56) // int32 | Hero IDs on the player's team (array) (optional)
	againstHeroId := int32(56) // int32 | Hero IDs against the player's team (array) (optional)
	significant := int32(56) // int32 | Whether the match was significant for aggregation purposes. Defaults to 1 (true), set this to 0 to return data for non-standard modes/matches. (optional)
	having := int32(56) // int32 | The minimum number of games played, for filtering hero stats (optional)
	sort := "sort_example" // string | The field to return matches sorted by in descending order (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectWordcloud(context.Background(), accountId).Limit(limit).Offset(offset).Win(win).Patch(patch).GameMode(gameMode).LobbyType(lobbyType).Region(region).Date(date).LaneRole(laneRole).HeroId(heroId).IsRadiant(isRadiant).IncludedAccountId(includedAccountId).ExcludedAccountId(excludedAccountId).WithHeroId(withHeroId).AgainstHeroId(againstHeroId).Significant(significant).Having(having).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayersByAccountIdSelectWordcloud``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayersByAccountIdSelectWordcloud`: PlayerWordCloudResponse
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayersByAccountIdSelectWordcloud`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Steam32 account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayersByAccountIdSelectWordcloudRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of matches to limit to | 
 **offset** | **int32** | Number of matches to offset start by | 
 **win** | **int32** | Whether the player won | 
 **patch** | **int32** | Patch ID, from dotaconstants | 
 **gameMode** | **int32** | Game Mode ID | 
 **lobbyType** | **int32** | Lobby type ID | 
 **region** | **int32** | Region ID | 
 **date** | **int32** | Days previous | 
 **laneRole** | **int32** | Lane Role ID | 
 **heroId** | **int32** | Hero ID | 
 **isRadiant** | **int32** | Whether the player was radiant | 
 **includedAccountId** | **int32** | Account IDs in the match (array) | 
 **excludedAccountId** | **int32** | Account IDs not in the match (array) | 
 **withHeroId** | **int32** | Hero IDs on the player&#39;s team (array) | 
 **againstHeroId** | **int32** | Hero IDs against the player&#39;s team (array) | 
 **significant** | **int32** | Whether the match was significant for aggregation purposes. Defaults to 1 (true), set this to 0 to return data for non-standard modes/matches. | 
 **having** | **int32** | The minimum number of games played, for filtering hero stats | 
 **sort** | **string** | The field to return matches sorted by in descending order | 

### Return type

[**PlayerWordCloudResponse**](PlayerWordCloudResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRefresh

> map[string]interface{} PostRefresh(ctx, accountId).Execute()

POST /players/{account_id}/refresh



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
	accountId := int32(56) // int32 | Steam32 account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.PostRefresh(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.PostRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRefresh`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.PostRefresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Steam32 account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

