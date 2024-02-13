# \ProPlayersAPI

All URIs are relative to *https://api.opendota.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProPlayers**](ProPlayersAPI.md#GetProPlayers) | **Get** /proPlayers | GET /proPlayers



## GetProPlayers

> []PlayerObjectResponse GetProPlayers(ctx).Execute()

GET /proPlayers



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
	resp, r, err := apiClient.ProPlayersAPI.GetProPlayers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProPlayersAPI.GetProPlayers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProPlayers`: []PlayerObjectResponse
	fmt.Fprintf(os.Stdout, "Response from `ProPlayersAPI.GetProPlayers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProPlayersRequest struct via the builder pattern


### Return type

[**[]PlayerObjectResponse**](PlayerObjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

