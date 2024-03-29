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


// PublicMatchesAPIService PublicMatchesAPI service
type PublicMatchesAPIService service

type ApiGetPublicMatchesRequest struct {
	ctx context.Context
	ApiService *PublicMatchesAPIService
	lessThanMatchId *int32
	minRank *int32
	maxRank *int32
	mmrAscending *int32
	mmrDescending *int32
}

// Get matches with a match ID lower than this value
func (r ApiGetPublicMatchesRequest) LessThanMatchId(lessThanMatchId int32) ApiGetPublicMatchesRequest {
	r.lessThanMatchId = &lessThanMatchId
	return r
}

// Minimum rank for the matches. Ranks are represented by integers (10-15: Herald, 20-25: Guardian, 30-35: Crusader, 40-45: Archon, 50-55: Legend, 60-65: Ancient, 70-75: Divine, 80-85: Immortal). Each increment represents an additional star.
func (r ApiGetPublicMatchesRequest) MinRank(minRank int32) ApiGetPublicMatchesRequest {
	r.minRank = &minRank
	return r
}

// Maximum rank for the matches. Ranks are represented by integers (10-15: Herald, 20-25: Guardian, 30-35: Crusader, 40-45: Archon, 50-55: Legend, 60-65: Ancient, 70-75: Divine, 80-85: Immortal). Each increment represents an additional star.
func (r ApiGetPublicMatchesRequest) MaxRank(maxRank int32) ApiGetPublicMatchesRequest {
	r.maxRank = &maxRank
	return r
}

// Order by average rank ascending
func (r ApiGetPublicMatchesRequest) MmrAscending(mmrAscending int32) ApiGetPublicMatchesRequest {
	r.mmrAscending = &mmrAscending
	return r
}

// Order by average rank descending
func (r ApiGetPublicMatchesRequest) MmrDescending(mmrDescending int32) ApiGetPublicMatchesRequest {
	r.mmrDescending = &mmrDescending
	return r
}

func (r ApiGetPublicMatchesRequest) Execute() ([]PublicMatchesResponse, *http.Response, error) {
	return r.ApiService.GetPublicMatchesExecute(r)
}

/*
GetPublicMatches GET /publicMatches

Get list of randomly sampled public matches

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPublicMatchesRequest
*/
func (a *PublicMatchesAPIService) GetPublicMatches(ctx context.Context) ApiGetPublicMatchesRequest {
	return ApiGetPublicMatchesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []PublicMatchesResponse
func (a *PublicMatchesAPIService) GetPublicMatchesExecute(r ApiGetPublicMatchesRequest) ([]PublicMatchesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []PublicMatchesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicMatchesAPIService.GetPublicMatches")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/publicMatches"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.lessThanMatchId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "less_than_match_id", r.lessThanMatchId, "")
	}
	if r.minRank != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "min_rank", r.minRank, "")
	}
	if r.maxRank != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max_rank", r.maxRank, "")
	}
	if r.mmrAscending != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mmr_ascending", r.mmrAscending, "")
	}
	if r.mmrDescending != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mmr_descending", r.mmrDescending, "")
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
