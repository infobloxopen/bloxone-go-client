/*
DNS Configuration API

Testing ViewAPIService

*/

package dns_config_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	openapiclient "github.com/infobloxopen/bloxone-go-client/dns_config"
	"github.com/infobloxopen/bloxone-go-client/internal"
)

var ConfigBulkCopyViewPost = openapiclient.ConfigBulkCopyView{
	AuthZoneConfig:      nil,
	ForwardZoneConfig:   nil,
	Recursive:           nil,
	Resources:           nil,
	SecondaryZoneConfig: nil,
	SkipOnError:         nil,
	Target:              "TargetPost",
}
var ConfigViewPost = openapiclient.ConfigView{
	Id:   openapiclient.PtrString("ConfigViewPost"),
	Tags: make(map[string]interface{}),
}
var ConfigViewPatch = openapiclient.ConfigView{
	Id:   openapiclient.PtrString("ConfigViewPatch"),
	Tags: make(map[string]interface{}),
}

func TestViewAPIService(t *testing.T) {

	t.Run("Test ViewAPIService ViewBulkCopy", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPost, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/view/bulk_copy", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))

			var reqBody openapiclient.ConfigBulkCopyView
			require.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			require.Equal(t, ConfigBulkCopyViewPost, reqBody)

			response := openapiclient.ConfigBulkCopyResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.ViewAPI.ViewBulkCopy(context.Background()).Body(ConfigBulkCopyViewPost).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test ViewAPIService ViewCreate", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPost, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/view", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))

			var reqBody openapiclient.ConfigView
			require.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			require.Equal(t, ConfigViewPost, reqBody)

			response := openapiclient.ConfigCreateViewResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.ViewAPI.ViewCreate(context.Background()).Body(ConfigViewPost).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test ViewAPIService ViewDelete", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodDelete, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/view/"+*ConfigViewPost.Id, req.URL.Path)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader([]byte{})),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		httpRes, err := apiClient.ViewAPI.ViewDelete(context.Background(), *ConfigViewPost.Id).Execute()
		require.Nil(t, err)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test ViewAPIService ViewList", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodGet, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/view", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			response := openapiclient.ConfigListViewResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.ViewAPI.ViewList(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test ViewAPIService ViewRead", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodGet, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/view/"+*ConfigViewPost.Id, req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			response := openapiclient.ConfigReadViewResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.ViewAPI.ViewRead(context.Background(), *ConfigViewPost.Id).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test ViewAPIService ViewUpdate", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPatch, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/view/"+*ConfigViewPatch.Id, req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))

			var reqBody openapiclient.ConfigView
			require.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			require.Equal(t, ConfigViewPatch, reqBody)

			response := openapiclient.ConfigUpdateViewResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.ViewAPI.ViewUpdate(context.Background(), *ConfigViewPatch.Id).Body(ConfigViewPatch).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

}
