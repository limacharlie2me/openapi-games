# \ParsedMatchesAPI

All URIs are relative to *https://api.opendota.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetParsedMatches**](ParsedMatchesAPI.md#GetParsedMatches) | **Get** /parsedMatches | GET /parsedMatches



## GetParsedMatches

> []ParsedMatchesResponse GetParsedMatches(ctx).LessThanMatchId(lessThanMatchId).Execute()

GET /parsedMatches



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
	resp, r, err := apiClient.ParsedMatchesAPI.GetParsedMatches(context.Background()).LessThanMatchId(lessThanMatchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParsedMatchesAPI.GetParsedMatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParsedMatches`: []ParsedMatchesResponse
	fmt.Fprintf(os.Stdout, "Response from `ParsedMatchesAPI.GetParsedMatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetParsedMatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lessThanMatchId** | **int32** | Get matches with a match ID lower than this value | 

### Return type

[**[]ParsedMatchesResponse**](ParsedMatchesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

