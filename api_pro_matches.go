/*
OpenDota API

# Introduction The OpenDota API provides Dota 2 related data including advanced match data extracted from match replays.  You can find data that can be used to convert hero and ability IDs and other information provided by the API from the [dotaconstants](https://github.com/odota/dotaconstants) repository.  Without a key, you can make 2,000 free calls per day at a rate limit of 60 requests/minute. We also offer a Premium Tier with unlimited API calls and higher rate limits. Check out the [API page](https://www.opendota.com/api-keys) to learn more.     

API version: 25.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ProMatchesAPIService ProMatchesAPI service
type ProMatchesAPIService service

type ApiGetProMatchesRequest struct {
	ctx context.Context
	ApiService *ProMatchesAPIService
	lessThanMatchId *int32
}

// Get matches with a match ID lower than this value
func (r ApiGetProMatchesRequest) LessThanMatchId(lessThanMatchId int32) ApiGetProMatchesRequest {
	r.lessThanMatchId = &lessThanMatchId
	return r
}

func (r ApiGetProMatchesRequest) Execute() ([]MatchObjectResponse, *http.Response, error) {
	return r.ApiService.GetProMatchesExecute(r)
}

/*
GetProMatches GET /proMatches

Get list of pro matches

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetProMatchesRequest
*/
func (a *ProMatchesAPIService) GetProMatches(ctx context.Context) ApiGetProMatchesRequest {
	return ApiGetProMatchesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []MatchObjectResponse
func (a *ProMatchesAPIService) GetProMatchesExecute(r ApiGetProMatchesRequest) ([]MatchObjectResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []MatchObjectResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProMatchesAPIService.GetProMatches")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/proMatches"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.lessThanMatchId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "less_than_match_id", r.lessThanMatchId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; charset=utf-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
