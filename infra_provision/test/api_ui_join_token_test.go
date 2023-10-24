/*
Host Activation Service

Testing UIJoinTokenAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package infra_provision

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	openapiclient "github.com/infobloxopen/bloxone-go-client/infra_provision"
	"github.com/infobloxopen/bloxone-go-client/internal"
)

func Test_infra_provision_UIJoinTokenAPIService(t *testing.T) {

	configuration := internal.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UIJoinTokenAPIService UIJoinTokenCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UIJoinTokenAPI.UIJoinTokenCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UIJoinTokenAPIService UIJoinTokenDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.UIJoinTokenAPI.UIJoinTokenDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UIJoinTokenAPIService UIJoinTokenDeleteSet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.UIJoinTokenAPI.UIJoinTokenDeleteSet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UIJoinTokenAPIService UIJoinTokenList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UIJoinTokenAPI.UIJoinTokenList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UIJoinTokenAPIService UIJoinTokenRead", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.UIJoinTokenAPI.UIJoinTokenRead(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UIJoinTokenAPIService UIJoinTokenUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.UIJoinTokenAPI.UIJoinTokenUpdate(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
