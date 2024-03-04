/*
DNS Configuration API

Testing AuthZoneAPIService

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

var ConfigCopyAuthZonePost = openapiclient.ConfigCopyAuthZone{
	Comment: openapiclient.PtrString("This is a copyping AuthZone for testing."),
	Id:      openapiclient.PtrString("dummyAuthZoneId"),
}
var ConfigAuthZonePost = openapiclient.ConfigAuthZone{
	Comment: openapiclient.PtrString("This is a dummy AuthZone for testing."),
	Id:      openapiclient.PtrString("dummyAuthZoneId"),
	Tags:    make(map[string]interface{}),
}
var ConfigAuthZonePatch = openapiclient.ConfigAuthZone{
	Comment: openapiclient.PtrString("This is an updated dummy AuthZone for testing."),
	Id:      ConfigAuthZonePost.Id,
	Tags:    make(map[string]interface{}),
}

func TestAuthZoneAPIService(t *testing.T) {

	t.Run("Test AuthZoneAPIService AuthZoneCopy", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPost, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/auth_zone/copy", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))

			var reqBody openapiclient.ConfigCopyAuthZone
			require.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			require.Equal(t, ConfigCopyAuthZonePost, reqBody)

			response := openapiclient.ConfigCopyAuthZoneResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.AuthZoneAPI.AuthZoneCopy(context.Background()).Body(ConfigCopyAuthZonePost).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test AuthZoneAPIService AuthZoneCreate", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPost, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/auth_zone", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))

			var reqBody openapiclient.ConfigAuthZone
			require.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			require.Equal(t, ConfigAuthZonePost, reqBody)

			response := openapiclient.ConfigCreateAuthZoneResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.AuthZoneAPI.AuthZoneCreate(context.Background()).Body(ConfigAuthZonePost).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test AuthZoneAPIService AuthZoneDelete", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodDelete, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/auth_zone/"+*ConfigAuthZonePost.Id, req.URL.Path)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader([]byte{})),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		httpRes, err := apiClient.AuthZoneAPI.AuthZoneDelete(context.Background(), *ConfigAuthZonePost.Id).Execute()
		require.Nil(t, err)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test AuthZoneAPIService AuthZoneList", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodGet, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/auth_zone", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			response := openapiclient.ConfigListAuthZoneResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.AuthZoneAPI.AuthZoneList(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test AuthZoneAPIService AuthZoneRead", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodGet, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/auth_zone/"+*ConfigAuthZonePost.Id, req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			response := openapiclient.ConfigReadAuthZoneResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.AuthZoneAPI.AuthZoneRead(context.Background(), *ConfigAuthZonePost.Id).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test AuthZoneAPIService AuthZoneUpdate", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPatch, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/auth_zone/"+*ConfigAuthZonePost.Id, req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))

			var reqBody openapiclient.ConfigAuthZone
			require.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			require.Equal(t, ConfigAuthZonePatch, reqBody)

			response := openapiclient.ConfigUpdateAuthZoneResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.AuthZoneAPI.AuthZoneUpdate(context.Background(), *ConfigAuthZonePatch.Id).Body(ConfigAuthZonePatch).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

}
