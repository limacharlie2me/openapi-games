# \MatchesAPI

All URIs are relative to *https://api.opendota.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMatchesByMatchId**](MatchesAPI.md#GetMatchesByMatchId) | **Get** /matches/{match_id} | GET /matches/{match_id}



## GetMatchesByMatchId

> MatchResponse GetMatchesByMatchId(ctx, matchId).Execute()

GET /matches/{match_id}



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
	matchId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MatchesAPI.GetMatchesByMatchId(context.Background(), matchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MatchesAPI.GetMatchesByMatchId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMatchesByMatchId`: MatchResponse
	fmt.Fprintf(os.Stdout, "Response from `MatchesAPI.GetMatchesByMatchId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**matchId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMatchesByMatchIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MatchResponse**](MatchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

