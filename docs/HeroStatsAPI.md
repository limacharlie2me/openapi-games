# \HeroStatsAPI

All URIs are relative to *https://api.opendota.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHeroStats**](HeroStatsAPI.md#GetHeroStats) | **Get** /heroStats | GET /heroStats



## GetHeroStats

> []HeroStatsResponse GetHeroStats(ctx).Execute()

GET /heroStats



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
	resp, r, err := apiClient.HeroStatsAPI.GetHeroStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HeroStatsAPI.GetHeroStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHeroStats`: []HeroStatsResponse
	fmt.Fprintf(os.Stdout, "Response from `HeroStatsAPI.GetHeroStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHeroStatsRequest struct via the builder pattern


### Return type

[**[]HeroStatsResponse**](HeroStatsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

