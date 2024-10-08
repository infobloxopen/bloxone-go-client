/*
Discovery Configuration API V2

Testing AccountsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package clouddiscovery

import (
	"context"
	"github.com/infobloxopen/bloxone-go-client/client"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountsAPIService(t *testing.T) {

	apiClient := client.NewAPIClient()

	t.Run("Test AccountsAPIService List", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DiscoveryConfigurationAPIV2.AccountsAPI.List(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
