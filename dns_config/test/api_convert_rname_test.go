/*
DNS Configuration API

Testing ConvertRnameAPIService

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

func TestConvertRnameAPIService(t *testing.T) {

	apiClient := dns_config.NewAPIClient()

	t.Run("Test ConvertRnameAPIService ConvertRnameConvertRName", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var emailAddress string

		resp, httpRes, err := apiClient.ConvertRnameAPI.ConvertRnameConvertRName(context.Background(), emailAddress).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
