/*
DDI Keys API

Testing GenerateTsigAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package keys

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/infobloxopen/bloxone-go-client/internal"
	openapiclient "github.com/infobloxopen/bloxone-go-client/keys"
)

func Test_keys_GenerateTsigAPIService(t *testing.T) {

	configuration := internal.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GenerateTsigAPIService GenerateTsigGenerateTSIG", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GenerateTsigAPI.GenerateTsigGenerateTSIG(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}