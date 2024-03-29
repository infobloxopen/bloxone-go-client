/*
Infrastructure Management API

Testing DetailAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package infra_mgmt

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	openapiclient "github.com/infobloxopen/bloxone-go-client/infra_mgmt"
	"github.com/infobloxopen/bloxone-go-client/internal"
)

func Test_infra_mgmt_DetailAPIService(t *testing.T) {

	configuration := internal.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DetailAPIService DetailHostsList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DetailAPI.DetailHostsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DetailAPIService DetailServicesList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DetailAPI.DetailServicesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
