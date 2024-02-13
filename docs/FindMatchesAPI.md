# \FindMatchesAPI

All URIs are relative to *https://api.opendota.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFindMatches**](FindMatchesAPI.md#GetFindMatches) | **Get** /findMatches | GET /



## GetFindMatches

> []map[string]interface{} GetFindMatches(ctx).TeamA(teamA).TeamB(teamB).Execute()

GET /



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
	teamA := []int32{int32(123)} // []int32 | Hero IDs on first team (array) (optional)
	teamB := []int32{int32(123)} // []int32 | Hero IDs on second team (array) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindMatchesAPI.GetFindMatches(context.Background()).TeamA(teamA).TeamB(teamB).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindMatchesAPI.GetFindMatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFindMatches`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FindMatchesAPI.GetFindMatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFindMatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamA** | **[]int32** | Hero IDs on first team (array) | 
 **teamB** | **[]int32** | Hero IDs on second team (array) | 

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

