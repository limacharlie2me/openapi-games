# \ScenariosAPI

All URIs are relative to *https://api.opendota.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetScenariosItemTimings**](ScenariosAPI.md#GetScenariosItemTimings) | **Get** /scenarios/itemTimings | GET /scenarios/itemTimings
[**GetScenariosLaneRoles**](ScenariosAPI.md#GetScenariosLaneRoles) | **Get** /scenarios/laneRoles | GET /scenarios/laneRoles
[**GetScenariosMisc**](ScenariosAPI.md#GetScenariosMisc) | **Get** /scenarios/misc | GET /scenarios/misc



## GetScenariosItemTimings

> []ScenarioItemTimingsResponse GetScenariosItemTimings(ctx).Item(item).HeroId(heroId).Execute()

GET /scenarios/itemTimings



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
	item := "item_example" // string | Filter by item name e.g. \"spirit_vessel\" (optional)
	heroId := int32(56) // int32 | Hero ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScenariosAPI.GetScenariosItemTimings(context.Background()).Item(item).HeroId(heroId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScenariosAPI.GetScenariosItemTimings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScenariosItemTimings`: []ScenarioItemTimingsResponse
	fmt.Fprintf(os.Stdout, "Response from `ScenariosAPI.GetScenariosItemTimings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetScenariosItemTimingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **item** | **string** | Filter by item name e.g. \&quot;spirit_vessel\&quot; | 
 **heroId** | **int32** | Hero ID | 

### Return type

[**[]ScenarioItemTimingsResponse**](ScenarioItemTimingsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScenariosLaneRoles

> []ScenarioLaneRolesResponse GetScenariosLaneRoles(ctx).LaneRole(laneRole).HeroId(heroId).Execute()

GET /scenarios/laneRoles



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
	laneRole := "laneRole_example" // string | Filter by lane role 1-4 (Safe, Mid, Off, Jungle) (optional)
	heroId := int32(56) // int32 | Hero ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScenariosAPI.GetScenariosLaneRoles(context.Background()).LaneRole(laneRole).HeroId(heroId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScenariosAPI.GetScenariosLaneRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScenariosLaneRoles`: []ScenarioLaneRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `ScenariosAPI.GetScenariosLaneRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetScenariosLaneRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **laneRole** | **string** | Filter by lane role 1-4 (Safe, Mid, Off, Jungle) | 
 **heroId** | **int32** | Hero ID | 

### Return type

[**[]ScenarioLaneRolesResponse**](ScenarioLaneRolesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScenariosMisc

> []ScenarioMiscResponse GetScenariosMisc(ctx).Scenario(scenario).Execute()

GET /scenarios/misc



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
	scenario := "scenario_example" // string | Name of the scenario (see teamScenariosQueryParams) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScenariosAPI.GetScenariosMisc(context.Background()).Scenario(scenario).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScenariosAPI.GetScenariosMisc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScenariosMisc`: []ScenarioMiscResponse
	fmt.Fprintf(os.Stdout, "Response from `ScenariosAPI.GetScenariosMisc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetScenariosMiscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scenario** | **string** | Name of the scenario (see teamScenariosQueryParams) | 

### Return type

[**[]ScenarioMiscResponse**](ScenarioMiscResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

