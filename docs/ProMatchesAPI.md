# \ProMatchesAPI

All URIs are relative to *https://api.opendota.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProMatches**](ProMatchesAPI.md#GetProMatches) | **Get** /proMatches | GET /proMatches



## GetProMatches

> []MatchObjectResponse GetProMatches(ctx).LessThanMatchId(lessThanMatchId).Execute()

GET /proMatches



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
	lessThanMatchId := int32(56) // int32 | Get matches with a match ID lower than this value (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProMatchesAPI.GetProMatches(context.Background()).LessThanMatchId(lessThanMatchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProMatchesAPI.GetProMatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProMatches`: []MatchObjectResponse
	fmt.Fprintf(os.Stdout, "Response from `ProMatchesAPI.GetProMatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProMatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lessThanMatchId** | **int32** | Get matches with a match ID lower than this value | 

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

