/*
DNS Configuration API

Testing CacheFlushAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package dns_config

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/infobloxopen/bloxone-go-client/dns_config"
)

func TestCacheFlushAPIService(t *testing.T) {

	apiClient := dns_config.NewAPIClient()

	t.Run("Test CacheFlushAPIService CacheFlushCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CacheFlushAPI.CacheFlushCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
