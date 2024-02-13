# \ConstantsAPI

All URIs are relative to *https://api.opendota.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConstants**](ConstantsAPI.md#GetConstants) | **Get** /constants | GET /constants
[**GetConstantsByResource**](ConstantsAPI.md#GetConstantsByResource) | **Get** /constants/{resource} | GET /constants



## GetConstants

> []string GetConstants(ctx).Execute()

GET /constants



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
	resp, r, err := apiClient.ConstantsAPI.GetConstants(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsAPI.GetConstants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConstants`: []string
	fmt.Fprintf(os.Stdout, "Response from `ConstantsAPI.GetConstants`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConstantsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConstantsByResource

> GetConstantsByResource200Response GetConstantsByResource(ctx, resource).Execute()

GET /constants



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
	resource := "resource_example" // string | Resource name e.g. `heroes`. [List of resources](https://github.com/odota/dotaconstants/tree/master/build)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstantsAPI.GetConstantsByResource(context.Background(), resource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsAPI.GetConstantsByResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConstantsByResource`: GetConstantsByResource200Response
	fmt.Fprintf(os.Stdout, "Response from `ConstantsAPI.GetConstantsByResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resource** | **string** | Resource name e.g. &#x60;heroes&#x60;. [List of resources](https://github.com/odota/dotaconstants/tree/master/build) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConstantsByResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetConstantsByResource200Response**](GetConstantsByResource200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

