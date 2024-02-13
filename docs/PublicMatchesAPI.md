# \PublicMatchesAPI

All URIs are relative to *https://api.opendota.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPublicMatches**](PublicMatchesAPI.md#GetPublicMatches) | **Get** /publicMatches | GET /publicMatches



## GetPublicMatches

> []PublicMatchesResponse GetPublicMatches(ctx).LessThanMatchId(lessThanMatchId).MinRank(minRank).MaxRank(maxRank).MmrAscending(mmrAscending).MmrDescending(mmrDescending).Execute()

GET /publicMatches



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
	minRank := int32(56) // int32 | Minimum rank for the matches. Ranks are represented by integers (10-15: Herald, 20-25: Guardian, 30-35: Crusader, 40-45: Archon, 50-55: Legend, 60-65: Ancient, 70-75: Divine, 80-85: Immortal). Each increment represents an additional star. (optional)
	maxRank := int32(56) // int32 | Maximum rank for the matches. Ranks are represented by integers (10-15: Herald, 20-25: Guardian, 30-35: Crusader, 40-45: Archon, 50-55: Legend, 60-65: Ancient, 70-75: Divine, 80-85: Immortal). Each increment represents an additional star. (optional)
	mmrAscending := int32(56) // int32 | Order by average rank ascending (optional)
	mmrDescending := int32(56) // int32 | Order by average rank descending (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicMatchesAPI.GetPublicMatches(context.Background()).LessThanMatchId(lessThanMatchId).MinRank(minRank).MaxRank(maxRank).MmrAscending(mmrAscending).MmrDescending(mmrDescending).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicMatchesAPI.GetPublicMatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicMatches`: []PublicMatchesResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicMatchesAPI.GetPublicMatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicMatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lessThanMatchId** | **int32** | Get matches with a match ID lower than this value | 
 **minRank** | **int32** | Minimum rank for the matches. Ranks are represented by integers (10-15: Herald, 20-25: Guardian, 30-35: Crusader, 40-45: Archon, 50-55: Legend, 60-65: Ancient, 70-75: Divine, 80-85: Immortal). Each increment represents an additional star. | 
 **maxRank** | **int32** | Maximum rank for the matches. Ranks are represented by integers (10-15: Herald, 20-25: Guardian, 30-35: Crusader, 40-45: Archon, 50-55: Legend, 60-65: Ancient, 70-75: Divine, 80-85: Immortal). Each increment represents an additional star. | 
 **mmrAscending** | **int32** | Order by average rank ascending | 
 **mmrDescending** | **int32** | Order by average rank descending | 

### Return type

[**[]PublicMatchesResponse**](PublicMatchesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

