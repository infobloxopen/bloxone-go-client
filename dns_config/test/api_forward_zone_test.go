/*
DNS Configuration API

Testing ForwardZoneAPIService

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

func TestForwardZoneAPIService(t *testing.T) {

	apiClient := dns_config.NewAPIClient()

	t.Run("Test ForwardZoneAPIService ForwardZoneCopy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ForwardZoneAPI.ForwardZoneCopy(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForwardZoneAPIService ForwardZoneCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ForwardZoneAPI.ForwardZoneCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForwardZoneAPIService ForwardZoneDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.ForwardZoneAPI.ForwardZoneDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForwardZoneAPIService ForwardZoneList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ForwardZoneAPI.ForwardZoneList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForwardZoneAPIService ForwardZoneRead", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.ForwardZoneAPI.ForwardZoneRead(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForwardZoneAPIService ForwardZoneUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.ForwardZoneAPI.ForwardZoneUpdate(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
