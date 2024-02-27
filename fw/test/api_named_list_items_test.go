/*
BloxOne FW API

Testing NamedListItemsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package fw

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	openapiclient "github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/bloxone-go-client/internal"
)

func Test_fw_NamedListItemsAPIService(t *testing.T) {

	configuration := internal.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NamedListItemsAPIService NamedListItemsDeleteNamedListItems", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		httpRes, err := apiClient.NamedListItemsAPI.NamedListItemsDeleteNamedListItems(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamedListItemsAPIService NamedListItemsInsertOrReplaceNamedListItems", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.NamedListItemsAPI.NamedListItemsInsertOrReplaceNamedListItems(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamedListItemsAPIService NamedListItemsNamedListItemsPartialUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.NamedListItemsAPI.NamedListItemsNamedListItemsPartialUpdate(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
