# \ExplorerAPI

All URIs are relative to *https://api.opendota.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExplorer**](ExplorerAPI.md#GetExplorer) | **Get** /explorer | GET /explorer



## GetExplorer

> map[string]interface{} GetExplorer(ctx).Sql(sql).Execute()

GET /explorer



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
	sql := "sql_example" // string | The PostgreSQL query as percent-encoded string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExplorerAPI.GetExplorer(context.Background()).Sql(sql).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExplorerAPI.GetExplorer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExplorer`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ExplorerAPI.GetExplorer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExplorerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sql** | **string** | The PostgreSQL query as percent-encoded string. | 

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

