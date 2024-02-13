# \RequestAPI

All URIs are relative to *https://api.opendota.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRequestByJobId**](RequestAPI.md#GetRequestByJobId) | **Get** /request/{jobId} | GET /request/{jobId}
[**PostRequestByJobId**](RequestAPI.md#PostRequestByJobId) | **Post** /request/{match_id} | POST /request/{match_id}



## GetRequestByJobId

> map[string]interface{} GetRequestByJobId(ctx, jobId).Execute()

GET /request/{jobId}



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
	jobId := "jobId_example" // string | The job ID to query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestAPI.GetRequestByJobId(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.GetRequestByJobId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequestByJobId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RequestAPI.GetRequestByJobId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The job ID to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestByJobIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PostRequestByJobId

> map[string]interface{} PostRequestByJobId(ctx, matchId).Execute()

POST /request/{match_id}



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
	resp, r, err := apiClient.RequestAPI.PostRequestByJobId(context.Background(), matchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.PostRequestByJobId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRequestByJobId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RequestAPI.PostRequestByJobId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**matchId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRequestByJobIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

