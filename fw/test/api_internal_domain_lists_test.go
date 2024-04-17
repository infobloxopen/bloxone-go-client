/*
BloxOne FW API

Testing InternalDomainListsAPIService

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

func TestInternalDomainListsAPIService(t *testing.T) {

	apiClient := fw.NewAPIClient()

	t.Run("Test InternalDomainListsAPIService InternalDomainListsCreateInternalDomains", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.InternalDomainListsAPI.InternalDomainListsCreateInternalDomains(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InternalDomainListsAPIService InternalDomainListsDeleteInternalDomains", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.InternalDomainListsAPI.InternalDomainListsDeleteInternalDomains(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InternalDomainListsAPIService InternalDomainListsDeleteSingleInternalDomains", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		httpRes, err := apiClient.InternalDomainListsAPI.InternalDomainListsDeleteSingleInternalDomains(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InternalDomainListsAPIService InternalDomainListsInternalDomainsItemsPartialUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.InternalDomainListsAPI.InternalDomainListsInternalDomainsItemsPartialUpdate(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InternalDomainListsAPIService InternalDomainListsListInternalDomains", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.InternalDomainListsAPI.InternalDomainListsListInternalDomains(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InternalDomainListsAPIService InternalDomainListsReadInternalDomains", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.InternalDomainListsAPI.InternalDomainListsReadInternalDomains(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InternalDomainListsAPIService InternalDomainListsUpdateInternalDomains", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.InternalDomainListsAPI.InternalDomainListsUpdateInternalDomains(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
