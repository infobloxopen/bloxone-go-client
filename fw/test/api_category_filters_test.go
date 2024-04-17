/*
BloxOne FW API

Testing CategoryFiltersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package fw

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/infobloxopen/bloxone-go-client/fw"
)

func TestCategoryFiltersAPIService(t *testing.T) {

	apiClient := fw.NewAPIClient()

	t.Run("Test CategoryFiltersAPIService CategoryFiltersCreateCategoryFilter", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CategoryFiltersAPI.CategoryFiltersCreateCategoryFilter(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CategoryFiltersAPIService CategoryFiltersDeleteCategoryFilters", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.CategoryFiltersAPI.CategoryFiltersDeleteCategoryFilters(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CategoryFiltersAPIService CategoryFiltersDeleteSingleCategoryFilters", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		httpRes, err := apiClient.CategoryFiltersAPI.CategoryFiltersDeleteSingleCategoryFilters(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CategoryFiltersAPIService CategoryFiltersListCategoryFilters", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CategoryFiltersAPI.CategoryFiltersListCategoryFilters(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CategoryFiltersAPIService CategoryFiltersReadCategoryFilter", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.CategoryFiltersAPI.CategoryFiltersReadCategoryFilter(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CategoryFiltersAPIService CategoryFiltersUpdateCategoryFilter", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.CategoryFiltersAPI.CategoryFiltersUpdateCategoryFilter(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
