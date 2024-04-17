/*
DDI Keys API

Testing UploadAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package keys

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/infobloxopen/bloxone-go-client/keys"
)

func TestUploadAPIService(t *testing.T) {

	apiClient := keys.NewAPIClient()

	t.Run("Test UploadAPIService UploadUpload", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UploadAPI.UploadUpload(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
