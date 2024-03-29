/*
OpenDota API

Testing PlayersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_PlayersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PlayersAPIService GetPlayersByAccountId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId int32

		resp, httpRes, err := apiClient.PlayersAPI.GetPlayersByAccountId(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlayersAPIService GetPlayersByAccountIdHistogramsByField", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId int32
		var field string

		resp, httpRes, err := apiClient.PlayersAPI.GetPlayersByAccountIdHistogramsByField(context.Background(), accountId, field).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlayersAPIService GetPlayersByAccountIdSelectCounts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId int32

		resp, httpRes, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectCounts(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlayersAPIService GetPlayersByAccountIdSelectHeroes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId int32

		resp, httpRes, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectHeroes(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlayersAPIService GetPlayersByAccountIdSelectMatches", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId int32

		resp, httpRes, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectMatches(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlayersAPIService GetPlayersByAccountIdSelectPeers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId int32

		resp, httpRes, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectPeers(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlayersAPIService GetPlayersByAccountIdSelectPros", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId int32

		resp, httpRes, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectPros(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlayersAPIService GetPlayersByAccountIdSelectRankings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId int32

		resp, httpRes, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectRankings(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlayersAPIService GetPlayersByAccountIdSelectRatings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId int32

		resp, httpRes, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectRatings(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlayersAPIService GetPlayersByAccountIdSelectRecentMatches", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId int32

		resp, httpRes, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectRecentMatches(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlayersAPIService GetPlayersByAccountIdSelectTotals", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId int32

		resp, httpRes, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectTotals(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlayersAPIService GetPlayersByAccountIdSelectWardmap", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId int32

		resp, httpRes, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectWardmap(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlayersAPIService GetPlayersByAccountIdSelectWl", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId int32

		resp, httpRes, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectWl(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlayersAPIService GetPlayersByAccountIdSelectWordcloud", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId int32

		resp, httpRes, err := apiClient.PlayersAPI.GetPlayersByAccountIdSelectWordcloud(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlayersAPIService PostRefresh", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId int32

		resp, httpRes, err := apiClient.PlayersAPI.PostRefresh(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
