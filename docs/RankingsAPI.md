# \RankingsAPI

All URIs are relative to *https://api.opendota.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRankings**](RankingsAPI.md#GetRankings) | **Get** /rankings | GET /rankings



## GetRankings

> RankingsResponse GetRankings(ctx).HeroId(heroId).Execute()

GET /rankings



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
	heroId := "heroId_example" // string | Hero ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RankingsAPI.GetRankings(context.Background()).HeroId(heroId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RankingsAPI.GetRankings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRankings`: RankingsResponse
	fmt.Fprintf(os.Stdout, "Response from `RankingsAPI.GetRankings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRankingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **heroId** | **string** | Hero ID | 

### Return type

[**RankingsResponse**](RankingsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

